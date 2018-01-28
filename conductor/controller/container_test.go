package controller

import (
	"fmt"
	"testing"
	"time"

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

func TestStartContainer(t *testing.T) {
	//container run
	sourceuuid, _ := createContainer("yotanagai/loop")
	sourcecontainer, _ := db.GetContainerByUUID(sourceuuid)
	startContainer(sourceuuid)
	//create restore container
	start := time.Now()
	targetuuid, _ := createContainer("yotanagai/loop")
	targetcontainer, _ := db.GetContainerByUUID(targetuuid)
	end := time.Now()
	fmt.Printf("create container took time = %v\n", end.Sub(start))
	start = time.Now()
	chkdir, _ := checkpointContainer(sourcecontainer.ContainerID, sourcecontainer.Host)
	end = time.Now()
	fmt.Printf("checkpoint took time = %v\n", end.Sub(start))
	start = time.Now()
	restoreContainer(targetcontainer.ContainerID, chkdir, targetcontainer.Host)
	end = time.Now()
	fmt.Printf("restore took time = %v\n", end.Sub(start))
	start = time.Now()
	deleteContainer(sourceuuid)
	end = time.Now()
	fmt.Printf("delete container took time = %v\n", end.Sub(start))
}
