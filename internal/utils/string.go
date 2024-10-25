package utils

import (
	"strings"
)

func SplitText(text string, maxLength int) []string {
	partsToAssemble := make([]string, 0)
	paragraphs := strings.Split(text, "\n")
	for _, paragraph := range paragraphs {
		if len(paragraph) > maxLength {
			partsToAssemble = append(partsToAssemble, strings.Split(paragraph, ` `)...)
		} else {
			partsToAssemble = append(partsToAssemble, paragraph)
		}
	}

	result := make([]string, 0)
	currentPart := ""
	for _, part := range partsToAssemble {
		if len(currentPart)+len(part) > maxLength {
			result = append(result, currentPart)
			currentPart = part
		} else {
			currentPart += part
		}
	}
	if currentPart != "" {
		result = append(result, currentPart)
	}

	return result
}
