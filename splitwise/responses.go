package splitwise

import (
	"encoding/json"
)

type balance struct {
	CurrencyCode string `json:"currency_code"`
	Amount       string `json:"amount"`
}

type Debt struct {
	From         int    `json:"from"`
	To           int    `json:"to"`
	Amount       string `json:"amount"`
	CurrencyCode string `json:"currency_code"`
}

type DeleteResponse struct {
	Success bool     `json:"success"`
	Errors  []string `json:"errors"`
}

func indentString(v interface{}) (string, error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}
