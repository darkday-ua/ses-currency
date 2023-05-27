package users
import
(
	"fmt"
//	"currency_service/config"
	"currency_service/db"
//	"currency_service/currency"
)

var	packageVersion string = "0.0.1"


func AddUser(email string) error {
	fmt.Println("AddUser",email)
	return nil
}

func init() {
	fmt.Println("..users package version ", packageVersion)
	db.Init()
}