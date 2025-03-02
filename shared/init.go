package shared

import "os/exec"

var IsNvAvailable bool

func init() {
	_, err := exec.Command("nvidia-smi").Output()
	if err != nil {
		IsNvAvailable = false
	} else {
		IsNvAvailable = true
	}
}
