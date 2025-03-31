package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
)

// docker run <image> <cmd> <params>
// go run main.go <cmd> <params>
func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("what?")
	}

}

func run() {
	fmt.Printf("Running %v as %d \n", os.Args[2:], os.Getpid())

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	//
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS,
	}

	cmd.Run()
}

func child() {
	fmt.Printf("Running %v as %d \n", os.Args[2:], os.Getpid())
	cg()
	syscall.Sethostname([]byte("amazon-linux"))
	syscall.Chroot("/home/ec2-user/workdir")
	syscall.Chdir("/")
	syscall.Mount("proc", "proc", "proc", 0, "")

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()

	syscall.Unmount("/proc", 0)
}

func cg() {
	cgroups := "/sys/fs/cgroup"
	pids := filepath.Join(cgroups, "pids")
	err := os.Mkdir(filepath.Join(pids, "test"), 0755)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}

	must(ioutil.WriteFile(filepath.Join(pids, "test/pids.max"), []byte("10"), 0700))
	must(ioutil.WriteFile(filepath.Join(pids, "test/cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700))
	must(ioutil.WriteFile(filepath.Join(pids, "test/notify_on_release"), []byte("1"), 0700))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

