package anatomy

type JWTUser struct {
	UserID   string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
