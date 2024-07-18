package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Date time.Time

// UnmarshalCSV parses the CSV string and assigns it to the Date
func (d *Date) UnmarshalCSV(csv string) (err error) {
	t, err := time.Parse("1/2/2006", csv)
	*d = Date(t)
	return err
}

// Value implements the driver Valuer interface for Date
func (d Date) Value() (driver.Value, error) {
	t := time.Time(d)
	return t.Format("2006-01-02"), nil
}

// Scan implements the sql Scanner interface for Date
func (d *Date) Scan(value interface{}) error {
	if value == nil {
		*d = Date(time.Time{})
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		*d = Date(v)
	case string:
		t, err := time.Parse("2006-01-02", v)
		if err != nil {
			return err
		}
		*d = Date(t)
	default:
		return fmt.Errorf("unsupported scan type for Date: %T", v)
	}
	return nil
}
