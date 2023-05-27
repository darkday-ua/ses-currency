
package main
import (
	"fmt"
	"currency_service/api"
)

var appVersion string

func main(){
	api.Init()
}

func init(){
	fmt.Println("\nCurrency-Service version ",appVersion)
}