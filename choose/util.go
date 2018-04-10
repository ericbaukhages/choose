package choose

import (
	"errors"
	"os"
	"os/exec"
	"runtime"
)

var (
	clear map[string]func()
)

// CallClear clears the terminal screen
func CallClear() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Your platform is unsupported, I can't clear the screen!")
	}
}

// CallClearInit initialzes the CallClear necessary components
func CallClearInit() {
	// set up the clear func
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["darwin"] = clear["linux"]
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// ValidPath checks if given path is valid
func ValidPath(path string) error {
	stat, err := os.Stat(path)
	if !(err == nil && stat.IsDir()) {
		return errors.New("path is not valid directory")
	}

	return nil
}
