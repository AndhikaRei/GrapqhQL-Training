package models

import (
	"time"
)

// Ktp is the implementation of Indonesian Ktp structure
type Ktp struct {
	ID           int       `json:"id"`
	Nik          string    `json:"nik"`
	Nama         string    `json:"nama"`
	Agama        string    `json:"agama"`
	JenisKelamin string    `gorm:"column:jenis_kelamin"`
	TanggalLahir time.Time `gorm:"column:tanggal_lahir"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}

// TODO: make new ktp not declared here, but in graphql generated schema
// NewKtp is the body ktp for update and create query
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
