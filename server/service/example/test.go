package example

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("python", "script.py")
	output, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(output))
}
