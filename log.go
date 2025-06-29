package u

import (
	"fmt"

	"github.com/fatih/color"
)

func LogError(s string) {
	fmt.Print("  ❌  ")
	red := color.New(color.FgRed).Sprint("Error: ")
	fmt.Printf("%s%s\n", red, s)
}

func LogWarn(s string) {
	fmt.Print("  ⚠️  ")
	yellow := color.New(color.FgYellow).Sprint("Warning: ")
	fmt.Printf("%s%s\n", yellow, s)
}

func LogInfo(s string) {
	fmt.Print("  ℹ️  ")
	blue := color.New(color.FgBlue).Sprint("Info: ")
	fmt.Printf("%s%s\n", blue, s)
}

func LogDebug(s string) {
	fmt.Print("  🔍  ")
	mgt := color.New(color.FgMagenta).Sprint("Debug: ")
	fmt.Printf("%s%s\n", mgt, s)
}

func LogSuccess(s string) {
	fmt.Print("  ✅  ")
	color.Green("%s", s)
}
