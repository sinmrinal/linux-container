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

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("bad command")
	}
}

func run() {
	fmt.Printf("Running %v as %d\n", os.Args[2:], os.Getpid())

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
		Unshareflags: syscall.CLONE_NEWNS,
	}
	cmd.Run()

	syscall.Sethostname([]byte("container"))
}
func child() {
	fmt.Printf("Running %v as %d\n", os.Args[2:], os.Getpid())

	controlGroup()
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	catch(syscall.Sethostname([]byte("container")))
	catch(syscall.Chroot("./fs/ubuntufs"))
	catch(syscall.Chdir("/"))
	catch(syscall.Mount("proc", "proc", "proc", 0, ""))

	cmd.Run()

	catch(syscall.Unmount("proc", 0))
}

func controlGroup() {
	cgroup := "/sys/fs/cgroup"
	pids := filepath.Join(cgroup, "pids")
	err := os.Mkdir(filepath.Join(pids, "cmrinal"), 0755)
	if err != nil && !os.IsExist(err) {
		panic(err)
	}

	catch(ioutil.WriteFile(filepath.Join(pids, "cmrinal/pids.max"), []byte("10"), 0700))
	catch(ioutil.WriteFile(filepath.Join(pids, "cmrinal/notify_on_release"), []byte("1"), 0700))
	catch(ioutil.WriteFile(filepath.Join(pids, "cmrinal/cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700))
}

func catch(err error) {
	if err != nil {
		panic(err)
	}
}
