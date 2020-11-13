package catdog_entry

import (
	"github.com/asim/nitro/v3/server"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type RestEntry interface {
	Entry
}

type RpcEntry interface {
	Entry
}

type Entry interface {
	Init() error
	Start() error
	Stop() error
	Description(description ...string) error
	Version(v string) error
	Flags(fn func(flags *pflag.FlagSet)) error
	Commands(commands ...*cobra.Command) error
	Options() Options
	Handler(handler interface{}, opts ...server.HandlerOption) error
}

type Option func(o *Options)
type Options struct {
	Initialized bool
	RestAddr    string
	Name        string
	Version     string
	RunCommand  *cobra.Command
	Command     *cobra.Command
}
