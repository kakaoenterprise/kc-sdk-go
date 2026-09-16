# \VolumeAPI

All URIs are relative to *https://volume.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVolume**](VolumeAPI.md#CreateVolume) | **Post** /api/v1/volumes | Create volume
[**DeleteVolume**](VolumeAPI.md#DeleteVolume) | **Delete** /api/v1/volumes/{volume_id} | Delete volume
[**ExtendVolume**](VolumeAPI.md#ExtendVolume) | **Post** /api/v1/volumes/{volume_id}/size | Extend volume
[**GetVolume**](VolumeAPI.md#GetVolume) | **Get** /api/v1/volumes/{volume_id} | Get volume
[**ListVolumeTypes**](VolumeAPI.md#ListVolumeTypes) | **Get** /api/v1/volumes/types | List volume types
[**ListVolumes**](VolumeAPI.md#ListVolumes) | **Get** /api/v1/volumes | List volumes
[**UpdateVolume**](VolumeAPI.md#UpdateVolume) | **Put** /api/v1/volumes/{volume_id} | Update volume



## CreateVolume

> CreateVolumeResponse CreateVolume(ctx).CreateVolumeRequest(createVolumeRequest).Execute()

Create volume



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
	createVolumeRequest := *openapiclient.NewCreateVolumeRequest(*openapiclient.NewCreateVolume("Name_example", int32(123), openapiclient.AvailabilityZone("kr-central-2-a"))) // CreateVolumeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeAPI.CreateVolume(context.Background()).CreateVolumeRequest(createVolumeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.CreateVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVolume`: CreateVolumeResponse
	fmt.Fprintf(os.Stdout, "Response from `VolumeAPI.CreateVolume`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createVolumeRequest** | [**CreateVolumeRequest**](CreateVolumeRequest.md) |  | 

### Return type

[**CreateVolumeResponse**](CreateVolumeResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVolume

> DeleteVolume(ctx, volumeId).Execute()

Delete volume



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VolumeAPI.DeleteVolume(context.Background(), volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.DeleteVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | 볼륨의 고유 ID - [List volumes](/openapi/bcs/list-volumes)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVolumeRequest struct via the builder pattern


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


## ExtendVolume

> ExtendVolume(ctx, volumeId).ExtendVolumeRequest(extendVolumeRequest).Execute()

Extend volume



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
	extendVolumeRequest := *openapiclient.NewExtendVolumeRequest(*openapiclient.NewExtendVolume(int32(123))) // ExtendVolumeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VolumeAPI.ExtendVolume(context.Background(), volumeId).ExtendVolumeRequest(extendVolumeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.ExtendVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | 볼륨의 고유 ID - [List volumes](/openapi/bcs/list-volumes)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiExtendVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **extendVolumeRequest** | [**ExtendVolumeRequest**](ExtendVolumeRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVolume

> GetVolumeResponse GetVolume(ctx, volumeId).Execute()

Get volume



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
	volumeId := "volumeId_example" // string | 볼륨의 고유 ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeAPI.GetVolume(context.Background(), volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.GetVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVolume`: GetVolumeResponse
	fmt.Fprintf(os.Stdout, "Response from `VolumeAPI.GetVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | 볼륨의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetVolumeResponse**](GetVolumeResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVolumeTypes

> ListVolumeTypesResponse ListVolumeTypes(ctx).Execute()

List volume types



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeAPI.ListVolumeTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.ListVolumeTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVolumeTypes`: ListVolumeTypesResponse
	fmt.Fprintf(os.Stdout, "Response from `VolumeAPI.ListVolumeTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVolumeTypesRequest struct via the builder pattern


### Return type

[**ListVolumeTypesResponse**](ListVolumeTypesResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVolumes

> ListVolumesResponse ListVolumes(ctx).Name(name).Id(id).Status(status).InstanceId(instanceId).MountPoint(mountPoint).Type_(type_).Size(size).AvailabilityZone(availabilityZone).InstanceName(instanceName).VolumeType(volumeType).AttachStatus(attachStatus).IsBootable(isBootable).IsEncrypted(isEncrypted).IsRoot(isRoot).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Offset(offset).Limit(limit).Execute()

List volumes



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
	name := "name_example" // string | 조회할 볼륨 이름 (optional)
	id := "id_example" // string | 조회할 볼륨 ID (optional)
	status := "status_example" // string | 조회할 볼륨 상태 - [볼륨 상태값](https://docs.kakaocloud.com/service/bcs/vm/vm-main#volume-states) (optional)
	instanceId := "instanceId_example" // string | 인스턴스의 고유 ID (optional)
	mountPoint := "mountPoint_example" // string | 볼륨이 인스턴스에 마운트된 경로 - 예시: `/dev/vdb` (optional)
	type_ := "type__example" // string | 볼륨 저장소 유형 (optional)
	size := int32(56) // int32 | 조회할 볼륨 크기 (GB 단위) (optional)
	availabilityZone := openapiclient.AvailabilityZone("kr-central-2-a") // AvailabilityZone | 볼륨이 생성된 가용 영역 (optional)
	instanceName := "instanceName_example" // string | 연결된 인스턴스 이름 (optional)
	volumeType := "volumeType_example" // string | 볼륨 유형 이름 (optional)
	attachStatus := "attachStatus_example" // string | 볼륨 연결 상태 - 예시: `attached`, `detached` (optional)
	isBootable := true // bool | 부팅 가능 여부 (optional)
	isEncrypted := true // bool | 암호화 여부 (optional)
	isRoot := true // bool | 볼륨이 루트 디스크인지 여부 (optional)
	createdAt := "createdAt_example" // string | 리소스가 생성된 시간 - ISO 8601 형식 - UTC 기준 (optional)
	updatedAt := "updatedAt_example" // string | 리소스가 마지막으로 수정된 시간 - ISO 8601 형식 - UTC 기준 (optional)
	sortKeys := "sortKeys_example" // string | 정렬할 필드를 콤마(,)로 구분 (optional)
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`) (optional)
	offset := int32(56) // int32 | 조회 시작 위치 (optional)
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeAPI.ListVolumes(context.Background()).Name(name).Id(id).Status(status).InstanceId(instanceId).MountPoint(mountPoint).Type_(type_).Size(size).AvailabilityZone(availabilityZone).InstanceName(instanceName).VolumeType(volumeType).AttachStatus(attachStatus).IsBootable(isBootable).IsEncrypted(isEncrypted).IsRoot(isRoot).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.ListVolumes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVolumes`: ListVolumesResponse
	fmt.Fprintf(os.Stdout, "Response from `VolumeAPI.ListVolumes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVolumesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | 조회할 볼륨 이름 | 
 **id** | **string** | 조회할 볼륨 ID | 
 **status** | **string** | 조회할 볼륨 상태 - [볼륨 상태값](https://docs.kakaocloud.com/service/bcs/vm/vm-main#volume-states) | 
 **instanceId** | **string** | 인스턴스의 고유 ID | 
 **mountPoint** | **string** | 볼륨이 인스턴스에 마운트된 경로 - 예시: &#x60;/dev/vdb&#x60; | 
 **type_** | **string** | 볼륨 저장소 유형 | 
 **size** | **int32** | 조회할 볼륨 크기 (GB 단위) | 
 **availabilityZone** | [**AvailabilityZone**](AvailabilityZone.md) | 볼륨이 생성된 가용 영역 | 
 **instanceName** | **string** | 연결된 인스턴스 이름 | 
 **volumeType** | **string** | 볼륨 유형 이름 | 
 **attachStatus** | **string** | 볼륨 연결 상태 - 예시: &#x60;attached&#x60;, &#x60;detached&#x60; | 
 **isBootable** | **bool** | 부팅 가능 여부 | 
 **isEncrypted** | **bool** | 암호화 여부 | 
 **isRoot** | **bool** | 볼륨이 루트 디스크인지 여부 | 
 **createdAt** | **string** | 리소스가 생성된 시간 - ISO 8601 형식 - UTC 기준 | 
 **updatedAt** | **string** | 리소스가 마지막으로 수정된 시간 - ISO 8601 형식 - UTC 기준 | 
 **sortKeys** | **string** | 정렬할 필드를 콤마(,)로 구분 | 
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;) | 
 **offset** | **int32** | 조회 시작 위치 | 
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | 

### Return type

[**ListVolumesResponse**](ListVolumesResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVolume

> UpdateVolumeResponse UpdateVolume(ctx, volumeId).UpdateVolumeRequest(updateVolumeRequest).Execute()

Update volume



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
	updateVolumeRequest := *openapiclient.NewUpdateVolumeRequest(*openapiclient.NewUpdateVolume()) // UpdateVolumeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeAPI.UpdateVolume(context.Background(), volumeId).UpdateVolumeRequest(updateVolumeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.UpdateVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateVolume`: UpdateVolumeResponse
	fmt.Fprintf(os.Stdout, "Response from `VolumeAPI.UpdateVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | 볼륨의 고유 ID - [List volumes](/openapi/bcs/list-volumes)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateVolumeRequest** | [**UpdateVolumeRequest**](UpdateVolumeRequest.md) |  | 

### Return type

[**UpdateVolumeResponse**](UpdateVolumeResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

