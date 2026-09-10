# \InstanceAPI

All URIs are relative to *https://bcs.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssociateNewPublicIp**](InstanceAPI.md#AssociateNewPublicIp) | **Post** /api/v1/instances/{instance_id}/network-interfaces/{network_interface_id}/public-ips | Associate new public IP
[**AssociatePublicIp**](InstanceAPI.md#AssociatePublicIp) | **Put** /api/v1/instances/{instance_id}/network-interfaces/{network_interface_id}/public-ips/{public_ip_id} | Associate public IP
[**AttachNetworkInterface**](InstanceAPI.md#AttachNetworkInterface) | **Post** /api/v1/instances/{instance_id}/network-interfaces/{network_interface_id} | Attach network interface
[**AttachSecurityGroup**](InstanceAPI.md#AttachSecurityGroup) | **Post** /api/v1/instances/{instance_id}/network-interfaces/{network_interface_id}/security-groups/{security_group_id} | Attach security group
[**AttachVolume**](InstanceAPI.md#AttachVolume) | **Post** /api/v1/instances/{instance_id}/volumes/{volume_id} | Attach volume
[**CreateInstance**](InstanceAPI.md#CreateInstance) | **Post** /api/v1/instances | Create instance
[**DeleteInstance**](InstanceAPI.md#DeleteInstance) | **Delete** /api/v1/instances/{instance_id} | Delete instance
[**DetachNetworkInterface**](InstanceAPI.md#DetachNetworkInterface) | **Delete** /api/v1/instances/{instance_id}/network-interfaces/{network_interface_id} | Detach network interface
[**DetachSecurityGroup**](InstanceAPI.md#DetachSecurityGroup) | **Delete** /api/v1/instances/{instance_id}/network-interfaces/{network_interface_id}/security-groups/{security_group_id} | Detach security group
[**DetachVolume**](InstanceAPI.md#DetachVolume) | **Delete** /api/v1/instances/{instance_id}/volumes/{volume_id} | Detach volume
[**GetInstance**](InstanceAPI.md#GetInstance) | **Get** /api/v1/instances/{instance_id} | Get instance
[**GetInstanceConsoleLog**](InstanceAPI.md#GetInstanceConsoleLog) | **Get** /api/v1/instances/{instance_id}/console-logs | Get instance console log
[**GetInstancePassword**](InstanceAPI.md#GetInstancePassword) | **Get** /api/v1/instances/{instance_id}/password | Get instance password
[**GetInstanceType**](InstanceAPI.md#GetInstanceType) | **Get** /api/v1/flavors/{flavor_id} | Get instance type
[**HardRebootInstance**](InstanceAPI.md#HardRebootInstance) | **Post** /api/v1/instances/{instance_id}/hard-reboot | Hard reboot instance
[**ListInstanceActionLogs**](InstanceAPI.md#ListInstanceActionLogs) | **Get** /api/v1/instances/{instance_id}/action-logs | List instance action logs 
[**ListInstanceNetworkInterfaces**](InstanceAPI.md#ListInstanceNetworkInterfaces) | **Get** /api/v1/instances/{instance_id}/network-interfaces | List instance network interfaces
[**ListInstanceTypes**](InstanceAPI.md#ListInstanceTypes) | **Get** /api/v1/flavors | List instance types (flavors) 
[**ListInstances**](InstanceAPI.md#ListInstances) | **Get** /api/v1/instances | List instances
[**OperateInstance**](InstanceAPI.md#OperateInstance) | **Post** /api/v1/instances/{instance_id}/actions | Operate instances
[**RebuildInstance**](InstanceAPI.md#RebuildInstance) | **Post** /api/v1/instances/{instance_id}/rebuild | Rebuild instance
[**RemovePublicIp**](InstanceAPI.md#RemovePublicIp) | **Delete** /api/v1/instances/{instance_id}/network-interfaces/{network_interface_id}/public-ips | Remove public IP
[**ResizeInstance**](InstanceAPI.md#ResizeInstance) | **Post** /api/v1/instances/{instance_id}/resize | Resize instance
[**ShelveInstance**](InstanceAPI.md#ShelveInstance) | **Post** /api/v1/instances/{instance_id}/shelve | Shelve instance
[**SoftRebootInstance**](InstanceAPI.md#SoftRebootInstance) | **Post** /api/v1/instances/{instance_id}/soft-reboot | Soft reboot instance 
[**StartInstance**](InstanceAPI.md#StartInstance) | **Post** /api/v1/instances/{instance_id}/start | Start instance
[**StopInstance**](InstanceAPI.md#StopInstance) | **Post** /api/v1/instances/{instance_id}/stop | Stop instance
[**UnshelveInstance**](InstanceAPI.md#UnshelveInstance) | **Post** /api/v1/instances/{instance_id}/unshelve | Unshelve instance
[**UpdateInstance**](InstanceAPI.md#UpdateInstance) | **Put** /api/v1/instances/{instance_id} | Update instance
[**UpdateInstanceVolume**](InstanceAPI.md#UpdateInstanceVolume) | **Put** /api/v1/instances/{instance_id}/volumes/{volume_id} | Update instance volume
[**UpdateNetworkInterfaceAllowedAddresses**](InstanceAPI.md#UpdateNetworkInterfaceAllowedAddresses) | **Put** /api/v1/instances/{instance_id}/network-interfaces/{network_interface_id}/allowed-address-pairs | Update network interface allowed addresses



## AssociateNewPublicIp

> AssociateNewPublicIpResponse AssociateNewPublicIp(ctx, instanceId, networkInterfaceId).XAuthToken(xAuthToken).Execute()

Associate new public IP



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/bcs"
)

func main() {
	instanceId := "instanceId_example" // string | 인스턴스의 고유 ID <br/> - [List instances](https://docs.kakaocloud.com/openapi/bcs/list-instances)에서 확인
	networkInterfaceId := "networkInterfaceId_example" // string | 퍼블릭 IP를 연결할 대상 네트워크 인터페이스 ID <br/> - [List instance network interfaces](https://docs.kakaocloud.com/openapi/bcs/list-instance-network-interfaces)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.AssociateNewPublicIp(context.Background(), instanceId, networkInterfaceId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.AssociateNewPublicIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssociateNewPublicIp`: AssociateNewPublicIpResponse
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.AssociateNewPublicIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID &lt;br/&gt; - [List instances](https://docs.kakaocloud.com/openapi/bcs/list-instances)에서 확인 | 
**networkInterfaceId** | **string** | 퍼블릭 IP를 연결할 대상 네트워크 인터페이스 ID &lt;br/&gt; - [List instance network interfaces](https://docs.kakaocloud.com/openapi/bcs/list-instance-network-interfaces)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateNewPublicIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**AssociateNewPublicIpResponse**](AssociateNewPublicIpResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssociatePublicIp

> AssociatePublicIpResponse AssociatePublicIp(ctx, instanceId, networkInterfaceId, publicIpId).XAuthToken(xAuthToken).Execute()

Associate public IP



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/bcs"
)

func main() {
	instanceId := "instanceId_example" // string | 인스턴스의 고유 ID <br/> - [List instances](https://docs.kakaocloud.com/openapi/bcs/list-instances)에서 확인
	networkInterfaceId := "networkInterfaceId_example" // string | 퍼블릭 IP를 연결할 대상 네트워크 인터페이스 ID <br/> - [List instance network interfaces](https://docs.kakaocloud.com/openapi/bcs/list-instance-network-interfaces)에서 확인
	publicIpId := "publicIpId_example" // string | 인스턴스에 연결할 퍼블릭 IP의 ID <br/> - [List public IPs](https://docs.kakaocloud.com/openapi/networking/vpc/list-public-ips)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.AssociatePublicIp(context.Background(), instanceId, networkInterfaceId, publicIpId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.AssociatePublicIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssociatePublicIp`: AssociatePublicIpResponse
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.AssociatePublicIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID &lt;br/&gt; - [List instances](https://docs.kakaocloud.com/openapi/bcs/list-instances)에서 확인 | 
**networkInterfaceId** | **string** | 퍼블릭 IP를 연결할 대상 네트워크 인터페이스 ID &lt;br/&gt; - [List instance network interfaces](https://docs.kakaocloud.com/openapi/bcs/list-instance-network-interfaces)에서 확인 | 
**publicIpId** | **string** | 인스턴스에 연결할 퍼블릭 IP의 ID &lt;br/&gt; - [List public IPs](https://docs.kakaocloud.com/openapi/networking/vpc/list-public-ips)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssociatePublicIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**AssociatePublicIpResponse**](AssociatePublicIpResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttachNetworkInterface

> AttachNetworkInterfaceResponse AttachNetworkInterface(ctx, instanceId, networkInterfaceId).XAuthToken(xAuthToken).Execute()

Attach network interface



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/bcs"
)

func main() {
	instanceId := "instanceId_example" // string | 인스턴스의 고유 ID <br/> - [List instances](https://docs.kakaocloud.com/openapi/bcs/list-instances)에서 확인
	networkInterfaceId := "networkInterfaceId_example" // string | 인스턴스에 연결할 네트워크 인터페이스 ID <br/> - [List network interfaces](https://docs.kakaocloud.com/openapi/networking/vpc/list-network-interfaces)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.AttachNetworkInterface(context.Background(), instanceId, networkInterfaceId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.AttachNetworkInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachNetworkInterface`: AttachNetworkInterfaceResponse
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.AttachNetworkInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID &lt;br/&gt; - [List instances](https://docs.kakaocloud.com/openapi/bcs/list-instances)에서 확인 | 
**networkInterfaceId** | **string** | 인스턴스에 연결할 네트워크 인터페이스 ID &lt;br/&gt; - [List network interfaces](https://docs.kakaocloud.com/openapi/networking/vpc/list-network-interfaces)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachNetworkInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**AttachNetworkInterfaceResponse**](AttachNetworkInterfaceResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttachSecurityGroup

> AttachSecurityGroupResponse AttachSecurityGroup(ctx, instanceId, networkInterfaceId, securityGroupId).XAuthToken(xAuthToken).Execute()

Attach security group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/bcs"
)

func main() {
	instanceId := "instanceId_example" // string | 인스턴스의 고유 ID <br/> - [List instances](https://docs.kakaocloud.com/openapi/bcs/list-instances)에서 확인
	networkInterfaceId := "networkInterfaceId_example" // string | 보안 그룹을 연결할 네트워크 인터페이스 ID <br/> - [List instance network interfaces](https://docs.kakaocloud.com/openapi/bcs/list-instance-network-interfaces)에서 확인
	securityGroupId := "securityGroupId_example" // string | 보안 그룹의 고유 ID <br/> - [List security groups](https://docs.kakaocloud.com/openapi/networking/vpc/list-security-groups)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.AttachSecurityGroup(context.Background(), instanceId, networkInterfaceId, securityGroupId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.AttachSecurityGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachSecurityGroup`: AttachSecurityGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.AttachSecurityGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID &lt;br/&gt; - [List instances](https://docs.kakaocloud.com/openapi/bcs/list-instances)에서 확인 | 
**networkInterfaceId** | **string** | 보안 그룹을 연결할 네트워크 인터페이스 ID &lt;br/&gt; - [List instance network interfaces](https://docs.kakaocloud.com/openapi/bcs/list-instance-network-interfaces)에서 확인 | 
**securityGroupId** | **string** | 보안 그룹의 고유 ID &lt;br/&gt; - [List security groups](https://docs.kakaocloud.com/openapi/networking/vpc/list-security-groups)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachSecurityGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**AttachSecurityGroupResponse**](AttachSecurityGroupResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AttachVolume

> AttachVolumeResponse AttachVolume(ctx, instanceId, volumeId).XAuthToken(xAuthToken).AttachVolumeRequest(attachVolumeRequest).Execute()

Attach volume



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/bcs"
)

func main() {
	instanceId := "instanceId_example" // string | 인스턴스의 고유 ID <br/> - [List instances](https://docs.kakaocloud.com/openapi/bcs/list-instances)에서 확인
	volumeId := "volumeId_example" // string | 볼륨의 고유 ID <br/> - [List volumes](https://docs.kakaocloud.com/openapi/bcs/list-volumes)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	attachVolumeRequest := *openapiclient.NewAttachVolumeRequest(*openapiclient.NewAttachVolume()) // AttachVolumeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.AttachVolume(context.Background(), instanceId, volumeId).XAuthToken(xAuthToken).AttachVolumeRequest(attachVolumeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.AttachVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AttachVolume`: AttachVolumeResponse
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.AttachVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID &lt;br/&gt; - [List instances](https://docs.kakaocloud.com/openapi/bcs/list-instances)에서 확인 | 
**volumeId** | **string** | 볼륨의 고유 ID &lt;br/&gt; - [List volumes](https://docs.kakaocloud.com/openapi/bcs/list-volumes)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **attachVolumeRequest** | [**AttachVolumeRequest**](AttachVolumeRequest.md) |  | 

### Return type

[**AttachVolumeResponse**](AttachVolumeResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInstance

> CreateInstanceResponse CreateInstance(ctx).XAuthToken(xAuthToken).CreateInstanceRequest(createInstanceRequest).Execute()

Create instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/bcs"
)

func main() {
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	createInstanceRequest := *openapiclient.NewCreateInstanceRequest(*openapiclient.NewCreateInstance("Name_example", "ImageId_example", "FlavorId_example", []openapiclient.SubnetRequest{*openapiclient.NewSubnetRequest()})) // CreateInstanceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.CreateInstance(context.Background()).XAuthToken(xAuthToken).CreateInstanceRequest(createInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.CreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInstance`: CreateInstanceResponse
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.CreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **createInstanceRequest** | [**CreateInstanceRequest**](CreateInstanceRequest.md) |  | 

### Return type

[**CreateInstanceResponse**](CreateInstanceResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInstance

> DeleteInstance(ctx, instanceId).XAuthToken(xAuthToken).Execute()

Delete instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/bcs"
)

func main() {
	instanceId := "instanceId_example" // string | 인스턴스의 고유 ID <br/> - [List instances](https://docs.kakaocloud.com/openapi/bcs/list-instances)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstanceAPI.DeleteInstance(context.Background(), instanceId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.DeleteInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID &lt;br/&gt; - [List instances](https://docs.kakaocloud.com/openapi/bcs/list-instances)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInstanceRequest struct via the builder pattern


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
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/bcs"
)

func main() {
	instanceId := "instanceId_example" // string | 인스턴스의 고유 ID <br/> - [List instances](https://docs.kakaocloud.com/openapi/bcs/list-instances)에서 확인
	networkInterfaceId := "networkInterfaceId_example" // string | 인스턴스에서 분리할 네트워크 인터페이스 ID <br/> - [List instance network interfaces](https://docs.kakaocloud.com/openapi/bcs/list-instance-network-interfaces)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.DetachNetworkInterface(context.Background(), instanceId, networkInterfaceId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.DetachNetworkInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DetachNetworkInterface`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.DetachNetworkInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID &lt;br/&gt; - [List instances](https://docs.kakaocloud.com/openapi/bcs/list-instances)에서 확인 | 
**networkInterfaceId** | **string** | 인스턴스에서 분리할 네트워크 인터페이스 ID &lt;br/&gt; - [List instance network interfaces](https://docs.kakaocloud.com/openapi/bcs/list-instance-network-interfaces)에서 확인 | 

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


## DetachSecurityGroup

> DetachSecurityGroupResponse DetachSecurityGroup(ctx, instanceId, networkInterfaceId, securityGroupId).XAuthToken(xAuthToken).Execute()

Detach security group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/bcs"
)

func main() {
	instanceId := "instanceId_example" // string | 인스턴스의 고유 ID <br/> - [List instances](https://docs.kakaocloud.com/openapi/bcs/list-instances)에서 확인
	networkInterfaceId := "networkInterfaceId_example" // string | 보안 그룹 연결을 해제할 네트워크 인터페이스 ID <br/> - [List instance network interfaces](https://docs.kakaocloud.com/openapi/bcs/list-instance-network-interfaces)에서 확인
	securityGroupId := "securityGroupId_example" // string | 보안 그룹의 고유 ID <br/> - [List security groups](https://docs.kakaocloud.com/openapi/networking/vpc/list-security-groups)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.DetachSecurityGroup(context.Background(), instanceId, networkInterfaceId, securityGroupId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.DetachSecurityGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DetachSecurityGroup`: DetachSecurityGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.DetachSecurityGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID &lt;br/&gt; - [List instances](https://docs.kakaocloud.com/openapi/bcs/list-instances)에서 확인 | 
**networkInterfaceId** | **string** | 보안 그룹 연결을 해제할 네트워크 인터페이스 ID &lt;br/&gt; - [List instance network interfaces](https://docs.kakaocloud.com/openapi/bcs/list-instance-network-interfaces)에서 확인 | 
**securityGroupId** | **string** | 보안 그룹의 고유 ID &lt;br/&gt; - [List security groups](https://docs.kakaocloud.com/openapi/networking/vpc/list-security-groups)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachSecurityGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**DetachSecurityGroupResponse**](DetachSecurityGroupResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
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
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/bcs"
)

func main() {
	instanceId := "instanceId_example" // string | 인스턴스의 고유 ID <br/> - [List instances](https://docs.kakaocloud.com/openapi/bcs/list-instances)에서 확인
	volumeId := "volumeId_example" // string | 볼륨의 고유 ID <br/> - [List volumes](https://docs.kakaocloud.com/openapi/bcs/list-volumes)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.DetachVolume(context.Background(), instanceId, volumeId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.DetachVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DetachVolume`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.DetachVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID &lt;br/&gt; - [List instances](https://docs.kakaocloud.com/openapi/bcs/list-instances)에서 확인 | 
**volumeId** | **string** | 볼륨의 고유 ID &lt;br/&gt; - [List volumes](https://docs.kakaocloud.com/openapi/bcs/list-volumes)에서 확인 | 

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


## GetInstance

> GetInstanceResponse GetInstance(ctx, instanceId).XAuthToken(xAuthToken).Execute()

Get instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/bcs"
)

func main() {
	instanceId := "instanceId_example" // string | 인스턴스의 고유 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.GetInstance(context.Background(), instanceId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.GetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstance`: GetInstanceResponse
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.GetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**GetInstanceResponse**](GetInstanceResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstanceConsoleLog

> GetInstanceConsoleLogResponse GetInstanceConsoleLog(ctx, instanceId).XAuthToken(xAuthToken).Lines(lines).Execute()

Get instance console log



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/bcs"
)

func main() {
	instanceId := "instanceId_example" // string | 인스턴스의 고유 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	lines := int32(56) // int32 | 콘솔 로그의 마지막 줄부터 조회할 라인(line)의 수 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.GetInstanceConsoleLog(context.Background(), instanceId).XAuthToken(xAuthToken).Lines(lines).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.GetInstanceConsoleLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstanceConsoleLog`: GetInstanceConsoleLogResponse
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.GetInstanceConsoleLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceConsoleLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **lines** | **int32** | 콘솔 로그의 마지막 줄부터 조회할 라인(line)의 수 | 

### Return type

[**GetInstanceConsoleLogResponse**](GetInstanceConsoleLogResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstancePassword

> GetInstancePasswordResponse GetInstancePassword(ctx, instanceId).XAuthToken(xAuthToken).Execute()

Get instance password



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/bcs"
)

func main() {
	instanceId := "instanceId_example" // string | 인스턴스의 고유 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.GetInstancePassword(context.Background(), instanceId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.GetInstancePassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstancePassword`: GetInstancePasswordResponse
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.GetInstancePassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstancePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**GetInstancePasswordResponse**](GetInstancePasswordResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstanceType

> GetInstanceTypeResponse GetInstanceType(ctx, flavorId).XAuthToken(xAuthToken).Execute()

Get instance type



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/bcs"
)

func main() {
	flavorId := "flavorId_example" // string | 인스턴스 유형 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.GetInstanceType(context.Background(), flavorId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.GetInstanceType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstanceType`: GetInstanceTypeResponse
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.GetInstanceType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flavorId** | **string** | 인스턴스 유형 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**GetInstanceTypeResponse**](GetInstanceTypeResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HardRebootInstance

> interface{} HardRebootInstance(ctx, instanceId).XAuthToken(xAuthToken).Execute()

Hard reboot instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/bcs"
)

func main() {
	instanceId := "instanceId_example" // string | 인스턴스의 고유 ID <br/> - [List instances](https://docs.kakaocloud.com/openapi/bcs/list-instances)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.HardRebootInstance(context.Background(), instanceId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.HardRebootInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HardRebootInstance`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.HardRebootInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID &lt;br/&gt; - [List instances](https://docs.kakaocloud.com/openapi/bcs/list-instances)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiHardRebootInstanceRequest struct via the builder pattern


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


## ListInstanceActionLogs

> ListInstanceActionLogsResponse ListInstanceActionLogs(ctx, instanceId).XAuthToken(xAuthToken).Execute()

List instance action logs 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/bcs"
)

func main() {
	instanceId := "instanceId_example" // string | 인스턴스의 고유 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.ListInstanceActionLogs(context.Background(), instanceId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.ListInstanceActionLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInstanceActionLogs`: ListInstanceActionLogsResponse
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.ListInstanceActionLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListInstanceActionLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**ListInstanceActionLogsResponse**](ListInstanceActionLogsResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInstanceNetworkInterfaces

> ListInstanceNetworkInterfacesResponse ListInstanceNetworkInterfaces(ctx, instanceId).XAuthToken(xAuthToken).Execute()

List instance network interfaces



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/bcs"
)

func main() {
	instanceId := "instanceId_example" // string | 인스턴스의 고유 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.ListInstanceNetworkInterfaces(context.Background(), instanceId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.ListInstanceNetworkInterfaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInstanceNetworkInterfaces`: ListInstanceNetworkInterfacesResponse
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.ListInstanceNetworkInterfaces`: %v\n", resp)
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

[**ListInstanceNetworkInterfacesResponse**](ListInstanceNetworkInterfacesResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInstanceTypes

> ListInstanceTypesResponse ListInstanceTypes(ctx).XAuthToken(xAuthToken).Id(id).Name(name).IsBurstable(isBurstable).Vcpus(vcpus).Architecture(architecture).MemoryMb(memoryMb).InstanceType(instanceType).InstanceFamily(instanceFamily).InstanceSize(instanceSize).Manufacturer(manufacturer).MaximumNetworkInterfaces(maximumNetworkInterfaces).Processor(processor).CreatedAt(createdAt).UpdatedAt(updatedAt).Offset(offset).Limit(limit).SortKeys(sortKeys).SortDirs(sortDirs).Execute()

List instance types (flavors) 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/bcs"
)

func main() {
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	id := "id_example" // string | 조회할 인스턴스 유형(Flavor) ID (optional)
	name := "name_example" // string | 인스턴스 유형의 이름 (optional)
	isBurstable := true // bool | [버스터블 인스턴스](https://docs.kakaocloud.com/service/bcs/bcs-specifications/general-purpose/burstable-main) 여부 (optional)
	vcpus := int32(56) // int32 | 가상 CPU 개수 (optional)
	architecture := "architecture_example" // string | 인스턴스 유형의 아키텍처(CPU 구조) (optional)
	memoryMb := int32(56) // int32 | 메모리 크기 (MB 단위) (optional)
	instanceType := "instanceType_example" // string | 인스턴스 유형 <br/> - `vm`: Virtual Machine 유형  <br/> - `bm`: Bare Metal Server 유형 <br/> - `gpu`: GPU 유형 (optional)
	instanceFamily := "instanceFamily_example" // string | [인스턴스 패밀리](https://docs.kakaocloud.com/service/bcs/instance-overview#instance-family) <br/> - 예시:  `r2a`, `c2a`  등 (optional)
	instanceSize := "instanceSize_example" // string | 인스턴스 크기 (optional)
	manufacturer := "manufacturer_example" // string | 제조사 정보 (optional)
	maximumNetworkInterfaces := int32(56) // int32 | 최대 네트워크 인터페이스 개수 (optional)
	processor := "processor_example" // string | 프로세서 이름 (optional)
	createdAt := "createdAt_example" // string | 리소스가 생성된 시간 <br/> - ISO 8601 형식 <br/> - UTC 기준 (optional)
	updatedAt := "updatedAt_example" // string | 리소스가 마지막으로 수정된 시간 <br/> - ISO 8601 형식 <br/> - UTC 기준 (optional)
	offset := int32(56) // int32 | 조회 시작 위치 (optional)
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional)
	sortKeys := "sortKeys_example" // string |  정렬할 필드를 콤마(,)로 구분   (optional)
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`)  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.ListInstanceTypes(context.Background()).XAuthToken(xAuthToken).Id(id).Name(name).IsBurstable(isBurstable).Vcpus(vcpus).Architecture(architecture).MemoryMb(memoryMb).InstanceType(instanceType).InstanceFamily(instanceFamily).InstanceSize(instanceSize).Manufacturer(manufacturer).MaximumNetworkInterfaces(maximumNetworkInterfaces).Processor(processor).CreatedAt(createdAt).UpdatedAt(updatedAt).Offset(offset).Limit(limit).SortKeys(sortKeys).SortDirs(sortDirs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.ListInstanceTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInstanceTypes`: ListInstanceTypesResponse
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.ListInstanceTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInstanceTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **id** | **string** | 조회할 인스턴스 유형(Flavor) ID | 
 **name** | **string** | 인스턴스 유형의 이름 | 
 **isBurstable** | **bool** | [버스터블 인스턴스](https://docs.kakaocloud.com/service/bcs/bcs-specifications/general-purpose/burstable-main) 여부 | 
 **vcpus** | **int32** | 가상 CPU 개수 | 
 **architecture** | **string** | 인스턴스 유형의 아키텍처(CPU 구조) | 
 **memoryMb** | **int32** | 메모리 크기 (MB 단위) | 
 **instanceType** | **string** | 인스턴스 유형 &lt;br/&gt; - &#x60;vm&#x60;: Virtual Machine 유형  &lt;br/&gt; - &#x60;bm&#x60;: Bare Metal Server 유형 &lt;br/&gt; - &#x60;gpu&#x60;: GPU 유형 | 
 **instanceFamily** | **string** | [인스턴스 패밀리](https://docs.kakaocloud.com/service/bcs/instance-overview#instance-family) &lt;br/&gt; - 예시:  &#x60;r2a&#x60;, &#x60;c2a&#x60;  등 | 
 **instanceSize** | **string** | 인스턴스 크기 | 
 **manufacturer** | **string** | 제조사 정보 | 
 **maximumNetworkInterfaces** | **int32** | 최대 네트워크 인터페이스 개수 | 
 **processor** | **string** | 프로세서 이름 | 
 **createdAt** | **string** | 리소스가 생성된 시간 &lt;br/&gt; - ISO 8601 형식 &lt;br/&gt; - UTC 기준 | 
 **updatedAt** | **string** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO 8601 형식 &lt;br/&gt; - UTC 기준 | 
 **offset** | **int32** | 조회 시작 위치 | 
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | 
 **sortKeys** | **string** |  정렬할 필드를 콤마(,)로 구분   | 
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;)  | 

### Return type

[**ListInstanceTypesResponse**](ListInstanceTypesResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInstances

> ListInstancesResponse ListInstances(ctx).XAuthToken(xAuthToken).Id(id).Name(name).VmState(vmState).FlavorName(flavorName).ImageName(imageName).PrivateIp(privateIp).PublicIp(publicIp).AvailabilityZone(availabilityZone).InstanceType(instanceType).Status(status).UserId(userId).Hostname(hostname).OsType(osType).IsHadoop(isHadoop).IsK8se(isK8se).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Offset(offset).Limit(limit).Execute()

List instances



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/bcs"
)

func main() {
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	id := "id_example" // string | 조회할 인스턴스의 ID (optional)
	name := "name_example" // string | 인스턴스 이름 <br/> - 특정 이름을 포함하는 인스턴스를 검색 시 사용 (optional)
	vmState := "vmState_example" // string | 인스턴스의 상태 코드 <br/> - 전체 [인스턴스 상태 코드](https://docs.kakaocloud.com/service/bcs/vm/vm-main#instance-state-and-billing) 참고 (optional)
	flavorName := "flavorName_example" // string | 인스턴스 유형 이름  (optional)
	imageName := "imageName_example" // string | 인스턴스를 생성할 때 사용한 이미지 이름 (optional)
	privateIp := "privateIp_example" // string | 인스턴스에 할당된 프라이빗 IP 주소 (IPv4 형식) (optional)
	publicIp := "publicIp_example" // string | 인스턴스에 연결된 퍼블릭 IP 주소 (optional)
	availabilityZone := openapiclient.AvailabilityZone("kr-central-2-a") // AvailabilityZone | 인스턴스가 위치한 가용 영역 (optional)
	instanceType := "instanceType_example" // string | 인스턴스 유형 <br/> - `vm`: Virtual Machine 유형  <br/> - `bm`: Bare Metal Server 유형 <br/> - `gpu`: GPU 유형 (optional)
	status := "status_example" // string | 인스턴스의 상세 상태 정보<br/> - 내부적으로 정의된 상태 값 (optional)
	userId := "userId_example" // string | 해당 인스턴스를 생성한 사용자 ID (optional)
	hostname := "hostname_example" // string | 인스턴스의 호스트 이름(내부 DNS 이름 등) (optional)
	osType := "osType_example" // string | 운영체제 유형 (optional)
	isHadoop := true // bool | Hadoop 환경용으로 생성된 인스턴스인지 여부 (optional)
	isK8se := true // bool | Kubernetes Engine 환경용으로 생성된 인스턴스인지 여부 (optional)
	createdAt := "createdAt_example" // string | 리소스가 생성된 시간 <br/> - ISO 8601 형식 <br/> - UTC 기준 (optional)
	updatedAt := "updatedAt_example" // string | 리소스가 마지막으로 수정된 시간 <br/> - ISO 8601 형식 <br/> - UTC 기준 (optional)
	sortKeys := "sortKeys_example" // string | 정렬할 필드를 콤마(,)로 구분   (optional)
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`)  (optional)
	offset := int32(56) // int32 | 조회 시작 위치 (optional)
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.ListInstances(context.Background()).XAuthToken(xAuthToken).Id(id).Name(name).VmState(vmState).FlavorName(flavorName).ImageName(imageName).PrivateIp(privateIp).PublicIp(publicIp).AvailabilityZone(availabilityZone).InstanceType(instanceType).Status(status).UserId(userId).Hostname(hostname).OsType(osType).IsHadoop(isHadoop).IsK8se(isK8se).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.ListInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInstances`: ListInstancesResponse
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.ListInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **id** | **string** | 조회할 인스턴스의 ID | 
 **name** | **string** | 인스턴스 이름 &lt;br/&gt; - 특정 이름을 포함하는 인스턴스를 검색 시 사용 | 
 **vmState** | **string** | 인스턴스의 상태 코드 &lt;br/&gt; - 전체 [인스턴스 상태 코드](https://docs.kakaocloud.com/service/bcs/vm/vm-main#instance-state-and-billing) 참고 | 
 **flavorName** | **string** | 인스턴스 유형 이름  | 
 **imageName** | **string** | 인스턴스를 생성할 때 사용한 이미지 이름 | 
 **privateIp** | **string** | 인스턴스에 할당된 프라이빗 IP 주소 (IPv4 형식) | 
 **publicIp** | **string** | 인스턴스에 연결된 퍼블릭 IP 주소 | 
 **availabilityZone** | [**AvailabilityZone**](AvailabilityZone.md) | 인스턴스가 위치한 가용 영역 | 
 **instanceType** | **string** | 인스턴스 유형 &lt;br/&gt; - &#x60;vm&#x60;: Virtual Machine 유형  &lt;br/&gt; - &#x60;bm&#x60;: Bare Metal Server 유형 &lt;br/&gt; - &#x60;gpu&#x60;: GPU 유형 | 
 **status** | **string** | 인스턴스의 상세 상태 정보&lt;br/&gt; - 내부적으로 정의된 상태 값 | 
 **userId** | **string** | 해당 인스턴스를 생성한 사용자 ID | 
 **hostname** | **string** | 인스턴스의 호스트 이름(내부 DNS 이름 등) | 
 **osType** | **string** | 운영체제 유형 | 
 **isHadoop** | **bool** | Hadoop 환경용으로 생성된 인스턴스인지 여부 | 
 **isK8se** | **bool** | Kubernetes Engine 환경용으로 생성된 인스턴스인지 여부 | 
 **createdAt** | **string** | 리소스가 생성된 시간 &lt;br/&gt; - ISO 8601 형식 &lt;br/&gt; - UTC 기준 | 
 **updatedAt** | **string** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO 8601 형식 &lt;br/&gt; - UTC 기준 | 
 **sortKeys** | **string** | 정렬할 필드를 콤마(,)로 구분   | 
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;)  | 
 **offset** | **int32** | 조회 시작 위치 | 
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | 

### Return type

[**ListInstancesResponse**](ListInstancesResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## OperateInstance

> interface{} OperateInstance(ctx, instanceId).Action(action).XAuthToken(xAuthToken).Execute()

Operate instances



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/bcs"
)

func main() {
	instanceId := "instanceId_example" // string | 액션을 수행할 대상 인스턴스 ID <br/> - [List instances](https://docs.kakaocloud.com/openapi/bcs/list-instances)에서 확인
	action := openapiclient.InstanceActionType("start") // InstanceActionType | 수행할 인스턴스 액션
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.OperateInstance(context.Background(), instanceId).Action(action).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.OperateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OperateInstance`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.OperateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 액션을 수행할 대상 인스턴스 ID &lt;br/&gt; - [List instances](https://docs.kakaocloud.com/openapi/bcs/list-instances)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiOperateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **action** | [**InstanceActionType**](InstanceActionType.md) | 수행할 인스턴스 액션 | 
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


## RebuildInstance

> RebuildInstanceResponse RebuildInstance(ctx, instanceId).XAuthToken(xAuthToken).RebuildInstanceRequest(rebuildInstanceRequest).Execute()

Rebuild instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/bcs"
)

func main() {
	instanceId := "instanceId_example" // string | 인스턴스의 고유 ID <br/> - [List instances](https://docs.kakaocloud.com/openapi/bcs/list-instances)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	rebuildInstanceRequest := *openapiclient.NewRebuildInstanceRequest(*openapiclient.NewRebuildInstance("ImageId_example")) // RebuildInstanceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.RebuildInstance(context.Background(), instanceId).XAuthToken(xAuthToken).RebuildInstanceRequest(rebuildInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.RebuildInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RebuildInstance`: RebuildInstanceResponse
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.RebuildInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID &lt;br/&gt; - [List instances](https://docs.kakaocloud.com/openapi/bcs/list-instances)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebuildInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **rebuildInstanceRequest** | [**RebuildInstanceRequest**](RebuildInstanceRequest.md) |  | 

### Return type

[**RebuildInstanceResponse**](RebuildInstanceResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemovePublicIp

> RemovePublicIpResponse RemovePublicIp(ctx, instanceId, networkInterfaceId).XAuthToken(xAuthToken).IsDelete(isDelete).Execute()

Remove public IP



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/bcs"
)

func main() {
	instanceId := "instanceId_example" // string | 인스턴스의 고유 ID <br/> - [List instances](https://docs.kakaocloud.com/openapi/bcs/list-instances)에서 확인
	networkInterfaceId := "networkInterfaceId_example" // string | 퍼블릭 IP를 분리할 네트워크 인터페이스 ID <br/> - [List instance network interfaces](https://docs.kakaocloud.com/openapi/bcs/list-instance-network-interfaces)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	isDelete := true // bool | 퍼블릭 IP 연결 해제 후 즉시 삭제할지 여부 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.RemovePublicIp(context.Background(), instanceId, networkInterfaceId).XAuthToken(xAuthToken).IsDelete(isDelete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.RemovePublicIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemovePublicIp`: RemovePublicIpResponse
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.RemovePublicIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID &lt;br/&gt; - [List instances](https://docs.kakaocloud.com/openapi/bcs/list-instances)에서 확인 | 
**networkInterfaceId** | **string** | 퍼블릭 IP를 분리할 네트워크 인터페이스 ID &lt;br/&gt; - [List instance network interfaces](https://docs.kakaocloud.com/openapi/bcs/list-instance-network-interfaces)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemovePublicIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **isDelete** | **bool** | 퍼블릭 IP 연결 해제 후 즉시 삭제할지 여부 | 

### Return type

[**RemovePublicIpResponse**](RemovePublicIpResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResizeInstance

> interface{} ResizeInstance(ctx, instanceId).XAuthToken(xAuthToken).ResizeInstanceRequest(resizeInstanceRequest).Execute()

Resize instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/bcs"
)

func main() {
	instanceId := "instanceId_example" // string | 인스턴스의 고유 ID <br/> - [List instances](https://docs.kakaocloud.com/openapi/bcs/list-instances)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	resizeInstanceRequest := *openapiclient.NewResizeInstanceRequest(*openapiclient.NewResizeInstance("Id_example")) // ResizeInstanceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.ResizeInstance(context.Background(), instanceId).XAuthToken(xAuthToken).ResizeInstanceRequest(resizeInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.ResizeInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResizeInstance`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.ResizeInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID &lt;br/&gt; - [List instances](https://docs.kakaocloud.com/openapi/bcs/list-instances)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiResizeInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **resizeInstanceRequest** | [**ResizeInstanceRequest**](ResizeInstanceRequest.md) |  | 

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


## ShelveInstance

> interface{} ShelveInstance(ctx, instanceId).XAuthToken(xAuthToken).Execute()

Shelve instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/bcs"
)

func main() {
	instanceId := "instanceId_example" // string | 인스턴스의 고유 ID <br/> - [List instances](https://docs.kakaocloud.com/openapi/bcs/list-instances)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.ShelveInstance(context.Background(), instanceId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.ShelveInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShelveInstance`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.ShelveInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID &lt;br/&gt; - [List instances](https://docs.kakaocloud.com/openapi/bcs/list-instances)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiShelveInstanceRequest struct via the builder pattern


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


## SoftRebootInstance

> interface{} SoftRebootInstance(ctx, instanceId).XAuthToken(xAuthToken).Execute()

Soft reboot instance 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/bcs"
)

func main() {
	instanceId := "instanceId_example" // string | 인스턴스의 고유 ID <br/> - [List instances](https://docs.kakaocloud.com/openapi/bcs/list-instances)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.SoftRebootInstance(context.Background(), instanceId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.SoftRebootInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SoftRebootInstance`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.SoftRebootInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID &lt;br/&gt; - [List instances](https://docs.kakaocloud.com/openapi/bcs/list-instances)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiSoftRebootInstanceRequest struct via the builder pattern


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


## StartInstance

> interface{} StartInstance(ctx, instanceId).XAuthToken(xAuthToken).Execute()

Start instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/bcs"
)

func main() {
	instanceId := "instanceId_example" // string | 인스턴스의 고유 ID <br/> - [List instances](https://docs.kakaocloud.com/openapi/bcs/list-instances)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.StartInstance(context.Background(), instanceId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.StartInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StartInstance`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.StartInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID &lt;br/&gt; - [List instances](https://docs.kakaocloud.com/openapi/bcs/list-instances)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartInstanceRequest struct via the builder pattern


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


## StopInstance

> interface{} StopInstance(ctx, instanceId).XAuthToken(xAuthToken).Execute()

Stop instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/bcs"
)

func main() {
	instanceId := "instanceId_example" // string | 인스턴스의 고유 ID <br/> - [List instances](https://docs.kakaocloud.com/openapi/bcs/list-instances)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.StopInstance(context.Background(), instanceId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.StopInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopInstance`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.StopInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID &lt;br/&gt; - [List instances](https://docs.kakaocloud.com/openapi/bcs/list-instances)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopInstanceRequest struct via the builder pattern


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


## UnshelveInstance

> interface{} UnshelveInstance(ctx, instanceId).XAuthToken(xAuthToken).Execute()

Unshelve instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/bcs"
)

func main() {
	instanceId := "instanceId_example" // string | 인스턴스의 고유 ID <br/> - [List instances](https://docs.kakaocloud.com/openapi/bcs/list-instances)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.UnshelveInstance(context.Background(), instanceId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.UnshelveInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnshelveInstance`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.UnshelveInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID &lt;br/&gt; - [List instances](https://docs.kakaocloud.com/openapi/bcs/list-instances)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnshelveInstanceRequest struct via the builder pattern


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


## UpdateInstance

> UpdateInstanceResponse UpdateInstance(ctx, instanceId).XAuthToken(xAuthToken).UpdateInstanceRequest(updateInstanceRequest).Execute()

Update instance



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/bcs"
)

func main() {
	instanceId := "instanceId_example" // string | 인스턴스의 고유 ID <br/> - [List instances](https://docs.kakaocloud.com/openapi/bcs/list-instances)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	updateInstanceRequest := *openapiclient.NewUpdateInstanceRequest(*openapiclient.NewUpdateInstance()) // UpdateInstanceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.UpdateInstance(context.Background(), instanceId).XAuthToken(xAuthToken).UpdateInstanceRequest(updateInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.UpdateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInstance`: UpdateInstanceResponse
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.UpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID &lt;br/&gt; - [List instances](https://docs.kakaocloud.com/openapi/bcs/list-instances)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **updateInstanceRequest** | [**UpdateInstanceRequest**](UpdateInstanceRequest.md) |  | 

### Return type

[**UpdateInstanceResponse**](UpdateInstanceResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInstanceVolume

> interface{} UpdateInstanceVolume(ctx, instanceId, volumeId).XAuthToken(xAuthToken).UpdateInstanceVolumeRequest(updateInstanceVolumeRequest).Execute()

Update instance volume



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/bcs"
)

func main() {
	instanceId := "instanceId_example" // string | 인스턴스의 고유 ID <br/> - [List instances](https://docs.kakaocloud.com/openapi/bcs/list-instances)에서 확인
	volumeId := "volumeId_example" // string | 볼륨의 고유 ID <br/> - [List volumes](https://docs.kakaocloud.com/openapi/bcs/list-volumes)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	updateInstanceVolumeRequest := *openapiclient.NewUpdateInstanceVolumeRequest(*openapiclient.NewUpdateInstanceVolume()) // UpdateInstanceVolumeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.UpdateInstanceVolume(context.Background(), instanceId, volumeId).XAuthToken(xAuthToken).UpdateInstanceVolumeRequest(updateInstanceVolumeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.UpdateInstanceVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInstanceVolume`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.UpdateInstanceVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID &lt;br/&gt; - [List instances](https://docs.kakaocloud.com/openapi/bcs/list-instances)에서 확인 | 
**volumeId** | **string** | 볼륨의 고유 ID &lt;br/&gt; - [List volumes](https://docs.kakaocloud.com/openapi/bcs/list-volumes)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInstanceVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **updateInstanceVolumeRequest** | [**UpdateInstanceVolumeRequest**](UpdateInstanceVolumeRequest.md) |  | 

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


## UpdateNetworkInterfaceAllowedAddresses

> UpdateNetworkInterfaceAllowedAddressesResponse UpdateNetworkInterfaceAllowedAddresses(ctx, instanceId, networkInterfaceId).XAuthToken(xAuthToken).UpdateNetworkInterfaceAllowedAddressesRequest(updateNetworkInterfaceAllowedAddressesRequest).Execute()

Update network interface allowed addresses



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/bcs"
)

func main() {
	instanceId := "instanceId_example" // string | 인스턴스의 고유 ID <br/> - [List instances](https://docs.kakaocloud.com/openapi/bcs/list-instances)에서 확인
	networkInterfaceId := "networkInterfaceId_example" // string | 인스턴스에 연결된 네트워크 인터페이스 ID <br/> - [List instance network interfaces](https://docs.kakaocloud.com/openapi/bcs/list-instance-network-interfaces)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	updateNetworkInterfaceAllowedAddressesRequest := *openapiclient.NewUpdateNetworkInterfaceAllowedAddressesRequest([]openapiclient.UpdateNetworkInterfaceAllowedAddresses{*openapiclient.NewUpdateNetworkInterfaceAllowedAddresses("IpAddress_example")}) // UpdateNetworkInterfaceAllowedAddressesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.UpdateNetworkInterfaceAllowedAddresses(context.Background(), instanceId, networkInterfaceId).XAuthToken(xAuthToken).UpdateNetworkInterfaceAllowedAddressesRequest(updateNetworkInterfaceAllowedAddressesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.UpdateNetworkInterfaceAllowedAddresses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkInterfaceAllowedAddresses`: UpdateNetworkInterfaceAllowedAddressesResponse
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.UpdateNetworkInterfaceAllowedAddresses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID &lt;br/&gt; - [List instances](https://docs.kakaocloud.com/openapi/bcs/list-instances)에서 확인 | 
**networkInterfaceId** | **string** | 인스턴스에 연결된 네트워크 인터페이스 ID &lt;br/&gt; - [List instance network interfaces](https://docs.kakaocloud.com/openapi/bcs/list-instance-network-interfaces)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkInterfaceAllowedAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **updateNetworkInterfaceAllowedAddressesRequest** | [**UpdateNetworkInterfaceAllowedAddressesRequest**](UpdateNetworkInterfaceAllowedAddressesRequest.md) |  | 

### Return type

[**UpdateNetworkInterfaceAllowedAddressesResponse**](UpdateNetworkInterfaceAllowedAddressesResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

