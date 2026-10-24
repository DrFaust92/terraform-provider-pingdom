package pingdom

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// fakePingdom is a minimal Pingdom API for a single HTTP check. It records the basic auth
// credentials it received so tests can check what was sent.
type fakePingdom struct {
	mu    sync.Mutex
	check url.Values
	auths []string
}

func (f *fakePingdom) RoundTrip(r *http.Request) (*http.Response, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	if r.Body != nil {
		raw, _ := io.ReadAll(r.Body)
		r.Body = io.NopCloser(strings.NewReader(string(raw)))
	}
	_ = r.ParseForm()

	path := strings.TrimPrefix(r.URL.Path, "/api/3.1")
	switch {
	case r.Method == http.MethodGet && path == "/checks":
		checks := []interface{}{}
		if f.check != nil {
			checks = append(checks, map[string]interface{}{"id": 1, "name": f.check.Get("name"), "hostname": f.check.Get("host"), "type": "http"})
		}
		return pingdomJSON(r, map[string]interface{}{"checks": checks}), nil
	case r.Method == http.MethodPost && path == "/checks":
		f.check = r.Form
		f.auths = append(f.auths, r.Form.Get("auth"))
		return pingdomJSON(r, map[string]interface{}{"check": map[string]interface{}{"id": 1, "name": f.check.Get("name")}}), nil
	case r.Method == http.MethodPut && path == "/checks/1" && f.check != nil:
		for k, v := range r.Form {
			f.check[k] = v
		}
		f.auths = append(f.auths, r.Form.Get("auth"))
		return pingdomJSON(r, map[string]interface{}{"message": "Modification of check was successful!"}), nil
	case r.Method == http.MethodDelete && path == "/checks/1":
		f.check = nil
		return pingdomJSON(r, map[string]interface{}{"message": "Deletion of check was successful!"}), nil
	case r.Method == http.MethodGet && path == "/checks/1" && f.check != nil:
		user, pass, _ := strings.Cut(f.check.Get("auth"), ":")
		atoi := func(k string) int { n, _ := strconv.Atoi(f.check.Get(k)); return n }
		return pingdomJSON(r, map[string]interface{}{"check": map[string]interface{}{
			"id":                       1,
			"name":                     f.check.Get("name"),
			"hostname":                 f.check.Get("host"),
			"resolution":               atoi("resolution"),
			"sendnotificationwhendown": atoi("sendnotificationwhendown"),
			"notifyagainevery":         atoi("notifyagainevery"),
			"notifywhenbackup":         f.check.Get("notifywhenbackup") == "true",
			"responsetime_threshold":   atoi("responsetime_threshold"),
			"status":                   "up",
			"type": map[string]interface{}{"http": map[string]interface{}{
				"url":                  f.check.Get("url"),
				"encryption":           f.check.Get("encryption") == "true",
				"port":                 atoi("port"),
				"username":             user,
				"password":             pass,
				"verify_certificate":   f.check.Get("verify_certificate") == "true",
				"ssl_down_days_before": atoi("ssl_down_days_before"),
			}},
		}}), nil
	}

	return pingdomJSON(r, map[string]interface{}{"error": map[string]interface{}{"statuscode": 404, "statusdesc": "Not Found", "errormessage": r.Method + " " + path}}), nil
}

func pingdomJSON(r *http.Request, v interface{}) *http.Response {
	body, _ := json.Marshal(v)
	status := http.StatusOK
	if strings.Contains(string(body), `"statuscode":404`) {
		status = http.StatusNotFound
	}
	return &http.Response{
		StatusCode: status,
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       io.NopCloser(strings.NewReader(string(body))),
		Request:    r,
	}
}

func (f *fakePingdom) testCheckAuth(want string) resource.TestCheckFunc {
	return func(*terraform.State) error {
		f.mu.Lock()
		defer f.mu.Unlock()
		if len(f.auths) == 0 || f.auths[len(f.auths)-1] != want {
			return fmt.Errorf("API received auth %q, want %q", f.auths, want)
		}
		return nil
	}
}

func testCheckSecretNotInState(secret string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			for k, v := range rs.Primary.Attributes {
				if strings.Contains(v, secret) {
					return fmt.Errorf("secret found in state: %s.%s", name, k)
				}
			}
		}
		return nil
	}
}

func TestWriteOnlyCheckPassword(t *testing.T) {
	fake := &fakePingdom{}
	original := http.DefaultTransport
	http.DefaultTransport = fake
	t.Cleanup(func() { http.DefaultTransport = original })
	t.Setenv("PINGDOM_API_TOKEN", "token")

	config := func(password, version string) string {
		return fmt.Sprintf(`
resource "pingdom_check" "test" {
  type                = "http"
  name                = "wo-test"
  host                = "example.com"
  resolution          = 5
  username            = "monitor"
  password_wo         = %q
  password_wo_version = %q
}
`, password, version)
	}

	resource.UnitTest(t, resource.TestCase{
		ProviderFactories: writeOnlyTestProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: config("s3cret", "1"),
				Check: resource.ComposeTestCheckFunc(
					fake.testCheckAuth("monitor:s3cret"),
					// The API returns the password, it must still not end up in state.
					testCheckSecretNotInState("s3cret"),
				),
			},
			{
				Config: config("rotated", "2"),
				Check: resource.ComposeTestCheckFunc(
					fake.testCheckAuth("monitor:rotated"),
					testCheckSecretNotInState("rotated"),
				),
			},
		},
	})
}

var writeOnlyTestProviderFactories = map[string]func() (*schema.Provider, error){
	"pingdom": func() (*schema.Provider, error) { return Provider(), nil },
}
