package u

import (
	"testing"
)

func TestLog(t *testing.T) {
	LogError("error occurred")
	LogWarn("i double dare you!")
	LogInfo("did you know...")
	LogSuccess("Yay!")
	LogDebug("Some info...")
}
