package models

import "time"

type RegisterUser struct {
	Name        string `json:"name" db:"name" validate:"required,min=3,max=50"`
	Email       string `json:"email" db:"email" validate:"required,email"`
	Role        string `json:"role" db:"role" validate:"required,oneof=admin employee project_manager asset_manager employee_manager"`
	RoleType    string `json:"roleType" db:"type" validate:"required,oneof=full_time intern freelancer" `
	PhoneNumber string `json:"phoneNumber" db:"phone_no" validate:"required,len=10"`
	Password    string `json:"password" db:"password" validate:"required,min=8,max=20"`
}

type LoginUser struct {
	Email    string `json:"email" db:"email" validate:"email"`
	Password string `json:"password" db:"password" validate:"required,alphanum,min=8,max=20"`
}

type LoginData struct {
	UserID       string `db:"user_id"`
	PasswordHash string `db:"password"`
}

type User struct {
	UserID      string    `json:"id" db:"user_id"`
	Name        string    `json:"name" db:"name"`
	Email       string    `json:"email" db:"email"`
	Role        string    `json:"role" db:"role" validate:"required,oneof=admin employee project_manager asset_manager employee_manager"`
	RoleType    string    `json:"roleType" db:"type" validate:"required,oneof=full_time intern freelancer" `
	PhoneNumber string    `json:"phoneNumber" db:"phone_no" validate:"required,len=10"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
}
