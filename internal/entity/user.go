package entity

type User struct {
	ID        int    `json:"id"       example:"1"`
	Username  string `json:"username"  example:"username"`
	Name      string `json:"name"  example:"name"`
	Password  string `json:"password"     example:"5f4dcc3b5aa765d61d8327deb882cf99"`
	CreatedAt string `json:"created_at"  example:"2025-04-02 11:31:01.314851"`
}
