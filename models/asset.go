package models

import "time"

type AssetRequest struct {
	Brand         string               `json:"brand" db:"brand" `
	Model         string               `json:"model" db:"model" `
	SerialNumber  string               `json:"serialNumber" db:"serial_number" `
	AssetType     string               `json:"assetType" db:"asset_type" validate:"required,oneof=laptop keyboard mouse mobile"`
	Status        string               `json:"status" db:"status" validate:"required,oneof=available assigned needs_repair under_repair damaged"`
	OwnerType     string               `json:"owner" db:"owner_type" validate:"required,oneof=client remotestate"`
	WarrantyStart time.Time            `json:"warrantyStart" db:"warranty_start" `
	WarrantyEnd   time.Time            `json:"warrantyEnd" db:"warranty_end" `
	Laptop        LaptopSpecsRequest   `json:"laptopSpecs,omitempty"`
	Keyboard      KeyboardSpecsRequest `json:"keyboardSpecs,omitempty"`
	Mouse         MouseSpecsRequest    `json:"mouseSpecs,omitempty"`
	Mobile        MobileSpecsRequest   `json:"mobileSpecs,omitempty"`
}
type LaptopSpecsRequest struct {
	Processor       string `json:"processor"`
	Ram             string `json:"ram" `
	Storage         string `json:"storage"`
	OperatingSystem string `json:"operatingSystem" `
	Charger         bool   `json:"charger" `
}
type LaptopSpecs struct {
	AssetID         string `json:"assetID" db:"asset_id"`
	Processor       string `json:"processor" db:"processor"`
	Ram             string `json:"ram" db:"ram"`
	Storage         string `json:"storage" db:"storage"`
	OperatingSystem string `json:"operatingSystem" db:"operating_system"`
	Charger         bool   `json:"charger" db:"charger"`
}

type KeyboardSpecsRequest struct {
	Layout         string `json:"layout" `
	ConnectionType string `json:"connectionType" `
}
type KeyboardSpecs struct {
	AssetID        string `json:"assetID" db:"asset_id"`
	Layout         string `json:"layout" db:"layout"`
	ConnectionType string `json:"connectionType" db:"connectivity"`
}

type MouseSpecsRequest struct {
	Dpi            int    `json:"dpi" `
	ConnectionType string `json:"connectionType"`
}
type MouseSpecs struct {
	AssetID        string `json:"assetID" db:"asset_id"`
	Dpi            int    `json:"dpi" db:"dpi"`
	ConnectionType string `json:"connectionType" db:"connectivity"`
}

type MobileSpecsRequest struct {
	OperatingSystem string `json:"operatingSystem" `
	Ram             string `json:"ram" `
	Storage         string `json:"storage" `
	Charger         bool   `json:"charger" `
}
type MobileSpecs struct {
	AssetID         string `json:"assetID" db:"asset_id"`
	OperatingSystem string `json:"operatingSystem" db:"operating_system"`
	Ram             string `json:"ram" db:"ram"`
	Storage         string `json:"storage" db:"storage"`
	Charger         bool   `json:"charger" db:"charger"`
}
