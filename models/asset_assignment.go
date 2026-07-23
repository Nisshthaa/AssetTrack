package models

import "time"

type AssignAssetRequest struct {
	AssetID string `json:"assetID" validate:"required,uuid"`
	UserID  string `json:"userID" validate:"required,uuid"`
}

type Dashboard struct {
	TotalAssets          int `json:"totalAssets"`
	TotalAssignedAssets  int `json:"totalAssignedAssets"`
	TotalAvailableAssets int `json:"totalAvailableAssets"`
	NeedsRepairAssets    int `json:"needsRepairAssets"`
	UnderRepairAssets    int `json:"underRepairAssets"`
	DamagedAssets        int `json:"damagedAssets"`
}

type AssetHistory struct {
	AssetID    string     `json:"assetId" db:"asset_id"`
	AssetType  string     `json:"assetType" db:"asset_type"`
	Brand      string     `json:"brand" db:"brand"`
	Model      string     `json:"model" db:"model"`
	UserID     string     `json:"userId" db:"user_id"`
	Name       string     `json:"name" db:"name"`
	AssignedOn time.Time  `json:"assignedOn" db:"assigned_on"`
	ReturnedAt *time.Time `json:"returnedAt" db:"returned_at"`
}
