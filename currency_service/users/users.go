package users
import
(
	"fmt"
	"log"
	"currency_service/db"
	"currency_service/config"
	"net/smtp"
	"currency_service/currencies"
)

var	packageVersion string = "0.0.2"
var sessionManager *db.SessionManager


func AddUser(email string) (string,bool) {
	dbSession:= sessionManager.GetDBSession()		
	desc,added:= dbSession.AddUser(email)
	if !added  {
		return desc,false
	}
	
	return "",true
}
// mock validation
func ValidateEmail(email string) bool {
	return email!="" && len(email)>3 && len(email)<30
}

func SendRate() (string,bool) {
	dbSession:= sessionManager.GetDBSession()		
	users:= dbSession.GetSubscribedUsers()
	//count:=len(users)

	// for real app we should use here multiprocessing or queue
	smtp_from:=config.Config.SMTP_FROM
	smtp_password:=config.Config.SMTP_PASSWORD
	smtp_user:=config.Config.SMTP_USER
	smtp_host:=config.Config.SMTP_HOST
	smtp_port:=config.Config.SMTP_PORT
	if smtp_from=="" || smtp_password=="" || smtp_user=="" || smtp_host=="" || smtp_port=="" {
		return "not able to send email, check smtp settings",false
	}
	auth := smtp.PlainAuth("", smtp_user, smtp_password, smtp_host)
	pair:= currencies.CurrencyPair{Currencies: [2]string{"BTC","UAH"}}
	rate,_:=pair.FetchRate()

	for _,user:=range users {
		log.Println("sending email to ",user)
		body:=fmt.Sprintf("To: "+user+"\r\nSubject: Currency rate\n%s / %s: %f\n",pair.Currencies[0],pair.Currencies[1],rate)
		log.Println(body)
		err := smtp.SendMail(smtp_host+":"+smtp_port, auth, smtp_from, []string{user}, []byte(body))
		if err!=nil{
			log.Println("Error sending email to ",user)
			log.Println(err)
		}
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