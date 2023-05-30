package s

type User struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

type AddUserReq struct {
	Name string `json:"name"`
}
