// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudcix

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/cloudcix-go/internal/apijson"
	"github.com/stainless-sdks/cloudcix-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/cloudcix-go/internal/encoding/json"
	"github.com/stainless-sdks/cloudcix-go/internal/requestconfig"
	"github.com/stainless-sdks/cloudcix-go/option"
	"github.com/stainless-sdks/cloudcix-go/packages/param"
	"github.com/stainless-sdks/cloudcix-go/packages/respjson"
)

// StorageVolumeService contains methods and other services that help with
// interacting with the cloudcix API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStorageVolumeService] method instead.
type StorageVolumeService struct {
	Options []option.RequestOption
}

// NewStorageVolumeService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStorageVolumeService(opts ...option.RequestOption) (r StorageVolumeService) {
	r = StorageVolumeService{}
	r.Options = opts
	return
}

// Create a new storage volume in a project. Specify volume type (Ceph or Hyper-V),
// storage capacity using SKUs, and configuration options. Ceph volumes can be
// mounted to multiple LXD instances. Hyper-V volumes are attached as secondary
// drives to Hyper-V instances.
func (r *StorageVolumeService) New(ctx context.Context, body StorageVolumeNewParams, opts ...option.RequestOption) (res *StorageVolumesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "storage/volumes/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve detailed information about a specific storage volume including its
// type, capacity, mount status, and associated instances.
func (r *StorageVolumeService) Get(ctx context.Context, pk int64, opts ...option.RequestOption) (res *StorageVolumesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("storage/volumes/%v/", pk)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a storage volume's configuration. Modify volume name, increase storage
// capacity (cannot decrease), mount/unmount from instances (Ceph), or trigger
// state changes for resource management.
func (r *StorageVolumeService) Update(ctx context.Context, pk int64, body StorageVolumeUpdateParams, opts ...option.RequestOption) (res *StorageVolumesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("storage/volumes/%v/", pk)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Retrieve a paginated list of storage volumes accessible to the requesting user.
// Filter by volume type (Ceph or Hyper-V), project, region, and other attributes.
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
func (r *StorageVolumeService) List(ctx context.Context, query StorageVolumeListParams, opts ...option.RequestOption) (res *StorageVolumeListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "storage/volumes/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type StorageVolumes struct {
	// The ID of the Storage Volume record
	ID int64 `json:"id,required"`
	// A list of Compute Instances the Storage Volume is mounted to.
	ContraInstances []StorageVolumesContraInstance `json:"contra_instances,required"`
	// Timestamp, in ISO format, of when the Storage Volume was created.
	Created string `json:"created,required"`
	// The "hyperv" Compute Instance the "hyperv" Storage Volume is attached to.
	Instance StorageVolumesInstance `json:"instance,required"`
	// The metadata object of "ceph" Storage Volumes
	Metadata StorageVolumesMetadata `json:"metadata,required"`
	// The user-friendly name given to this Storage Volume
	Name string `json:"name,required"`
	// The ID of the Project that this Storage Volume belongs to
	ProjectID int64 `json:"project_id,required"`
	// An array of the specs for the Storage Volume
	Specs []Bom `json:"specs,required"`
	// The current state of the Storage Volume
	State int64 `json:"state,required"`
	// The type of the Storage Volume
	Type string `json:"type,required"`
	// Timestamp, in ISO format, of when the Storage Volume was last updated.
	Updated string `json:"updated,required"`
	// URL that can be used to run methods in the API associated with the Storage
	// Volumes instance.
	Uri string `json:"uri,required" format:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		ContraInstances respjson.Field
		Created         respjson.Field
		Instance        respjson.Field
		Metadata        respjson.Field
		Name            respjson.Field
		ProjectID       respjson.Field
		Specs           respjson.Field
		State           respjson.Field
		Type            respjson.Field
		Updated         respjson.Field
		Uri             respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageVolumes) RawJSON() string { return r.JSON.raw }
func (r *StorageVolumes) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageVolumesContraInstance struct {
	// The ID of the Compute Instance the Storage Volume is attached to.
	ID int64 `json:"id"`
	// The user-friendly name given to the Compute Instance the Storage Volume is
	// attached to.
	Name string `json:"name"`
	// The current state of the Compute Instance the Storage Volume is attached to.
	State int64 `json:"state"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		State       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageVolumesContraInstance) RawJSON() string { return r.JSON.raw }
func (r *StorageVolumesContraInstance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The "hyperv" Compute Instance the "hyperv" Storage Volume is attached to.
type StorageVolumesInstance struct {
	// The ID of the "hyperv" Compute Instance the "hyperv" Storage Volume is attached
	// to.
	ID int64 `json:"id"`
	// The user-friendly name of the "hyperv" Compute Instance the "hyperv" Storage
	// Volume is attached to.
	Name string `json:"name"`
	// The current state of the "hyperv" Compute Instance the "hyperv" Storage Volume
	// is attached to.
	State int64 `json:"state"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		State       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageVolumesInstance) RawJSON() string { return r.JSON.raw }
func (r *StorageVolumesInstance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The metadata object of "ceph" Storage Volumes
type StorageVolumesMetadata struct {
	// List of IDs of "lxd" Compute instances which the "ceph" Storage Volume should be
	// attached to.
	AttachInstanceIDs []int64 `json:"attach_instance_ids"`
	// List of IDs of "lxd" Compute instances which the "ceph" Storage Volume should be
	// detached from.
	DetachInstanceIDs []int64 `json:"detach_instance_ids"`
	// The mpunt path of the "ceph" Storage Volume on the "lxd" Compute instances.
	MountPath string `json:"mount_path"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AttachInstanceIDs respjson.Field
		DetachInstanceIDs respjson.Field
		MountPath         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageVolumesMetadata) RawJSON() string { return r.JSON.raw }
func (r *StorageVolumesMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageVolumesResponse struct {
	Content StorageVolumes `json:"content"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StorageVolumesResponse) RawJSON() string { return r.JSON.raw }
func (r *StorageVolumesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageVolumesUpdateParam struct {
	// The user-friendly name for the Storage Volume. If not sent, it will default to
	// current name.
	Name param.Opt[string] `json:"name,omitzero"`
	// Change the state of the Storage Volume, triggering the CloudCIX Robot to perform
	// the requested action. Users can only request state changes from certain current
	// states:
	//
	// - running -> update_running or delete
	State param.Opt[string] `json:"state,omitzero"`
	// Required if type is "cephfs".
	//
	// Metadata for the Storage Volume drive with the type "cephfs".
	Metadata StorageVolumesUpdateMetadataParam `json:"metadata,omitzero"`
	// List of specs (SKUs) for the Storage Volume drive. Only required if increasing
	// the current size of the Storage Volume.
	Specs []StorageVolumesUpdateSpecParam `json:"specs,omitzero"`
	paramObj
}

func (r StorageVolumesUpdateParam) MarshalJSON() (data []byte, err error) {
	type shadow StorageVolumesUpdateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StorageVolumesUpdateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Required if type is "cephfs".
//
// Metadata for the Storage Volume drive with the type "cephfs".
type StorageVolumesUpdateMetadataParam struct {
	// The mount path for the Ceph file system volume inside the LXC instance.
	MountPath param.Opt[string] `json:"mount_path,omitzero"`
	// A list of IDs for running or stopped Compute Instances with the type "lxd" in
	// the project to mount this Ceph file system volume to. If not sent, it will
	// default to an empty list.
	AttachInstanceIDs []int64 `json:"attach_instance_ids,omitzero"`
	// A list of IDs for running or stopped Compute Instances with the type "lxd" in
	// the project to unmount this Ceph file system volume from. If not sent, it will
	// default to an empty list.
	DetachInstanceIDs []int64 `json:"detach_instance_ids,omitzero"`
	paramObj
}

func (r StorageVolumesUpdateMetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow StorageVolumesUpdateMetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StorageVolumesUpdateMetadataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageVolumesUpdateSpecParam struct {
	// The quantity (GB) of the SKU to configure the Storage Volume drive with.
	Quantity param.Opt[int64] `json:"quantity,omitzero"`
	// The name of the SKU for the Storage Volume drive.
	SKUName param.Opt[string] `json:"sku_name,omitzero"`
	paramObj
}

func (r StorageVolumesUpdateSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow StorageVolumesUpdateSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StorageVolumesUpdateSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageVolumeListResponse struct {
	Metadata ListMetadata     `json:"_metadata"`
	Content  []StorageVolumes `json:"content"`
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
func (r StorageVolumeListResponse) RawJSON() string { return r.JSON.raw }
func (r *StorageVolumeListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageVolumeNewParams struct {
	// The ID of the Project which this Storage Volume should be in.
	ProjectID int64 `json:"project_id,required"`
	// List of specs (SKUs) for the Storage Volume drive.
	Specs []StorageVolumeNewParamsSpec `json:"specs,omitzero,required"`
	// Required if type is "hyperv".
	//
	// The ID of a Compute Instance with the type "hyperv" the Storage Volume is to be
	// mounted to.
	InstanceID param.Opt[int64] `json:"instance_id,omitzero"`
	// The user-friendly name for the Storage Volume type. If not sent and the type is
	// "cephfs", it will default to the name 'Ceph'. If not sent and the type is
	// "hyperv", it will default to the name 'Storage HyperV'.
	Name param.Opt[string] `json:"name,omitzero"`
	// The type of Storage Volume to create. Valid options are:
	//
	//   - "cephfs" A Ceph file system volume which can be mounted to one or more Compute
	//     Instances of the type "lxd"
	//   - "hyperv" A volume which can be mounted as a secondary drive to a Compute
	//     Instance of the type "hyperv"
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

func (r StorageVolumeNewParams) MarshalJSON() (data []byte, err error) {
	type shadow StorageVolumeNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StorageVolumeNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageVolumeNewParamsSpec struct {
	// The quantity (GB) of the SKU to configure the Storage Volume drive with.
	Quantity param.Opt[int64] `json:"quantity,omitzero"`
	// The name of the SKU for the Storage Volume drive.
	SKUName param.Opt[string] `json:"sku_name,omitzero"`
	paramObj
}

func (r StorageVolumeNewParamsSpec) MarshalJSON() (data []byte, err error) {
	type shadow StorageVolumeNewParamsSpec
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StorageVolumeNewParamsSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StorageVolumeUpdateParams struct {
	StorageVolumesUpdate StorageVolumesUpdateParam
	paramObj
}

func (r StorageVolumeUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.StorageVolumesUpdate)
}
func (r *StorageVolumeUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.StorageVolumesUpdate)
}

type StorageVolumeListParams struct {
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

// URLQuery serializes [StorageVolumeListParams]'s query parameters as
// `url.Values`.
func (r StorageVolumeListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
