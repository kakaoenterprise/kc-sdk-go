# \SubnetAPI

All URIs are relative to *https://vpc.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubnet**](SubnetAPI.md#CreateSubnet) | **Post** /api/v1/subnets | Create subnet
[**DeleteSubnet**](SubnetAPI.md#DeleteSubnet) | **Delete** /api/v1/subnets/{subnet_id} | Delete subnet
[**GetSubnet**](SubnetAPI.md#GetSubnet) | **Get** /api/v1/subnets/{subnet_id} | Get subnet
[**ListSubnetSharedProjects**](SubnetAPI.md#ListSubnetSharedProjects) | **Get** /api/v1/subnets/{subnet_id}/projects | List subnet shared projects
[**ListSubnets**](SubnetAPI.md#ListSubnets) | **Get** /api/v1/subnets | List subnets
[**ShareSubnet**](SubnetAPI.md#ShareSubnet) | **Post** /api/v1/subnets/{subnet_id}/projects/{project_id} | Share subnet
[**UnshareSubnet**](SubnetAPI.md#UnshareSubnet) | **Delete** /api/v1/subnets/{subnet_id}/projects/{project_id} | Unshare subnet
[**UpdateSubnet**](SubnetAPI.md#UpdateSubnet) | **Put** /api/v1/subnets/{subnet_id} | Update subnet



## CreateSubnet

> CreateSubnetResponse CreateSubnet(ctx).CreateSubnetRequest(createSubnetRequest).Execute()

Create subnet



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
	createSubnetRequest := *openapiclient.NewCreateSubnetRequest(*openapiclient.NewCreateSubnet(openapiclient.AvailabilityZone("kr-central-2-a"), "CidrBlock_example", "Name_example", "VpcId_example")) // CreateSubnetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubnetAPI.CreateSubnet(context.Background()).CreateSubnetRequest(createSubnetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubnetAPI.CreateSubnet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSubnet`: CreateSubnetResponse
	fmt.Fprintf(os.Stdout, "Response from `SubnetAPI.CreateSubnet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubnetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSubnetRequest** | [**CreateSubnetRequest**](CreateSubnetRequest.md) |  | 

### Return type

[**CreateSubnetResponse**](CreateSubnetResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSubnet

> DeleteSubnet(ctx, subnetId).Execute()

Delete subnet



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
	subnetId := "subnetId_example" // string | 삭제할 서브넷 ID - [List subnets](/openapi/networking/vpc/list-subnets)에서 확인

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SubnetAPI.DeleteSubnet(context.Background(), subnetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubnetAPI.DeleteSubnet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subnetId** | **string** | 삭제할 서브넷 ID - [List subnets](/openapi/networking/vpc/list-subnets)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubnetRequest struct via the builder pattern


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


## GetSubnet

> GetSubnetResponse GetSubnet(ctx, subnetId).Execute()

Get subnet



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
	subnetId := "subnetId_example" // string | 조회할 서브넷 ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubnetAPI.GetSubnet(context.Background(), subnetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubnetAPI.GetSubnet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubnet`: GetSubnetResponse
	fmt.Fprintf(os.Stdout, "Response from `SubnetAPI.GetSubnet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subnetId** | **string** | 조회할 서브넷 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubnetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetSubnetResponse**](GetSubnetResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubnetSharedProjects

> ListSubnetSharedProjectsResponse ListSubnetSharedProjects(ctx, subnetId).Execute()

List subnet shared projects



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
	subnetId := "subnetId_example" // string | 공유받은 프로젝트를 조회할 대상 서브넷 ID - [List subnets](/openapi/networking/vpc/list-subnets)에서 확인

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubnetAPI.ListSubnetSharedProjects(context.Background(), subnetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubnetAPI.ListSubnetSharedProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSubnetSharedProjects`: ListSubnetSharedProjectsResponse
	fmt.Fprintf(os.Stdout, "Response from `SubnetAPI.ListSubnetSharedProjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subnetId** | **string** | 공유받은 프로젝트를 조회할 대상 서브넷 ID - [List subnets](/openapi/networking/vpc/list-subnets)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSubnetSharedProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListSubnetSharedProjectsResponse**](ListSubnetSharedProjectsResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubnets

> ListSubnetsResponse ListSubnets(ctx).Id(id).Name(name).AvailabilityZone(availabilityZone).ProvisioningStatus(provisioningStatus).CidrBlock(cidrBlock).VpcId(vpcId).VpcName(vpcName).RouteTableId(routeTableId).RouteTableName(routeTableName).IsShared(isShared).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Offset(offset).Limit(limit).Execute()

List subnets



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
	id := "id_example" // string | 조회할 서브넷 ID (optional)
	name := "name_example" // string | 조회할 서브넷 이름 (optional)
	availabilityZone := openapiclient.AvailabilityZone("kr-central-2-a") // AvailabilityZone | 서브넷이 위치한 가용 영역 (optional)
	provisioningStatus := openapiclient.ProvisioningStatus("ACTIVE") // ProvisioningStatus | 프로비저닝 상태 (optional)
	cidrBlock := "cidrBlock_example" // string | 서브넷의 IPv4 CIDR 블록 (예: `10.0.1.0/24`) (optional)
	vpcId := "vpcId_example" // string | 연결된 VPC ID (optional)
	vpcName := "vpcName_example" // string | 연결된 VPC 이름 (optional)
	routeTableId := "routeTableId_example" // string | 연결된 라우팅 테이블 ID (optional)
	routeTableName := "routeTableName_example" // string | 연결된 라우팅 테이블 이름 (optional)
	isShared := true // bool | 공유 여부 (optional)
	createdAt := "createdAt_example" // string | 리소스가 생성된 시간 - ISO 8601 형식 - UTC 기준 (optional)
	updatedAt := "updatedAt_example" // string | 리소스가 마지막으로 수정된 시간 - ISO 8601 형식 - UTC 기준 (optional)
	sortKeys := "sortKeys_example" // string | 정렬할 필드를 콤마(,)로 구분 (optional)
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`) (optional)
	offset := int32(56) // int32 | 조회 시작 위치 (optional)
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubnetAPI.ListSubnets(context.Background()).Id(id).Name(name).AvailabilityZone(availabilityZone).ProvisioningStatus(provisioningStatus).CidrBlock(cidrBlock).VpcId(vpcId).VpcName(vpcName).RouteTableId(routeTableId).RouteTableName(routeTableName).IsShared(isShared).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubnetAPI.ListSubnets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSubnets`: ListSubnetsResponse
	fmt.Fprintf(os.Stdout, "Response from `SubnetAPI.ListSubnets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSubnetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | 조회할 서브넷 ID | 
 **name** | **string** | 조회할 서브넷 이름 | 
 **availabilityZone** | [**AvailabilityZone**](AvailabilityZone.md) | 서브넷이 위치한 가용 영역 | 
 **provisioningStatus** | [**ProvisioningStatus**](ProvisioningStatus.md) | 프로비저닝 상태 | 
 **cidrBlock** | **string** | 서브넷의 IPv4 CIDR 블록 (예: &#x60;10.0.1.0/24&#x60;) | 
 **vpcId** | **string** | 연결된 VPC ID | 
 **vpcName** | **string** | 연결된 VPC 이름 | 
 **routeTableId** | **string** | 연결된 라우팅 테이블 ID | 
 **routeTableName** | **string** | 연결된 라우팅 테이블 이름 | 
 **isShared** | **bool** | 공유 여부 | 
 **createdAt** | **string** | 리소스가 생성된 시간 - ISO 8601 형식 - UTC 기준 | 
 **updatedAt** | **string** | 리소스가 마지막으로 수정된 시간 - ISO 8601 형식 - UTC 기준 | 
 **sortKeys** | **string** | 정렬할 필드를 콤마(,)로 구분 | 
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;) | 
 **offset** | **int32** | 조회 시작 위치 | 
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | 

### Return type

[**ListSubnetsResponse**](ListSubnetsResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShareSubnet

> ShareSubnet(ctx, subnetId, projectId).Execute()

Share subnet



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
	subnetId := "subnetId_example" // string | 공유할 서브넷 ID
	projectId := "projectId_example" // string | 서브넷을 공유받을 대상 프로젝트 ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SubnetAPI.ShareSubnet(context.Background(), subnetId, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubnetAPI.ShareSubnet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subnetId** | **string** | 공유할 서브넷 ID | 
**projectId** | **string** | 서브넷을 공유받을 대상 프로젝트 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiShareSubnetRequest struct via the builder pattern


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


## UnshareSubnet

> UnshareSubnet(ctx, subnetId, projectId).Execute()

Unshare subnet



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
	subnetId := "subnetId_example" // string | 공유 해제할 서브넷 ID
	projectId := "projectId_example" // string | 공유를 해제할 대상 프로젝트 ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SubnetAPI.UnshareSubnet(context.Background(), subnetId, projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubnetAPI.UnshareSubnet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subnetId** | **string** | 공유 해제할 서브넷 ID | 
**projectId** | **string** | 공유를 해제할 대상 프로젝트 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnshareSubnetRequest struct via the builder pattern


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


## UpdateSubnet

> UpdateSubnetResponse UpdateSubnet(ctx, subnetId).UpdateSubnetRequest(updateSubnetRequest).Execute()

Update subnet



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
	subnetId := "subnetId_example" // string | 이름을 수정할 서브넷의 ID - [List subnets](/openapi/networking/vpc/list-subnets)에서 확인
	updateSubnetRequest := *openapiclient.NewUpdateSubnetRequest(*openapiclient.NewUpdateSubnet()) // UpdateSubnetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubnetAPI.UpdateSubnet(context.Background(), subnetId).UpdateSubnetRequest(updateSubnetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubnetAPI.UpdateSubnet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSubnet`: UpdateSubnetResponse
	fmt.Fprintf(os.Stdout, "Response from `SubnetAPI.UpdateSubnet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subnetId** | **string** | 이름을 수정할 서브넷의 ID - [List subnets](/openapi/networking/vpc/list-subnets)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSubnetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSubnetRequest** | [**UpdateSubnetRequest**](UpdateSubnetRequest.md) |  | 

### Return type

[**UpdateSubnetResponse**](UpdateSubnetResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

