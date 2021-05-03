package domain

type Student struct {
	Id    int    `json:"id"`
	Name  string `json:"string"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type CreateStudent struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
