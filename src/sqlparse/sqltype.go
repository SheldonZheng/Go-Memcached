package sqlparse

type SqlType int

const (
	INSERT = iota
	DELETE
	UPDATE
	CREATE
	TRUNCATE
)
