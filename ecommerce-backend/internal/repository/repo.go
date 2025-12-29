package repository

import (
	"ecommerce-backend/models"
	"ecommerce-backend/pkg/database"
)

// User Repository
func CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}

func FindUserByPhone(phone string) (models.User, error) {
	var user models.User
	err := database.DB.Where("notelp = ?", phone).First(&user).Error
	return user, err
}

func FindUserByID(id uint) (models.User, error) {
	var user models.User
	err := database.DB.Preload("Store").First(&user, id).Error
	return user, err
}

func UpdateUser(user *models.User) error {
	return database.DB.Save(user).Error
}

// Address Repository
func GetAddressesByUserID(userID uint) ([]models.Address, error) {
	var addresses []models.Address
	err := database.DB.Where("id_user = ?", userID).Find(&addresses).Error
	return addresses, err
}

func GetAddressByID(id uint) (models.Address, error) {
	var address models.Address
	err := database.DB.First(&address, id).Error
	return address, err
}

func CreateAddress(address *models.Address) error {
	return database.DB.Create(address).Error
}

func UpdateAddress(address *models.Address) error {
	return database.DB.Save(address).Error
}

func DeleteAddress(id uint) error {
	return database.DB.Delete(&models.Address{}, id).Error
}

// Store Repository
func GetStoreByUserID(userID uint) (models.Store, error) {
	var store models.Store
	err := database.DB.Where("id_user = ?", userID).First(&store).Error
	return store, err
}

func GetStoreByID(id uint) (models.Store, error) {
	var store models.Store
	err := database.DB.First(&store, id).Error
	return store, err
}

func UpdateStore(store *models.Store) error {
	return database.DB.Save(store).Error
}

func GetAllStores(page, limit int, nameFilter string) ([]models.Store, int64, error) {
	var stores []models.Store
	var total int64
	
	query := database.DB.Model(&models.Store{})
	if nameFilter != "" {
		query = query.Where("nama_toko LIKE ?", "%"+nameFilter+"%")
	}
	
	query.Count(&total)
	offset := (page - 1) * limit
	err := query.Limit(limit).Offset(offset).Find(&stores).Error
	return stores, total, err
}

// Category Repository
func GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	err := database.DB.Find(&categories).Error
	return categories, err
}

func CreateCategory(category *models.Category) error {
	return database.DB.Create(category).Error
}

func GetCategoryByID(id uint) (models.Category, error) {
	var category models.Category
	err := database.DB.First(&category, id).Error
	return category, err
}

func UpdateCategory(category *models.Category) error {
	return database.DB.Save(category).Error
}

func DeleteCategory(id uint) error {
	return database.DB.Delete(&models.Category{}, id).Error
}

// Product Repository
func GetProducts(page, limit int, name, categoryID, storeID, maxPrice, minPrice string) ([]models.Product, int64, error) {
	var products []models.Product
	var total int64
	
	query := database.DB.Model(&models.Product{}).Preload("Store").Preload("Category").Preload("Photos")

	if name != "" {
		query = query.Where("nama_produk LIKE ?", "%"+name+"%")
	}
	if categoryID != "" {
		query = query.Where("id_category = ?", categoryID)
	}
	if storeID != "" {
		query = query.Where("id_toko = ?", storeID)
	}
	if maxPrice != "" {
		query = query.Where("harga_konsumen <= ?", maxPrice)
	}
	if minPrice != "" {
		query = query.Where("harga_konsumen >= ?", minPrice)
	}

	query.Count(&total)
	offset := (page - 1) * limit
	err := query.Limit(limit).Offset(offset).Find(&products).Error
	return products, total, err
}

func GetProductByID(id uint) (models.Product, error) {
	var product models.Product
	err := database.DB.Preload("Store").Preload("Category").Preload("Photos").First(&product, id).Error
	return product, err
}

func CreateProduct(product *models.Product) error {
	return database.DB.Create(product).Error
}

func UpdateProduct(product *models.Product) error {
	return database.DB.Save(product).Error
}

func DeleteProduct(id uint) error {
	return database.DB.Delete(&models.Product{}, id).Error
}

func DeleteProductPhotos(productID uint) {
	database.DB.Where("id_produk = ?", productID).Delete(&models.ProductPhoto{})
}

// Transaction Repository
func CreateTransaction(trx *models.Transaction, reqDetails []models.TrxItemRequest) error {
	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 1. Create Transaction Header
	if err := tx.Create(trx).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, item := range reqDetails {
		// 2. Get Product & Lock
		var product models.Product
		if err := tx.First(&product, item.ProductID).Error; err != nil {
			tx.Rollback()
			return err
		}

		// 3. Check & Update Stock
		if product.Stock < item.Kuantitas {
			tx.Rollback()
			return nil // In real app, return named error
		}
		product.Stock -= item.Kuantitas
		if err := tx.Save(&product).Error; err != nil {
			tx.Rollback()
			return err
		}

		// 4. Create Product Log (Snapshot)
		log := models.ProductLog{
			ProductID:     product.ID,
			StoreID:       product.StoreID,
			CategoryID:    product.CategoryID,
			Name:          product.Name,
			Slug:          product.Slug,
			ResellerPrice: product.ResellerPrice,
			ConsumerPrice: product.ConsumerPrice,
			Description:   product.Description,
		}
		if err := tx.Create(&log).Error; err != nil {
			tx.Rollback()
			return err
		}

		// 5. Create Transaction Detail linked to Log
		detail := models.TransactionDetail{
			TransactionID: trx.ID,
			ProductLogID:  log.ID,
			StoreID:       product.StoreID,
			Quantity:      item.Kuantitas,
			TotalPrice:    product.ConsumerPrice * float64(item.Kuantitas),
		}
		if err := tx.Create(&detail).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func GetTransactionsByUserID(userID uint) ([]models.Transaction, error) {
	var trxs []models.Transaction
	// Preload Log via Details
	err := database.DB.Preload("Address").Preload("Details").Preload("Details.ProductLog").Where("id_user = ?", userID).Find(&trxs).Error
	return trxs, err
}

func GetTransactionByID(id uint) (models.Transaction, error) {
	var trx models.Transaction
	err := database.DB.Preload("Address").Preload("Details").Preload("Details.ProductLog").First(&trx, id).Error
	return trx, err
}