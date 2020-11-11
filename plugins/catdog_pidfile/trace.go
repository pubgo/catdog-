package catdog_pidfile

import (
	"fmt"
	"github.com/pubgo/catdog/catdog_config"
	"github.com/pubgo/catdog/internal/catdog_abc"
	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog"
)

func init() {
	xerror.Exit(catdog_abc.WithAfterStart(func() {
		if !catdog_config.Trace {
			return
		}

		pid, err := GetPid()
		xerror.Panic(err)
		xlog.Debugf("path, pid trace")
		fmt.Println(GetPidPath(), pid)
	}))
}
