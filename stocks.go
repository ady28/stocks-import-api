package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"stocks/alphadata"
	"stocks/stocksdb"
	"stocks/yahoodata"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type Response struct {
	Success bool   `json:"status"`
	Message string `json:"message"`
}

var mongoDBServerName = os.Getenv("MONGODBSERVERNAME")
var mongoDBServerPort = os.Getenv("MONGODBSERVERPORT")
var PORT = os.Getenv("PORT")

func main() {

	handleRequests()
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/v1/import/{ticker}", importStock)
	s := &http.Server{
		Addr:           ":" + PORT,
		Handler:        myRouter,
		ReadTimeout:    1 * time.Minute,
		WriteTimeout:   1 * time.Minute,
		MaxHeaderBytes: 0,
	}
	s.ListenAndServe()
}

func importStock(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	key := vars["ticker"]
	key = strings.ToUpper(key)

	var c *alphadata.OverviewData
	var d *yahoodata.YahooData

	var akey *stocksdb.Key
	var ykey *stocksdb.Key

	ret := ""

	mongoDBAdminUser, mongoDBAdminUserPassword := getDBCredentials()

	akey = stocksdb.GetKey("alpha", mongoDBServerName, mongoDBServerPort, mongoDBAdminUser, mongoDBAdminUserPassword)
	ykey = stocksdb.GetKey("yahoo", mongoDBServerName, mongoDBServerPort, mongoDBAdminUser, mongoDBAdminUserPassword)
	if akey.Key == "" {
		w.Header().Set("Content-Type", "application/json")
		message := "Error getting API key for Alpha."
		json.NewEncoder(w).Encode(&Response{false, message})
	} else {
		if ykey.Key == "" {
			w.Header().Set("Content-Type", "application/json")
			message := "Error getting API key for Yahoo."
			json.NewEncoder(w).Encode(&Response{false, message})
		}
	}

	if akey.Key != "" && ykey.Key != "" {
		c = alphadata.NewOverviewData(akey.Key, key)
		if c.Note != "" {
			w.Header().Set("Content-Type", "application/json")
			message := "Error getting " + key + " with AlphaVantage API. " + c.Note
			json.NewEncoder(w).Encode(&Response{false, message})
		} else if c.Symbol == "" {
			w.Header().Set("Content-Type", "application/json")
			message := "Error getting " + key + ". Check Alpha API."
			json.NewEncoder(w).Encode(&Response{false, message})
		} else {
			d = yahoodata.NewData(ykey.Key, key)
			if d == nil {
				w.Header().Set("Content-Type", "application/json")
				message := "Error getting " + key + ". Check Yahoo API."
				json.NewEncoder(w).Encode(&Response{false, message})
			} else if len(d.QuoteSummary.Result) == 0 {
				w.Header().Set("Content-Type", "application/json")
				message := "Error getting " + key + ". Stock not found."
				json.NewEncoder(w).Encode(&Response{false, message})
			} else if d.QuoteSummary.Result[0].AssetProfile.Country == "" {
				w.Header().Set("Content-Type", "application/json")
				message := "Error getting " + key + ". Stock not found."
				json.NewEncoder(w).Encode(&Response{false, message})
			} else {
				if stocksdb.FindStock(key, mongoDBServerName, mongoDBServerPort, mongoDBAdminUser, mongoDBAdminUserPassword) {
					ret = "Stock " + key + " already exists. Updating relevant data"
					stocksdb.UpdateStock(c, d, key, mongoDBServerName, mongoDBServerPort, mongoDBAdminUser, mongoDBAdminUserPassword)
				} else {
					ret = "Getting and inserting new stock " + key
					stocksdb.NewStock(c, d, mongoDBServerName, mongoDBServerPort, mongoDBAdminUser, mongoDBAdminUserPassword)
				}

				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(&Response{true, ret})
			}
		}
	}
}

func getDBCredentials() (string, string) {
	dbuser, _ := ioutil.ReadFile("/run/secrets/stocksmongouser")
	dbpass, _ := ioutil.ReadFile("/run/secrets/stocksmongopassword")

	du := string(dbuser)
	du = strings.TrimRight(du, "\n")
	dp := string(dbpass)
	dp = strings.TrimRight(dp, "\n")

	return du, dp
}
