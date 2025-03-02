package shared

import "os/exec"

func GetSmi() string {
	if !IsNvAvailable {
		return "NV SMI is not available."
	}
	bs, err := exec.Command("nvidia-smi").Output()
	if err != nil {
		return "NV SMI is not available."
	}
	return string(bs)
}
