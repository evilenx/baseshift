package helper

import (
	"os"
	"os/exec"
)

func Clear() {
  c := exec.Command("clear")
  c.Stdout = os.Stdout
  c.Run()
}
