package responses

import "time"

type AccountResponse struct {
	ID                int       `json:"id"`
	AccountName       string    `json:"account_name"`
	AccountNumber     int64     `json:"account_number"`
	AccountCode       string    `json:"account_code"`
	AccountHolderName string    `json:"account_holder_name"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
