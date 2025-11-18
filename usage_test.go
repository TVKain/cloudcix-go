// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudcix_test

import (
	"context"
	"os"
	"testing"

	"github.com/TVKain/cloudcix-go"
	"github.com/TVKain/cloudcix-go/internal/testutil"
	"github.com/TVKain/cloudcix-go/option"
)

func TestUsage(t *testing.T) {
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
	backups, err := client.Compute.Backups.List(context.TODO(), cloudcix.ComputeBackupListParams{})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", backups.Metadata)
}
