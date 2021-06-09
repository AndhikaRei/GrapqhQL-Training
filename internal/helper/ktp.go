package helper

import (
	"encoding/base64"
	"errors"
	"strconv"
)

// EncodeID is a function that will turn ID (int) to base64 string representation
// This function will return string representation of base65 encoded id
func EncodeID(ID int) string {

	stringID := strconv.Itoa(ID)
	return base64.StdEncoding.EncodeToString([]byte(stringID))
}

// DecodeID is a function that will turn base64 string representation of ID to integer
// This function will return decoded id in integer
func DecodeID(ID string) (int, error) {

	bytes, err := base64.StdEncoding.DecodeString(ID)
	if err != nil {
		return -1, err
	}

	vals := string(bytes)
	id, err := strconv.Atoi(vals)
	if err != nil {
		return -1, errors.New("invalid_cursor")
	}

	return id, nil
}
