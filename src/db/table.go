package db

type Table struct {
	Fields []Field
	Name string
}

type TableHolder struct {
	table Table
	rows []Row
}

var tableHolder map[string]TableHolder = make(map[string]TableHolder)
