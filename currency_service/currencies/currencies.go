package currencies
import
(
	"fmt"
)

var	packageVersion string = "0.0.1"

func About() string {return fmt.Sprint("..currencies package version ", packageVersion)}

type CurrencyPair struct {
	Currencies [2]string
}



func (pair CurrencyPair) FetchRate() (float64, error) {
	fmt.Println("Fetching rate for ",pair.Currencies[0],"/",pair.Currencies[1])
	return 1.0/2.0,nil // Example rate value
}

func init() {}