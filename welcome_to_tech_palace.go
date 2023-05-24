package techpalace
import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return fmt.Sprintf("Welcome to the Tech Palace, " + strings.ToUpper(customer))
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	repeatedStars := strings.Repeat("*", numStarsPerLine)
	fancyWelcome := fmt.Sprintf("%s\n%s\n%s", repeatedStars, welcomeMsg, repeatedStars)
	return fancyWelcome 
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	newMsg := strings.ReplaceAll(oldMsg, "*", "")
	return strings.TrimSpace(newMsg)
}



