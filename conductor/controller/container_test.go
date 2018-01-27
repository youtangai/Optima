package controller

import "testing"

func TestCreateContainer(t *testing.T) {
	uuid, err := createContainer("yotanagai/loop")
	if err != nil {
		t.Fatalf("cannot create container err = %v", err)
	}
	t.Logf("uuid = %s", uuid)
}
