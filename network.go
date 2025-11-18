// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudcix

import (
	"github.com/stainless-sdks/cloudcix-go/option"
)

// NetworkService contains methods and other services that help with interacting
// with the cloudcix API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNetworkService] method instead.
type NetworkService struct {
	Options   []option.RequestOption
	Firewalls NetworkFirewallService
	IPGroups  NetworkIPGroupService
	Routers   NetworkRouterService
	Vpns      NetworkVpnService
}

// NewNetworkService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewNetworkService(opts ...option.RequestOption) (r NetworkService) {
	r = NetworkService{}
	r.Options = opts
	r.Firewalls = NewNetworkFirewallService(opts...)
	r.IPGroups = NewNetworkIPGroupService(opts...)
	r.Routers = NewNetworkRouterService(opts...)
	r.Vpns = NewNetworkVpnService(opts...)
	return
}
