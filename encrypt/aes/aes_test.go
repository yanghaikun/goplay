package aes

import (
	"testing"
	"fmt"
)

func TestCreateHash(t *testing.T){
	var passwd = "hello"
	key := createHash(passwd)

	fmt.Println("key is %s", key)
}
