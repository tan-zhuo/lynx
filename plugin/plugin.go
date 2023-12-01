package plugin

import (
	"errors"
	"github.com/go-kratos/kratos/v2/config"
)

var (
	factory = newGlobalPluginFactory()
)

func GlobalPluginFactory() *Factory {
	return factory
}

type Plugin interface {
	Weight() int
	Name() string
	Load(config.Value) (Plugin, error)
	Unload() error
}

type Manger interface {
	LoadPlugins(map[string]config.Value)
	UnloadPlugins()
	LoadSpecificPlugins(plugins []string, config map[string]config.Value)
	UnloadSpecificPlugins(plugins []string)
	GetPlugin(name string) Plugin
}

type Factory struct {
	creators map[string]func() Plugin
}

func newGlobalPluginFactory() *Factory {
	return &Factory{
		creators: make(map[string]func() Plugin),
	}
}

func (f *Factory) Register(name string, creator func() Plugin) {
	f.creators[name] = creator
}

func (f *Factory) Create(name string) (Plugin, error) {
	creator, exists := f.creators[name]
	if !exists {
		return nil, errors.New("invalid plugin name")
	}
	return creator(), nil
}

func (f *Factory) Exists(name string) bool {
	_, exists := f.creators[name]
	return exists
}
