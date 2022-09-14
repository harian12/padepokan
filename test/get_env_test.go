package test

import (
	"padepokan/helpers"
	"testing"
)

func TestGetEnv(t *testing.T) {
	appName := helpers.GetEnv("APP_NAME", "padepokan")

	if appName != "padepokan" {
		t.Error("app name not match")
	}
}
