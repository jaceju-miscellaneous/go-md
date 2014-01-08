package markdown

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

// func TestParse(t *testing.T) {
// 	m := new(Markdown)
// 	text := "Hello **World**"
// 	expected := "<p>Hello <strong>World</strong></p>"
// 	result := m.parse(text)
// 	if result != expected {
// 		t.Errorf("'%s' expected but was '%s'.", expected, result)
// 	}
// }

func TestParse(t *testing.T) {
	files, _ := ioutil.ReadDir("./data")
	for _, f := range files {
		// @todo Get file path smart.
		s := []string{"./data/", f.Name()}
		path := strings.Join(s, "")
		content, _ := ioutil.ReadFile(path)
		fmt.Println(string(content))
	}
}

// func getFilePath(f *File) string {
// 	s := []string{"./data/", f.Name()}
// 	return strings.Join(s, "")
// }
