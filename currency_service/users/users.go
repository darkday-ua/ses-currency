package users
import
(
	"fmt"
	"log"
	"currency_service/db"
//	"currency_service/currency"
)

var	packageVersion string = "0.0.1"
var sessionManager *db.SessionManager


func AddUser(email string) (string,bool) {
	dbSession:= sessionManager.GetDBSession()		
	desc,added:= dbSession.AddUser(email)
	if !added  {
		return desc,false
	}
	
	return "",true
}

func init() {
	fmt.Println("..users package version ", packageVersion)
	var err error
	sessionManager,err = db.GetSessionManager()
	if err!=nil{
		log.Fatal("Error getting session manager: ",err)
	}

}