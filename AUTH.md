# CloudCIX Go SDK Authentication

The CloudCIX Go SDK provides automatic authentication and token management for the CloudCIX API. The SDK handles token creation, renewal, and retry logic automatically.

## Configuration

### Settings File

Create a settings file (e.g., `my_settings`) with your CloudCIX credentials:

```env
CLOUDCIX_API_URL="https://compute.api.cloudcix.com/"
CLOUDCIX_API_V2_URL="https://membership.api.cloudcix.com/"
CLOUDCIX_API_VERSION="5.0"
CLOUDCIX_API_USERNAME="your_email@example.com"
CLOUDCIX_API_PASSWORD="your_password"
CLOUDCIX_API_KEY="your_api_key"
```

### Environment Variables

Alternatively, set environment variables:

```bash
export CLOUDCIX_API_USERNAME="your_email@example.com"
export CLOUDCIX_API_PASSWORD="your_password"
export CLOUDCIX_API_KEY="your_api_key"
```

## Usage

### Basic Usage

```go
package main

import (
    "context"
    "log"
    
    "github.com/TVKain/cloudcix-go"
)

func main() {
    // Option 1: Auto-detect credentials from environment variables
    // Will use automatic authentication if credentials are available
    client := cloudcix.NewClient()
    
    // Option 2: Require authentication (will fail if credentials missing)
    client, err := cloudcix.NewClientWithAuth()
    if err != nil {
        log.Fatal(err)
    }
    
    // Option 3: Load from specific settings file
    client, err = cloudcix.NewClientFromFile("my_settings")
    if err != nil {
        log.Fatal(err)
    }
    
    // Make authenticated requests - token management is automatic
    ctx := context.Background()
    backups, err := client.Compute.Backups.List(ctx, cloudcix.ComputeBackupListParams{})
    if err != nil {
        log.Fatal(err)
    }
}
```

### Manual Settings

```go
package main

import (
    "github.com/TVKain/cloudcix-go"
    "github.com/TVKain/cloudcix-go/config"
)

func main() {
    settings := &config.Settings{
        CLOUDCIX_API_URL:      "https://compute.api.cloudcix.com/",
        CLOUDCIX_API_V2_URL:   "https://membership.api.cloudcix.com/",
        CLOUDCIX_API_VERSION:  "5.0",
        CLOUDCIX_API_USERNAME: "your_email@example.com",
        CLOUDCIX_API_PASSWORD: "your_password",
        CLOUDCIX_API_KEY:      "your_api_key",
    }
    
    client := cloudcix.NewClientWithSettings(settings)
    
    // Use client normally...
}
```

## Automatic Features

The SDK automatically handles:

1. **Token Creation**: Obtains authentication tokens from the membership API
2. **Token Renewal**: Refreshes tokens every 2 hours automatically
3. **Retry Logic**: Automatically retries failed requests with fresh tokens
4. **Thread Safety**: Token management is safe for concurrent use
5. **Error Handling**: Comprehensive error handling for authentication failures

## API Endpoints

- **Compute API**: Uses `CLOUDCIX_API_URL` (legacy API)
- **Membership API**: Uses `CLOUDCIX_API_V2_URL` for authentication

The SDK automatically routes requests to the appropriate API endpoint.

## Manual Token Management

```go
// Manually refresh token
err := client.RefreshToken(context.Background())
if err != nil {
    log.Printf("Failed to refresh token: %v", err)
}

// Get a token directly (without client)
token, err := auth.GetToken("email", "password", "api_key")
if err != nil {
    log.Fatal(err)
}

// Get admin token from environment/settings
adminToken, err := auth.GetAdminToken("my_settings")
if err != nil {
    log.Fatal(err)
}
```

## Error Handling

The SDK will automatically handle authentication errors and retry requests. If authentication repeatedly fails, the SDK will return the underlying authentication error.

```go
backups, err := client.Compute.Backups.List(ctx, params)
if err != nil {
    // This could be an authentication error, network error, or API error
    log.Printf("Request failed: %v", err)
}
```

## Security Considerations

- Store credentials securely (environment variables or protected files)
- Use HTTPS endpoints only
- Rotate API keys regularly
- Monitor token usage and authentication logs
