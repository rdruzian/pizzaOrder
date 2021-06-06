package customer

const (
	Street = iota
	Avenue
	Road
)

type Customer struct {
	ID int64 `json:"idCustomer"`
	Name string `json:"name"`
	Address Address `json:"address"`
	Orders []int64 `json:"order"`
	Favorite string `json:"favorite"`
	Login string `json:"login"`
	Password string `json:"pass"`
}

type Address struct {
	Type int64 `json:"type"`
	PublicPlace string `json:"publicPlace"`
	Number int64 `json:"number"`
	Complement string `json:"complement"`
}