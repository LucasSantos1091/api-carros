package models

type Car struct {
	ID    int    `json:"id"`
	Make  string `json:"make"`
	Model string `json:"model"`
	Year  int    `json:"year"`
}

var Cars = []Car{
	{ID: 1, Make: "Toyota", Model: "Corolla", Year: 2020},
	{ID: 2, Make: "Honda", Model: "Civic", Year: 2019},
}
