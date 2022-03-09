// Provider.go will have the resource server function calls.
package klayers

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Create the provider
func Provider() *schema.Provider {
	return &schema.Provider{
		// Declare the terraform resource/data name with the associated go function
		DataSourcesMap: map[string]*schema.Resource{
			"klayers_package_latest_version": dataLatestKlayersPackage(),
			"klayers_package_all_versions": dataVersionsKlayersPackage(),
		},
	}
}
