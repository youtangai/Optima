package controller

import "testing"

func TestCheckpoint(t *testing.T) {
	containerID := "0825008df297"
	result, err := checkpoint(containerID)
	if err != nil {
		t.Fatalf("failed with err = %v", err)
	}
	t.Logf("result = %s", result)
}

func TestScpCheckpointDir(t *testing.T) {
	targetIP := "192.168.64.12"
	sourceDir := "/tmp/0825008df2970c7cd25c5a7dafca2b0bfe8dabbc5a80ba0cb299ceea1083de79/chk"
	err := scpCheckpointDir(targetIP, sourceDir)
	if err != nil {
		t.Fatalf("failed with err = %v", err)
	}
}
