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

> ResponseImageMemberModel AddImageShare(ctx, imageId, memberId).XAuthToken(xAuthToken).Execute()

Add image share



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
	imageId := "imageId_example" // string | 이미지의 고유 ID
	memberId := "memberId_example" // string | 이미지를 공유할 다른 프로젝트의 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.AddImageShare(context.Background(), imageId, memberId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.AddImageShare``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddImageShare`: ResponseImageMemberModel
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.AddImageShare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string** | 이미지의 고유 ID | 
**memberId** | **string** | 이미지를 공유할 다른 프로젝트의 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddImageShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**ResponseImageMemberModel**](ResponseImageMemberModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteImage

> DeleteImage(ctx, imageId).XAuthToken(xAuthToken).Execute()

Delete image



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
	imageId := "imageId_example" // string | 이미지의 고유 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ImageAPI.DeleteImage(context.Background(), imageId).XAuthToken(xAuthToken).Execute()
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
**imageId** | **string** | 이미지의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

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

> BcsImageV1ApiGetImageModelResponseImageModel GetImage(ctx, imageId).XAuthToken(xAuthToken).Execute()

Get image



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
	imageId := "imageId_example" // string | 이미지의 고유 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.GetImage(context.Background(), imageId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.GetImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetImage`: BcsImageV1ApiGetImageModelResponseImageModel
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

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**BcsImageV1ApiGetImageModelResponseImageModel**](BcsImageV1ApiGetImageModelResponseImageModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListImageSharedProjects

> ImageMemberListModel ListImageSharedProjects(ctx, imageId).XAuthToken(xAuthToken).Limit(limit).Offset(offset).Execute()

List image shared projects



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
	imageId := "imageId_example" // string | 이미지의 고유 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional) (default to 20)
	offset := int32(56) // int32 | 조회 시작 위치 (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.ListImageSharedProjects(context.Background(), imageId).XAuthToken(xAuthToken).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.ListImageSharedProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListImageSharedProjects`: ImageMemberListModel
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

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | [default to 20]
 **offset** | **int32** | 조회 시작 위치 | [default to 0]

### Return type

[**ImageMemberListModel**](ImageMemberListModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListImages

> ImageListModel ListImages(ctx).XAuthToken(xAuthToken).Id(id).Name(name).InstanceType(instanceType).ImageType(imageType).Size(size).MinDisk(minDisk).DiskFormat(diskFormat).Status(status).OsType(osType).Visibility(visibility).ImageMemberStatus(imageMemberStatus).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Limit(limit).Offset(offset).Execute()

List images



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
	id := "id_example" // string | 이미지의 고유 ID (optional)
	name := "name_example" // string | 이미지 이름 (optional)
	instanceType := openapiclient.ImageInstanceType("vm") // ImageInstanceType | 이미지와 호환 가능한 인스턴스 유형 <br/> - `vm`: Virtual Machine 유형  <br/> - `bm`: Bare Metal Server 유형 (optional)
	imageType := openapiclient.ImageVisibilityType("basic") // ImageVisibilityType | 이미지 제공 유형<br/> - `basic`: 카카오클라우드에서 기본적으로 제공하는 [기본 이미지](https://docs.kakaocloud.com/service/bcs/vm/vm-main#default-images-available)로 Linux 및 Windows 이미지를 제공 <br/> - `my`: 프로젝트 사용자가 생성한 커스텀 이미지 (optional)
	size := int64(789) // int64 | 이미지 크기 (bytes 단위) (optional)
	minDisk := int32(56) // int32 | 이미지를 사용할 때 필요한 최소 디스크 크기(GB) (optional)
	diskFormat := "diskFormat_example" // string | 이미지의 디스크 포맷 (optional)
	status := "status_example" // string | 이미지의 상태 (optional)
	osType := "osType_example" // string | 운영체제 유형 (optional)
	visibility := "visibility_example" // string | 이미지의 가시성 (optional)
	imageMemberStatus := "imageMemberStatus_example" // string | 공유된 이미지의 멤버 상태 (optional)
	createdAt := "createdAt_example" // string | 리소스가 생성된 시간 <br/> - ISO_8601 형식  <br/> - UTC 기준 (optional)
	updatedAt := "updatedAt_example" // string | 리소스가 마지막으로 수정된 시간 <br/> - ISO_8601 형식  <br/> - UTC 기준 (optional)
	sortKeys := "sortKeys_example" // string | 정렬할 필드를 콤마(,)로 구분   (optional) (default to "created_at")
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`)   (optional) (default to "desc")
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional) (default to 20)
	offset := int32(56) // int32 | 조회 시작 위치 (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.ListImages(context.Background()).XAuthToken(xAuthToken).Id(id).Name(name).InstanceType(instanceType).ImageType(imageType).Size(size).MinDisk(minDisk).DiskFormat(diskFormat).Status(status).OsType(osType).Visibility(visibility).ImageMemberStatus(imageMemberStatus).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.ListImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListImages`: ImageListModel
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.ListImages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **id** | **string** | 이미지의 고유 ID | 
 **name** | **string** | 이미지 이름 | 
 **instanceType** | [**ImageInstanceType**](ImageInstanceType.md) | 이미지와 호환 가능한 인스턴스 유형 &lt;br/&gt; - &#x60;vm&#x60;: Virtual Machine 유형  &lt;br/&gt; - &#x60;bm&#x60;: Bare Metal Server 유형 | 
 **imageType** | [**ImageVisibilityType**](ImageVisibilityType.md) | 이미지 제공 유형&lt;br/&gt; - &#x60;basic&#x60;: 카카오클라우드에서 기본적으로 제공하는 [기본 이미지](https://docs.kakaocloud.com/service/bcs/vm/vm-main#default-images-available)로 Linux 및 Windows 이미지를 제공 &lt;br/&gt; - &#x60;my&#x60;: 프로젝트 사용자가 생성한 커스텀 이미지 | 
 **size** | **int64** | 이미지 크기 (bytes 단위) | 
 **minDisk** | **int32** | 이미지를 사용할 때 필요한 최소 디스크 크기(GB) | 
 **diskFormat** | **string** | 이미지의 디스크 포맷 | 
 **status** | **string** | 이미지의 상태 | 
 **osType** | **string** | 운영체제 유형 | 
 **visibility** | **string** | 이미지의 가시성 | 
 **imageMemberStatus** | **string** | 공유된 이미지의 멤버 상태 | 
 **createdAt** | **string** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
 **updatedAt** | **string** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
 **sortKeys** | **string** | 정렬할 필드를 콤마(,)로 구분   | [default to &quot;created_at&quot;]
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;)   | [default to &quot;desc&quot;]
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | [default to 20]
 **offset** | **int32** | 조회 시작 위치 | [default to 0]

### Return type

[**ImageListModel**](ImageListModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveImageShare

> RemoveImageShare(ctx, imageId, memberId).XAuthToken(xAuthToken).Execute()

Remove image share



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
	imageId := "imageId_example" // string | 이미지의 고유 ID
	memberId := "memberId_example" // string | 이미지 공유를 해제할 멤버(프로젝트) ID 
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ImageAPI.RemoveImageShare(context.Background(), imageId, memberId).XAuthToken(xAuthToken).Execute()
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
**imageId** | **string** | 이미지의 고유 ID | 
**memberId** | **string** | 이미지 공유를 해제할 멤버(프로젝트) ID  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveImageShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

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

> BcsImageV1ApiUpdateImageModelResponseImageModel UpdateImage(ctx, imageId).XAuthToken(xAuthToken).BodyUpdateImage(bodyUpdateImage).Execute()

Update image



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
	imageId := "imageId_example" // string | 이미지의 고유 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bodyUpdateImage := *openapiclient.NewBodyUpdateImage(*openapiclient.NewEditImageModel()) // BodyUpdateImage | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.UpdateImage(context.Background(), imageId).XAuthToken(xAuthToken).BodyUpdateImage(bodyUpdateImage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.UpdateImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateImage`: BcsImageV1ApiUpdateImageModelResponseImageModel
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.UpdateImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**imageId** | **string** | 이미지의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyUpdateImage** | [**BodyUpdateImage**](BodyUpdateImage.md) |  | 

### Return type

[**BcsImageV1ApiUpdateImageModelResponseImageModel**](BcsImageV1ApiUpdateImageModelResponseImageModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

