/*
------------------------------------------------------------------------------------------------------------------------
####### datamap ####### (c) 2020-2021 mls-361 ###################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package datamap

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/mls-361/failure"
)

var (
	// ErrBadType AFAIRE.
	ErrBadType = errors.New("bad type")
	// ErrNotFound AFAIRE.
	ErrNotFound = errors.New("not found")
)

// DataMap AFAIRE.
type DataMap map[string]interface{}

// New AFAIRE.
func New() DataMap {
	return make(DataMap)
}

func (dm DataMap) retrieve(dPath string, keys []string) (interface{}, error) {
	dPath += "/" + keys[0]

	value, ok := dm[keys[0]]
	if !ok {
		return nil, failure.New(ErrNotFound).
			Set("path", dPath).
			Msg("this data path does not exist") ///////////////////////////////////////////////////////////////////////

	}

	if len(keys) == 1 {
		return value, nil
	}

	vdm, ok := value.(DataMap)
	if !ok {
		return nil, failure.New(ErrBadType).
			Set("path", dPath).
			Msg("this data path does not refer to a data map") /////////////////////////////////////////////////////////
	}

	return vdm.retrieve(dPath, keys[1:])
}

// Retrieve AFAIRE.
func (dm DataMap) Retrieve(keys ...string) (interface{}, error) {
	if len(keys) == 0 {
		return dm, nil
	}

	return dm.retrieve("", keys)
}

// RetrieveWD AFAIRE.
func (dm DataMap) RetrieveWD(d interface{}, keys ...string) (interface{}, error) {
	value, err := dm.Retrieve(keys...)
	if err == nil {
		return value, nil
	}

	if errors.Is(err, ErrNotFound) {
		return d, nil
	}

	return nil, err
}

func errBadType(t string, keys ...string) error {
	return failure.New(ErrBadType).
		Set("path", strings.Join(keys, "/")).
		Msg("this data path does not refer to a " + t) /////////////////////////////////////////////////////////////////
}

// Bool AFAIRE.
func (dm DataMap) Bool(keys ...string) (bool, error) {
	value, err := dm.Retrieve(keys...)
	if err != nil {
		return false, err
	}

	switch v := value.(type) {
	case bool:
		return v, nil
	case string:
		return strconv.ParseBool(v)
	default:
		return false, errBadType("boolean", keys...)
	}
}

// BoolWD AFAIRE.
func (dm DataMap) BoolWD(d bool, keys ...string) (bool, error) {
	value, err := dm.Bool(keys...)
	if err == nil {
		return value, nil
	}

	if errors.Is(err, ErrNotFound) {
		return d, nil
	}

	return false, err
}

// Int AFAIRE.
func (dm DataMap) Int(keys ...string) (int, error) {
	value, err := dm.Retrieve(keys...)
	if err != nil {
		return 0, err
	}

	switch v := value.(type) {
	case int:
		return v, nil
	case string:
		return strconv.Atoi(v)
	default:
		return 0, errBadType("integer", keys...)
	}
}

// IntWD AFAIRE.
func (dm DataMap) IntWD(d int, keys ...string) (int, error) {
	value, err := dm.Int(keys...)
	if err == nil {
		return value, nil
	}

	if errors.Is(err, ErrNotFound) {
		return d, nil
	}

	return 0, err
}

// String AFAIRE.
func (dm DataMap) String(keys ...string) (string, error) {
	value, err := dm.Retrieve(keys...)
	if err != nil {
		return "", err
	}

	if v, ok := value.(string); ok {
		return v, nil
	}

	return "", errBadType("string", keys...)
}

// StringWD AFAIRE.
func (dm DataMap) StringWD(d string, keys ...string) (string, error) {
	value, err := dm.String(keys...)
	if err == nil {
		return value, nil
	}

	if errors.Is(err, ErrNotFound) {
		return d, nil
	}

	return "", err
}

// Duration AFAIRE.
func (dm DataMap) Duration(keys ...string) (time.Duration, error) {
	value, err := dm.Retrieve(keys...)
	if err != nil {
		return 0, err
	}

	s, ok := value.(string)
	if ok {
		v, err := time.ParseDuration(s)
		if err == nil {
			return v, nil
		}
	}

	return 0, errBadType("duration", keys...)
}

// DurationWD AFAIRE.
func (dm DataMap) DurationWD(d time.Duration, keys ...string) (time.Duration, error) {
	value, err := dm.Duration(keys...)
	if err == nil {
		return value, nil
	}

	if errors.Is(err, ErrNotFound) {
		return d, nil
	}

	return 0, err
}

/*
######################################################################################################## @(°_°)@ #######
*/
