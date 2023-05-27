package api
import(
	"encoding/json"
	"log"
	"net/http"
	"fmt"
)

var	packageVersion string = "0.0.1"

func About() string {return fmt.Sprint("\n\tapi package version ", packageVersion)}

func Init() {
	http.HandleFunc("/api/rate", getRateHandler)
	http.HandleFunc("/api/subscribe", subscribeHandler)
	http.HandleFunc("/api/sendEmails", sendEmailsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getRateHandler(w http.ResponseWriter, r *http.Request) {	
	
}

func subscribeHandler(w http.ResponseWriter, r *http.Request) {
	
}

func sendEmailsHandler(w http.ResponseWriter, r *http.Request) {
	
}
