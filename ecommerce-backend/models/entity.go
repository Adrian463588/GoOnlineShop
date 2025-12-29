package models

import (
	"time"

	"gorm.io/gorm"
)

// User Entity
type User struct {
	ID           uint           `gorm:"primaryKey;column:id" json:"id"`
	Name         string         `gorm:"column:nama" json:"nama"`
	Password     string         `gorm:"column:kata_sandi" json:"-"`
	Phone        string         `gorm:"unique;column:notelp" json:"no_telp"`
	Email        string         `gorm:"unique;column:email" json:"email"`
	DOB          string         `gorm:"column:tanggal_lahir" json:"tanggal_lahir"`
	Gender       string         `gorm:"column:jenis_kelamin" json:"jenis_kelamin"`
	About        string         `gorm:"column:tentang" json:"tentang"`
	Job          string         `gorm:"column:pekerjaan" json:"pekerjaan"`
	ProvinceID   string         `gorm:"column:id_provinsi" json:"id_provinsi"`
	CityID       string         `gorm:"column:id_kota" json:"id_kota"`
	IsAdmin      bool           `gorm:"default:false;column:isAdmin" json:"-"`
	Store        Store          `gorm:"foreignKey:UserID;references:ID" json:"toko,omitempty"`
	CreatedAt    time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

// Address Entity
type Address struct {
	ID           uint      `gorm:"primaryKey;column:id" json:"id"`
	UserID       uint      `gorm:"column:id_user" json:"id_user"`
	Title        string    `gorm:"column:judul_alamat" json:"judul_alamat"`
	ReceiverName string    `gorm:"column:nama_penerima" json:"nama_penerima"`
	Phone        string    `gorm:"column:no_telp" json:"no_telp"`
	Detail       string    `gorm:"column:detail_alamat" json:"detail_alamat"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"-"`
}

// Store Entity
type Store struct {
	ID        uint      `gorm:"primaryKey;column:id" json:"id"`
	UserID    uint      `gorm:"column:id_user" json:"id_user"`
	Name      string    `gorm:"column:nama_toko" json:"nama_toko"`
	PhotoURL  string    `gorm:"column:url_foto" json:"url_foto"`
	CreatedAt time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"-"`
}

// Category Entity
type Category struct {
	ID        uint      `gorm:"primaryKey;column:id" json:"id"`
	Name      string    `gorm:"column:nama_category" json:"nama_category"`
	CreatedAt time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"-"`
}

// Product Entity
type Product struct {
	ID            uint           `gorm:"primaryKey;column:id" json:"id"`
	StoreID       uint           `gorm:"column:id_toko" json:"toko_id"`
	CategoryID    uint           `gorm:"column:id_category" json:"category_id"`
	Name          string         `gorm:"column:nama_produk" json:"nama_produk"`
	Slug          string         `gorm:"column:slug" json:"slug"`
	ResellerPrice float64        `gorm:"column:harga_reseller" json:"harga_reseller"`
	ConsumerPrice float64        `gorm:"column:harga_konsumen" json:"harga_konsumen"`
	Stock         int            `gorm:"column:stok" json:"stok"`
	Description   string         `gorm:"column:deskripsi" json:"deskripsi"`
	Store         Store          `gorm:"foreignKey:StoreID" json:"toko"`
	Category      Category       `gorm:"foreignKey:CategoryID" json:"category"`
	Photos        []ProductPhoto `gorm:"foreignKey:ProductID" json:"photos"`
	CreatedAt     time.Time      `gorm:"column:created_at" json:"-"`
	UpdatedAt     time.Time      `gorm:"column:updated_at" json:"-"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

// Product Photo Entity
type ProductPhoto struct {
	ID        uint      `gorm:"primaryKey;column:id" json:"id"`
	ProductID uint      `gorm:"column:id_produk" json:"product_id"`
	URL       string    `gorm:"column:url" json:"url"`
	CreatedAt time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"-"`
}

// Transaction Entity
type Transaction struct {
	ID            uint                `gorm:"primaryKey;column:id" json:"id"`
	UserID        uint                `gorm:"column:id_user" json:"user_id"`
	AddressID     uint                `gorm:"column:alamat_pengiriman" json:"alamat_kirim"`
	TotalPrice    float64             `gorm:"column:harga_total" json:"harga_total"`
	InvoiceCode   string              `gorm:"column:kode_invoice" json:"kode_invoice"`
	PaymentMethod string              `gorm:"column:method_bayar" json:"method_bayar"`
	Address       Address             `gorm:"foreignKey:AddressID" json:"detail_alamat"`
	Details       []TransactionDetail `gorm:"foreignKey:TransactionID" json:"detail_trx"`
	CreatedAt     time.Time           `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time           `gorm:"column:updated_at" json:"updated_at"`
}

// Transaction Detail Entity
type TransactionDetail struct {
	ID            uint       `gorm:"primaryKey;column:id" json:"id"`
	TransactionID uint       `gorm:"column:id_trx" json:"trx_id"`
	ProductLogID  uint       `gorm:"column:id_log_produk" json:"log_product_id"`
	StoreID       uint       `gorm:"column:id_toko" json:"store_id"`
	Quantity      int        `gorm:"column:kuantitas" json:"kuantitas"`
	TotalPrice    float64    `gorm:"column:harga_total" json:"harga_total"`
	ProductLog    ProductLog `gorm:"foreignKey:ProductLogID" json:"product"`
	CreatedAt     time.Time  `gorm:"column:created_at" json:"-"`
	UpdatedAt     time.Time  `gorm:"column:updated_at" json:"-"`
}

// Product Log Entity (Snapshot)
type ProductLog struct {
	ID            uint      `gorm:"primaryKey;column:id" json:"id"`
	ProductID     uint      `gorm:"column:id_produk" json:"product_id"`
	StoreID       uint      `gorm:"column:id_toko" json:"store_id"`
	CategoryID    uint      `gorm:"column:id_category" json:"category_id"`
	Name          string    `gorm:"column:nama_produk" json:"nama_produk"`
	Slug          string    `gorm:"column:slug" json:"slug"`
	ResellerPrice float64   `gorm:"column:harga_reseller" json:"harga_reseller"`
	ConsumerPrice float64   `gorm:"column:harga_konsumen" json:"harga_konsumen"`
	Description   string    `gorm:"column:deskripsi" json:"deskripsi"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// API Response Wrappers
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

type Pagination struct {
	Page  int         `json:"page"`
	Limit int         `json:"limit"`
	Data  interface{} `json:"data"`
}

// Request Binding Structs
type RegisterRequest struct {
	Name       string `json:"nama" binding:"required"`
	Phone      string `json:"no_telp" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"kata_sandi" binding:"required,min=6"`
	DOB        string `json:"tanggal_lahir"`
	Gender     string `json:"jenis_kelamin"`
	About      string `json:"tentang"`
	Job        string `json:"pekerjaan"`
	ProvinceID string `json:"id_provinsi"`
	CityID     string `json:"id_kota"`
}

type LoginRequest struct {
	Phone    string `json:"no_telp" binding:"required"`
	Password string `json:"kata_sandi" binding:"required"`
}

// TrxItemRequest is a strict struct for transaction items
type TrxItemRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
	Kuantitas int  `json:"kuantitas" binding:"required,gt=0"`
}

type TrxRequest struct {
	MethodBayar string           `json:"method_bayar" binding:"required"`
	AlamatKirim uint             `json:"alamat_kirim" binding:"required"`
	DetailTrx   []TrxItemRequest `json:"detail_trx" binding:"required,dive"`
}