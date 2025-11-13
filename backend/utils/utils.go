package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"
)

func GenerateRandomCode(n uint8) string {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		fmt.Println("Error when Generating Code:", err)
		return ""
	}

	return strings.ToUpper(hex.EncodeToString(bytes))
}
