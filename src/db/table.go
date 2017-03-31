package db

type Table struct {
	Fields []Field
	Name string
}

type TableHolder struct {
	table Table
	rows []Row
}
