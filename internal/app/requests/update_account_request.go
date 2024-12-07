package requests

type UpdateAccountRequest struct {
	ID            int    `validate:"required"`
	UserID        int    `validate:"required" json:"user_id"`
	AccountName   string `validate:"required,max=255,min=3" json:"account_name"`
	AccountNumber int64  `validate:"required" json:"account_number"`
	AccountCode   string `validate:"required" json:"account_code"`
}
