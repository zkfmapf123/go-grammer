package repo

type book struct {
	product   Product
	createdAt interface{}
}

func NewBook(title string, price int, published interface{}) book {
	b := book{
		Product{title, price},
		published,
	}

	return b
}
