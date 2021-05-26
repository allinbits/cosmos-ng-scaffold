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

// ModuleContext keeps mutable configuration for the module.
type ModuleContext struct {
	Command *cobra.Command
	Keeper  *Keeper
}

// Genesis represents the genesis.
type Genesis struct{}

// Message defines message type.
type Message interface{}

// QueryRequest defines query request type.
type QueryRequest interface{}

// QueryResponse defines response request type.
type QueryResponse interface{}

// Handler defines handler interface.
type Handler interface {
	// Context provides module's context to the type for advanced configuration.
	Context(*ModuleContext)

	// Command optionally adds another command for the module for this message type.
	Command() *cobra.Command

	// Handler must implement one of the Handle functions depending on if it is a message request or
	// a query request, otherwise RegisterQuery or RegisterMessage will panic.
	//   Handle(*ModuleContext, T<Message>) (error)
	//   Handle(*ModuleContext, T<QueryRequest>) (T<QueryResponse>, error)
}

// Request represents a new request related to a Type.
type Request struct {
	Payload interface{}
}

func (r *Request) Context() *ModuleContext {
	return &ModuleContext{}
}

type Keeper struct {
}

func (k *Keeper) Save(key string, val interface{}) error { return nil }

func (k *Keeper) Get(key string, out interface{}) error { return nil }

// RegisterModule registers a module.
func RegisterModule(Module) {}

// RegisterMessage registers a messages.
func RegisterMessage(Message, Handler) {}

// RegisterQuery registers a query.
func RegisterQuery(QueryRequest, QueryResponse, Handler) {}

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
