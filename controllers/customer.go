package controllers

import (
	"Trx-service/helpers"
	"Trx-service/models"

	"github.com/astaxie/beego"
)

//CustomerController ...
type CustomerController struct {
	beego.Controller
}

// InquiryBalance ...
func (cc *CustomerController) InquiryBalance() {
	acc := cc.Ctx.Input.Param(":acc")
	if len(models.InquiryBalance(acc)) <= 0 {
		cc.Data["json"] = helpers.ResponseCustomer{1, "Transaction Failed: Account not exist", models.Customer{}}
	} else {
		cc.Data["json"] = models.InquiryBalance(acc)
	}
	cc.ServeJSON()
}
