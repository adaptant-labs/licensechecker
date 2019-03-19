package main

import (
	"github.com/rendon/testcli"
	"testing"
)

func TestLicenseChecker(t *testing.T) {
	testcli.Run("lc", "LICENSE")
	if !testcli.Success() {
		t.Fatalf("Expected to succeed, but failed: %s\n", testcli.Error())
	}

	if !testcli.StdoutContains("MIT") {
		t.Fatalf("Expected to find MIT License\n")
	}
}
