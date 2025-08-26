package usercontroller

import (
	"crud-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /users
func Index(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"users": users})
}

// GET /users/:id
func Show(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if err := models.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// POST /users
func Create(c *gin.Context) {
	var input models.User

	// Bind JSON ke struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Simpan user baru
	user := models.User{
		Email:    input.Email,
		Username: input.Username,
		Name:     input.Name,
		Password: input.Password, // ⚠️ sebaiknya di-hash sebelum simpan
	}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// PUT /users/:id
func Update(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	// Cari user
	if err := models.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Bind data baru
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update field
	models.DB.Model(&user).Updates(models.User{
		Email:    input.Email,
		Username: input.Username,
		Name:     input.Name,
		Password: input.Password,
	})

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// DELETE /users/:id
func Delete(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	// Cari user
	if err := models.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Hapus user
	models.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}