package db

import(
	"fmt"
	"os"
	"log"
	"currency_service/config"

)

var	packageVersion string = "0.0.1"

var db_in_memory bool = false

func init() {
	fmt.Println("..db package version ", packageVersion)
	file, err:= os.OpenFile(config.Config.DB_PATH,os.O_CREATE, 0666)
    if err != nil {
        fmt.Printf("Unable to create/open db file at  %s, in-memory db will be used\n",config.Config.DB_PATH)
		db_in_memory = true
    } else {
		fmt.Println("Database file created at ",config.Config.DB_PATH)
	}
    file.Close()
}

