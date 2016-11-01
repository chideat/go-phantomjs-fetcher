package phantomjs

import (
	"syscall"
)

func killProcess(pid int, handlePtr uintptr) {
	syscall.Kill(pid, syscall.SIGKILL)
	syscall.Wait4(pid, nil, 0, nil)
}
