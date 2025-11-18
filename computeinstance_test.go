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

func TestComputeInstanceNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Compute.Instances.New(context.TODO(), cloudcix.ComputeInstanceNewParams{
		Metadata: cloudcix.ComputeInstanceNewParamsMetadata{
			DNS:          cloudcix.String("dns"),
			InstanceType: cloudcix.String("instance_type"),
			Userdata:     cloudcix.String("userdata"),
		},
		ProjectID: 1,
		Specs: []cloudcix.ComputeInstanceNewParamsSpec{{
			Quantity: cloudcix.Int(8),
			SKUName:  cloudcix.String("RAM_001"),
		}, {
			Quantity: cloudcix.Int(4),
			SKUName:  cloudcix.String("vCPU_002"),
		}, {
			Quantity: cloudcix.Int(100),
			SKUName:  cloudcix.String("SSD_001"),
		}, {
			Quantity: cloudcix.Int(1),
			SKUName:  cloudcix.String("MSServer2022"),
		}},
		GracePeriod: cloudcix.Int(72),
		Interfaces: []cloudcix.ComputeInstanceNewParamsInterface{{
			Gateway: cloudcix.Bool(true),
			Ipv4Addresses: []cloudcix.ComputeInstanceNewParamsInterfaceIpv4Address{{
				Address: cloudcix.String("10.0.0.10"),
				Nat:     cloudcix.Bool(true),
			}},
			Ipv6Addresses: []cloudcix.ComputeInstanceNewParamsInterfaceIpv6Address{{
				Address: cloudcix.String("2a02:2078:9:1234::20"),
			}},
		}},
		Name: cloudcix.String("windows-server"),
		Type: cloudcix.String("hyperv"),
	})
	if err != nil {
		var apierr *cloudcix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestComputeInstanceGet(t *testing.T) {
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
	_, err := client.Compute.Instances.Get(context.TODO(), 0)
	if err != nil {
		var apierr *cloudcix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestComputeInstanceUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Compute.Instances.Update(
		context.TODO(),
		0,
		cloudcix.ComputeInstanceUpdateParams{
			ComputeInstanceUpdate: cloudcix.ComputeInstanceUpdateParam{
				Metadata: cloudcix.ComputeInstanceUpdateMetadataParam{
					DNS:          cloudcix.String("dns"),
					InstanceType: cloudcix.String("instance_type"),
					Userdata:     cloudcix.String("userdata"),
				},
				Specs: []cloudcix.ComputeInstanceUpdateSpecParam{{
					Quantity: cloudcix.Int(0),
					SKUName:  cloudcix.String("sku_name"),
				}},
				GracePeriod: cloudcix.Int(0),
				Interfaces: []cloudcix.ComputeInstanceUpdateInterfaceParam{{
					Gateway: cloudcix.Bool(true),
					Ipv4Addresses: []cloudcix.ComputeInstanceUpdateInterfaceIpv4AddressParam{{
						Address: cloudcix.String("address"),
						Nat:     cloudcix.Bool(true),
					}},
					Ipv6Addresses: []cloudcix.ComputeInstanceUpdateInterfaceIpv6AddressParam{{
						Address: cloudcix.String("address"),
					}},
				}},
				Name:  cloudcix.String("name"),
				State: cloudcix.String("delete"),
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

func TestComputeInstanceListWithOptionalParams(t *testing.T) {
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
	_, err := client.Compute.Instances.List(context.TODO(), cloudcix.ComputeInstanceListParams{
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
