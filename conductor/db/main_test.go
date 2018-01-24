package db

import (
	"testing"

	"github.com/youtangai/Optima/conductor/model"
)

func TestRegistLoadIndicator(t *testing.T) {
	entity := new(model.LoadIndicatorJson)
	entity.HostName = "testhost"
	entity.HostIP = "10.0.0.1"
	entity.LoadIndicator = 0.1546
	err := RegistLoadIndicator(*entity)
	if err != nil {
		t.Fatalf("test failed %v", err)
	}
	entity.LoadIndicator = 0.234
	err = RegistLoadIndicator(*entity)
	if err != nil {
		t.Fatalf("test failed %v", err)
	}
}
