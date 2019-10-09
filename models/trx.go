package models

import (
	"errors"
	"strings"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

// Trx ...
type Trx struct {
	ID      int       `json:"id" orm:"pk;column(id);auto"`
	FromAcc string    `json:"from_account" orm:"column(from_account)"`
	ToAcc   string    `json:"to_account" orm:"column(to_account)"`
	Amount  float64   `json:"amount" orm:"column(amount)"`
	Desc    string    `json:"desc" orm:"column(desc)"`
	Type    string    `json:"type" orm:"column(type)"`
	Time    time.Time `json:"time" orm:"column(time)"`
	Date    time.Time `json:"date" orm:"column(date)"`
}

func init() {
	orm.RegisterModel(new(Trx))
}

// GetAll ...
func GetAll() []*Trx {
	var trxs []*Trx

	qb := []string{
		"SELECT",
		"*",
		"FROM", "trx",
		"ORDER BY date,time DESC",
	}
	sql := strings.Join(qb, " ")

	o := orm.NewOrm()
	_, err := o.Raw(sql).QueryRows(&trxs)

	if err != nil {
		l := logs.GetLogger()
		l.Println(err)
		panic(1)
	}

	//l := logs.GetLogger()
	//l.Println(&trxs)

	return trxs
}

// Deposit ...
func Deposit(trx Trx) (err error) {
	if len(trx.ToAcc) != 11 {
		err = errors.New("wrong input for account number")
	} else {
		cekcustomer := GetCustomerID(trx.ToAcc)
		if len(cekcustomer) <= 0 {
			err = errors.New("customer not exist")
		} else {

			o := orm.NewOrm()
			trx.FromAcc = "TELLER"
			trx.Time = time.Now()
			trx.Date = time.Now()
			trx.Desc = "Deposit from OLD CBS"
			trx.Type = "DEBIT"

			_, err = o.Insert(&trx)
			if err == nil {
				err2 := UpdateBalance(trx.Amount, trx.ToAcc, "+")
				if err2 != nil {
					err = err2
				}
			}

		}
	}
	return err
}

// InternalTransfer ...
func InternalTransfer(trx Trx) (err error) {
	if len(trx.ToAcc) != 11 || len(trx.FromAcc) != 11 {
		err = errors.New("wrong input for account number")
	} else {
		cekcustomerkirim := GetCustomerID(trx.FromAcc)
		cekcustomerterima := GetCustomerID(trx.ToAcc)
		if len(cekcustomerkirim) <= 0 || len(cekcustomerterima) <= 0 {
			err = errors.New("customer not exist")
		} else {
			if trx.Amount > cekcustomerkirim[0].Balance {
				err = errors.New("Balance not sufficient to transfer")
			} else {
				o := orm.NewOrm()

				//insert trx yg kirim
				trx.Time = time.Now()
				trx.Date = time.Now()
				trx.Desc = "Internal transfer to " + trx.ToAcc
				trx.Type = "KREDIT"

				//insert trx yg terima
				var trx2 Trx
				trx2.Time = time.Now()
				trx2.Date = time.Now()
				trx2.ToAcc = trx.ToAcc
				trx2.FromAcc = trx.FromAcc
				trx2.Desc = "Internal transfer from " + trx.FromAcc
				trx2.Type = "DEBIT"
				trx2.Amount = trx.Amount

				trxs := []Trx{trx, trx2}
				_, err := o.InsertMulti(2, trxs)
				if err == nil {
					err2 := UpdateBalance(trx.Amount, trx.ToAcc, "+")
					err2 = UpdateBalance(trx.Amount, trx.FromAcc, "-")
					if err2 != nil {
						err = err2
					}
				}
			}
		}
	}
	return err
}

// InquiryHistory ...
func InquiryHistory(acc string) []*Trx {
	var trxs []*Trx

	qb := []string{
		"SELECT",
		"*",
		"FROM", "trx",
		"WHERE to_account = ?",
		"ORDER BY date,time DESC",
	}
	sql := strings.Join(qb, " ")

	o := orm.NewOrm()
	_, err := o.Raw(sql, acc).QueryRows(&trxs)

	if err != nil {
		l := logs.GetLogger()
		l.Println(err)
		panic(1)
	}

	//l := logs.GetLogger()
	//l.Println(&trxs)

	return trxs
}
