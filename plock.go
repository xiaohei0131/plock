package plock

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"syscall"
)

const (
	p_file    = "p.lock"
	no_record = -1
)

func Lock() {
	pid := getProcessIdInRecord()
	if pid != no_record {
		sysType := runtime.GOOS
		exist_flag := false
		if sysType == "linux" {
			// LINUX系统
			exist_flag = checkPlock(pid)
		}

		if sysType == "windows" {
			// windows系统
			exist_flag = checkPlock(pid)
		}
		if exist_flag {
			panic(fmt.Sprintf("程序已启动 pid = %d", pid))
		}
	}
	curpid := os.Getpid()
	//fmt.Printf("进程 PID: %d \n", curpid)
	// 记录进程id
	recordProcessId(curpid)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs,syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		//sig := <-sigs
		//fmt.Println(sig)
		<-sigs
		//删除进程id记录文件
		removeFile()
		os.Exit(0)
	}()
}

/**
读取记录的进程id，没有记录返回-1
 */
func getProcessIdInRecord() int {
	b, err := ioutil.ReadFile(p_file)
	if err != nil {
		return no_record
	}
	sb := string(b)
	pid, err := strconv.Atoi(sb)
	if err != nil {
		return no_record
	}
	return pid
}

/**
记录进程id
 */
func recordProcessId(pid int) {
	f, err := os.OpenFile(p_file, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)
	if err != nil {
		panic("记录pid失败")
	}
	data := strconv.Itoa(pid)
	f.WriteString(data) //以字符串写入
	f.Close()
}

/**
删除进程号记录文件
 */
func removeFile() {
	os.Remove(p_file)
}

func UnLock()  {
	removeFile()
}
