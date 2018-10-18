package scheduler

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
	"syscall"
)

// Enable adds cron lines to the user's crontab and returns whether the
// crontab was updated.
func Enable() (crontabWasAdded bool, err error) {
	currentCrontab, err := getCurrentCrontab()
	if err != nil {
		return false, fmt.Errorf("error getting crontab: %v", err)
	}

	if !hasFluidkeysCronLines(currentCrontab) {
		newCrontab := addCrontabLinesWithoutRepeating(currentCrontab)
		err = writeCrontab(newCrontab)
		if err != nil {
			return false, err
		}
		crontabWasAdded = true
	} else {
		crontabWasAdded = false
	}

	return
}

// Disable parses the crontab (output of `crontab -l`) and removes Fluidkeys'
// cron lines if present.
// If the remaining crontab is empty, the crontab is removed with `crontab -r`
func Disable() (cronLinesWereRemoved bool, err error) {
	currentCrontab, err := getCurrentCrontab()
	if err != nil {
		return false, fmt.Errorf("error getting crontab: %v", err)
	}

	if hasFluidkeysCronLines(currentCrontab) {
		cronLinesWereRemoved = true
		newCrontab := removeCrontabLines(currentCrontab)
		err = writeCrontab(newCrontab)
		return
	} else {
		cronLinesWereRemoved = false
		return
	}
}

func hasFluidkeysCronLines(crontab string) bool {
	return strings.Contains(crontab, cronLines)
}

func getCurrentCrontab() (string, error) {
	output, err := runCrontab("-l")
	if err != nil {
		if isExitStatusOne(err) && strings.Contains(output, "no crontab for") {
			return "", nil
		}
	}
	return output, err
}

func writeCrontab(newCrontab string) error {
	if isEmptyCrontab(newCrontab) {
		_, err := runCrontab("-r") // remove the user's crontab
		return err
	} else {
		f, err := ioutil.TempFile("", "")
		if err != nil {
			return err
		}

		f.Write([]byte(newCrontab))
		f.Close()

		_, err = runCrontab(f.Name())
		return err
	}
}

func isEmptyCrontab(crontab string) bool {
	// TODO: strip newlines
	return crontab == ""
}

func isExitStatusOne(err error) bool {
	if exiterr, ok := err.(*exec.ExitError); ok {
		if _, ok := exiterr.Sys().(syscall.WaitStatus); ok {
			return true
		}
	}
	return false
}

func runCrontab(arguments ...string) (string, error) {
	cmd := exec.Command(crontab, arguments...)

	out, err := cmd.CombinedOutput()

	outString := string(out)

	if err != nil {
		return outString, err
	}
	return outString, nil
}

func addCrontabLinesWithoutRepeating(crontab string) string {
	return removeCrontabLines(crontab) + cronLines
}

func removeCrontabLines(crontab string) string {
	return strings.Replace(crontab, cronLines, "", -1)
}

const crontab string = "crontab"
const cronLines string = `
# Fluidkeys added the following line. To disable, edit your Fluidkeys configuration file.
@hourly /usr/local/bin/fk key rotate --automatic --cron-output
`
