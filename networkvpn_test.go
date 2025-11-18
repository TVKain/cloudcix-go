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

func TestNetworkVpnNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Network.Vpns.New(context.TODO(), cloudcix.NetworkVpnNewParams{
		ProjectID: 1,
		Metadata: cloudcix.NetworkVpnNewParamsMetadata{
			ChildSas: []cloudcix.NetworkVpnNewParamsMetadataChildSa{{
				LocalTs:  cloudcix.String("10.0.0.0/24"),
				RemoteTs: cloudcix.String("172.16.10.0/27"),
			}},
			IkeAuthentication:   cloudcix.String("SHA256"),
			IkeDhGroups:         cloudcix.String("Group 24"),
			IkeEncryption:       cloudcix.String("256 bit AES-CBC"),
			IkeGatewayType:      cloudcix.String("hostname"),
			IkeGatewayValue:     cloudcix.String("hostname.example.com"),
			IkeLifetime:         cloudcix.Int(0),
			IkePreSharedKey:     cloudcix.String("R4nd0mKey!"),
			IkeVersion:          cloudcix.String("v2-only"),
			IpsecAuthentication: cloudcix.String("SHA256"),
			IpsecEncryption:     cloudcix.String("AES 256 GCM"),
			IpsecEstablishTime:  cloudcix.String("On Traffic"),
			IpsecLifetime:       cloudcix.Int(3000),
			IpsecPfsGroups:      cloudcix.String("Group 24"),
			TrafficSelector:     cloudcix.Bool(true),
		},
		Name: cloudcix.String("Cork Office to Dublin Office"),
		Type: cloudcix.String("site-to-site"),
	})
	if err != nil {
		var apierr *cloudcix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNetworkVpnGet(t *testing.T) {
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
	_, err := client.Network.Vpns.Get(context.TODO(), 0)
	if err != nil {
		var apierr *cloudcix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNetworkVpnUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Network.Vpns.Update(
		context.TODO(),
		0,
		cloudcix.NetworkVpnUpdateParams{
			NetworkVpnUpdate: cloudcix.NetworkVpnUpdateParam{
				Metadata: cloudcix.NetworkVpnUpdateMetadataParam{
					ChildSas: []cloudcix.NetworkVpnUpdateMetadataChildSaParam{{
						LocalTs:  cloudcix.String("local_ts"),
						RemoteTs: cloudcix.String("remote_ts"),
					}},
					IkeAuthentication:   cloudcix.String("ike_authentication"),
					IkeDhGroups:         cloudcix.String("ike_dh_groups"),
					IkeEncryption:       cloudcix.String("ike_encryption"),
					IkeGatewayType:      cloudcix.String("ike_gateway_type"),
					IkeGatewayValue:     cloudcix.String("ike_gateway_value"),
					IkeLifetime:         cloudcix.Int(0),
					IkePreSharedKey:     cloudcix.String("ike_pre_shared_key"),
					IkeVersion:          cloudcix.String("ike_version"),
					IpsecAuthentication: cloudcix.String("ipsec_authentication"),
					IpsecEncryption:     cloudcix.String("ipsec_encryption"),
					IpsecEstablishTime:  cloudcix.String("ipsec_establish_time"),
					IpsecLifetime:       cloudcix.Int(0),
					IpsecPfsGroups:      cloudcix.String("ipsec_pfs_groups"),
					TrafficSelector:     cloudcix.Bool(true),
				},
				Name:  cloudcix.String("Cork Office to Dublin Office"),
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

func TestNetworkVpnListWithOptionalParams(t *testing.T) {
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
	_, err := client.Network.Vpns.List(context.TODO(), cloudcix.NetworkVpnListParams{
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
