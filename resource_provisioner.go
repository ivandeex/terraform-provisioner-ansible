package main

import (
	"context"
	"fmt"

	"github.com/radekg/terraform-provisioner-ansible/dbg"
	"github.com/radekg/terraform-provisioner-ansible/mode"
	"github.com/radekg/terraform-provisioner-ansible/types"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

type provisioner struct {
	defaults           *types.Defaults
	plays              []*types.Play
	ansibleSSHSettings *types.AnsibleSSHSettings
	remote             *types.RemoteSettings
}

// Provisioner describes this provisioner configuration.
func Provisioner() terraform.ResourceProvisioner {
	return &schema.Provisioner{
		Schema: map[string]*schema.Schema{
			"plays":                types.NewPlaySchema(),
			"defaults":             types.NewDefaultsSchema(),
			"remote":               types.NewRemoteSchema(),
			"ansible_ssh_settings": types.NewAnsibleSSHSettingsSchema(),
		},
		ValidateFunc: validateFn,
		ApplyFunc:    applyFn,
	}
}

// This function should be removed before release
func investigateResourceTypes(c *terraform.ResourceConfig) {
	desc := func(p string) {
		isset := c.IsSet(p)
		val, okval := c.Get(p)
		raw, okraw := c.Get(p)
		num, oknum := c.Get(p + ".#")
		comp := c.IsComputed(p)
		dbg.Log("%q (num=[%d] comp=[%v] type=[%T] raw=[%T] ok=[set=%v val=%v raw=%v num=%v])\n<val>\t%#v\n<raw>\t%#v",
				p, num, comp, val, raw, isset, okval, okraw, oknum, val, raw)
	}

	desc("plays")
	desc("plays.0")
	desc("plays.0.extra_args")
	desc("plays.0.extra_args.0")
	desc("plays.0.extra_args.1")
	desc("plays.0.extra_args.2")
	desc("plays.0.module")
	desc("plays.0.playbook")
	desc("plays.0.playbook.0")
	desc("plays.0.playbook.0.file_path")
	desc("plays.0.playbook.0.roles_path")
	desc("plays.0.playbook.0.roles_path.0")
	desc("plays.0.playbook.0.roles_path.1")
	desc("plays.0.playbook.0.roles_path.2")
	desc("plays.0.playbook.0.roles_path.3")
}

func validateFn(c *terraform.ResourceConfig) (ws []string, es []error) {
	investigateResourceTypes(c)

	defer func() {
		if r := recover(); r != nil {
			es = append(es, fmt.Errorf("error while validating the provisioner, reason: %+v", r))
		}
	}()

	validPlaysCount := 0

	if plays, hasPlays := c.Get("plays"); hasPlays {
		for _, vPlay := range plays.([]map[string]interface{}) {

			currentErrorCount := len(es)

			vPlaybook, playHasPlaybook := vPlay["playbook"]
			_, playHasModule := vPlay["module"]

			if playHasPlaybook && playHasModule {
				es = append(es, fmt.Errorf("playbook and module can't be used together"))
			} else if !playHasPlaybook && !playHasModule {
				es = append(es, fmt.Errorf("playbook or module must be set"))
			} else {

				if playHasPlaybook {
					vPlaybookTyped := vPlaybook.([]map[string]interface{})
					rolesPath, hasRolesPath := vPlaybookTyped[0]["roles_path"]
					if hasRolesPath {
						for _, singlePath := range rolesPath.([]interface{}) {
							vws, ves := types.VfPathDirectory(singlePath, "roles_path")

							for _, w := range vws {
								ws = append(ws, w)
							}
							for _, e := range ves {
								es = append(es, e)
							}
						}
					}
				}

			}

			if currentErrorCount == len(es) {
				validPlaysCount++
			}
		}

		if validPlaysCount == 0 {
			ws = append(ws, "nothing to play")
		}

	} else {
		ws = append(ws, "nothing to play")
	}

	return ws, es
}

func applyFn(ctx context.Context) error {

	o := ctx.Value(schema.ProvOutputKey).(terraform.UIOutput)
	s := ctx.Value(schema.ProvRawStateKey).(*terraform.InstanceState)
	d := ctx.Value(schema.ProvConfigDataKey).(*schema.ResourceData)

	// Decode the provisioner config
	p, err := decodeConfig(d)
	if err != nil {
		return err
	}

	if p.remote.IsRemoteInUse() {
		remoteMode, err := mode.NewRemoteMode(o, s, p.remote)
		if err != nil {
			o.Output(fmt.Sprintf("%+v", err))
			return err
		}
		return remoteMode.Run(p.plays)
	}

	localMode, err := mode.NewLocalMode(o, s)
	if err != nil {
		o.Output(fmt.Sprintf("%+v", err))
		return err
	}
	return localMode.Run(p.plays, p.ansibleSSHSettings)

}

func decodeConfig(d *schema.ResourceData) (*provisioner, error) {

	vRemoteSettings := types.NewRemoteSettingsFromInterface(d.GetOk("remote"))
	vAnsibleSSHSettings := types.NewAnsibleSSHSettingsFromInterface(d.GetOk("ansible_ssh_settings"))
	vDefaults := types.NewDefaultsFromInterface(d.GetOk("defaults"))

	plays := make([]*types.Play, 0)
	if rawPlays, ok := d.GetOk("plays"); ok {
		playSchema := types.NewPlaySchema()
		for _, iface := range rawPlays.([]interface{}) {
			plays = append(plays, types.NewPlayFromInterface(schema.NewSet(schema.HashResource(playSchema.Elem.(*schema.Resource)), []interface{}{iface}), vDefaults))
		}
	}
	return &provisioner{
		defaults:           vDefaults,
		remote:             vRemoteSettings,
		ansibleSSHSettings: vAnsibleSSHSettings,
		plays:              plays,
	}, nil
}
