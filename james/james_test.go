package james

import (
	"os"
	"testing"
	"time"
)

func TestLookFor(t *testing.T) {
	LookFor("error", "found error")

	if len(actions) != 1 {
		// should have set length to 1 ... if not, something else happened
		t.Errorf("Unable to add 'error' to actions")
	}

	if action, exists := actions["error"]; !exists {
		// verify error exists
		t.Errorf("Unable to add 'error' to actions")
	} else {
		// if error exists, verify that the execute command is set properly
		if action.cmd != "found error" {
			t.Errorf("Command for 'error' should be 'found error'")
		}
	}

	// add another
	LookFor("warning", "found warning")
	if len(actions) != 2 {
		t.Errorf("Unable to add 'warning'")
	}
}

func TestExecute(t *testing.T) {
	// clean up from the last round, if need be
	os.Remove("/tmp/error.txt")
	os.Remove("/tmp/warning.txt")

	// add an action
	LookFor("(error)$", "echo :match > /tmp/error.txt")
	LookFor("^(warning)", "echo :match > /tmp/warning.txt")

	// try it out!
	if Execute("This is my really long log line") {
		t.Errorf("Should return false, as nothing matched")
	}

	if !Execute("warning error") {
		// both should have been triggered!
		t.Errorf("warning|error were both found, but nothing was done")
	} else {
		// Give bond a chance to execute properly ...
		<-time.After(time.Second)
		if _, err := os.Stat("/tmp/error.txt"); err != nil {
			t.Errorf("/tmp/error.txt should have been created")
		}
		if _, err := os.Stat("/tmp/warning.txt"); err != nil {
			t.Errorf("/tmp/warning.txt should have been created")
		}
	}

	t.Errorf("...")
}
