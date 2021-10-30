package model

// User model .
type User struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Name string `json:"name"`
	Age  string `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Type uint `json:"type"`
}
// TableName create name table
func (User) TableName() string { return "users" }
