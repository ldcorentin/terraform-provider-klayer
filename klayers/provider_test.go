package klayers

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"testing"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"klayers": testAccProvider,
	}
}
