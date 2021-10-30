package model

type Product truct {
	ID 	uint
	Name	string
	Price	float32
	Urls	string
	Images	string


	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

// TableName create name table
func (Product) TableName() string {return "products"}; 
