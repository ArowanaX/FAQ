package entity

type Faq struct {
	Category Category
	Question string
	Answer   string
}

type Category struct {
	Title string
}
