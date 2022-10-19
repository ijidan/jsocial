package uuid

import (
	"github.com/gofrs/uuid"
)

func GetUUId() (string,error) {
	u2, err := uuid.NewV4()
	if err != nil {
		return "",err
	}
	return u2.String(),nil
}
