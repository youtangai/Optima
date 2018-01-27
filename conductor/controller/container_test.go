package controller

import (
	"fmt"
	"testing"
)

func TestCreateContainer(t *testing.T) {
	uuid, err := createContainer("yotanagai/loop")
	if err != nil {
		t.Fatalf("cannot create container err = %v", err)
	}
	t.Logf("uuid = %s", uuid)
}

func TestAuthKeyStone(t *testing.T) {
	token, err := authKeyStone()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(token)
}
