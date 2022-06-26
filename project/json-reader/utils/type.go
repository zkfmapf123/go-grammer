package utils

type Language struct {
	Greet   string `json:"greet"`
	Bye     string `json:"bye"`
	Cheerup string `json:"cheerup"`
}

type Book struct {
	Ko Language `json:"ko"`
	En Language `json:"en"`
	Jp Language `json:"jp"`
}
