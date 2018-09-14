package main

import (
	"net/http"
	"sync/atomic"

	"rlmm-platform/plugin"
	"fmt"
)

type Configuration struct {
	Enabled  bool
	Secret   string
	UserName string
}

type Plugin struct {
	api           plugin.API
	configuration atomic.Value
}

func (p *Plugin) OnActivate(api plugin.API) error {
	p.api = api
	return p.OnConfigurationChange()
}

func (p *Plugin) config() *Configuration {
	return p.configuration.Load().(*Configuration)
}

func (p *Plugin) OnConfigurationChange() error {
	var configuration Configuration
	err := p.api.LoadPluginConfiguration(&configuration)
	p.configuration.Store(&configuration)
	return err
}

func (p *Plugin) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello world")
}
