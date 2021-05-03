package domain

type Course struct {
	Name        string `json:"name" db:"name"`
	Professor   string `json:"professor" db:"professor"`
	Description string `json:"description" db:"description"`
}
