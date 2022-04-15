package utils

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func B2i(b bool) int {
	if b {
		return 2
	}
	return 0
}

// GetCurrentAbPath 最终方案-全兼容
func GetCurrentAbPath() string {
	dir := getCurrentAbPathByExecutable()
	tmpDir, _ := filepath.EvalSymlinks(os.TempDir())
	if strings.Contains(dir, tmpDir) {
		return filepath.Dir(getCurrentAbPathByCaller())
	}
	return dir
}

// 获取当前执行文件绝对路径
func getCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// 获取当前执行文件绝对路径（go run）
func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

func Shell(cmd string) (cmdRes string, err error) {
	command := exec.Command("bash", "-c", cmd)
	outinfo := bytes.Buffer{}
	outerr := bytes.Buffer{}
	command.Stdout = &outinfo
	command.Stderr = &outerr

	if err = command.Start(); err != nil {
		return "", errors.New(fmt.Sprintf("command start outErr: %v, err: %v", outerr.String(), err.Error()))
	}

	if err = command.Wait(); err != nil {
		return "", errors.New(fmt.Sprintf("command wait outErr: %v, err: %v, pid: %v", outerr.String(), err.Error(), command.Process.Pid))
	}

	cmdRes = outinfo.String()
	return
}

func GetMachineId() (id string, err error) {
	machineId, err := ioutil.ReadFile("/etc/machine-id")
	id = strings.TrimSpace(string(machineId))
	if err != nil {
		return "", err
	}

	return
}
