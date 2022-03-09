package klayers

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccItem_BasicLatest(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckItemBasicLatest(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.klayers_package_latest_version.pandas", "name", "pandas"),
					resource.TestCheckResourceAttr(
						"data.klayers_package_latest_version.pandas", "region", "eu-west-1"),
					resource.TestCheckResourceAttr(
						"data.klayers_package_latest_version.pandas", "python_version", "3.9"),
				),
			},
		},
	})
}

func testAccCheckItemBasicLatest() string {
	return fmt.Sprintf(`
		data "klayers_package_latest_version" "pandas" {
  			name   = "pandas"
  			region = "eu-west-1"
		}
	`)
}
