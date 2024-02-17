package app

import (
	"fmt"
	"github.com/fatih/color"
	"strings"
)

type InfoBlock struct {
	title  string
	params map[string]interface{}
}

func (a *App) PrintAppInfo(blocks []InfoBlock) {
	name := color.New(color.FgCyan, color.Bold).SprintFunc()
	version := color.New(color.FgYellow, color.Bold).SprintFunc()
	debug := color.New(color.FgRed, color.Bold).SprintFunc()

	debugStatus := "false"
	if a.Config.Debug {
		debugStatus = "true"
	}

	blocks = append([]InfoBlock{
		{
			title: "App",
			params: map[string]interface{}{
				"name":    name(a.Config.Name),
				"version": version(a.Config.Version),
				"debug":   debug(debugStatus),
			},
		},
	}, blocks...)
	var builder strings.Builder

	border := "================================\n"

	builder.WriteString(border)

	for _, block := range blocks {
		builder.WriteString(fmt.Sprintf("| %s:\n", block.title))

		maxKeyLength := 0
		for key := range block.params {
			if len(key) > maxKeyLength {
				maxKeyLength = len(key)
			}
		}

		for key, value := range block.params {
			spaces := strings.Repeat(" ", maxKeyLength-len(key))
			builder.WriteString(fmt.Sprintf("| - %s:%s %v\n", key, spaces, value))
		}
	}

	builder.WriteString(border)

	builder.WriteString("\nApplication Logs:\n")

	fmt.Print(builder.String())
}
