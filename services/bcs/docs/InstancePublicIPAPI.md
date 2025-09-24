# \InstancePublicIPAPI

All URIs are relative to *https://bcs.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssociateNewPublicIp**](InstancePublicIPAPI.md#AssociateNewPublicIp) | **Post** /api/v1/instances/{instance_id}/network-interfaces/{network_interface_id}/public-ips | Associate new public IP
[**AssociatePublicIp**](InstancePublicIPAPI.md#AssociatePublicIp) | **Put** /api/v1/instances/{instance_id}/network-interfaces/{network_interface_id}/public-ips/{public_ip_id} | Associate public IP
[**RemovePublicIp**](InstancePublicIPAPI.md#RemovePublicIp) | **Delete** /api/v1/instances/{instance_id}/network-interfaces/{network_interface_id}/public-ips | Remove public IP



## AssociateNewPublicIp

> BcsInstanceV1ApiAssociateNewPublicIpModelResponsePublicIpModel AssociateNewPublicIp(ctx, instanceId, networkInterfaceId).XAuthToken(xAuthToken).Execute()

Associate new public IP



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
	networkInterfaceId := "networkInterfaceId_example" // string | 퍼블릭 IP를 연결할 대상 네트워크 인터페이스 ID 
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancePublicIPAPI.AssociateNewPublicIp(context.Background(), instanceId, networkInterfaceId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancePublicIPAPI.AssociateNewPublicIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssociateNewPublicIp`: BcsInstanceV1ApiAssociateNewPublicIpModelResponsePublicIpModel
	fmt.Fprintf(os.Stdout, "Response from `InstancePublicIPAPI.AssociateNewPublicIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID | 
**networkInterfaceId** | **string** | 퍼블릭 IP를 연결할 대상 네트워크 인터페이스 ID  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateNewPublicIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**BcsInstanceV1ApiAssociateNewPublicIpModelResponsePublicIpModel**](BcsInstanceV1ApiAssociateNewPublicIpModelResponsePublicIpModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssociatePublicIp

> BcsInstanceV1ApiAssociatePublicIpModelResponsePublicIpModel AssociatePublicIp(ctx, instanceId, networkInterfaceId, publicIpId).XAuthToken(xAuthToken).Execute()

Associate public IP



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
	networkInterfaceId := "networkInterfaceId_example" // string | 퍼블릭 IP를 연결할 대상 네트워크 인터페이스 ID 
	publicIpId := "publicIpId_example" // string | 인스턴스에 연결할 퍼블릭 IP의 ID 
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancePublicIPAPI.AssociatePublicIp(context.Background(), instanceId, networkInterfaceId, publicIpId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancePublicIPAPI.AssociatePublicIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssociatePublicIp`: BcsInstanceV1ApiAssociatePublicIpModelResponsePublicIpModel
	fmt.Fprintf(os.Stdout, "Response from `InstancePublicIPAPI.AssociatePublicIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID | 
**networkInterfaceId** | **string** | 퍼블릭 IP를 연결할 대상 네트워크 인터페이스 ID  | 
**publicIpId** | **string** | 인스턴스에 연결할 퍼블릭 IP의 ID  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssociatePublicIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**BcsInstanceV1ApiAssociatePublicIpModelResponsePublicIpModel**](BcsInstanceV1ApiAssociatePublicIpModelResponsePublicIpModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemovePublicIp

> BcsInstanceV1ApiRemovePublicIpModelResponsePublicIpModel RemovePublicIp(ctx, instanceId, networkInterfaceId).XAuthToken(xAuthToken).IsDelete(isDelete).Execute()

Remove public IP



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
	networkInterfaceId := "networkInterfaceId_example" // string | 퍼블릭 IP를 분리할 네트워크 인터페이스 ID 
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	isDelete := true // bool | 퍼블릭 IP 연결 해제 후 즉시 삭제할지 여부 (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstancePublicIPAPI.RemovePublicIp(context.Background(), instanceId, networkInterfaceId).XAuthToken(xAuthToken).IsDelete(isDelete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstancePublicIPAPI.RemovePublicIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemovePublicIp`: BcsInstanceV1ApiRemovePublicIpModelResponsePublicIpModel
	fmt.Fprintf(os.Stdout, "Response from `InstancePublicIPAPI.RemovePublicIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID | 
**networkInterfaceId** | **string** | 퍼블릭 IP를 분리할 네트워크 인터페이스 ID  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemovePublicIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **isDelete** | **bool** | 퍼블릭 IP 연결 해제 후 즉시 삭제할지 여부 | [default to false]

### Return type

[**BcsInstanceV1ApiRemovePublicIpModelResponsePublicIpModel**](BcsInstanceV1ApiRemovePublicIpModelResponsePublicIpModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

