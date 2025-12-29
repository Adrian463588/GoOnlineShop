package handler

import (
	"ecommerce-backend/internal/repository"
	"ecommerce-backend/models"
	"ecommerce-backend/pkg/database"
	"ecommerce-backend/pkg/utils"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// --- Auth Handlers ---

func Register(c *gin.Context) {
	var input models.RegisterRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.APIResponse(c, http.StatusBadRequest, false, "Failed to POST data", nil, []string{err.Error()})
		return
	}

	hashedPwd, _ := utils.HashPassword(input.Password)
	user := models.User{
		Name: input.Name, Phone: input.Phone, Email: input.Email, 
		Password: hashedPwd, DOB: input.DOB, Job: input.Job, 
		Gender: input.Gender, About: input.About, // New Fields
		ProvinceID: input.ProvinceID, CityID: input.CityID,
	}

	if err := repository.CreateUser(&user); err != nil {
		utils.APIResponse(c, http.StatusBadRequest, false, "Failed to POST data", nil, []string{err.Error()})
		return
	}

	// Auto create store
	store := models.Store{UserID: user.ID, Name: user.Name + "'s Store"}
	repository.UpdateStore(&store)

	utils.APIResponse(c, http.StatusOK, true, "Succeed to POST data", "Register Succeed", nil)
}

func Login(c *gin.Context) {
	var input models.LoginRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.APIResponse(c, http.StatusBadRequest, false, "Failed to POST data", nil, []string{err.Error()})
		return
	}

	user, err := repository.FindUserByPhone(input.Phone)
	if err != nil || !utils.CheckPassword(user.Password, input.Password) {
		utils.APIResponse(c, http.StatusUnauthorized, false, "Failed to POST data", nil, []string{"No Telp atau kata sandi salah"})
		return
	}

	token, _ := utils.GenerateToken(user.ID, user.IsAdmin)
	
	response := map[string]interface{}{
		"nama": user.Name, "no_telp": user.Phone, "email": user.Email, "token": token,
	}
	utils.APIResponse(c, http.StatusOK, true, "Succeed to POST data", response, nil)
}

// --- User/Profile Handlers ---

func GetProfile(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	user, err := repository.FindUserByID(userID)
	if err != nil {
		utils.APIResponse(c, http.StatusNotFound, false, "Failed to GET data", nil, []string{"User not found"})
		return
	}
	utils.APIResponse(c, http.StatusOK, true, "Succeed to GET data", user, nil)
}

func UpdateProfile(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.APIResponse(c, http.StatusBadRequest, false, "Failed to UPDATE data", nil, []string{err.Error()})
		return
	}

	user, err := repository.FindUserByID(userID)
	if err != nil {
		utils.APIResponse(c, http.StatusNotFound, false, "User not found", nil, nil)
		return
	}

	// Update fields
	user.Name = input.Name
	user.Job = input.Job
	user.About = input.About
	user.Gender = input.Gender
	if input.Password != "" {
		hash, _ := utils.HashPassword(input.Password)
		user.Password = hash
	}
	
	repository.UpdateUser(&user)
	utils.APIResponse(c, http.StatusOK, true, "Succeed to UPDATE data", "", nil)
}

// --- Address Handlers ---

func GetMyAddress(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	addresses, _ := repository.GetAddressesByUserID(userID)
	utils.APIResponse(c, http.StatusOK, true, "Succeed to GET data", addresses, nil)
}

func GetAddressByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	address, err := repository.GetAddressByID(uint(id))
	if err != nil {
		utils.APIResponse(c, http.StatusNotFound, false, "Failed to GET data", nil, []string{"Address not found"})
		return
	}

	userID := c.MustGet("user_id").(uint)
	if address.UserID != userID {
		utils.APIResponse(c, http.StatusForbidden, false, "Failed", nil, []string{"Forbidden"})
		return
	}

	utils.APIResponse(c, http.StatusOK, true, "Succeed to GET data", address, nil)
}

func CreateAddress(c *gin.Context) {
	var input models.Address
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.APIResponse(c, http.StatusBadRequest, false, "Failed", nil, []string{err.Error()})
		return
	}
	input.UserID = c.MustGet("user_id").(uint)
	repository.CreateAddress(&input)
	utils.APIResponse(c, http.StatusOK, true, "Succeed to POST data", 1, nil)
}

func UpdateAddress(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input models.Address
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.APIResponse(c, http.StatusBadRequest, false, "Failed", nil, []string{err.Error()})
		return
	}

	address, err := repository.GetAddressByID(uint(id))
	if err != nil {
		utils.APIResponse(c, http.StatusBadRequest, false, "Failed to GET data", nil, []string{"record not found"})
		return
	}

	if address.UserID != c.MustGet("user_id").(uint) {
		utils.APIResponse(c, http.StatusForbidden, false, "Forbidden", nil, nil)
		return
	}

	address.ReceiverName = input.ReceiverName
	address.Phone = input.Phone
	address.Detail = input.Detail
	repository.UpdateAddress(&address)
	utils.APIResponse(c, http.StatusOK, true, "Succeed to GET data", "", nil)
}

func DeleteAddress(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	address, err := repository.GetAddressByID(uint(id))
	if err != nil {
		utils.APIResponse(c, http.StatusBadRequest, false, "Failed to GET data", nil, []string{"record not found"})
		return
	}

	if address.UserID != c.MustGet("user_id").(uint) {
		utils.APIResponse(c, http.StatusForbidden, false, "Forbidden", nil, nil)
		return
	}

	repository.DeleteAddress(address.ID)
	utils.APIResponse(c, http.StatusOK, true, "Succeed to GET data", "", nil)
}

// --- Store Handlers ---

func GetMyStore(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	store, err := repository.GetStoreByUserID(userID)
	if err != nil {
		utils.APIResponse(c, http.StatusNotFound, false, "Store not found", nil, nil)
		return
	}
	utils.APIResponse(c, http.StatusOK, true, "Succeed to GET data", store, nil)
}

func UpdateStore(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id_toko"))
	store, err := repository.GetStoreByID(uint(id))
	
	if err != nil {
		utils.APIResponse(c, http.StatusNotFound, false, "Store not found", nil, nil)
		return
	}
	
	if store.UserID != c.MustGet("user_id").(uint) {
		utils.APIResponse(c, http.StatusForbidden, false, "Forbidden", nil, nil)
		return
	}

	store.Name = c.PostForm("nama_toko")
	filename, err := utils.SaveUploadedFile(c, "photo")
	if err == nil {
		store.PhotoURL = filename
	}

	repository.UpdateStore(&store)
	utils.APIResponse(c, http.StatusOK, true, "Succeed to UPDATE data", "Update toko succeed", nil)
}

func GetStoreByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id_toko"))
	store, err := repository.GetStoreByID(uint(id))
	if err != nil {
		utils.APIResponse(c, http.StatusNotFound, false, "Failed to GET data", nil, []string{"Toko tidak ditemukan"})
		return
	}
	utils.APIResponse(c, http.StatusOK, true, "Succeed to GET data", store, nil)
}

func GetAllStores(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	name := c.Query("nama")

	stores, _, _ := repository.GetAllStores(page, limit, name)
	utils.APIResponse(c, http.StatusOK, true, "Succeed to GET data", models.Pagination{Page: page, Limit: limit, Data: stores}, nil)
}

// --- Category Handlers (Admin) ---

func GetAllCategory(c *gin.Context) {
	cats, _ := repository.GetAllCategories()
	utils.APIResponse(c, http.StatusOK, true, "Succeed to GET data", cats, nil)
}

func CreateCategory(c *gin.Context) {
	var input models.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.APIResponse(c, http.StatusBadRequest, false, "Failed", nil, nil)
		return
	}
	repository.CreateCategory(&input)
	utils.APIResponse(c, http.StatusOK, true, "Succeed to POST data", input.ID, nil)
}

func GetCategoryByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	cat, err := repository.GetCategoryByID(uint(id))
	if err != nil {
		utils.APIResponse(c, http.StatusNotFound, false, "Failed to GET data", nil, []string{"No Data Category"})
		return
	}
	utils.APIResponse(c, http.StatusOK, true, "Succeed to GET data", cat, nil)
}

func UpdateCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input models.Category
	c.ShouldBindJSON(&input)
	
	cat, err := repository.GetCategoryByID(uint(id))
	if err != nil {
		utils.APIResponse(c, http.StatusNotFound, false, "Not Found", nil, nil)
		return
	}
	cat.Name = input.Name
	repository.UpdateCategory(&cat)
	utils.APIResponse(c, http.StatusOK, true, "Succeed to GET data", "", nil)
}

func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := repository.DeleteCategory(uint(id))
	if err != nil {
		utils.APIResponse(c, http.StatusBadRequest, false, "Failed to GET data", nil, []string{"record not found"})
		return
	}
	utils.APIResponse(c, http.StatusOK, true, "Succeed to GET data", "", nil)
}

// --- Product Handlers ---

func GetAllProducts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	name := c.Query("nama_produk")
	catID := c.Query("category_id")
	storeID := c.Query("toko_id")
	maxP := c.Query("max_harga")
	minP := c.Query("min_harga")

	products, _, _ := repository.GetProducts(page, limit, name, catID, storeID, maxP, minP)
	utils.APIResponse(c, http.StatusOK, true, "Succeed to GET data", models.Pagination{Page: page, Limit: limit, Data: products}, nil)
}

func GetProductByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	prod, err := repository.GetProductByID(uint(id))
	if err != nil {
		utils.APIResponse(c, http.StatusNotFound, false, "Failed to GET data", nil, []string{"No Data Product"})
		return
	}
	utils.APIResponse(c, http.StatusOK, true, "Succeed to GET data", prod, nil)
}

func CreateProduct(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	store, err := repository.GetStoreByUserID(userID)
	if err != nil {
		utils.APIResponse(c, http.StatusBadRequest, false, "User has no store", nil, nil)
		return
	}

	priceRes, _ := strconv.ParseFloat(c.PostForm("harga_reseller"), 64)
	priceCons, _ := strconv.ParseFloat(c.PostForm("harga_konsumen"), 64)
	stock, _ := strconv.Atoi(c.PostForm("stok"))
	catID, _ := strconv.Atoi(c.PostForm("category_id"))

	product := models.Product{
		Name:          c.PostForm("nama_produk"),
		CategoryID:    uint(catID),
		StoreID:       store.ID,
		ResellerPrice: priceRes,
		ConsumerPrice: priceCons,
		Stock:         stock,
		Description:   c.PostForm("deskripsi"),
		Slug:          utils.Slugify(c.PostForm("nama_produk")),
	}

	if err := repository.CreateProduct(&product); err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, false, "Error", nil, nil)
		return
	}

	form, _ := c.MultipartForm()
	files := form.File["photos"]
	for _, file := range files {
		filename := fmt.Sprintf("%d-%s", time.Now().Unix(), file.Filename)
		c.SaveUploadedFile(file, "public/uploads/"+filename)
		database.DB.Create(&models.ProductPhoto{ProductID: product.ID, URL: filename})
	}

	utils.APIResponse(c, http.StatusOK, true, "Succeed to POST data", product.ID, nil)
}

func UpdateProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := repository.GetProductByID(uint(id))
	
	if err != nil {
		utils.APIResponse(c, http.StatusNotFound, false, "Not Found", nil, nil)
		return
	}

	userID := c.MustGet("user_id").(uint)
	store, _ := repository.GetStoreByUserID(userID)
	if product.StoreID != store.ID {
		utils.APIResponse(c, http.StatusForbidden, false, "Forbidden", nil, nil)
		return
	}

	product.Name = c.PostForm("nama_produk")
	if val := c.PostForm("nama_produk"); val != "" { product.Name = val }
	
	repository.UpdateProduct(&product)
	utils.APIResponse(c, http.StatusOK, true, "Succeed to GET data", "", nil)
}

func DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := repository.GetProductByID(uint(id))
	
	if err != nil {
		utils.APIResponse(c, http.StatusBadRequest, false, "Failed to GET data", nil, []string{"record not found"})
		return
	}

	userID := c.MustGet("user_id").(uint)
	store, _ := repository.GetStoreByUserID(userID)
	if product.StoreID != store.ID {
		utils.APIResponse(c, http.StatusForbidden, false, "Forbidden", nil, nil)
		return
	}

	repository.DeleteProduct(product.ID)
	utils.APIResponse(c, http.StatusOK, true, "Succeed to GET data", "", nil)
}

// --- Transaction Handlers ---

func CreateTrx(c *gin.Context) {
	var input models.TrxRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.APIResponse(c, http.StatusBadRequest, false, "Failed", nil, []string{err.Error()})
		return
	}

	userID := c.MustGet("user_id").(uint)

	addr, err := repository.GetAddressByID(input.AlamatKirim)
	if err != nil || addr.UserID != userID {
		utils.APIResponse(c, http.StatusBadRequest, false, "Invalid Address", nil, nil)
		return
	}

	// Calculate Total strictly for header
	var total float64
	for _, item := range input.DetailTrx {
		prod, err := repository.GetProductByID(item.ProductID)
		if err != nil {
			utils.APIResponse(c, http.StatusBadRequest, false, "Product Unavailable", nil, nil)
			return
		}
		total += float64(item.Kuantitas) * prod.ConsumerPrice
	}

	trx := models.Transaction{
		UserID:        userID,
		AddressID:     input.AlamatKirim,
		TotalPrice:    total,
		InvoiceCode:   fmt.Sprintf("INV-%d", time.Now().Unix()),
		PaymentMethod: input.MethodBayar,
	}

	// Now passing slice directly because Repo accepts []models.TrxItemRequest
	if err := repository.CreateTransaction(&trx, input.DetailTrx); err != nil {
		utils.APIResponse(c, http.StatusInternalServerError, false, "Failed to create transaction", nil, []string{err.Error()})
		return
	}

	utils.APIResponse(c, http.StatusOK, true, "Succeed to POST data", len(input.DetailTrx), nil)
}

func GetTrxByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	trx, err := repository.GetTransactionByID(uint(id))
	
	if err != nil {
		utils.APIResponse(c, http.StatusNotFound, false, "Failed to GET data", nil, []string{"No Data Trx"})
		return
	}

	userID := c.MustGet("user_id").(uint)
	if trx.UserID != userID {
		utils.APIResponse(c, http.StatusForbidden, false, "Forbidden", nil, nil)
		return
	}

	utils.APIResponse(c, http.StatusOK, true, "Succeed to GET data", trx, nil)
}

func GetAllTrx(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	trxs, _ := repository.GetTransactionsByUserID(userID)
	utils.APIResponse(c, http.StatusOK, true, "Succeed to GET data", models.Pagination{Data: trxs}, nil)
}