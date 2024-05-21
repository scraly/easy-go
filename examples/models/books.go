package models

type Size struct {
	Width  int
	Height int
}

type Books struct {
	Name string
	Size Size
}
