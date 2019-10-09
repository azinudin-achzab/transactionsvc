package routers

import (
	"Trx-service/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/trx/getAll", &controllers.TrxController{}, "post:GetAll")
	beego.Router("/trx/deposit", &controllers.TrxController{}, "post:Deposit")
	beego.Router("/trx/transferin", &controllers.TrxController{}, "post:InternalTransfer")
	beego.Router("/trx/inquiryhistory/:acc", &controllers.TrxController{}, "get:InquiryHistory")
	// Customer
	beego.Router("/customer/inquirybalance/:acc", &controllers.CustomerController{}, "get:InquiryBalance")
}
