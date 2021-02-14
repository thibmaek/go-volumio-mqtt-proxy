package volumio

import "os/exec"

// Execute an MPC command
func Mpc(command string) ([]byte, error) {
	cmd := exec.Command("mpc", command)
	return cmd.CombinedOutput()
}
