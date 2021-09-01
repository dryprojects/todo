package utils

import "fmt"

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

func TextRed(text string) string {
	return fmt.Sprintf("%s%s%s", colorRed, text, colorReset)
}

func TextGreen(text string) string {
	return fmt.Sprintf("%s%s%s", colorGreen, text, colorReset)
}

func TextYellow(text string) string {
	return fmt.Sprintf("%s%s%s", colorYellow, text, colorReset)
}

func TextBlue(text string) string {
	return fmt.Sprintf("%s%s%s", colorBlue, text, colorReset)
}

func TextPurple(text string) string {
	return fmt.Sprintf("%s%s%s", colorPurple, text, colorReset)
}

func TextCyan(text string) string {
	return fmt.Sprintf("%s%s%s", colorCyan, text, colorReset)
}

func TextWhite(text string) string {
	return fmt.Sprintf("%s%s%s", colorWhite, text, colorReset)
}
