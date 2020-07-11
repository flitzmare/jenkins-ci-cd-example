package jenkins_ci_cd_example

import (
	"testing"
)

func Test_main(t *testing.T) {
	if Print() != "Hello World" {
		t.Error("Print() doesn't return Hello World")
	}
}
