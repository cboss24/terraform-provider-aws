package aws

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAwsBatchComputeEnvironment() *schema.Resource {
	return &schema.Resource{
		Create: resourceAwsBatchComputeEnvironmentCreate,
		Read:   resourceAwsBatchComputeEnvironmentRead,
		Delete: resourceAwsBatchComputeEnvironmentDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceAwsBatchComputeEnvironmentCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceAwsBatchComputeEnvironmentRead(d, meta)
}

func resourceAwsBatchComputeEnvironmentRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceAwsBatchComputeEnvironmentDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
