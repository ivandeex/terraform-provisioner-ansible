package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
	"github.com/radekg/terraform-provisioner-ansible/dbg"
)

func main() {
	dbg.StartLog()
	defer dbg.StopLog()
	plugin.Serve(&plugin.ServeOpts{
		ProvisionerFunc: func() terraform.ResourceProvisioner {
			return Provisioner()
		},
	})
}
