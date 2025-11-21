# CloudCIX API Architecture

CloudCIX uses **two separate APIs** to handle different aspects of the platform:

## 1. Membership API
- **URL**: `https://membership.api.cloudcix.com/`
- **Purpose**: Authentication and user management
- **Used for**: 
  - User login and authentication
  - Token generation and management
  - User account operations
- **Authentication endpoint**: `/auth/login/`
- **Request format**: POST with email, password, and api_key

## 2. Compute API
- **URL**: `https://compute.api.cloudcix.com/`
- **Purpose**: Cloud infrastructure management
- **Used for**:
  - Creating and managing compute instances
  - Project management
  - Storage volumes
  - Network configurations
  - Backups and snapshots
- **Authentication**: Uses tokens obtained from Membership API
- **Header**: `X-Auth-Token: <token>`

## How the Go SDK Handles Both APIs

The CloudCIX Go SDK automatically manages both APIs:

1. **Authentication Flow**:
   ```
   User Credentials → Membership API → Auth Token → Compute API Requests
   ```

2. **Automatic Routing**:
   - Authentication requests → `{base_url}/membership/auth/login/`
   - All other operations → `{base_url}/compute/`

3. **Token Management**:
   - Tokens expire every 2 hours
   - SDK automatically refreshes tokens before expiry
   - Failed requests with expired tokens are automatically retried

## Configuration

In your `my_settings` file:

```env
# Base API URL - service endpoints are constructed automatically
CLOUDCIX_API_URL="https://api.cloudcix.com/"

# Your region ID (required for creating resources)
CLOUDCIX_REGION_ID=12345

# Your CloudCIX credentials
CLOUDCIX_API_USERNAME="your_email@example.com"
CLOUDCIX_API_PASSWORD="your_password"
CLOUDCIX_API_KEY="your_api_key"
```

The SDK automatically constructs:

- **Membership API**: `https://membership.api.cloudcix.com/` (for authentication)
- **Compute API**: `https://compute.api.cloudcix.com/` (for compute operations)

## API Endpoints Summary

### Membership API Endpoints
- `POST /auth/login/` - Authenticate and get token

### Compute API Endpoints
- `GET /project/` - List projects
- `POST /project/` - Create project
- `GET /compute/instances/` - List instances
- `POST /compute/instances/` - Create instance
- `GET /compute/backups/` - List backups
- And many more...

## Why Two APIs?

This separation allows CloudCIX to:
- **Scale independently** - auth and compute can be scaled separately
- **Security isolation** - credentials are only sent to membership API
- **Service specialization** - each API can be optimized for its purpose
- **Flexibility** - different teams can work on different APIs

## SDK Benefits

The Go SDK abstracts away this complexity:
- ✅ **Single client** - one client handles both APIs automatically
- ✅ **Automatic routing** - requests go to the right API
- ✅ **Token management** - no manual token handling needed
- ✅ **Error handling** - automatic retry with fresh tokens
- ✅ **Type safety** - strongly typed parameters and responses