# \NetworkInterfaceAPI

All URIs are relative to *https://network.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssociatePublicIp**](NetworkInterfaceAPI.md#AssociatePublicIp) | **Put** /api/v1/public-ips/{public_ip_id}/network-interfaces/{network_interface_id} | Associate public IP
[**DisassociatePublicIp**](NetworkInterfaceAPI.md#DisassociatePublicIp) | **Delete** /api/v1/public-ips/{public_ip_id}/network-interfaces/{network_interface_id} | Disassociate public IP



## AssociatePublicIp

> AssociatePublicIpResponse AssociatePublicIp(ctx, publicIpId, networkInterfaceId).Execute()

Associate public IP



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/network"
)

func main() {
	publicIpId := "publicIpId_example" // string | 연결할 퍼블릭 IP의 ID - [List public IPs](/openapi/networking/vpc/list-public-ips)에서 확인
	networkInterfaceId := "networkInterfaceId_example" // string | 퍼블릭 IP를 연결할 대상 네트워크 인터페이스의 ID - [List network interfaces](/openapi/networking/vpc/list-network-interfaces)에서 확인

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkInterfaceAPI.AssociatePublicIp(context.Background(), publicIpId, networkInterfaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkInterfaceAPI.AssociatePublicIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssociatePublicIp`: AssociatePublicIpResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworkInterfaceAPI.AssociatePublicIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicIpId** | **string** | 연결할 퍼블릭 IP의 ID - [List public IPs](/openapi/networking/vpc/list-public-ips)에서 확인 | 
**networkInterfaceId** | **string** | 퍼블릭 IP를 연결할 대상 네트워크 인터페이스의 ID - [List network interfaces](/openapi/networking/vpc/list-network-interfaces)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssociatePublicIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## DisassociatePublicIp

> DisassociatePublicIpResponse DisassociatePublicIp(ctx, publicIpId, networkInterfaceId).Execute()

Disassociate public IP



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/network"
)

func main() {
	publicIpId := "publicIpId_example" // string | 연결 해제할 퍼블릭 IP의 ID - [List public IPs](/openapi/networking/vpc/list-public-ips)에서 확인
	networkInterfaceId := "networkInterfaceId_example" // string | 연결 해제 대상 네트워크 인터페이스의 ID - [List network interfaces](/openapi/networking/vpc/list-network-interfaces)에서 확인

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkInterfaceAPI.DisassociatePublicIp(context.Background(), publicIpId, networkInterfaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkInterfaceAPI.DisassociatePublicIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DisassociatePublicIp`: DisassociatePublicIpResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworkInterfaceAPI.DisassociatePublicIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**publicIpId** | **string** | 연결 해제할 퍼블릭 IP의 ID - [List public IPs](/openapi/networking/vpc/list-public-ips)에서 확인 | 
**networkInterfaceId** | **string** | 연결 해제 대상 네트워크 인터페이스의 ID - [List network interfaces](/openapi/networking/vpc/list-network-interfaces)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisassociatePublicIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DisassociatePublicIpResponse**](DisassociatePublicIpResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

