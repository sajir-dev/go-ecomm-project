package domain

// Product struct
type Product struct {
	ProductID   string  `json:"id"`
	ProductName string  `json:"product_name"`
	Category    string  `json:"category"`
	Price       float32 `json:"price"`
	Discount    float32 `json:"discout"`
	Description string  `json:"description"`
	Rating      float32 `json:"rating"`
	// Stock       int     `json:"stock"`
	// Images      string  `json:"images"`
	// Size        string  `json:"size"`
}

// insert into products (product_name, category, price, discount, description, rating) values ('test name','test cat', 34.5, 12, 'test description', 4.5);
