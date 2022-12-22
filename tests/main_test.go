package tests

import (
	"go_code/gintest/bootstrap"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	bootstrap.Initialize("../config/dev_conf.yaml")
	os.Exit(m.Run())
}
