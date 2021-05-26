package framework

import (
	"context"

	"github.com/spf13/cobra"
)

// Module defines the interface for framework modules to implement.
type Module interface {
	Name() string
	Genesis(*Genesis)
	AddCommands(root *cobra.Command)
}

// Genesis represents the genesis.
type Genesis struct{}

// Message defines message interface
type Message interface {
	Handle(Request) error
}

// Request represents a new request related to a Type.
type Request struct {
}

func (r *Request) Keeper() *Keeper {
	return &Keeper{}
}

type Keeper struct {
}

func (k *Keeper) Save(key string, val interface{}) error { return nil }

// RegisterModule registers a module.
func RegisterModule(Module) {}

// RegisterMessage registers a messages.
func RegisterMessage(Message) {}

type Command int

const (
	StartCommand Command = iota
	InitCommand
	AddAccountCommand
)

// EnableCommand enables a top level command.
func EnableCommand(Command) {
}

// Run runs the app.
func Run(ctx context.Context) error { return nil }
