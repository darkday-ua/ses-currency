
package main
import (
	"fmt"
	"currency_service/currencies"
	"currency_service/users"
	"currency_service/api"
)

var appVersion string

func main(){
	fmt.Printf("Currency-Service version %s",appVersion)
	fmt.Printf(currencies.About())
	fmt.Printf(users.About())
	fmt.Printf(api.About())
	api.Init()
}