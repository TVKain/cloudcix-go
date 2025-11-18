// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudcix

import (
	"github.com/stainless-sdks/cloudcix-go/option"
)

// ComputeService contains methods and other services that help with interacting
// with the cloudcix API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewComputeService] method instead.
type ComputeService struct {
	Options   []option.RequestOption
	Backups   ComputeBackupService
	GPUs      ComputeGPUService
	Images    ComputeImageService
	Instances ComputeInstanceService
	Snapshots ComputeSnapshotService
}

// NewComputeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewComputeService(opts ...option.RequestOption) (r ComputeService) {
	r = ComputeService{}
	r.Options = opts
	r.Backups = NewComputeBackupService(opts...)
	r.GPUs = NewComputeGPUService(opts...)
	r.Images = NewComputeImageService(opts...)
	r.Instances = NewComputeInstanceService(opts...)
	r.Snapshots = NewComputeSnapshotService(opts...)
	return
}
