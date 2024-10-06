package datastructures

import "encoding/json"

type JsonSlice []string

func (o *JsonSlice) UnmarshalJSON(data []byte) error {
	// Try to unmarshal data as []string
	var arr []string
	if err := json.Unmarshal(data, &arr); err == nil {
		*o = arr
		return nil
	}

	// If that fails, try to unmarshal data as a string
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}

	// Unmarshal the string as JSON array
	if err := json.Unmarshal([]byte(str), &arr); err != nil {
		return err
	}

	*o = arr
	return nil
}

type ClobReward struct {
	ID               string `json:"id"`
	ConditionID      string `json:"conditionId"`
	AssetAddress     string `json:"assetAddress"`
	RewardsAmount    int    `json:"rewardsAmount"`
	RewardsDailyRate int    `json:"rewardsDailyRate"`
	StartDate        string `json:"startDate"`
	EndDate          string `json:"endDate"`
}

type Market struct {
	ID                           string       `json:"id"`
	Question                     string       `json:"question"`
	ConditionID                  string       `json:"conditionId"`
	Slug                         string       `json:"slug"`
	ResolutionSource             string       `json:"resolutionSource"`
	EndDate                      string       `json:"endDate"`
	Liquidity                    float64      `json:"liquidity,string"`
	StartDate                    string       `json:"startDate"`
	Image                        string       `json:"image"`
	Icon                         string       `json:"icon"`
	Description                  string       `json:"description"`
	Outcomes                     JsonSlice    `json:"outcomes"`
	OutcomePrices                JsonSlice    `json:"outcomePrices"`
	Volume                       float64      `json:"volume,string"`
	Active                       bool         `json:"active"`
	Closed                       bool         `json:"closed"`
	MarketMakerAddress           string       `json:"marketMakerAddress"`
	CreatedAt                    string       `json:"createdAt"`
	UpdatedAt                    string       `json:"updatedAt"`
	New                          bool         `json:"new"`
	Featured                     bool         `json:"featured"`
	SubmittedBy                  string       `json:"submitted_by"`
	Archived                     bool         `json:"archived"`
	ResolvedBy                   string       `json:"resolvedBy"`
	Restricted                   bool         `json:"restricted"`
	GroupItemTitle               string       `json:"groupItemTitle"`
	GroupItemThreshold           string       `json:"groupItemThreshold"`
	QuestionID                   string       `json:"questionID"`
	EnableOrderBook              bool         `json:"enableOrderBook"`
	OrderPriceMinTickSize        float64      `json:"orderPriceMinTickSize"`
	OrderMinSize                 int          `json:"orderMinSize"`
	VolumeNum                    float64      `json:"volumeNum"`
	LiquidityNum                 float64      `json:"liquidityNum"`
	EndDateIso                   string       `json:"endDateIso"`
	StartDateIso                 string       `json:"startDateIso"`
	HasReviewedDates             bool         `json:"hasReviewedDates"`
	Volume24hr                   float64      `json:"volume24hr"`
	ClobTokenIds                 JsonSlice    `json:"clobTokenIds"`
	UmaBond                      float64      `json:"umaBond,string"`
	UmaReward                    float64      `json:"umaReward,string"`
	Volume24hrClob               float64      `json:"volume24hrClob"`
	VolumeClob                   float64      `json:"volumeClob"`
	LiquidityClob                float64      `json:"liquidityClob"`
	AcceptingOrders              bool         `json:"acceptingOrders"`
	NegRisk                      bool         `json:"negRisk"`
	NegRiskMarketID              string       `json:"negRiskMarketID"`
	NegRiskRequestID             string       `json:"negRiskRequestID"`
	CommentCount                 int          `json:"commentCount"`
	Sync                         bool         `json:"_sync"`
	Ready                        bool         `json:"ready"`
	Funded                       bool         `json:"funded"`
	AcceptingOrdersTimestamp     string       `json:"acceptingOrdersTimestamp"`
	Cyom                         bool         `json:"cyom"`
	Competitive                  float64      `json:"competitive"`
	PagerDutyNotificationEnabled bool         `json:"pagerDutyNotificationEnabled"`
	Approved                     bool         `json:"approved"`
	ClobRewards                  []ClobReward `json:"clobRewards"`
	RewardsMinSize               int          `json:"rewardsMinSize"`
	RewardsMaxSpread             float64      `json:"rewardsMaxSpread"`
	Spread                       float64      `json:"spread"`
	OneDayPriceChange            float64      `json:"oneDayPriceChange"`
	LastTradePrice               float64      `json:"lastTradePrice"`
	BestBid                      float64      `json:"bestBid"`
	BestAsk                      float64      `json:"bestAsk"`
	AutomaticallyActive          bool         `json:"automaticallyActive"`
	ClearBookOnStart             bool         `json:"clearBookOnStart"`
}
