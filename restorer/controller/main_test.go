package controller

import (
	"testing"

	"github.com/youtangai/Optima/restorer/config"
)

const ID = "e950dfd850befa39f089ce09e91cbf51afad06e126ca0c2dc3c239c952616947"

func TestRestore(t *testing.T) {
	containerID := ID
	sourceID := "e950dfd850befa39f089ce09e91cbf51afad06e126ca0c2dc3c239c952616947"
	err := restore(containerID, sourceID)
	if err != nil {
		t.Fatalf("failed with err = %v", err)
	}

}

func TestScpRestoreDir(t *testing.T) {
	config.SetSecretKeyPath("/var/optima/optima_key")
	config.SetControllerIP("192.168.64.12")
	restoreDir := "/var/optima/sandbox/e950dfd850befa39f089ce09e91cbf51afad06e126ca0c2dc3c239c952616947"
	sourceID, err := scpRestoreDir(restoreDir)
	if err != nil {
		t.Fatalf("failed with err = %v", err)
	}
	t.Logf("sourceID = %s", sourceID)
}