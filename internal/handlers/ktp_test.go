package handlers

import (
	"context"
	"strconv"

	"ktp-fix/configs/database"
	"ktp-fix/internal/models"
	"log"
	"testing"

	"github.com/machinebox/graphql"
	"github.com/stretchr/testify/assert"
)

// TestCreateKtp is a function that will test create new ktp with valid data
func TestCreateKtp(t *testing.T) {

	client := graphql.NewClient("http://localhost:8080/query")

	req := graphql.NewRequest(`
		mutation {
			createKtp(input: {
				nik: "1335", 
				tanggal_lahir: "2023-05-01", 
				nama: "reihan",
				agama: "islam", 
				jenis_kelamin: "laki-laki"
			}){
				nik,
				nama,
				agama
			}
		  }	
	`)

	ctx := context.Background()
	var respData map[string]interface{}
	if err := client.Run(ctx, req, &respData); err != nil {
		t.Error(err)

	}

	// Assertion
	nik := respData["createKtp"].(map[string]interface{})["nik"]
	assert.Equal(t, nik, "1335", "nik should be equal")

	nama := respData["createKtp"].(map[string]interface{})["nama"]
	assert.Equal(t, nama, "reihan", "nama should be equal")

	agama := respData["createKtp"].(map[string]interface{})["agama"]
	assert.Equal(t, agama, "islam", "agama should be equal")
}

// TestCreateKtpFail is a function that will test create new ktp with invalid data
func TestCreateKtpFail(t *testing.T) {

	client := graphql.NewClient("http://localhost:8080/query")

	req := graphql.NewRequest(`
		mutation {
			createKtp(input: {
				nik: "1335", 
				tanggal_lahir: "2023-05-01", 
				nama: "reihan",
				jenis_kelamin: "laki-laki"
			}){
				nik,
				nama,
				agama
			}
		  }	
	`)

	ctx := context.Background()

	var respData map[string]interface{}
	err := client.Run(ctx, req, &respData)

	// Assertion
	assert.NotNil(t, err)
}

// TestGetKtp is a function that will test getting ktp with valid id on database
func TestGetKtp(t *testing.T) {

	err := database.ConnectDB()
	if err != nil {
		log.Fatal("Error happened")
	}

	client := graphql.NewClient("http://localhost:8080/query")

	ktp := models.Ktp{}
	rs := database.DB.Where("nik = ?", "1335").First(&ktp)
	assert.Nil(t, rs.Error)
	id := strconv.Itoa(ktp.ID)

	req := graphql.NewRequest(`
		query  {
			getKtp(id:` + id + `) {
				nik,
				nama,
				agama
			}
		}	
	`)

	ctx := context.Background()

	var respData map[string]interface{}
	if err := client.Run(ctx, req, &respData); err != nil {
		t.Error(err)
	}

	// Assertion
	nik := respData["getKtp"].(map[string]interface{})["nik"]
	assert.Equal(t, nik, "1335", "nik should be equal")

	nama := respData["getKtp"].(map[string]interface{})["nama"]
	assert.Equal(t, nama, "reihan", "nama should be equal")

	agama := respData["getKtp"].(map[string]interface{})["agama"]
	assert.Equal(t, agama, "islam", "agama should be equal")
}

// TestGetKtpFail is a function that will test getting ktp with invalid id on database
func TestGetKtpFail(t *testing.T) {

	client := graphql.NewClient("http://localhost:8080/query")

	req := graphql.NewRequest(`
		query  {
			getKtp(id:"` + `abcdefghijklm` + `") {
				nik,
				nama,
				agama
			}
		}	
	`)

	ctx := context.Background()

	var respData map[string]interface{}
	err := client.Run(ctx, req, &respData)

	// Assertion
	assert.NotNil(t, err)
}

// TestGetAllKtp is a function that will test getting al Ktp in database
func TestGetAllKtp(t *testing.T) {

	client := graphql.NewClient("http://localhost:8080/query")

	req := graphql.NewRequest(`
		query  {
			getAllKtp {
		  		nik
			}
		}
	`)

	ctx := context.Background()

	var respData map[string]interface{}
	if err := client.Run(ctx, req, &respData); err != nil {
		t.Error(err)
	}

	// Assertion
	assert.NotEmpty(t, respData["getAllKtp"], "There must be at least one data")
}

// TestUpdateKtp is a function that will test updating ktp with valid id and valid data
func TestUpdateKtp(t *testing.T) {

	err := database.ConnectDB()
	if err != nil {
		log.Fatal("Error happened")
	}

	client := graphql.NewClient("http://localhost:8080/query")

	ktp := models.Ktp{}
	rs := database.DB.Where("nik = ?", "1335").First(&ktp)
	assert.Nil(t, rs.Error)
	id := strconv.Itoa(ktp.ID)

	req := graphql.NewRequest(`
		mutation {
			updateKtp(id:"` + id + `"` + ` ,input:{
				nik: "13352", 
				tanggal_lahir: "2023-05-01", 
				nama: "reihanUpdated2",
				agama: "islamUpdated", 
				jenis_kelamin: "laki-laki"
			})
		  }	
	`)

	ctx := context.Background()

	var respData map[string]interface{}
	if err := client.Run(ctx, req, &respData); err != nil {
		t.Error(err)
	}

	// Assertion
	isUpdate := respData["updateKtp"]
	assert.Equal(t, isUpdate, true, "they should be equal")
}

// TestUpdateKtpFail is a function that will test updating ktp with invalid id and invalid data
func TestUpdateKtpFail(t *testing.T) {

	client := graphql.NewClient("http://localhost:8080/query")

	req := graphql.NewRequest(`
		mutation {
			updateKtp(id:"` + "afwwfafewsas" + `"` + ` ,input:{
				nik: "13352", 
				tanggal_lahir: "2023-05-01", 
				nama: "reihanUpdated2",
				agama: "islamUpdated", 
			})
		  }	
	`)

	ctx := context.Background()

	var respData map[string]interface{}
	err := client.Run(ctx, req, &respData)

	// Assertion
	assert.NotNil(t, err)
}

// TestDeleteKtp is a function that will test deleting ktp with valid id
func TestDeleteKtp(t *testing.T) {

	err := database.ConnectDB()
	if err != nil {
		log.Fatal("Error happened")
	}

	client := graphql.NewClient("http://localhost:8080/query")

	ktp := models.Ktp{}
	rs := database.DB.Where("nik = ?", "13352").First(&ktp)
	assert.Nil(t, rs.Error)
	id := strconv.Itoa(ktp.ID)

	req := graphql.NewRequest(`
		mutation {
			deleteKtp(id:"` + id + `")
		}	
	`)

	ctx := context.Background()

	var respData map[string]interface{}
	if err := client.Run(ctx, req, &respData); err != nil {
		t.Error(err)
	}

	isDeleted := respData["deleteKtp"]

	// Assertion
	assert.Equal(t, isDeleted, true, "Must be deleted")
}

// TestDeleteKtpFail is a function that will test deleting ktp with invalid id
func TestDeleteKtpFail(t *testing.T) {

	client := graphql.NewClient("http://localhost:8080/query")

	req := graphql.NewRequest(`
		mutation {
			deleteKtp(id:"` + "qfdqwfweefdweaewfw" + `")
		}	
	`)

	ctx := context.Background()

	var respData map[string]interface{}
	err := client.Run(ctx, req, &respData)

	// Assertion
	assert.NotNil(t, err)
}

// Testing pagination implementation of Ktp
func TestPaginationKtp(t *testing.T) {

	err := database.ConnectDB()
	if err != nil {
		log.Fatal("Error happened")
	}

	client := graphql.NewClient("http://localhost:8080/query")

	req := graphql.NewRequest(`
		mutation {
			createKtp(input: {
				nik: "1335", 
				tanggal_lahir: "2023-05-01", 
				nama: "reihan",
				jenis_kelamin: "laki-laki",
				agama: "islam" 
			}){
				nik,
				nama,
				agama
			}
		  }	
	`)

	ctx := context.Background()
	var respData map[string]interface{}
	for i := 0; i < 10; i++ {
		if err := client.Run(ctx, req, &respData); err != nil {
			t.Error(err)
		}
	}

	req2 := graphql.NewRequest(`
		query {
			paginateKtp(input: {
				first:5,
				offset:3,
				after:"NQ==",
				query:"reihan",
				sort: ["id"],
		}){
			totalcount,
			pageInfo{
				endCursor,
				hasNextPAge
			},
			edges{
				cursor, 
				node{
					jenis_kelamin,
					nik,
					nama
				}
			}
		}}
	`)
	ctx = context.Background()
	var respData2 map[string]interface{}
	if err2 := client.Run(ctx, req2, &respData2); err2 != nil {
		t.Error(err)
	}

	// Assertion
	respData2 = respData2["paginateKtp"].(map[string]interface{})
	pageInfo := respData2["pageInfo"].(map[string]interface{})
	hasNextPage := pageInfo["hasNextPAge"].(bool)
	endCursor := pageInfo["endCursor"].(string)
	assert.False(t, hasNextPage)
	assert.Equal(t, "MTE=", endCursor)

	totalCount := respData2["totalcount"].(float64)
	assert.Equal(t, float64(7), totalCount)

}

// Testing pagination implementation of Ktp with invalid input
func TestPaginationKtpFail(t *testing.T) {

	client := graphql.NewClient("http://localhost:8080/query")
	req := graphql.NewRequest(`
		query {
			paginateKtp(input: {
				first:"sdwef",
				after:"ewffwe",
				query:"reihan",
				sort: ["id asc"],
		}){
			totalcount,
			pageInfo{
				endCursor,
				hasNextPAge
			},
			edges{
				cursor, 
				node{
					jenis_kelamin,
					nik,
					nama
				}
			}
		}}
	`)
	ctx := context.Background()
	var respData map[string]interface{}
	err := client.Run(ctx, req, &respData)

	// Assertion
	assert.NotNil(t, err)
}
