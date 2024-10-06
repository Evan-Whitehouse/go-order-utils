package datastructures

type tags struct {
	ID          string `json:"id"`
	Label       string `json:"label"`
	Slug        string `json:"slug"`
	ForceShow   bool   `json:"forceShow"`
	PublishedAt string `json:"published_at"`
	UpdatedBy   int    `json:"updatedBy"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	Sync        bool   `json:"_sync"`
}

type Event struct {
	ID                  string   `json:"id"`
	Ticker              string   `json:"ticker"`
	Slug                string   `json:"slug"`
	Title               string   `json:"title"`
	Description         string   `json:"description"`
	ResolutionSource    string   `json:"resolutionSource"`
	StartDate           string   `json:"startDate"`
	CreationDate        string   `json:"creationDate"`
	EndDate             string   `json:"endDate"`
	Image               string   `json:"image"`
	Icon                string   `json:"icon"`
	Active              bool     `json:"active"`
	Closed              bool     `json:"closed"`
	Archived            bool     `json:"archived"`
	New                 bool     `json:"new"`
	Featured            bool     `json:"featured"`
	Restricted          bool     `json:"restricted"`
	Liquidity           float64  `json:"liquidity"`
	Volume              float64  `json:"volume"`
	OpenInterest        float64  `json:"openInterest"`
	SortBy              string   `json:"sortBy"`
	PublishedAt         string   `json:"published_at"`
	UpdatedBy           string   `json:"updatedBy"`
	CreatedAt           string   `json:"createdAt"`
	UpdatedAt           string   `json:"updatedAt"`
	CommentsEnabled     bool     `json:"commentsEnabled"`
	Competitive         float64  `json:"competitive"`
	Volume24hr          float64  `json:"volume24hr"`
	FeaturedImage       string   `json:"featuredImage"`
	EnableOrderBook     bool     `json:"enableOrderBook"`
	LiquidityClob       float64  `json:"liquidityClob"`
	Sync                bool     `json:"_sync"`
	NegRisk             bool     `json:"negRisk"`
	NegRiskMarketID     string   `json:"negRiskMarketID"`
	NegRiskFeeBips      int      `json:"negRiskFeeBips"`
	CommentCount        int      `json:"commentCount"`
	Markets             []Market `json:"markets"`
	Tags                []tags   `json:"tags"`
	Cyom                bool     `json:"cyom"`
	ShowAllOutcomes     bool     `json:"showAllOutcomes"`
	ShowMarketImages    bool     `json:"showMarketImages"`
	EnableNegRisk       bool     `json:"enableNegRisk"`
	AutomaticallyActive bool     `json:"automaticallyActive"`
}
