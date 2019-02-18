package reportRUsage

import (
	"testing"
	"time"
)

// reportRUsage_test.go tests  ...

func Test(t *testing.T) {

	reportRUsage("ExampleThing", time.Now())
	// t.Errorf("%s failed, %#v\n", "descr", nil)
}

