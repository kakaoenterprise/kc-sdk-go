# \HighAvailabilityGroupAPI

All URIs are relative to *https://load-balancer.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHaGroup**](HighAvailabilityGroupAPI.md#CreateHaGroup) | **Post** /api/v1/beyond-load-balancers | Create HA group 
[**DeleteHaGroup**](HighAvailabilityGroupAPI.md#DeleteHaGroup) | **Delete** /api/v1/beyond-load-balancers/{beyond_load_balancer_id} | Delete HA group 
[**DetachHaGroupLoadBalancer**](HighAvailabilityGroupAPI.md#DetachHaGroupLoadBalancer) | **Delete** /api/v1/beyond-load-balancers/{beyond_load_balancer_id}/subnets/{load_balancer_id} | Detach HA group load balancer 
[**GetHaGroup**](HighAvailabilityGroupAPI.md#GetHaGroup) | **Get** /api/v1/beyond-load-balancers/{beyond_load_balancer_id} | Get HA group 
[**ListHaGroups**](HighAvailabilityGroupAPI.md#ListHaGroups) | **Get** /api/v1/beyond-load-balancers | List HA groups 
[**UpdateHaGroup**](HighAvailabilityGroupAPI.md#UpdateHaGroup) | **Put** /api/v1/beyond-load-balancers/{beyond_load_balancer_id} | Update HA group 
[**UpdateHaGroupLoadBalancer**](HighAvailabilityGroupAPI.md#UpdateHaGroupLoadBalancer) | **Post** /api/v1/beyond-load-balancers/{beyond_load_balancer_id}/subnets | Update HA group load balancer 



## CreateHaGroup

> CreateHaGroupResponse CreateHaGroup(ctx).CreateHaGroupRequest(createHaGroupRequest).Execute()

Create HA group 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/loadbalancer"
)

func main() {
	createHaGroupRequest := *openapiclient.NewCreateHaGroupRequest(*openapiclient.NewCreateHaGroup("Name_example", "TypeId_example", openapiclient.BeyondLoadBalancerScheme("internet-facing"), "VpcId_example", []openapiclient.VpcSubnetRequest{*openapiclient.NewVpcSubnetRequest("LoadBalancerId_example", openapiclient.AvailabilityZone("kr-central-2-a"))})) // CreateHaGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HighAvailabilityGroupAPI.CreateHaGroup(context.Background()).CreateHaGroupRequest(createHaGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HighAvailabilityGroupAPI.CreateHaGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateHaGroup`: CreateHaGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `HighAvailabilityGroupAPI.CreateHaGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateHaGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createHaGroupRequest** | [**CreateHaGroupRequest**](CreateHaGroupRequest.md) |  | 

### Return type

[**CreateHaGroupResponse**](CreateHaGroupResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteHaGroup

> DeleteHaGroup(ctx, beyondLoadBalancerId).Execute()

Delete HA group 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/loadbalancer"
)

func main() {
	beyondLoadBalancerId := "beyondLoadBalancerId_example" // string | 삭제할 고가용성 그룹의 ID - [List HA groups](/openapi/networking/lb/list-ha-groups)에서 확인

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HighAvailabilityGroupAPI.DeleteHaGroup(context.Background(), beyondLoadBalancerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HighAvailabilityGroupAPI.DeleteHaGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**beyondLoadBalancerId** | **string** | 삭제할 고가용성 그룹의 ID - [List HA groups](/openapi/networking/lb/list-ha-groups)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHaGroupRequest struct via the builder pattern


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


## DetachHaGroupLoadBalancer

> DetachHaGroupLoadBalancer(ctx, beyondLoadBalancerId, loadBalancerId).Execute()

Detach HA group load balancer 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/loadbalancer"
)

func main() {
	beyondLoadBalancerId := "beyondLoadBalancerId_example" // string | 고가용성 그룹의 ID - [List HA groups](/openapi/networking/lb/list-ha-groups)에서 확인
	loadBalancerId := "loadBalancerId_example" // string | 연결 해제할 로드 밸런서의 ID - [List load balancers](/openapi/networking/lb/list-load-balancers)에서 확인

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HighAvailabilityGroupAPI.DetachHaGroupLoadBalancer(context.Background(), beyondLoadBalancerId, loadBalancerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HighAvailabilityGroupAPI.DetachHaGroupLoadBalancer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**beyondLoadBalancerId** | **string** | 고가용성 그룹의 ID - [List HA groups](/openapi/networking/lb/list-ha-groups)에서 확인 | 
**loadBalancerId** | **string** | 연결 해제할 로드 밸런서의 ID - [List load balancers](/openapi/networking/lb/list-load-balancers)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachHaGroupLoadBalancerRequest struct via the builder pattern


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


## GetHaGroup

> GetHaGroupResponse GetHaGroup(ctx, beyondLoadBalancerId).Execute()

Get HA group 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/loadbalancer"
)

func main() {
	beyondLoadBalancerId := "beyondLoadBalancerId_example" // string | 조회할 고가용성 그룹의 ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HighAvailabilityGroupAPI.GetHaGroup(context.Background(), beyondLoadBalancerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HighAvailabilityGroupAPI.GetHaGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHaGroup`: GetHaGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `HighAvailabilityGroupAPI.GetHaGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**beyondLoadBalancerId** | **string** | 조회할 고가용성 그룹의 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHaGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetHaGroupResponse**](GetHaGroupResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHaGroups

> ListHaGroupsResponse ListHaGroups(ctx).Id(id).Name(name).DnsName(dnsName).Scheme(scheme).ProvisioningStatus(provisioningStatus).OperatingStatus(operatingStatus).Type_(type_).VpcName(vpcName).VpcId(vpcId).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Offset(offset).Limit(limit).Execute()

List HA groups 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/loadbalancer"
)

func main() {
	id := "id_example" // string | 조회할 고가용성 그룹 ID (optional)
	name := "name_example" // string | 고가용성 그룹 이름 (optional)
	dnsName := "dnsName_example" // string | DNS 이름 (optional)
	scheme := openapiclient.BeyondLoadBalancerScheme("internet-facing") // BeyondLoadBalancerScheme | 접근 방식 (optional)
	provisioningStatus := openapiclient.ProvisioningStatus("ACTIVE") // ProvisioningStatus | 프로비저닝 상태 (optional)
	operatingStatus := openapiclient.LoadBalancerOperatingStatus("ONLINE") // LoadBalancerOperatingStatus | 운영 상태 (optional)
	type_ := openapiclient.LoadBalancerType("ALB") // LoadBalancerType | 로드 밸런서 유형 (optional)
	vpcName := "vpcName_example" // string | 연결된 VPC 이름 (optional)
	vpcId := "vpcId_example" // string | 연결된 VPC ID (optional)
	createdAt := "createdAt_example" // string | 리소스가 생성된 시간 - ISO 8601 형식 - UTC 기준 (optional)
	updatedAt := "updatedAt_example" // string | 리소스가 마지막으로 수정된 시간 - ISO 8601 형식 - UTC 기준 (optional)
	sortKeys := "sortKeys_example" // string | 정렬할 필드를 콤마(,)로 구분 (optional)
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`) (optional)
	offset := int32(56) // int32 | 조회 시작 위치 (optional)
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HighAvailabilityGroupAPI.ListHaGroups(context.Background()).Id(id).Name(name).DnsName(dnsName).Scheme(scheme).ProvisioningStatus(provisioningStatus).OperatingStatus(operatingStatus).Type_(type_).VpcName(vpcName).VpcId(vpcId).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HighAvailabilityGroupAPI.ListHaGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHaGroups`: ListHaGroupsResponse
	fmt.Fprintf(os.Stdout, "Response from `HighAvailabilityGroupAPI.ListHaGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHaGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | 조회할 고가용성 그룹 ID | 
 **name** | **string** | 고가용성 그룹 이름 | 
 **dnsName** | **string** | DNS 이름 | 
 **scheme** | [**BeyondLoadBalancerScheme**](BeyondLoadBalancerScheme.md) | 접근 방식 | 
 **provisioningStatus** | [**ProvisioningStatus**](ProvisioningStatus.md) | 프로비저닝 상태 | 
 **operatingStatus** | [**LoadBalancerOperatingStatus**](LoadBalancerOperatingStatus.md) | 운영 상태 | 
 **type_** | [**LoadBalancerType**](LoadBalancerType.md) | 로드 밸런서 유형 | 
 **vpcName** | **string** | 연결된 VPC 이름 | 
 **vpcId** | **string** | 연결된 VPC ID | 
 **createdAt** | **string** | 리소스가 생성된 시간 - ISO 8601 형식 - UTC 기준 | 
 **updatedAt** | **string** | 리소스가 마지막으로 수정된 시간 - ISO 8601 형식 - UTC 기준 | 
 **sortKeys** | **string** | 정렬할 필드를 콤마(,)로 구분 | 
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;) | 
 **offset** | **int32** | 조회 시작 위치 | 
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | 

### Return type

[**ListHaGroupsResponse**](ListHaGroupsResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHaGroup

> UpdateHaGroupResponse UpdateHaGroup(ctx, beyondLoadBalancerId).UpdateHaGroupRequest(updateHaGroupRequest).Execute()

Update HA group 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/loadbalancer"
)

func main() {
	beyondLoadBalancerId := "beyondLoadBalancerId_example" // string | 수정할 고가용성 로드 밸런서 ID - [List HA groups](/openapi/networking/lb/list-ha-groups)에서 확인
	updateHaGroupRequest := *openapiclient.NewUpdateHaGroupRequest(*openapiclient.NewUpdateHaGroup()) // UpdateHaGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HighAvailabilityGroupAPI.UpdateHaGroup(context.Background(), beyondLoadBalancerId).UpdateHaGroupRequest(updateHaGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HighAvailabilityGroupAPI.UpdateHaGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateHaGroup`: UpdateHaGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `HighAvailabilityGroupAPI.UpdateHaGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**beyondLoadBalancerId** | **string** | 수정할 고가용성 로드 밸런서 ID - [List HA groups](/openapi/networking/lb/list-ha-groups)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHaGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateHaGroupRequest** | [**UpdateHaGroupRequest**](UpdateHaGroupRequest.md) |  | 

### Return type

[**UpdateHaGroupResponse**](UpdateHaGroupResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHaGroupLoadBalancer

> UpdateHaGroupLoadBalancerResponse UpdateHaGroupLoadBalancer(ctx, beyondLoadBalancerId).UpdateHaGroupLoadBalancerRequest(updateHaGroupLoadBalancerRequest).Execute()

Update HA group load balancer 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/loadbalancer"
)

func main() {
	beyondLoadBalancerId := "beyondLoadBalancerId_example" // string | 연결 대상 고가용성 로드 밸런서 ID - [List HA groups](/openapi/networking/lb/list-ha-groups)에서 확인
	updateHaGroupLoadBalancerRequest := *openapiclient.NewUpdateHaGroupLoadBalancerRequest(*openapiclient.NewUpdateHaGroupLoadBalancer([]openapiclient.VpcSubnetRequest{*openapiclient.NewVpcSubnetRequest("LoadBalancerId_example", openapiclient.AvailabilityZone("kr-central-2-a"))})) // UpdateHaGroupLoadBalancerRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HighAvailabilityGroupAPI.UpdateHaGroupLoadBalancer(context.Background(), beyondLoadBalancerId).UpdateHaGroupLoadBalancerRequest(updateHaGroupLoadBalancerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HighAvailabilityGroupAPI.UpdateHaGroupLoadBalancer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateHaGroupLoadBalancer`: UpdateHaGroupLoadBalancerResponse
	fmt.Fprintf(os.Stdout, "Response from `HighAvailabilityGroupAPI.UpdateHaGroupLoadBalancer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**beyondLoadBalancerId** | **string** | 연결 대상 고가용성 로드 밸런서 ID - [List HA groups](/openapi/networking/lb/list-ha-groups)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHaGroupLoadBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateHaGroupLoadBalancerRequest** | [**UpdateHaGroupLoadBalancerRequest**](UpdateHaGroupLoadBalancerRequest.md) |  | 

### Return type

[**UpdateHaGroupLoadBalancerResponse**](UpdateHaGroupLoadBalancerResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

