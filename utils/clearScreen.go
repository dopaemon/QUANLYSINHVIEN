package utils

import (
	"runtime"
	"os"
	"os/exec"
)

var (
)

func ClearScreen() {
        clearCmd := "clear"
        if runtime.GOOS == "windows" {
                clearCmd = "cls"
        }

        cmd := exec.Command(clearCmd)
        cmd.Stdout = os.Stdout
        cmd.Run()
}
