# \NetworkInterfaceAPI

All URIs are relative to *https://vpc.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkInterface**](NetworkInterfaceAPI.md#CreateNetworkInterface) | **Post** /api/v1/network-interfaces | Create network interface
[**DeleteNetworkInterface**](NetworkInterfaceAPI.md#DeleteNetworkInterface) | **Delete** /api/v1/network-interfaces/{network_interface_id} | Delete network interface
[**GetNetworkInterface**](NetworkInterfaceAPI.md#GetNetworkInterface) | **Get** /api/v1/network-interfaces/{network_interface_id} | Get network interface
[**ListNetworkInterfaces**](NetworkInterfaceAPI.md#ListNetworkInterfaces) | **Get** /api/v1/network-interfaces | List network interfaces
[**UpdateNetworkInterface**](NetworkInterfaceAPI.md#UpdateNetworkInterface) | **Put** /api/v1/network-interfaces/{network_interface_id} | Update network interface
[**UpdateNetworkInterfaceAllowedAddresses**](NetworkInterfaceAPI.md#UpdateNetworkInterfaceAllowedAddresses) | **Put** /api/v1/network-interfaces/{network_interface_id}/allowed-address-pairs | Update network interface allowed addresses



## CreateNetworkInterface

> CreateNetworkInterfaceResponse CreateNetworkInterface(ctx).CreateNetworkInterfaceRequest(createNetworkInterfaceRequest).Execute()

Create network interface



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/vpc"
)

func main() {
	createNetworkInterfaceRequest := *openapiclient.NewCreateNetworkInterfaceRequest(*openapiclient.NewCreateNetworkInterface("Name_example", "SubnetId_example", []string{"SecurityGroups_example"})) // CreateNetworkInterfaceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkInterfaceAPI.CreateNetworkInterface(context.Background()).CreateNetworkInterfaceRequest(createNetworkInterfaceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkInterfaceAPI.CreateNetworkInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkInterface`: CreateNetworkInterfaceResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworkInterfaceAPI.CreateNetworkInterface`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNetworkInterfaceRequest** | [**CreateNetworkInterfaceRequest**](CreateNetworkInterfaceRequest.md) |  | 

### Return type

[**CreateNetworkInterfaceResponse**](CreateNetworkInterfaceResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkInterface

> DeleteNetworkInterface(ctx, networkInterfaceId).Execute()

Delete network interface



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/vpc"
)

func main() {
	networkInterfaceId := "networkInterfaceId_example" // string | 삭제할 네트워크 인터페이스 ID - [List network interfaces](/openapi/networking/vpc/list-network-interfaces)에서 확인

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkInterfaceAPI.DeleteNetworkInterface(context.Background(), networkInterfaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkInterfaceAPI.DeleteNetworkInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkInterfaceId** | **string** | 삭제할 네트워크 인터페이스 ID - [List network interfaces](/openapi/networking/vpc/list-network-interfaces)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkInterfaceRequest struct via the builder pattern


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


## GetNetworkInterface

> GetNetworkInterfaceResponse GetNetworkInterface(ctx, networkInterfaceId).Execute()

Get network interface



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/vpc"
)

func main() {
	networkInterfaceId := "networkInterfaceId_example" // string | 조회할 네트워크 인터페이스 ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkInterfaceAPI.GetNetworkInterface(context.Background(), networkInterfaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkInterfaceAPI.GetNetworkInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkInterface`: GetNetworkInterfaceResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworkInterfaceAPI.GetNetworkInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkInterfaceId** | **string** | 조회할 네트워크 인터페이스 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNetworkInterfaceResponse**](GetNetworkInterfaceResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkInterfaces

> ListNetworkInterfacesResponse ListNetworkInterfaces(ctx).Id(id).Name(name).Status(status).PrivateIp(privateIp).PublicIp(publicIp).DeviceOwner(deviceOwner).DeviceId(deviceId).SubnetId(subnetId).MacAddress(macAddress).SecurityGroupId(securityGroupId).SecurityGroupName(securityGroupName).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Offset(offset).Limit(limit).Execute()

List network interfaces



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/vpc"
)

func main() {
	id := "id_example" // string | 네트워크 인터페이스 ID (optional)
	name := "name_example" // string | 네트워크 인터페이스 이름 (optional)
	status := openapiclient.NetworkInterfaceStatus("available") // NetworkInterfaceStatus | 네트워크 인터페이스의 상태 (optional)
	privateIp := "privateIp_example" // string | 프라이빗 IP 주소 (IPv4 형식) (optional)
	publicIp := "publicIp_example" // string | 퍼블릭 IP 주소 (optional)
	deviceOwner := "deviceOwner_example" // string | 인터페이스가 연결된 리소스의 소유자 유형 (예: 인스턴스, 라우터 등) (optional)
	deviceId := "deviceId_example" // string | 연결된 디바이스 ID (예: 인스턴스 ID 등) (optional)
	subnetId := "subnetId_example" // string | 연결된 서브넷 ID (optional)
	macAddress := "macAddress_example" // string | 네트워크 인터페이스의 MAC 주소 (optional)
	securityGroupId := "securityGroupId_example" // string | 보안 그룹의 고유 ID (optional)
	securityGroupName := "securityGroupName_example" // string | 보안 그룹 이름 (optional)
	createdAt := "createdAt_example" // string | 리소스가 생성된 시간 - ISO 8601 형식 - UTC 기준 (optional)
	updatedAt := "updatedAt_example" // string | 리소스가 마지막으로 수정된 시간 - ISO 8601 형식 - UTC 기준 (optional)
	sortKeys := "sortKeys_example" // string | 정렬할 필드를 콤마(,)로 구분 (optional)
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`) (optional)
	offset := int32(56) // int32 | 조회 시작 위치 (optional)
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkInterfaceAPI.ListNetworkInterfaces(context.Background()).Id(id).Name(name).Status(status).PrivateIp(privateIp).PublicIp(publicIp).DeviceOwner(deviceOwner).DeviceId(deviceId).SubnetId(subnetId).MacAddress(macAddress).SecurityGroupId(securityGroupId).SecurityGroupName(securityGroupName).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkInterfaceAPI.ListNetworkInterfaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNetworkInterfaces`: ListNetworkInterfacesResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworkInterfaceAPI.ListNetworkInterfaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkInterfacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | 네트워크 인터페이스 ID | 
 **name** | **string** | 네트워크 인터페이스 이름 | 
 **status** | [**NetworkInterfaceStatus**](NetworkInterfaceStatus.md) | 네트워크 인터페이스의 상태 | 
 **privateIp** | **string** | 프라이빗 IP 주소 (IPv4 형식) | 
 **publicIp** | **string** | 퍼블릭 IP 주소 | 
 **deviceOwner** | **string** | 인터페이스가 연결된 리소스의 소유자 유형 (예: 인스턴스, 라우터 등) | 
 **deviceId** | **string** | 연결된 디바이스 ID (예: 인스턴스 ID 등) | 
 **subnetId** | **string** | 연결된 서브넷 ID | 
 **macAddress** | **string** | 네트워크 인터페이스의 MAC 주소 | 
 **securityGroupId** | **string** | 보안 그룹의 고유 ID | 
 **securityGroupName** | **string** | 보안 그룹 이름 | 
 **createdAt** | **string** | 리소스가 생성된 시간 - ISO 8601 형식 - UTC 기준 | 
 **updatedAt** | **string** | 리소스가 마지막으로 수정된 시간 - ISO 8601 형식 - UTC 기준 | 
 **sortKeys** | **string** | 정렬할 필드를 콤마(,)로 구분 | 
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;) | 
 **offset** | **int32** | 조회 시작 위치 | 
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | 

### Return type

[**ListNetworkInterfacesResponse**](ListNetworkInterfacesResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkInterface

> UpdateNetworkInterfaceResponse UpdateNetworkInterface(ctx, networkInterfaceId).UpdateNetworkInterfaceRequest(updateNetworkInterfaceRequest).Execute()

Update network interface



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/vpc"
)

func main() {
	networkInterfaceId := "networkInterfaceId_example" // string | 수정할 네트워크 인터페이스 ID - [List network interfaces](/openapi/networking/vpc/list-network-interfaces)에서 확인
	updateNetworkInterfaceRequest := *openapiclient.NewUpdateNetworkInterfaceRequest(*openapiclient.NewUpdateNetworkInterface()) // UpdateNetworkInterfaceRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkInterfaceAPI.UpdateNetworkInterface(context.Background(), networkInterfaceId).UpdateNetworkInterfaceRequest(updateNetworkInterfaceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkInterfaceAPI.UpdateNetworkInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkInterface`: UpdateNetworkInterfaceResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworkInterfaceAPI.UpdateNetworkInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkInterfaceId** | **string** | 수정할 네트워크 인터페이스 ID - [List network interfaces](/openapi/networking/vpc/list-network-interfaces)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkInterfaceRequest** | [**UpdateNetworkInterfaceRequest**](UpdateNetworkInterfaceRequest.md) |  | 

### Return type

[**UpdateNetworkInterfaceResponse**](UpdateNetworkInterfaceResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkInterfaceAllowedAddresses

> UpdateNetworkInterfaceAllowedAddressesResponse UpdateNetworkInterfaceAllowedAddresses(ctx, networkInterfaceId).UpdateNetworkInterfaceAllowedAddressesRequest(updateNetworkInterfaceAllowedAddressesRequest).Execute()

Update network interface allowed addresses



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/vpc"
)

func main() {
	networkInterfaceId := "networkInterfaceId_example" // string | 설정할 대상 네트워크 인터페이스 ID - [List network interfaces](/openapi/networking/vpc/list-network-interfaces)에서 확인
	updateNetworkInterfaceAllowedAddressesRequest := *openapiclient.NewUpdateNetworkInterfaceAllowedAddressesRequest([]openapiclient.UpdateNetworkInterfaceAllowedAddresses{*openapiclient.NewUpdateNetworkInterfaceAllowedAddresses("IpAddress_example")}) // UpdateNetworkInterfaceAllowedAddressesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkInterfaceAPI.UpdateNetworkInterfaceAllowedAddresses(context.Background(), networkInterfaceId).UpdateNetworkInterfaceAllowedAddressesRequest(updateNetworkInterfaceAllowedAddressesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkInterfaceAPI.UpdateNetworkInterfaceAllowedAddresses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkInterfaceAllowedAddresses`: UpdateNetworkInterfaceAllowedAddressesResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworkInterfaceAPI.UpdateNetworkInterfaceAllowedAddresses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkInterfaceId** | **string** | 설정할 대상 네트워크 인터페이스 ID - [List network interfaces](/openapi/networking/vpc/list-network-interfaces)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkInterfaceAllowedAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

