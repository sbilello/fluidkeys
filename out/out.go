// Copyright 2018 Paul Furley and Ian Drysdale
//
// This file is part of Fluidkeys Client which makes it simple to use OpenPGP.
//
// Fluidkeys Client is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Fluidkeys Client is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with Fluidkeys Client.  If not, see <https://www.gnu.org/licenses/>.

package out

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/fluidkeys/fluidkeys/colour"
)

var outputter outputterInterface
// NoLogCharacter can be added to an output message to prevent that line from
// being saved to the log file.
const NoLogCharacter string = "🤫"

func init() {
	// this is necessary for tests to use out.Print (they don't init
	// through the main function)
	SetOutputToTerminal()
}

func Load(logDirectory string) error {
	if logDirectory == "" {
		return fmt.Errorf("missing log directory")
	}
	SetOutputToTerminal()

	logFilename := filepath.Join(logDirectory, "debug.log")

	if f, err := os.OpenFile(logFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644); err != nil {
		return fmt.Errorf("failed to open '%s' for writing: %v", logFilename, err)
	} else {
		// configure global logger to output to this file
		log.SetOutput(f)
	}
	return nil
}

func SetOutputToTerminal() {
	outputter = &terminalOutputter{}
}

func SetOutputToBuffer() {
	outputter = &bufferOutputter{}
}

func Print(message string) {
	outputter.print(message)

	if lines := splitIntoLogLines(message); len(lines) > 0 {
		for _, line := range lines {
			log.Print(line)
		}
	}
}

func PrintDontLog(message string) {
	outputter.print(message)
}

func PrintTheBuffer() {
	if theBufferOutputter, ok := outputter.(*bufferOutputter); ok {
		theBufferOutputter.printTheBuffer()
	}
}

// splitIntoLogLines takes an output message (with newlines) and returns a
// slice of lines for suitable for outputting with log.Print, namely:
// * remove colour codes
// * remove blank lines
// * redact lines containing NoLogCharacter
func splitIntoLogLines(message string) []string {
	message = colour.StripAllColourCodes(message)

	outLines := []string{}
	for _, line := range strings.Split(message, "\n") {
		if strings.Trim(line, "\n") == "" {
			continue
		}
		outLines = append(outLines, line)
	}

	if strings.Contains(message, NoLogCharacter) {
		outLines = []string{fmt.Sprintf("*** %d lines redacted from log file ***", len(outLines))}
	}

	return outLines
}

type outputterInterface interface {
	print(message string)
}

type terminalOutputter struct{}

func (o *terminalOutputter) print(message string) {
	fmt.Print(message)
}

type bufferOutputter struct {
	buffer string
}

func (o *bufferOutputter) print(message string) {
	o.buffer += message
}

func (o *bufferOutputter) printTheBuffer() {
	fmt.Print(o.buffer)
}
