package db

import(
	"fmt"
	"os"
	"log"
	"currency_service/config"

)

var	packageVersion string = "0.0.1"

func init() {
	fmt.Println("..db package version ", packageVersion)
	file, err:= os.OpenFile(config.Config.DB_PATH,os.O_CREATE, 0666)
    if err != nil {
        log.Fatal(err)
    } else {
		fmt.Println("Database file created at ",config.Config.DB_PATH)
	}
    file.Close()
}

