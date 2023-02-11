package yahoodata

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var YBASEURL = "https://yfapi.net/v11/finance/quoteSummary/<Ticker>?modules=price,summaryDetail,earningsTrend,earnings,earningsHistory,defaultKeyStatistics,esgScores,quoteType,majorHoldersBreakdown,majorDirectHolders,fundOwnership,balanceSheetHistoryQuarterly,recommendationTrend,institutionOwnership,upgradeDowngradeHistory,sectorTrend,indexTrend,balanceSheetHistory,cashflowStatementHistory,cashflowStatementHistoryQuarterly,incomeStatementHistoryQuarterly,incomeStatementHistory,calendarEvents,financialData,assetProfile"

type YahooData struct {
	QuoteSummary yahooDataResult `json:"quoteSummary"`
}
type yahooDataResult struct {
	Result []yahooDataResultObj `json:"result"`
	Error  errorObj             `json:"error"`
}
type errorObj struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}
type yahooDataResultObj struct {
	AssetProfile                      yahooDataAssetProfileObj         `json:"assetProfile"`
	RecommendationTrend               yahooDataRecommTrendObj          `json:"recommendationTrend"`
	CashflowStatementHistory          yahooDataCashFlowStmH            `json:"cashflowStatementHistory"`
	DefaultKeyStatistics              yahooDataDefaultKeyStatisticsObj `json:"defaultKeyStatistics"`
	IncomeStatementHistory            yahooDataIncomeStmH              `json:"incomeStatementHistory"`
	SummaryDetail                     yahooDataSummaryDatailObj        `json:"summaryDetail"`
	CalendarEvents                    yahooDataCalendarEventsObj       `json:"calendarEvents"`
	BalanceSheetHistory               yahooDataBalanceSheetStmH        `json:"balanceSheetHistory"`
	EarningsTrend                     yahooDataEarningsTrendObj        `json:"earningsTrend"`
	BalanceSheetHistoryQuarterly      yahooDataBalanceSheetStmH        `json:"balanceSheetHistoryQuarterly"`
	IncomeStatementHistoryQuarterly   yahooDataIncomeStmH              `json:"incomeStatementHistoryQuarterly"`
	CashflowStatementHistoryQuarterly yahooDataCashFlowStmH            `json:"cashflowStatementHistoryQuarterly"`
	FinancialData                     yahooDataFinancialDataObj        `json:"financialData"`
	Price                             yahooDataPriceObj                `json:"price"`
}
type yahooDataAssetProfileObj struct {
	Address1            string `json:"address1"`
	City                string `json:"city"`
	Country             string `json:"country"`
	FullTimeEmployees   int64  `json:"fullTimeEmployees"`
	Industry            string `json:"industry"`
	LongBusinessSummary string `json:"longBusinessSummary"`
	Sector              string `json:"sector"`
	State               string `json:"state"`
}
type yahooDataRecommTrendObj struct {
	Trend []struct {
		Period     string `json:"period"`
		StrongBuy  int64  `json:"strongBuy"`
		Buy        int64  `json:"buy"`
		Hold       int64  `json:"hold"`
		Sell       int64  `json:"sell"`
		StrongSell int64  `json:"strongSell"`
	} `json:"trend"`
}
type yahooDataCashFlowStmH struct {
	CashflowStatements []struct {
		CapitalExpenditures struct {
			Fmt     string `json:"fmt"`
			LongFmt string `json:"longFmt"`
			Raw     int64  `json:"raw"`
		} `json:"capitalExpenditures"`
		ChangeInCash struct {
			Fmt     string `json:"fmt"`
			LongFmt string `json:"longFmt"`
			Raw     int64  `json:"raw"`
		} `json:"changeInCash"`
		ChangeToAccountReceivables struct {
			Fmt     string `json:"fmt"`
			LongFmt string `json:"longFmt"`
			Raw     int64  `json:"raw"`
		} `json:"changeToAccountReceivables"`
		ChangeToInventory struct {
			Fmt     string `json:"fmt"`
			LongFmt string `json:"longFmt"`
			Raw     int64  `json:"raw"`
		} `json:"changeToInventory"`
		ChangeToLiabilities struct {
			Fmt     string `json:"fmt"`
			LongFmt string `json:"longFmt"`
			Raw     int64  `json:"raw"`
		} `json:"changeToLiabilities"`
		ChangeToNetincome struct {
			Fmt     string `json:"fmt"`
			LongFmt string `json:"longFmt"`
			Raw     int64  `json:"raw"`
		} `json:"changeToNetincome"`
		Depreciation struct {
			Fmt     string `json:"fmt"`
			LongFmt string `json:"longFmt"`
			Raw     int64  `json:"raw"`
		} `json:"depreciation"`
		EffectOfExchangeRate struct {
			Fmt     string `json:"fmt"`
			LongFmt string `json:"longFmt"`
			Raw     int64  `json:"raw"`
		} `json:"effectOfExchangeRate"`
		EndDate struct {
			Fmt string `json:'fmt'`
			Raw int64  `json:'raw'`
		} `json:'endDate'`
		Investments struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'investments'`
		NetBorrowings struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'netBorrowings'`
		NetIncome struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'netIncome'`
		OtherCashflowsFromFinancingActivities struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'otherCashflowsFromFinancingActivities'`
		OtherCashflowsFromInvestingActivities struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'otherCashflowsFromInvestingActivities'`
		RepurchaseOfStock struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'repurchaseOfStock'`
		TotalCashflowsFromInvestingActivities struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'totalCashflowsFromInvestingActivities'`
		TotalCashFromFinancingActivities struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'totalCashFromFinancingActivities'`
		TotalCashFromOperatingActivities struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'totalCashFromOperatingActivities'`
	} `json:"cashflowStatements"`
}
type yahooDataDefaultKeyStatisticsObj struct {
	PriceHint struct {
		Fmt     string `json:'fmt'`
		LongFmt string `json:'longFmt'`
		Raw     int64  `json:'raw'`
	} `json:'priceHint'`
	EnterpriseValue struct {
		Fmt     string `json:'fmt'`
		LongFmt string `json:'longFmt'`
		Raw     int64  `json:'raw'`
	} `json:'enterpriseValue'`
	ForwardPE struct {
		Fmt string  `json:'fmt'`
		Raw float64 `json:'raw'`
	} `json:'forwardPE'`
	ProfitMargins struct {
		Fmt string  `json:'fmt'`
		Raw float64 `json:'raw'`
	} `json:'profitMargins'`
	FloatShares struct {
		Fmt     string `json:'fmt'`
		LongFmt string `json:'longFmt'`
		Raw     int64  `json:'raw'`
	} `json:'floatShares'`
	SharesOutstanding struct {
		Fmt     string `json:'fmt'`
		LongFmt string `json:'longFmt'`
		Raw     int64  `json:'raw'`
	} `json:'sharesOutstanding'`
	SharesShort struct {
		Fmt     string `json:'fmt'`
		LongFmt string `json:'longFmt'`
		Raw     int64  `json:'raw'`
	} `json:'sharesShort'`
	SharesShortPriorMonth struct {
		Fmt     string `json:'fmt'`
		LongFmt string `json:'longFmt'`
		Raw     int64  `json:'raw'`
	} `json:'sharesShortPriorMonth'`
	SharesShortPreviousMonthDate struct {
		Fmt string `json:'fmt'`
		Raw int64  `json:'raw'`
	} `json:'sharesShortPreviousMonthDate'`
	DateShortInterest struct {
		Fmt string `json:'fmt'`
		Raw int64  `json:'raw'`
	} `json:'dateShortInterest'`
	SharesPercentSharesOut struct {
		Fmt string  `json:'fmt'`
		Raw float64 `json:'raw'`
	} `json:'sharesPercentSharesOut'`
	HeldPercentInsiders struct {
		Fmt string  `json:'fmt'`
		Raw float64 `json:'raw'`
	} `json:'heldPercentInsiders'`
	HeldPercentInstitutions struct {
		Fmt string  `json:'fmt'`
		Raw float64 `json:'raw'`
	} `json:'heldPercentInstitutions'`
	ShortRatio struct {
		Fmt string  `json:'fmt'`
		Raw float64 `json:'raw'`
	} `json:'shortRatio'`
	ShortPercentOfFloat struct {
		Fmt string  `json:'fmt'`
		Raw float64 `json:'raw'`
	} `json:'shortPercentOfFloat'`
	Beta struct {
		Fmt string  `json:'fmt'`
		Raw float64 `json:'raw'`
	} `json:'beta'`
	ImpliedSharesOutstanding struct {
		Fmt     string `json:'fmt'`
		LongFmt string `json:'longFmt'`
		Raw     int64  `json:'raw'`
	} `json:'impliedSharesOutstanding'`
	BookValue struct {
		Fmt string  `json:'fmt'`
		Raw float64 `json:'raw'`
	} `json:'bookValue'`
	PriceToBook struct {
		Fmt string  `json:'fmt'`
		Raw float64 `json:'raw'`
	} `json:'priceToBook'`
	LastFiscalYearEnd struct {
		Fmt string `json:'fmt'`
		Raw int64  `json:'raw'`
	} `json:'lastFiscalYearEnd'`
	NextFiscalYearEnd struct {
		Fmt string `json:'fmt'`
		Raw int64  `json:'raw'`
	} `json:'nextFiscalYearEnd'`
	MostRecentQuarter struct {
		Fmt string `json:'fmt'`
		Raw int64  `json:'raw'`
	} `json:'mostRecentQuarter'`
	EarningsQuarterlyGrowth struct {
		Fmt string  `json:'fmt'`
		Raw float64 `json:'raw'`
	} `json:'earningsQuarterlyGrowth'`
	NetIncomeToCommon struct {
		Fmt     string `json:'fmt'`
		LongFmt string `json:'longFmt'`
		Raw     int64  `json:'raw'`
	} `json:'netIncomeToCommon'`
	TrailingEps struct {
		Fmt string  `json:'fmt'`
		Raw float64 `json:'raw'`
	} `json:'trailingEps'`
	ForwardEps struct {
		Fmt string  `json:'fmt'`
		Raw float64 `json:'raw'`
	} `json:'forwardEps'`
	PegRatio struct {
		Fmt string  `json:'fmt'`
		Raw float64 `json:'raw'`
	} `json:'pegRatio'`
	LastSplitFactor string `json:'lastSplitFactor'`
	LastSplitDate   struct {
		Fmt string `json:'fmt'`
		Raw int64  `json:'raw'`
	} `json:'lastSplitDate'`
	EnterpriseToRevenue struct {
		Fmt string  `json:'fmt'`
		Raw float64 `json:'raw'`
	} `json:'enterpriseToRevenue'`
	EnterpriseToEbitda struct {
		Fmt string  `json:'fmt'`
		Raw float64 `json:'raw'`
	} `json:'enterpriseToEbitda'`
	WeekChange52 struct {
		Fmt string  `json:'fmt'`
		Raw float64 `json:'raw'`
	} `json:"52WeekChange"`
	SandP52WeekChange struct {
		Fmt string  `json:'fmt'`
		Raw float64 `json:'raw'`
	} `json:'SandP52WeekChange'`
	LastDividendValue struct {
		Fmt string  `json:'fmt'`
		Raw float64 `json:'raw'`
	} `json:'lastDividendValue'`
	LastDividendDate struct {
		Fmt string `json:'fmt'`
		Raw int64  `json:'raw'`
	} `json:'lastDividendDate'`
}
type yahooDataIncomeStmH struct {
	IncomeStatementHistory []struct {
		TotalRevenue struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'totalRevenue'`
		CostOfRevenue struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'costOfRevenue'`
		GrossProfit struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'grossProfit'`
		ResearchDevelopment struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'researchDevelopment'`
		SellingGeneralAdministrative struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'sellingGeneralAdministrative'`
		NonRecurring struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'nonRecurring'`
		OtherOperatingExpenses struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'otherOperatingExpenses'`
		TotalOperatingExpenses struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'totalOperatingExpenses'`
		EndDate struct {
			Fmt string `json:'fmt'`
			Raw int64  `json:'raw'`
		} `json:'endDate'`
		OperatingIncome struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'operatingIncome'`
		TotalOtherIncomeExpenseNet struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'totalOtherIncomeExpenseNet'`
		Ebit struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'ebit'`
		InterestExpense struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'interestExpense'`
		IncomeBeforeTax struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'incomeBeforeTax'`
		IncomeTaxExpense struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'incomeTaxExpense'`
		MinorityInterest struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'minorityInterest'`
		NetIncomeFromContinuingOps struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'netIncomeFromContinuingOps'`
		DiscontinuedOperations struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'discontinuedOperations'`
		ExtraordinaryItems struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'extraordinaryItems'`
		EffectOfAccountingCharges struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'effectOfAccountingCharges'`
		OtherItems struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'otherItems'`
		NetIncome struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'netIncome'`
		NetIncomeApplicableToCommonShares struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'netIncomeApplicableToCommonShares'`
	} `json:'incomeStatementHistory'`
}
type yahooDataSummaryDatailObj struct {
	Currency       string `json:'currency'`
	ExDividendDate struct {
		Fmt string `json:'fmt'`
		Raw int64  `json:'raw'`
	} `json:'exDividendDate'`
	DividendRate struct {
		Fmt string  `json:'fmt'`
		Raw float64 `json:'raw'`
	} `json:'dividendRate'`
	DividendYield struct {
		Fmt string  `json:'fmt'`
		Raw float64 `json:'raw'`
	} `json:'dividendYield'`
	PayoutRatio struct {
		Fmt string  `json:'fmt'`
		Raw float64 `json:'raw'`
	} `json:'payoutRatio'`
	Beta struct {
		Fmt string  `json:'fmt'`
		Raw float64 `json:'raw'`
	} `json:'beta'`
	TrailingPE struct {
		Fmt string  `json:'fmt'`
		Raw float64 `json:'raw'`
	} `json:'trailingPE'`
	ForwardPE struct {
		Fmt string  `json:'fmt'`
		Raw float64 `json:'raw'`
	} `json:'forwardPE'`
	MarketCap struct {
		Fmt     string `json:'fmt'`
		LongFmt string `json:'longFmt'`
		Raw     int64  `json:'raw'`
	} `json:'marketCap'`
	FiftyTwoWeekLow struct {
		Fmt string  `json:'fmt'`
		Raw float64 `json:'raw'`
	} `json:'fiftyTwoWeekLow'`
	FiftyTwoWeekHigh struct {
		Fmt string  `json:'fmt'`
		Raw float64 `json:'raw'`
	} `json:'fiftyTwoWeekHigh'`
}
type yahooDataCalendarEventsObj struct {
	Earnings struct {
		EarningsDate []struct {
			Fmt string `json:'fmt'`
			Raw int64  `json:'raw'`
		} `json:'earningsDate'`
		EarningsAverage struct {
			Fmt string  `json:'fmt'`
			Raw float64 `json:'raw'`
		} `json:'earningsAverage'`
		EarningsLow struct {
			Fmt string  `json:'fmt'`
			Raw float64 `json:'raw'`
		} `json:'earningsLow'`
		EarningsHigh struct {
			Fmt string  `json:'fmt'`
			Raw float64 `json:'raw'`
		} `json:'earningsHigh'`
		RevenueAverage struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'revenueAverage'`
		RevenueLow struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'revenueLow'`
		RevenueHigh struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'revenueHigh'`
	} `json:'earnings'`
}
type yahooDataBalanceSheetStmH struct {
	BalanceSheetStatements []struct {
		Cash struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'cash'`
		ShortTermInvestments struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'shortTermInvestments'`
		NetReceivables struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'netReceivables'`
		Inventory struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'inventory'`
		OtherCurrentAssets struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'otherCurrentAssets'`
		TotalCurrentAssets struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'totalCurrentAssets'`
		LongTermInvestments struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'longTermInvestments'`
		PropertyPlantEquipment struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'propertyPlantEquipment'`
		EndDate struct {
			Fmt string `json:'fmt'`
			Raw int64  `json:'raw'`
		} `json:'endDate'`
		OtherAssets struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'otherAssets'`
		TotalAssets struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'totalAssets'`
		AccountsPayable struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'accountsPayable'`
		ShortLongTermDebt struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'shortLongTermDebt'`
		OtherCurrentLiab struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'otherCurrentLiab'`
		LongTermDebt struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'longTermDebt'`
		OtherLiab struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'otherLiab'`
		TotalCurrentLiabilities struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'totalCurrentLiabilities'`
		TotalLiab struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'totalLiab'`
		CommonStock struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'commonStock'`
		RetainedEarnings struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'retainedEarnings'`
		TreasuryStock struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'treasuryStock'`
		OtherStockholderEquity struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'otherStockholderEquity'`
		TotalStockholderEquity struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'totalStockholderEquity'`
		NetTangibleAssets struct {
			Fmt     string `json:'fmt'`
			LongFmt string `json:'longFmt'`
			Raw     int64  `json:'raw'`
		} `json:'netTangibleAssets'`
	} `json:'balanceSheetStatements'`
}
type yahooDataEarningsTrendObj struct {
	Trend []struct {
		Period  string `json:'period'`
		EndDate string `json:'endDate'`
		Growth  struct {
			Fmt string  `json:'fmt'`
			Raw float64 `json:'raw'`
		} `json:'growth'`
	} `json:'trend'`
}
type yahooDataFinancialDataObj struct {
	CurrentPrice struct {
		Fmt string  `json:'fmt'`
		Raw float64 `json:'raw'`
	} `json:'currentPrice'`
	TargetHighPrice struct {
		Fmt string  `json:'fmt'`
		Raw float64 `json:'raw'`
	} `json:'targetHighPrice'`
	TargetLowPrice struct {
		Fmt string  `json:'fmt'`
		Raw float64 `json:'raw'`
	} `json:'targetLowPrice'`
	TargetMedianPrice struct {
		Fmt string  `json:'fmt'`
		Raw float64 `json:'raw'`
	} `json:'targetMedianPrice'`
	RecommendationMean struct {
		Fmt string  `json:'fmt'`
		Raw float64 `json:'raw'`
	} `json:'recommendationMean'`
	RecommendationKey       string `json:'recommendationKey'`
	NumberOfAnalystOpinions struct {
		Fmt     string `json:'fmt'`
		LongFmt string `json:'longFmt'`
		Raw     int64  `json:'raw'`
	} `json:'numberOfAnalystOpinions'`
	TotalCash struct {
		Fmt     string `json:'fmt'`
		LongFmt string `json:'longFmt'`
		Raw     int64  `json:'raw'`
	} `json:'totalCash'`
	TotalCashPerShare struct {
		Fmt string  `json:'fmt'`
		Raw float64 `json:'raw'`
	} `json:'totalCashPerShare'`
	Ebitda struct {
		Fmt     string `json:'fmt'`
		LongFmt string `json:'longFmt'`
		Raw     int64  `json:'raw'`
	} `json:'ebitda'`
	TotalDebt struct {
		Fmt     string `json:'fmt'`
		LongFmt string `json:'longFmt'`
		Raw     int64  `json:'raw'`
	} `json:'totalDebt'`
	QuickRatio struct {
		Fmt string  `json:'fmt'`
		Raw float64 `json:'raw'`
	} `json:'quickRatio'`
	CurrentRatio struct {
		Fmt string  `json:'fmt'`
		Raw float64 `json:'raw'`
	} `json:'currentRatio'`
	TotalRevenue struct {
		Fmt     string `json:'fmt'`
		LongFmt string `json:'longFmt'`
		Raw     int64  `json:'raw'`
	} `json:'totalRevenue'`
	DebtToEquity struct {
		Fmt string  `json:'fmt'`
		Raw float64 `json:'raw'`
	} `json:'debtToEquity'`
	RevenuePerShare struct {
		Fmt string  `json:'fmt'`
		Raw float64 `json:'raw'`
	} `json:'revenuePerShare'`
	ReturnOnAssets struct {
		Fmt string  `json:'fmt'`
		Raw float64 `json:'raw'`
	} `json:'returnOnAssets'`
	ReturnOnEquity struct {
		Fmt string  `json:'fmt'`
		Raw float64 `json:'raw'`
	} `json:'returnOnEquity'`
	GrossProfits struct {
		Fmt     string `json:'fmt'`
		LongFmt string `json:'longFmt'`
		Raw     int64  `json:'raw'`
	} `json:'grossProfits'`
	FreeCashflow struct {
		Fmt     string `json:'fmt'`
		LongFmt string `json:'longFmt'`
		Raw     int64  `json:'raw'`
	} `json:'freeCashflow'`
	OperatingCashflow struct {
		Fmt     string `json:'fmt'`
		LongFmt string `json:'longFmt'`
		Raw     int64  `json:'raw'`
	} `json:'operatingCashflow'`
	GrossMargins struct {
		Fmt string  `json:'fmt'`
		Raw float64 `json:'raw'`
	} `json:'grossMargins'`
	EbitdaMargins struct {
		Fmt string  `json:'fmt'`
		Raw float64 `json:'raw'`
	} `json:'ebitdaMargins'`
	OperatingMargins struct {
		Fmt string  `json:"fmt"`
		Raw float64 `json:"raw"`
	} `json:"operatingMargins"`
	ProfitMargins struct {
		Fmt string  `json:"fmt"`
		Raw float64 `json:"raw"`
	} `json:"profitMargins"`
}
type yahooDataPriceObj struct {
	ExchangeName string `json:"exchangeName"`
	Symbol       string `json:"symbol"`
	ShortName    string `json:"shortName"`
	LongName     string `json:"longName"`
}

func NewData(apikey string, ticker string) *YahooData {
	p := new(YahooData)

	yLink := strings.Replace(YBASEURL, "<Ticker>", ticker, -1)

	req, err := http.NewRequest("GET", yLink, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("x-api-key", apikey)
	client := &http.Client{}
	yresp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if yresp.StatusCode != http.StatusOK {
		return nil
	}

	ybody, err := ioutil.ReadAll(yresp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(ybody, &p)
	if err != nil {
		log.Fatal(err)
	}

	return p
}
