package customer

const (
	Street = iota
	Avenue
	Road
)

type Customer struct {
	ID int64 `json:'idCustomer'`
	Name string `json:'name'`
	Address []address `json:'address'`
	Orders []int64 `json:'order'`
	Favorite string `json:'favorite'`
}

type address struct {
	Type int64 `json:'type'`
	PublicPlace string `json:'publicPlace'`
	Number int64 `json:'number'`
	Complement string `json:'complement'`
}