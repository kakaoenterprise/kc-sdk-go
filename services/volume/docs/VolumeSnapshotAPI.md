# \VolumeSnapshotAPI

All URIs are relative to *https://volume.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSnapshot**](VolumeSnapshotAPI.md#DeleteSnapshot) | **Delete** /api/v1/snapshots/{snapshot_id} | Delete snapshot
[**GetSnapshot**](VolumeSnapshotAPI.md#GetSnapshot) | **Get** /api/v1/snapshots/{snapshot_id} | Get snapshot
[**ListSnapshots**](VolumeSnapshotAPI.md#ListSnapshots) | **Get** /api/v1/snapshots | List snapshots
[**RestoreSnapshot**](VolumeSnapshotAPI.md#RestoreSnapshot) | **Post** /api/v1/snapshots/{snapshot_id}/restore | Restore snapshot
[**UpdateSnapshot**](VolumeSnapshotAPI.md#UpdateSnapshot) | **Put** /api/v1/snapshots/{snapshot_id} | Update snapshot



## DeleteSnapshot

> interface{} DeleteSnapshot(ctx, snapshotId).XAuthToken(xAuthToken).Execute()

Delete snapshot



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	snapshotId := "snapshotId_example" // string | 스냅샷의 고유 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeSnapshotAPI.DeleteSnapshot(context.Background(), snapshotId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeSnapshotAPI.DeleteSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSnapshot`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `VolumeSnapshotAPI.DeleteSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string** | 스냅샷의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

**interface{}**

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSnapshot

> BcsVolumeV1ApiGetSnapshotModelResponseVolumeSnapshotModel GetSnapshot(ctx, snapshotId).XAuthToken(xAuthToken).Execute()

Get snapshot



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	snapshotId := "snapshotId_example" // string | 스냅샷의 고유 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeSnapshotAPI.GetSnapshot(context.Background(), snapshotId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeSnapshotAPI.GetSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSnapshot`: BcsVolumeV1ApiGetSnapshotModelResponseVolumeSnapshotModel
	fmt.Fprintf(os.Stdout, "Response from `VolumeSnapshotAPI.GetSnapshot`: %v\n", resp)
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

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**BcsVolumeV1ApiGetSnapshotModelResponseVolumeSnapshotModel**](BcsVolumeV1ApiGetSnapshotModelResponseVolumeSnapshotModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSnapshots

> VolumeSnapshotListModel ListSnapshots(ctx).XAuthToken(xAuthToken).Id(id).Name(name).Status(status).VolumeId(volumeId).IsIncremental(isIncremental).IsDependentSnapshot(isDependentSnapshot).ScheduleId(scheduleId).ParentId(parentId).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Limit(limit).Offset(offset).Execute()

List snapshots



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	id := "id_example" // string | 스냅샷의 ID (optional)
	name := "name_example" // string | 스냅샷의 이름 (optional)
	status := "status_example" // string | 스냅샷의 현재 상태  (optional)
	volumeId := "volumeId_example" // string | 볼륨의 고유 ID (optional)
	isIncremental := true // bool | 증분 스냅샷인지 여부 <br/>- `true`: 이전 스냅샷 이후 변경된 데이터만 포함  <br/>- `false`: 전체 데이터를 포함하는 스냅샷 (optional)
	isDependentSnapshot := true // bool | 해당 스냅샷이 다른 스냅샷에 의존하는지 여부  <br/>- `true`: 부모 스냅샷 없이 단독 복원이 불가능  <br/>- `false`: 독립적으로 복원이 가능한 스냅샷 (optional)
	scheduleId := "scheduleId_example" // string | 스냅샷 생성 시 사용한 스케줄의 ID  (optional)
	parentId := "parentId_example" // string | 부모 스냅샷의 ID <br/> - 증분 스냅샷인 경우 해당 필드로 부모 스냅샷을 추적 가능 (optional)
	createdAt := "createdAt_example" // string | 리소스가 생성된 시간 <br/> - ISO_8601 형식  <br/> - UTC 기준 (optional)
	updatedAt := "updatedAt_example" // string | 리소스가 마지막으로 수정된 시간 <br/> - ISO_8601 형식  <br/> - UTC 기준 (optional)
	sortKeys := "sortKeys_example" // string | 정렬할 필드를 콤마(,)로 구분   (optional) (default to "created_at")
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`)   (optional) (default to "desc")
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional) (default to 20)
	offset := int32(56) // int32 | 조회 시작 위치 (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeSnapshotAPI.ListSnapshots(context.Background()).XAuthToken(xAuthToken).Id(id).Name(name).Status(status).VolumeId(volumeId).IsIncremental(isIncremental).IsDependentSnapshot(isDependentSnapshot).ScheduleId(scheduleId).ParentId(parentId).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeSnapshotAPI.ListSnapshots``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSnapshots`: VolumeSnapshotListModel
	fmt.Fprintf(os.Stdout, "Response from `VolumeSnapshotAPI.ListSnapshots`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **id** | **string** | 스냅샷의 ID | 
 **name** | **string** | 스냅샷의 이름 | 
 **status** | **string** | 스냅샷의 현재 상태  | 
 **volumeId** | **string** | 볼륨의 고유 ID | 
 **isIncremental** | **bool** | 증분 스냅샷인지 여부 &lt;br/&gt;- &#x60;true&#x60;: 이전 스냅샷 이후 변경된 데이터만 포함  &lt;br/&gt;- &#x60;false&#x60;: 전체 데이터를 포함하는 스냅샷 | 
 **isDependentSnapshot** | **bool** | 해당 스냅샷이 다른 스냅샷에 의존하는지 여부  &lt;br/&gt;- &#x60;true&#x60;: 부모 스냅샷 없이 단독 복원이 불가능  &lt;br/&gt;- &#x60;false&#x60;: 독립적으로 복원이 가능한 스냅샷 | 
 **scheduleId** | **string** | 스냅샷 생성 시 사용한 스케줄의 ID  | 
 **parentId** | **string** | 부모 스냅샷의 ID &lt;br/&gt; - 증분 스냅샷인 경우 해당 필드로 부모 스냅샷을 추적 가능 | 
 **createdAt** | **string** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
 **updatedAt** | **string** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
 **sortKeys** | **string** | 정렬할 필드를 콤마(,)로 구분   | [default to &quot;created_at&quot;]
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;)   | [default to &quot;desc&quot;]
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | [default to 20]
 **offset** | **int32** | 조회 시작 위치 | [default to 0]

### Return type

[**VolumeSnapshotListModel**](VolumeSnapshotListModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreSnapshot

> ResponseRestoreVolumeSnapshotModel RestoreSnapshot(ctx, snapshotId).XAuthToken(xAuthToken).BodyRestoreSnapshot(bodyRestoreSnapshot).Execute()

Restore snapshot



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	snapshotId := "snapshotId_example" // string | 스냅샷의 고유 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bodyRestoreSnapshot := *openapiclient.NewBodyRestoreSnapshot(*openapiclient.NewRequestRestoreVolumeSnapshotModel("Name_example", openapiclient.AvailabilityZone("kr-central-2-a"))) // BodyRestoreSnapshot | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeSnapshotAPI.RestoreSnapshot(context.Background(), snapshotId).XAuthToken(xAuthToken).BodyRestoreSnapshot(bodyRestoreSnapshot).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeSnapshotAPI.RestoreSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestoreSnapshot`: ResponseRestoreVolumeSnapshotModel
	fmt.Fprintf(os.Stdout, "Response from `VolumeSnapshotAPI.RestoreSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string** | 스냅샷의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyRestoreSnapshot** | [**BodyRestoreSnapshot**](BodyRestoreSnapshot.md) |  | 

### Return type

[**ResponseRestoreVolumeSnapshotModel**](ResponseRestoreVolumeSnapshotModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSnapshot

> BcsVolumeV1ApiUpdateSnapshotModelResponseVolumeSnapshotModel UpdateSnapshot(ctx, snapshotId).XAuthToken(xAuthToken).BodyUpdateSnapshot(bodyUpdateSnapshot).Execute()

Update snapshot



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	snapshotId := "snapshotId_example" // string | 스냅샷의 고유 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bodyUpdateSnapshot := *openapiclient.NewBodyUpdateSnapshot(*openapiclient.NewEditVolumeSnapshotModel()) // BodyUpdateSnapshot | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeSnapshotAPI.UpdateSnapshot(context.Background(), snapshotId).XAuthToken(xAuthToken).BodyUpdateSnapshot(bodyUpdateSnapshot).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeSnapshotAPI.UpdateSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSnapshot`: BcsVolumeV1ApiUpdateSnapshotModelResponseVolumeSnapshotModel
	fmt.Fprintf(os.Stdout, "Response from `VolumeSnapshotAPI.UpdateSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**snapshotId** | **string** | 스냅샷의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyUpdateSnapshot** | [**BodyUpdateSnapshot**](BodyUpdateSnapshot.md) |  | 

### Return type

[**BcsVolumeV1ApiUpdateSnapshotModelResponseVolumeSnapshotModel**](BcsVolumeV1ApiUpdateSnapshotModelResponseVolumeSnapshotModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

