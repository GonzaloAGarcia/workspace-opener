package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {

	workspace_flag := flag.String("name", "", "set workspace")
	workspace := make(map[string]string)
	flag.Parse()

	// ToDo: get workspaces from file and add new ones there.
	if *workspace_flag == "go_course" {
		workspace = map[string]string{
			`C:\Users\gonza\AppData\Local\Obsidian\obsidian.exe`:                             "explorer",
			`C:\Users\gonza\go\Code\golang\googles_go_course`:                                "code",
			`C:\Users\gonza\Downloads\course+outline+-+learn+to+code+go+-+v3.18.0.pdf`:       "explorer",
			`https://www.udemy.com/course/learn-how-to-code/learn/lecture/38376704#overview`: "explorer",
		}
	} else {
		fmt.Println("No workspace matches that flag.")
		// ToDo: Print available workspaces
		os.Exit(1)
	}

	for filepath, command := range workspace {
		cmd := exec.Command(command, filepath)
		err := cmd.Start()
		if err != nil {
			fmt.Println("Error running the command:", err)
		}
	}
	os.Exit(0)
}
