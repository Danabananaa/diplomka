package model

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserImage struct {
	UserID    int64
	ImageName string
}

type ImageData struct {
	Data string `json:"data"`
}
