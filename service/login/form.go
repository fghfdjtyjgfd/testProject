package login

type LoginForm struct {
	Email    string `json:"email"`
	ID       int    `json:"id"`
	Password string `json:"password"`
}
