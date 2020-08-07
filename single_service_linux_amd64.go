// +build linux,amd64
package plock

func checkPlock(pid int) bool {
	if err := syscall.Kill(pid, 0); err == nil {
		//  进程存在
		return true
	}
	return false
}