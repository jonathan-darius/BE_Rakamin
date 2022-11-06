package controllers

import (
	"BE_Rakamin/databases"
	"BE_Rakamin/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	sqlClient = databases.DBInit()
)

// CreatePerson godoc
// @Summary      Add Person
// @Description  post Person
// @Tags         UserManagement
// @Accept       json
// @Produce      json
// @Param        user body	models.SwagPerson true "Register Person"
// @Router       /AddPerson [post]
func CreatePerson() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			user models.Person
		)
		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := sqlClient.Create(&user).Error; err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{
				"Status":  "Failed",
				"Message": err,
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"Status":  "Success",
			"Data":    &user,
			"Message": "OK",
		})
	}
}

// GetUser godoc
// @Summary      Get a User
// @Description  Get a User
// @Tags         UserManagement
// @Produce      json
// @Param		 user_id  path      string  true  "search user by id"
// @Router       /GetUser/{user_id} [get]
func GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var person models.Person
		userId := ctx.Param("user_id")
		if err := sqlClient.Where("id = ?", userId).First(&person).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Status":  "Failed",
				"Message": "Data Not Found",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"Status":  "Success",
			"Data":    &person,
			"Message": "OK",
		})
	}
}

// GetAllUser godoc
// @Summary      All User
// @Description  All User
// @Tags         UserManagement
// @Produce      json
// @Router      /GetUsers [get]
func GetAllUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var persons []models.Person
		if err := sqlClient.Find(&persons).Error; err != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{
				"Status":  "Failed",
				"Message": "Failed Request Data",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"Status":  "Success",
			"Total":   len(persons),
			"Data":    &persons,
			"Message": "OK",
		})
	}
}

// DeleteUser godoc
// @Summary      Delete a User
// @Description  Delete a User
// @Tags         UserManagement
// @Produce      json
// @Param		 user_id  path      string  true  "Delete user by id"
// @Router       /DeleteUser/{user_id} [delete]
func DeleteUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.Param("user_id")

		if err := sqlClient.First(&models.Person{}, userId).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Status":  "Failed",
				"Message": "Data Not Found",
			})
			return
		}
		if err := sqlClient.Delete(&models.Person{}, userId).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Status":  "Failed",
				"Message": "Failed to Delete",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"Status":  "Success",
			"Message": "OK",
		})
	}
}

// UpdateUser godoc
// @Summary      Update a User
// @Description  Update a User
// @Tags         UserManagement
// @Accept       json
// @Produce      json
// @Param		 user_id  path      string  true  "Update user by id"
// @Param        user body	models.SwagPerson true "Update Person"
// @Router       /UpdateUser/{user_id} [patch]
func UpdateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.Param("user_id")
		var (
			Olduser models.Person
			user    models.Person
		)

		if err := sqlClient.First(&Olduser, userId).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Status":  "Failed",
				"Message": "Data Not Found",
			})
			return
		}

		if err := ctx.BindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := sqlClient.Model(&Olduser).Updates(&user).Error; err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Status":  "Failed",
				"Message": "Error Update",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"Status":  "Success",
			"data":    sqlClient.First(&models.Person{}, userId),
			"Message": "OK",
		})
	}
}
