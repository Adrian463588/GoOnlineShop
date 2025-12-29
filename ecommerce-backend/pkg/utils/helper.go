package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var SecretKey = []byte("supersecretkey") // Use ENV in production

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateToken(userID uint, isAdmin bool) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  userID,
		"is_admin": isAdmin,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SecretKey)
}

func APIResponse(ctx *gin.Context, code int, status bool, message string, data interface{}, errors interface{}) {
	ctx.JSON(code, gin.H{
		"status":  status,
		"message": message,
		"errors":  errors,
		"data":    data,
	})
}

// File Upload Helper
func SaveUploadedFile(ctx *gin.Context, paramName string) (string, error) {
	file, err := ctx.FormFile(paramName)
	if err != nil {
		return "", err
	}

	// Create uploads directory if not exists
	uploadPath := "public/uploads"
	if _, err := os.Stat(uploadPath); os.IsNotExist(err) {
		os.MkdirAll(uploadPath, 0755)
	}

	// Generate unique filename
	filename := fmt.Sprintf("%d-%s", time.Now().Unix(), filepath.Base(file.Filename))
	dst := filepath.Join(uploadPath, filename)

	if err := ctx.SaveUploadedFile(file, dst); err != nil {
		return "", err
	}
	
	return filename, nil
}

// Slugify Helper (Added this function)
func Slugify(s string) string {
	return strings.ToLower(strings.ReplaceAll(s, " ", "-"))
}