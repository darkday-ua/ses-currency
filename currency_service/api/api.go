package api
import(
	"encoding/json"
	"log"
	"net/http"
	"fmt"
	"currency_service/currencies"
)

var	packageVersion string = "0.0.1"

func About() string {return fmt.Sprint("..api package version ", packageVersion)}

func Init() {
	http.HandleFunc("/api/rate", getRateHandler)
	http.HandleFunc("/api/subscribe", subscribeHandler)
	http.HandleFunc("/api/sendEmails", sendEmailsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getRateHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	
	pair:= currencies.CurrencyPair{Currencies: [2]string{"BTC","UAH"}}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := struct {
		Rate float64 `json:"rate"`
	}{
		Rate: pair.FetchRate(),
	}
	json.NewEncoder(w).Encode(response)
}

func subscribeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func sendEmailsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)

}
