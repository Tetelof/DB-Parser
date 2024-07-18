package tables

import (
	"database/sql/driver"
	"fmt"
)

type DataType struct {
	Name   string
	Length int
}

const (
	// Numeric Types
	// Integer types

	INTEGER   = "integer"
	INT       = "int"
	SMALLINT  = "smallint"
	TINYINT   = "tinyint"
	MEDIUMINT = "mediumint"
	BIGINT    = "bigint"

	// Fixed-Point types

	DECIMAL = "decimal"
	NUMERIC = "numeric"

	// Float-Point types

	FLOAT  = "float"
	DOUBLE = "double"

	// String Types
	// Text types

	TEXT       = "text"
	TINYTEXT   = "tinytext"
	MEDIUMTEXT = "mediumtext"
	LONGTEXT   = "longtext"

	// Date and Time Data types

	TIME     = "time"
	DATE     = "date"
	DATETIME = "datetime"
)

var MappedDataTypes = map[string]string{
	TINYINT:   "int8",
	SMALLINT:  "int16",
	MEDIUMINT: "int32",
	INT:       "int32",
	INTEGER:   "int32",
	BIGINT:    "int64",
	DATE:      "time.Time",
	TIME:      "time.Time",
	DATETIME:  "time.Time",
	LONGTEXT:  "string",
}

func (dt DataType) Value() (driver.Value, error) {
	return dt.Name, nil
}

func (dt *DataType) Scan(value interface{}) error {
	switch v := value.(type) {
	case string:
		dt.Name = value.(string)
	case []uint8:
		dt.Name = string(value.([]uint8))
	default:
		return fmt.Errorf("unsupported scan type for DataType: %T", v)
	}
	return nil
}

func (dt *DataType) Import() string {
	return map[string]string{
		DATE:     "time",
		TIME:     "time",
		DATETIME: "time",
	}[dt.Name]
}
