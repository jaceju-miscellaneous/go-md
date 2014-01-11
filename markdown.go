package markdown

import (
	"fmt"
	"regexp"
	"strings"
)

func Parse(text string) string {

	// removes UTF-8 BOM and marker characters
	re := regexp.MustCompile("^\xEF\xBB\xBF|\x1A")
	text = re.ReplaceAllString(text, "")

	escape_sequences := []string{"\\\\", "\\`", "\\*", "\\_", "\\{", "\\}", "\\[", "\\]", "\\(", "\\)", "\\>", "\\#", "\\+", "\\-", "\\.", "\\!"}
	escape_sequence_map := map[string]string{}

	// removes \r characters
	text = strings.Replace(text, "\r\n", "\n", -1)
	text = strings.Replace(text, "\r", "\n", -1)

	// replaces tabs with spaces
	text = strings.Replace(text, "\t", "    ", -1)

	// encodes escape sequences

	if strings.Index(text, "\\") >= 0 {
		for i, v := range escape_sequences {
			if strings.Index(text, v) >= 0 {
				code := strings.Join([]string{"\x1A", "\\", string(i), ";"}, "")
				text = strings.Replace(text, v, code, -1)
				escape_sequence_map[code] = v
			}
		}
	}

	// ~
	re = regexp.MustCompile("\\n\\s*\\n")
	text = re.ReplaceAllString(text, "\n\n")
	text = strings.TrimPrefix(text, "\n")
	text = strings.TrimSuffix(text, "\n")
	lines := strings.Split(text, "\n")
	text = _ParseBlockElements(lines)

	// decodes escape sequences
	for code, escape_sequence := range escape_sequences {
		text = strings.Replace(text, string(code), string(escape_sequence[1]), -1)
	}

	// ~
	text = strings.TrimSuffix(text, "\n")

	fmt.Println(text)

	return text
}

func _ParseBlockElements(lines []string) string {
	return strings.Join(lines, "\n")
}
