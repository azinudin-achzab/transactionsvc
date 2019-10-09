package controllers

import (
	"Trx-service/helpers"
	"Trx-service/models"
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

//TrxController ...
type TrxController struct {
	beego.Controller
}

// GetAll function gets all the users
func (tc *TrxController) GetAll() {
	//uc.Data["json"] = models.GetAllUsers()
	tc.Data["json"] = models.GetAll()
	tc.ServeJSON()
}

// Deposit ...
func (tc *TrxController) Deposit() {
	var trx models.Trx
	json.Unmarshal(tc.Ctx.Input.RequestBody, &trx)
	if helpers.ValidateTrx(&trx) {
		tc.Data["json"] = helpers.ResponseTrx{0, "Transaction was succesful", trx}
		err := models.Deposit(trx)
		if err != nil {
			tc.Data["json"] = helpers.ResponseTrx{1, "Transaction was failed: " + err.Error(), models.Trx{}}
		} else {
			tc.Data["json"] = helpers.ResponseTrx{0, "Transaction was succesful", trx}
		}
	} else {
		tc.Data["json"] = helpers.ResponseTrx{0, "Transaction was failed: required field cant be empty", models.Trx{}}
	}
	l := logs.GetLogger()
	l.Println(trx)
	tc.ServeJSON()
}

// InternalTransfer ...
func (tc *TrxController) InternalTransfer() {
	var trx models.Trx
	json.Unmarshal(tc.Ctx.Input.RequestBody, &trx)
	if helpers.ValidateTrx2(&trx) {
		tc.Data["json"] = helpers.ResponseTrx{0, "Transaction was succesful", trx}
		err := models.InternalTransfer(trx)
		if err != nil {
			tc.Data["json"] = helpers.ResponseTrx{1, "Transaction was failed: " + err.Error(), models.Trx{}}
		} else {
			tc.Data["json"] = helpers.ResponseTrx{0, "Transaction was succesful", trx}
		}
	} else {
		tc.Data["json"] = helpers.ResponseTrx{0, "Transaction was failed: required field cant be empty", models.Trx{}}
	}
	l := logs.GetLogger()
	l.Println(trx)
	tc.ServeJSON()
}

// InquiryHistory
func (tc *TrxController) InquiryHistory() {
	acc := tc.Ctx.Input.Param(":acc")
	if len(models.InquiryHistory(acc)) <= 0 {
		tc.Data["json"] = helpers.ResponseTrx{1, "Transaction Failed: Account not exist", models.Trx{}}
	} else {
		tc.Data["json"] = models.InquiryHistory(acc)
	}
	tc.ServeJSON()
}
