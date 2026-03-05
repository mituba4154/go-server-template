package example

import goapi "github.com/mituba4154/goapi-sdk"

// Register registers all event handlers and commands for this module.
// Call this from main.go.
func Register() {
	goapi.On("PlayerJoinEvent", onJoin)
}
