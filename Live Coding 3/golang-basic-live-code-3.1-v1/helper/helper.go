package helper

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func InputScan(message string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s", message)
	res, _ := reader.ReadString('\n')
	res = strings.TrimSpace(res)
	return res
}

func ClearScreen() {
	osName := runtime.GOOS

	switch osName {
	case "linux", "darwin": // Untuk Linux dan MacOS
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows": // Untuk Windows 10
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		fmt.Println("Clear screen tidak didukung pada sistem operasi ini")
	}
}

func Delay(duration int) {
	for i := duration; i >= 1; i-- {
		fmt.Printf("\r%d seconds...", i)
		time.Sleep(1 * time.Second)
	}
	fmt.Print("\r")
}
