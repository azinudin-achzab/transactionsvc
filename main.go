package main

import (
	_ "Trx-service/routers"
	"fmt"
	"log"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	} else {

		// get database configuration from environment variables
		dbUser := os.Getenv("db_user")
		dbPwd := os.Getenv("db_password")
		dbName := os.Getenv("db_name")
		dbHost := os.Getenv("db_host")
		dbPort := os.Getenv("db_port")
		dbString := "user=" + dbUser + " password=" + dbPwd + " host=" + dbHost + " port=" + dbPort + " dbname=" + dbName + " sslmode=disable"

		// Register Driver
		orm.RegisterDriver("postgres", orm.DRPostgres)

		// Register default database
		orm.RegisterDataBase("default", "postgres", dbString)

		// autosync
		// db alias
		name := "default"

		// drop table and re-create
		force := false

		// print log
		verbose := true

		// error
		err := orm.RunSyncdb(name, force, verbose)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	/*o := orm.NewOrm()
	o.Using("default")

	cust := new(Customer)
	cust.Cif = "123123123124"
	cust.AccountNumber = "12312312311"
	cust.AccountStatus = "ACTIVE"
	cust.FirstName = "Azinudin"
	cust.LastName = "Achzab"
	cust.ProductCode = "AA002"
	cust.OpeningDate = time.Now()
	cust.Balance = 0

	o.Insert((cust))*/

	beego.Run()
}
