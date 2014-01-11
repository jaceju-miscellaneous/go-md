package markdown

import (
	"io/ioutil"
	"path"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	files, _ := ioutil.ReadDir("./data")
	for _, f := range files {
		fileName := f.Name()
		filePath := path.Join("./data/", fileName)
		fileExt := path.Ext(filePath)
		if ".md" != fileExt {
			continue
		}
		markdown, _ := ioutil.ReadFile(filePath)
		input := string(markdown)

		basename := strings.Replace(fileName, fileExt, "", 1)
		htmlPath := strings.Join([]string{"./data/", basename, ".html"}, "")

		html, _ := ioutil.ReadFile(htmlPath)
		expected := string(html)

		result := Parse(input)
		if result != expected {
			t.Errorf("'%s' expected but was '%s'.", expected, result)
		}

	}
}
