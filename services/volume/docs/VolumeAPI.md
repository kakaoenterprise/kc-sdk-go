# \VolumeAPI

All URIs are relative to *https://volume.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateImage**](VolumeAPI.md#CreateImage) | **Post** /api/v1/volumes/{volume_id}/image | Create image
[**CreateSnapshot**](VolumeAPI.md#CreateSnapshot) | **Post** /api/v1/volumes/{volume_id}/snapshots | Create snapshot
[**CreateVolume**](VolumeAPI.md#CreateVolume) | **Post** /api/v1/volumes | Create volume
[**DeleteVolume**](VolumeAPI.md#DeleteVolume) | **Delete** /api/v1/volumes/{volume_id} | Delete volume
[**ExtendVolume**](VolumeAPI.md#ExtendVolume) | **Post** /api/v1/volumes/{volume_id}/size | Extend volume
[**GetVolume**](VolumeAPI.md#GetVolume) | **Get** /api/v1/volumes/{volume_id} | Get volume
[**ListVolumeTypes**](VolumeAPI.md#ListVolumeTypes) | **Get** /api/v1/volumes/types | List volume types
[**ListVolumes**](VolumeAPI.md#ListVolumes) | **Get** /api/v1/volumes | List volumes
[**UpdateVolume**](VolumeAPI.md#UpdateVolume) | **Put** /api/v1/volumes/{volume_id} | Update volume



## CreateImage

> ResponseVolumeImageModel CreateImage(ctx, volumeId).XAuthToken(xAuthToken).BodyCreateImage(bodyCreateImage).Execute()

Create image



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
	volumeId := "volumeId_example" // string | 볼륨의 고유 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bodyCreateImage := *openapiclient.NewBodyCreateImage(*openapiclient.NewCreateVolumeImageModel("Name_example")) // BodyCreateImage | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeAPI.CreateImage(context.Background(), volumeId).XAuthToken(xAuthToken).BodyCreateImage(bodyCreateImage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.CreateImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateImage`: ResponseVolumeImageModel
	fmt.Fprintf(os.Stdout, "Response from `VolumeAPI.CreateImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | 볼륨의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyCreateImage** | [**BodyCreateImage**](BodyCreateImage.md) |  | 

### Return type

[**ResponseVolumeImageModel**](ResponseVolumeImageModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSnapshot

> ResponseCreateVolumeSnapshotModel CreateSnapshot(ctx, volumeId).XAuthToken(xAuthToken).BodyCreateSnapshot(bodyCreateSnapshot).Execute()

Create snapshot



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
	volumeId := "volumeId_example" // string | 볼륨의 고유 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bodyCreateSnapshot := *openapiclient.NewBodyCreateSnapshot(*openapiclient.NewCreateVolumeSnapshotModel("Name_example", false)) // BodyCreateSnapshot | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeAPI.CreateSnapshot(context.Background(), volumeId).XAuthToken(xAuthToken).BodyCreateSnapshot(bodyCreateSnapshot).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.CreateSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSnapshot`: ResponseCreateVolumeSnapshotModel
	fmt.Fprintf(os.Stdout, "Response from `VolumeAPI.CreateSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | 볼륨의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyCreateSnapshot** | [**BodyCreateSnapshot**](BodyCreateSnapshot.md) |  | 

### Return type

[**ResponseCreateVolumeSnapshotModel**](ResponseCreateVolumeSnapshotModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVolume

> BcsVolumeV1ApiCreateVolumeModelResponseVolumeModel CreateVolume(ctx).XAuthToken(xAuthToken).BodyCreateVolume(bodyCreateVolume).Execute()

Create volume



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
	bodyCreateVolume := *openapiclient.NewBodyCreateVolume(*openapiclient.NewCreateVolumeModel("Name_example", int32(123), openapiclient.AvailabilityZone("kr-central-2-a"))) // BodyCreateVolume | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeAPI.CreateVolume(context.Background()).XAuthToken(xAuthToken).BodyCreateVolume(bodyCreateVolume).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.CreateVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVolume`: BcsVolumeV1ApiCreateVolumeModelResponseVolumeModel
	fmt.Fprintf(os.Stdout, "Response from `VolumeAPI.CreateVolume`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyCreateVolume** | [**BodyCreateVolume**](BodyCreateVolume.md) |  | 

### Return type

[**BcsVolumeV1ApiCreateVolumeModelResponseVolumeModel**](BcsVolumeV1ApiCreateVolumeModelResponseVolumeModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVolume

> interface{} DeleteVolume(ctx, volumeId).XAuthToken(xAuthToken).Execute()

Delete volume



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
	volumeId := "volumeId_example" // string | 볼륨의 고유 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeAPI.DeleteVolume(context.Background(), volumeId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.DeleteVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVolume`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `VolumeAPI.DeleteVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | 볼륨의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVolumeRequest struct via the builder pattern


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


## ExtendVolume

> interface{} ExtendVolume(ctx, volumeId).XAuthToken(xAuthToken).BodyExtendVolume(bodyExtendVolume).Execute()

Extend volume



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
	volumeId := "volumeId_example" // string | 볼륨의 고유 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bodyExtendVolume := *openapiclient.NewBodyExtendVolume(*openapiclient.NewExtendVolumeModel(int32(123))) // BodyExtendVolume | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeAPI.ExtendVolume(context.Background(), volumeId).XAuthToken(xAuthToken).BodyExtendVolume(bodyExtendVolume).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.ExtendVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExtendVolume`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `VolumeAPI.ExtendVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | 볼륨의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiExtendVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyExtendVolume** | [**BodyExtendVolume**](BodyExtendVolume.md) |  | 

### Return type

**interface{}**

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVolume

> BcsVolumeV1ApiGetVolumeModelResponseVolumeModel GetVolume(ctx, volumeId).XAuthToken(xAuthToken).Execute()

Get volume



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
	volumeId := "volumeId_example" // string | 볼륨의 고유 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeAPI.GetVolume(context.Background(), volumeId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.GetVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVolume`: BcsVolumeV1ApiGetVolumeModelResponseVolumeModel
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

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**BcsVolumeV1ApiGetVolumeModelResponseVolumeModel**](BcsVolumeV1ApiGetVolumeModelResponseVolumeModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVolumeTypes

> VolumeTypeListModel ListVolumeTypes(ctx).XAuthToken(xAuthToken).Execute()

List volume types



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeAPI.ListVolumeTypes(context.Background()).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.ListVolumeTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVolumeTypes`: VolumeTypeListModel
	fmt.Fprintf(os.Stdout, "Response from `VolumeAPI.ListVolumeTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVolumeTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**VolumeTypeListModel**](VolumeTypeListModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVolumes

> VolumeListModel ListVolumes(ctx).XAuthToken(xAuthToken).Id(id).Name(name).Status(status).InstanceId(instanceId).MountPoint(mountPoint).Type_(type_).Size(size).AvailabilityZone(availabilityZone).InstanceName(instanceName).VolumeType(volumeType).AttachStatus(attachStatus).IsBootable(isBootable).IsEncrypted(isEncrypted).IsRoot(isRoot).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Limit(limit).Offset(offset).Execute()

List volumes



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
	id := "id_example" // string | 조회할 볼륨 ID  (optional)
	name := "name_example" // string | 조회할 볼륨 이름  (optional)
	status := "status_example" // string | 조회할 볼륨 상태  <br/> - [볼륨 상태값](https://docs.kakaocloud.com/service/bcs/vm/vm-main#volume-states)  (optional)
	instanceId := "instanceId_example" // string | 인스턴스의 고유 ID (optional)
	mountPoint := "mountPoint_example" // string | 볼륨이 인스턴스에 마운트된 경로 <br/> - 예시: `/dev/vdb` (optional)
	type_ := "type__example" // string | 볼륨 저장소 유형 (optional)
	size := int32(56) // int32 | 조회할 볼륨 크기 (GB 단위)  (optional)
	availabilityZone := openapiclient.AvailabilityZone("kr-central-2-a") // AvailabilityZone | 볼륨이 생성된 가용 영역<br/> - `kr-central-2-a`: kr-central-2-a 가용 영역 <br/> - `kr-central-2-b`: kr-central-2-b 가용 영역 <br/> - `kr-central-2-c`: kr-central-2-c 가용 영역 (optional)
	instanceName := "instanceName_example" // string | 연결된 인스턴스 이름 (optional)
	volumeType := "volumeType_example" // string | 볼륨 유형 이름  (optional)
	attachStatus := "attachStatus_example" // string | 볼륨 연결 상태 <br/> - 예시: `attached`, `detached` (optional)
	isBootable := true // bool | 부팅 가능 여부  (optional)
	isEncrypted := true // bool | 암호화 여부  (optional)
	isRoot := true // bool | 볼륨이 루트 디스크인지 여부 (optional)
	createdAt := "createdAt_example" // string | 리소스가 생성된 시간 <br/> - ISO_8601 형식  <br/> - UTC 기준 (optional)
	updatedAt := "updatedAt_example" // string | 리소스가 마지막으로 수정된 시간 <br/> - ISO_8601 형식  <br/> - UTC 기준 (optional)
	sortKeys := "sortKeys_example" // string | 정렬할 필드를 콤마(,)로 구분   (optional) (default to "created_at")
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`)   (optional) (default to "desc")
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional) (default to 20)
	offset := int32(56) // int32 | 조회 시작 위치 (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeAPI.ListVolumes(context.Background()).XAuthToken(xAuthToken).Id(id).Name(name).Status(status).InstanceId(instanceId).MountPoint(mountPoint).Type_(type_).Size(size).AvailabilityZone(availabilityZone).InstanceName(instanceName).VolumeType(volumeType).AttachStatus(attachStatus).IsBootable(isBootable).IsEncrypted(isEncrypted).IsRoot(isRoot).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.ListVolumes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVolumes`: VolumeListModel
	fmt.Fprintf(os.Stdout, "Response from `VolumeAPI.ListVolumes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVolumesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **id** | **string** | 조회할 볼륨 ID  | 
 **name** | **string** | 조회할 볼륨 이름  | 
 **status** | **string** | 조회할 볼륨 상태  &lt;br/&gt; - [볼륨 상태값](https://docs.kakaocloud.com/service/bcs/vm/vm-main#volume-states)  | 
 **instanceId** | **string** | 인스턴스의 고유 ID | 
 **mountPoint** | **string** | 볼륨이 인스턴스에 마운트된 경로 &lt;br/&gt; - 예시: &#x60;/dev/vdb&#x60; | 
 **type_** | **string** | 볼륨 저장소 유형 | 
 **size** | **int32** | 조회할 볼륨 크기 (GB 단위)  | 
 **availabilityZone** | [**AvailabilityZone**](AvailabilityZone.md) | 볼륨이 생성된 가용 영역&lt;br/&gt; - &#x60;kr-central-2-a&#x60;: kr-central-2-a 가용 영역 &lt;br/&gt; - &#x60;kr-central-2-b&#x60;: kr-central-2-b 가용 영역 &lt;br/&gt; - &#x60;kr-central-2-c&#x60;: kr-central-2-c 가용 영역 | 
 **instanceName** | **string** | 연결된 인스턴스 이름 | 
 **volumeType** | **string** | 볼륨 유형 이름  | 
 **attachStatus** | **string** | 볼륨 연결 상태 &lt;br/&gt; - 예시: &#x60;attached&#x60;, &#x60;detached&#x60; | 
 **isBootable** | **bool** | 부팅 가능 여부  | 
 **isEncrypted** | **bool** | 암호화 여부  | 
 **isRoot** | **bool** | 볼륨이 루트 디스크인지 여부 | 
 **createdAt** | **string** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
 **updatedAt** | **string** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
 **sortKeys** | **string** | 정렬할 필드를 콤마(,)로 구분   | [default to &quot;created_at&quot;]
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;)   | [default to &quot;desc&quot;]
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | [default to 20]
 **offset** | **int32** | 조회 시작 위치 | [default to 0]

### Return type

[**VolumeListModel**](VolumeListModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVolume

> BcsVolumeV1ApiUpdateVolumeModelResponseVolumeModel UpdateVolume(ctx, volumeId).XAuthToken(xAuthToken).BodyUpdateVolume(bodyUpdateVolume).Execute()

Update volume



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
	volumeId := "volumeId_example" // string | 볼륨의 고유 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bodyUpdateVolume := *openapiclient.NewBodyUpdateVolume(*openapiclient.NewEditVolumeModel("Name_example")) // BodyUpdateVolume | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumeAPI.UpdateVolume(context.Background(), volumeId).XAuthToken(xAuthToken).BodyUpdateVolume(bodyUpdateVolume).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumeAPI.UpdateVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateVolume`: BcsVolumeV1ApiUpdateVolumeModelResponseVolumeModel
	fmt.Fprintf(os.Stdout, "Response from `VolumeAPI.UpdateVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | 볼륨의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyUpdateVolume** | [**BodyUpdateVolume**](BodyUpdateVolume.md) |  | 

### Return type

[**BcsVolumeV1ApiUpdateVolumeModelResponseVolumeModel**](BcsVolumeV1ApiUpdateVolumeModelResponseVolumeModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

