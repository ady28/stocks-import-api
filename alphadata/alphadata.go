package alphadata

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

var ABASEURL = "https://www.alphavantage.co/query"
var AOVERVIEW = ABASEURL + "?function=OVERVIEW"

type OverviewData struct {
	Symbol               string `json:"Symbol"`
	AssetType            string `json:"AssetType"`
	Name                 string `json:"Name"`
	Description          string `json:"Description"`
	CIK                  int64  `json:"CIK,string"`
	Exchange             string `json:"Exchange"`
	Currency             string `json:"Currency"`
	Country              string `json:"Country"`
	Sector               string `json:"Sector"`
	Industry             string `json:"Industry"`
	Address              string `json:"Address"`
	FiscalYearEnd        string `json:"FiscalYearEnd"`
	LatestQuarter        string `json:"LatestQuarter"`
	MarketCapitalization string `json:"MarketCapitalization"`
	EBITDA               string `json:"EBITDA"`
	PERatio              string `json:"PERatio"`
	PEGRatio             string `json:"PEGRatio"`
	BookValue            string `json:"BookValue"`
	DividendPerShare     string `json:"DividendPerShare"`
	DividendYield        string `json:"DividendYield"`
	EPS                  string `json:"EPS"`
	RevenuePerShareTTM   string `json:"RevenuePerShareTTM"`
	ProfitMargin         string `json:"ProfitMargin"`
	OperatingMarginTTM   string `json:"OperatingMarginTTM"`
	ReturnOnAssetsTTM    string `json:"ReturnOnAssetsTTM"`
	ReturnOnEquityTTM    string `json:"ReturnOnEquityTTM"`
	RevenueTTM           string `json:"RevenueTTM"`
	GrossProfitTTM       string `json:"GrossProfitTTM"`
	DilutedEPSTTM        string `json:"DilutedEPSTTM"`
	QEarningsGrthYOY     string `json:"QuarterlyEarningsGrowthYOY"`
	QRevenueGrthYOY      string `json:"QuarterlyRevenueGrowthYOY"`
	AnalystTargetPrice   string `json:"AnalystTargetPrice"`
	TrailingPE           string `json:"TrailingPE"`
	ForwardPE            string `json:"ForwardPE"`
	PriceToSalesRatioTTM string `json:"PriceToSalesRatioTTM"`
	PriceToBookRatio     string `json:"PriceToBookRatio"`
	EVToRevenue          string `json:"EVToRevenue"`
	EVToEBITDA           string `json:"EVToEBITDA"`
	Beta                 string `json:"Beta"`
	WeekHigh52           string `json:"52WeekHigh"`
	WeekLow52            string `json:"52WeekLow"`
	DayMovingAverage50   string `json:"50DayMovingAverage"`
	DayMovingAverage200  string `json:"200DayMovingAverage"`
	SharesOutstanding    string `json:"SharesOutstanding"`
	DividendDate         string `json:"DividendDate"`
	ExDividendDate       string `json:"ExDividendDate"`
	Note                 string `json:"Note"`
}

func NewOverviewData(apikey string, ticker string) *OverviewData {
	p := new(OverviewData)

	overviewLink := AOVERVIEW + "&symbol=" + ticker + "&apikey=" + apikey

	req, err := http.NewRequest("GET", overviewLink, nil)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	oresp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	obody, err := ioutil.ReadAll(oresp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(obody, &p)
	if err != nil {
		log.Fatal(err)
	}

	return p
}
