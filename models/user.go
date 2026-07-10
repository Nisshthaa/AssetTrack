package models

type RegisterUser struct {
	Name        string `json:"name" db:"name" validate:"required,min=3,max=50"`
	Email       string `json:"email" db:"email" validate:"required,email"`
	Role        string `json:"role" db:"role" validate:"required,oneof=admin employee project_manager asset_manager employee_manager"`
	RoleType    string `json:"roleType" db:"type" validate:"required,oneof=full_time intern freelancer" `
	PhoneNumber string `json:"phoneNumber" db:"phone_number" validate:"required,len=10"`
	Password    string `json:"password" db:"password" validate:"required,min=8,max=20"`
}
