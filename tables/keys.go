package tables

import (
	"database/sql/driver"
	"fmt"
)

type Key struct {
	Name     string
	Relation string
}

func (t Key) Value() (driver.Value, error) {
	return t.Name, nil
}

func (t *Key) Scan(value interface{}) error {
	switch v := value.(type) {
	case string:
		t.Name = value.(string)
	case []uint8:
		t.Name = string(value.([]uint8))
	default:
		return fmt.Errorf("unsupported scan type for Key: %T", v)
	}
	return nil
}
