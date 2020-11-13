package catdog_pidfile

import (
	"github.com/pubgo/catdog/catdog_util"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/pubgo/catdog/catdog_config"
	"github.com/pubgo/catdog/internal/catdog_abc"
	"github.com/pubgo/xerror"
)

const pidPerm os.FileMode = 0755

func GetPid() (pid int, _ error) {
	pidData, err := ioutil.ReadFile(GetPidPath())
	if err != nil {
		return 0, xerror.Wrap(err)
	}

	pid, err = strconv.Atoi(string(pidData))
	return pid, xerror.Wrap(err)
}

func SavePid() error {
	pidBytes := []byte(strconv.Itoa(os.Getpid()))
	return xerror.Wrap(ioutil.WriteFile(GetPidPath(), pidBytes, pidPerm))
}

func GetPidPath() string {
	return filepath.Join(catdog_config.Home, "pidfile", catdog_config.Domain+"."+catdog_config.Project+".pid")
}

func init() {
	// 检查存放pid的目录是否存在, 不存在就创建
	xerror.Panic(catdog_abc.WithBeforeStart(func() {
		pidPath := filepath.Dir(GetPidPath())
		if !catdog_util.PathExist(pidPath) {
			xerror.Exit(os.MkdirAll(pidPath, pidPerm))
		}
	}))

	// 保存pid到文件当中
	xerror.Panic(catdog_abc.WithAfterStart(func() {
		xerror.Panic(SavePid())
	}))
}

func Init() {}