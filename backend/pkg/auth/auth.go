package auth

type LoginObj struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpObj struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
	Birthday string `json:"birthday"`
}
