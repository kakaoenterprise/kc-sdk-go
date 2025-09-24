# \InstanceAttachedVolumeAPI

All URIs are relative to *https://bcs.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachVolume**](InstanceAttachedVolumeAPI.md#AttachVolume) | **Post** /api/v1/instances/{instance_id}/volumes/{volume_id} | Attach volume
[**DetachVolume**](InstanceAttachedVolumeAPI.md#DetachVolume) | **Delete** /api/v1/instances/{instance_id}/volumes/{volume_id} | Detach volume
[**UpdateInstanceVolume**](InstanceAttachedVolumeAPI.md#UpdateInstanceVolume) | **Put** /api/v1/instances/{instance_id}/volumes/{volume_id} | Update instance volume



## AttachVolume

> InstanceAttachedVolumeModelResponse AttachVolume(ctx, instanceId, volumeId).XAuthToken(xAuthToken).BodyAttachVolume(bodyAttachVolume).Execute()

Attach volume



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
	instanceId := "instanceId_example" // string | 인스턴스의 고유 ID
	volumeId := "volumeId_example" // string | 볼륨의 고유 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bodyAttachVolume := *openapiclient.NewBodyAttachVolume(*openapiclient.NewCreateVolumeModel()) // BodyAttachVolume | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAttachedVolumeAPI.AttachVolume(context.Background(), instanceId, volumeId).XAuthToken(xAuthToken).BodyAttachVolume(bodyAttachVolume).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAttachedVolumeAPI.AttachVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachVolume`: InstanceAttachedVolumeModelResponse
	fmt.Fprintf(os.Stdout, "Response from `InstanceAttachedVolumeAPI.AttachVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID | 
**volumeId** | **string** | 볼륨의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyAttachVolume** | [**BodyAttachVolume**](BodyAttachVolume.md) |  | 

### Return type

[**InstanceAttachedVolumeModelResponse**](InstanceAttachedVolumeModelResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DetachVolume

> interface{} DetachVolume(ctx, instanceId, volumeId).XAuthToken(xAuthToken).Execute()

Detach volume



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
	instanceId := "instanceId_example" // string | 인스턴스의 고유 ID
	volumeId := "volumeId_example" // string | 볼륨의 고유 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAttachedVolumeAPI.DetachVolume(context.Background(), instanceId, volumeId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAttachedVolumeAPI.DetachVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DetachVolume`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `InstanceAttachedVolumeAPI.DetachVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID | 
**volumeId** | **string** | 볼륨의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachVolumeRequest struct via the builder pattern


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


## UpdateInstanceVolume

> interface{} UpdateInstanceVolume(ctx, instanceId, volumeId).XAuthToken(xAuthToken).BodyUpdateInstanceVolume(bodyUpdateInstanceVolume).Execute()

Update instance volume



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
	instanceId := "instanceId_example" // string | 인스턴스의 고유 ID
	volumeId := "volumeId_example" // string | 볼륨의 고유 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bodyUpdateInstanceVolume := *openapiclient.NewBodyUpdateInstanceVolume(*openapiclient.NewEditVolumeModel()) // BodyUpdateInstanceVolume | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAttachedVolumeAPI.UpdateInstanceVolume(context.Background(), instanceId, volumeId).XAuthToken(xAuthToken).BodyUpdateInstanceVolume(bodyUpdateInstanceVolume).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAttachedVolumeAPI.UpdateInstanceVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInstanceVolume`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `InstanceAttachedVolumeAPI.UpdateInstanceVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID | 
**volumeId** | **string** | 볼륨의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInstanceVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyUpdateInstanceVolume** | [**BodyUpdateInstanceVolume**](BodyUpdateInstanceVolume.md) |  | 

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

