package klayers

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceVersionsKlayersRead(d *schema.ResourceData, m interface{}) error {

	name := d.Get("name").(string)
	region := d.Get("region").(string)
	python_version := d.Get("python_version").(string)
	
	arns := getKlayerArns(name, region, python_version).([]map[string]string)

	var lArns []string
	for _, layer := range arns {
		latestArn, _ := layer["arn"]
		lArns = append(lArns, latestArn)
    }

	if len(lArns) == 0 {
		return fmt.Errorf("Package %s not found for python version %s and region %s", name, python_version, region)
	}

	d.Set("region", region)
	d.Set("api", baseLatestURL)
	d.Set("arns", lArns)
	d.Set("name", name)

	d.SetId(fmt.Sprintf("%s-%s", name, region))
	return nil
}

func dataVersionsKlayersPackage() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVersionsKlayersRead,
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
			"arns": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{ 
					Type: schema.TypeString,
				},
			},
		},
	}
}