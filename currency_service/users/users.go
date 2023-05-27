package users
import
(
	"fmt"
)

var	packageVersion string = "0.0.1"

func About() string {return fmt.Sprint("\n\tusers package version ", packageVersion)}

func init() {}