package models

import (
	"strings"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

// Customer ...
type Customer struct {
	Cif           string    `json:"cif" orm:"column(cif);pk"`
	AccountNumber string    `json:"account_number,omitempty" orm:"column(account_number)"`
	FirstName     string    `json:"first_name,omitempty" orm:"column(first_name)"`
	MiddleName    string    `json:"middle_name,omitempty" orm:"column(middle_name)"`
	LastName      string    `json:"last_name,omitempty" orm:"column(last_name)"`
	ProductCode   string    `json:"product_code,omitempty" orm:"column(product_code)"`
	OpeningDate   time.Time `json:"opening_date" orm:"column(opening_date)"`
	AccountStatus string    `json:"account_status,omitempty" orm:"column(account_status)"`
	Balance       float64   `json:"balance,omitempty" orm:"column(balance)"`
	//omitempty suapayaa bisa kosong
}

func init() {
	orm.RegisterModel(new(Customer))
}

// GetCustomerID ...
func GetCustomerID(acc string) []Customer {
	var cust []Customer

	qb := []string{
		"SELECT",
		"*",
		"FROM", "customer",
		"WHERE account_number = ?",
	}
	sql := strings.Join(qb, " ")

	o := orm.NewOrm()
	_, err := o.Raw(sql, acc).QueryRows(&cust)

	if err != nil {
		l := logs.GetLogger()
		l.Println(err)
		panic(1)
	}

	return cust
}

// UpdateBalance ...
func UpdateBalance(amount float64, acc string, kind string) error {
	qb := []string{
		"UPDATE",
		"customer",
		"SET",
		"balance = balance " + kind + " ?",
		"WHERE account_number = ?",
	}
	sql := strings.Join(qb, " ")

	o := orm.NewOrm()
	_, err := o.Raw(sql, amount, acc).Exec()
	return err
}

// InquiryBalance ...
func InquiryBalance(acc string) []*Customer {
	var custs []*Customer
	//l := logs.GetLogger()
	//l.Println(acc)
	qb := []string{
		"SELECT",
		"balance",
		"FROM", "customer",
		"WHERE account_number = ?",
	}
	sql := strings.Join(qb, " ")

	o := orm.NewOrm()
	_, err := o.Raw(sql, acc).QueryRows(&custs)

	if err != nil {
		custs = []*Customer{}
	}

	return custs
}
