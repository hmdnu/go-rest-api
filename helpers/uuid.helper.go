package helper

import (
	"encoding/hex"
	"log"

	"github.com/google/uuid"
)

func GenerateUUID() string {
	newUUID, err := uuid.NewUUID()

	if err != nil {
		log.Fatal(err)
	}

	v, err := newUUID.MarshalBinary()

	if err != nil {
		log.Fatal(err)
	}

	return hex.EncodeToString(v)
}
