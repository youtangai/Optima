package controller

import "testing"

const ID = "9277231bbe1634591433a66b4c907b32899a3b658641326494b83309bd784f2b"

func TestCheckpoint(t *testing.T) {
	containerID := ID
	result, err := checkpoint(containerID)
	if err != nil {
		t.Fatalf("failed with err = %v", err)
	}
	t.Logf("result = %s", result)
}

func TestScpCheckpointDir(t *testing.T) {
	targetIP := "192.168.64.12"
	sourceDir := "/var/lib/docker/containers/" + ID + "/checkpoints/chk"
	err := scpCheckpointDir(targetIP, sourceDir)
	if err != nil {
		t.Fatalf("failed with err = %v", err)
	}
}
