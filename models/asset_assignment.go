package models

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
