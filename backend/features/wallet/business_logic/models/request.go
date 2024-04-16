package models

type WalletRequest struct {
	Name          string `json:"name" form:"name"`
	Type          string `json:"type" form:"type" validate:"required"`
	AccountNumber string `json:"account_number" form:"account_number" validate:"required,min=3,max=30"`
	AccountScheme string `json:"account_scheme" form:"account_scheme" validate:"required,min=3,max=30"`
	Owner         string `json:"owner" form:"owner" validate:"required,uuid"`
}
