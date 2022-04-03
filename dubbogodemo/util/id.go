package util

import (
	"github.com/google/uuid"
	"strings"
)

func GenUuid()string{
	u1 := uuid.NewString()
	return strings.Replace(u1, "-", "", -1)
}