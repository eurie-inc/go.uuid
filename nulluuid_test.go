package uuid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNullUUID_Value(t *testing.T) {
	var u NullUUID
	var err error

	u.UUID, err = FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	if err != nil {
		t.Errorf("Error parsing UUID from string: %s", err)
	}
	u.Valid = true
	val, err := u.Value()
	assert.Nil(t, err, "err shoud be nil")
	assert.Equal(t, u.UUID.String(), val, "In the case of Valid == true, they should be equal.")

	u.Valid = false
	val, err = u.Value()
	assert.Nil(t, err, "err should be nil")
	assert.Nil(t, val, "In the case of Valid == false, val shoud be nil.")
}

func TestNullUUID_Scan(t *testing.T) {

	uuidString := "6ba7b810-9dad-11d1-80b4-00c04fd430c8"
	e1 := UUID{0x6b, 0xa7, 0xb8, 0x10, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8}

	var u NullUUID

	err := u.Scan(uuidString)
	assert.Nil(t, err, "err should be nil.")
	assert.Equal(t, e1, u.UUID)
	assert.True(t, u.Valid)

	uuidBytes := []byte{0x6b, 0xa7, 0xb8, 0x10, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8}

	err = u.Scan(uuidBytes)
	assert.Nil(t, err, "err should be nil.")
	assert.Equal(t, e1, u.UUID)
	assert.True(t, u.Valid)

	err = u.Scan(nil)
	assert.Nil(t, err, "err should be nil.")
	assert.Equal(t, Nil, u.UUID)
	assert.False(t, u.Valid)
}

func TestNullUUID_MarshalJSON(t *testing.T) {
	e1 := []byte{0x22, 0x45, 0x4a, 0x33, 0x34, 0x6b, 0x43, 0x56, 0x78, 0x78, 0x46, 0x39, 0x6a, 0x48, 0x4d, 0x4b, 0x44, 0x34, 0x45, 0x67, 0x72, 0x41, 0x4b, 0x22}
	uuidBase58 := "EJ34kCVxxF9jHMKD4EgrAK"

	var u NullUUID
	u.UUID, _ = FromBase58String(uuidBase58)
	u.Valid = true

	a1, err := u.MarshalJSON()
	assert.Nil(t, err, "err should be nil.")
	assert.Equal(t, e1, a1, "They should be equal.")

	e2 := []byte("null")
	u.Valid = false

	a2, err := u.MarshalJSON()
	assert.Nil(t, err, "err should be nil.")
	assert.Equal(t, e2, a2, "In the case of valid == false, MarshalJSON() should return 'null'.")
}

func TestNullUUID_UnmarshalJSON(t *testing.T) {
	e1 := UUID{0x6b, 0xa7, 0xb8, 0x10, 0x9d, 0xad, 0x11, 0xd1, 0x80, 0xb4, 0x00, 0xc0, 0x4f, 0xd4, 0x30, 0xc8}

	var u NullUUID
	err := u.UnmarshalJSON([]byte("EJ34kCVxxF9jHMKD4EgrAK"))
	assert.Nil(t, err, "err should be nil.")
	assert.Equal(t, e1, u.UUID, "They should be equal.")
	assert.True(t, u.Valid, "Valid shoud be true")

	err = u.UnmarshalJSON(nil)
	assert.Nil(t, err, "err should be nil.")
	assert.Equal(t, Nil, u.UUID, "In the case of text is nil, UUID should be Nil.")
	assert.False(t, u.Valid, "In the case of text is nil, Valid should be false.")

}
