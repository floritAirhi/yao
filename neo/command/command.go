package command

import (
	"fmt"

	"github.com/yaoapp/gou/connector"
	"github.com/yaoapp/yao/aigc"
	"github.com/yaoapp/yao/neo/command/driver"
	"github.com/yaoapp/yao/neo/command/query"
	"github.com/yaoapp/yao/openai"
)

// DefaultStore the default store driver
var DefaultStore Store

// SetStore the driver interface
func SetStore(store Store) {
	DefaultStore = store
}

// Match the command from the content
func Match(sid string, query query.Param, input string) (string, error) {

	if DefaultStore == nil {
		return "", fmt.Errorf("command store is not set")
	}

	// Check the command from the store
	if id, cid, has := DefaultStore.GetRequest(sid); has {
		fmt.Println("Match Requst:", id)
		return cid, nil
	}

	return DefaultStore.Match(query, input)
}

// save the command to the store
func (cmd *Command) save() error {
	if DefaultStore == nil {
		return nil
	}

	args := []map[string]interface{}{}
	for _, arg := range cmd.Args {
		args = append(args, map[string]interface{}{
			"name":        arg.Name,
			"description": arg.Description,
			"type":        arg.Type,
			"required":    arg.Required,
		})
	}

	return DefaultStore.Set(cmd.ID, driver.Command{
		ID:          cmd.ID,
		Description: cmd.Description,
		Args:        args,
		Stack:       cmd.Stack,
		Path:        cmd.Path,
	})
}

// NewAI create a new AI
func (cmd *Command) newAI() (aigc.AI, error) {

	if cmd.Connector == "" {
		return nil, fmt.Errorf("%s connector is required", cmd.ID)
	}

	conn, err := connector.Select(cmd.Connector)
	if err != nil {
		return nil, err
	}

	if conn.Is(connector.OPENAI) {
		return openai.New(cmd.Connector)
	}

	return nil, fmt.Errorf("%s connector %s not support, should be a openai", cmd.ID, cmd.Connector)
}
