# \InstanceNetworkInterfaceAPI

All URIs are relative to *https://bcs.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachNetworkInterface**](InstanceNetworkInterfaceAPI.md#AttachNetworkInterface) | **Post** /api/v1/instances/{instance_id}/network-interfaces/{network_interface_id} | Attach network interface
[**DetachNetworkInterface**](InstanceNetworkInterfaceAPI.md#DetachNetworkInterface) | **Delete** /api/v1/instances/{instance_id}/network-interfaces/{network_interface_id} | Detach network interface
[**ListInstanceNetworkInterfaces**](InstanceNetworkInterfaceAPI.md#ListInstanceNetworkInterfaces) | **Get** /api/v1/instances/{instance_id}/network-interfaces | List instance network interfaces
[**UpdateNetworkInterfaceAllowedAddresses**](InstanceNetworkInterfaceAPI.md#UpdateNetworkInterfaceAllowedAddresses) | **Put** /api/v1/instances/{instance_id}/network-interfaces/{network_interface_id}/allowed-address-pairs | Update network interface allowed addresses



## AttachNetworkInterface

> BcsInstanceV1ApiAttachNetworkInterfaceModelResponseInstanceNetworkInterfaceModel AttachNetworkInterface(ctx, instanceId, networkInterfaceId).XAuthToken(xAuthToken).Execute()

Attach network interface



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
	networkInterfaceId := "networkInterfaceId_example" // string | 인스턴스에 연결할 네트워크 인터페이스 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceNetworkInterfaceAPI.AttachNetworkInterface(context.Background(), instanceId, networkInterfaceId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceNetworkInterfaceAPI.AttachNetworkInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachNetworkInterface`: BcsInstanceV1ApiAttachNetworkInterfaceModelResponseInstanceNetworkInterfaceModel
	fmt.Fprintf(os.Stdout, "Response from `InstanceNetworkInterfaceAPI.AttachNetworkInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID | 
**networkInterfaceId** | **string** | 인스턴스에 연결할 네트워크 인터페이스 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachNetworkInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**BcsInstanceV1ApiAttachNetworkInterfaceModelResponseInstanceNetworkInterfaceModel**](BcsInstanceV1ApiAttachNetworkInterfaceModelResponseInstanceNetworkInterfaceModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DetachNetworkInterface

> interface{} DetachNetworkInterface(ctx, instanceId, networkInterfaceId).XAuthToken(xAuthToken).Execute()

Detach network interface



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
	networkInterfaceId := "networkInterfaceId_example" // string | 인스턴스에서 분리할 네트워크 인터페이스 ID 
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceNetworkInterfaceAPI.DetachNetworkInterface(context.Background(), instanceId, networkInterfaceId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceNetworkInterfaceAPI.DetachNetworkInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DetachNetworkInterface`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `InstanceNetworkInterfaceAPI.DetachNetworkInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID | 
**networkInterfaceId** | **string** | 인스턴스에서 분리할 네트워크 인터페이스 ID  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachNetworkInterfaceRequest struct via the builder pattern


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


## ListInstanceNetworkInterfaces

> InstanceNetworkInterfaceListModel ListInstanceNetworkInterfaces(ctx, instanceId).XAuthToken(xAuthToken).Execute()

List instance network interfaces



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
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceNetworkInterfaceAPI.ListInstanceNetworkInterfaces(context.Background(), instanceId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceNetworkInterfaceAPI.ListInstanceNetworkInterfaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInstanceNetworkInterfaces`: InstanceNetworkInterfaceListModel
	fmt.Fprintf(os.Stdout, "Response from `InstanceNetworkInterfaceAPI.ListInstanceNetworkInterfaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListInstanceNetworkInterfacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**InstanceNetworkInterfaceListModel**](InstanceNetworkInterfaceListModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkInterfaceAllowedAddresses

> BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelResponseInstanceNetworkInterfaceModel UpdateNetworkInterfaceAllowedAddresses(ctx, instanceId, networkInterfaceId).XAuthToken(xAuthToken).BodyUpdateNetworkInterfaceAllowedAddresses(bodyUpdateNetworkInterfaceAllowedAddresses).Execute()

Update network interface allowed addresses



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
	networkInterfaceId := "networkInterfaceId_example" // string | 인스턴스에 연결된 네트워크 인터페이스 ID 
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bodyUpdateNetworkInterfaceAllowedAddresses := *openapiclient.NewBodyUpdateNetworkInterfaceAllowedAddresses([]openapiclient.AllowedAddressPairModelInput{*openapiclient.NewAllowedAddressPairModelInput("IpAddress_example")}) // BodyUpdateNetworkInterfaceAllowedAddresses | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceNetworkInterfaceAPI.UpdateNetworkInterfaceAllowedAddresses(context.Background(), instanceId, networkInterfaceId).XAuthToken(xAuthToken).BodyUpdateNetworkInterfaceAllowedAddresses(bodyUpdateNetworkInterfaceAllowedAddresses).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceNetworkInterfaceAPI.UpdateNetworkInterfaceAllowedAddresses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkInterfaceAllowedAddresses`: BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelResponseInstanceNetworkInterfaceModel
	fmt.Fprintf(os.Stdout, "Response from `InstanceNetworkInterfaceAPI.UpdateNetworkInterfaceAllowedAddresses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID | 
**networkInterfaceId** | **string** | 인스턴스에 연결된 네트워크 인터페이스 ID  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkInterfaceAllowedAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyUpdateNetworkInterfaceAllowedAddresses** | [**BodyUpdateNetworkInterfaceAllowedAddresses**](BodyUpdateNetworkInterfaceAllowedAddresses.md) |  | 

### Return type

[**BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelResponseInstanceNetworkInterfaceModel**](BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelResponseInstanceNetworkInterfaceModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

