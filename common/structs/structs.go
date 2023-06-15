package s

import "time"

type User struct {
	ID             string    `json:"id"`
	FullName       string    `json:"full_name"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	Avatar         string    `json:"avatar"`
	HashedPassword string    `json:"hashed_password"`
	Balance        []Balance `json:"balance"`
	Role           string    `json:"role"`
	CreatedAt      time.Time `json:"created_at"`
	CreatedBy      string    `json:"created_by"` //this is ID of the user who created this user
	ModifiedAt     time.Time `json:"modified_at"`
	ModifiedBy     string    `json:"modified_by"` //this is ID of the user who modified this user at any time

}

type Balance struct {
	UsedCurrency string `json:"currency"`
	Amount       string `json:"amount"`
}

type AddUserReq struct {
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Password string `json:"hashed_password"`
	Role     string `json:"role"`
}

type AddUserRes struct {
	ID       string `json:"id"`
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Role     string `json:"role"`
}
