package main

type Api struct {
	Query
}

type Query struct {
	Pages
}

type Pages struct {
	Page
}

type Page struct {
	Title string "attr"
	Langlinks
}

type Langlinks struct {
	Ll []Ll
}

type Ll struct {
	Lang string "attr"
	Translation string "chardata"
}
