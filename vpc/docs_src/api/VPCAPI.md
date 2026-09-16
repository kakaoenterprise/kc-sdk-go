# \VPCAPI

All URIs are relative to *https://vpc.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVpc**](VPCAPI.md#CreateVpc) | **Post** /api/v1/vpcs | Create VPC 
[**DeleteVpc**](VPCAPI.md#DeleteVpc) | **Delete** /api/v1/vpcs/{vpc_id} | Delete VPC
[**GetVpc**](VPCAPI.md#GetVpc) | **Get** /api/v1/vpcs/{vpc_id} | Get VPC
[**ListVpcs**](VPCAPI.md#ListVpcs) | **Get** /api/v1/vpcs | List VPCs
[**UpdateVpc**](VPCAPI.md#UpdateVpc) | **Put** /api/v1/vpcs/{vpc_id} | Update VPC



## CreateVpc

> CreateVpcResponse CreateVpc(ctx).CreateVpcRequest(createVpcRequest).Execute()

Create VPC 



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
	createVpcRequest := *openapiclient.NewCreateVpcRequest(*openapiclient.NewCreateVpc("Name_example", "CidrBlock_example")) // CreateVpcRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPCAPI.CreateVpc(context.Background()).CreateVpcRequest(createVpcRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCAPI.CreateVpc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVpc`: CreateVpcResponse
	fmt.Fprintf(os.Stdout, "Response from `VPCAPI.CreateVpc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVpcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createVpcRequest** | [**CreateVpcRequest**](CreateVpcRequest.md) |  | 

### Return type

[**CreateVpcResponse**](CreateVpcResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVpc

> DeleteVpc(ctx, vpcId).Execute()

Delete VPC



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
	vpcId := "vpcId_example" // string | 삭제할 VPC ID - [List VPCs](/openapi/networking/vpc/list-vpcs)에서 확인

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VPCAPI.DeleteVpc(context.Background(), vpcId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCAPI.DeleteVpc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpcId** | **string** | 삭제할 VPC ID - [List VPCs](/openapi/networking/vpc/list-vpcs)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVpcRequest struct via the builder pattern


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


## GetVpc

> GetVpcResponse GetVpc(ctx, vpcId).Execute()

Get VPC



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
	vpcId := "vpcId_example" // string | 조회할 VPC ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPCAPI.GetVpc(context.Background(), vpcId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCAPI.GetVpc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVpc`: GetVpcResponse
	fmt.Fprintf(os.Stdout, "Response from `VPCAPI.GetVpc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpcId** | **string** | 조회할 VPC ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVpcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetVpcResponse**](GetVpcResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVpcs

> ListVpcsResponse ListVpcs(ctx).Id(id).Name(name).CidrBlock(cidrBlock).ProvisioningStatus(provisioningStatus).IsDefault(isDefault).SortKeys(sortKeys).SortDirs(sortDirs).CreatedAt(createdAt).UpdatedAt(updatedAt).Offset(offset).Limit(limit).Execute()

List VPCs



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
	id := "id_example" // string | 조회할 VPC ID (optional)
	name := "name_example" // string | 조회할 VPC 이름 (optional)
	cidrBlock := "cidrBlock_example" // string | VPC의 IPv4 CIDR 블록 (예: `10.0.0.0/16`) (optional)
	provisioningStatus := openapiclient.ProvisioningStatus("ACTIVE") // ProvisioningStatus | 프로비저닝 상태 (optional)
	isDefault := true // bool | 기본 VPC 여부 (optional)
	sortKeys := "sortKeys_example" // string | 정렬할 필드를 콤마(,)로 구분 (optional)
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`) (optional)
	createdAt := "createdAt_example" // string | 리소스가 생성된 시간 - ISO 8601 형식 - UTC 기준 (optional)
	updatedAt := "updatedAt_example" // string | 리소스가 마지막으로 수정된 시간 - ISO 8601 형식 - UTC 기준 (optional)
	offset := int32(56) // int32 | 조회 시작 위치 (optional)
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPCAPI.ListVpcs(context.Background()).Id(id).Name(name).CidrBlock(cidrBlock).ProvisioningStatus(provisioningStatus).IsDefault(isDefault).SortKeys(sortKeys).SortDirs(sortDirs).CreatedAt(createdAt).UpdatedAt(updatedAt).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCAPI.ListVpcs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVpcs`: ListVpcsResponse
	fmt.Fprintf(os.Stdout, "Response from `VPCAPI.ListVpcs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVpcsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | 조회할 VPC ID | 
 **name** | **string** | 조회할 VPC 이름 | 
 **cidrBlock** | **string** | VPC의 IPv4 CIDR 블록 (예: &#x60;10.0.0.0/16&#x60;) | 
 **provisioningStatus** | [**ProvisioningStatus**](ProvisioningStatus.md) | 프로비저닝 상태 | 
 **isDefault** | **bool** | 기본 VPC 여부 | 
 **sortKeys** | **string** | 정렬할 필드를 콤마(,)로 구분 | 
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;) | 
 **createdAt** | **string** | 리소스가 생성된 시간 - ISO 8601 형식 - UTC 기준 | 
 **updatedAt** | **string** | 리소스가 마지막으로 수정된 시간 - ISO 8601 형식 - UTC 기준 | 
 **offset** | **int32** | 조회 시작 위치 | 
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | 

### Return type

[**ListVpcsResponse**](ListVpcsResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVpc

> UpdateVpcResponse UpdateVpc(ctx, vpcId).UpdateVpcRequest(updateVpcRequest).Execute()

Update VPC



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
	vpcId := "vpcId_example" // string | 수정할 VPC ID - [List VPCs](/openapi/networking/vpc/list-vpcs)에서 확인
	updateVpcRequest := *openapiclient.NewUpdateVpcRequest(*openapiclient.NewUpdateVpc()) // UpdateVpcRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPCAPI.UpdateVpc(context.Background(), vpcId).UpdateVpcRequest(updateVpcRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCAPI.UpdateVpc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateVpc`: UpdateVpcResponse
	fmt.Fprintf(os.Stdout, "Response from `VPCAPI.UpdateVpc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpcId** | **string** | 수정할 VPC ID - [List VPCs](/openapi/networking/vpc/list-vpcs)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVpcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateVpcRequest** | [**UpdateVpcRequest**](UpdateVpcRequest.md) |  | 

### Return type

[**UpdateVpcResponse**](UpdateVpcResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

