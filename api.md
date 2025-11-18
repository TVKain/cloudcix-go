# Compute

## Backups

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeBackupUpdateParam">ComputeBackupUpdateParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeBackup">ComputeBackup</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeBackupResponse">ComputeBackupResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ListMetadata">ListMetadata</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeBackupListResponse">ComputeBackupListResponse</a>

Methods:

- <code title="post /compute/backups/">client.Compute.Backups.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeBackupService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeBackupNewParams">ComputeBackupNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeBackupResponse">ComputeBackupResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /compute/backups/{pk}/">client.Compute.Backups.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeBackupService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, pk <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeBackupResponse">ComputeBackupResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /compute/backups/{pk}/">client.Compute.Backups.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeBackupService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, pk <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeBackupUpdateParams">ComputeBackupUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeBackupResponse">ComputeBackupResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /compute/backups/">client.Compute.Backups.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeBackupService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeBackupListParams">ComputeBackupListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeBackupListResponse">ComputeBackupListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## GPUs

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeGPUUpdateParam">ComputeGPUUpdateParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeGPU">ComputeGPU</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeGPUResponse">ComputeGPUResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeGPUListResponse">ComputeGPUListResponse</a>

Methods:

- <code title="get /compute/gpus/{pk}/">client.Compute.GPUs.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeGPUService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, pk <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeGPUResponse">ComputeGPUResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /compute/gpus/{pk}/">client.Compute.GPUs.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeGPUService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, pk <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeGPUUpdateParams">ComputeGPUUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeGPUResponse">ComputeGPUResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /compute/gpus/">client.Compute.GPUs.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeGPUService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeGPUListParams">ComputeGPUListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeGPUListResponse">ComputeGPUListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /compute/gpus/">client.Compute.GPUs.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeGPUService.Attach">Attach</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeGPUAttachParams">ComputeGPUAttachParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeGPUResponse">ComputeGPUResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Images

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeImage">ComputeImage</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeImageGetResponse">ComputeImageGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeImageListResponse">ComputeImageListResponse</a>

Methods:

- <code title="get /compute/images/{pk}/">client.Compute.Images.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeImageService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, pk <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeImageGetResponse">ComputeImageGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /compute/images/">client.Compute.Images.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeImageService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeImageListParams">ComputeImageListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeImageListResponse">ComputeImageListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Instances

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeInstanceUpdateParam">ComputeInstanceUpdateParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#Bom">Bom</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeInstance">ComputeInstance</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeInstanceResponse">ComputeInstanceResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeInstanceListResponse">ComputeInstanceListResponse</a>

Methods:

- <code title="post /compute/instances/">client.Compute.Instances.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeInstanceService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeInstanceNewParams">ComputeInstanceNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeInstanceResponse">ComputeInstanceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /compute/instances/{pk}/">client.Compute.Instances.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeInstanceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, pk <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeInstanceResponse">ComputeInstanceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /compute/instances/{pk}/">client.Compute.Instances.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeInstanceService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, pk <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeInstanceUpdateParams">ComputeInstanceUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeInstanceResponse">ComputeInstanceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /compute/instances/">client.Compute.Instances.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeInstanceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeInstanceListParams">ComputeInstanceListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeInstanceListResponse">ComputeInstanceListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Snapshots

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeSnapshotUpdateParam">ComputeSnapshotUpdateParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeSnapshot">ComputeSnapshot</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeSnapshotResponse">ComputeSnapshotResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeSnapshotListResponse">ComputeSnapshotListResponse</a>

Methods:

- <code title="post /compute/snapshots/">client.Compute.Snapshots.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeSnapshotService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeSnapshotNewParams">ComputeSnapshotNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeSnapshotResponse">ComputeSnapshotResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /compute/snapshots/{pk}/">client.Compute.Snapshots.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeSnapshotService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, pk <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeSnapshotResponse">ComputeSnapshotResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /compute/snapshots/{pk}/">client.Compute.Snapshots.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeSnapshotService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, pk <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeSnapshotUpdateParams">ComputeSnapshotUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeSnapshotResponse">ComputeSnapshotResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /compute/snapshots/">client.Compute.Snapshots.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeSnapshotService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeSnapshotListParams">ComputeSnapshotListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ComputeSnapshotListResponse">ComputeSnapshotListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Network

## Firewalls

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkFirewallUpdateParam">NetworkFirewallUpdateParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkFirewall">NetworkFirewall</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkFirewallResponse">NetworkFirewallResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkFirewallListResponse">NetworkFirewallListResponse</a>

Methods:

- <code title="post /network/firewalls/">client.Network.Firewalls.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkFirewallService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkFirewallNewParams">NetworkFirewallNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkFirewallResponse">NetworkFirewallResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /network/firewalls/{pk}/">client.Network.Firewalls.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkFirewallService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, pk <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkFirewallResponse">NetworkFirewallResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /network/firewalls/{pk}/">client.Network.Firewalls.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkFirewallService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, pk <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkFirewallUpdateParams">NetworkFirewallUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkFirewallResponse">NetworkFirewallResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /network/firewalls/">client.Network.Firewalls.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkFirewallService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkFirewallListParams">NetworkFirewallListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkFirewallListResponse">NetworkFirewallListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## IPGroups

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkIPGroupUpdateParam">NetworkIPGroupUpdateParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkIPGroup">NetworkIPGroup</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkIPGroupResponse">NetworkIPGroupResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkIPGroupListResponse">NetworkIPGroupListResponse</a>

Methods:

- <code title="post /network/ip_groups/">client.Network.IPGroups.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkIPGroupService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkIPGroupNewParams">NetworkIPGroupNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkIPGroupResponse">NetworkIPGroupResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /network/ip_groups/{pk}/">client.Network.IPGroups.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkIPGroupService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, pk <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkIPGroupResponse">NetworkIPGroupResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /network/ip_groups/{pk}/">client.Network.IPGroups.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkIPGroupService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, pk <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkIPGroupUpdateParams">NetworkIPGroupUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkIPGroupResponse">NetworkIPGroupResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /network/ip_groups/">client.Network.IPGroups.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkIPGroupService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkIPGroupListParams">NetworkIPGroupListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkIPGroupListResponse">NetworkIPGroupListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /network/ip_groups/{pk}/">client.Network.IPGroups.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkIPGroupService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, pk <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Routers

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkRouterUpdateParam">NetworkRouterUpdateParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#BaseIPAddress">BaseIPAddress</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkRouter">NetworkRouter</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkRouterResponse">NetworkRouterResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkRouterListResponse">NetworkRouterListResponse</a>

Methods:

- <code title="post /network/routers/">client.Network.Routers.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkRouterService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkRouterNewParams">NetworkRouterNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkRouterResponse">NetworkRouterResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /network/routers/{pk}/">client.Network.Routers.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkRouterService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, pk <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkRouterResponse">NetworkRouterResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /network/routers/{pk}/">client.Network.Routers.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkRouterService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, pk <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkRouterUpdateParams">NetworkRouterUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkRouterResponse">NetworkRouterResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /network/routers/">client.Network.Routers.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkRouterService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkRouterListParams">NetworkRouterListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkRouterListResponse">NetworkRouterListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Vpns

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkVpnUpdateParam">NetworkVpnUpdateParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkVpn">NetworkVpn</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkVpnResponse">NetworkVpnResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkVpnListResponse">NetworkVpnListResponse</a>

Methods:

- <code title="post /network/vpns/">client.Network.Vpns.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkVpnService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkVpnNewParams">NetworkVpnNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkVpnResponse">NetworkVpnResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /network/vpns/{pk}/">client.Network.Vpns.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkVpnService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, pk <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkVpnResponse">NetworkVpnResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /network/vpns/{pk}/">client.Network.Vpns.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkVpnService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, pk <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkVpnUpdateParams">NetworkVpnUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkVpnResponse">NetworkVpnResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /network/vpns/">client.Network.Vpns.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkVpnService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkVpnListParams">NetworkVpnListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#NetworkVpnListResponse">NetworkVpnListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Project

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ProjectUpdateParam">ProjectUpdateParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#Project">Project</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ProjectResponse">ProjectResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ProjectListResponse">ProjectListResponse</a>

Methods:

- <code title="post /project/">client.Project.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ProjectService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ProjectNewParams">ProjectNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ProjectResponse">ProjectResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /project/{pk}/">client.Project.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ProjectService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, pk <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ProjectResponse">ProjectResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /project/{pk}/">client.Project.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ProjectService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, pk <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ProjectUpdateParams">ProjectUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ProjectResponse">ProjectResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /project/">client.Project.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ProjectService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ProjectListParams">ProjectListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#ProjectListResponse">ProjectListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Storage

## Volumes

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#StorageVolumesUpdateParam">StorageVolumesUpdateParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#StorageVolumes">StorageVolumes</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#StorageVolumesResponse">StorageVolumesResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#StorageVolumeListResponse">StorageVolumeListResponse</a>

Methods:

- <code title="post /storage/volumes/">client.Storage.Volumes.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#StorageVolumeService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#StorageVolumeNewParams">StorageVolumeNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#StorageVolumesResponse">StorageVolumesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /storage/volumes/{pk}/">client.Storage.Volumes.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#StorageVolumeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, pk <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#StorageVolumesResponse">StorageVolumesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /storage/volumes/{pk}/">client.Storage.Volumes.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#StorageVolumeService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, pk <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#StorageVolumeUpdateParams">StorageVolumeUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#StorageVolumesResponse">StorageVolumesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /storage/volumes/">client.Storage.Volumes.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#StorageVolumeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#StorageVolumeListParams">StorageVolumeListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go">cloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/cloudcix-go#StorageVolumeListResponse">StorageVolumeListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
