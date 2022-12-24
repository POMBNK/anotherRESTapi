package todo

// User... Сущность пользователь. Поля полностью совпадают со структоурой БД.
type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
