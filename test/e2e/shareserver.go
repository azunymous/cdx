//+build e2e

package e2e

import (
	"log"
	"net"
	"os/exec"
	"strconv"
)

func StartCdxShareServer() (string, func()) {
	port, err := GetFreePort()
	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command(CDX, "share", "start", "--insecure", "--port", strconv.Itoa(port))
	_ = cmd.Start()
	return "localhost:" + strconv.Itoa(port), func() {
		_ = cmd.Process.Kill()
	}
}

// GetFreePort asks the kernel for a free open port that is ready to use.
func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}
