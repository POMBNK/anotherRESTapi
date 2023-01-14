package todo

// TodoList  Todo списки
type TodoList struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// UserList сущность для связи пользователя и списка задач M to M
type UserList struct {
	ID     int
	UserId int
	ListId int
}

// TodoItem сущность элементов из списка задач
type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

// ListItem сущность для связи списков и задач M to M
type ListItem struct {
	Id     int
	ListId int
	ItemId int
}
