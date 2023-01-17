// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccFirebaseAndroidApp_firebaseAndroidAppBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        getTestOrgFromEnv(t),
		"project_id":    getTestProjectFromEnv(),
		"package_name":  "android.package.app" + randString(t, 4),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckFirebaseAndroidAppDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccFirebaseAndroidApp_firebaseAndroidAppBasicExample(context),
			},
			{
				ResourceName:            "google_firebase_android_app.basic",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"deletion_policy", "project"},
			},
		},
	})
}

func testAccFirebaseAndroidApp_firebaseAndroidAppBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_firebase_android_app" "basic" {
  provider = google-beta
  project = "%{project_id}"
  display_name = "Display Name Basic%{random_suffix}"
  package_name = "%{package_name}"
  sha1_hashes = ["2145bdf698b8715039bd0e83f2069bed435ac21c"]
  sha256_hashes = ["2145bdf698b8715039bd0e83f2069bed435ac21ca1b2c3d4e5f6123456789abc"]
}
`, context)
}

func testAccCheckFirebaseAndroidAppDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_firebase_android_app" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{FirebaseBasePath}}{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("FirebaseAndroidApp still exists at %s", url)
			}
		}

		return nil
	}
}
