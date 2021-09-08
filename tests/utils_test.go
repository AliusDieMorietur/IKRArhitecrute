package tests

import (
	"os"
	"testing"

	Utils "server/lib/utils"
)

func TestGetPort(t *testing.T) {
	port := Utils.GetPort()
	envPort := os.Getenv("PORT")
	if port != envPort && port != "8765" {
		t.Fatalf(`error`)
	}
}
