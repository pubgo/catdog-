package rpc_entry

import (
	grpcC "github.com/asim/nitro-plugins/client/grpc/v3"
	"github.com/asim/nitro-plugins/server/grpc/v3"
	"github.com/asim/nitro/v3/client"
	"github.com/pubgo/catdog/catdog_entry"
	"github.com/pubgo/catdog/catdog_entry/base_entry"
)

type entry struct {
	catdog_entry.Entry
	c client.Client
}

func newEntry(name string) *entry {
	ent := &entry{
		Entry: base_entry.New(
			name,
			grpc.NewServer(),
		),
		c: grpcC.NewClient(),
	}

	return ent
}

func New(name string) catdog_entry.Entry {
	return newEntry(name)
}
