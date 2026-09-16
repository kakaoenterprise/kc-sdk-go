# \ImageAPI

All URIs are relative to *https://image.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddImageShare**](ImageAPI.md#AddImageShare) | **Post** /api/v1/images/{image_id}/members/{member_id} | Add image share
[**DeleteImage**](ImageAPI.md#DeleteImage) | **Delete** /api/v1/images/{image_id} | Delete image
[**GetImage**](ImageAPI.md#GetImage) | **Get** /api/v1/images/{image_id} | Get image
[**ListImageSharedProjects**](ImageAPI.md#ListImageSharedProjects) | **Get** /api/v1/images/{image_id}/members | List image shared projects
[**ListImages**](ImageAPI.md#ListImages) | **Get** /api/v1/images | List images
[**RemoveImageShare**](ImageAPI.md#RemoveImageShare) | **Delete** /api/v1/images/{image_id}/members/{member_id} | Remove image share
[**UpdateImage**](ImageAPI.md#UpdateImage) | **Put** /api/v1/images/{image_id} | Update image



## AddImageShare

> AddImageShareResponse AddImageShare(ctx, imageId, memberId).Execute()

Add image share



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/image"
)

func main() {
	imageId := "imageId_example" // string | 이미지의 고유 ID - [List images](/openapi/bcs/list-images)에서 확인
	memberId := "memberId_example" // string | 이미지를 공유할 다른 프로젝트의 ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.AddImageShare(context.Background(), imageId, memberId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.AddImageShare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddImageShare`: AddImageShareResponse
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.AddImageShare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string** | 이미지의 고유 ID - [List images](/openapi/bcs/list-images)에서 확인 | 
**memberId** | **string** | 이미지를 공유할 다른 프로젝트의 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddImageShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AddImageShareResponse**](AddImageShareResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteImage

> DeleteImage(ctx, imageId).Execute()

Delete image



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/image"
)

func main() {
	imageId := "imageId_example" // string | 이미지의 고유 ID - [List images](/openapi/bcs/list-images)에서 확인

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ImageAPI.DeleteImage(context.Background(), imageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.DeleteImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string** | 이미지의 고유 ID - [List images](/openapi/bcs/list-images)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteImageRequest struct via the builder pattern


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


## GetImage

> GetImageResponse GetImage(ctx, imageId).Execute()

Get image



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/image"
)

func main() {
	imageId := "imageId_example" // string | 이미지의 고유 ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.GetImage(context.Background(), imageId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.GetImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImage`: GetImageResponse
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.GetImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string** | 이미지의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetImageResponse**](GetImageResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListImageSharedProjects

> ListImageSharedProjectsResponse ListImageSharedProjects(ctx, imageId).Offset(offset).Limit(limit).Execute()

List image shared projects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/image"
)

func main() {
	imageId := "imageId_example" // string | 이미지의 고유 ID
	offset := int32(56) // int32 | 조회 시작 위치 (optional)
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.ListImageSharedProjects(context.Background(), imageId).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.ListImageSharedProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListImageSharedProjects`: ListImageSharedProjectsResponse
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.ListImageSharedProjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string** | 이미지의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListImageSharedProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | 조회 시작 위치 | 
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | 

### Return type

[**ListImageSharedProjectsResponse**](ListImageSharedProjectsResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListImages

> ListImagesResponse ListImages(ctx).InstanceType(instanceType).ImageType(imageType).OsType(osType).Name(name).Id(id).Size(size).MinDisk(minDisk).DiskFormat(diskFormat).Status(status).Visibility(visibility).ImageMemberStatus(imageMemberStatus).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Offset(offset).Limit(limit).Execute()

List images



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/image"
)

func main() {
	instanceType := openapiclient.ImageInstanceType("vm") // ImageInstanceType | 이미지와 호환 가능한 인스턴스 유형 (optional)
	imageType := openapiclient.ImageType("basic") // ImageType | 이미지 제공 유형 (optional)
	osType := "osType_example" // string | 운영체제 유형 (optional)
	name := "name_example" // string | 이미지 이름 (optional)
	id := "id_example" // string | 이미지의 고유 ID (optional)
	size := int64(789) // int64 | 이미지 크기 (bytes 단위) (optional)
	minDisk := int32(56) // int32 | 이미지를 사용할 때 필요한 최소 디스크 크기(GB) (optional)
	diskFormat := "diskFormat_example" // string | 이미지의 디스크 포맷 (optional)
	status := "status_example" // string | 이미지의 상태 (optional)
	visibility := "visibility_example" // string | 이미지의 가시성 (optional)
	imageMemberStatus := "imageMemberStatus_example" // string | 공유된 이미지의 멤버 상태 (optional)
	createdAt := "createdAt_example" // string | 리소스가 생성된 시간 - ISO 8601 형식 - UTC 기준 (optional)
	updatedAt := "updatedAt_example" // string | 리소스가 마지막으로 수정된 시간 - ISO 8601 형식 - UTC 기준 (optional)
	sortKeys := "sortKeys_example" // string | 정렬할 필드를 콤마(,)로 구분 (optional)
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`) (optional)
	offset := int32(56) // int32 | 조회 시작 위치 (optional)
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.ListImages(context.Background()).InstanceType(instanceType).ImageType(imageType).OsType(osType).Name(name).Id(id).Size(size).MinDisk(minDisk).DiskFormat(diskFormat).Status(status).Visibility(visibility).ImageMemberStatus(imageMemberStatus).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.ListImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListImages`: ListImagesResponse
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.ListImages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **instanceType** | [**ImageInstanceType**](ImageInstanceType.md) | 이미지와 호환 가능한 인스턴스 유형 | 
 **imageType** | [**ImageType**](ImageType.md) | 이미지 제공 유형 | 
 **osType** | **string** | 운영체제 유형 | 
 **name** | **string** | 이미지 이름 | 
 **id** | **string** | 이미지의 고유 ID | 
 **size** | **int64** | 이미지 크기 (bytes 단위) | 
 **minDisk** | **int32** | 이미지를 사용할 때 필요한 최소 디스크 크기(GB) | 
 **diskFormat** | **string** | 이미지의 디스크 포맷 | 
 **status** | **string** | 이미지의 상태 | 
 **visibility** | **string** | 이미지의 가시성 | 
 **imageMemberStatus** | **string** | 공유된 이미지의 멤버 상태 | 
 **createdAt** | **string** | 리소스가 생성된 시간 - ISO 8601 형식 - UTC 기준 | 
 **updatedAt** | **string** | 리소스가 마지막으로 수정된 시간 - ISO 8601 형식 - UTC 기준 | 
 **sortKeys** | **string** | 정렬할 필드를 콤마(,)로 구분 | 
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;) | 
 **offset** | **int32** | 조회 시작 위치 | 
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | 

### Return type

[**ListImagesResponse**](ListImagesResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveImageShare

> RemoveImageShare(ctx, imageId, memberId).Execute()

Remove image share



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/image"
)

func main() {
	imageId := "imageId_example" // string | 이미지의 고유 ID - [List images](/openapi/bcs/list-images)에서 확인
	memberId := "memberId_example" // string | 공유 대상 프로젝트 ID - [List image shared projects](/openapi/bcs/list-image-shared-projects)에서 확인

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ImageAPI.RemoveImageShare(context.Background(), imageId, memberId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.RemoveImageShare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string** | 이미지의 고유 ID - [List images](/openapi/bcs/list-images)에서 확인 | 
**memberId** | **string** | 공유 대상 프로젝트 ID - [List image shared projects](/openapi/bcs/list-image-shared-projects)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveImageShareRequest struct via the builder pattern


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


## UpdateImage

> UpdateImageResponse UpdateImage(ctx, imageId).UpdateImageRequest(updateImageRequest).Execute()

Update image



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/image"
)

func main() {
	imageId := "imageId_example" // string | 이미지의 고유 ID - [List images](/openapi/bcs/list-images)에서 확인
	updateImageRequest := *openapiclient.NewUpdateImageRequest(*openapiclient.NewUpdateImage()) // UpdateImageRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.UpdateImage(context.Background(), imageId).UpdateImageRequest(updateImageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.UpdateImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateImage`: UpdateImageResponse
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.UpdateImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string** | 이미지의 고유 ID - [List images](/openapi/bcs/list-images)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateImageRequest** | [**UpdateImageRequest**](UpdateImageRequest.md) |  | 

### Return type

[**UpdateImageResponse**](UpdateImageResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

