package main

type Api struct {
	Query
	Any Continue //This field should be named Query-continue
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
	Lang        string "attr"
	Translation string "chardata"
}

type Continue struct {
	Langlinks LanglinksC
}

type LanglinksC struct {
	Llcontinue string "attr"
}
