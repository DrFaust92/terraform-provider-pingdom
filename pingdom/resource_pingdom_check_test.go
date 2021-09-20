package pingdom

import (
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/russellcardullo/go-pingdom/pingdom"
)

func TestAccPingdomCheck_basic(t *testing.T) {
	var check pingdom.CheckResponse

	rName := acctest.RandomWithPrefix("tf-acc-test")
	resourceName := "pingdom_check.test"
	resource.Test(t, resource.TestCase{
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckPingdomCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPingdomBasicConfig(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckPingdomCheckExists(resourceName, &check),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "type", "http"),
					resource.TestCheckResourceAttr(resourceName, "host", "example.com"),
					resource.TestCheckResourceAttr(resourceName, "url", "/"),
					resource.TestCheckResourceAttr(resourceName, "verify_certificate", "false"),
					resource.TestCheckResourceAttr(resourceName, "ssl_down_days_before", "0"),
					resource.TestCheckResourceAttr(resourceName, "sendnotificationwhendown", "2"),
					resource.TestCheckResourceAttr(resourceName, "resolution", "5"),
					resource.TestCheckResourceAttr(resourceName, "notifywhenbackup", "false"),
					resource.TestCheckResourceAttr(resourceName, "notifyagainevery", "0"),
					resource.TestCheckResourceAttr(resourceName, "encryption", "false"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckPingdomCheckExists(n string, checkId *pingdom.CheckResponse) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No key ID is set")
		}

		client := testAccProvider.Meta().(*pingdom.Client)
		keys, err := client.Checks.List()
		if err != nil {
			return err
		}

		for _, key := range keys {
			if fmt.Sprint(key.ID) == rs.Primary.ID {
				*checkId = key
				break
			}
		}
		return nil
	}
}

func testAccCheckPingdomCheckDestroy(s *terraform.State) error {
	client := testAccProvider.Meta().(*pingdom.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "pingdom_check" {
			continue
		}

		keys, err := client.Checks.List()
		if err != nil {
			return err
		}
		if err == nil {
			for _, key := range keys {
				if fmt.Sprint(key.ID) == rs.Primary.ID {
					return errors.New("Key still exists")
				}
			}
		}
		return nil
	}
	return nil
}

func testAccPingdomBasicConfig(rName string) string {
	return fmt.Sprintf(`
resource "pingdom_check" "test" {
  name = %[1]q
  host = "example.com"
  type = "http"
}
`, rName)
}
