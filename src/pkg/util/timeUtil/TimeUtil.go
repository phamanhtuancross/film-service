package timeUtil

import "os/exec"

func GenerateIdString() string {
	uuidIdString, err := exec.Command("uuidgen").Output()

	if err != nil {
		return ""
	}

	return string(uuidIdString)
}