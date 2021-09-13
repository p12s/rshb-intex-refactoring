package model

type Book struct {
	Id       int     `json:"id" db:"id"`
	Title    string  `json:"title" db:"title"`
	AuthorId int     `json:"author_id" db:"author_id"`
	Cost     float64 `json:"cost" db:"cost"`
}

type Author struct {
	Id         int    `json:"id" db:"id"`
	FirstName  string `json:"first_name" db:"first_name"`
	SecondName string `json:"second_name" db:"second_name"`
}
