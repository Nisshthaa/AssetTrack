package models

import "time"

type AssetRequest struct {
	Brand         string               `json:"brand" db:"brand" validate:"required"`
	Model         string               `json:"model" db:"model" validate:"required"`
	SerialNumber  string               `json:"serialNumber" db:"serial_number" validate:"required"`
	AssetType     string               `json:"assetType" db:"asset_type" validate:"required,oneof=laptop keyboard mouse mobile"`
	Status        string               `json:"status" db:"status" validate:"required,oneof=available assigned needs_repair under_repair damaged"`
	OwnerType     string               `json:"owner" db:"owner_type" validate:"required,oneof=client remotestate"`
	WarrantyStart time.Time            `json:"warrantyStart" db:"warranty_start" validate:"required"`
	WarrantyEnd   time.Time            `json:"warrantyEnd" db:"warranty_end" validate:"required"`
	Laptop        LaptopSpecsRequest   `json:"laptopSpecs,omitempty"`
	Keyboard      KeyboardSpecsRequest `json:"keyboardSpecs,omitempty"`
	Mouse         MouseSpecsRequest    `json:"mouseSpecs,omitempty"`
	Mobile        MobileSpecsRequest   `json:"mobileSpecs,omitempty"`
}
type LaptopSpecsRequest struct {
	Processor       string `json:"processor" validate:"required"`
	Ram             string `json:"ram" validate:"required"`
	Storage         string `json:"storage" validate:"required"`
	OperatingSystem string `json:"operatingSystem" validate:"required"`
	Charger         string `json:"charger" validate:"required"`
}
type LaptopSpecs struct {
	AssetID         string `json:"assetID" db:"asset_id"`
	Processor       string `json:"processor" db:"processor"`
	Ram             string `json:"ram" db:"ram"`
	Storage         string `json:"storage" db:"storage"`
	OperatingSystem string `json:"operatingSystem" db:"operating_system"`
	Charger         string `json:"charger" db:"charger"`
}

type KeyboardSpecsRequest struct {
	Layout       string `json:"layout" validate:"required"`
	Connectivity string `json:"connectivity" validate:"required"`
}
type KeyboardSpecs struct {
	AssetID      string `json:"assetID" db:"asset_id"`
	Layout       string `json:"layout" db:"layout"`
	Connectivity string `json:"connectivity" db:"connectivity"`
}

type MouseSpecsRequest struct {
	Dpi          int    `json:"dpi" validate:"required,gt=0"`
	Connectivity string `json:"connectivity" validate:"required,oneof=wired wireless bluetooth"`
}
type MouseSpecs struct {
	AssetID      string `json:"assetID" db:"asset_id"`
	Dpi          int    `json:"dpi" db:"dpi"`
	Connectivity string `json:"connectivity" db:"connectivity"`
}

type MobileSpecsRequest struct {
	OperatingSystem string `json:"operatingSystem" validate:"required"`
	Ram             string `json:"ram" validate:"required"`
	Storage         string `json:"storage" validate:"required"`
	Charger         string `json:"charger" validate:"required"`
}
type MobileSpecs struct {
	AssetID         string `json:"assetID" db:"asset_id"`
	OperatingSystem string `json:"operatingSystem" db:"operating_system"`
	Ram             string `json:"ram" db:"ram"`
	Storage         string `json:"storage" db:"storage"`
	Charger         string `json:"charger" db:"charger"`
}
