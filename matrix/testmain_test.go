package matrix

import (
	"os"
	"testing"

	"github.com/365admin/sharepoint-webparts/util"
	//"github.com/koksmat-com/koksmat/config"
)

func TestMain(m *testing.M) {
	util.Setup("../test.env")
	code := m.Run()

	os.Exit(code)
}
