package user

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name,omitempty" binding:"required"`
	Surname  string `json:"surname,omitempty" binding:"required"`
	Username string `json:"username,omitempty" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
	RoleName string `json:"role_name" binding:"required"`
}
