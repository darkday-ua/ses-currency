package currencies

import (
	"bytes"
	"currency_service/config"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

var	packageVersion string = "0.0.2"


type CurrencyPair struct {
	Currencies [2]string
}



func (pair CurrencyPair) FetchRate() (float64, error) {
	fmt.Println("Fetching rate for ",pair.Currencies[0],"/",pair.Currencies[1])
	request_template:= map[string]any{"currency": pair.Currencies[1], "code": pair.Currencies[0],"meta":false}
	request_json, err := json.Marshal(request_template)
	if err != nil {
        log.Fatal(err)
    }
	// Hardcoded until future of the task will be clear
	client := &http.Client{Timeout: 3*time.Second}
    r, _ := http.NewRequest("POST", "https://api.livecoinwatch.com/coins/single", bytes.NewBuffer(request_json)) // URL-encoded payload
    r.Header.Set("x-api-key", config.Config.LIVE_COIN_WATCH_API_KEY)
    r.Header.Set("content-type", "application/json")
    resp, err := client.Do(r)

	if err != nil {
        log.Println(err)
		if config.Config.EMULATE_RATE=="yes"{
			return 123456.7,nil
		}else{
			return 0.0,fmt.Errorf("can't fetch rate")
		}
    }
	defer resp.Body.Close()
	if resp.StatusCode!=200 {
		log.Println(err,resp)
		if config.Config.EMULATE_RATE=="yes"{
			return 123456.7,nil
		}else{
			return 0.0,fmt.Errorf("can't fetch rate")
		}
		
	}
	var res map[string]interface{}

    json.NewDecoder(resp.Body).Decode(&res)
	rate,ok:= res["rate"].(float64)
	//rate,err:= strconv.ParseFloat(string(res["rate"]),	64)
	if !ok {
		return 0.0,fmt.Errorf("can't parse rate value")
	}
	return rate,nil // Example rate value
}

func init() {
	fmt.Println("..currencies package version ", packageVersion)
}