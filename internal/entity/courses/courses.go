package courses

type Course struct {
	Id          int    `json:"-" db:"id"`
	Name        string `json:"name,omitempty" binding:"required"`
	Description string `json:"description,omitempty" binding:"required"`
}
