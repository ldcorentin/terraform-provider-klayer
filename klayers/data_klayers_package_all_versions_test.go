package klayers

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccItem_BasicVersions(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckItemBasicVersions(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.klayers_package_all_versions.requests", "name", "requests"),
					resource.TestCheckResourceAttr(
						"data.klayers_package_all_versions.requests", "region", "eu-west-1"),
					resource.TestCheckResourceAttr(
						"data.klayers_package_all_versions.requests", "python_version", "3.8"),
				),
			},
		},
	})
}

func testAccCheckItemBasicVersions() string {
	return fmt.Sprintf(`
		data "klayers_package_all_versions" "requests" {
  			name   			= "requests"
  			region 			= "eu-west-1"
			python_version  = "3.8"
		}
	`)
}
