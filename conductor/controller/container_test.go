package controller

import (
	"fmt"
	"testing"

	"github.com/youtangai/Optima/conductor/db"
)

func TestCreateContainer(t *testing.T) {
	uuid, err := createContainer("yotanagai/loop")
	if err != nil {
		t.Fatalf("cannot create container err = %v", err)
	}
	fmt.Printf("uuid = %s", uuid)
}

func TestAuthKeyStone(t *testing.T) {
	token, err := authKeyStone()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(token)
}

func TestDeleteContainer(t *testing.T) {
	uuid, _ := createContainer("cirros")
	err := deleteContainer(uuid)
	if err != nil {
		t.Fatalf("delete container err = %v", err)
	}
}
func TestCheckpointContainer(t *testing.T) {
	chkdir, err := checkpointContainer("containerid", "hostname")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(chkdir)
}

func TestRestoreContainer(t *testing.T) {
	uuid, _ := createContainer("yotanagai/loop")
	container, _ := db.GetContainerByUUID(uuid)
	fmt.Printf("container info = %+v", container)
	err := restoreContainer(container.ContainerID, "restoredir", container.Host)
	if err != nil {
		t.Fatal(err)
	}
}
