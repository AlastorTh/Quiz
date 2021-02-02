package models

// Field is a row in the csv
type Field struct {
	Question string
	Answer   int
}

// Container for records ..
type Container struct {
	Field []*Field
}
