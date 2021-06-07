package handlers

import (
	"context"
	"ktp-fix/configs/database"
	"ktp-fix/internal/models"
	"log"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

//  TODO: fix not null bug
/*
Menambah data Ktp baru
	@param input (models.NewKtp): data Ktp baru yang akan dimasukkan
	@param ctx (context.Context): context dari graphQl

	@return (*models.Ktp): apabila terjadi error maka isinya nil, jika tidak isinya data ktp baru
	@return (error): apabila terjadi error maka isinya error, jika tidak isinya nil
*/
func CreateKtpHandler(ctx context.Context, input models.NewKtp) (*models.Ktp, error) {
	/*
		Parsing tanggal lahir + generate unique id
	*/
	tanggal_lahir, err := time.Parse("2006-01-02", input.TanggalLahir)
	if err != nil {
		return nil, err
	}
	tanggal_lahir = tanggal_lahir.In(time.Local)

	id, err := gonanoid.New()
	if err != nil {
		return nil, err
	}
	log.Print(id)

	/*
		Assignment + Penyimpanan di database
	*/
	var ktp models.Ktp = models.Ktp{
		ID:           id,
		Nik:          input.Nik,
		Nama:         input.Nama,
		Agama:        input.Agama,
		JenisKelamin: input.JenisKelamin,
		TanggalLahir: tanggal_lahir,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	rs := database.DB.Create(&ktp)
	if rs.Error == nil {
		return nil, rs.Error
	}

	return &ktp, nil
}

/*
Mendapatkan Ktp berdasarkan id
	@param id (string): id dari ktp yang akan dicari
	@param ctx (context.Context): context dari graphQl

	@return (*models.Ktp): apabila data tidak ditemukan maka isinya nil, jika tidak isinya data ktp berdasar id
	@return (error): apabila data tidak ditemukan maka isinya error, jika tidak isinya nil
*/
func GetKtpHandler(ctx context.Context, id string) (*models.Ktp, error) {
	selectedKtp := &models.Ktp{}
	rs := database.DB.First(selectedKtp, "id = ?", id)

	if rs.Error != nil {
		return nil, rs.Error
	}

	return selectedKtp, nil
}

/*
Mendapatkan semua data Ktp
	@param ctx (context.Context): context dari graphQl

	@return ([]*models.Ktp): apabila data tidak ditemukan maka isinya nil, jika tidak isinya data semua ktp
	@return (error): apabila data tidak ditemukan maka isinya error, jika tidak isinya nil
*/
func GetAllKtp(ctx context.Context) ([]*models.Ktp, error) {
	allUser := []*models.Ktp{}
	rs := database.DB.Find(&allUser)

	if rs.Error != nil {
		return nil, rs.Error
	}

	return allUser, nil
}

/*
Mengubah data Ktp dengan id tertentu
	@param input (*models.NewKtp): data Ktp yang akan diedit
	@param id (string): id dari Ktp yang akan diubah
	@param ctx (context.Context): context dari graphQl

	@return (boolean): apabila gagal dalam update maka isinya false, jika tidak isinya true
	@return (error): apabila gagal dalam update maka isinya error, jika tidak isinya nil
*/
func UpdateKtp(ctx context.Context, id string, input *models.NewKtp) (bool, error) {
	/*
		Parsing tanggal lahir + generate unique id
	*/
	if input == nil {
		return true, nil
	}

	tanggal_lahir, err := time.Parse("2006-01-02", input.TanggalLahir)
	if err != nil {
		return false, err
	}
	tanggal_lahir = tanggal_lahir.In(time.Local)

	updatedKtp := &models.Ktp{}
	var newKtp models.Ktp = models.Ktp{
		Nik:          input.Nik,
		Nama:         input.Nama,
		Agama:        input.Agama,
		JenisKelamin: input.JenisKelamin,
		TanggalLahir: tanggal_lahir,
		UpdatedAt:    time.Now(),
	}

	rs := database.DB.Model(updatedKtp).Where("id = ?", id).Updates(newKtp)
	if rs.Error != nil {
		return false, rs.Error
	}

	database.DB.Save(updatedKtp)

	return true, nil
}

/*
Menghapus data Ktp dengan id tertentu
	@param id (string): id dari Ktp yang akan dihapus
	@param ctx (context.Context): context dari graphQl

	@return (bool): gagal dalam update menghapus data maka isinya false, jika tidak isinya true
	@return (error): apabila gagal dalam menghapus data maka isinya error, jika tidak isinya nil
*/
func DeleteKtpHandler(ctx context.Context, id string) (bool, error) {
	deletedKtp, err := GetKtpHandler(ctx, id)

	if err != nil {
		return false, err
	}

	rs := database.DB.Delete(deletedKtp, "id = ?", id)

	if rs.Error != nil {
		return false, rs.Error
	}

	return true, nil
}
