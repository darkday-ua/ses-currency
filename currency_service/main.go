
package main
import (
	"fmt"
	"currency_service/currencies"
	"currency_service/users"
	"currency_service/api"
)

var appVersion string

func main(){
	fmt.Printf("\nCurrency-Service version %s\n",appVersion)
	fmt.Println(currencies.About())
	fmt.Println(users.About())
	fmt.Println(api.About())
	api.Init()
}