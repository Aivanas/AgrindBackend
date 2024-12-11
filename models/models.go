package models

import "time"

type Vehicle struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	VIN      string `json:"vin"`
	Category int    `json:"category"`
}

type VehicleCategory struct {
	ID           int    `json:"id" gorm:"primaryKey"`
	CategoryName string `json:"category_name"`
}

type VehicleStatus struct {
	ID              int       `json:"id" gorm:"primaryKey"`
	VehicleID       int       `json:"vehicle_id"`
	VehicleStatusID int       `json:"vehicle_status_id"`
	EmployeeID      int       `json:"employee_id"`
	Action          string    `json:"action"`
	ActionTime      time.Time `json:"action_time"`
}

type Employee struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

type EmployeeRole struct {
	ID           int `json:"id" gorm:"primaryKey"`
	EmployeeID   int `json:"employee_id"`
	EmployeeRole int `json:"employee_role"`
}

type Role struct {
	ID       int    `json:"id" gorm:"primaryKey"`
	RoleName string `json:"role_name"`
}

type Field struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

type FieldStatusInfo struct {
	ID         int `json:"id" gorm:"primaryKey"`
	FieldID    int `json:"field_id"`
	StatusID   int `json:"status_id"`
	EmployeeID int `json:"employee_id"`
	CropID     int `json:"crop_id"`
}

type FieldStatusList struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	StatusName string `json:"status_name"`
}

type Crop struct {
	ID               int    `json:"id" gorm:"primaryKey"`
	Name             string `json:"name"`
	HarvestingMonth  int    `json:"harvesting_month"`
	IrrigationPeriod int    `json:"irrigation_period"`
	LandingMonth     int    `json:"landing_month"`
}
