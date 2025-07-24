package executer

import (
	"fmt"
)

type EngineInterface interface {
	Set(key string, value []byte)
	Get(key string) ([]byte, bool)
	Delete(key string)
}

type Executer interface {
	Execute() (*string, error)
}

type SetOp struct {
	Key    string
	Value  string
	Engine EngineInterface
}

func NewSetOp(key, value string, eng EngineInterface) Executer {
	return &SetOp{
		Key:    key,
		Value:  value,
		Engine: eng,
	}
}

func (s *SetOp) Execute() (*string, error) {
	s.Engine.Set(s.Key, []byte(s.Value))
	resp := "OK\n"
	return &resp, nil
}

type GetOp struct {
	Key    string
	Engine EngineInterface
}

func NewGetOp(key string, eng EngineInterface) Executer {
	return &GetOp{
		Key:    key,
		Engine: eng,
	}
}

func (g *GetOp) Execute() (*string, error) {
	val, ok := g.Engine.Get(g.Key)
	if !ok {
		resp := "Key not found\n"
		return &resp, nil
	}
	resp := fmt.Sprintf("%s\n", string(val))
	return &resp, nil
}

type DeleteOp struct {
	Key    string
	Engine EngineInterface
}

func NewDeleteOp(key string, eng EngineInterface) Executer {
	return &DeleteOp{
		Key:    key,
		Engine: eng,
	}
}

func (d *DeleteOp) Execute() (*string, error) {
	d.Engine.Delete(d.Key)
	resp := "Deleted\n"
	return &resp, nil
}
