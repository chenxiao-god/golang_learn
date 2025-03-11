package module

type User struct {
	ID       int64  `json:"id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Address  string `json:"address"`
}

func (User) TableName() string {
	return "user"
}
