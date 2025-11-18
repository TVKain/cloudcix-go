// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudcix

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/TVKain/cloudcix-go/internal/apijson"
	"github.com/TVKain/cloudcix-go/internal/apiquery"
	shimjson "github.com/TVKain/cloudcix-go/internal/encoding/json"
	"github.com/TVKain/cloudcix-go/internal/requestconfig"
	"github.com/TVKain/cloudcix-go/option"
	"github.com/TVKain/cloudcix-go/packages/param"
	"github.com/TVKain/cloudcix-go/packages/respjson"
)

// NetworkFirewallService contains methods and other services that help with
// interacting with the cloudcix API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNetworkFirewallService] method instead.
type NetworkFirewallService struct {
	Options []option.RequestOption
}

// NewNetworkFirewallService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNetworkFirewallService(opts ...option.RequestOption) (r NetworkFirewallService) {
	r = NetworkFirewallService{}
	r.Options = opts
	return
}

// Create a new firewall configuration with one or more rules. Specify the firewall
// type (project for fine-grained rules or geo for country-based filtering),
// project, router, and the list of rules to apply. Rules are evaluated in the
// order provided.
func (r *NetworkFirewallService) New(ctx context.Context, body NetworkFirewallNewParams, opts ...option.RequestOption) (res *NetworkFirewallResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "network/firewalls/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve detailed information about a specific firewall configuration, including
// its type (project or geo), associated router, complete list of rules, and
// current state.
func (r *NetworkFirewallService) Get(ctx context.Context, pk int64, opts ...option.RequestOption) (res *NetworkFirewallResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("network/firewalls/%v/", pk)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a firewall configuration to modify its name, rules, or state. You can
// replace the entire rule list or change the firewall state to update_running or
// delete.
func (r *NetworkFirewallService) Update(ctx context.Context, pk int64, body NetworkFirewallUpdateParams, opts ...option.RequestOption) (res *NetworkFirewallResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("network/firewalls/%v/", pk)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Retrieve a paginated list of network firewall configurations across your
// projects. Supports filtering by type (project or geo), project, state, and other
// attributes. Results can be ordered and paginated.
//
// ## Filtering
//
// The following fields and modifiers can be used to filter records from the list;
//
// - created (gt, gte, in, isnull, lt, lte, range)
// - id (gt, gte, in, isnull, lt, lte, range)
// - name (in, icontains, iendswith, iexact, istartswith)
// - project_id (gt, gte, in, isnull, lt, lte, range)
// - project\_\_address_id (gt, gte, in, isnull, lt, lte, range)
// - project\_\_name (in, icontains, iendswith, iexact, istartswith)
// - project\_\_region_id (gt, gte, in, isnull, lt, lte, range)
// - project\_\_reseller_id (gt, gte, in, isnull, lt, lte, range)
// - state (gt, gte, in, isnull, lt, lte, range)
// - type
// - updated (gt, gte, in, isnull, lt, lte, range)
//
// To search, simply add `?search[field]=value` to include records that match the
// request, or `?exclude[field]=value` to exclude them. To use modifiers, simply
// add `?search[field__modifier]` and `?exclude[field__modifier]`
//
// ## Ordering
//
// The following fields can be used to order the results of the list;
//
// - name (default)
// - id
// - project_id
//
// To reverse the ordering, simply prepend a `-` character to the request. So
// `?order=field` orders by `field` in ascending order, while `?order=-field`
// orders in descending order instead.
func (r *NetworkFirewallService) List(ctx context.Context, query NetworkFirewallListParams, opts ...option.RequestOption) (res *NetworkFirewallListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "network/firewalls/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type NetworkFirewall struct {
	// The ID of the Network Firewall record
	ID int64 `json:"id,required"`
	// Timestamp, in ISO format, of when the Network Firewall record was created.
	Created string `json:"created,required"`
	// The user-friendly name given to this Network Firewall instance
	Name string `json:"name,required"`
	// The id of the Project that this Network Firewall belongs to
	ProjectID int64 `json:"project_id,required"`
	// List of rules for this Network Firewall.
	Rules []NetworkFirewallRule `json:"rules,required"`
	// An array of the specs for the Network Firewall
	Specs []Bom `json:"specs,required"`
	// The current state of the Network Firewall
	State int64 `json:"state,required"`
	// The type of the Network Firewall
	Type string `json:"type,required"`
	// Timestamp, in ISO format, of when the Network Firewall record was last updated.
	Updated string `json:"updated,required"`
	// URL that can be used to run methods in the API associated with the Network
	// Firewall instance.
	Uri string `json:"uri,required" format:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Created     respjson.Field
		Name        respjson.Field
		ProjectID   respjson.Field
		Rules       respjson.Field
		Specs       respjson.Field
		State       respjson.Field
		Type        respjson.Field
		Updated     respjson.Field
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkFirewall) RawJSON() string { return r.JSON.raw }
func (r *NetworkFirewall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkFirewallRule struct {
	// True to allow traffic, False to deny.
	Allow bool `json:"allow"`
	// Optional description of the rule. Returned if the type is "project".
	Description string `json:"description"`
	// Destination address or subnet. Use \* for any. Returned if the type is
	// "project".
	Destination string `json:"destination"`
	// The name of the Geo IP Address Group. Returned if the type is "geo".
	GroupName string `json:"group_name"`
	// True if the rule applies to inbound traffic.
	Inbound bool `json:"inbound"`
	// Order of rule evaluation (lower runs first). Returned if the type is "project".
	Order int64 `json:"order"`
	// Port or port range (e.g. 80, 443, 1000-2000). Not required for ICMP or ANY.
	// Returned if the type is "project".
	Port string `json:"port"`
	// Network protocol (any, icmp, tcp, udp). Returned if the type is "project".
	Protocol string `json:"protocol"`
	// Source address or subnet. Use \* for any. Returned if the type is "project".
	Source string `json:"source"`
	// IP version (4 or 6). Returned if the type is "project".
	Version int64 `json:"version"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Allow       respjson.Field
		Description respjson.Field
		Destination respjson.Field
		GroupName   respjson.Field
		Inbound     respjson.Field
		Order       respjson.Field
		Port        respjson.Field
		Protocol    respjson.Field
		Source      respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkFirewallRule) RawJSON() string { return r.JSON.raw }
func (r *NetworkFirewallRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkFirewallResponse struct {
	Content NetworkFirewall `json:"content"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkFirewallResponse) RawJSON() string { return r.JSON.raw }
func (r *NetworkFirewallResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property State is required.
type NetworkFirewallUpdateParam struct {
	// Change the state of the Network Firewall, triggering the CloudCIX Robot to
	// perform the requested action. Users can only request state changes from certain
	// current states:
	//
	// - running -> update_running or delete
	State string `json:"state,required"`
	// The user-friendly name for the Network Firewall type. If not sent, it will
	// default to current name.
	Name param.Opt[string] `json:"name,omitzero"`
	// CRITICAL WARNING: This completely replaces ALL existing firewall rules. Any
	// rules not included in this update will be permanently deleted. You must include
	// the complete list of all rules you want to keep, both existing and new ones.
	//
	// A list of the rules to be configured in the Network Firewall type. They will be
	// applied in the order they are sent.
	Rules []NetworkFirewallUpdateRuleParam `json:"rules,omitzero"`
	paramObj
}

func (r NetworkFirewallUpdateParam) MarshalJSON() (data []byte, err error) {
	type shadow NetworkFirewallUpdateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NetworkFirewallUpdateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkFirewallUpdateRuleParam struct {
	// Optional. A flag that states whether traffic matching this rule should be
	// allowed to pass through the Network Firewall Type. If not sent, it wlll default
	// to False.
	Allow param.Opt[bool] `json:"allow,omitzero"`
	// Optional description of what the rule is for.
	Description param.Opt[string] `json:"description,omitzero"`
	// Required if type is "project".
	//
	// A Subnet, IP Address or the name of a Project Network IP Group with `@`
	// prepended that indicates what the destination of traffic should be in order to
	// match this rule.
	//
	// It can also be just a `*` character, which will indicate that any destination is
	// allowed.
	//
	// Both source and destination must use the same IP Version.
	Destination param.Opt[string] `json:"destination,omitzero"`
	// Required if type is "geo".
	//
	// The name of a Geo Network IP Group with `@` prepended.
	GroupName param.Opt[string] `json:"group_name,omitzero"`
	// Optional. Flag indicating the direction of traffic this rule applies to:
	//
	// - true: Inbound rule (traffic coming INTO your project/network)
	// - false: Outbound rule (traffic going OUT FROM your project/network)
	//
	// If not sent, it will default to False (outbound rule).
	Inbound param.Opt[bool] `json:"inbound,omitzero"`
	// Required if type is "project".
	//
	// A string that indicates what the destination port of traffic should be in order
	// to match this rule. The range for valid ports are between 1 - 65535 inclusive.
	//
	// Allowed Values: `22`: Only port 22 is allowed `20-25`: Ports between 20 and 25
	// inclusive are allowed `22,24-25,444`: Combination of one or more single port and
	// single port ranges with comma separated are allowed “: No port is required if
	// the protocol is 'any' or 'icmp'
	Port param.Opt[string] `json:"port,omitzero"`
	// Required if type is "project".
	//
	// A string that indicates what protocol traffic should be using in order to match
	// this rule. The supported protocols are; - 'icmp' - 'tcp' - 'udp'
	//
	// The special case protocol 'any' is allowed and allows any protocol through.
	Protocol param.Opt[string] `json:"protocol,omitzero"`
	// Required if type is "project".
	//
	// A Subnet, IP Address or the name of a Project Network IP Group with `@`
	// prepended that indicates what the destination of traffic should be in order to
	// match this rule.
	//
	// It can also be just a `*` character, which will indicate that any source is
	// allowed.
	//
	// Both source and destination must use the same IP Version.
	Source param.Opt[string] `json:"source,omitzero"`
	paramObj
}

func (r NetworkFirewallUpdateRuleParam) MarshalJSON() (data []byte, err error) {
	type shadow NetworkFirewallUpdateRuleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NetworkFirewallUpdateRuleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkFirewallListResponse struct {
	Metadata ListMetadata      `json:"_metadata"`
	Content  []NetworkFirewall `json:"content"`
	// Maximum number of records returned per page.
	Count int64 `json:"count"`
	// Page number of the current results.
	Page int64 `json:"page"`
	// Total number of matching records.
	Total int64 `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metadata    respjson.Field
		Content     respjson.Field
		Count       respjson.Field
		Page        respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkFirewallListResponse) RawJSON() string { return r.JSON.raw }
func (r *NetworkFirewallListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkFirewallNewParams struct {
	// The ID of the Project which this Network Firewall should be created in. Each
	// project can have exactly ONE project firewall and ONE geo firewall maximum.
	ProjectID int64 `json:"project_id,required"`
	// The user-friendly name for the Network Firewall type. If not sent and the type
	// is "geo", it will default to the name 'Geofilter'. If not sent and the type is
	// "project", it will default to the name 'Firewall'.
	Name param.Opt[string] `json:"name,omitzero"`
	// The type of Network Firewall to create. Each project can have exactly ONE of
	// each type. Valid options are:
	//
	//   - "geo" A Geofilter Firewall to allow or block traffic from/to specific
	//     countries using global IP Address Groups (member_id = 0) that contain country
	//     IP ranges.
	//   - "project" A Project Firewall with fine-grained rules for specific
	//     source/destination IPs, ports, and protocols. Can reference your member's IP
	//     Address Groups using '@groupname' syntax.
	Type param.Opt[string] `json:"type,omitzero"`
	// A list of the rules to be configured in the Network Firewall type. They will be
	// applied in the order they are sent.
	Rules []NetworkFirewallNewParamsRule `json:"rules,omitzero"`
	paramObj
}

func (r NetworkFirewallNewParams) MarshalJSON() (data []byte, err error) {
	type shadow NetworkFirewallNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NetworkFirewallNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkFirewallNewParamsRule struct {
	// Optional. A flag that states whether traffic matching this rule should be
	// allowed to pass through the Network Firewall Type. If not sent, it wlll default
	// to False.
	Allow param.Opt[bool] `json:"allow,omitzero"`
	// Optional description of what the rule is for.
	Description param.Opt[string] `json:"description,omitzero"`
	// Required if type is "project".
	//
	// A Subnet, IP Address or the name of a Project Network IP Group with `@`
	// prepended that indicates what the destination of traffic should be in order to
	// match this rule.
	//
	// It can also be just a `*` character, which will indicate that any destination is
	// allowed.
	//
	// Both source and destination must use the same IP Version.
	Destination param.Opt[string] `json:"destination,omitzero"`
	// Required if type is "geo".
	//
	// The name of a Geo Network IP Group with `@` prepended.
	GroupName param.Opt[string] `json:"group_name,omitzero"`
	// Optional. Flag indicating the direction of traffic this rule applies to:
	//
	// - true: Inbound rule (traffic coming INTO your project/network)
	// - false: Outbound rule (traffic going OUT FROM your project/network)
	//
	// If not sent, it will default to False (outbound rule).
	Inbound param.Opt[bool] `json:"inbound,omitzero"`
	// Required if type is "project".
	//
	// A string that indicates what the destination port of traffic should be in order
	// to match this rule. The range for valid ports are between 1 - 65535 inclusive.
	//
	// Allowed Values: `22`: Only port 22 is allowed `20-25`: Ports between 20 and 25
	// inclusive are allowed `22,24-25,444`: Combination of one or more single port and
	// single port ranges with comma separated are allowed “: No port is required if
	// the protocol is 'any' or 'icmp'
	Port param.Opt[string] `json:"port,omitzero"`
	// Required if type is "project".
	//
	// A string that indicates what protocol traffic should be using in order to match
	// this rule. The supported protocols are; - 'icmp' - 'tcp' - 'udp'
	//
	// The special case protocol 'any' is allowed and allows any protocol through.
	Protocol param.Opt[string] `json:"protocol,omitzero"`
	// Required if type is "project".
	//
	// A Subnet, IP Address or the name of a Project Network IP Group with `@`
	// prepended that indicates what the destination of traffic should be in order to
	// match this rule.
	//
	// It can also be just a `*` character, which will indicate that any source is
	// allowed.
	//
	// Both source and destination must use the same IP Version.
	Source param.Opt[string] `json:"source,omitzero"`
	paramObj
}

func (r NetworkFirewallNewParamsRule) MarshalJSON() (data []byte, err error) {
	type shadow NetworkFirewallNewParamsRule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NetworkFirewallNewParamsRule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkFirewallUpdateParams struct {
	NetworkFirewallUpdate NetworkFirewallUpdateParam
	paramObj
}

func (r NetworkFirewallUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.NetworkFirewallUpdate)
}
func (r *NetworkFirewallUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.NetworkFirewallUpdate)
}

type NetworkFirewallListParams struct {
	// The limit of the number of objects returned per page
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The field to use for ordering. Possible fields and the default are outlined in
	// the individual method descriptions.
	Order param.Opt[string] `query:"order,omitzero" json:"-"`
	// The page of records to return, assuming `limit` number of records per page.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Filter the result to objects that do not match the specified filters. Possible
	// filters are outlined in the individual list method descriptions.
	Exclude any `query:"exclude,omitzero" json:"-"`
	// Filter the result to objects that match the specified filters. Possible filters
	// are outlined in the individual list method descriptions.
	Search any `query:"search,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [NetworkFirewallListParams]'s query parameters as
// `url.Values`.
func (r NetworkFirewallListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
