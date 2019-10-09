package helpers

import (
	"Trx-service/models"
)

// ResponseTrx of crud process
type ResponseTrx struct {
	Errcode int
	Message string
	Trx     models.Trx
}

// ResponseCustomer of crud process
type ResponseCustomer struct {
	Errcode  int
	Message  string
	Customer models.Customer
}

// ValidateTrx ...
func ValidateTrx(trx *models.Trx) bool {
	var result bool
	if trx.ToAcc == "" || trx.Amount <= 0 {
		result = false
	} else {
		result = true
	}
	return result
}

// ValidateTrx2 ...
func ValidateTrx2(trx *models.Trx) bool {
	var result bool
	if trx.ToAcc == "" || trx.FromAcc == "" || trx.Amount <= 0 {
		result = false
	} else {
		result = true
	}
	return result
}
