package catdog_watcher

import (
	"context"
	"github.com/asim/nitro-plugins/config/source/etcd/v3"
	"github.com/asim/nitro/v3/config"
	"github.com/asim/nitro/v3/config/reader"
	"github.com/pubgo/catdog/catdog_config"
	"github.com/pubgo/catdog/internal/catdog_abc"
	"github.com/pubgo/xerror"
	"github.com/pubgo/xlog"
	"github.com/pubgo/xprocess"
	"strings"
)

func init() {
	// 检查是否有watcher
	xerror.Exit(catdog_abc.WithBeforeStart(func() {
		// 获取
		cfg := catdog_config.GetCfg()
		_, err := cfg.Load("watcher")
		if err != nil {
			xlog.Debugf("config [watcher] is error: %v", err)
			return
		}

		return

		//uri := cfgMap[""]

		xerror.Panic(cfg.Init(
			config.WithSource(
				etcd.NewSource(
					// optionally specify etcd address; default to localhost:8500
					etcd.WithAddress("10.0.0.10:8500"),
					// optionally specify prefix; defaults to /micro/config
					etcd.WithPrefix("/my/prefix"),
					// optionally strip the provided prefix from the keys, defaults to false
					etcd.StripPrefix(true),
				),
			)))
	}))
}

func Watch(name string, watcher func(r reader.Value) error) error {
	if name == "" {
		return xerror.Fmt("[name] should not be empty")
	}

	if watcher == nil {
		return xerror.Fmt("[watcher] should not be nil")
	}

	return xerror.Wrap(catdog_abc.WithBeforeStart(func() {
		key := strings.Join([]string{catdog_config.Project, name}, ".")
		resp := xerror.PanicErr(catdog_config.Load(key)).(reader.Value)
		if resp.Bytes() != nil {
			xerror.Panic(watcher(resp))
		}

		xlog.Debugf("Start Watch Config, Key: %s", key)
		w := xerror.PanicErr(catdog_config.GetCfg().Watch(key)).(config.Watcher)

		// 开启监听配置
		cancel := xprocess.Go(func(ctx context.Context) (err error) {
			defer xerror.RespErr(&err)
			defer func() {
				xlog.Debugf("Stop Watch Config, Key: %s", key)
				xerror.Panic(w.Stop())
			}()

			for {
				select {
				case <-ctx.Done():
					return nil
				default:
					r, err := w.Next()
					if err != nil && strings.Contains(err.Error(), "stopped") {
						return nil
					}
					xerror.Panic(err)
					xerror.Panic(watcher(r))
				}
			}
		})

		// 关闭监听配置变化
		xerror.Exit(catdog_abc.WithAfterStart(func() { xerror.Exit(cancel()) }))
	}))
}