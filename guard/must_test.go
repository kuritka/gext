package guard

import (
	"os"
	"os/exec"
	"testing"

	"github.com/pkg/errors"
)

func TestCommExec(t *testing.T) {
	err := errors.New("test")
	if os.Getenv("MUST") == "1" {
		Must(err)
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestCommExec")
	cmd.Env = append(os.Environ(), "MUST=1")
	err = cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}

func TestCommExec2(t *testing.T) {
	TestCommExec(nil)
}
