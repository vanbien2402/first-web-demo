package identifier

import "github.com/google/uuid"

//Generate generate a new id
func Generate() string {
	return uuid.New().String()
}
