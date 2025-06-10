package docxtemplate

import (
	"bytes"
	"strings"

	"baliance.com/gooxml/document"
)

// ReplacePlaceholders открывает docx-файл по пути inputPath,
// заменяет плейсхолдеры вида {{key}} значениями из replacements,
// и возвращает результат как []byte.
func ReplacePlaceholders(inputPath string, replacements map[string]string) ([]byte, error) {
	doc, err := document.Open(inputPath)
	if err != nil {
		return nil, err
	}

	for _, para := range doc.Paragraphs() {
		runs := para.Runs()
		if len(runs) == 0 {
			continue
		}

		var fullText string
		for _, run := range runs {
			fullText += run.Text()
		}

		updatedText := fullText
		for key, val := range replacements {
			placeholder := "{{" + key + "}}"
			updatedText = strings.ReplaceAll(updatedText, placeholder, val)
		}

		if updatedText != fullText {
			for i := len(runs) - 1; i >= 0; i-- {
				para.RemoveRun(runs[i])
			}
			para.AddRun().AddText(updatedText)
		}
	}

	// Сохраняем документ в буфер
	var buf bytes.Buffer
	err = doc.Save(&buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
