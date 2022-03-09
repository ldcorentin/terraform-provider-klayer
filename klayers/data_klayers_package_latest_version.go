// This file will have the resource function declaration and definition like create, delete, update etc, it also gets the input params required to create resources.
package klayers

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Create
// An empty interface may hold values of any type. (Every type implements at least zero methods.)
// Empty interfaces are used by code that handles values of unknown type.
// m is the interface{} returned from the ConfigureFunc in provider.go
// In the example below, interface{} is being used to pass us our Provider object without expressing anything about what type it is.
// This allows a fixed function signature in the Terraform library, while letting us runtime cast to whatever type our Provider object actually is.
// In our case then, meta interface{} refers to the ProviderStruct defined elsewhere in our source code.
// This gives us access to all the properties defined on our Provider from within the lifecycle handling for our data object.
func dataSourceKlayersRead(d *schema.ResourceData, m interface{}) error {

	name := d.Get("name").(string)
	region := d.Get("region").(string)
	python_version := d.Get("python_version").(string)
	
	arns := getKlayerArns(name, region, python_version).([]map[string]string)

	var latestArn string
	for _, layer := range arns {
		if value, _ := layer["deployStatus"]; value == "latest" {
			latestArn, _ = layer["arn"]
			break
		}
    }

	if latestArn == "" {
		return fmt.Errorf("Package %s not found for python version %s and region %s", name, python_version, region)
	}

	d.Set("region", region)
	d.Set("api", baseLatestURL)
	d.Set("arn", latestArn)
	d.Set("name", name)

	d.SetId(fmt.Sprintf("%s-%s", name, region))
	return nil
}

func dataLatestKlayersPackage() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceKlayersRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"region": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"python_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default: "3.9",
			},
			"api": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"arn": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}