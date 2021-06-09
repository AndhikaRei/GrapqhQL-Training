package handlers

import (
	"context"
	"encoding/base64"
	"fmt"
	"ktp-fix/configs/database"
	"ktp-fix/graph/model"
	"ktp-fix/internal/helper"
	"ktp-fix/internal/models"
	"strconv"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
)

// CreateKtpHandler is a function that create new Ktp in database
// This function will return the models of the created ktp and error
func CreateKtpHandler(ctx context.Context, input models.NewKtp) (*models.Ktp, error) {

	// date parse + generate unique id
	tanggal_lahir, err := time.Parse("2006-01-02", input.TanggalLahir)
	if err != nil {
		return nil, err
	}
	tanggal_lahir = tanggal_lahir.In(time.Local)

	// Assignment + saving in database
	ktp := models.Ktp{
		Nik:          input.Nik,
		Nama:         input.Nama,
		Agama:        input.Agama,
		JenisKelamin: input.JenisKelamin,
		TanggalLahir: tanggal_lahir,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	rs := database.DB.Create(&ktp)
	if rs.Error != nil {
		return nil, rs.Error
	}

	return &ktp, nil
}

// GetKtpHandler is a function that get Ktp with specific id
// This function will return models of the Ktp ( if not error ) and error ( if not exist )
func GetKtpHandler(ctx context.Context, id string) (*models.Ktp, error) {

	selectedKtp := &models.Ktp{}
	intId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	rs := database.DB.First(selectedKtp, "id = ?", intId)
	if rs.Error != nil {
		return nil, rs.Error
	}

	return selectedKtp, nil
}

// GetAllKtp is a function that get all ktp in database
// This function will return array of pointer to Ktp
func GetAllKtp(ctx context.Context) ([]*models.Ktp, error) {

	allUser := []*models.Ktp{}
	rs := database.DB.Find(&allUser)

	if rs.Error != nil {
		return nil, rs.Error
	}

	return allUser, nil
}

// UpdateKtp will update ktp with specific Id with new data
// This function will return boolean (success/not) and error (if exist)
func UpdateKtp(ctx context.Context, id string, input *models.NewKtp) (bool, error) {

	// Date parsing
	if input == nil {
		return true, nil
	}

	if _, err := GetKtpHandler(ctx, id); err != nil {
		return false, err
	}

	tanggal_lahir, err := time.Parse("2006-01-02", input.TanggalLahir)
	if err != nil {
		return false, err
	}
	tanggal_lahir = tanggal_lahir.In(time.Local)

	newKtp := models.Ktp{
		Nik:          input.Nik,
		Nama:         input.Nama,
		Agama:        input.Agama,
		JenisKelamin: input.JenisKelamin,
		TanggalLahir: tanggal_lahir,
		UpdatedAt:    time.Now(),
	}

	rs := database.DB.Model(&models.Ktp{}).Where("id = ?", id).Updates(newKtp)
	if rs.Error != nil {
		return false, rs.Error
	}

	return true, nil
}

// DeleteKtpHandler is a function that delete ktp with specific id
// This function will return boolean (success/not) and error (if exist)
func DeleteKtpHandler(ctx context.Context, id string) (bool, error) {

	deletedKtp, err := GetKtpHandler(ctx, id)
	if err != nil {
		return false, err
	}

	intId, _ := strconv.Atoi(id)
	rs := database.DB.Delete(deletedKtp, "id = ?", intId)
	if rs.Error != nil {
		return false, rs.Error
	}

	return true, nil
}

// PaginateKtpHandler is a function that return pagination of ktp
// This function will return paginate result
func PaginateKtpHandler(ctx context.Context, input model.Pagination) (*model.PaginationResultKtp, error) {

	// Generate SQL Query
	// query evaluation
	ktps := sq.Select("*").From("ktps")
	full := sq.Select("COUNT(id)").From("ktps")
	if input.Query != "" {
		ktps = ktps.Where("(nama LIKE ? OR nik LIKE ?)", fmt.Sprint("%", input.Query, "%"), fmt.Sprint("%", input.Query, "%"))
		full = full.Where("(nama LIKE ? OR nik LIKE ?)", fmt.Sprint("%", input.Query, "%"), fmt.Sprint("%", input.Query, "%"))
	}

	// after evaluation
	if input.After != nil {
		id, err := helper.DecodeID(*input.After)
		if err != nil {
			return nil, err
		}
		ktps = ktps.Where(sq.GtOrEq{"id": id})
		full = full.Where(sq.GtOrEq{"id": id})
	}

	// sort evaluation
	for _, val := range input.Sort {
		desc := strings.HasPrefix(val, "-")
		if desc {
			ktps = ktps.OrderBy(strings.Replace(val, "-", "", 1) + " desc")
		} else {
			ktps = ktps.OrderBy(val + " asc")
		}
	}

	// fetch full data
	sql1, args1, _ := full.ToSql()
	var fullDataId int
	res := database.DB.Raw(sql1, args1...).Scan(&fullDataId)
	if res.Error != nil {
		return nil, res.Error
	}

	// offset and limit evaluation
	ktps = ktps.Offset(uint64(input.Offset)).Limit(uint64(input.First))
	sql2, args2, _ := ktps.ToSql()
	limitedData := []models.Ktp{}
	res2 := database.DB.Raw(sql2, args2...).Scan(&limitedData)
	if res2.Error != nil {
		return nil, res.Error
	}

	// Construct Pagination Result
	// Construct Edge
	paginationEdges := []*model.PaginationEdgeKtp{}
	for _, ktp := range limitedData {
		nktp := ktp
		cursor := helper.EncodeID(ktp.ID)
		paginationEdge := model.PaginationEdgeKtp{
			Node:   &nktp,
			Cursor: cursor,
		}
		paginationEdges = append(paginationEdges, &paginationEdge)
	}

	// Construct dataCount
	dataCount := fullDataId

	// Construct Pagination Info
	endCursor := base64.StdEncoding.EncodeToString([]byte("0"))
	if len(paginationEdges) > 0 {
		endCursor = paginationEdges[len(paginationEdges)-1].Cursor
	}

	paginationInfo := model.PaginationInfoKtp{
		EndCursor:   endCursor,
		HasNextPAge: fullDataId-input.Offset-input.First > 0,
	}

	// Construct PaginationResultKtp
	paginationResultKtp := model.PaginationResultKtp{
		Totalcount: dataCount,
		Edges:      paginationEdges,
		PageInfo:   &paginationInfo,
	}

	return &paginationResultKtp, nil
}
