package domain

import "strings"

func ValidateFilename(filename Filename) bool {
	return strings.HasSuffix(filename.Name, ".txt")
}
