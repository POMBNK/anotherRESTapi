package todo

// User Сущность пользователь. Поля полностью совпадают со структурой БД.
type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
