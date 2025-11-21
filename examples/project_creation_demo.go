package main

import (
	"context"
	"fmt"
	"log"

	"github.com/TVKain/cloudcix-go"
)

func main() {
	// Load credentials and settings from my_settings file
	client, err := cloudcix.NewClientFromFile("my_settings")
	if err != nil {
		log.Printf("Failed to load from settings file: %v", err)
		log.Println("Make sure you have a 'my_settings' file with your CloudCIX credentials.")
		log.Println("Example my_settings file:")
		log.Println("")
		log.Println("# Base API URL")
		log.Println(`CLOUDCIX_API_URL="https://api.cloudcix.com/"`)
		log.Println("")
		log.Println("# Default region ID")
		log.Println(`CLOUDCIX_REGION_ID=12345`)
		log.Println("")
		log.Println("# Your credentials")
		log.Println(`CLOUDCIX_API_USERNAME="your_email@example.com"`)
		log.Println(`CLOUDCIX_API_PASSWORD="your_password"`)
		log.Println(`CLOUDCIX_API_KEY="your_api_key"`)
		return
	}

	// All operations will automatically handle authentication and token renewal
	ctx := context.Background()

	fmt.Println("=== CloudCIX Go SDK - Project Creation Demo ===")
	
	// First, let's list existing projects to see what we have
	fmt.Println("\n1. Listing existing projects...")
	projects, err := client.Project.List(ctx, cloudcix.ProjectListParams{})
	if err != nil {
		log.Printf("Failed to list projects: %v", err)
	} else {
		fmt.Printf("Found %d existing projects:\n", len(projects.Content))
		for _, project := range projects.Content {
			fmt.Printf("  - ID: %d, Name: %s, Region: %d\n", project.ID, project.Name, project.RegionID)
		}
	}

	// Create a new project
	fmt.Println("\n2. Creating a new project...")
	
	// Get region ID from settings - this must be configured
	if client.GetSettings() == nil || client.GetSettings().CLOUDCIX_REGION_ID == 0 {
		log.Println("Error: CLOUDCIX_REGION_ID is not configured in my_settings file.")
		log.Println("Please add your region ID to the my_settings file:")
		log.Println("CLOUDCIX_REGION_ID=your_region_id")
		return
	}
	
	regionID := client.GetSettings().CLOUDCIX_REGION_ID
	fmt.Printf("Using region ID: %d\n", regionID)
	
	newProject, err := client.Project.New(ctx, cloudcix.ProjectNewParams{
		Name:     "go-sdk-demo-project-2",
		RegionID: int64(regionID),
		Note:     cloudcix.String("Demo project created using CloudCIX Go SDK"),
	})
	if err != nil {
		log.Printf("Failed to create project: %v", err)
		log.Println("Note: Make sure the region ID is valid and you have permissions to create projects.")
	} else {
		fmt.Printf("✅ Successfully created project!\n")
		fmt.Printf("  - Project ID: %d\n", newProject.Content.ID)
		fmt.Printf("  - Name: %s\n", newProject.Content.Name)
		fmt.Printf("  - Region ID: %d\n", newProject.Content.RegionID)
		fmt.Printf("  - Note: %s\n", newProject.Content.Note)
		fmt.Printf("  - Created: %s\n", newProject.Content.Created)
		fmt.Printf("  - Manager ID: %d\n", newProject.Content.ManagerID)
		fmt.Printf("  - Address ID: %d\n", newProject.Content.AddressID)

		// Get the specific project we just created to verify
		fmt.Println("\n3. Retrieving the created project...")
		retrievedProject, err := client.Project.Get(ctx, newProject.Content.ID)
		if err != nil {
			log.Printf("Failed to retrieve project: %v", err)
		} else {
			fmt.Printf("✅ Retrieved project details:\n")
			fmt.Printf("  - ID: %d\n", retrievedProject.Content.ID)
			fmt.Printf("  - Name: %s\n", retrievedProject.Content.Name)
			fmt.Printf("  - Closed: %v\n", retrievedProject.Content.Closed)
			fmt.Printf("  - Region ID: %d\n", retrievedProject.Content.RegionID)
		}
	}

	fmt.Println("\n=== Demo completed! ===")
	fmt.Println("\nThis demo showed:")
	fmt.Println("✓ Loading credentials from my_settings file")
	fmt.Println("✓ Automatic dual-API management") 
	fmt.Println("✓ Listing existing projects")
	fmt.Println("✓ Creating a new project")
	fmt.Println("✓ Retrieving project details")
	fmt.Println("\nThe SDK automatically handled:")
	fmt.Println("✓ Authentication via Membership API (https://membership.api.cloudcix.com/)")
	fmt.Println("✓ Token management and renewal") 
	fmt.Println("✓ Project operations via Compute API (https://compute.api.cloudcix.com/)")
	fmt.Println("✓ Error handling and retries")
}