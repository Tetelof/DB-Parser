package tables

import (
	"database/sql/driver"
	"fmt"
	"strings"
)

type Extra struct {
	Names []string
}

func (t Extra) Value() (driver.Value, error) {
	return strings.Join(t.Names, ","), nil
}

func (t *Extra) Scan(value interface{}) error {
	switch v := value.(type) {
	case string:
		strs := strings.Split(value.(string), ",")
		if len(strs) == 0 {
			return nil
		}
		t.Names = strs
	case []uint8:
		strs := strings.Split(string(value.([]uint8)), ",")
		if len(strs) == 0 {
			return nil
		}
		t.Names = strs
	default:
		return fmt.Errorf("unsupported scan type for Extra: %T", v)
	}
	return nil
}
