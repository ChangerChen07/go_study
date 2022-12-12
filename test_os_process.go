/* package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// 获取当前进程id
	fmt.Printf("os.Getpid(): %v\n", os.Getpid())
	// 获取当前进程父进程id
	fmt.Printf("os.Getppid(): %v\n", os.Getppid())

	// 设置新进程属性
	fmt.Println("通过startprocess启动进程")
	attr := &os.ProcAttr{
		// files指定新进程继承的活动文件对象
		//
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		Env:   os.Environ(),
	}

	fmt.Printf("attr: %v\n", attr)
	p, err := os.StartProcess("C:\\Windows\\System32\\notepad.exe", []string{"C:\\Windows\\System32\\notepad.exe", "E:\\a.txt"}, attr)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("p: %v\n", p)
	fmt.Printf("进程id是: %v\n", p.Pid)

	// 通过进程id查找进程
	fmt.Println("通过进程id查找进程")
	p2, _ := os.FindProcess(p.Pid)
	fmt.Printf("p2: %v\n", p2)

	// 等待10s执行函数
	fmt.Println("10s后，通过os.kill杀掉进程")
	time.AfterFunc(time.Second*10, func() {
		p2.Signal(os.Kill)
	})
	// 等待进程退出并返回状态，阻塞状态。
	ps, _ := p2.Wait()
	fmt.Printf("ps: %v\n", ps)
}
*/