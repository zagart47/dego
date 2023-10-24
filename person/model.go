package person

type Person struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Surname    string    `json:"surname"`
	Patronymic string    `json:"patronymic"`
	Age        int       `json:"age"`
	Gender     string    `json:"gender"`
	Country    []Country `json:"country"`
}

type Country struct {
	PersonId    int     `json:"person_id"`
	CountryId   string  `json:"country_id"`
	Probability float32 `json:"probability"`
}
