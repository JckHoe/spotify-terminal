package internal

type Page struct {
	Name string

	IsInit   bool
	Items    []Item
	Cursor   int
	Selected Item
}

type Item struct {
	DisplayName string
	ID          string
	Active      bool
}
