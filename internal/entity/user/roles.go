package user

type Role struct {
	RoleID   int    `json:"role_id"`
	RoleName string `json:"role_name" binding:"required"`
}
