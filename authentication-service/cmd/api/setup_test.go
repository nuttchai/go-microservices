// must be named as setup_test.go to run the test before starting the app
package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
