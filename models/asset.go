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
	Processor       string `json:"processor"`
	Ram             string `json:"ram" `
	Storage         string `json:"storage"`
	OperatingSystem string `json:"operatingSystem" `
	Charger         bool   `json:"charger" `
}
type KeyboardSpecsRequest struct {
	Layout         string `json:"layout" `
	ConnectionType string `json:"connectionType" `
}
type MouseSpecsRequest struct {
	Dpi            int    `json:"dpi" `
	ConnectionType string `json:"connectionType"`
}
type MobileSpecsRequest struct {
	OperatingSystem string `json:"operatingSystem" `
	Ram             string `json:"ram" `
	Storage         string `json:"storage" `
	Charger         bool   `json:"charger" `
}

type UpdateAssetRequest struct {
	Brand         *string    `json:"brand"`
	Model         *string    `json:"model"`
	SerialNumber  *string    `json:"serialNumber"`
	Status        *string    `json:"status" validate:"omitempty,oneof=available assigned needs_repair under_repair damaged"`
	OwnerType     *string    `json:"owner" validate:"omitempty,oneof=client remotestate"`
	WarrantyStart *time.Time `json:"warrantyStart"`
	WarrantyEnd   *time.Time `json:"warrantyEnd"`

	Laptop   *LaptopSpecsRequest   `json:"laptop,omitempty"`
	Keyboard *KeyboardSpecsRequest `json:"keyboard,omitempty"`
	Mouse    *MouseSpecsRequest    `json:"mouse,omitempty"`
	Mobile   *MobileSpecsRequest   `json:"mobile,omitempty"`
}

type RepairAssetRequest struct {
	AssetID string `json:"assetID" validate:"required,uuid"`
}
