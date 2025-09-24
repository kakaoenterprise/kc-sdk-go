# \InstanceSecurityGroupAPI

All URIs are relative to *https://bcs.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachSecurityGroup**](InstanceSecurityGroupAPI.md#AttachSecurityGroup) | **Post** /api/v1/instances/{instance_id}/network-interfaces/{network_interface_id}/security-groups/{security_group_id} | Attach security group
[**DetachSecurityGroup**](InstanceSecurityGroupAPI.md#DetachSecurityGroup) | **Delete** /api/v1/instances/{instance_id}/network-interfaces/{network_interface_id}/security-groups/{security_group_id} | Detach security group



## AttachSecurityGroup

> ResponseSecurityGroup AttachSecurityGroup(ctx, instanceId, networkInterfaceId, securityGroupId).XAuthToken(xAuthToken).Execute()

Attach security group



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
	networkInterfaceId := "networkInterfaceId_example" // string | 보안 그룹을 연결할 네트워크 인터페이스 ID
	securityGroupId := "securityGroupId_example" // string | 보안 그룹의 고유 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceSecurityGroupAPI.AttachSecurityGroup(context.Background(), instanceId, networkInterfaceId, securityGroupId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceSecurityGroupAPI.AttachSecurityGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachSecurityGroup`: ResponseSecurityGroup
	fmt.Fprintf(os.Stdout, "Response from `InstanceSecurityGroupAPI.AttachSecurityGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID | 
**networkInterfaceId** | **string** | 보안 그룹을 연결할 네트워크 인터페이스 ID | 
**securityGroupId** | **string** | 보안 그룹의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachSecurityGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**ResponseSecurityGroup**](ResponseSecurityGroup.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DetachSecurityGroup

> ResponseSecurityGroup DetachSecurityGroup(ctx, instanceId, networkInterfaceId, securityGroupId).XAuthToken(xAuthToken).Execute()

Detach security group



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
	networkInterfaceId := "networkInterfaceId_example" // string | 보안 그룹 연결을 해제할 네트워크 인터페이스 ID
	securityGroupId := "securityGroupId_example" // string | 보안 그룹의 고유 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceSecurityGroupAPI.DetachSecurityGroup(context.Background(), instanceId, networkInterfaceId, securityGroupId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceSecurityGroupAPI.DetachSecurityGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DetachSecurityGroup`: ResponseSecurityGroup
	fmt.Fprintf(os.Stdout, "Response from `InstanceSecurityGroupAPI.DetachSecurityGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID | 
**networkInterfaceId** | **string** | 보안 그룹 연결을 해제할 네트워크 인터페이스 ID | 
**securityGroupId** | **string** | 보안 그룹의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachSecurityGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**ResponseSecurityGroup**](ResponseSecurityGroup.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

