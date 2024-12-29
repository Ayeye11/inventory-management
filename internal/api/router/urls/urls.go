package urls

//inventory-management
const (
	//inventory
	Inventory string = "/inventory"
	//products
	Product         string = "/products"
	Product_paramId string = "/products/{id}"
	//categories
	Categories         string = "/categories"
	categories_paramId string = "/categories/{id}"
	//suppliers
	Supplier          string = "/categories"
	Supplier__paramId string = "/categories/{id}"
	//stock-movements
	StockMovement          string = "/stock-movement"
	StockMovement__paramId string = "/stock-movement/{id}"
)

//auth
const (
	Auth_register string = "/auth/register"
	Auth_login    string = "/auth/login"
	Auth_logout   string = "/auth/logout"
	Auth_verify   string = "/auth/verify"
)
