package controller

import "testing"

func TestCheckpoint(t *testing.T) {
	containerID := "thisistest"
	result, err := checkpoint(containerID)
	if err != nil {
		t.Fatalf("failed with err = %v", err)
	}
	t.Logf("result = %s", result)
}

func TestScpCheckpointDir(t *testing.T) {
	targetIP := "192.168.64.12"
	sourceDir := "agagagaga"
	err := scpCheckpointDir(targetIP, sourceDir)
	if err != nil {
		t.Fatalf("failed with err = %v", err)
	}
}
