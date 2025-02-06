package models

import (
	"AgringBackend/config"
	"time"
)

type Vehicle struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	VIN      string `json:"vin"`
	Category int    `json:"category" gorm:"foreignKey:ID"`
}

type VehicleCategory struct {
	ID           int    `json:"id" gorm:"primaryKey"`
	CategoryName string `json:"category_name"`
}

type VehicleStatus struct {
	ID              int       `json:"id" gorm:"primaryKey"`
	VehicleID       int       `json:"vehicle_id" gorm:"foreignKey:VehicleID"`
	VehicleStatusID int       `json:"vehicle_status_id" gorm:"foreignKey:ID"`
	EmployeeID      int       `json:"employee_id" gorm:"foreignKey:ID"`
	Action          string    `json:"action"`
	ActionTime      time.Time `json:"action_time"`
}

type User struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Username    string `json:"username" binding:"required"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

type UsersRoles struct {
	ID           int `json:"id" gorm:"primaryKey"`
	EmployeeID   int `json:"employee_id" gorm:"foreignKey:ID"`
	EmployeeRole int `json:"employee_role" gorm:"foreignKey:ID"`
}

type Role struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	RoleName string `json:"role_name"`
}

type WorkType struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null;unique"`
}

type FieldWorkers struct {
	EmployeeID int `json:"employee_id" gorm:"primaryKey; unique; not null; foreignKey:IDl; required"`
	FieldID    int `json:"field_id" gorm:"foreignKey:ID; required"`
	WorkType   int `json:"work_type" gorm:"foreignKey:ID; required"` // ✅ Тип int
}

type Field struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" binding:"required"`
	StatusID int    `json:"status_id" gorm:"foreignKey:ID"`
	CropID   int    `json:"crop_id" gorm:"foreignKey:ID"`
}

type FieldStatus struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	StatusName string `json:"status_name" binding:"required" gorm:"unique"`
}

type Crop struct {
	ID               int    `json:"id" gorm:"primaryKey"`
	Name             string `json:"name"`
	HarvestingMonth  int    `json:"harvesting_month"`
	IrrigationPeriod int    `json:"irrigation_period"`
	LandingMonth     int    `json:"landing_month"`
}

type VirtualFieldWorker struct {
	EmployeeName    string
	EmployeeSurname string
	FieldID         int
	WorkTypeName    string
}

func GetVirtualFieldWorkers() ([]VirtualFieldWorker, error) {
	var virtualWorkers []VirtualFieldWorker
	query := `
        SELECT u.name AS employee_name, u.surname AS employee_surname, fw.field_id, wt.name AS work_type_name
        FROM field_workers fw
        JOIN users u ON fw.employee_id = u.id
        JOIN work_types wt ON fw.work_type = wt.id
    `
	result := config.DB.Raw(query).Scan(&virtualWorkers)
	return virtualWorkers, result.Error
}
