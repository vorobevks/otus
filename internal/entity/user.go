package entity

type User struct {
	ID        int    `json:"id"       example:"1"`
	Username  string `json:"username"  example:"username"`
	FirstName string `json:"firstName"  example:"firstName"`
	LastName  string `json:"lastName"  example:"lastName"`
	Email     string `json:"email"  example:"email@example.com"`
	Phone     string `json:"phone"  example:"+79999999999"`
}
