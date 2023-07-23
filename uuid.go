package uuid

import (
	"database/sql/driver"

	"github.com/google/uuid"
)

var Nil = UUID{uuid.Nil}

type UUID struct {
	uuid.UUID
}

/*
NewSHA1 returns a new SHA1 (Version 5) UUID based on the supplied name space and data
*/
func NewSha1(namespace UUID, data []byte) UUID {
	return UUID{UUID: uuid.NewSHA1(namespace.UUID, data)}
}

/*
New creates a new random UUID or panics.
*/
func New() UUID {
	return UUID{uuid.New()}
}

/*
StringOrBlank String returns the string form of uuid but retrurn blank string in uuid == uuid.Nil
*/
func (uuid UUID) StringOrBlank() string {
	if uuid == Nil {
		return ""
	}
	return uuid.String()
}

/*
Value can implements sql.Scanner so UUIDs can be read from databases transparently. return nil if uuid == uuid.Nil
*/
func (uuid UUID) Value() (driver.Value, error) {
	if uuid == Nil {
		return nil, nil
	}
	return uuid.String(), nil
}

/*
Less compare two uuid for order
*/
func (uuid UUID) Less(than UUID) bool {
	return uuid.UUID[0] < than.UUID[0] &&
		uuid.UUID[1] < than.UUID[1] &&
		uuid.UUID[2] < than.UUID[2] &&
		uuid.UUID[3] < than.UUID[3] &&
		uuid.UUID[4] < than.UUID[4] &&
		uuid.UUID[5] < than.UUID[5] &&
		uuid.UUID[6] < than.UUID[6] &&
		uuid.UUID[7] < than.UUID[7] &&
		uuid.UUID[8] < than.UUID[8] &&
		uuid.UUID[9] < than.UUID[9] &&
		uuid.UUID[10] < than.UUID[10] &&
		uuid.UUID[11] < than.UUID[11] &&
		uuid.UUID[12] < than.UUID[12] &&
		uuid.UUID[13] < than.UUID[13] &&
		uuid.UUID[14] < than.UUID[14] &&
		uuid.UUID[15] < than.UUID[15]
}

/*
Parse decodes string into a UUID or returns an error
*/
func Parse(v string) (UUID, error) {
	_uuid, err := uuid.Parse(v)
	return UUID{_uuid}, err
}

/*
Parse decodes string into a UUID or uuid.Nil
*/
func ParseOrNil(s string) UUID {
	if u, err := Parse(s); err == nil {
		return u
	} else {
		return Nil
	}
}

/*
Strings encode UUIDs to strings
*/
func Strings(args []UUID) []string {
	ret := make([]string, len(args))
	for i := range args {
		ret[i] = args[i].String()
	}
	return ret
}

/*
Parse decodes strings into a UUIDs
*/
func Parses(args []string) []UUID {
	ret := []UUID{}
	for i := range args {
		tmp, err := Parse(args[i])
		if err != nil {
			continue
		}
		ret = append(ret, tmp)
	}
	return ret
}

/*
Containe return true if uuid is found in given slice
*/
func Containe(arr []UUID, id UUID) bool {
	for i := range arr {
		if arr[i] == id {
			return true
		}
	}
	return false
}

/*
FilterUnique remove uuid diplicate
*/
func FilterUnique(arr []UUID) []UUID {
	ret := []UUID{}
	for i := range arr {
		if !Containe(ret, arr[i]) {
			ret = append(ret, arr[i])
		}
	}
	return ret
}

/*
Remove return array without uuid passed
*/
func Remove(arr []UUID, id UUID) []UUID {
	ret := []UUID{}
	for i := range arr {
		if arr[i] == id {
			continue
		}
		ret = append(ret, arr[i])
	}
	return ret
}

/*
Parse decodes string into a UUID or return uuid.Nil
*/
func FromBytes(b []byte) (UUID, error) {
	_uuid, err := uuid.FromBytes(b)
	if err != nil {
		return Nil, err
	}
	return UUID{UUID: _uuid}, nil
}
