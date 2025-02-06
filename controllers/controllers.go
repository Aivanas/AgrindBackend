package controllers

import (
	"AgringBackend/config"
	"AgringBackend/models"
	"AgringBackend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetOK(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}

func GetVehicles(c *gin.Context) {
	var vehicles []models.Vehicle
	config.DB.Find(&vehicles)
	c.JSON(http.StatusOK, vehicles)
}

func GetVehicle(c *gin.Context) {
	id := c.Param("id")
	var vehicle models.Vehicle
	if err := config.DB.First(&vehicle, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Status not found"})
		return
	}
	c.JSON(http.StatusOK, vehicle)
}

func CreateVehicle(c *gin.Context) {
	var vehicle models.Vehicle
	if err := c.ShouldBindJSON(&vehicle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&vehicle)
	c.JSON(http.StatusCreated, vehicle)
}

func UpdateVehicle(c *gin.Context) {
	id := c.Param("id")
	var vehicle models.Vehicle
	if err := config.DB.First(&vehicle, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vehicle not found"})
		return
	}
	if err := c.ShouldBindJSON(&vehicle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&vehicle)
	c.JSON(http.StatusOK, vehicle)
}

func DeleteVehicle(c *gin.Context) {
	id := c.Param("id")
	var vehicle models.Vehicle
	if err := config.DB.First(&vehicle, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vehicle not found"})
	}
	config.DB.Delete(&vehicle)
	c.JSON(http.StatusNoContent, nil)
}

func GetVehicleCategories(c *gin.Context) {
	var categories []models.VehicleCategory
	config.DB.Find(&categories)
	c.JSON(http.StatusOK, categories)
}

func CreateVehicleCategory(c *gin.Context) {
	var category models.VehicleCategory
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&category)
	c.JSON(http.StatusCreated, category)
}

func GetVehicleCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.VehicleCategory
	if err := config.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	c.JSON(http.StatusOK, category)
}

func UpdateVehicleCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.VehicleCategory
	if err := config.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&category)
	c.JSON(http.StatusOK, category)
}

func DeleteVehicleCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.VehicleCategory
	if err := config.DB.First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	config.DB.Delete(&category)
	c.JSON(http.StatusNoContent, nil)
}

func GetVehicleStatuses(c *gin.Context) {
	var statuses []models.VehicleStatus
	config.DB.Find(&statuses)
	c.JSON(http.StatusOK, statuses)
}

func CreateVehicleStatus(c *gin.Context) {
	var status models.VehicleStatus
	if err := c.ShouldBindJSON(&status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&status)
	c.JSON(http.StatusCreated, status)
}

func GetVehicleStatus(c *gin.Context) {
	id := c.Param("id")
	var status models.VehicleStatus
	if err := config.DB.First(&status, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Status not found"})
		return
	}
	c.JSON(http.StatusOK, status)
}

func UpdateVehicleStatus(c *gin.Context) {
	id := c.Param("id")
	var status models.VehicleStatus
	if err := config.DB.First(&status, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Status not found"})
		return
	}
	if err := c.ShouldBindJSON(&status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&status)
	c.JSON(http.StatusOK, status)
}

func DeleteVehicleStatus(c *gin.Context) {
	id := c.Param("id")
	var status models.VehicleStatus
	if err := config.DB.First(&status, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Status not found"})
		return
	}
	config.DB.Delete(&status)
	c.JSON(http.StatusNoContent, nil)
}

func GetUsers(c *gin.Context) {
	var Users []models.User
	config.DB.Find(&Users)
	c.JSON(http.StatusOK, Users)
}

func CreateUser(c *gin.Context) {
	var User models.User
	if err := c.ShouldBindJSON(&User); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&User)
	c.JSON(http.StatusCreated, User)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	var User models.User
	if err := config.DB.First(&User, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, User)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var User models.User
	if err := config.DB.First(&User, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	if err := c.ShouldBindJSON(&User); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&User)
	c.JSON(http.StatusOK, User)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var User models.User
	if err := config.DB.First(&User, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	config.DB.Delete(&User)
	c.JSON(http.StatusNoContent, nil)
}

func GetUsersRoless(c *gin.Context) {
	var roles []models.UsersRoles
	config.DB.Find(&roles)
	c.JSON(http.StatusOK, roles)
}

func CreateUsersRoles(c *gin.Context) {
	var role models.UsersRoles
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&role)
	c.JSON(http.StatusCreated, role)
}

func GetUsersRoles(c *gin.Context) {
	id := c.Param("id")
	var role models.UsersRoles
	if err := config.DB.First(&role, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}
	c.JSON(http.StatusOK, role)
}

func UpdateUsersRoles(c *gin.Context) {
	id := c.Param("id")
	var role models.UsersRoles
	if err := config.DB.First(&role, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&role)
	c.JSON(http.StatusOK, role)
}

func DeleteUsersRoles(c *gin.Context) {
	id := c.Param("id")
	var role models.UsersRoles
	if err := config.DB.First(&role, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}
	config.DB.Delete(&role)
	c.JSON(http.StatusNoContent, nil)
}

func GetRoles(c *gin.Context) {
	var roles []models.Role
	config.DB.Find(&roles)
	c.JSON(http.StatusOK, roles)
}

func CreateRole(c *gin.Context) {
	var role models.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&role)
	c.JSON(http.StatusCreated, role)
}

func GetRole(c *gin.Context) {
	id := c.Param("id")
	var role models.Role
	if err := config.DB.First(&role, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}
	c.JSON(http.StatusOK, role)
}

func UpdateRole(c *gin.Context) {
	id := c.Param("id")
	var role models.Role
	if err := config.DB.First(&role, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&role)
	c.JSON(http.StatusOK, role)
}

func DeleteRole(c *gin.Context) {
	id := c.Param("id")
	var role models.Role
	if err := config.DB.First(&role, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}
	config.DB.Delete(&role)
	c.JSON(http.StatusNoContent, nil)
}

func GetFields(c *gin.Context) {
	var fields []models.Field
	config.DB.Find(&fields)
	c.JSON(http.StatusOK, fields)
}

func CreateField(c *gin.Context) {
	var field models.Field
	if err := c.ShouldBindJSON(&field); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&field)
	c.JSON(http.StatusCreated, field)
}

func GetField(c *gin.Context) {
	id := c.Param("id")
	var field models.Field
	if err := config.DB.First(&field, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Field not found"})
		return
	}
	c.JSON(http.StatusOK, field)
}

func UpdateField(c *gin.Context) {
	id := c.Param("id")
	var field models.Field
	if err := config.DB.First(&field, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Field not found"})
		return
	}
	if err := c.ShouldBindJSON(&field); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&field)
	c.JSON(http.StatusOK, field)
}

func DeleteField(c *gin.Context) {
	id := c.Param("id")
	var field models.Field
	if err := config.DB.First(&field, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Field not found"})
		return
	}
	config.DB.Delete(&field)
	c.JSON(http.StatusNoContent, nil)
}

func GetFieldWorkers(c *gin.Context) {
	var infos []models.FieldWorkers
	config.DB.Find(&infos)
	c.JSON(http.StatusOK, infos)
}

func CreateFieldWorkers(c *gin.Context) {
	var info models.FieldWorkers
	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&info)
	c.JSON(http.StatusCreated, info)
}

func GetFieldsWorker(c *gin.Context) {
	id := c.Param("id")
	var info models.FieldWorkers
	if err := config.DB.First(&info, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "FieldStatusInfo not found"})
		return
	}
	c.JSON(http.StatusOK, info)
}

func UpdateFieldWorker(c *gin.Context) {
	id := c.Param("id")
	var info models.FieldWorkers
	if err := config.DB.First(&info, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "FieldStatusInfo not found"})
		return
	}
	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&info)
	c.JSON(http.StatusOK, info)
}

func DeleteFieldWorker(c *gin.Context) {
	id := c.Param("id")
	var info models.FieldWorkers
	if err := config.DB.First(&info, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "FieldStatusInfo not found"})
		return
	}
	config.DB.Delete(&info)
	c.JSON(http.StatusNoContent, nil)
}

func GetFieldStatuses(c *gin.Context) {
	var lists []models.FieldStatus
	config.DB.Find(&lists)
	c.JSON(http.StatusOK, lists)
}

func CreateFieldStatus(c *gin.Context) {
	var list models.FieldStatus
	if err := c.ShouldBindJSON(&list); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&list)
	c.JSON(http.StatusCreated, list)
}

func GetFieldStatus(c *gin.Context) {
	id := c.Param("id")
	var list models.FieldStatus
	if err := config.DB.First(&list, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "FieldStatus not found"})
		return
	}
	c.JSON(http.StatusOK, list)
}

func UpdateFieldStatus(c *gin.Context) {
	id := c.Param("id")
	var list models.FieldStatus
	if err := config.DB.First(&list, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "FieldStatus not found"})
		return
	}
	if err := c.ShouldBindJSON(&list); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Save(&list)
	c.JSON(http.StatusOK, list)
}

func DeleteFieldStatus(c *gin.Context) {
	id := c.Param("id")
	var list models.FieldStatus
	if err := config.DB.First(&list, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "FieldStatus not found"})
		return
	}
	config.DB.Delete(&list)
	c.JSON(http.StatusNoContent, nil)
}

func GetCrops(c *gin.Context) {
	var crops []models.Crop
	config.DB.Find(&crops)
	c.JSON(http.StatusOK, crops)
}

func CreateCrop(c *gin.Context) {
	var crop models.Crop
	if err := c.ShouldBindJSON(&crop); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&crop)
	c.JSON(http.StatusCreated, crop)
}

func GetCrop(c *gin.Context) {
	id := c.Param("id")
	var crop models.Crop
	if err := config.DB.First(&crop, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Crop not found"})
		return
	}
	c.JSON(http.StatusOK, crop)
}

func UpdateCrop(c *gin.Context) {
	id := c.Param("id")
	var crop models.Crop
	if err := config.DB.First(&crop, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Crop not found"})
	}
}

func DeleteCrop(c *gin.Context) {
	id := c.Param("id")
	var crop models.Crop
	if err := config.DB.First(&crop, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "FieldStatusInfo not found"})
		return
	}
	config.DB.Delete(&crop)
	c.JSON(http.StatusNoContent, nil)
}

func Register(c *gin.Context) {
	var user models.User
	fmt.Printf("Received data: %+v\n", user)
	// Универсальный биндинг (работает и с JSON, и с form-data)
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// Хэшируем пароль
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
		return
	}
	user.Password = hashedPassword

	// Сохраняем пользователя в базе данных
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	var user models.User
	var loginData struct {
		Username string `form:"username" json:"username" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}

	// Универсальный биндинг (работает и с JSON, и с form-data)
	if err := c.ShouldBind(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	// Ищем пользователя в базе данных
	if err := config.DB.Where("username = ?", loginData.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	// Проверяем пароль
	if !utils.CheckPasswordHash(loginData.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	// Генерируем JWT
	token, err := utils.GenerateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func GetWorkTypes(c *gin.Context) {
	var workTypes []models.WorkType
	if err := config.DB.Find(&workTypes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve work types"})
		return
	}
	c.JSON(http.StatusOK, workTypes)
}

// Получение одного типа работы по ID
func GetWorkTypeByID(c *gin.Context) {
	id := c.Param("id")
	var workType models.WorkType
	if err := config.DB.First(&workType, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Work type not found"})
		return
	}
	c.JSON(http.StatusOK, workType)
}

// Создание нового типа работы
func CreateWorkType(c *gin.Context) {
	var workType models.WorkType
	if err := c.ShouldBindJSON(&workType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	if err := config.DB.Create(&workType).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create work type"})
		return
	}
	c.JSON(http.StatusCreated, workType)
}

// Обновление типа работы по ID
func UpdateWorkType(c *gin.Context) {
	id := c.Param("id")
	var workType models.WorkType

	// Проверяем, существует ли запись
	if err := config.DB.First(&workType, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Work type not found"})
		return
	}

	// Применяем изменения
	if err := c.ShouldBindJSON(&workType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	// Сохраняем обновлённую запись
	if err := config.DB.Save(&workType).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update work type"})
		return
	}

	c.JSON(http.StatusOK, workType)
}

// Удаление типа работы по ID
func DeleteWorkType(c *gin.Context) {
	id := c.Param("id")
	var workType models.WorkType

	// Проверяем, существует ли запись
	if err := config.DB.First(&workType, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Work type not found"})
		return
	}

	// Удаляем запись
	if err := config.DB.Delete(&workType).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete work type"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Work type deleted successfully"})
}
