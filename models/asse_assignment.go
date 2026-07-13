package models

type AssignAssetRequest struct {
	AssetID string `json:"assetID" validate:"required,uuid"`
	UserID  string `json:"userID" validate:"required,uuid"`
}
