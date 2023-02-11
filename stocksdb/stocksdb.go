package stocksdb

import (
	"context"
	"log"
	"math"
	"net/http"
	"os"
	"stocks/alphadata"
	"stocks/yahoodata"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Stock struct {
	ID                          primitive.ObjectID `bson:"_id,omitempty"`
	Name                        string             `bson:"name"`
	Ticker                      string             `bson:"ticker"`
	Beta                        float64            `bson:"beta"`
	Industry                    string             `bson:"industry"`
	Sector                      string             `bson:"sector"`
	Address                     string             `bson:"address"`
	City                        string             `bson:"city"`
	Country                     string             `bson:"country"`
	EmployeeNo                  int64              `bson:"employeeno"`
	RecommTrend                 recommTrend        `bson:"recommtrend"`
	CashFlowH                   []cashFlowH        `bson:"cashflowh"`
	EnterpriseValue             int64              `bson:"enterprisevalue"`
	EnterpriseValueNice         string             `bson:"enterprisevaluenice"`
	ForwardPE                   float64            `bson:"forwardpe"`
	ForwardPENice               string             `bson:"forwardpenice"`
	ProfitMargins               float64            `bson:"profitmargins"`
	ProfitMarginsNice           string             `bson:"profitmarginsnice"`
	FloatShares                 int64              `bson:"floatshares"`
	FloatSharesNice             string             `bson:"floatsharesnice"`
	SharesOutstanding           int64              `bson:"sharesoutstanding"`
	SharesOutstandingNice       string             `bson:"sharesoutstandingnice"`
	SharesShort                 int64              `bson:"sharesshort"`
	SharesShortNice             string             `bson:"sharesshortnice"`
	HeldPercentInsiders         float64            `bson:"heldpercentinsiders"`
	HeldPercentInsidersNice     string             `bson:"heldpercentinsidersnice"`
	HeldPercentInstitutions     float64            `bson:"heldpercentinstitutions"`
	HeldPercentInstitutionsNice string             `bson:"heldpercentinstitutionsnice"`
	ShortRatio                  float64            `bson:"shortratio"`
	ShortRatioNice              string             `bson:"shortrationice"`
	ShortPercentOfFloat         float64            `bson:"shortpercentoffloat"`
	ShortPercentOfFloatNice     string             `bson:"shortpercentoffloatnice"`
	BookValue                   float64            `bson:"bookvalue"`
	BookValueNice               string             `bson:"bookvaluenice"`
	PriceToBook                 float64            `bson:"pricetobook"`
	PriceToBookNice             string             `bson:"pricetobooknice"`
	LastFiscalYearEnd           string             `bson:"lastfiscalyearend"`
	MostRecentQuarter           string             `bson:"mostrecentquarter"`
	NetIncomeToCommon           int64              `bson:"netincometocommon"`
	NetIncomeToCommonNice       string             `bson:"netincometocommonnice"`
	TrailingEps                 float64            `bson:"trailingeps"`
	TrailingEpsNice             string             `bson:"trailingepsnice"`
	ForwardEps                  float64            `bson:"forwardeps"`
	ForwardEpsNice              string             `bson:"forwardepsnice"`
	PegRatio                    float64            `bson:"pegratio"`
	PegRatioNice                string             `bson:"pegrationice"`
	LastSplitFactor             string             `bson:"lastsplitfactor"`
	LastSplitDate               string             `bson:"lastsplitdate"`
	EnterpriseToRevenue         float64            `bson:"enterprisetorevenue"`
	EnterpriseToRevenueNice     string             `bson:"enterprisetorevenuenice"`
	EnterpriseToEbitda          float64            `bson:"enterprisetoebitda"`
	EnterpriseToEbitdaNice      string             `bson:"enterprisetoebitdanice"`
	WeekChange52                float64            `bson:"weekchange52"`
	WeekChange52Nice            string             `bson:"weekchange52nice"`
	Exchange                    string             `bson:"exchange"`
	IncomeH                     []incomeH          `bson:"incomeh"`
	Currency                    string             `bson:"currency"`
	ExDividendDate              string             `bson:"exdividenddate"`
	DividendRate                float64            `bson:"dividendrate"`
	DividendRateNice            string             `bson:"dividendratenice"`
	DividendYield               float64            `bson:"dividendyield"`
	DividendYieldNice           string             `bson:"dividendyieldnice"`
	PayoutRatio                 float64            `bson:"payoutratio"`
	PayoutRatioNice             string             `bson:"payoutrationice"`
	TrailingPE                  float64            `bson:"trailingpe"`
	TrailingPENice              string             `bson:"trailingpenice"`
	MarketCap                   int64              `bson:"marketcap"`
	MarketCapNice               string             `bson:"marketcapnice"`
	EarningsNext                earningsN          `bson:"earningsnext"`
	BalanceH                    []balanceH         `bson:"balanceh"`
	Growth5y                    float64            `bson:"growth5y"`
	Growth5yNice                string             `bson:"growth5ynice"`
	BalanceHQ                   []balanceH         `bson:"balancehq"`
	IncomeHQ                    []incomeH          `bson:"incomehq"`
	CashFlowHQ                  []cashFlowH        `bson:"cashflowhq"`
	Price                       float64            `bson:"price"`
	TargetHighPrice             float64            `bson:"targethighprice"`
	TargetLowPrice              float64            `bson:"targetlowprice"`
	TargetMedianPrice           float64            `bson:"targetmedianprice"`
	RecommendationKey           string             `bson:"recommendationkey"`
	TotalCash                   int64              `bson:"totalcash"`
	TotalCashNice               string             `bson:"totalcashnice"`
	TotalCashPerShare           float64            `bson:"totalcashpershare"`
	TotalCashPerShareNice       string             `bson:"totalcashpersharenice"`
	Ebitda                      int64              `bson:"ebitda"`
	EbitdaNice                  string             `bson:"ebitdanice"`
	TotalDebt                   int64              `bson:"totaldebt"`
	TotalDebtNice               string             `bson:"totaldebtnice"`
	QuickRatio                  float64            `bson:"quickratio"`
	QuickRatioNice              string             `bson:"quickrationice"`
	CurrentRatio                float64            `bson:"currentratio"`
	CurrentRatioNice            string             `bson:"currentrationice"`
	TotalRevenue                int64              `bson:"totalrevenue"`
	TotalRevenueNice            string             `bson:"totalrevenuenice"`
	DebtToEquity                float64            `bson:"debttoequity"`
	RevenuePerShare             float64            `bson:"revenuepershare"`
	RevenuePerShareNice         string             `bson:"revenuepersharenice"`
	ReturnOnAssets              float64            `bson:"returnonassets"`
	ReturnOnAssetsNice          string             `bson:"returnonassetsnice"`
	ReturnOnEquity              float64            `bson:"returnonequity"`
	ReturnOnEquityNice          string             `bson:"returnonequitynice"`
	GrossProfits                int64              `bson:"grossprofits"`
	GrossProfitsNice            string             `bson:"grossprofitsnice"`
	FreeCashflow                int64              `bson:"freecashflow"`
	FreeCashflowNice            string             `bson:"freecashflownice"`
	OperatingCashflow           int64              `bson:"operatingcashflow"`
	OperatingCashflowNice       string             `bson:"operatingcashflownice"`
	GrossMargins                float64            `bson:"grossmargins"`
	GrossMarginsNice            string             `bson:"grossmarginsnice"`
	EbitdaMargins               float64            `bson:"ebitdamargins"`
	EbitdaMarginsNice           string             `bson:"ebitdamarginsnice"`
	OperatingMargins            float64            `bson:"operatingmargins"`
	OperatingMarginsNice        string             `bson:"operatingmarginsnice"`
	ROIC                        float64            `bson:"roic"`
	WorkingCapital              int64              `bson:"workingcapital"`
	EnterpriseToEbit            float64            `bson:"enterprisetoebit"`
	LastUpdated                 time.Time          `bson:"lastupdated"`
}
type recommTrend struct {
	StrongBuy  int64 `bson:"strongbuy"`
	Buy        int64 `bson:"buy"`
	Hold       int64 `bson:"hold"`
	Sell       int64 `bson:"sell"`
	StrongSell int64 `bson:"strongsell"`
}
type cashFlowH struct {
	CapEx                        int64  `bson:"capex"`
	CapExNice                    string `bson:"capexnice"`
	ChangeCash                   int64  `bson:"changecash"`
	ChangeCashNice               string `bson:"changecashnice"`
	ChangeAccountReceivables     int64  `bson:"changeaccountreceivables"`
	ChangeAccountReceivablesNice string `bson:"changeaccountreceivablesnice"`
	ChangeInventory              int64  `bson:"changeinventory"`
	ChangeInventoryNice          string `bson:"changeinventorynice"`
	ChangeLiabilities            int64  `bson:"changeliabilities"`
	ChangeLiabilitiesNice        string `bson:"changeliabilitiesnice"`
	ChangeNetIncome              int64  `bson:"changenetincome"`
	ChangeNetIncomeNice          string `bson:"changenetincomenice"`
	Depreciation                 int64  `bson:"depreciation"`
	DepreciationNice             string `bson:"depreciationnice"`
	EffectExchangeRate           int64  `bson:"effectexchangerate"`
	EffectExchangeRateNice       string `bson:"effectexchangeratenice"`
	EndDate                      string `bson:"enddate"`
	EndDateY                     string `bson:"enddatey"`
	Investments                  int64  `bson:"investments"`
	InvestmentsNice              string `bson:"investmentsnice"`
	NetBorrowings                int64  `bson:"netborrowings"`
	NetBorrowingsNice            string `bson:"netborrowingsnice"`
	NetIncome                    int64  `bson:"netincome"`
	NetIncomeNice                string `bson:"netincomenice"`
	OtherCashFinancing           int64  `bson:"othercashfinancing"`
	OtherCashFinancingNice       string `bson:"othercashfinancingnice"`
	OtherCashInvesting           int64  `bson:"othercashinvesting"`
	OtherCashInvestingNice       string `bson:"othercashinvestingnice"`
	RepurchaseStock              int64  `bson:"repurchasestock"`
	RepurchaseStockNice          string `bson:"repurchasestocknice"`
	TotalCashInvesting           int64  `bson:"totalcashinvesting"`
	TotalCashInvestingNice       string `bson:"totalcashinvestingnice"`
	TotalCashFinancing           int64  `bson:"totalcashfinancing"`
	TotalCashFinancingNice       string `bson:"totalcashfinancingnice"`
	TotalCashOperating           int64  `bson:"totalcashoperating"`
	TotalCashOperatingNice       string `bson:"totalcashoperatingnice"`
}
type incomeH struct {
	TotalRevenue                     int64  `bson:"totalrevenue"`
	TotalRevenueNice                 string `bson:"totalrevenuenice"`
	CostOfRevenue                    int64  `bson:"costofrevenue"`
	CostOfRevenueNice                string `bson:"costofrevenuenice"`
	GrossProfit                      int64  `bson:"grossprofit"`
	GrossProfitNice                  string `bson:"grossprofitnice"`
	ResearchDevelopment              int64  `bson:"researchdevelopment"`
	ResearchDevelopmentNice          string `bson:"researchdevelopmentnice"`
	SellingGeneralAdministrative     int64  `bson:"sellinggeneraladministrative"`
	SellingGeneralAdministrativeNice string `bson:"sellinggeneraladministrativenice"`
	NonRecurring                     int64  `bson:"nonrecurring"`
	NonRecurringNice                 string `bson:"nonrecurringnice"`
	OtherOperatingExpenses           int64  `bson:"otheroperatingexpenses"`
	OtherOperatingExpensesNice       string `bson:"otheroperatingexpensesnice"`
	TotalOperatingExpenses           int64  `bson:"totaloperatingexpenses"`
	TotalOperatingExpensesNice       string `bson:"totaloperatingexpensesnice"`
	EndDate                          string `bson:"enddate"`
	EndDateY                         string `bson:"enddatey"`
	OperatingIncome                  int64  `bson:"operatingincome"`
	OperatingIncomeNice              string `bson:"operatingincomenice"`
	TotalOtherIncomeExpenseNet       int64  `bson:"totalotherincomeexpensenet"`
	TotalOtherIncomeExpenseNetNice   string `bson:"totalotherincomeexpensenetnice"`
	Ebit                             int64  `bson:"ebit"`
	EbitNice                         string `bson:"ebitnice"`
	InterestExpense                  int64  `bson:"interestexpense"`
	InterestExpenseNice              string `bson:"interestexpensenice"`
	IncomeBeforeTax                  int64  `bson:"incomebeforetax"`
	IncomeBeforeTaxNice              string `bson:"incomebeforetaxnice"`
	IncomeTaxExpense                 int64  `bson:"incometaxexpense"`
	IncomeTaxExpenseNice             string `bson:"incometaxexpensenice"`
	MinorityInterest                 int64  `bson:"minorityinterest"`
	MinorityInterestNice             string `bson:"minorityinterestnice"`
	NetIncomeFromContinuingOps       int64  `bson:"netincomefromcontinuingops"`
	NetIncomeFromContinuingOpsNice   string `bson:"netincomefromcontinuingopsnice"`
	DiscontinuedOperations           int64  `bson:"discontinuedoperations"`
	DiscontinuedOperationsNice       string `bson:"discontinuedoperationsnice"`
	ExtraordinaryItems               int64  `bson:"extraordinaryitems"`
	ExtraordinaryItemsNice           string `bson:"extraordinaryitemsnice"`
	EffectOfAccountingCharges        int64  `bson:"effectofaccountingcharges"`
	EffectOfAccountingChargesNice    string `bson:"effectofaccountingchargesnice"`
	OtherItems                       int64  `bson:"otheritems"`
	OtherItemsNice                   string `bson:"otheritemsnice"`
	NetIncome                        int64  `bson:"netincome"`
	NetIncomeNice                    string `bson:"netincomenice"`
	NetIncomeCommonShares            int64  `bson:"netincomecommonshares"`
	NetIncomeCommonSharesNice        string `bson:"netincomecommonsharesnice"`
}
type earningsN struct {
	Date1               string  `bson:"date1"`
	Date2               string  `bson:"date2"`
	EarningsAverage     float64 `bson:"earningsaverage"`
	EarningsAverageNice string  `bson:"earningsaveragenice"`
	EarningsLow         float64 `bson:"earningslow"`
	EarningsLowNice     string  `bson:"earningslownice"`
	EarningsHigh        float64 `bson:"earningshigh"`
	EarningsHighNice    string  `bson:"earningshighnice"`
	RevenueAverage      int64   `bson:"revenueaverage"`
	RevenueAverageNice  string  `bson:"revenueaveragenice"`
	RevenueLow          int64   `bson:"revenuelow"`
	RevenueLowNice      string  `bson:"revenuelownice"`
	RevenueHigh         int64   `bson:"revenuehigh"`
	RevenueHighNice     string  `bson:"revenuehighnice"`
}
type balanceH struct {
	Cash                        int64  `bson:"cash"`
	CashNice                    string `bson:"cashnice"`
	ShortTermInvestments        int64  `bson:"shortterminvestments"`
	ShortTermInvestmentsNice    string `bson:"shortterminvestmentsnice"`
	NetReceivables              int64  `bson:"netreceivables"`
	NetReceivablesNice          string `bson:"netreceivablesnice"`
	Inventory                   int64  `bson:"inventory"`
	InventoryNice               string `bson:"inventorynice"`
	OtherCurrentAssets          int64  `bson:"othercurrentassets"`
	OtherCurrentAssetsNice      string `bson:"othercurrentassetsnice"`
	TotalCurrentAssets          int64  `bson:"totalcurrentassets"`
	TotalCurrentAssetsNice      string `bson:"totalcurrentassetsnice"`
	LongTermInvestments         int64  `bson:"longterminvestments"`
	LongTermInvestmentsNice     string `bson:"longterminvestmentsnice"`
	PropertyPlantEquipment      int64  `bson:"propertyplantequipment"`
	PropertyPlantEquipmentNice  string `bson:"propertyplantequipmentnice"`
	EndDate                     string `bson:"enddate"`
	EndDateY                    string `bson:"enddatey"`
	OtherAssets                 int64  `bson:"otherassets"`
	OtherAssetsNice             string `bson:"otherassetsnice"`
	TotalAssets                 int64  `bson:"totalassets"`
	TotalAssetsNice             string `bson:"totalassetsnice"`
	AccountsPayable             int64  `bson:"accountspayable"`
	AccountsPayableNice         string `bson:"accountspayablenice"`
	ShortLongTermDebt           int64  `bson:"shortlongtermdebt"`
	ShortLongTermDebtNice       string `bson:"shortlongtermdebtnice"`
	OtherCurrentLiab            int64  `bson:"othercurrentliab"`
	OtherCurrentLiabNice        string `bson:"othercurrentliabnice"`
	LongTermDebt                int64  `bson:"longtermdebt"`
	LongTermDebtNice            string `bson:"longtermdebtnice"`
	OtherLiab                   int64  `bson:"otherliab"`
	OtherLiabNice               string `bson:"otherliabnice"`
	TotalCurrentLiabilities     int64  `bson:"totalcurrentliabilities"`
	TotalCurrentLiabilitiesNice string `bson:"totalcurrentliabilitiesnice"`
	TotalLiab                   int64  `bson:"totalliab"`
	TotalLiabNice               string `bson:"totalliabnice"`
	CommonStock                 int64  `bson:"commonstock"`
	CommonStockNice             string `bson:"commonstocknice"`
	RetainedEarnings            int64  `bson:"retainedearnings"`
	RetainedEarningsNice        string `bson:"retainedearningsnice"`
	TreasuryStock               int64  `bson:"treasurystock"`
	TreasuryStockNice           string `bson:"treasurystocknice"`
	OtherStockholderEquity      int64  `bson:"otherstockholderequity"`
	OtherStockholderEquityNice  string `bson:"otherstockholderequitynice"`
	TotalStockholderEquity      int64  `bson:"totalstockholderequity"`
	TotalStockholderEquityNice  string `bson:"totalstockholderequitynice"`
	NetTangibleAssets           int64  `bson:"nettangibleassets"`
	NetTangibleAssetsNice       string `bson:"nettangibleassetsnice"`
}

type Key struct {
	Name string `bson:"name"`
	Key  string `bson:"key"`
}

var stocksDataBase = "stocks"
var stocksColl = "stocks"
var keyColl = "keys"

func MongoDBConnect(mongoDBServerName, mongoDBServerPort, mongoDBAdminUser, mongoDBAdminUserPassword string) (*mongo.Client, context.Context, context.CancelFunc) {
	mongoString := "mongodb://" + mongoDBAdminUser + ":" + mongoDBAdminUserPassword + "@" + mongoDBServerName + ":" + mongoDBServerPort
	ctx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoString))
	if err != nil {
		log.Fatal(err)
	}

	return client, ctx, ctxCancel
}

func SetCompetitors(ticker, exchange string) {
	var competitorsServerName = os.Getenv("COMPETITORS_NAME")
	var competitorsServerPort = os.Getenv("COMPETITORS_PORT")

	var competitorsLink = "http://" + competitorsServerName + ":" + competitorsServerPort + "/competitors?ticker=" + ticker + "&exchange=" + exchange
	//#nosec G107 -- This is a false positive
	_, err := http.Get(competitorsLink)
	if err != nil {
		log.Fatal(err)
	}
}

func NewStock(ca *alphadata.OverviewData, cy *yahoodata.YahooData, dbServer, dbPort, dbUser, dbPass string) {
	client, ctx, ctxCancel := MongoDBConnect(dbServer, dbPort, dbUser, dbPass)

	defer client.Disconnect(ctx)
	defer ctxCancel()

	var stock Stock
	stock.Name = ca.Name
	stock.Ticker = ca.Symbol
	stock.Exchange = ca.Exchange
	stock.Beta, _ = strconv.ParseFloat(cy.QuoteSummary.Result[0].DefaultKeyStatistics.Beta.Fmt, 64)
	stock.Industry = cy.QuoteSummary.Result[0].AssetProfile.Industry
	stock.Address = cy.QuoteSummary.Result[0].AssetProfile.Address1
	stock.City = cy.QuoteSummary.Result[0].AssetProfile.City
	stock.Country = cy.QuoteSummary.Result[0].AssetProfile.Country
	stock.EmployeeNo = cy.QuoteSummary.Result[0].AssetProfile.FullTimeEmployees
	stock.Sector = cy.QuoteSummary.Result[0].AssetProfile.Sector
	stock.RecommTrend = findStockRecomm(cy)
	insertStockDatabyDate(cy, &stock, "CashFlow")
	stock.EnterpriseValue = cy.QuoteSummary.Result[0].DefaultKeyStatistics.EnterpriseValue.Raw
	stock.EnterpriseValueNice = cy.QuoteSummary.Result[0].DefaultKeyStatistics.EnterpriseValue.Fmt
	stock.ForwardPE = cy.QuoteSummary.Result[0].DefaultKeyStatistics.ForwardPE.Raw
	stock.ForwardPENice = cy.QuoteSummary.Result[0].DefaultKeyStatistics.ForwardPE.Fmt
	stock.ProfitMargins = cy.QuoteSummary.Result[0].DefaultKeyStatistics.ProfitMargins.Raw
	stock.ProfitMarginsNice = cy.QuoteSummary.Result[0].DefaultKeyStatistics.ProfitMargins.Fmt
	stock.FloatShares = cy.QuoteSummary.Result[0].DefaultKeyStatistics.FloatShares.Raw
	stock.FloatSharesNice = cy.QuoteSummary.Result[0].DefaultKeyStatistics.FloatShares.Fmt
	stock.SharesOutstanding = cy.QuoteSummary.Result[0].DefaultKeyStatistics.SharesOutstanding.Raw
	stock.SharesOutstandingNice = cy.QuoteSummary.Result[0].DefaultKeyStatistics.SharesOutstanding.Fmt
	stock.SharesShort = cy.QuoteSummary.Result[0].DefaultKeyStatistics.SharesShort.Raw
	stock.SharesShortNice = cy.QuoteSummary.Result[0].DefaultKeyStatistics.SharesShort.Fmt
	stock.HeldPercentInsiders = cy.QuoteSummary.Result[0].DefaultKeyStatistics.HeldPercentInsiders.Raw
	stock.HeldPercentInsidersNice = cy.QuoteSummary.Result[0].DefaultKeyStatistics.HeldPercentInsiders.Fmt
	stock.HeldPercentInstitutions = cy.QuoteSummary.Result[0].DefaultKeyStatistics.HeldPercentInstitutions.Raw
	stock.HeldPercentInstitutionsNice = cy.QuoteSummary.Result[0].DefaultKeyStatistics.HeldPercentInstitutions.Fmt
	stock.ShortRatio = cy.QuoteSummary.Result[0].DefaultKeyStatistics.ShortRatio.Raw
	stock.ShortRatioNice = cy.QuoteSummary.Result[0].DefaultKeyStatistics.ShortRatio.Fmt
	stock.ShortPercentOfFloat = cy.QuoteSummary.Result[0].DefaultKeyStatistics.ShortPercentOfFloat.Raw
	stock.ShortPercentOfFloatNice = cy.QuoteSummary.Result[0].DefaultKeyStatistics.ShortPercentOfFloat.Fmt
	stock.BookValue = cy.QuoteSummary.Result[0].DefaultKeyStatistics.BookValue.Raw
	stock.BookValueNice = cy.QuoteSummary.Result[0].DefaultKeyStatistics.BookValue.Fmt
	stock.PriceToBook = cy.QuoteSummary.Result[0].DefaultKeyStatistics.PriceToBook.Raw
	stock.PriceToBookNice = cy.QuoteSummary.Result[0].DefaultKeyStatistics.PriceToBook.Fmt
	stock.LastFiscalYearEnd = cy.QuoteSummary.Result[0].DefaultKeyStatistics.LastFiscalYearEnd.Fmt
	stock.MostRecentQuarter = cy.QuoteSummary.Result[0].DefaultKeyStatistics.MostRecentQuarter.Fmt
	stock.NetIncomeToCommon = cy.QuoteSummary.Result[0].DefaultKeyStatistics.NetIncomeToCommon.Raw
	stock.NetIncomeToCommonNice = cy.QuoteSummary.Result[0].DefaultKeyStatistics.NetIncomeToCommon.Fmt
	stock.TrailingEps = cy.QuoteSummary.Result[0].DefaultKeyStatistics.TrailingEps.Raw
	stock.TrailingEpsNice = cy.QuoteSummary.Result[0].DefaultKeyStatistics.TrailingEps.Fmt
	stock.ForwardEps = cy.QuoteSummary.Result[0].DefaultKeyStatistics.ForwardEps.Raw
	stock.ForwardEpsNice = cy.QuoteSummary.Result[0].DefaultKeyStatistics.ForwardEps.Fmt
	stock.PegRatio = cy.QuoteSummary.Result[0].DefaultKeyStatistics.PegRatio.Raw
	stock.PegRatioNice = cy.QuoteSummary.Result[0].DefaultKeyStatistics.PegRatio.Fmt
	stock.LastSplitFactor = cy.QuoteSummary.Result[0].DefaultKeyStatistics.LastSplitFactor
	stock.LastSplitDate = cy.QuoteSummary.Result[0].DefaultKeyStatistics.LastSplitDate.Fmt
	stock.EnterpriseToRevenue = cy.QuoteSummary.Result[0].DefaultKeyStatistics.EnterpriseToRevenue.Raw
	stock.EnterpriseToRevenueNice = cy.QuoteSummary.Result[0].DefaultKeyStatistics.EnterpriseToRevenue.Fmt
	stock.EnterpriseToEbitda = cy.QuoteSummary.Result[0].DefaultKeyStatistics.EnterpriseToEbitda.Raw
	stock.EnterpriseToEbitdaNice = cy.QuoteSummary.Result[0].DefaultKeyStatistics.EnterpriseToEbitda.Fmt
	stock.WeekChange52 = cy.QuoteSummary.Result[0].DefaultKeyStatistics.WeekChange52.Raw
	stock.WeekChange52Nice = cy.QuoteSummary.Result[0].DefaultKeyStatistics.WeekChange52.Fmt
	insertStockDatabyDate(cy, &stock, "Income")
	stock.Currency = cy.QuoteSummary.Result[0].SummaryDetail.Currency
	stock.ExDividendDate = cy.QuoteSummary.Result[0].SummaryDetail.ExDividendDate.Fmt
	stock.DividendRate = cy.QuoteSummary.Result[0].SummaryDetail.DividendRate.Raw
	stock.DividendRateNice = cy.QuoteSummary.Result[0].SummaryDetail.DividendRate.Fmt
	stock.DividendYield = cy.QuoteSummary.Result[0].SummaryDetail.DividendYield.Raw
	stock.DividendYieldNice = cy.QuoteSummary.Result[0].SummaryDetail.DividendYield.Fmt
	stock.PayoutRatio = cy.QuoteSummary.Result[0].SummaryDetail.PayoutRatio.Raw
	stock.PayoutRatioNice = cy.QuoteSummary.Result[0].SummaryDetail.PayoutRatio.Fmt
	stock.TrailingPE = cy.QuoteSummary.Result[0].SummaryDetail.TrailingPE.Raw
	stock.TrailingPENice = cy.QuoteSummary.Result[0].SummaryDetail.TrailingPE.Fmt
	stock.MarketCap = cy.QuoteSummary.Result[0].SummaryDetail.MarketCap.Raw
	stock.MarketCapNice = cy.QuoteSummary.Result[0].SummaryDetail.MarketCap.Fmt
	if len(cy.QuoteSummary.Result[0].CalendarEvents.Earnings.EarningsDate) >= 1 {
		stock.EarningsNext.Date1 = cy.QuoteSummary.Result[0].CalendarEvents.Earnings.EarningsDate[0].Fmt
	}
	if len(cy.QuoteSummary.Result[0].CalendarEvents.Earnings.EarningsDate) >= 2 {
		stock.EarningsNext.Date2 = cy.QuoteSummary.Result[0].CalendarEvents.Earnings.EarningsDate[1].Fmt
	}
	stock.EarningsNext.EarningsAverage = cy.QuoteSummary.Result[0].CalendarEvents.Earnings.EarningsAverage.Raw
	stock.EarningsNext.EarningsAverageNice = cy.QuoteSummary.Result[0].CalendarEvents.Earnings.EarningsAverage.Fmt
	stock.EarningsNext.EarningsLow = cy.QuoteSummary.Result[0].CalendarEvents.Earnings.EarningsLow.Raw
	stock.EarningsNext.EarningsLowNice = cy.QuoteSummary.Result[0].CalendarEvents.Earnings.EarningsLow.Fmt
	stock.EarningsNext.EarningsHigh = cy.QuoteSummary.Result[0].CalendarEvents.Earnings.EarningsHigh.Raw
	stock.EarningsNext.EarningsHighNice = cy.QuoteSummary.Result[0].CalendarEvents.Earnings.EarningsHigh.Fmt
	stock.EarningsNext.RevenueAverage = cy.QuoteSummary.Result[0].CalendarEvents.Earnings.RevenueAverage.Raw
	stock.EarningsNext.RevenueAverageNice = cy.QuoteSummary.Result[0].CalendarEvents.Earnings.RevenueAverage.Fmt
	stock.EarningsNext.RevenueLow = cy.QuoteSummary.Result[0].CalendarEvents.Earnings.RevenueLow.Raw
	stock.EarningsNext.RevenueLowNice = cy.QuoteSummary.Result[0].CalendarEvents.Earnings.RevenueLow.Fmt
	stock.EarningsNext.RevenueHigh = cy.QuoteSummary.Result[0].CalendarEvents.Earnings.RevenueHigh.Raw
	stock.EarningsNext.RevenueHighNice = cy.QuoteSummary.Result[0].CalendarEvents.Earnings.RevenueHigh.Fmt
	insertStockDatabyDate(cy, &stock, "Balance")
	stock.Growth5y, stock.Growth5yNice = findGrowth(cy)
	insertStockDatabyDate(cy, &stock, "BalanceQ")
	insertStockDatabyDate(cy, &stock, "IncomeQ")
	insertStockDatabyDate(cy, &stock, "CashFlowQ")
	stock.Price = cy.QuoteSummary.Result[0].FinancialData.CurrentPrice.Raw
	stock.TargetHighPrice = cy.QuoteSummary.Result[0].FinancialData.TargetHighPrice.Raw
	stock.TargetLowPrice = cy.QuoteSummary.Result[0].FinancialData.TargetLowPrice.Raw
	stock.TargetMedianPrice = cy.QuoteSummary.Result[0].FinancialData.TargetMedianPrice.Raw
	stock.RecommendationKey = cy.QuoteSummary.Result[0].FinancialData.RecommendationKey
	stock.TotalCash = cy.QuoteSummary.Result[0].FinancialData.TotalCash.Raw
	stock.TotalCashNice = cy.QuoteSummary.Result[0].FinancialData.TotalCash.Fmt
	stock.TotalCashPerShare = cy.QuoteSummary.Result[0].FinancialData.TotalCashPerShare.Raw
	stock.TotalCashPerShareNice = cy.QuoteSummary.Result[0].FinancialData.TotalCashPerShare.Fmt
	stock.Ebitda = cy.QuoteSummary.Result[0].FinancialData.Ebitda.Raw
	stock.EbitdaNice = cy.QuoteSummary.Result[0].FinancialData.Ebitda.Fmt
	stock.TotalDebt = cy.QuoteSummary.Result[0].FinancialData.TotalDebt.Raw
	stock.TotalDebtNice = cy.QuoteSummary.Result[0].FinancialData.TotalDebt.Fmt
	stock.QuickRatio = cy.QuoteSummary.Result[0].FinancialData.QuickRatio.Raw
	stock.QuickRatioNice = cy.QuoteSummary.Result[0].FinancialData.QuickRatio.Fmt
	stock.CurrentRatio = cy.QuoteSummary.Result[0].FinancialData.CurrentRatio.Raw
	stock.CurrentRatioNice = cy.QuoteSummary.Result[0].FinancialData.CurrentRatio.Fmt
	stock.TotalRevenue = cy.QuoteSummary.Result[0].FinancialData.TotalRevenue.Raw
	stock.TotalRevenueNice = cy.QuoteSummary.Result[0].FinancialData.TotalRevenue.Fmt
	stock.RevenuePerShare = cy.QuoteSummary.Result[0].FinancialData.RevenuePerShare.Raw
	stock.RevenuePerShareNice = cy.QuoteSummary.Result[0].FinancialData.RevenuePerShare.Fmt
	stock.ReturnOnAssets = cy.QuoteSummary.Result[0].FinancialData.ReturnOnAssets.Raw
	stock.ReturnOnAssetsNice = cy.QuoteSummary.Result[0].FinancialData.ReturnOnAssets.Fmt
	stock.ReturnOnEquity = cy.QuoteSummary.Result[0].FinancialData.ReturnOnEquity.Raw
	stock.ReturnOnEquityNice = cy.QuoteSummary.Result[0].FinancialData.ReturnOnEquity.Fmt
	stock.GrossProfits = cy.QuoteSummary.Result[0].FinancialData.GrossProfits.Raw
	stock.GrossProfitsNice = cy.QuoteSummary.Result[0].FinancialData.GrossProfits.Fmt
	stock.FreeCashflow = cy.QuoteSummary.Result[0].FinancialData.FreeCashflow.Raw
	stock.FreeCashflowNice = cy.QuoteSummary.Result[0].FinancialData.FreeCashflow.Fmt
	stock.OperatingCashflow = cy.QuoteSummary.Result[0].FinancialData.OperatingCashflow.Raw
	stock.OperatingCashflowNice = cy.QuoteSummary.Result[0].FinancialData.OperatingCashflow.Fmt
	stock.GrossMargins = cy.QuoteSummary.Result[0].FinancialData.GrossMargins.Raw
	stock.GrossMarginsNice = cy.QuoteSummary.Result[0].FinancialData.GrossMargins.Fmt
	stock.EbitdaMargins = cy.QuoteSummary.Result[0].FinancialData.EbitdaMargins.Raw
	stock.EbitdaMarginsNice = cy.QuoteSummary.Result[0].FinancialData.EbitdaMargins.Fmt
	stock.OperatingMargins = cy.QuoteSummary.Result[0].FinancialData.OperatingMargins.Raw
	stock.OperatingMarginsNice = cy.QuoteSummary.Result[0].FinancialData.OperatingMargins.Fmt
	stock.DebtToEquity = getDE(cy)
	stock.ROIC = getROIC(cy)
	stock.WorkingCapital = getWC(cy)
	stock.EnterpriseToEbit = getEVToEbit(cy)
	stock.LastUpdated = time.Now()

	collection := client.Database(stocksDataBase).Collection(stocksColl)
	_, err := collection.InsertOne(ctx, stock)
	if err != nil {
		log.Fatal(err)
	}
	SetCompetitors(stock.Ticker, stock.Exchange)
}

func FindStock(ticker, dbServer, dbPort, dbUser, dbPass string) bool {

	var stock Stock

	client, ctx, ctxCancel := MongoDBConnect(dbServer, dbPort, dbUser, dbPass)

	defer client.Disconnect(ctx)
	defer ctxCancel()

	collection := client.Database(stocksDataBase).Collection(stocksColl)

	if err := collection.FindOne(ctx, bson.M{"ticker": ticker}).Decode(&stock); err != nil {
		return false
	} else {
		return true
	}
}

func GetStock(ticker, dbServer, dbPort, dbUser, dbPass string) *Stock {

	var stock *Stock

	client, ctx, ctxCancel := MongoDBConnect(dbServer, dbPort, dbUser, dbPass)

	defer client.Disconnect(ctx)
	defer ctxCancel()

	collection := client.Database(stocksDataBase).Collection(stocksColl)

	collection.FindOne(ctx, bson.M{"ticker": ticker}).Decode(&stock)
	return stock
}

func GetKey(name, dbServer, dbPort, dbUser, dbPass string) *Key {

	var key *Key

	client, ctx, ctxCancel := MongoDBConnect(dbServer, dbPort, dbUser, dbPass)

	defer client.Disconnect(ctx)
	defer ctxCancel()

	collection := client.Database(stocksDataBase).Collection(keyColl)

	collection.FindOne(ctx, bson.M{"name": name}).Decode(&key)
	return key
}

func UpdateStock(ca *alphadata.OverviewData, cy *yahoodata.YahooData, ticker, dbServer, dbPort, dbUser, dbPass string) {

	client, ctx, ctxCancel := MongoDBConnect(dbServer, dbPort, dbUser, dbPass)

	defer client.Disconnect(ctx)
	defer ctxCancel()

	currentStock := GetStock(ticker, dbServer, dbPort, dbUser, dbPass)
	if currentStock == nil {
		log.Fatal("Stock " + ticker + " should exist but was not found")
	}

	currentStock.Name = ca.Name
	currentStock.Exchange = ca.Exchange
	currentStock.Beta, _ = strconv.ParseFloat(cy.QuoteSummary.Result[0].DefaultKeyStatistics.Beta.Fmt, 64)
	currentStock.Industry = cy.QuoteSummary.Result[0].AssetProfile.Industry
	currentStock.Address = cy.QuoteSummary.Result[0].AssetProfile.Address1
	currentStock.City = cy.QuoteSummary.Result[0].AssetProfile.City
	currentStock.Country = cy.QuoteSummary.Result[0].AssetProfile.Country
	currentStock.EmployeeNo = cy.QuoteSummary.Result[0].AssetProfile.FullTimeEmployees
	currentStock.Sector = cy.QuoteSummary.Result[0].AssetProfile.Sector
	currentStock.RecommTrend = findStockRecomm(cy)
	insertStockDatabyDate(cy, currentStock, "CashFlow")
	currentStock.EnterpriseValue = cy.QuoteSummary.Result[0].DefaultKeyStatistics.EnterpriseValue.Raw
	currentStock.EnterpriseValueNice = cy.QuoteSummary.Result[0].DefaultKeyStatistics.EnterpriseValue.Fmt
	currentStock.ForwardPE = cy.QuoteSummary.Result[0].DefaultKeyStatistics.ForwardPE.Raw
	currentStock.ForwardPENice = cy.QuoteSummary.Result[0].DefaultKeyStatistics.ForwardPE.Fmt
	currentStock.ProfitMargins = cy.QuoteSummary.Result[0].DefaultKeyStatistics.ProfitMargins.Raw
	currentStock.ProfitMarginsNice = cy.QuoteSummary.Result[0].DefaultKeyStatistics.ProfitMargins.Fmt
	currentStock.FloatShares = cy.QuoteSummary.Result[0].DefaultKeyStatistics.FloatShares.Raw
	currentStock.FloatSharesNice = cy.QuoteSummary.Result[0].DefaultKeyStatistics.FloatShares.Fmt
	currentStock.SharesOutstanding = cy.QuoteSummary.Result[0].DefaultKeyStatistics.SharesOutstanding.Raw
	currentStock.SharesOutstandingNice = cy.QuoteSummary.Result[0].DefaultKeyStatistics.SharesOutstanding.Fmt
	currentStock.SharesShort = cy.QuoteSummary.Result[0].DefaultKeyStatistics.SharesShort.Raw
	currentStock.SharesShortNice = cy.QuoteSummary.Result[0].DefaultKeyStatistics.SharesShort.Fmt
	currentStock.HeldPercentInsiders = cy.QuoteSummary.Result[0].DefaultKeyStatistics.HeldPercentInsiders.Raw
	currentStock.HeldPercentInsidersNice = cy.QuoteSummary.Result[0].DefaultKeyStatistics.HeldPercentInsiders.Fmt
	currentStock.HeldPercentInstitutions = cy.QuoteSummary.Result[0].DefaultKeyStatistics.HeldPercentInstitutions.Raw
	currentStock.HeldPercentInstitutionsNice = cy.QuoteSummary.Result[0].DefaultKeyStatistics.HeldPercentInstitutions.Fmt
	currentStock.ShortRatio = cy.QuoteSummary.Result[0].DefaultKeyStatistics.ShortRatio.Raw
	currentStock.ShortRatioNice = cy.QuoteSummary.Result[0].DefaultKeyStatistics.ShortRatio.Fmt
	currentStock.ShortPercentOfFloat = cy.QuoteSummary.Result[0].DefaultKeyStatistics.ShortPercentOfFloat.Raw
	currentStock.ShortPercentOfFloatNice = cy.QuoteSummary.Result[0].DefaultKeyStatistics.ShortPercentOfFloat.Fmt
	currentStock.BookValue = cy.QuoteSummary.Result[0].DefaultKeyStatistics.BookValue.Raw
	currentStock.BookValueNice = cy.QuoteSummary.Result[0].DefaultKeyStatistics.BookValue.Fmt
	currentStock.PriceToBook = cy.QuoteSummary.Result[0].DefaultKeyStatistics.PriceToBook.Raw
	currentStock.PriceToBookNice = cy.QuoteSummary.Result[0].DefaultKeyStatistics.PriceToBook.Fmt
	currentStock.LastFiscalYearEnd = cy.QuoteSummary.Result[0].DefaultKeyStatistics.LastFiscalYearEnd.Fmt
	currentStock.MostRecentQuarter = cy.QuoteSummary.Result[0].DefaultKeyStatistics.MostRecentQuarter.Fmt
	currentStock.NetIncomeToCommon = cy.QuoteSummary.Result[0].DefaultKeyStatistics.NetIncomeToCommon.Raw
	currentStock.NetIncomeToCommonNice = cy.QuoteSummary.Result[0].DefaultKeyStatistics.NetIncomeToCommon.Fmt
	currentStock.TrailingEps = cy.QuoteSummary.Result[0].DefaultKeyStatistics.TrailingEps.Raw
	currentStock.TrailingEpsNice = cy.QuoteSummary.Result[0].DefaultKeyStatistics.TrailingEps.Fmt
	currentStock.ForwardEps = cy.QuoteSummary.Result[0].DefaultKeyStatistics.ForwardEps.Raw
	currentStock.ForwardEpsNice = cy.QuoteSummary.Result[0].DefaultKeyStatistics.ForwardEps.Fmt
	currentStock.PegRatio = cy.QuoteSummary.Result[0].DefaultKeyStatistics.PegRatio.Raw
	currentStock.PegRatioNice = cy.QuoteSummary.Result[0].DefaultKeyStatistics.PegRatio.Fmt
	currentStock.LastSplitFactor = cy.QuoteSummary.Result[0].DefaultKeyStatistics.LastSplitFactor
	currentStock.LastSplitDate = cy.QuoteSummary.Result[0].DefaultKeyStatistics.LastSplitDate.Fmt
	currentStock.EnterpriseToRevenue = cy.QuoteSummary.Result[0].DefaultKeyStatistics.EnterpriseToRevenue.Raw
	currentStock.EnterpriseToRevenueNice = cy.QuoteSummary.Result[0].DefaultKeyStatistics.EnterpriseToRevenue.Fmt
	currentStock.EnterpriseToEbitda = cy.QuoteSummary.Result[0].DefaultKeyStatistics.EnterpriseToEbitda.Raw
	currentStock.EnterpriseToEbitdaNice = cy.QuoteSummary.Result[0].DefaultKeyStatistics.EnterpriseToEbitda.Fmt
	currentStock.WeekChange52 = cy.QuoteSummary.Result[0].DefaultKeyStatistics.WeekChange52.Raw
	currentStock.WeekChange52Nice = cy.QuoteSummary.Result[0].DefaultKeyStatistics.WeekChange52.Fmt
	insertStockDatabyDate(cy, currentStock, "Income")
	currentStock.Currency = cy.QuoteSummary.Result[0].SummaryDetail.Currency
	currentStock.ExDividendDate = cy.QuoteSummary.Result[0].SummaryDetail.ExDividendDate.Fmt
	currentStock.DividendRate = cy.QuoteSummary.Result[0].SummaryDetail.DividendRate.Raw
	currentStock.DividendRateNice = cy.QuoteSummary.Result[0].SummaryDetail.DividendRate.Fmt
	currentStock.DividendYield = cy.QuoteSummary.Result[0].SummaryDetail.DividendYield.Raw
	currentStock.DividendYieldNice = cy.QuoteSummary.Result[0].SummaryDetail.DividendYield.Fmt
	currentStock.PayoutRatio = cy.QuoteSummary.Result[0].SummaryDetail.PayoutRatio.Raw
	currentStock.PayoutRatioNice = cy.QuoteSummary.Result[0].SummaryDetail.PayoutRatio.Fmt
	currentStock.TrailingPE = cy.QuoteSummary.Result[0].SummaryDetail.TrailingPE.Raw
	currentStock.TrailingPENice = cy.QuoteSummary.Result[0].SummaryDetail.TrailingPE.Fmt
	currentStock.MarketCap = cy.QuoteSummary.Result[0].SummaryDetail.MarketCap.Raw
	currentStock.MarketCapNice = cy.QuoteSummary.Result[0].SummaryDetail.MarketCap.Fmt
	if len(cy.QuoteSummary.Result[0].CalendarEvents.Earnings.EarningsDate) >= 1 {
		currentStock.EarningsNext.Date1 = cy.QuoteSummary.Result[0].CalendarEvents.Earnings.EarningsDate[0].Fmt
	}
	if len(cy.QuoteSummary.Result[0].CalendarEvents.Earnings.EarningsDate) >= 2 {
		currentStock.EarningsNext.Date2 = cy.QuoteSummary.Result[0].CalendarEvents.Earnings.EarningsDate[1].Fmt
	}
	currentStock.EarningsNext.EarningsAverage = cy.QuoteSummary.Result[0].CalendarEvents.Earnings.EarningsAverage.Raw
	currentStock.EarningsNext.EarningsAverageNice = cy.QuoteSummary.Result[0].CalendarEvents.Earnings.EarningsAverage.Fmt
	currentStock.EarningsNext.EarningsLow = cy.QuoteSummary.Result[0].CalendarEvents.Earnings.EarningsLow.Raw
	currentStock.EarningsNext.EarningsLowNice = cy.QuoteSummary.Result[0].CalendarEvents.Earnings.EarningsLow.Fmt
	currentStock.EarningsNext.EarningsHigh = cy.QuoteSummary.Result[0].CalendarEvents.Earnings.EarningsHigh.Raw
	currentStock.EarningsNext.EarningsHighNice = cy.QuoteSummary.Result[0].CalendarEvents.Earnings.EarningsHigh.Fmt
	currentStock.EarningsNext.RevenueAverage = cy.QuoteSummary.Result[0].CalendarEvents.Earnings.RevenueAverage.Raw
	currentStock.EarningsNext.RevenueAverageNice = cy.QuoteSummary.Result[0].CalendarEvents.Earnings.RevenueAverage.Fmt
	currentStock.EarningsNext.RevenueLow = cy.QuoteSummary.Result[0].CalendarEvents.Earnings.RevenueLow.Raw
	currentStock.EarningsNext.RevenueLowNice = cy.QuoteSummary.Result[0].CalendarEvents.Earnings.RevenueLow.Fmt
	currentStock.EarningsNext.RevenueHigh = cy.QuoteSummary.Result[0].CalendarEvents.Earnings.RevenueHigh.Raw
	currentStock.EarningsNext.RevenueHighNice = cy.QuoteSummary.Result[0].CalendarEvents.Earnings.RevenueHigh.Fmt
	insertStockDatabyDate(cy, currentStock, "Balance")
	currentStock.Growth5y, currentStock.Growth5yNice = findGrowth(cy)
	insertStockDatabyDate(cy, currentStock, "BalanceQ")
	insertStockDatabyDate(cy, currentStock, "IncomeQ")
	insertStockDatabyDate(cy, currentStock, "CashFlowQ")
	currentStock.Price = cy.QuoteSummary.Result[0].FinancialData.CurrentPrice.Raw
	currentStock.TargetHighPrice = cy.QuoteSummary.Result[0].FinancialData.TargetHighPrice.Raw
	currentStock.TargetLowPrice = cy.QuoteSummary.Result[0].FinancialData.TargetLowPrice.Raw
	currentStock.TargetMedianPrice = cy.QuoteSummary.Result[0].FinancialData.TargetMedianPrice.Raw
	currentStock.RecommendationKey = cy.QuoteSummary.Result[0].FinancialData.RecommendationKey
	currentStock.TotalCash = cy.QuoteSummary.Result[0].FinancialData.TotalCash.Raw
	currentStock.TotalCashNice = cy.QuoteSummary.Result[0].FinancialData.TotalCash.Fmt
	currentStock.TotalCashPerShare = cy.QuoteSummary.Result[0].FinancialData.TotalCashPerShare.Raw
	currentStock.TotalCashPerShareNice = cy.QuoteSummary.Result[0].FinancialData.TotalCashPerShare.Fmt
	currentStock.Ebitda = cy.QuoteSummary.Result[0].FinancialData.Ebitda.Raw
	currentStock.EbitdaNice = cy.QuoteSummary.Result[0].FinancialData.Ebitda.Fmt
	currentStock.TotalDebt = cy.QuoteSummary.Result[0].FinancialData.TotalDebt.Raw
	currentStock.TotalDebtNice = cy.QuoteSummary.Result[0].FinancialData.TotalDebt.Fmt
	currentStock.QuickRatio = cy.QuoteSummary.Result[0].FinancialData.QuickRatio.Raw
	currentStock.QuickRatioNice = cy.QuoteSummary.Result[0].FinancialData.QuickRatio.Fmt
	currentStock.CurrentRatio = cy.QuoteSummary.Result[0].FinancialData.CurrentRatio.Raw
	currentStock.CurrentRatioNice = cy.QuoteSummary.Result[0].FinancialData.CurrentRatio.Fmt
	currentStock.TotalRevenue = cy.QuoteSummary.Result[0].FinancialData.TotalRevenue.Raw
	currentStock.TotalRevenueNice = cy.QuoteSummary.Result[0].FinancialData.TotalRevenue.Fmt
	currentStock.RevenuePerShare = cy.QuoteSummary.Result[0].FinancialData.RevenuePerShare.Raw
	currentStock.RevenuePerShareNice = cy.QuoteSummary.Result[0].FinancialData.RevenuePerShare.Fmt
	currentStock.ReturnOnAssets = cy.QuoteSummary.Result[0].FinancialData.ReturnOnAssets.Raw
	currentStock.ReturnOnAssetsNice = cy.QuoteSummary.Result[0].FinancialData.ReturnOnAssets.Fmt
	currentStock.ReturnOnEquity = cy.QuoteSummary.Result[0].FinancialData.ReturnOnEquity.Raw
	currentStock.ReturnOnEquityNice = cy.QuoteSummary.Result[0].FinancialData.ReturnOnEquity.Fmt
	currentStock.GrossProfits = cy.QuoteSummary.Result[0].FinancialData.GrossProfits.Raw
	currentStock.GrossProfitsNice = cy.QuoteSummary.Result[0].FinancialData.GrossProfits.Fmt
	currentStock.FreeCashflow = cy.QuoteSummary.Result[0].FinancialData.FreeCashflow.Raw
	currentStock.FreeCashflowNice = cy.QuoteSummary.Result[0].FinancialData.FreeCashflow.Fmt
	currentStock.OperatingCashflow = cy.QuoteSummary.Result[0].FinancialData.OperatingCashflow.Raw
	currentStock.OperatingCashflowNice = cy.QuoteSummary.Result[0].FinancialData.OperatingCashflow.Fmt
	currentStock.GrossMargins = cy.QuoteSummary.Result[0].FinancialData.GrossMargins.Raw
	currentStock.GrossMarginsNice = cy.QuoteSummary.Result[0].FinancialData.GrossMargins.Fmt
	currentStock.EbitdaMargins = cy.QuoteSummary.Result[0].FinancialData.EbitdaMargins.Raw
	currentStock.EbitdaMarginsNice = cy.QuoteSummary.Result[0].FinancialData.EbitdaMargins.Fmt
	currentStock.OperatingMargins = cy.QuoteSummary.Result[0].FinancialData.OperatingMargins.Raw
	currentStock.OperatingMarginsNice = cy.QuoteSummary.Result[0].FinancialData.OperatingMargins.Fmt
	currentStock.DebtToEquity = getDE(cy)
	currentStock.ROIC = getROIC(cy)
	currentStock.WorkingCapital = getWC(cy)
	currentStock.EnterpriseToEbit = getEVToEbit(cy)
	currentStock.LastUpdated = time.Now()

	pByte, err := bson.Marshal(currentStock)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database(stocksDataBase).Collection(stocksColl)
	filter := bson.M{"ticker": bson.M{"$eq": ticker}}
	var update bson.M
	err = bson.Unmarshal(pByte, &update)
	if err != nil {
		log.Fatal(err)
	}
	_, err = collection.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: update}})

	if err != nil {
		log.Fatal(err)
	}

	SetCompetitors(currentStock.Ticker, currentStock.Exchange)
}

func findStockRecomm(cy *yahoodata.YahooData) recommTrend {
	var ret recommTrend

	for _, elem := range cy.QuoteSummary.Result[0].RecommendationTrend.Trend {
		if elem.Period == "0m" {
			ret.StrongBuy = elem.StrongBuy
			ret.Buy = elem.Buy
			ret.Hold = elem.Hold
			ret.StrongSell = elem.StrongSell
			ret.Sell = elem.Sell

			break
		}
	}

	return ret
}

func findGrowth(cy *yahoodata.YahooData) (float64, string) {
	var r1 float64
	var r2 string
	for _, elem := range cy.QuoteSummary.Result[0].EarningsTrend.Trend {
		if elem.Period == "+5y" {
			r1 = elem.Growth.Raw
			r2 = elem.Growth.Fmt
		}
	}
	return r1, r2
}

func getDE(cy *yahoodata.YahooData) float64 {

	n1 := cy.QuoteSummary.Result[0].BalanceSheetHistoryQuarterly.BalanceSheetStatements[0].TotalLiab.Raw
	n2 := cy.QuoteSummary.Result[0].BalanceSheetHistoryQuarterly.BalanceSheetStatements[0].TotalStockholderEquity.Raw

	return math.Round((float64(n1)/float64(n2))*100) / 100
}

func getROIC(cy *yahoodata.YahooData) float64 {

	ebit := cy.QuoteSummary.Result[0].IncomeStatementHistory.IncomeStatementHistory[0].Ebit.Raw
	taxexpense := cy.QuoteSummary.Result[0].IncomeStatementHistory.IncomeStatementHistory[0].IncomeTaxExpense.Raw
	longdebt := cy.QuoteSummary.Result[0].BalanceSheetHistoryQuarterly.BalanceSheetStatements[0].LongTermDebt.Raw
	stockequity := cy.QuoteSummary.Result[0].BalanceSheetHistoryQuarterly.BalanceSheetStatements[0].TotalStockholderEquity.Raw
	cash := cy.QuoteSummary.Result[0].BalanceSheetHistoryQuarterly.BalanceSheetStatements[0].Cash.Raw

	return (math.Round(((float64(ebit)-float64(taxexpense))/(float64(longdebt)+float64(stockequity)-float64(cash)))*100) / 100) * 100
}

func getWC(cy *yahoodata.YahooData) int64 {

	cl := cy.QuoteSummary.Result[0].BalanceSheetHistoryQuarterly.BalanceSheetStatements[0].TotalCurrentLiabilities.Raw
	ca := cy.QuoteSummary.Result[0].BalanceSheetHistoryQuarterly.BalanceSheetStatements[0].TotalCurrentAssets.Raw

	return ca - cl
}

func getEVToEbit(cy *yahoodata.YahooData) float64 {
	ev := cy.QuoteSummary.Result[0].DefaultKeyStatistics.EnterpriseValue.Raw
	ebit := cy.QuoteSummary.Result[0].IncomeStatementHistory.IncomeStatementHistory[0].Ebit.Raw

	return math.Round(float64(ev)/float64(ebit)*100) / 100
}

func getYear(endDate string) string {
	date, _ := time.Parse("2006-01-02", endDate)
	//subtract one month
	date = date.AddDate(0, -2, -15)
	year := date.Year()
	return strconv.FormatInt(int64(year), 10)
}
func getQuarter(endDate string) string {
	date, _ := time.Parse("2006-01-02", endDate)
	//subtract one month and a couple of days
	date = date.AddDate(0, -2, -1)
	year := strconv.FormatInt(int64(date.Year()), 10)
	month := int(date.Month())
	quarter := "Q4"
	if month >= 1 && month <= 3 {
		quarter = "Q1"
	} else if month >= 4 && month <= 6 {
		quarter = "Q2"
	} else if month >= 7 && month <= 9 {
		quarter = "Q3"
	}
	return year + "-" + quarter
}

func insertStockDatabyDate(cy *yahoodata.YahooData, s *Stock, t string) {

	switch t {
	case "CashFlow":
		for i := len(cy.QuoteSummary.Result[0].CashflowStatementHistory.CashflowStatements) - 1; i >= 0; i-- {
			elem := cy.QuoteSummary.Result[0].CashflowStatementHistory.CashflowStatements[i]
			found := false
			for i, elemdb := range s.CashFlowH {
				s.CashFlowH[i].EndDateY = getYear(elemdb.EndDate)
				if elem.EndDate.Fmt == elemdb.EndDate {
					found = true
					break
				}
			}
			if !found {
				var c cashFlowH
				c.CapEx = elem.CapitalExpenditures.Raw
				c.CapExNice = elem.CapitalExpenditures.Fmt
				c.ChangeCash = elem.ChangeInCash.Raw
				c.ChangeCashNice = elem.ChangeInCash.Fmt
				c.ChangeAccountReceivables = elem.ChangeToAccountReceivables.Raw
				c.ChangeAccountReceivablesNice = elem.ChangeToAccountReceivables.Fmt
				c.ChangeInventory = elem.ChangeToInventory.Raw
				c.ChangeInventoryNice = elem.ChangeToInventory.Fmt
				c.ChangeLiabilities = elem.ChangeToLiabilities.Raw
				c.ChangeLiabilitiesNice = elem.ChangeToLiabilities.Fmt
				c.ChangeNetIncome = elem.ChangeToNetincome.Raw
				c.ChangeNetIncomeNice = elem.ChangeToNetincome.Fmt
				c.Depreciation = elem.Depreciation.Raw
				c.DepreciationNice = elem.Depreciation.Fmt
				c.EffectExchangeRate = elem.EffectOfExchangeRate.Raw
				c.EffectExchangeRateNice = elem.EffectOfExchangeRate.Fmt
				c.EndDate = elem.EndDate.Fmt
				c.Investments = elem.Investments.Raw
				c.InvestmentsNice = elem.Investments.Fmt
				c.NetBorrowings = elem.NetBorrowings.Raw
				c.NetBorrowingsNice = elem.NetBorrowings.Fmt
				c.NetIncome = elem.NetIncome.Raw
				c.NetIncomeNice = elem.NetIncome.Fmt
				c.OtherCashFinancing = elem.OtherCashflowsFromFinancingActivities.Raw
				c.OtherCashFinancingNice = elem.OtherCashflowsFromFinancingActivities.Fmt
				c.OtherCashInvesting = elem.OtherCashflowsFromInvestingActivities.Raw
				c.OtherCashInvestingNice = elem.OtherCashflowsFromInvestingActivities.Fmt
				c.RepurchaseStock = elem.RepurchaseOfStock.Raw
				c.RepurchaseStockNice = elem.RepurchaseOfStock.Fmt
				c.TotalCashInvesting = elem.TotalCashflowsFromInvestingActivities.Raw
				c.TotalCashInvestingNice = elem.TotalCashflowsFromInvestingActivities.Fmt
				c.TotalCashFinancing = elem.TotalCashFromFinancingActivities.Raw
				c.TotalCashFinancingNice = elem.TotalCashFromFinancingActivities.Fmt
				c.TotalCashOperating = elem.TotalCashFromOperatingActivities.Raw
				c.TotalCashOperatingNice = elem.TotalCashFromOperatingActivities.Fmt
				c.EndDateY = getYear(elem.EndDate.Fmt)

				if c.TotalCashOperating != 0 {
					s.CashFlowH = append(s.CashFlowH, c)
				}
			}
		}
	case "CashFlowQ":
		for i := len(cy.QuoteSummary.Result[0].CashflowStatementHistoryQuarterly.CashflowStatements) - 1; i >= 0; i-- {
			elem := cy.QuoteSummary.Result[0].CashflowStatementHistoryQuarterly.CashflowStatements[i]
			found := false
			for i, elemdb := range s.CashFlowHQ {
				s.CashFlowHQ[i].EndDateY = getQuarter(elemdb.EndDate)
				if elem.EndDate.Fmt == elemdb.EndDate {
					found = true
					break
				}
			}
			if !found {
				var c cashFlowH
				c.CapEx = elem.CapitalExpenditures.Raw
				c.CapExNice = elem.CapitalExpenditures.Fmt
				c.ChangeCash = elem.ChangeInCash.Raw
				c.ChangeCashNice = elem.ChangeInCash.Fmt
				c.ChangeAccountReceivables = elem.ChangeToAccountReceivables.Raw
				c.ChangeAccountReceivablesNice = elem.ChangeToAccountReceivables.Fmt
				c.ChangeInventory = elem.ChangeToInventory.Raw
				c.ChangeInventoryNice = elem.ChangeToInventory.Fmt
				c.ChangeLiabilities = elem.ChangeToLiabilities.Raw
				c.ChangeLiabilitiesNice = elem.ChangeToLiabilities.Fmt
				c.ChangeNetIncome = elem.ChangeToNetincome.Raw
				c.ChangeNetIncomeNice = elem.ChangeToNetincome.Fmt
				c.Depreciation = elem.Depreciation.Raw
				c.DepreciationNice = elem.Depreciation.Fmt
				c.EffectExchangeRate = elem.EffectOfExchangeRate.Raw
				c.EffectExchangeRateNice = elem.EffectOfExchangeRate.Fmt
				c.EndDate = elem.EndDate.Fmt
				c.Investments = elem.Investments.Raw
				c.InvestmentsNice = elem.Investments.Fmt
				c.NetBorrowings = elem.NetBorrowings.Raw
				c.NetBorrowingsNice = elem.NetBorrowings.Fmt
				c.NetIncome = elem.NetIncome.Raw
				c.NetIncomeNice = elem.NetIncome.Fmt
				c.OtherCashFinancing = elem.OtherCashflowsFromFinancingActivities.Raw
				c.OtherCashFinancingNice = elem.OtherCashflowsFromFinancingActivities.Fmt
				c.OtherCashInvesting = elem.OtherCashflowsFromInvestingActivities.Raw
				c.OtherCashInvestingNice = elem.OtherCashflowsFromInvestingActivities.Fmt
				c.RepurchaseStock = elem.RepurchaseOfStock.Raw
				c.RepurchaseStockNice = elem.RepurchaseOfStock.Fmt
				c.TotalCashInvesting = elem.TotalCashflowsFromInvestingActivities.Raw
				c.TotalCashInvestingNice = elem.TotalCashflowsFromInvestingActivities.Fmt
				c.TotalCashFinancing = elem.TotalCashFromFinancingActivities.Raw
				c.TotalCashFinancingNice = elem.TotalCashFromFinancingActivities.Fmt
				c.TotalCashOperating = elem.TotalCashFromOperatingActivities.Raw
				c.TotalCashOperatingNice = elem.TotalCashFromOperatingActivities.Fmt
				c.EndDateY = getQuarter(elem.EndDate.Fmt)

				if c.TotalCashOperating != 0 {
					s.CashFlowHQ = append(s.CashFlowHQ, c)
				}
			}
		}
	case "Income":
		for i := len(cy.QuoteSummary.Result[0].IncomeStatementHistory.IncomeStatementHistory) - 1; i >= 0; i-- {
			elem := cy.QuoteSummary.Result[0].IncomeStatementHistory.IncomeStatementHistory[i]
			found := false
			for i, elemdb := range s.IncomeH {
				s.IncomeH[i].EndDateY = getYear(elemdb.EndDate)
				if elem.EndDate.Fmt == elemdb.EndDate {
					found = true
					break
				}
			}
			if !found {
				var c incomeH
				c.TotalRevenue = elem.TotalRevenue.Raw
				c.TotalRevenueNice = elem.TotalRevenue.Fmt
				c.CostOfRevenue = elem.CostOfRevenue.Raw
				c.CostOfRevenueNice = elem.CostOfRevenue.Fmt
				c.GrossProfit = elem.GrossProfit.Raw
				c.GrossProfitNice = elem.GrossProfit.Fmt
				c.ResearchDevelopment = elem.ResearchDevelopment.Raw
				c.ResearchDevelopmentNice = elem.ResearchDevelopment.Fmt
				c.SellingGeneralAdministrative = elem.SellingGeneralAdministrative.Raw
				c.SellingGeneralAdministrativeNice = elem.SellingGeneralAdministrative.Fmt
				c.NonRecurring = elem.NonRecurring.Raw
				c.NonRecurringNice = elem.NonRecurring.Fmt
				c.OtherOperatingExpenses = elem.OtherOperatingExpenses.Raw
				c.OtherOperatingExpensesNice = elem.OtherOperatingExpenses.Fmt
				c.TotalOperatingExpenses = elem.TotalOperatingExpenses.Raw
				c.TotalOperatingExpensesNice = elem.TotalOperatingExpenses.Fmt
				c.EndDate = elem.EndDate.Fmt
				c.OperatingIncome = elem.OperatingIncome.Raw
				c.OperatingIncomeNice = elem.OperatingIncome.Fmt
				c.TotalOtherIncomeExpenseNet = elem.TotalOtherIncomeExpenseNet.Raw
				c.TotalOtherIncomeExpenseNetNice = elem.TotalOtherIncomeExpenseNet.Fmt
				c.Ebit = elem.Ebit.Raw
				c.EbitNice = elem.Ebit.Fmt
				c.InterestExpense = elem.InterestExpense.Raw
				c.InterestExpenseNice = elem.InterestExpense.Fmt
				c.IncomeBeforeTax = elem.IncomeBeforeTax.Raw
				c.IncomeBeforeTaxNice = elem.IncomeBeforeTax.Fmt
				c.IncomeTaxExpense = elem.IncomeTaxExpense.Raw
				c.IncomeTaxExpenseNice = elem.IncomeTaxExpense.Fmt
				c.MinorityInterest = elem.MinorityInterest.Raw
				c.MinorityInterestNice = elem.MinorityInterest.Fmt
				c.NetIncomeFromContinuingOps = elem.NetIncomeFromContinuingOps.Raw
				c.NetIncomeFromContinuingOpsNice = elem.NetIncomeFromContinuingOps.Fmt
				c.DiscontinuedOperations = elem.DiscontinuedOperations.Raw
				c.DiscontinuedOperationsNice = elem.DiscontinuedOperations.Fmt
				c.ExtraordinaryItems = elem.ExtraordinaryItems.Raw
				c.ExtraordinaryItemsNice = elem.ExtraordinaryItems.Fmt
				c.EffectOfAccountingCharges = elem.EffectOfAccountingCharges.Raw
				c.EffectOfAccountingChargesNice = elem.EffectOfAccountingCharges.Fmt
				c.OtherItems = elem.OtherItems.Raw
				c.OtherItemsNice = elem.OtherItems.Fmt
				c.NetIncome = elem.NetIncome.Raw
				c.NetIncomeNice = elem.NetIncome.Fmt
				c.NetIncomeCommonShares = elem.NetIncomeApplicableToCommonShares.Raw
				c.NetIncomeCommonSharesNice = elem.NetIncomeApplicableToCommonShares.Fmt
				c.EndDateY = getYear(elem.EndDate.Fmt)

				if c.NetIncome != 0 {
					s.IncomeH = append(s.IncomeH, c)
				}
			}
		}
	case "IncomeQ":
		for i := len(cy.QuoteSummary.Result[0].IncomeStatementHistoryQuarterly.IncomeStatementHistory) - 1; i >= 0; i-- {
			elem := cy.QuoteSummary.Result[0].IncomeStatementHistoryQuarterly.IncomeStatementHistory[i]
			found := false
			for i, elemdb := range s.IncomeHQ {
				s.IncomeHQ[i].EndDateY = getQuarter(elemdb.EndDate)
				if elem.EndDate.Fmt == elemdb.EndDate {
					found = true
					break
				}
			}
			if !found {
				var c incomeH
				c.TotalRevenue = elem.TotalRevenue.Raw
				c.TotalRevenueNice = elem.TotalRevenue.Fmt
				c.CostOfRevenue = elem.CostOfRevenue.Raw
				c.CostOfRevenueNice = elem.CostOfRevenue.Fmt
				c.GrossProfit = elem.GrossProfit.Raw
				c.GrossProfitNice = elem.GrossProfit.Fmt
				c.ResearchDevelopment = elem.ResearchDevelopment.Raw
				c.ResearchDevelopmentNice = elem.ResearchDevelopment.Fmt
				c.SellingGeneralAdministrative = elem.SellingGeneralAdministrative.Raw
				c.SellingGeneralAdministrativeNice = elem.SellingGeneralAdministrative.Fmt
				c.NonRecurring = elem.NonRecurring.Raw
				c.NonRecurringNice = elem.NonRecurring.Fmt
				c.OtherOperatingExpenses = elem.OtherOperatingExpenses.Raw
				c.OtherOperatingExpensesNice = elem.OtherOperatingExpenses.Fmt
				c.TotalOperatingExpenses = elem.TotalOperatingExpenses.Raw
				c.TotalOperatingExpensesNice = elem.TotalOperatingExpenses.Fmt
				c.EndDate = elem.EndDate.Fmt
				c.OperatingIncome = elem.OperatingIncome.Raw
				c.OperatingIncomeNice = elem.OperatingIncome.Fmt
				c.TotalOtherIncomeExpenseNet = elem.TotalOtherIncomeExpenseNet.Raw
				c.TotalOtherIncomeExpenseNetNice = elem.TotalOtherIncomeExpenseNet.Fmt
				c.Ebit = elem.Ebit.Raw
				c.EbitNice = elem.Ebit.Fmt
				c.InterestExpense = elem.InterestExpense.Raw
				c.InterestExpenseNice = elem.InterestExpense.Fmt
				c.IncomeBeforeTax = elem.IncomeBeforeTax.Raw
				c.IncomeBeforeTaxNice = elem.IncomeBeforeTax.Fmt
				c.IncomeTaxExpense = elem.IncomeTaxExpense.Raw
				c.IncomeTaxExpenseNice = elem.IncomeTaxExpense.Fmt
				c.MinorityInterest = elem.MinorityInterest.Raw
				c.MinorityInterestNice = elem.MinorityInterest.Fmt
				c.NetIncomeFromContinuingOps = elem.NetIncomeFromContinuingOps.Raw
				c.NetIncomeFromContinuingOpsNice = elem.NetIncomeFromContinuingOps.Fmt
				c.DiscontinuedOperations = elem.DiscontinuedOperations.Raw
				c.DiscontinuedOperationsNice = elem.DiscontinuedOperations.Fmt
				c.ExtraordinaryItems = elem.ExtraordinaryItems.Raw
				c.ExtraordinaryItemsNice = elem.ExtraordinaryItems.Fmt
				c.EffectOfAccountingCharges = elem.EffectOfAccountingCharges.Raw
				c.EffectOfAccountingChargesNice = elem.EffectOfAccountingCharges.Fmt
				c.OtherItems = elem.OtherItems.Raw
				c.OtherItemsNice = elem.OtherItems.Fmt
				c.NetIncome = elem.NetIncome.Raw
				c.NetIncomeNice = elem.NetIncome.Fmt
				c.NetIncomeCommonShares = elem.NetIncomeApplicableToCommonShares.Raw
				c.NetIncomeCommonSharesNice = elem.NetIncomeApplicableToCommonShares.Fmt
				c.EndDateY = getQuarter(elem.EndDate.Fmt)

				if c.NetIncome != 0 {
					s.IncomeHQ = append(s.IncomeHQ, c)
				}
			}
		}
	case "Balance":
		for i := len(cy.QuoteSummary.Result[0].BalanceSheetHistory.BalanceSheetStatements) - 1; i >= 0; i-- {
			elem := cy.QuoteSummary.Result[0].BalanceSheetHistory.BalanceSheetStatements[i]
			found := false
			for i, elemdb := range s.BalanceH {
				s.BalanceH[i].EndDateY = getYear(elemdb.EndDate)
				if elem.EndDate.Fmt == elemdb.EndDate {
					found = true
					break
				}
			}
			if !found {
				var c balanceH
				c.Cash = elem.Cash.Raw
				c.CashNice = elem.Cash.Fmt
				c.ShortTermInvestments = elem.ShortTermInvestments.Raw
				c.ShortTermInvestmentsNice = elem.ShortTermInvestments.Fmt
				c.NetReceivables = elem.NetReceivables.Raw
				c.NetReceivablesNice = elem.NetReceivables.Fmt
				c.Inventory = elem.Inventory.Raw
				c.InventoryNice = elem.Inventory.Fmt
				c.OtherCurrentAssets = elem.OtherCurrentAssets.Raw
				c.OtherCurrentAssetsNice = elem.OtherCurrentAssets.Fmt
				c.TotalCurrentAssets = elem.TotalCurrentAssets.Raw
				c.TotalCurrentAssetsNice = elem.TotalCurrentAssets.Fmt
				c.LongTermInvestments = elem.LongTermInvestments.Raw
				c.LongTermInvestmentsNice = elem.LongTermInvestments.Fmt
				c.PropertyPlantEquipment = elem.PropertyPlantEquipment.Raw
				c.PropertyPlantEquipmentNice = elem.PropertyPlantEquipment.Fmt
				c.EndDate = elem.EndDate.Fmt
				c.OtherAssets = elem.OtherAssets.Raw
				c.OtherAssetsNice = elem.OtherAssets.Fmt
				c.TotalAssets = elem.TotalAssets.Raw
				c.TotalAssetsNice = elem.TotalAssets.Fmt
				c.AccountsPayable = elem.AccountsPayable.Raw
				c.AccountsPayableNice = elem.AccountsPayable.Fmt
				c.ShortLongTermDebt = elem.ShortLongTermDebt.Raw
				c.ShortLongTermDebtNice = elem.ShortLongTermDebt.Fmt
				c.OtherCurrentLiab = elem.OtherCurrentLiab.Raw
				c.OtherCurrentLiabNice = elem.OtherCurrentLiab.Fmt
				c.LongTermDebt = elem.LongTermDebt.Raw
				c.LongTermDebtNice = elem.LongTermDebt.Fmt
				c.OtherLiab = elem.OtherLiab.Raw
				c.OtherLiabNice = elem.OtherLiab.Fmt
				c.TotalCurrentLiabilities = elem.TotalCurrentLiabilities.Raw
				c.TotalCurrentLiabilitiesNice = elem.TotalCurrentLiabilities.Fmt
				c.TotalLiab = elem.TotalLiab.Raw
				c.TotalLiabNice = elem.TotalLiab.Fmt
				c.CommonStock = elem.CommonStock.Raw
				c.CommonStockNice = elem.CommonStock.Fmt
				c.RetainedEarnings = elem.RetainedEarnings.Raw
				c.RetainedEarningsNice = elem.RetainedEarnings.Fmt
				c.TreasuryStock = elem.TreasuryStock.Raw
				c.TreasuryStockNice = elem.TreasuryStock.Fmt
				c.OtherStockholderEquity = elem.OtherStockholderEquity.Raw
				c.OtherStockholderEquityNice = elem.OtherStockholderEquity.Fmt
				c.TotalStockholderEquity = elem.TotalStockholderEquity.Raw
				c.TotalStockholderEquityNice = elem.TotalStockholderEquity.Fmt
				c.NetTangibleAssets = elem.NetTangibleAssets.Raw
				c.NetTangibleAssetsNice = elem.NetTangibleAssets.Fmt
				c.EndDateY = getYear(elem.EndDate.Fmt)

				if c.Cash != 0 {
					s.BalanceH = append(s.BalanceH, c)
				}
			}
		}
	case "BalanceQ":
		for i := len(cy.QuoteSummary.Result[0].BalanceSheetHistoryQuarterly.BalanceSheetStatements) - 1; i >= 0; i-- {
			elem := cy.QuoteSummary.Result[0].BalanceSheetHistoryQuarterly.BalanceSheetStatements[i]
			found := false
			for i, elemdb := range s.BalanceHQ {
				s.BalanceHQ[i].EndDateY = getQuarter(elemdb.EndDate)
				if elem.EndDate.Fmt == elemdb.EndDate {
					found = true
					break
				}
			}
			if !found {
				var c balanceH
				c.Cash = elem.Cash.Raw
				c.CashNice = elem.Cash.Fmt
				c.ShortTermInvestments = elem.ShortTermInvestments.Raw
				c.ShortTermInvestmentsNice = elem.ShortTermInvestments.Fmt
				c.NetReceivables = elem.NetReceivables.Raw
				c.NetReceivablesNice = elem.NetReceivables.Fmt
				c.Inventory = elem.Inventory.Raw
				c.InventoryNice = elem.Inventory.Fmt
				c.OtherCurrentAssets = elem.OtherCurrentAssets.Raw
				c.OtherCurrentAssetsNice = elem.OtherCurrentAssets.Fmt
				c.TotalCurrentAssets = elem.TotalCurrentAssets.Raw
				c.TotalCurrentAssetsNice = elem.TotalCurrentAssets.Fmt
				c.LongTermInvestments = elem.LongTermInvestments.Raw
				c.LongTermInvestmentsNice = elem.LongTermInvestments.Fmt
				c.PropertyPlantEquipment = elem.PropertyPlantEquipment.Raw
				c.PropertyPlantEquipmentNice = elem.PropertyPlantEquipment.Fmt
				c.EndDate = elem.EndDate.Fmt
				c.OtherAssets = elem.OtherAssets.Raw
				c.OtherAssetsNice = elem.OtherAssets.Fmt
				c.TotalAssets = elem.TotalAssets.Raw
				c.TotalAssetsNice = elem.TotalAssets.Fmt
				c.AccountsPayable = elem.AccountsPayable.Raw
				c.AccountsPayableNice = elem.AccountsPayable.Fmt
				c.ShortLongTermDebt = elem.ShortLongTermDebt.Raw
				c.ShortLongTermDebtNice = elem.ShortLongTermDebt.Fmt
				c.OtherCurrentLiab = elem.OtherCurrentLiab.Raw
				c.OtherCurrentLiabNice = elem.OtherCurrentLiab.Fmt
				c.LongTermDebt = elem.LongTermDebt.Raw
				c.LongTermDebtNice = elem.LongTermDebt.Fmt
				c.OtherLiab = elem.OtherLiab.Raw
				c.OtherLiabNice = elem.OtherLiab.Fmt
				c.TotalCurrentLiabilities = elem.TotalCurrentLiabilities.Raw
				c.TotalCurrentLiabilitiesNice = elem.TotalCurrentLiabilities.Fmt
				c.TotalLiab = elem.TotalLiab.Raw
				c.TotalLiabNice = elem.TotalLiab.Fmt
				c.CommonStock = elem.CommonStock.Raw
				c.CommonStockNice = elem.CommonStock.Fmt
				c.RetainedEarnings = elem.RetainedEarnings.Raw
				c.RetainedEarningsNice = elem.RetainedEarnings.Fmt
				c.TreasuryStock = elem.TreasuryStock.Raw
				c.TreasuryStockNice = elem.TreasuryStock.Fmt
				c.OtherStockholderEquity = elem.OtherStockholderEquity.Raw
				c.OtherStockholderEquityNice = elem.OtherStockholderEquity.Fmt
				c.TotalStockholderEquity = elem.TotalStockholderEquity.Raw
				c.TotalStockholderEquityNice = elem.TotalStockholderEquity.Fmt
				c.NetTangibleAssets = elem.NetTangibleAssets.Raw
				c.NetTangibleAssetsNice = elem.NetTangibleAssets.Fmt
				c.EndDateY = getQuarter(elem.EndDate.Fmt)

				if c.Cash != 0 {
					s.BalanceHQ = append(s.BalanceHQ, c)
				}
			}
		}
	}
}
