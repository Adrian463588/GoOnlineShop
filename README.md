# Go Online Shop Backend Service

![Go](https://img.shields.io/badge/Go-1.20%2B-blue)
![Gin](https://img.shields.io/badge/Framework-Gin-blue)
![MySQL](https://img.shields.io/badge/Database-MySQL-blue)
![License](https://img.shields.io/badge/License-MIT-green)

## 📖 Project Overview

**Go Online Shop Backend Service** adalah backend RESTful API untuk aplikasi e-commerce marketplace yang dibangun menggunakan **Golang**, **Gin Web Framework**, dan **GORM** dengan database **MySQL**.

Proyek ini menerapkan **Clean Architecture** untuk memisahkan concerns antara handlers, services/repositories, dan models.

### 🎯 Fitur Utama

- ✅ **User Management**: Registrasi, Login (JWT Authentication), Profile updates
- ✅ **Store Management**: Auto-create store saat registrasi, kelola detail toko
- ✅ **Product Management**: CRUD operasi, upload gambar, manajemen stok
- ✅ **Category Management**: Admin-only category management
- ✅ **Address Management**: Manajemen alamat pengiriman
- ✅ **Transaction System**: Purchase transactions, stock deduction, transaction logs
- ✅ **Authentication**: JWT-based authentication dengan role-based access control

**Developer**: Adrian Syah Abidin  
**Repository**: [GoOnlineShop](https://github.com/Adrian463588/GoOnlineShop)

---

## 📂 Project Structure

```
ecommerce-backend/
├── main.go                 # Entry point aplikasi
├── go.mod                  # Go module definitions
├── go.sum                  # Go module checksums
├── internal/
│   ├── handler/            # HTTP Handlers (Controllers)
│   │   └── handler.go      # Business logic untuk endpoints
│   └── repository/         # Database interaction layer
│       └── repo.go         # GORM queries dan DB operations
├── models/
│   └── entity.go           # Database structs & request models
├── pkg/
│   ├── database/
│   │   └── mysql.go        # Database connection & auto-migration
│   ├── middleware/
│   │   ├── auth.go         # JWT Authentication middleware
│   │   └── admin.go        # Admin-only middleware
│   └── utils/
│       └── helper.go       # Utility functions (Password hashing, Response formatting)
└── public/
    └── uploads/            # Directory untuk product images
```

### 🏗️ Arsitektur

Proyek ini menerapkan **Clean Architecture**:

```
┌─────────────────────────────────────┐
│      HTTP Handler Layer (Gin)       │
│  (Request Validation & Response)    │
└──────────────┬──────────────────────┘
               │
┌──────────────▼──────────────────────┐
│     Business Logic / Services        │
│     (Handlers dalam handler.go)      │
└──────────────┬──────────────────────┘
               │
┌──────────────▼──────────────────────┐
│   Repository Layer (Database Layer)  │
│   (GORM Queries & DB Operations)    │
└──────────────┬──────────────────────┘
               │
┌──────────────▼──────────────────────┐
│        Database (MySQL)              │
└─────────────────────────────────────┘
```

---

## 📋 Prerequisites

Pastikan Anda sudah menginstal:

| Software | Versi | Deskripsi |
|----------|-------|-----------|
| **Go** | 1.20+ | Programming language |
| **MySQL** | 5.7+ | Database server |
| **Postman** | Latest | API testing tool |
| **Git** | Latest | Version control |

### Instalasi

- **Windows**: Download dari [golang.org](https://golang.org/dl) dan [mysql.com](https://www.mysql.com/downloads/)
- **macOS**: `brew install go mysql`
- **Linux**: `sudo apt install golang-go mysql-server`

---

## 🚀 Setup & Installation

### 1️⃣ Clone Repository

```bash
git clone https://github.com/Adrian463588/GoOnlineShop.git
cd ecommerce-backend
```

### 2️⃣ Database Setup

#### Option A: Manual Setup

1. Buka MySQL client (MySQL Workbench, phpMyAdmin, atau command line)
2. Buat database baru:

```sql
CREATE DATABASE evermos_db;
```

#### Option B: Docker (Opsional)

```bash
docker run --name mysql-ecommerce \
  -e MYSQL_ROOT_PASSWORD=root \
  -e MYSQL_DATABASE=evermos_db \
  -p 3306:3306 \
  -d mysql:8.0
```

**Catatan**: Anda tidak perlu membuat tabel secara manual. Aplikasi akan otomatis melakukan auto-migration saat dijalankan.

---

## 🔧 Database Configuration

### Buka File Konfigurasi

File: `pkg/database/mysql.go`

### Format DSN (Data Source Name)

```
username:password@tcp(host:port)/database?charset=utf8mb4&parseTime=True&loc=Local
```

### Contoh Konfigurasi

**Default (tanpa password)**:
```go
dsn := "root:@tcp(127.0.0.1:3306)/evermos_db?charset=utf8mb4&parseTime=True&loc=Local"
```

**Dengan password**:
```go
dsn := "root:password123@tcp(127.0.0.1:3306)/evermos_db?charset=utf8mb4&parseTime=True&loc=Local"
```

**Remote MySQL**:
```go
dsn := "username:password@tcp(remote-host:3306)/evermos_db?charset=utf8mb4&parseTime=True&loc=Local"
```

---

## ▶️ Running the Application

### Step 1: Install Dependencies

```bash
go mod tidy
```

**Output yang diharapkan**:
```
go: downloading ...
go: downloading ...
```

### Step 2: Run Server

```bash
go run main.go
```

### Step 3: Verifikasi Server Berjalan

Jika berhasil, Anda akan melihat output:

```
Database connected successfully
[GIN-debug] Listening and serving HTTP on :8000
```

✅ Server sekarang berjalan di `http://localhost:8000`

### Troubleshooting

Jika ada error:

```bash
# Error: Database connection failed
# Solusi: Pastikan MySQL running dan DSN benar

# Error: Port 8000 already in use
# Solusi: Change port di main.go atau kill process yang menggunakan port

# Error: Module not found
# Solusi: Run 'go mod tidy' dan 'go mod download'
```

---

## 🧪 Panduan Testing API dengan Postman

### 📦 Import Postman Collection

Anda dapat menggunakan Postman collection berikut:

**Postman Collection**: [View/Import Collection](https://www.postman.com/testerteam1230/workspace/testing/collection/19770079-2b5d4541-f0f6-413e-9d1b-b222a1f1d055?action=share&source=copy-link&creator=19770079)

---

### ⚙️ Setup Environment di Postman

#### 1. Buat Environment Baru

1. Klik **Environments** → **Create Environment**
2. Nama: `Local Development`

#### 2. Tambahkan Variables

```json
{
  "base_url": "http://localhost:8000/api/v1",
  "token": "",
  "user_id": "",
  "is_admin": false,
  "address_id": "",
  "product_id": "",
  "store_id": "",
  "category_id": "",
  "transaction_id": ""
}
```

#### 3. Save Environment

Klik **Save** dan pastikan environment `Local Development` sudah aktif di dropdown kanan atas.

---

### 🔑 Authentication & Token Management

#### Endpoint: Register User Baru

**Method**: `POST`  
**URL**: `{{base_url}}/auth/register`

**Request Body** (Raw - JSON):
```json
{
  "name": "John Doe",
  "phone": "081234567890",
  "email": "john@example.com",
  "password": "securePassword123",
  "dob": "1990-01-15",
  "job": "Software Engineer",
  "gender": "Male",
  "about": "I love coding",
  "province_id": 1,
  "city_id": 1
}
```

**Expected Response** (200):
```json
{
  "status": true,
  "message": "Succeed to POST data",
  "data": "Register Succeed"
}
```

---

#### Endpoint: Login & Dapatkan Token

**Method**: `POST`  
**URL**: `{{base_url}}/auth/login`

**Request Body** (Raw - JSON):
```json
{
  "phone": "081234567890",
  "password": "securePassword123"
}
```

**Expected Response** (200):
```json
{
  "status": true,
  "message": "Succeed to POST data",
  "data": {
    "nama": "John Doe",
    "notelp": "081234567890",
    "email": "john@example.com",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

---

#### 💾 Script: Simpan Token Otomatis

Di tab **Tests** endpoint Login, tambahkan script berikut:

```javascript
// Simpan token ke environment variable
if (pm.response.code === 200) {
    var jsonData = pm.response.json();
    
    // Simpan token
    pm.environment.set("token", jsonData.data.token);
    
    // Simpan user info
    pm.environment.set("user_phone", jsonData.data.notelp);
    pm.environment.set("user_name", jsonData.data.nama);
    pm.environment.set("user_id", jsonData.data.id);
    
    console.log("✓ Token berhasil disimpan!");
    console.log("User: " + jsonData.data.nama);
}
```

---

### 🔐 Setup Authorization Header

Semua endpoint terproteksi memerlukan JWT token. Ada 3 cara untuk menambahkan token:

#### Cara 1: Authorization Tab (Recommended)

1. Di tab **Authorization**
2. Pilih **Type**: `Bearer Token`
3. **Token**: `{{token}}`

#### Cara 2: Headers Manual

Di tab **Headers**, tambahkan:

```
Authorization: Bearer {{token}}
```

#### Cara 3: Custom Header

```
token: {{token}}
```

---

### 📊 API Endpoint Summary

#### Public Endpoints (Tanpa Token)

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/auth/register` | Register user baru |
| POST | `/auth/login` | Login & dapatkan token |
| GET | `/category` | Lihat semua kategori |
| GET | `/category/:id` | Lihat kategori spesifik |
| GET | `/product` | Lihat semua produk (dengan filter) |
| GET | `/product/:id` | Lihat produk spesifik |
| GET | `/toko` | Lihat semua toko |
| GET | `/toko/:id` | Lihat toko spesifik |

#### Protected Endpoints (Butuh Token)

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET | `/user` | Get profil user |
| PUT | `/user` | Update profil user |
| GET | `/user/alamat` | Get semua alamat |
| POST | `/user/alamat` | Create alamat baru |
| GET | `/user/alamat/:id` | Get alamat spesifik |
| PUT | `/user/alamat/:id` | Update alamat |
| DELETE | `/user/alamat/:id` | Delete alamat |
| GET | `/toko/my` | Get toko saya |
| PUT | `/toko/:id` | Update toko |
| POST | `/product` | Create produk |
| PUT | `/product/:id` | Update produk |
| DELETE | `/product/:id` | Delete produk |
| GET | `/trx` | Get semua transaksi |
| POST | `/trx` | Create transaksi |
| GET | `/trx/:id` | Get transaksi spesifik |

#### Admin Endpoints (Butuh Token + Admin Role)

| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/category` | Create kategori |
| PUT | `/category/:id` | Update kategori |
| DELETE | `/category/:id` | Delete kategori |

---

## 👑 Admin Access Setup

Untuk mengelola Categories, user harus menjadi Admin. Karena tidak ada "Register as Admin" endpoint untuk keamanan, ikuti langkah berikut:

### Step 1: Register User Normal

Register user biasa melalui endpoint `/auth/register`:

```json
{
  "name": "Admin User",
  "phone": "081234567890",
  "email": "admin@example.com",
  "password": "securePassword123",
  "dob": "1990-01-15",
  "job": "Administrator",
  "gender": "Male",
  "about": "Admin account",
  "province_id": 1,
  "city_id": 1
}
```

### Step 2: Buka Database Manager

Buka MySQL Workbench atau phpMyAdmin

### Step 3: Pilih Database

**MySQL Workbench**:
```
1. Double-click "evermos_db" di sidebar kiri hingga menjadi bold
2. Klik create table icon (atau Ctrl+T)
```

**phpMyAdmin**:
```
1. Pilih database "evermos_db"
2. Klik tab "SQL"
```

**Command Line**:
```bash
mysql -u root -p evermos_db
```

### Step 4: Jalankan Query untuk Promote User ke Admin

```sql
UPDATE users SET is_admin = 1 WHERE email = 'admin@example.com';
```

### Step 5: Verifikasi Update Berhasil

```sql
SELECT id, email, name, is_admin FROM users WHERE email = 'admin@example.com';
```

**Output yang diharapkan**:
```
| id | email             | name       | is_admin |
|----|-------------------|-----------|----------|
| 1  | admin@example.com | Admin User | 1        |
```

### Step 6: Login Ulang untuk Mendapatkan Token Admin

Setelah update `is_admin` ke 1, Anda harus login ulang agar JWT token berisi claim `is_admin: true`.

```
POST /auth/login
Body: {
  "phone": "081234567890",
  "password": "securePassword123"
}
```

Sekarang Anda dapat mengakses admin-only endpoints! ✅

---

## 📚 API Documentation Detail

### 1. Authentication Endpoints

#### Register
```
POST /auth/register
Content-Type: application/json

Request:
{
  "name": "string",
  "phone": "string (required, unique)",
  "email": "string (required, unique)",
  "password": "string (min 6 chars)",
  "dob": "string (YYYY-MM-DD)",
  "job": "string",
  "gender": "Male|Female",
  "about": "string",
  "province_id": "integer",
  "city_id": "integer"
}

Response: 200 OK
{
  "status": true,
  "message": "Succeed to POST data",
  "data": "Register Succeed"
}
```

#### Login
```
POST /auth/login
Content-Type: application/json

Request:
{
  "phone": "string (required)",
  "password": "string (required)"
}

Response: 200 OK
{
  "status": true,
  "message": "Succeed to POST data",
  "data": {
    "nama": "string",
    "notelp": "string",
    "email": "string",
    "token": "jwt_token"
  }
}

Error: 401 Unauthorized
{
  "status": false,
  "message": "Failed to POST data",
  "data": "No Telp atau kata sandi salah"
}
```

---

### 2. User Profile Endpoints

#### Get Profil
```
GET /user
Authorization: Bearer {token}

Response: 200 OK
{
  "status": true,
  "message": "Succeed to GET data",
  "data": {
    "id": "integer",
    "name": "string",
    "phone": "string",
    "email": "string",
    "dob": "string",
    "job": "string",
    "gender": "string",
    "about": "string",
    "is_admin": "boolean",
    "province_id": "integer",
    "city_id": "integer"
  }
}
```

#### Update Profil
```
PUT /user
Authorization: Bearer {token}
Content-Type: application/json

Request:
{
  "name": "string",
  "job": "string",
  "about": "string",
  "gender": "string",
  "password": "string (optional)"
}

Response: 200 OK
{
  "status": true,
  "message": "Succeed to UPDATE data"
}
```

---

### 3. Address Endpoints

#### Get All Address
```
GET /user/alamat
Authorization: Bearer {token}

Response: 200 OK
{
  "status": true,
  "message": "Succeed to GET data",
  "data": [
    {
      "id": "integer",
      "user_id": "integer",
      "receiver_name": "string",
      "phone": "string",
      "detail": "string",
      "created_at": "timestamp"
    }
  ]
}
```

#### Get Address by ID
```
GET /user/alamat/:id
Authorization: Bearer {token}

Response: 200 OK
{
  "status": true,
  "message": "Succeed to GET data",
  "data": {
    "id": "integer",
    "user_id": "integer",
    "receiver_name": "string",
    "phone": "string",
    "detail": "string"
  }
}
```

#### Create Address
```
POST /user/alamat
Authorization: Bearer {token}
Content-Type: application/json

Request:
{
  "receiver_name": "string",
  "phone": "string",
  "detail": "string"
}

Response: 200 OK
{
  "status": true,
  "message": "Succeed to POST data",
  "data": 1
}
```

#### Update Address
```
PUT /user/alamat/:id
Authorization: Bearer {token}
Content-Type: application/json

Request:
{
  "receiver_name": "string",
  "phone": "string",
  "detail": "string"
}

Response: 200 OK
{
  "status": true,
  "message": "Succeed to UPDATE data"
}
```

#### Delete Address
```
DELETE /user/alamat/:id
Authorization: Bearer {token}

Response: 200 OK
{
  "status": true,
  "message": "Succeed to DELETE data"
}
```

---

### 4. Product Endpoints

#### Get All Products (dengan filter & pagination)
```
GET /product?page=1&limit=10&nama_produk=laptop&category_id=1&min_harga=1000000&max_harga=10000000
Authorization: Optional

Query Parameters:
- page: integer (default: 1)
- limit: integer (default: 10)
- nama_produk: string (optional)
- category_id: integer (optional)
- toko_id: integer (optional)
- min_harga: integer (optional)
- max_harga: integer (optional)

Response: 200 OK
{
  "status": true,
  "message": "Succeed to GET data",
  "data": {
    "page": 1,
    "limit": 10,
    "data": [
      {
        "id": "integer",
        "name": "string",
        "category_id": "integer",
        "store_id": "integer",
        "reseller_price": "float",
        "consumer_price": "float",
        "stock": "integer",
        "description": "string"
      }
    ]
  }
}
```

#### Get Product by ID
```
GET /product/:id
Authorization: Optional

Response: 200 OK
{
  "status": true,
  "message": "Succeed to GET data",
  "data": {
    "id": "integer",
    "name": "string",
    "category_id": "integer",
    "store_id": "integer",
    "reseller_price": "float",
    "consumer_price": "float",
    "stock": "integer",
    "description": "string"
  }
}
```

#### Create Product
```
POST /product
Authorization: Bearer {token}
Content-Type: multipart/form-data

Form Data:
- nama_produk: string (required)
- harga_reseller: float (required)
- harga_konsumen: float (required)
- stok: integer (required)
- category_id: integer (required)
- deskripsi: string (required)
- photos: file[] (required, multiple files)

Response: 200 OK
{
  "status": true,
  "message": "Succeed to POST data",
  "data": 1
}
```

#### Update Product
```
PUT /product/:id
Authorization: Bearer {token}
Content-Type: multipart/form-data

Form Data:
- nama_produk: string
- harga_reseller: float
- harga_konsumen: float
- stok: integer
- deskripsi: string

Response: 200 OK
{
  "status": true,
  "message": "Succeed to UPDATE data"
}
```

#### Delete Product
```
DELETE /product/:id
Authorization: Bearer {token}

Response: 200 OK
{
  "status": true,
  "message": "Succeed to DELETE data"
}
```

---

### 5. Store Endpoints

#### Get My Store
```
GET /toko/my
Authorization: Bearer {token}

Response: 200 OK
{
  "status": true,
  "message": "Succeed to GET data",
  "data": {
    "id": "integer",
    "user_id": "integer",
    "name": "string",
    "photo_url": "string",
    "created_at": "timestamp"
  }
}
```

#### Update Store
```
PUT /toko/:id
Authorization: Bearer {token}
Content-Type: multipart/form-data

Form Data:
- nama_toko: string
- photo: file (optional)

Response: 200 OK
{
  "status": true,
  "message": "Succeed to UPDATE data",
  "data": "Update toko succeed"
}
```

#### Get All Stores
```
GET /toko?page=1&limit=10&nama=string
Authorization: Optional

Response: 200 OK
{
  "status": true,
  "message": "Succeed to GET data",
  "data": {
    "page": 1,
    "limit": 10,
    "data": [
      {
        "id": "integer",
        "user_id": "integer",
        "name": "string",
        "photo_url": "string"
      }
    ]
  }
}
```

#### Get Store by ID
```
GET /toko/:id
Authorization: Optional

Response: 200 OK
{
  "status": true,
  "message": "Succeed to GET data",
  "data": {
    "id": "integer",
    "user_id": "integer",
    "name": "string"
  }
}
```

---

### 6. Category Endpoints (Admin Only)

#### Get All Categories
```
GET /category
Authorization: Optional

Response: 200 OK
{
  "status": true,
  "message": "Succeed to GET data",
  "data": [
    {
      "id": "integer",
      "name": "string",
      "created_at": "timestamp"
    }
  ]
}
```

#### Get Category by ID
```
GET /category/:id
Authorization: Optional

Response: 200 OK
{
  "status": true,
  "message": "Succeed to GET data",
  "data": {
    "id": "integer",
    "name": "string"
  }
}
```

#### Create Category (Admin Only)
```
POST /category
Authorization: Bearer {token} (Admin)
Content-Type: application/json

Request:
{
  "name": "string"
}

Response: 200 OK
{
  "status": true,
  "message": "Succeed to POST data",
  "data": 1
}

Error: 403 Forbidden
{
  "status": false,
  "message": "Forbidden",
  "data": "Admin access required"
}
```

#### Update Category (Admin Only)
```
PUT /category/:id
Authorization: Bearer {token} (Admin)
Content-Type: application/json

Request:
{
  "name": "string"
}

Response: 200 OK
{
  "status": true,
  "message": "Succeed to UPDATE data"
}
```

#### Delete Category (Admin Only)
```
DELETE /category/:id
Authorization: Bearer {token} (Admin)

Response: 200 OK
{
  "status": true,
  "message": "Succeed to DELETE data"
}
```

---

### 7. Transaction Endpoints

#### Create Transaction
```
POST /trx
Authorization: Bearer {token}
Content-Type: application/json

Request:
{
  "alamat_kirim": "integer (address_id)",
  "method_bayar": "string (Transfer, Cash, etc)",
  "detail_trx": [
    {
      "product_id": "integer",
      "kuantitas": "integer"
    }
  ]
}

Response: 200 OK
{
  "status": true,
  "message": "Succeed to POST data",
  "data": 1
}
```

#### Get All Transactions
```
GET /trx
Authorization: Bearer {token}

Response: 200 OK
{
  "status": true,
  "message": "Succeed to GET data",
  "data": [
    {
      "id": "integer",
      "user_id": "integer",
      "address_id": "integer",
      "invoice_code": "string",
      "total_price": "float",
      "payment_method": "string",
      "status": "string",
      "created_at": "timestamp"
    }
  ]
}
```

#### Get Transaction by ID
```
GET /trx/:id
Authorization: Bearer {token}

Response: 200 OK
{
  "status": true,
  "message": "Succeed to GET data",
  "data": {
    "id": "integer",
    "user_id": "integer",
    "address_id": "integer",
    "invoice_code": "string",
    "total_price": "float",
    "payment_method": "string",
    "status": "string"
  }
}
```

---

## 🧪 Testing Workflow Rekomendasi

Ikuti workflow ini untuk testing yang sistematis:

### **Stage 1: Authentication** ✅
```
1. POST /auth/register        → Create user baru
2. POST /auth/login           → Get token
3. Save token ke environment
```

### **Stage 2: Public Endpoints** ✅
```
1. GET /category              → Lihat semua kategori
2. GET /category/:id          → Lihat kategori spesifik
3. GET /product               → Lihat semua produk
4. GET /product/:id           → Lihat produk spesifik
5. GET /toko                  → Lihat semua toko
6. GET /toko/:id              → Lihat toko spesifik
```

### **Stage 3: User Profile & Address** ✅
```
1. GET /user                  → Get profil
2. PUT /user                  → Update profil
3. GET /user/alamat           → Get semua alamat
4. POST /user/alamat          → Create alamat baru
5. GET /user/alamat/:id       → Get alamat spesifik
6. PUT /user/alamat/:id       → Update alamat
7. DELETE /user/alamat/:id    → Delete alamat
```

### **Stage 4: Store & Product Management** ✅
```
1. GET /toko/my               → Get toko saya
2. PUT /toko/:id              → Update toko
3. POST /product              → Create produk
4. PUT /product/:id           → Update produk
5. DELETE /product/:id        → Delete produk
```

### **Stage 5: Transactions** ✅
```
1. POST /trx                  → Create transaksi
2. GET /trx                   → Get semua transaksi
3. GET /trx/:id               → Get transaksi spesifik
```

### **Stage 6: Admin Features** ✅
```
(Gunakan admin account setelah setup admin access)
1. POST /category             → Create kategori
2. PUT /category/:id          → Update kategori
3. DELETE /category/:id       → Delete kategori
```

---

## 🐛 HTTP Status Codes

| Status Code | Meaning | Action |
|-------------|---------|--------|
| **200** | OK | Request berhasil |
| **201** | Created | Resource berhasil dibuat |
| **400** | Bad Request | Data invalid, cek request body |
| **401** | Unauthorized | Token tidak valid atau expired |
| **403** | Forbidden | Access denied (bukan admin/owner) |
| **404** | Not Found | Resource tidak ditemukan |
| **500** | Server Error | Error di server, cek logs |

---

## 🔍 Troubleshooting

### Database Connection Error

```
Error: dial tcp 127.0.0.1:3306: connection refused
```

**Solusi**:
- Pastikan MySQL service running
- Cek DSN di `pkg/database/mysql.go`
- Verifikasi host, port, username, password

```bash
# Windows
net start MySQL80

# macOS
brew services start mysql

# Linux
sudo systemctl start mysql
```

### Port Already in Use

```
Error: listen tcp :8000: bind: address already in use
```

**Solusi**:
- Kill process yang menggunakan port 8000
- Atau change port di `main.go`: `r.Run(":9000")`

```bash
# Windows
netstat -ano | findstr :8000
taskkill /PID {PID} /F

# macOS/Linux
lsof -i :8000
kill -9 {PID}
```

### Module Not Found

```
Error: cannot find module
```

**Solusi**:
```bash
go mod tidy
go mod download
```

### JWT Token Error

```
Error: Unauthorized - Invalid token
```

**Solusi**:
- Login ulang untuk dapatkan token baru
- Pastikan token sudah disimpan di environment Postman
- Cek format: `Authorization: Bearer {token}`

### File Upload Error

```
Error: Failed to save uploaded file
```

**Solusi**:
- Pastikan folder `public/uploads/` ada
- Create folder jika belum ada: `mkdir -p public/uploads`
- Check folder permissions

---

## 📝 Kode Snippet Penting

### Environment Variable Postman

**Test Script untuk simpan ID otomatis**:

```javascript
// Simpan response data ke environment
if (pm.response.code === 200 || pm.response.code === 201) {
    var jsonData = pm.response.json();
    pm.environment.set("last_id", jsonData.data);
    console.log("✓ ID disimpan: " + jsonData.data);
}
```

**Pre-request Script untuk auto-add token**:

```javascript
// Auto-add Bearer token
if (pm.environment.get("token")) {
    pm.request.headers.add({
        key: "Authorization",
        value: "Bearer " + pm.environment.get("token")
    });
}
```

---

## 🚀 Performance Tips

1. **Pagination**: Selalu gunakan `page` dan `limit` untuk query large datasets
2. **Filter**: Gunakan filter untuk mengurangi data yang diambil
3. **Indexing**: Database sudah ter-index, cukup efficient
4. **Caching**: Pertimbangkan untuk implement Redis caching di production



---

## 👨‍💻 Author

**Adrian Syah Abidin**  


---


