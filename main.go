package main

import (
	"fmt"
	"os"
	"os/exec"
)

const script = `
	on run argv
	tell application "System Events"
        tell appearance preferences
            set dark mode to (item 1 of argv)
        end tell
    end tell
	end run
`

func main() {
	var arg string
	switch os.Args[1] {
	case "dark":
		arg = "true"
	case "light":
		arg = "false"
	default:
		fmt.Println("expected 'dark' or 'light'")
		os.Exit(1)
	}

	cmd := exec.Command("osascript", "-e", script, arg)
	_, err := cmd.StderrPipe()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := cmd.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
