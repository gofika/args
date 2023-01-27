package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

func main() {
	argsWithProg := os.Args
	appPath := argsWithProg[0]
	argsWithoutProg := argsWithProg[1:]
	windowsExeFlag := ".exe"
	originalFlag := ".actual"

	fmt.Println("Startup Args:")
	cmdline := ""
	for _, s := range argsWithoutProg {
		if len(cmdline) > 0 {
			cmdline += " "
		}
		if strings.Contains(s, " ") {
			cmdline += "\"" + s + "\""
		} else {
			cmdline += s
		}
	}
	fmt.Println(cmdline)
	appName := filepath.Base(appPath)
	logFile := appPath + ".laststartup.log"
	_ = os.WriteFile(logFile, []byte(cmdline), os.ModeAppend)
	if strings.ToLower(filepath.Ext(appPath)) == windowsExeFlag {
		// windows
		appName = appName[:len(appName)-4] + originalFlag + windowsExeFlag
	} else {
		// linux
		appName = appName + originalFlag
	}

	origPath := path.Join(filepath.ToSlash(filepath.Dir(appPath)), appName)

	fmt.Println("Press Enter to run original (" + origPath + ")...")
	var wait string
	fmt.Scanln(&wait)
	cmd := exec.Command(origPath, argsWithoutProg...)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}
