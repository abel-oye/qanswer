package tap

import "os/exec"

func SendSwipe(x string,y string){
	err := exec.Command("adb", "shell", "input", "tap", x,y).Run()
	if err != nil {
		return
	}
}
