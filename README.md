Example
===

## Commands

app
--

- app/app.go
and a command line utility to start the node

module
---

- x/{{module}}/module.go
- x/{{module}}/client/client.go with Client type definition
- x/{{module}}/genesis.go

extends:
- app/app.go with a line to register a module. not safe to regenerate, developer may change the process in a non-predictable way
(simplest change would be to change TxDecoder).

type (state object)
---

- x/{{module}}/pb/{{type}}.proto
- x/{{module}}/{{type}}.go registers StateObject in module

Doesn't have to modify `module.go` file.

message (state transition)
---

- x/{{module}}/pb/{{message}}.proto
- x/{{module}}/{{message}}.go StateTransition handler and register in the module

Doesn't have to modify `module.go` file.

For one of the Create/Update/Delete we can have generic implementation in StateTransition handler.
Otherwise just embdy body (see the difference between `create_post` and `change_author`).