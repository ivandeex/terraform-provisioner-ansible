package main

import (
	"context"
	"errors"
	"fmt"
	"log"

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

func validateFn(c *terraform.ResourceConfig) (ws []string, es []error) {

	defer func() {
		if r := recover(); r != nil {
			es = append(es, fmt.Errorf("error while validating the provisioner, reason: %+v", r))
		}
	}()

	if !c.IsSet("plays") {
		ws = append(ws, "nothing to play")
		return
	}

	value, _ := c.Get("plays")
	_, isArrayOfMaps := value.([]map[string]interface{}) // in Terraform 0.11.x types are concrete
	_, isArrayOfIfs := value.([]interface{})             // in Terraform 0.12+ types are wrapped in interface{}
	value, _ = c.Get("plays.#")
	numPlays, ok := value.(int)
	if (!isArrayOfIfs && !isArrayOfMaps) || value == nil || !ok {
		es = append(es, fmt.Errorf("`plays` must be an array isarr=%v/%v ok=%v val=%v num=%d", isArrayOfIfs, isArrayOfMaps, ok, value, numPlays))
		return
	}

	validPlaysCount := 0
	for playNo := 0; playNo < numPlays; playNo++ {
		playLoc := fmt.Sprintf("plays.%d", playNo)

		playHasPlaybook := c.IsSet(playLoc + ".playbook")
		playHasModule := c.IsSet(playLoc + ".module")
		if playHasPlaybook && playHasModule {
			es = append(es, errors.New("playbook and module can't be used together"))
			continue
		}
		if !playHasPlaybook && !playHasModule {
			es = append(es, errors.New("playbook or module must be set"))
			continue
		}
		if playHasModule {
			validPlaysCount++
			continue
		}

		value, _ := c.Get(playLoc + ".playbook")
		_, isArrayOfMaps = value.([]map[string]interface{}) // in Terraform 0.11.x types are concrete
		_, isArrayOfIfs = value.([]interface{})             // in Terraform 0.12+ types are wrapped in interface{}
		value, _ = c.Get(playLoc + ".playbook.#")
		numPlaybooks, ok := value.(int)
		if (!isArrayOfIfs && !isArrayOfMaps) || value == nil || !ok {
			es = append(es, errors.New("`plays.playbook` must be an array"))
			continue
		}

		playIsOK := true
		for playbookNo := 0; playbookNo < numPlaybooks; playbookNo++ {
			playbookLoc := fmt.Sprintf("%s.playbook.%d", playLoc, playbookNo)
			if !c.IsSet(playbookLoc + ".roles_path") {
				continue
			}

			// TODO investigate this:
			// If a value in roles_path is computed, an attempt to query the
			// array length using the `*.roles_path.#` returns a UUID string.
			// As a temporary workaround, lets use len() of the castedArray.
			value, _ := c.Get(playbookLoc + ".roles_path")
			rolesPathAsArray, rolesPathArrayIsOK := value.([]interface{})
			if !rolesPathArrayIsOK {
				es = append(es, errors.New("`plays.playbook.roles_path` must be an array"))
				playIsOK = false
				continue
			}

			for pathNo := 0; pathNo < len(rolesPathAsArray); pathNo++ {
				// Terraform 0.11.x attempts to interpolate everything ASAP, but
				// Terraform 0.12+ will lazily return an UUID for computed values.
				if c.IsComputed(fmt.Sprintf("%s.roles_path.%d", playbookLoc, pathNo)) {
					//ws = append(ws, "Cannot validate roles_path as it is computed")
					log.Printf("[WARN] Cannot validate roles_path as it is computed")
					continue
				}
				if value, ok = rolesPathAsArray[pathNo].(string); !ok {
					es = append(es, errors.New("`plays.playbook.roles_path` must be an array of strings"))
					playIsOK = false
					continue
				}
				wsDir, esDir := types.VfPathDirectory(value, "roles_path")
				ws = append(ws, wsDir...)
				es = append(es, esDir...)
				if esDir != nil {
					playIsOK = false
				}
			}
		}
		if playIsOK {
			validPlaysCount++
		}
	}

	if validPlaysCount == 0 {
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
