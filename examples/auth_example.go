//go:build ignore
// +build ignore

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/TVKain/cloudcix-go"
	"github.com/TVKain/cloudcix-go/config"
)

func main() {
	// Example 1: Auto-load from environment variables
	// Set these environment variables:
	// CLOUDCIX_API_USERNAME=your_email@example.com
	// CLOUDCIX_API_PASSWORD=your_password
	// CLOUDCIX_API_KEY=your_api_key
	client := cloudcix.NewClient()

	// Example 2: Require authentication (will error if credentials missing)
	clientWithRequiredAuth, authErr := cloudcix.NewClientWithAuth()
	if authErr != nil {
		log.Printf("Authentication required but not configured: %v", authErr)
		// This is expected if you don't have credentials set up
	}

	// Example 3: Load from settings file
	clientFromFile, fileErr := cloudcix.NewClientFromFile("my_settings")
	if fileErr != nil {
		log.Printf("Failed to load from file: %v", fileErr)
	}

	// Example 4: Manual settings
	settings := &config.Settings{
		CLOUDCIX_API_URL:      "https://api.cloudcix.com/",
		CLOUDCIX_API_VERSION:  "5.0",
		CLOUDCIX_API_USERNAME: "your_email@example.com",
		CLOUDCIX_API_PASSWORD: "your_password",
		CLOUDCIX_API_KEY:      "your_api_key",
	}
	clientWithSettings := cloudcix.NewClientWithSettings(settings)

	// All these clients will automatically handle authentication and token renewal
	ctx := context.Background()

	// Make requests - authentication is handled automatically
	// The client will:
	// 1. Automatically get a token on first request
	// 2. Refresh tokens when they expire (every 2 hours)
	// 3. Retry failed requests with fresh tokens
	// 4. Handle all authentication errors transparently

	fmt.Println("Testing compute backups...")
	backups, err := client.Compute.Backups.List(ctx, cloudcix.ComputeBackupListParams{})
	if err != nil {
		log.Printf("Failed to get backups: %v", err)
	} else {
		fmt.Printf("Found %d backups\n", len(backups.Content))
	}

	fmt.Println("Testing compute instances with auth-required client...")
	if authErr == nil {  // If clientWithRequiredAuth was created successfully
		instances, err := clientWithRequiredAuth.Compute.Instances.List(ctx, cloudcix.ComputeInstanceListParams{})
		if err != nil {
			log.Printf("Failed to get instances: %v", err)
		} else {
			fmt.Printf("Found %d instances\n", len(instances.Content))
		}
	}

	fmt.Println("Testing with settings file client...")
	if fileErr == nil {
		backups2, err := clientFromFile.Compute.Backups.List(ctx, cloudcix.ComputeBackupListParams{})
		if err != nil {
			log.Printf("Failed to get backups from file client: %v", err)
		} else {
			fmt.Printf("Found %d backups with file client\n", len(backups2.Content))
		}
	}

	// You can also manually refresh if needed
	fmt.Println("Manually refreshing token...")
	err = clientWithSettings.RefreshToken(ctx)
	if err != nil {
		log.Printf("Failed to refresh token: %v", err)
	} else {
		fmt.Println("Token refreshed successfully!")
	}

	fmt.Println("Authentication demo completed successfully!")
}