package model
// Brand model .
type Brand struct{
	ID		uint		'gorm:"primary", json:"id"'	
	Name 		string		'json:"name"'
	Slogan		string		'json:"'
	Description	string		'json:"description"'
	Logo		string		'json:"logo"'
	Banner		string		'json:"banner"'
	CreatedAt 	time.Time 	`json:"created_at"`
	UpdatedAt 	time.Time 	`json:"updated_at"`
	DeletedAt 	time.Time 	`json:"deleted_at"`
}	     
// TableName create name table
func (Brand) TableName() string { return "brands" }

