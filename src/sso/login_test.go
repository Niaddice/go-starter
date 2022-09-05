package sso

import (
	"fmt"
	"go-starter/src/utils"
	"testing"
)

func TestLogin(t *testing.T) {
	salt := utils.RandSalt(4)
	password := utils.GenerateHashPassword("123456", salt)
	fmt.Printf("salt: %s, password: %s\n", salt, password)
}
