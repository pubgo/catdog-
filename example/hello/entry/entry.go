package entry

import (
	"github.com/pubgo/catdog"
	"github.com/pubgo/xerror"

	"github.com/pubgo/catdog/example/hello/handler"
	"github.com/pubgo/catdog/version"
)

func GetEntry() catdog.Entry {
	//catdog_config.Domain = "dev"

	ent := catdog.NewEntry()
	xerror.Exit(ent.Name("hello", "hello 服务"))
	xerror.Exit(ent.Version(version.Version))

	xerror.Exit(ent.Handler(handler.NewHelloworld()))
	xerror.Exit(ent.Handler(handler.NewTestAPIHandler()))
	//xerror.Exit(ent.Handler(handler.NewTestAPIHandler1()))
	return ent
}
