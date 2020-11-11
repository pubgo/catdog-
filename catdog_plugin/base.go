package catdog_plugin

import (
	"github.com/pubgo/xerror"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"github.com/asim/nitro/v3/config/reader"
	"github.com/pubgo/catdog/catdog_handler"
)

var _ Plugin = (*Base)(nil)

type Base struct {
	Name       string
	Enabled    bool `yaml:"enabled"`
	OnInit     func()
	OnWatch    func(r reader.Value)
	OnCommands func(cmd *cobra.Command)
	OnHandler  func() *catdog_handler.Handler
	OnFlags    func(flags *pflag.FlagSet)
}

func (p *Base) Init(r reader.Value) (err error) {
	defer xerror.RespErr(&err)

	xerror.Panic(r.Scan(p))

	if !p.Enabled {
		return nil
	}

	if p.OnInit != nil {
		p.OnInit()
	}
	return nil
}

func (p *Base) Watch(r reader.Value) (err error) {
	defer xerror.RespErr(&err)

	if !p.Enabled {
		return nil
	}

	if p.OnWatch != nil {
		p.OnWatch(r)
	}
	return nil
}

func (p *Base) Commands() *cobra.Command {
	if !p.Enabled {
		return nil
	}

	if p.OnCommands != nil {
		cmd := &cobra.Command{Use: p.Name}
		p.OnCommands(cmd)
		return cmd
	}
	return nil
}

func (p *Base) Handler() *catdog_handler.Handler {
	if !p.Enabled {
		return nil
	}

	if p.OnHandler != nil {
		return p.OnHandler()
	}
	return nil
}

func (p *Base) String() string {
	return p.Name
}

func (p *Base) Flags() *pflag.FlagSet {
	if !p.Enabled {
		return nil
	}

	flags := pflag.NewFlagSet(p.Name, pflag.PanicOnError)
	if p.OnFlags != nil {
		p.OnFlags(flags)
	}
	return flags
}