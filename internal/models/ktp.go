package models

import (
	"time"
)

//  Struktur data KTP (FULL)
type Ktp struct {
	ID           string    `json:"id"`
	Nik          string    `json:"nik"`
	Nama         string    `json:"nama"`
	Agama        string    `json:"agama"`
	JenisKelamin string    `gorm:"column:jenis_kelamin"`
	TanggalLahir time.Time `gorm:"column:tanggal_lahir"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}

//  TODO: make new ktp not declared here, but in graphql generated schema
//  Struktur data body KTP (Untuk keperluan create dan update)
type NewKtp struct {
	Nik          string `json:"nik"`
	Nama         string `json:"nama"`
	Agama        string `json:"agama"`
	JenisKelamin string `json:"jenis_kelamin"`
	TanggalLahir string `json:"tanggal_lahir"`
}

// TableName is name for database table
func (Ktp) TableName() string {
	return "ktps"
}
