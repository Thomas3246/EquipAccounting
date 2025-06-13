package docxtemplate

import (
	"bytes"
	"strings"

	"baliance.com/gooxml/document"
	"baliance.com/gooxml/measurement"
)

func ReplacePlaceholders(inputPath string, replacements map[string]string) ([]byte, error) {
	doc, err := document.Open(inputPath)
	if err != nil {
		return nil, err
	}

	addFormattedRun := func(para *document.Paragraph, text string) {
		r := para.AddRun()
		r.AddText(text)
		r.Properties().SetFontFamily("Times New Roman")
		r.Properties().SetSize(14 * measurement.Point)
	}

	paragraphs := doc.Paragraphs()
	for i := range paragraphs {
		para := &paragraphs[i]

		var fullTextBuilder strings.Builder
		for _, run := range para.Runs() {
			fullTextBuilder.WriteString(run.Text())
		}
		fullText := fullTextBuilder.String()

		hasReplacement := false
		for key := range replacements {
			if strings.Contains(fullText, "{{"+key+"}}") {
				hasReplacement = true
				break
			}
		}

		if !hasReplacement {
			continue
		}

		for len(para.Runs()) > 0 {
			para.RemoveRun(para.Runs()[0])
		}

		start := 0
		for start < len(fullText) {
			beginIdx := strings.Index(fullText[start:], "{{")
			if beginIdx == -1 {
				break
			}
			beginIdx += start
			endIdx := strings.Index(fullText[beginIdx:], "}}")
			if endIdx == -1 {
				break
			}
			endIdx += beginIdx + 2

			if beginIdx > start {
				normalText := fullText[start:beginIdx]
				addFormattedRun(para, normalText)
			}

			placeholder := fullText[beginIdx:endIdx]
			key := strings.TrimSuffix(strings.TrimPrefix(placeholder, "{{"), "}}")
			if value, ok := replacements[key]; ok {
				addFormattedRun(para, value)
			} else {
				addFormattedRun(para, placeholder)
			}

			start = endIdx
		}

		if start < len(fullText) {
			remainingText := fullText[start:]
			addFormattedRun(para, remainingText)
		}
	}

	var buf bytes.Buffer
	if err := doc.Save(&buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
