package markdown

import (
	"fmt"
	"regexp"
	"strings"
)

var escape_sequences = []string{"\\\\", "\\`", "\\*", "\\_", "\\{", "\\}", "\\[", "\\]", "\\(", "\\)", "\\>", "\\#", "\\+", "\\-", "\\.", "\\!"}
var escape_sequence_map = map[string]string{}

type Element struct {
	_type        string
	_text        string
	_level       int
	_closed      bool
	_interrupted bool
	_fence       []string
	_lines       []string
}

func Parse(text string) string {

	text = _PrepareText(text)

	text = _EncodeText(text)

	lines := _TextToLines(text)

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

func _PrepareText(text string) string {
	// removes UTF-8 BOM and marker characters
	re := regexp.MustCompile("^\xEF\xBB\xBF|\x1A")
	text = re.ReplaceAllString(text, "")

	// removes \r characters
	text = strings.Replace(text, "\r\n", "\n", -1)
	text = strings.Replace(text, "\r", "\n", -1)

	// replaces tabs with spaces
	text = strings.Replace(text, "\t", "    ", -1)

	return text
}

func _EncodeText(text string) string {
	if strings.Index(text, "\\") >= 0 {
		for i, v := range escape_sequences {
			if strings.Index(text, v) >= 0 {
				code := strings.Join([]string{"\x1A", "\\", string(i), ";"}, "")
				text = strings.Replace(text, v, code, -1)
				escape_sequence_map[code] = v
			}
		}
	}
	return text
}

func _TextToLines(text string) []string {
	re := regexp.MustCompile("\\n\\s*\\n")
	text = re.ReplaceAllString(text, "\n\n")
	text = strings.TrimPrefix(text, "\n")
	text = strings.TrimSuffix(text, "\n")
	return strings.Split(text, "\n")
}

func _ParseBlockElements(lines []string) string {
	elements := []Element{}

	element := new(Element)
	element._type = ""
	element._text = ""
	element._level = 0
	element._closed = false
	element._interrupted = false
	element._fence = []string{}

Start:
	for _, line := range lines {
		fmt.Println(line)

		if "" == line {
			element._interrupted = true
			continue Start
		}

		// indentation sensitive types

		// deindented_line := line

		switch string(line[0]) {
		case "#":
			// atx heading (#)

			// if (preg_match('/^(#{1,6})[ ]*(.+?)[ ]*#*$/', $line, $matches)) {
			//     $elements []= $element;

			//     $level = strlen($matches[1]);

			//     $element = array(
			//         'type' => 'h.',
			//         'text' => $matches[2],
			//         'level' => $level,
			//     );

			//     continue 2;
			// }

			// break;
		}

		elements = append(elements, *element)
	}

	return strings.Join(lines, "\n")
}
