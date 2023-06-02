package main

import (
	"strings"

	"github.com/nsf/termbox-go"
	"github.com/russross/blackfriday/v2"
)

func main() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}
	defer termbox.Close()

	input := ""
	preview := ""

	// Render the preview initially
	preview = renderMarkdown(input)

	render(input, preview)

mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyCtrlC, termbox.KeyEsc:
				break mainloop
			case termbox.KeyBackspace, termbox.KeyBackspace2:
				if len(input) > 0 {
					input = input[:len(input)-1]
				}
			case termbox.KeyEnter:
				input += "\n"
			case termbox.KeySpace:
				input += " "
			default:
				if ev.Ch != 0 {
					input += string(ev.Ch)
				}
			}

			preview = renderMarkdown(input)
			render(input, preview)
		case termbox.EventResize:
			render(input, preview)
		}
	}
}

func renderMarkdown(input string) string {
	// Render the Markdown using blackfriday package
	output := blackfriday.Run([]byte(input))
	return string(output)
}

func render(input, preview string) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	// Render the input area
	for i, line := range strings.Split(input, "\n") {
		for j, ch := range line {
			termbox.SetCell(j, i, ch, termbox.ColorDefault, termbox.ColorDefault)
		}
	}

	// Render the preview area
	previewLines := strings.Split(preview, "\n")
	for i, line := range previewLines {
		for j, ch := range line {
			style := termbox.ColorDefault
			if isHeader(line) {
				style = termbox.ColorGreen // Set header color
			} else if isList(line) {
				style = termbox.ColorBlue // Set list color
			}
			termbox.SetCell(j, i+len(strings.Split(input, "\n"))+1, ch, style, termbox.ColorDefault)
		}
	}

	termbox.Flush()
}

func isHeader(line string) bool {
	// Check if the line is a Markdown header (starts with '#')
	return strings.HasPrefix(line, "#")
}

func isList(line string) bool {
	// Check if the line is a Markdown list item (starts with '-')
	return strings.HasPrefix(line, "-")
}
