package controller

import (
	"testing"

	"github.com/youtangai/Optima/checkpointer/config"
)

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
	config.SetSecretKeyPath("/var/optima/optima_key")
	config.SetControllerIP("192.168.64.12")
	sourceDir := "/var/lib/docker/containers/" + ID + "/checkpoints/chk"
	err := scpCheckpointDir(sourceDir)
	if err != nil {
		t.Fatalf("failed with err = %v", err)
	}
}
