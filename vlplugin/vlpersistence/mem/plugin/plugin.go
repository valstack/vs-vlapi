package main

import (
	"github.com/VolantMQ/vlapi/vlplugin"
	persistenceMem "github.com/VolantMQ/vlapi/vlplugin/vlpersistence/mem"
)

type persistencePlugin struct {
	vlplugin.Descriptor
}

var _ vlplugin.Plugin = (*persistencePlugin)(nil)
var _ vlplugin.Info = (*persistencePlugin)(nil)

// Plugin symbol
var Plugin persistencePlugin

func init() {
	Plugin.V = vlplugin.Version()
	Plugin.N = "mem"
	Plugin.T = "persistence"
}

// Load plugin
func (pl *persistencePlugin) Load(c interface{}, params *vlplugin.SysParams) (interface{}, error) {
	return persistenceMem.Load(c, params)
}

// Info plugin info
func (pl *persistencePlugin) Info() vlplugin.Info {
	return pl
}

func main() {
	panic("this is a plugin, build it as with -buildmode=plugin")
}
