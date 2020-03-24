package entities

// User Represents a user account
type User struct {
	Email    string `json:"email" gorm:"primary_key"`
	Password string `json:"password"`
}
