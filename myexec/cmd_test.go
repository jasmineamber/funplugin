package myexec

import "testing"

func TestRunShell(t *testing.T) {
	exitCode, err := RunShell("echo hello world")
	if err != nil {
		t.Fatal(err)
	}
	if exitCode != 0 {
		t.Fatalf("expected exit code 0, got %d", exitCode)
	}

	exitCode, err = RunShell("for i in {1..5}; do echo $i; sleep 1; done")
	if err != nil {
		t.Fatal(err)
	}
	if exitCode != 0 {
		t.Fatalf("expected exit code 0, got %d", exitCode)
	}
}