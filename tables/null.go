package tables

import (
	"database/sql/driver"
	"fmt"
)

type Null bool

func (t Null) Value() (driver.Value, error) {
	if t {
		return "YES", nil
	}
	return "NO", nil
}

func (t *Null) Scan(value interface{}) error {
	switch v := value.(type) {
	case string:
		*t = value.(string) == "YES"
	case []uint8:
		*t = string(value.([]uint8))  == "YES"
	default:
		return fmt.Errorf("unsupported scan type for Null: %T", v)
	}
	return nil
}
