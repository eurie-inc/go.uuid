package uuid

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/jbenet/go-base58"
)

var JSONNil = []byte("null")

type NullUUID struct {
	UUID  UUID
	Valid bool
}

// Value implements the driver.Valuer interface.
func (u NullUUID) Value() (driver.Value, error) {
	if !u.Valid {
		return nil, nil
	}
	return u.UUID.Value()
}

func (u *NullUUID) Scan(src interface{}) error {

	if src == nil {
		u.UUID, u.Valid = Nil, false
		return nil
	}

	u.Valid = true
	switch src := src.(type) {
	case []byte:
		if len(src) == 16 {
			return u.UUID.UnmarshalBinary(src)
		}
		return u.UUID.UnmarshalText(src)

	case string:
		return u.UUID.UnmarshalText([]byte(src))
	}

	return fmt.Errorf("uuid: cannot convert %T to UUID", src)
}

func (u NullUUID) MarshalJSON() ([]byte, error) {
	if u.Valid {
		return json.Marshal(u.UUID.Base58String())
	}
	return JSONNil, nil
}

func (u *NullUUID) UnmarshalJSON(text []byte) error {
	if text == nil || reflect.DeepEqual(text, JSONNil) {
		u.UUID = Nil
		u.Valid = false
		return nil
	}
	buf := base58.Decode(strings.Replace(string(text), "\"", "", -1))

	u.Valid = true
	return u.UUID.UnmarshalBinary(buf)
}
