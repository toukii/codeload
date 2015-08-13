package execc

import (
	"log"
	"os/exec"
	"strings"
)

func ExecuteCmdHere(command string) ([]byte, bool) {
	log.Println(command)
	cmdWithArgs := strings.Split(command, " ")
	var cmd *exec.Cmd
	cmdLength := len(cmdWithArgs)
	realCmd := cmdWithArgs[0]
	var args []string
	if cmdLength > 1 {
		args = cmdWithArgs[1:cmdLength]
		cmd = exec.Command(realCmd, args...)
	} else {
		cmd = exec.Command(realCmd)
	}
	result, err := cmd.CombinedOutput()
	if err != nil {
		log.Panicln(string(result))
		return nil, false
	}
	return result, true
}
