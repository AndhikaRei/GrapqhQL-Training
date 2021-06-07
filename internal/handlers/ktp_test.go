package handlers

import (
	"context"
	"fmt"

	"ktp-fix/configs/database"
	"ktp-fix/internal/models"
	"log"
	"testing"

	"github.com/machinebox/graphql"
	"github.com/stretchr/testify/assert"
)

/*
Test membuat ktp dengan data valid
	@param t (*testing.T):

	@output: akan gagal apabila ktp tidak berhasil dibuat
	@output: akan sukses apabila ktp baru berhasil dibuat
*/
func TestCreateKtp(t *testing.T) {
	var client = graphql.NewClient("http://localhost:8080/query")

	var req = graphql.NewRequest(`
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

	/*
		Assertion
	*/
	nik := respData["createKtp"].(map[string]interface{})["nik"]
	assert.Equal(t, nik, "1335", "nik should be equal")

	nama := respData["createKtp"].(map[string]interface{})["nama"]
	assert.Equal(t, nama, "reihan", "nama should be equal")

	agama := respData["createKtp"].(map[string]interface{})["agama"]
	assert.Equal(t, agama, "islam", "agama should be equal")
}

/*
Test membuat ktp dengan data tidak valid
	@param t (*testing.T):

	@output: akan gagal apabila ktp baru berhasil dibuat
	@output: akan sukses apabila ktp tidak berhasil dibuat
*/
func TestCreateKtpFail(t *testing.T) {
	var client = graphql.NewClient("http://localhost:8080/query")

	var req = graphql.NewRequest(`
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

	/*
		Assertion
	*/
	assert.NotNil(t, err)
}

/*
Test mendapatkan ktp dengan id yang ada di database
	@param t (*testing.T):

	@output: akan gagal apabila tidak menemukan ktp dengan id tersebut
	@output: akan sukses apabila ktp dengan id tersebut ditemukan
*/
func TestGetKtp(t *testing.T) {
	err := database.ConnectDB()
	if err != nil {
		log.Fatal("Error happened")
	}

	var client = graphql.NewClient("http://localhost:8080/query")

	ktp := models.Ktp{}
	rs := database.DB.Where("nik = ?", "1335").First(&ktp)
	assert.Nil(t, rs.Error)
	id := ktp.ID

	var req = graphql.NewRequest(`
		query  {
			getKtp(id:"` + id + `") {
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

	/*
		Assertion
	*/
	fmt.Print(respData)

	nik := respData["getKtp"].(map[string]interface{})["nik"]
	assert.Equal(t, nik, "1335", "nik should be equal")

	nama := respData["getKtp"].(map[string]interface{})["nama"]
	assert.Equal(t, nama, "reihan", "nama should be equal")

	agama := respData["getKtp"].(map[string]interface{})["agama"]
	assert.Equal(t, agama, "islam", "agama should be equal")
}

/*
Test mendapatkan ktp dengan id yang tidak ada di database
	@param t (*testing.T):

	@output: akan gagal apabila ktp dengan id tersebut ditemukan
	@output: akan sukses apabila tidak menemukan ktp dengan id tersebut
*/
func TestGetKtpFail(t *testing.T) {
	err := database.ConnectDB()
	if err != nil {
		log.Fatal("Error happened")
	}

	var client = graphql.NewClient("http://localhost:8080/query")

	var req = graphql.NewRequest(`
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
	err = client.Run(ctx, req, &respData)

	/*
		Assertion
	*/
	assert.NotNil(t, err)
}

/*
Test mendapatkan semua ktp di database dan ada ktp di database
	@param t (*testing.T):

	@output: akan gagal apabila tidak ada ktp yang diambil
	@output: akan sukses apabila ada ktp yang diambil
*/
func TestGetAllKtp(t *testing.T) {
	var client = graphql.NewClient("http://localhost:8080/query")

	var req = graphql.NewRequest(`
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

	/*
		Assertion
	*/
	fmt.Print(respData)
	assert.NotEmpty(t, respData["getAllKtp"], "There must be at least one data")
}

/*
Test mengupdate ktp dengan id tertentu yang sudah ada di database dengan data yang lengkap
	@param t (*testing.T):

	@output: akan gagal apabila ktp gagal di update (tidak ditemukan atau alasan lain)
	@output: akan sukses apabila ktp berhasil di update
*/
func TestUpdateKtp(t *testing.T) {
	err := database.ConnectDB()
	if err != nil {
		log.Fatal("Error happened")
	}

	var client = graphql.NewClient("http://localhost:8080/query")

	ktp := models.Ktp{}
	rs := database.DB.Where("nik = ?", "1335").First(&ktp)
	assert.Nil(t, rs.Error)
	id := ktp.ID

	var req = graphql.NewRequest(`
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

	/*
		Assertion
	*/
	fmt.Print(respData)
	isUpdate := respData["updateKtp"]
	assert.Equal(t, isUpdate, true, "they should be equal")
}

/*
Test menghapus ktp dengan id tertentu dan id tersebut ada di database
	@param t (*testing.T):

	@output: akan gagal apabila tidak berhasil menghapus ktp dengan id tersebut (tidak ditemukan/alasan lain)
	@output: akan sukses apabila ktp dengan id tersebut ditemukan dan berhasil dihapus
*/
func TestDeleteKtp(t *testing.T) {

	err := database.ConnectDB()
	if err != nil {
		log.Fatal("Error happened")
	}

	var client = graphql.NewClient("http://localhost:8080/query")

	ktp := models.Ktp{}
	rs := database.DB.Where("nik = ?", "13352").First(&ktp)
	assert.Nil(t, rs.Error)
	id := ktp.ID

	var req = graphql.NewRequest(`
		mutation {
			deleteKtp(id:"` + id + `")
		}	
	`)

	ctx := context.Background()

	var respData map[string]interface{}
	if err := client.Run(ctx, req, &respData); err != nil {
		t.Error(err)
	}

	/*
		Assertion
	*/
	fmt.Print(respData)
	isDeleted := respData["deleteKtp"]
	assert.Equal(t, isDeleted, true, "Must be deleted")
}
