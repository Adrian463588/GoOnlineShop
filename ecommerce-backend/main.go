package main

import (
	"ecommerce-backend/internal/handler"
	"ecommerce-backend/pkg/database"
	"ecommerce-backend/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()
	r.Static("/public", "./public")

	api := r.Group("/api/v1")
	{
		// Auth
		api.POST("/auth/register", handler.Register)
		api.POST("/auth/login", handler.Login)

		// Public Product
		api.GET("/product", handler.GetAllProducts)
		api.GET("/product/:id", handler.GetProductByID)

		// Public Category
		api.GET("/category", handler.GetAllCategory)
		api.GET("/category/:id", handler.GetCategoryByID)

		// Public Store
		api.GET("/toko", handler.GetAllStores)
		api.GET("/toko/:id_toko", handler.GetStoreByID)

		// Protected Routes
		authorized := api.Group("/")
		authorized.Use(middleware.AuthMiddleware())
		{
			// User
			authorized.GET("/user", handler.GetProfile)
			authorized.PUT("/user", handler.UpdateProfile)
			
			// Alamat
			authorized.GET("/user/alamat", handler.GetMyAddress)
			authorized.GET("/user/alamat/:id", handler.GetAddressByID)
			authorized.POST("/user/alamat", handler.CreateAddress)
			authorized.PUT("/user/alamat/:id", handler.UpdateAddress)
			authorized.DELETE("/user/alamat/:id", handler.DeleteAddress)

			// Store Management (My Store)
			authorized.GET("/toko/my", handler.GetMyStore)
			authorized.PUT("/toko/:id_toko", handler.UpdateStore)

			// Product Management
			authorized.POST("/product", handler.CreateProduct)
			authorized.PUT("/product/:id", handler.UpdateProduct)
			authorized.DELETE("/product/:id", handler.DeleteProduct)

			// Transaction
			authorized.GET("/trx", handler.GetAllTrx)
			authorized.GET("/trx/:id", handler.GetTrxByID)
			authorized.POST("/trx", handler.CreateTrx)

			// Admin Only
			admin := authorized.Group("/")
			admin.Use(middleware.AdminOnly())
			{
				admin.POST("/category", handler.CreateCategory)
				admin.PUT("/category/:id", handler.UpdateCategory)
				admin.DELETE("/category/:id", handler.DeleteCategory)
			}
		}
		
		// ProvCity (Placeholder for regional data)
		api.GET("/provcity/listprovincies", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": true, "data": []string{"Data fetched from external API"}})
		})
	}

	r.Run(":8000")
}