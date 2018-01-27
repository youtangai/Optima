package db

import (
	"testing"

	"github.com/youtangai/Optima/conductor/model"
)

const (
	TEST_HOST_NAME  = "zun1"
	TEST_HOST_IP    = "10.0.0.2"
	TEST_UUID       = "cb94b9de-f6d2-4135-befe-10642d42559b"
	TEST_CHKDIR     = "/var/optima/zun1/hoge"
	TEST_IMAGE_NAME = "yotanagai/loop"
)

func TestRegistLoadIndicator(t *testing.T) {
	t.Log("regist load indicator test")
	entity := new(model.LoadIndicatorJson)
	entity.HostName = TEST_HOST_NAME
	entity.HostIP = TEST_HOST_IP
	entity.LoadIndicator = 0.234
	err := RegistLoadIndicator(*entity)
	if err != nil {
		t.Fatalf("cannot regist load indicator err = %v", err)
	}
}

func TestGetIPAddrByHostName(t *testing.T) {
	t.Log("get ip addr by hostname test")
	ipAddr, err := GetIPAddrByHostName(TEST_HOST_NAME)
	if err != nil {
		t.Fatalf("cant get ip addr by hostname err = %v", err)
	}
	if ipAddr != TEST_HOST_IP {
		t.Fatalf("gotten host ip not match test host ip = %s", ipAddr)
	}
	t.Log(ipAddr)
}

func TestGetContainersByHostName(t *testing.T) {
	t.Log("get containers by hostname")
	containers, err := GetContainersByHostName(TEST_HOST_NAME)
	if err != nil {
		t.Fatalf("connot get containers err = %v", err)
	}
	t.Logf("get containers = %v", containers)
}

func TestGetContainerByUUID(t *testing.T) {
	t.Log("get container by uuid")
	container, err := GetContainerByUUID(TEST_UUID)
	if err != nil {
		t.Fatalf("cannot get container by uuid err = %v", err)
	}
	t.Logf("get container = %v", container)
}

func TestDeleteLoadIndicator(t *testing.T) {
	t.Log("test delete load indicator")
	err := DeleteLoadIndicator(TEST_HOST_NAME)
	if err != nil {
		t.Logf("cannot delete load indicator err = %v", err)
	}
	t.Logf("delete loadindicator complete")
}

func TestRegistCheckPointDir(t *testing.T) {
	t.Log("test register checkpointdir")
	err := RegistCheckPointDir(TEST_CHKDIR, TEST_IMAGE_NAME)
	if err != nil {
		t.Fatalf("cannot regist checkpointdir err = %v", err)
	}
	t.Log("regist chkdir")
}

func TestGetCheckPointDirs(t *testing.T) {
	t.Log("test getchkpointdir")
	checkpoints, err := GetCheckPointDirs()
	if err != nil {
		t.Fatalf("cannnot get chkdir err = %v", err)
	}
	t.Logf("get checkpoihnts = %v", checkpoints)
}

func TestDeleteCheckPointDir(t *testing.T) {
	t.Log("test delete chkdir")
	checkpointsAddr, _ := GetCheckPointDirs()
	checkpoints := *checkpointsAddr
	checkpoint := checkpoints[0]
	err := DeleteCheckPointDir(checkpoint)
	if err != nil {
		t.Fatalf("cannot delete chkdir err = %v", err)
	}
	t.Log("delete chkdir complete")
}

func TestGetHostOrderByLoadIndicator(t *testing.T) {
	hosts, err := GetHostOrderByLoadIndicator()
	if err != nil {
		t.Fatalf("cannot get host orderby load indicator = %v", err)
	}
	t.Logf("get hosts = %v", hosts)
}
