package models

type Image struct {
	Title        string 	`json:"title"`
	Description  string 	`json:"description"`
	FileLocation string		`json:"fileLocation"`
	Price		 float64 	`json:"price"`
	Public		 bool		`json:"public"`
}

