package api
import(
	"encoding/json"
	"log"
	"net/http"
	"fmt"
	"currency_service/currencies"
	"currency_service/config"
)

var	packageVersion string = "0.0.1"


func Init() {
	fmt.Println("..api package version ", packageVersion)
	// For the sake of simplicity we build router in explicit way
	// later we could use dynamic configuration if it is possible in Go, though)
	http.HandleFunc("/api/rate", getRateHandler)
	http.HandleFunc("/api/subscribe", subscribeHandler)
	http.HandleFunc("/api/sendEmails", sendEmailsHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s",config.Config.APP_PORT), nil))
}

func getRateHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// Here we can obtain the currency pair(s) from the request but for now lets follow the task
	// If we not require HFT we could use a cache for rates
	pair:= currencies.CurrencyPair{Currencies: [2]string{"BTC","UAH"}}
	rate,err:=pair.FetchRate()
	// actually we should return HTTP 503 Service Unavailable
	// But according to the task's definition we should return HTTP 400 Bad Request
	if err!=nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	response := struct {
		Rate float64 `json:"rate"`
	}{
		Rate: rate,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
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
