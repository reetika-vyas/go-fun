package tools

import (
	"fmt"
	"strings"
	"github.com/amanhigh/go-fun/util"
	"github.com/amanhigh/go-fun/kohan/commander"
)

const TIMEOUT = 5

func Jcurl(url string, pipe string) (output string) {
	if commander.IsDebugMode() {
		util.PrintPink(url)
	}

	if pipe == "" {
		output = RunCommandPrintError(fmt.Sprintf("curl -m %v -s '%v' | jq .", TIMEOUT, url))
	} else {
		output = RunCommandPrintError(fmt.Sprintf("curl -m %v -s '%v' | jq . | %v", TIMEOUT, url, pipe))
	}
	return
}

func ContentPiperSplit(content string, pipe string) ([]string) {
	output := ContentPiper(content, pipe)
	return util.FilterEmptyLines(strings.Split(output, "\n"))
}

func ContentPiper(content string, pipe string) (string) {
	output := RunCommandPrintError(fmt.Sprintf("echo '%v' | %v", content, pipe))
	return output
}
