package models

type Product struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	СategoryID   string `json:"category_id"`
	Subscription bool   `json:"subscription"`
	Level        int    `json:"level"`
	Icon         string `json:"icon"`
	Url          string `json:"url"`

	ISArchive bool `json:"archive"`
}
