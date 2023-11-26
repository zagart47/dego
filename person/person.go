package person

type Person struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Surname    string    `json:"surname"`
	Patronymic string    `json:"patronymic,omitempty"`
	Age        int16     `json:"age"`
	Gender     string    `json:"gender"`
	Country    []Country `json:"country"`
}

type Country struct {
	CountryId   string  `json:"country_id"`
	Probability float32 `json:"probability"`
}
