# CloudCIX Go SDK Examples

This directory contains examples demonstrating how to use the CloudCIX Go SDK with automatic authentication.

## Examples

### 1. **Project Creation Demo** (`project_creation_demo.go`)

Complete example showing:

- Loading credentials from settings file  
- Listing existing projects
- Creating new projects
- Retrieving project details

### 2. **Authentication Examples** (`auth_example.go`)

Demonstrates different authentication methods:

- Environment variable loading
- Settings file loading  
- Manual credential configuration
- Multiple client creation patterns

### 3. **Authentication Test** (`auth_verification.go`)

Simple verification script to test:

- Settings file loading
- Token authentication
- API connectivity

## Setup

1. **Configure your credentials** in the `my_settings` file:

```env
CLOUDCIX_API_URL="https://api.cloudcix.com/"
CLOUDCIX_API_VERSION="5.0"
CLOUDCIX_REGION_ID=12345
CLOUDCIX_API_USERNAME="your_email@example.com"
CLOUDCIX_API_PASSWORD="your_password"
CLOUDCIX_API_KEY="your_api_key"
```

1. **Run any example**:

```bash
go run project_creation_demo.go
go run auth_example.go
go run auth_verification.go
```

## What the Demo Does

1. **Loads credentials** from `my_settings` file
2. **Lists existing projects** in your CloudCIX account
3. **Creates a new project** named "go-sdk-demo-project"
4. **Retrieves the created project** to verify it was created successfully

## Authentication Features Demonstrated

✅ **Automatic credential loading** from settings file
✅ **Automatic token management** - gets and refreshes tokens as needed
✅ **Error handling** - graceful handling of authentication and API errors
✅ **Transparent operation** - no manual token handling required

## Project Creation Parameters

The demo creates a project with:

- **Name**: `go-sdk-demo-project`
- **Region ID**: `1` (you may need to adjust this for your account)
- **Note**: `Demo project created using CloudCIX Go SDK`

## Output Example

```text
=== CloudCIX Go SDK - Project Creation Demo ===

1. Listing existing projects...
Found 3 existing projects:
  - ID: 123, Name: production-env, Region: 1
  - ID: 124, Name: development-env, Region: 1
  - ID: 125, Name: staging-env, Region: 2

2. Creating a new project...
✅ Successfully created project!
  - Project ID: 126
  - Name: go-sdk-demo-project
  - Region ID: 1
  - Note: Demo project created using CloudCIX Go SDK
  - Created: 2025-11-21T13:45:00Z
  - Manager ID: 456
  - Address ID: 789

3. Retrieving the created project...
✅ Retrieved project details:
  - ID: 126
  - Name: go-sdk-demo-project
  - Closed: false
  - Region ID: 1

=== Demo completed! ===
```

## Error Handling

The demo includes comprehensive error handling for:

- Missing or invalid credentials
- Network connectivity issues
- Invalid region IDs
- Permission errors
- API rate limiting

## Behind the Scenes

The SDK automatically:

1. **Authenticates** with CloudCIX membership API using your credentials
2. **Manages tokens** - creates, stores, and refreshes them every 2 hours
3. **Retries requests** - if a token expires, it gets a new one and retries
4. **Routes requests** - uses V2 API for auth, legacy API for compute operations
5. **Handles errors** - provides clear error messages and suggestions

All of this happens transparently - you just need to provide your credentials once!
