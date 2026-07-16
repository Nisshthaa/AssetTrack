package models

import "time"

type AssetRequest struct {
	Brand         string               `json:"brand" `
	Model         string               `json:"model"  `
	SerialNumber  string               `json:"serialNumber"  `
	AssetType     string               `json:"assetType" validate:"required,oneof=laptop keyboard mouse mobile"`
	Status        string               `json:"status" validate:"required,oneof=available assigned needs_repair under_repair damaged"`
	OwnerType     string               `json:"owner" validate:"required,oneof=client remotestate"`
	WarrantyStart time.Time            `json:"warrantyStart"  `
	WarrantyEnd   time.Time            `json:"warrantyEnd" `
	Laptop        LaptopSpecsRequest   `json:"laptopSpecs,omitempty"`
	Keyboard      KeyboardSpecsRequest `json:"keyboardSpecs,omitempty"`
	Mouse         MouseSpecsRequest    `json:"mouseSpecs,omitempty"`
	Mobile        MobileSpecsRequest   `json:"mobileSpecs,omitempty"`
}

type AssetDetails struct {
	AssetID       string    `json:"assetID" db:"asset_id"`
	SerialNumber  string    `json:"serialNumber" db:"serial_number"`
	Brand         string    `json:"brand" db:"brand"`
	Model         string    `json:"model" db:"model"`
	AssetType     string    `json:"assetType" db:"asset_type"`
	Status        string    `json:"status" db:"status"`
	OwnerType     string    `json:"owner" db:"owner_type"`
	AssignedTo    string    `json:"assignedTo" db:"assigned_to"`
	WarrantyStart time.Time `json:"warrantyStart" db:"warranty_start" `
	WarrantyEnd   time.Time `json:"warrantyEnd" db:"warranty_end"`

	Laptop   *LaptopSpecsRequest   `json:"laptopSpecs,omitempty"`
	Keyboard *KeyboardSpecsRequest `json:"keyboardSpecs,omitempty"`
	Mouse    *MouseSpecsRequest    `json:"mouseSpecs,omitempty"`
	Mobile   *MobileSpecsRequest   `json:"mobileSpecs,omitempty"`
}
type LaptopSpecsRequest struct {
	Processor       *string `json:"processor" db:"processor"`
	Ram             *string `json:"ram" db:"ram" `
	Storage         *string `json:"storage" db:"storage"`
	OperatingSystem *string `json:"operatingSystem"  db:"operating_system"`
	Charger         *bool   `json:"charger" db:"charger"`
}
type KeyboardSpecsRequest struct {
	Layout         *string `json:"layout" db:"layout"`
	ConnectionType *string `json:"connectionType" db:"connection_type"`
}
type MouseSpecsRequest struct {
	Dpi            *int    `json:"dpi" db:"dpi"`
	ConnectionType *string `json:"connectionType" db:"connection_type"`
}
type MobileSpecsRequest struct {
	OperatingSystem *string `json:"operatingSystem" db:"layout"`
	Ram             *string `json:"ram" db:"ram"`
	Storage         *string `json:"storage" db:"storage"`
	Charger         *bool   `json:"charger" db:"layout"`
}

type UpdateAssetRequest struct {
	Brand         *string               `json:"brand"`
	Model         *string               `json:"model"`
	SerialNumber  *string               `json:"serialNumber"`
	Status        *string               `json:"status" validate:"omitempty,oneof=available assigned needs_repair under_repair damaged"`
	OwnerType     *string               `json:"owner" validate:"omitempty,oneof=client remotestate"`
	WarrantyStart *time.Time            `json:"warrantyStart"`
	WarrantyEnd   *time.Time            `json:"warrantyEnd"`
	Laptop        *LaptopSpecsRequest   `json:"laptopSpecs,omitempty"`
	Keyboard      *KeyboardSpecsRequest `json:"keyboardSpecs,omitempty"`
	Mouse         *MouseSpecsRequest    `json:"mouseSpecs,omitempty"`
	Mobile        *MobileSpecsRequest   `json:"mobileSpecs,omitempty"`
}

type RepairAssetRequest struct {
	AssetID string `json:"assetID" validate:"required,uuid"`
}
