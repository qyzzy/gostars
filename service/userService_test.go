package service

import (
	"fmt"
	"math/rand"
	"testing"
)

func randomName() string {
	length := rand.Int() % 10
	name := make([]byte, length)
	for i := 0; i < length; i++ {
		bt := byte(rand.Int()%26 + 'a')
		if rand.Int()%2 == 0 {
			bt = byte(rand.Int()%26 + 'A')
		}
		name[i] = bt
	}
	return string(name)
}

func TestUserService_Register(t *testing.T) {
	fmt.Println(randomName())
}
