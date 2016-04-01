package ini

import "strings"

// ParseINI converts s into a map of sections and property/value pairs.
// It assumes s is in proper INI format.
func Parse(s string) map[string]map[string]string {
	result := make(map[string]map[string]string)

	var currentSection string
	for _, line := range strings.Split(s, "\n") {
		if len(line) == 0 || line[0] == '#' || line[0] == ';' {
			continue
		}

		if line[0] == '[' {
			startIndex, endIndex := strings.Index(line, "["), strings.Index(line, "]")
			currentSection = line[startIndex+1 : endIndex]
			result[currentSection] = make(map[string]string)
		} else if delimiter := strings.Index(line, "="); delimiter != -1 {
			endIndex := len(line)
			if commentStart := strings.IndexAny(line, ";#"); commentStart != -1 {
				endIndex = commentStart
			}

			property, value := line[:delimiter], line[delimiter+1:endIndex]
			result[currentSection][property] = strings.TrimSpace(value)
		}
	}

	return result
}
