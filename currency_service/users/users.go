package users
import
(
	"fmt"
)

var	packageVersion string = "0.0.1"

func About() string {return fmt.Sprint("..users package version ", packageVersion)}

func init() {}