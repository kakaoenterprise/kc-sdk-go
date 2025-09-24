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

> BnsVpcV1ApiCreateNetworkInterfaceModelResponseNetworkInterfaceModel CreateNetworkInterface(ctx).XAuthToken(xAuthToken).BodyCreateNetworkInterface(bodyCreateNetworkInterface).Execute()

Create network interface



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
	bodyCreateNetworkInterface := *openapiclient.NewBodyCreateNetworkInterface(*openapiclient.NewCreateNetworkInterfaceModel("Name_example", "SubnetId_example", []string{"SecurityGroups_example"})) // BodyCreateNetworkInterface | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkInterfaceAPI.CreateNetworkInterface(context.Background()).XAuthToken(xAuthToken).BodyCreateNetworkInterface(bodyCreateNetworkInterface).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkInterfaceAPI.CreateNetworkInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNetworkInterface`: BnsVpcV1ApiCreateNetworkInterfaceModelResponseNetworkInterfaceModel
	fmt.Fprintf(os.Stdout, "Response from `NetworkInterfaceAPI.CreateNetworkInterface`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyCreateNetworkInterface** | [**BodyCreateNetworkInterface**](BodyCreateNetworkInterface.md) |  | 

### Return type

[**BnsVpcV1ApiCreateNetworkInterfaceModelResponseNetworkInterfaceModel**](BnsVpcV1ApiCreateNetworkInterfaceModelResponseNetworkInterfaceModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkInterface

> DeleteNetworkInterface(ctx, networkInterfaceId).XAuthToken(xAuthToken).Execute()

Delete network interface



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
	networkInterfaceId := "networkInterfaceId_example" // string | 삭제할 네트워크 인터페이스 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.NetworkInterfaceAPI.DeleteNetworkInterface(context.Background(), networkInterfaceId).XAuthToken(xAuthToken).Execute()
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
**networkInterfaceId** | **string** | 삭제할 네트워크 인터페이스 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkInterfaceRequest struct via the builder pattern


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


## GetNetworkInterface

> BnsVpcV1ApiGetNetworkInterfaceModelResponseNetworkInterfaceModel GetNetworkInterface(ctx, networkInterfaceId).XAuthToken(xAuthToken).Execute()

Get network interface



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
	networkInterfaceId := "networkInterfaceId_example" // string | 조회할 네트워크 인터페이스 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkInterfaceAPI.GetNetworkInterface(context.Background(), networkInterfaceId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkInterfaceAPI.GetNetworkInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworkInterface`: BnsVpcV1ApiGetNetworkInterfaceModelResponseNetworkInterfaceModel
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

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**BnsVpcV1ApiGetNetworkInterfaceModelResponseNetworkInterfaceModel**](BnsVpcV1ApiGetNetworkInterfaceModelResponseNetworkInterfaceModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkInterfaces

> NetworkInterfaceListModel ListNetworkInterfaces(ctx).XAuthToken(xAuthToken).Id(id).Name(name).Status(status).PrivateIp(privateIp).PublicIp(publicIp).DeviceId(deviceId).DeviceOwner(deviceOwner).SubnetId(subnetId).MacAddress(macAddress).SecurityGroupId(securityGroupId).SecurityGroupName(securityGroupName).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Limit(limit).Offset(offset).Execute()

List network interfaces



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
	id := "id_example" // string | 네트워크 인터페이스 ID (optional)
	name := "name_example" // string | 네트워크 인터페이스 이름 (optional)
	status := openapiclient.NetworkInterfaceStatus("available") // NetworkInterfaceStatus | 네트워크 인터페이스의 상태 <br/> - `available`: 사용 가능 상태<br/> - `in_use`: 사용 중 상태 (optional)
	privateIp := "privateIp_example" // string | 프라이빗 IP 주소 (IPv4 형식)  (optional)
	publicIp := "publicIp_example" // string | 퍼블릭 IP 주소  (optional)
	deviceId := "deviceId_example" // string | 연결된 디바이스 ID (예: 인스턴스 ID 등)  (optional)
	deviceOwner := "deviceOwner_example" // string | 인터페이스가 연결된 리소스의 소유자 유형 (예: 인스턴스, 라우터 등)  (optional)
	subnetId := "subnetId_example" // string | 연결된 서브넷 ID  (optional)
	macAddress := "macAddress_example" // string | 네트워크 인터페이스의 MAC 주소 (optional)
	securityGroupId := "securityGroupId_example" // string | 보안 그룹의 고유 ID (optional)
	securityGroupName := "securityGroupName_example" // string | 보안 그룹 이름  (optional)
	createdAt := "createdAt_example" // string | 리소스가 생성된 시간 <br/> - ISO_8601 형식  <br/> - UTC 기준 (optional)
	updatedAt := "updatedAt_example" // string | 리소스가 마지막으로 수정된 시간 <br/> - ISO_8601 형식  <br/> - UTC 기준 (optional)
	sortKeys := "sortKeys_example" // string | 정렬할 필드를 콤마(,)로 구분     (optional) (default to "created_at")
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`)   (optional) (default to "desc")
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional) (default to 20)
	offset := int32(56) // int32 | 조회 시작 위치 (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkInterfaceAPI.ListNetworkInterfaces(context.Background()).XAuthToken(xAuthToken).Id(id).Name(name).Status(status).PrivateIp(privateIp).PublicIp(publicIp).DeviceId(deviceId).DeviceOwner(deviceOwner).SubnetId(subnetId).MacAddress(macAddress).SecurityGroupId(securityGroupId).SecurityGroupName(securityGroupName).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkInterfaceAPI.ListNetworkInterfaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNetworkInterfaces`: NetworkInterfaceListModel
	fmt.Fprintf(os.Stdout, "Response from `NetworkInterfaceAPI.ListNetworkInterfaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkInterfacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **id** | **string** | 네트워크 인터페이스 ID | 
 **name** | **string** | 네트워크 인터페이스 이름 | 
 **status** | [**NetworkInterfaceStatus**](NetworkInterfaceStatus.md) | 네트워크 인터페이스의 상태 &lt;br/&gt; - &#x60;available&#x60;: 사용 가능 상태&lt;br/&gt; - &#x60;in_use&#x60;: 사용 중 상태 | 
 **privateIp** | **string** | 프라이빗 IP 주소 (IPv4 형식)  | 
 **publicIp** | **string** | 퍼블릭 IP 주소  | 
 **deviceId** | **string** | 연결된 디바이스 ID (예: 인스턴스 ID 등)  | 
 **deviceOwner** | **string** | 인터페이스가 연결된 리소스의 소유자 유형 (예: 인스턴스, 라우터 등)  | 
 **subnetId** | **string** | 연결된 서브넷 ID  | 
 **macAddress** | **string** | 네트워크 인터페이스의 MAC 주소 | 
 **securityGroupId** | **string** | 보안 그룹의 고유 ID | 
 **securityGroupName** | **string** | 보안 그룹 이름  | 
 **createdAt** | **string** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
 **updatedAt** | **string** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
 **sortKeys** | **string** | 정렬할 필드를 콤마(,)로 구분     | [default to &quot;created_at&quot;]
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;)   | [default to &quot;desc&quot;]
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | [default to 20]
 **offset** | **int32** | 조회 시작 위치 | [default to 0]

### Return type

[**NetworkInterfaceListModel**](NetworkInterfaceListModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkInterface

> BnsVpcV1ApiUpdateNetworkInterfaceModelResponseNetworkInterfaceModel UpdateNetworkInterface(ctx, networkInterfaceId).XAuthToken(xAuthToken).BodyUpdateNetworkInterface(bodyUpdateNetworkInterface).Execute()

Update network interface



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
	networkInterfaceId := "networkInterfaceId_example" // string | 수정할 네트워크 인터페이스 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bodyUpdateNetworkInterface := *openapiclient.NewBodyUpdateNetworkInterface(*openapiclient.NewEditNetworkInterfaceModel()) // BodyUpdateNetworkInterface | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkInterfaceAPI.UpdateNetworkInterface(context.Background(), networkInterfaceId).XAuthToken(xAuthToken).BodyUpdateNetworkInterface(bodyUpdateNetworkInterface).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkInterfaceAPI.UpdateNetworkInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkInterface`: BnsVpcV1ApiUpdateNetworkInterfaceModelResponseNetworkInterfaceModel
	fmt.Fprintf(os.Stdout, "Response from `NetworkInterfaceAPI.UpdateNetworkInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkInterfaceId** | **string** | 수정할 네트워크 인터페이스 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyUpdateNetworkInterface** | [**BodyUpdateNetworkInterface**](BodyUpdateNetworkInterface.md) |  | 

### Return type

[**BnsVpcV1ApiUpdateNetworkInterfaceModelResponseNetworkInterfaceModel**](BnsVpcV1ApiUpdateNetworkInterfaceModelResponseNetworkInterfaceModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkInterfaceAllowedAddresses

> BnsVpcV1ApiUpdateNetworkInterfaceAllowedAddressesModelResponseNetworkInterfaceModel UpdateNetworkInterfaceAllowedAddresses(ctx, networkInterfaceId).XAuthToken(xAuthToken).BodyUpdateNetworkInterfaceAllowedAddresses(bodyUpdateNetworkInterfaceAllowedAddresses).Execute()

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
	networkInterfaceId := "networkInterfaceId_example" // string | 설정할 대상 네트워크 인터페이스 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bodyUpdateNetworkInterfaceAllowedAddresses := *openapiclient.NewBodyUpdateNetworkInterfaceAllowedAddresses([]openapiclient.EditAllowedAddressPairModel{*openapiclient.NewEditAllowedAddressPairModel()}) // BodyUpdateNetworkInterfaceAllowedAddresses | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkInterfaceAPI.UpdateNetworkInterfaceAllowedAddresses(context.Background(), networkInterfaceId).XAuthToken(xAuthToken).BodyUpdateNetworkInterfaceAllowedAddresses(bodyUpdateNetworkInterfaceAllowedAddresses).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkInterfaceAPI.UpdateNetworkInterfaceAllowedAddresses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNetworkInterfaceAllowedAddresses`: BnsVpcV1ApiUpdateNetworkInterfaceAllowedAddressesModelResponseNetworkInterfaceModel
	fmt.Fprintf(os.Stdout, "Response from `NetworkInterfaceAPI.UpdateNetworkInterfaceAllowedAddresses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkInterfaceId** | **string** | 설정할 대상 네트워크 인터페이스 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkInterfaceAllowedAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyUpdateNetworkInterfaceAllowedAddresses** | [**BodyUpdateNetworkInterfaceAllowedAddresses**](BodyUpdateNetworkInterfaceAllowedAddresses.md) |  | 

### Return type

[**BnsVpcV1ApiUpdateNetworkInterfaceAllowedAddressesModelResponseNetworkInterfaceModel**](BnsVpcV1ApiUpdateNetworkInterfaceAllowedAddressesModelResponseNetworkInterfaceModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

