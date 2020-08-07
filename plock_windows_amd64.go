package plock

import "os"

func checkPlock(pid int) bool {
	p,err := os.FindProcess(pid)
	if p != nil && err == nil{
		//  进程存在
		return true
	}
	return false
}