package model

// UserRegister ...
type UserRegister struct {
	ID              uint64 `json:"_id"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
	UUID            string `json:"uuid"`
}

//UserLogin ...
type UserLogin struct {
	ID       uint64 `json:"_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	UUID     string `json:"uuid"`
}

//UserLogout ...
type UserLogout struct {
	Email string `json:"email"`
}

//UserVerify ...
type UserVerify struct {
	VerificationCode string `json:"verification_code"`
}
