package currencies
import
(
	"fmt"
)

var	packageVersion string = "0.0.1"

func About() string {return fmt.Sprint("\n\tcurrencies package version ", packageVersion)}

func init() {}