package db

type Struct struct {
	fields []Field
}

func (self *Struct) GetFields() []Field {
	return self.fields;
}

