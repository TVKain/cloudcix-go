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

func TestNetworkIPGroupNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Network.IPGroups.New(context.TODO(), cloudcix.NetworkIPGroupNewParams{
		Cidrs:   []string{"91.103.3.0/24", "90.103.2.36", "185.94.188.0/24"},
		Name:    "office-networks",
		Version: cloudcix.Int(4),
	})
	if err != nil {
		var apierr *cloudcix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNetworkIPGroupGet(t *testing.T) {
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
	_, err := client.Network.IPGroups.Get(context.TODO(), 0)
	if err != nil {
		var apierr *cloudcix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNetworkIPGroupUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Network.IPGroups.Update(
		context.TODO(),
		0,
		cloudcix.NetworkIPGroupUpdateParams{
			NetworkIPGroupUpdate: cloudcix.NetworkIPGroupUpdateParam{
				Cidrs:   []string{"91.103.3.0/24", "90.103.2.36", "185.94.188.0/24"},
				Name:    "office-networks",
				Version: cloudcix.Int(4),
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

func TestNetworkIPGroupListWithOptionalParams(t *testing.T) {
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
	_, err := client.Network.IPGroups.List(context.TODO(), cloudcix.NetworkIPGroupListParams{
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

func TestNetworkIPGroupDelete(t *testing.T) {
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
	err := client.Network.IPGroups.Delete(context.TODO(), 0)
	if err != nil {
		var apierr *cloudcix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
