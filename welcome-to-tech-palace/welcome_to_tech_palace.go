package techpalace

import (
	"strings"
)

var star string = "*"

func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	horizontalBorder := strings.Repeat(star, numStarsPerLine)
	var builder strings.Builder

	builder.WriteString(horizontalBorder + "\n")
	builder.WriteString(welcomeMsg + "\n")
	builder.WriteString(horizontalBorder)

	return builder.String()
}

func CleanupMessage(oldMsg string) string {
	cleanupString := strings.ReplaceAll(oldMsg, "\n", "")
	cleanupString = strings.ReplaceAll(cleanupString, star, "")

	return strings.TrimSpace(cleanupString)
}
