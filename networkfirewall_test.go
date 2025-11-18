// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudcix_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/cloudcix-go"
	"github.com/stainless-sdks/cloudcix-go/internal/testutil"
	"github.com/stainless-sdks/cloudcix-go/option"
)

func TestNetworkFirewallNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Network.Firewalls.New(context.TODO(), cloudcix.NetworkFirewallNewParams{
		ProjectID: 1,
		Name:      cloudcix.String("Allow Traffic from Ireland"),
		Rules: []cloudcix.NetworkFirewallNewParamsRule{{
			Allow:       cloudcix.Bool(true),
			Description: cloudcix.String("description"),
			Destination: cloudcix.String("destination"),
			GroupName:   cloudcix.String("@IE_v4"),
			Inbound:     cloudcix.Bool(true),
			Port:        cloudcix.String("port"),
			Protocol:    cloudcix.String("protocol"),
			Source:      cloudcix.String("source"),
		}},
		Type: cloudcix.String("geo"),
	})
	if err != nil {
		var apierr *cloudcix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNetworkFirewallGet(t *testing.T) {
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
	_, err := client.Network.Firewalls.Get(context.TODO(), 0)
	if err != nil {
		var apierr *cloudcix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNetworkFirewallUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Network.Firewalls.Update(
		context.TODO(),
		0,
		cloudcix.NetworkFirewallUpdateParams{
			NetworkFirewallUpdate: cloudcix.NetworkFirewallUpdateParam{
				State: "delete",
				Name:  cloudcix.String("Public Website Firewall"),
				Rules: []cloudcix.NetworkFirewallUpdateRuleParam{{
					Allow:       cloudcix.Bool(true),
					Description: cloudcix.String("description"),
					Destination: cloudcix.String("destination"),
					GroupName:   cloudcix.String("group_name"),
					Inbound:     cloudcix.Bool(true),
					Port:        cloudcix.String("port"),
					Protocol:    cloudcix.String("protocol"),
					Source:      cloudcix.String("source"),
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

func TestNetworkFirewallListWithOptionalParams(t *testing.T) {
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
	_, err := client.Network.Firewalls.List(context.TODO(), cloudcix.NetworkFirewallListParams{
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
