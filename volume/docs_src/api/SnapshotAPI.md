# \SnapshotAPI

All URIs are relative to *https://volume.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSnapshot**](SnapshotAPI.md#CreateSnapshot) | **Post** /api/v1/volumes/{volume_id}/snapshots | Create snapshot
[**DeleteSnapshot**](SnapshotAPI.md#DeleteSnapshot) | **Delete** /api/v1/snapshots/{snapshot_id} | Delete snapshot
[**GetSnapshot**](SnapshotAPI.md#GetSnapshot) | **Get** /api/v1/snapshots/{snapshot_id} | Get snapshot
[**ListSnapshots**](SnapshotAPI.md#ListSnapshots) | **Get** /api/v1/snapshots | List snapshots
[**RestoreSnapshot**](SnapshotAPI.md#RestoreSnapshot) | **Post** /api/v1/snapshots/{snapshot_id}/restore | Restore snapshot
[**UpdateSnapshot**](SnapshotAPI.md#UpdateSnapshot) | **Put** /api/v1/snapshots/{snapshot_id} | Update snapshot



## CreateSnapshot

> CreateSnapshotResponse CreateSnapshot(ctx, volumeId).CreateSnapshotRequest(createSnapshotRequest).Execute()

Create snapshot



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/volume"
)

func main() {
	volumeId := "volumeId_example" // string | 볼륨의 고유 ID - [List volumes](/openapi/bcs/list-volumes)에서 확인
	createSnapshotRequest := *openapiclient.NewCreateSnapshotRequest(*openapiclient.NewCreateSnapshot(false, "Name_example")) // CreateSnapshotRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnapshotAPI.CreateSnapshot(context.Background(), volumeId).CreateSnapshotRequest(createSnapshotRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotAPI.CreateSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSnapshot`: CreateSnapshotResponse
	fmt.Fprintf(os.Stdout, "Response from `SnapshotAPI.CreateSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | 볼륨의 고유 ID - [List volumes](/openapi/bcs/list-volumes)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createSnapshotRequest** | [**CreateSnapshotRequest**](CreateSnapshotRequest.md) |  | 

### Return type

[**CreateSnapshotResponse**](CreateSnapshotResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSnapshot

> DeleteSnapshot(ctx, snapshotId).Execute()

Delete snapshot



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/volume"
)

func main() {
	snapshotId := "snapshotId_example" // string | 스냅샷의 고유 ID - [List snapshots](/openapi/bcs/list-snapshots)에서 확인

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SnapshotAPI.DeleteSnapshot(context.Background(), snapshotId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotAPI.DeleteSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string** | 스냅샷의 고유 ID - [List snapshots](/openapi/bcs/list-snapshots)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSnapshot

> GetSnapshotResponse GetSnapshot(ctx, snapshotId).Execute()

Get snapshot



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/volume"
)

func main() {
	snapshotId := "snapshotId_example" // string | 스냅샷의 고유 ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnapshotAPI.GetSnapshot(context.Background(), snapshotId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotAPI.GetSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSnapshot`: GetSnapshotResponse
	fmt.Fprintf(os.Stdout, "Response from `SnapshotAPI.GetSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string** | 스냅샷의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetSnapshotResponse**](GetSnapshotResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSnapshots

> ListSnapshotsResponse ListSnapshots(ctx).Id(id).IsIncremental(isIncremental).Name(name).Status(status).VolumeId(volumeId).IsDependentSnapshot(isDependentSnapshot).ScheduleId(scheduleId).ParentId(parentId).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Offset(offset).Limit(limit).Execute()

List snapshots



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/volume"
)

func main() {
	id := "id_example" // string | 스냅샷의 ID (optional)
	isIncremental := true // bool | 증분 스냅샷인지 여부 - `true`: 이전 스냅샷 이후 변경된 데이터만 포함 - `false`: 전체 데이터를 포함하는 스냅샷 (optional)
	name := "name_example" // string | 스냅샷의 이름 (optional)
	status := "status_example" // string | 스냅샷의 현재 상태 (optional)
	volumeId := "volumeId_example" // string | 볼륨의 고유 ID (optional)
	isDependentSnapshot := true // bool | 해당 스냅샷이 다른 스냅샷에 의존하는지 여부 - `true`: 부모 스냅샷 없이 단독 복원이 불가능 - `false`: 독립적으로 복원이 가능한 스냅샷 (optional)
	scheduleId := "scheduleId_example" // string | 스냅샷 생성 시 사용한 스케줄의 ID (optional)
	parentId := "parentId_example" // string | 부모 스냅샷의 ID - 증분 스냅샷인 경우 해당 필드로 부모 스냅샷을 추적 가능 (optional)
	createdAt := "createdAt_example" // string | 리소스가 생성된 시간 - ISO 8601 형식 - UTC 기준 (optional)
	updatedAt := "updatedAt_example" // string | 리소스가 마지막으로 수정된 시간 - ISO 8601 형식 - UTC 기준 (optional)
	sortKeys := "sortKeys_example" // string | 정렬할 필드를 콤마(,)로 구분 (optional)
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`) (optional)
	offset := int32(56) // int32 | 조회 시작 위치 (optional)
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnapshotAPI.ListSnapshots(context.Background()).Id(id).IsIncremental(isIncremental).Name(name).Status(status).VolumeId(volumeId).IsDependentSnapshot(isDependentSnapshot).ScheduleId(scheduleId).ParentId(parentId).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotAPI.ListSnapshots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSnapshots`: ListSnapshotsResponse
	fmt.Fprintf(os.Stdout, "Response from `SnapshotAPI.ListSnapshots`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | 스냅샷의 ID | 
 **isIncremental** | **bool** | 증분 스냅샷인지 여부 - &#x60;true&#x60;: 이전 스냅샷 이후 변경된 데이터만 포함 - &#x60;false&#x60;: 전체 데이터를 포함하는 스냅샷 | 
 **name** | **string** | 스냅샷의 이름 | 
 **status** | **string** | 스냅샷의 현재 상태 | 
 **volumeId** | **string** | 볼륨의 고유 ID | 
 **isDependentSnapshot** | **bool** | 해당 스냅샷이 다른 스냅샷에 의존하는지 여부 - &#x60;true&#x60;: 부모 스냅샷 없이 단독 복원이 불가능 - &#x60;false&#x60;: 독립적으로 복원이 가능한 스냅샷 | 
 **scheduleId** | **string** | 스냅샷 생성 시 사용한 스케줄의 ID | 
 **parentId** | **string** | 부모 스냅샷의 ID - 증분 스냅샷인 경우 해당 필드로 부모 스냅샷을 추적 가능 | 
 **createdAt** | **string** | 리소스가 생성된 시간 - ISO 8601 형식 - UTC 기준 | 
 **updatedAt** | **string** | 리소스가 마지막으로 수정된 시간 - ISO 8601 형식 - UTC 기준 | 
 **sortKeys** | **string** | 정렬할 필드를 콤마(,)로 구분 | 
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;) | 
 **offset** | **int32** | 조회 시작 위치 | 
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | 

### Return type

[**ListSnapshotsResponse**](ListSnapshotsResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreSnapshot

> RestoreSnapshotResponse RestoreSnapshot(ctx, snapshotId).RestoreSnapshotRequest(restoreSnapshotRequest).Execute()

Restore snapshot



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/volume"
)

func main() {
	snapshotId := "snapshotId_example" // string | 스냅샷의 고유 ID - [List snapshots](/openapi/bcs/list-snapshots)에서 확인
	restoreSnapshotRequest := *openapiclient.NewRestoreSnapshotRequest(*openapiclient.NewRestoreSnapshot("Name_example", openapiclient.AvailabilityZone("kr-central-2-a"))) // RestoreSnapshotRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnapshotAPI.RestoreSnapshot(context.Background(), snapshotId).RestoreSnapshotRequest(restoreSnapshotRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotAPI.RestoreSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestoreSnapshot`: RestoreSnapshotResponse
	fmt.Fprintf(os.Stdout, "Response from `SnapshotAPI.RestoreSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string** | 스냅샷의 고유 ID - [List snapshots](/openapi/bcs/list-snapshots)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **restoreSnapshotRequest** | [**RestoreSnapshotRequest**](RestoreSnapshotRequest.md) |  | 

### Return type

[**RestoreSnapshotResponse**](RestoreSnapshotResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSnapshot

> UpdateSnapshotResponse UpdateSnapshot(ctx, snapshotId).UpdateSnapshotRequest(updateSnapshotRequest).Execute()

Update snapshot



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/volume"
)

func main() {
	snapshotId := "snapshotId_example" // string | 스냅샷의 고유 ID - [List snapshots](/openapi/bcs/list-snapshots)에서 확인
	updateSnapshotRequest := *openapiclient.NewUpdateSnapshotRequest(*openapiclient.NewUpdateSnapshot()) // UpdateSnapshotRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SnapshotAPI.UpdateSnapshot(context.Background(), snapshotId).UpdateSnapshotRequest(updateSnapshotRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SnapshotAPI.UpdateSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSnapshot`: UpdateSnapshotResponse
	fmt.Fprintf(os.Stdout, "Response from `SnapshotAPI.UpdateSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string** | 스냅샷의 고유 ID - [List snapshots](/openapi/bcs/list-snapshots)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSnapshotRequest** | [**UpdateSnapshotRequest**](UpdateSnapshotRequest.md) |  | 

### Return type

[**UpdateSnapshotResponse**](UpdateSnapshotResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

