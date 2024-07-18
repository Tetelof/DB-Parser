package tables

import "github.com/iancoleman/strcase"

type Column struct {
	Field   string
	Type    DataType
	Null    Null
	Key     Key
	Default string
	Extra   Extra
}

func (c *Column) ParseField() string {
	name := strcase.ToCamel(c.Field)
	go_type := MappedDataTypes[c.Type.Name]
	// TODO: tags
	if c.Null {
		return name + tab + pointer_str + go_type
	}
	return name + tab + go_type
}
