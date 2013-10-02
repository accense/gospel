package gospel

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"strings"
)

type Formatter interface {
	Started(*Example)
	Failed(*Example, string, interface{}, interface{})
	Succeeded(*Example)
}

type DotFormatter struct {}

func (formatter *DotFormatter) Started(example *Example) {
}

func (formatter *DotFormatter) Failed(example *Example, message string, actual, expected interface{}) {
	fmt.Print(red("F"))
	if example.Describing.Result == "" {
		example.Describing.Result += "\n\n"
	}
	_, filename, line, _ := runtime.Caller(4)
	buffer, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(buffer), "\n")[line-2:line+2]
	example.Describing.Result += fmt.Sprintf(
		red("  %s\n") +
		grey("  Expected `%v` to %s `%v`\n") +
		grey("  %s:%d\n") +
		grey("  %4d.%s\n") +
		grey("  %4d.%s\n") +
		grey("  %4d.%s\n") +
		"\n",
		example.FullDescription(),
		actual, message, expected,
		filename, line,
		line - 1, strings.Replace(lines[0], "\t", "  ", -1),
		line + 0, strings.Replace(lines[1], "\t", "  ", -1),
		line + 1, strings.Replace(lines[2], "\t", "  ", -1),
	)
}

func (formatter *DotFormatter) Succeeded(example *Example) {
	fmt.Print(green("."))
}

type DocumentFormatter struct {}

func (formatter *DocumentFormatter) Started(example *Example) {
	fullMessage := ""
	if example.DoneExamplesCount == 0 {
		fullMessage += example.Description + "\n"
	}
	if len(example.SubDescriptions) > 0 {
		differenceIsFound := false
		for i, subscription := range example.SubDescriptions {
			if !differenceIsFound && i <= len(example.PreviousSubDescriptions) - 1 {
				if subscription == example.PreviousSubDescriptions[i] {
					continue
				}
			}
			fullMessage += strings.Repeat("  ", i + 1) + subscription + "\n"
			differenceIsFound = true
		}
	}
	fmt.Print(fullMessage)
}

func (formatter *DocumentFormatter) Succeeded(example *Example) {
	margin := strings.Repeat("  ", len(example.SubDescriptions) + 1)
	fmt.Println(margin + green(example.Message))
}

func (formatter *DocumentFormatter) Failed(example *Example, message string, actual, expected interface{}) {
	_, filename, line, _ := runtime.Caller(4)
	buffer, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(buffer), "\n")[line-2:line+2]
	margin := strings.Repeat("  ", len(example.SubDescriptions) + 1)
	fmt.Printf(
		red("%s%s\n") +
		grey("%sExpected `%v` to %s `%v`\n") +
		grey("%s%s:%d\n") +
		grey("%s%4d.%s\n") +
		grey("%s%4d.%s\n") +
		grey("%s%4d.%s\n"),
		margin, example.Message,
		margin, actual, message, expected,
		margin, filename, line,
		margin, line - 1, strings.Replace(lines[0], "\t", "  ", -1),
		margin, line + 0, strings.Replace(lines[1], "\t", "  ", -1),
		margin, line + 1, strings.Replace(lines[2], "\t", "  ", -1),
	)
}

// Add red terminal ANSI color
func red(str string) string {
	return "\033[31m\033[1m" + str + "\033[0m"
}

// Add green terminal ANSI color
func green(str string) string {
	return "\033[32m\033[1m" + str + "\033[0m"
}

// Add grey terminal ANSI color
func grey(str string) string {
	return "\x1B[90m" + str + "\033[0m"
}
