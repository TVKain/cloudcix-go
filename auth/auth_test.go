package auth

import (
	"testing"

	"github.com/TVKain/cloudcix-go/config"
)

func TestTokenManager(t *testing.T) {
	settings := &config.Settings{
		CLOUDCIX_API_URL:      "https://compute.api.cloudcix.com/",
		CLOUDCIX_API_V2_URL:   "https://membership.api.cloudcix.com/",
		CLOUDCIX_API_VERSION:  "5.0",
		CLOUDCIX_API_USERNAME: "test@example.com",
		CLOUDCIX_API_PASSWORD: "testpass",
		CLOUDCIX_API_KEY:      "testkey",
	}

	tokenManager := NewTokenManager(settings)

	if tokenManager == nil {
		t.Fatal("TokenManager should not be nil")
	}

	if tokenManager.settings != settings {
		t.Fatal("TokenManager should store settings")
	}

	if tokenManager.IsTokenValid() {
		t.Fatal("TokenManager should not have valid token initially")
	}
}

func TestSettingsValidation(t *testing.T) {
	// Test valid settings
	validSettings := &config.Settings{
		CLOUDCIX_API_USERNAME: "test@example.com",
		CLOUDCIX_API_PASSWORD: "testpass",
		CLOUDCIX_API_KEY:      "testkey",
		CLOUDCIX_REGION_ID:    1,
	}

	if err := validSettings.Validate(); err != nil {
		t.Fatalf("Valid settings should pass validation: %v", err)
	}

	// Test missing username
	invalidSettings := &config.Settings{
		CLOUDCIX_API_PASSWORD: "testpass",
		CLOUDCIX_API_KEY:      "testkey",
	}

	if err := invalidSettings.Validate(); err == nil {
		t.Fatal("Settings without username should fail validation")
	}

	// Test missing password
	invalidSettings2 := &config.Settings{
		CLOUDCIX_API_USERNAME: "test@example.com",
		CLOUDCIX_API_KEY:      "testkey",
		CLOUDCIX_REGION_ID:    1,
	}

	if err := invalidSettings2.Validate(); err == nil {
		t.Fatal("Settings without password should fail validation")
	}

	// Test missing API key
	invalidSettings3 := &config.Settings{
		CLOUDCIX_API_USERNAME: "test@example.com",
		CLOUDCIX_API_PASSWORD: "testpass",
		CLOUDCIX_REGION_ID:    1,
	}

	if err := invalidSettings3.Validate(); err == nil {
		t.Fatal("Settings without API key should fail validation")
	}

	// Test missing region ID
	invalidSettings4 := &config.Settings{
		CLOUDCIX_API_USERNAME: "test@example.com",
		CLOUDCIX_API_PASSWORD: "testpass",
		CLOUDCIX_API_KEY:      "testkey",
		// CLOUDCIX_REGION_ID not set (defaults to 0)
	}

	if err := invalidSettings4.Validate(); err == nil {
		t.Fatal("Settings without region ID should fail validation")
	}
}