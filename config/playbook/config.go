package playbook

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("ansible_playbook", func(r *config.Resource) {
		r.ShortGroup = "playbook"
	})
}
