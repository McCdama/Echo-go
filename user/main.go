package user

type User struct {
	Password string `json:"password" form:"password"`
	Name     string `json:"name" form:"name"`
}
