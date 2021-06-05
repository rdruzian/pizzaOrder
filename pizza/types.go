package pizza

type Pizza struct {
	ID          int64    `json:"pizzaId"`
	FlavorName  string   `json:"flavorName"`
	Ingredients []string `json:"ingredients"`
	Price       float64  `json:"price"`
}
