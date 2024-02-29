package cmd

import (
	"testing"
)

func TestCreateTerm(t *testing.T) {
	if err := RunModel("localhost", 8900); err != nil {
		t.Error(err)
	}
}
