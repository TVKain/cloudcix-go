// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudcix_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/TVKain/cloudcix-go"
	"github.com/TVKain/cloudcix-go/internal/testutil"
	"github.com/TVKain/cloudcix-go/option"
)

func TestNetworkRouterNewWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudcix.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Network.Routers.New(context.TODO(), cloudcix.NetworkRouterNewParams{
		ProjectID: 1,
		Metadata: cloudcix.NetworkRouterNewParamsMetadata{
			Destination: cloudcix.String("destination"),
			Nat:         cloudcix.Bool(true),
			Nexthop:     cloudcix.String("nexthop"),
		},
		Name: cloudcix.String("Public Website Router"),
		Networks: []cloudcix.NetworkRouterNewParamsNetwork{{
			Ipv4: cloudcix.String("10.0.1.0/24"),
			Name: cloudcix.String("web-tier"),
		}, {
			Ipv4: cloudcix.String("10.0.2.0/24"),
			Name: cloudcix.String("app-tier"),
		}, {
			Ipv4: cloudcix.String("10.0.3.0/24"),
			Name: cloudcix.String("db-tier"),
		}},
		Type: cloudcix.String("router"),
	})
	if err != nil {
		var apierr *cloudcix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNetworkRouterGet(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudcix.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Network.Routers.Get(context.TODO(), 0)
	if err != nil {
		var apierr *cloudcix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNetworkRouterUpdateWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudcix.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Network.Routers.Update(
		context.TODO(),
		0,
		cloudcix.NetworkRouterUpdateParams{
			NetworkRouterUpdate: cloudcix.NetworkRouterUpdateParam{
				State: "state",
				Metadata: cloudcix.NetworkRouterUpdateMetadataParam{
					Destination: cloudcix.String("destination"),
					Nat:         cloudcix.Bool(true),
					Nexthop:     cloudcix.String("nexthop"),
				},
				Name: cloudcix.String("name"),
				Networks: []cloudcix.NetworkRouterUpdateNetworkParam{{
					Ipv4: cloudcix.String("10.0.1.0/24"),
					Ipv6: cloudcix.String("fd00::/64"),
					Name: cloudcix.String("existing-network"),
					Vlan: cloudcix.Int(100),
				}, {
					Ipv4: cloudcix.String("10.0.50.0/24"),
					Ipv6: cloudcix.String("ipv6"),
					Name: cloudcix.String("new-database-network"),
					Vlan: cloudcix.Int(0),
				}},
			},
		},
	)
	if err != nil {
		var apierr *cloudcix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNetworkRouterListWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := cloudcix.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Network.Routers.List(context.TODO(), cloudcix.NetworkRouterListParams{
		Exclude: map[string]interface{}{},
		Limit:   cloudcix.Int(0),
		Order:   cloudcix.String("order"),
		Page:    cloudcix.Int(0),
		Search:  map[string]interface{}{},
	})
	if err != nil {
		var apierr *cloudcix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
