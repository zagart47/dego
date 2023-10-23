package model

type Person struct {
	Name         string    `json:"name"`
	Surname      string    `json:"surname"`
	Patronymic   string    `json:"patronymic"`
	Age          int       `json:"age"`
	Gender       string    `json:"gender"`
	CountryField []Country `json:"country"`
}

type Country struct {
	CountryId   string  `json:"country_id"`
	Probability float32 `json:"probability"`
}
