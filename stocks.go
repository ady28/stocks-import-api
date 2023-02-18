package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
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
	myRouter.HandleFunc("/health", health)
	s := &http.Server{
		Addr:           ":" + PORT,
		Handler:        myRouter,
		ReadTimeout:    1 * time.Minute,
		WriteTimeout:   1 * time.Minute,
		MaxHeaderBytes: 0,
	}
	s.ListenAndServe()
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&Response{true, "OK"})
}

func importStock(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	key := vars["ticker"]
	key = strings.ToUpper(key)

	var d *yahoodata.YahooData
	var ykey *stocksdb.Key

	ret := ""

	mongoDBAdminUser, mongoDBAdminUserPassword := getDBCredentials()

	ykey = stocksdb.GetKey("yahoo", mongoDBServerName, mongoDBServerPort, mongoDBAdminUser, mongoDBAdminUserPassword)
	if ykey.Key == "" {
		w.Header().Set("Content-Type", "application/json")
		message := "Error getting API key for Yahoo."
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
				stocksdb.UpdateStock(d, key, mongoDBServerName, mongoDBServerPort, mongoDBAdminUser, mongoDBAdminUserPassword)
			} else {
				ret = "Getting and inserting new stock " + key
				stocksdb.NewStock(d, mongoDBServerName, mongoDBServerPort, mongoDBAdminUser, mongoDBAdminUserPassword)
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(&Response{true, ret})
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
