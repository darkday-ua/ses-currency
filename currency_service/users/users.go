package users
import
(
	"fmt"
)

var	packageVersion string = "0.0.1"


func init() {
	fmt.Println("..users package version ", packageVersion)
}