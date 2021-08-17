package utils

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// bash utils
func ExecLocalCommand(command string) error {
	args := strings.Split(command, " ")
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Env = os.Environ()
	res, err := cmd.Output()
	if err != nil {
		logrus.Errorf("Fail to execute command: %s, error:%v \n", command, err)
		return err
	}
	logrus.Infof("command: %s executed, \n result: \n %s \n", command, string(res))
	return nil
}

func MkdirIfNotExists(dir string) error {
	cmd := fmt.Sprintf("mkdir -p %s", dir)
	return ExecLocalCommand(cmd)
}

func Copy(src, dest string) error {
	cmdStr := fmt.Sprintf("cp %s %s", src, dest)
	cmd := exec.Command("/bin/sh", "-c", cmdStr)
	res, err := cmd.Output()
	if err != nil {
		logrus.Errorf("Fail to execute command: %s, error:%v \n", cmdStr, err)
		return err
	}
	logrus.Infof("command: %s executed, \n result: \n %s \n", cmdStr, string(res))
	return nil
}

// Get preferred outbound ip of this machine
func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP
}

func GetUUID() string {
	return uuid.New().String()
}

// RenameFile rename a file of name before within dir to name after
func RenameFile(dir, before, after string) error {
	return os.Rename(filepath.Join(dir, before), filepath.Join(dir, after))
}

func ToYaml(input interface{}) []byte {
	yaml, err := yaml.Marshal(input)
	if err != nil {
		panic(err)
	}
	return yaml
}
