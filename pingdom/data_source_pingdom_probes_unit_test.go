package pingdom

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const testProbesResponse = `{"probes": [
  {"id": 32, "country": "United States", "city": "Los Angeles", "name": "Los Angeles, CA", "active": true,
   "hostname": "s410.pingdom.com", "ip": "204.152.200.42", "countryiso": "US", "ipv6": "2607:fcd0:100:8d00::410", "region": "NA"},
  {"id": 184, "country": "Brazil", "city": "Sao Paulo", "name": "Sao Paulo 2, Brazil", "active": true,
   "hostname": "s4028.pingdom.com", "ip": "52.67.148.55", "countryiso": "BR", "ipv6": "", "region": "LATAM"}
]}`

type fakeProbesAPI struct {
	mu    sync.Mutex
	query []string
}

func (f *fakeProbesAPI) RoundTrip(r *http.Request) (*http.Response, error) {
	f.mu.Lock()
	f.query = append(f.query, r.URL.RawQuery)
	f.mu.Unlock()

	body := `{"error": {"statuscode": 404, "statusdesc": "Not Found", "errormessage": "not found"}}`
	status := http.StatusNotFound
	if strings.HasSuffix(r.URL.Path, "/probes") {
		body, status = testProbesResponse, http.StatusOK
	}

	return &http.Response{
		StatusCode: status,
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       io.NopCloser(strings.NewReader(body)),
		Request:    r,
	}, nil
}

func TestDataSourcePingdomProbes(t *testing.T) {
	fake := &fakeProbesAPI{}
	original := http.DefaultTransport
	http.DefaultTransport = fake
	t.Cleanup(func() { http.DefaultTransport = original })
	t.Setenv("PINGDOM_API_TOKEN", "token")

	resource.UnitTest(t, resource.TestCase{
		ProviderFactories: map[string]func() (*schema.Provider, error){
			"pingdom": func() (*schema.Provider, error) { return Provider(), nil },
		},
		Steps: []resource.TestStep{
			{
				Config: `
data "pingdom_probes" "all" {
  only_active = true
}

data "pingdom_probes" "na" {
  region = "NA"
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.pingdom_probes.all", "probes.#", "2"),
					resource.TestCheckResourceAttr("data.pingdom_probes.all", "probes.1.country_iso", "BR"),
					resource.TestCheckResourceAttr("data.pingdom_probes.all", "ipv4_addresses.#", "2"),
					resource.TestCheckResourceAttr("data.pingdom_probes.all", "ipv4_addresses.1", "52.67.148.55"),
					// Probes without an IPv6 address are skipped.
					resource.TestCheckResourceAttr("data.pingdom_probes.all", "ipv6_addresses.#", "1"),
					resource.TestCheckResourceAttr("data.pingdom_probes.na", "probes.#", "1"),
					resource.TestCheckResourceAttr("data.pingdom_probes.na", "probes.0.hostname", "s410.pingdom.com"),
					resource.TestCheckResourceAttr("data.pingdom_probes.na", "ipv4_addresses.0", "204.152.200.42"),
					func(*terraform.State) error {
						fake.mu.Lock()
						defer fake.mu.Unlock()
						for _, q := range fake.query {
							if strings.Contains(q, "onlyactive=true") {
								return nil
							}
						}
						return fmt.Errorf("onlyactive=true not sent, queries: %q", fake.query)
					},
				),
			},
		},
	})
}
