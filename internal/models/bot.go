package models


type Bot struct {
	ID string `json:"id"`
	Owner string `json:"owner"`
	BotType string `json:"bot_type"`
	Exchange string `json:"exchange"`
	APIKey string `json:"api_key"`
	Balance float64 `json:"balance"`
	PurchaseAmount float64 `json:"purchase_amount"`
	NumberOfPurchases int64 `json:"number_of_purchases"`
	BalanceReplenishment bool `json:"balance_replenishment"`
	AddProfit bool `json:"add_profit"`

} 