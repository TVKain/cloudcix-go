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

func TestComputeGPUGet(t *testing.T) {
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
	_, err := client.Compute.GPUs.Get(context.TODO(), 0)
	if err != nil {
		var apierr *cloudcix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestComputeGPUUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Compute.GPUs.Update(
		context.TODO(),
		0,
		cloudcix.ComputeGPUUpdateParams{
			ComputeGPUUpdate: cloudcix.ComputeGPUUpdateParam{
				Name:  cloudcix.String("HighPerformance-GPU"),
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

func TestComputeGPUListWithOptionalParams(t *testing.T) {
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
	_, err := client.Compute.GPUs.List(context.TODO(), cloudcix.ComputeGPUListParams{
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

func TestComputeGPUAttachWithOptionalParams(t *testing.T) {
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
	_, err := client.Compute.GPUs.Attach(context.TODO(), cloudcix.ComputeGPUAttachParams{
		InstanceID: 456,
		ProjectID:  123,
		Specs: []cloudcix.ComputeGPUAttachParamsSpec{{
			SKUName: cloudcix.String("A100_GPU"),
		}},
		Name: cloudcix.String("HighPerformance-GPU"),
	})
	if err != nil {
		var apierr *cloudcix.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
