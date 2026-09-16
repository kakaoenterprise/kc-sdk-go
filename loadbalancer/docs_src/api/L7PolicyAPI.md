# \L7PolicyAPI

All URIs are relative to *https://load-balancer.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateL7Policy**](L7PolicyAPI.md#CreateL7Policy) | **Post** /api/v1/load-balancers/l7policies | Create L7 policy
[**DeleteL7Policy**](L7PolicyAPI.md#DeleteL7Policy) | **Delete** /api/v1/load-balancers/l7policies/{l7policy_id} | Delete L7 policy
[**GetL7Policy**](L7PolicyAPI.md#GetL7Policy) | **Get** /api/v1/load-balancers/l7policies/{l7policy_id} | Get L7 policy
[**ListL7Policies**](L7PolicyAPI.md#ListL7Policies) | **Get** /api/v1/load-balancers/{load_balancer_id}/listeners/{listener_id}/l7policies | List L7 policies
[**UpdateL7Policy**](L7PolicyAPI.md#UpdateL7Policy) | **Put** /api/v1/load-balancers/l7policies/{l7policy_id} | Update L7 policy



## CreateL7Policy

> CreateL7PolicyResponse CreateL7Policy(ctx).CreateL7PolicyRequest(createL7PolicyRequest).Execute()

Create L7 policy



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
	createL7PolicyRequest := *openapiclient.NewCreateL7PolicyRequest(*openapiclient.NewCreateL7Policy(openapiclient.L7PolicyAction("REDIRECT_PREFIX"), "ListenerId_example")) // CreateL7PolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.L7PolicyAPI.CreateL7Policy(context.Background()).CreateL7PolicyRequest(createL7PolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `L7PolicyAPI.CreateL7Policy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateL7Policy`: CreateL7PolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `L7PolicyAPI.CreateL7Policy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateL7PolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createL7PolicyRequest** | [**CreateL7PolicyRequest**](CreateL7PolicyRequest.md) |  | 

### Return type

[**CreateL7PolicyResponse**](CreateL7PolicyResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteL7Policy

> DeleteL7Policy(ctx, l7policyId).Execute()

Delete L7 policy



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
	l7policyId := "l7policyId_example" // string | 삭제할 L7 정책의 ID - [List L7 policies](/openapi/networking/lb/list-l-7-policies)에서 확인

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.L7PolicyAPI.DeleteL7Policy(context.Background(), l7policyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `L7PolicyAPI.DeleteL7Policy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**l7policyId** | **string** | 삭제할 L7 정책의 ID - [List L7 policies](/openapi/networking/lb/list-l-7-policies)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteL7PolicyRequest struct via the builder pattern


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


## GetL7Policy

> GetL7PolicyResponse GetL7Policy(ctx, l7policyId).Execute()

Get L7 policy



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
	l7policyId := "l7policyId_example" // string | 조회할 L7 정책 ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.L7PolicyAPI.GetL7Policy(context.Background(), l7policyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `L7PolicyAPI.GetL7Policy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetL7Policy`: GetL7PolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `L7PolicyAPI.GetL7Policy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**l7policyId** | **string** | 조회할 L7 정책 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetL7PolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetL7PolicyResponse**](GetL7PolicyResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListL7Policies

> ListL7PoliciesResponse ListL7Policies(ctx, loadBalancerId, listenerId).Id(id).Position(position).Action(action).ProvisioningStatus(provisioningStatus).OperatingStatus(operatingStatus).Name(name).SortKeys(sortKeys).SortDirs(sortDirs).Offset(offset).Limit(limit).Execute()

List L7 policies



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
	loadBalancerId := "loadBalancerId_example" // string | 조회할 로드 밸런서 ID - [List load balancers](/openapi/networking/lb/list-load-balancers)에서 확인
	listenerId := "listenerId_example" // string | 조회할 리스너 ID - [List listeners](/openapi/networking/lb/list-listeners)에서 확인
	id := "id_example" // string | L7 정책 ID (optional)
	position := int32(56) // int32 | 정책 적용 우선순위 (숫자가 작을수록 우선순위 높음) (optional)
	action := openapiclient.L7PolicyAction("REDIRECT_PREFIX") // L7PolicyAction | 정책 동작 유형 (optional)
	provisioningStatus := openapiclient.ProvisioningStatus("ACTIVE") // ProvisioningStatus | 프로비저닝 상태 (optional)
	operatingStatus := openapiclient.LoadBalancerOperatingStatus("ONLINE") // LoadBalancerOperatingStatus | 운영 상태 (optional)
	name := "name_example" // string | L7 정책 이름 (optional)
	sortKeys := "sortKeys_example" // string | 정렬할 필드를 콤마(,)로 구분 (optional)
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`) (optional)
	offset := int32(56) // int32 | 조회 시작 위치 (optional)
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.L7PolicyAPI.ListL7Policies(context.Background(), loadBalancerId, listenerId).Id(id).Position(position).Action(action).ProvisioningStatus(provisioningStatus).OperatingStatus(operatingStatus).Name(name).SortKeys(sortKeys).SortDirs(sortDirs).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `L7PolicyAPI.ListL7Policies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListL7Policies`: ListL7PoliciesResponse
	fmt.Fprintf(os.Stdout, "Response from `L7PolicyAPI.ListL7Policies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | 조회할 로드 밸런서 ID - [List load balancers](/openapi/networking/lb/list-load-balancers)에서 확인 | 
**listenerId** | **string** | 조회할 리스너 ID - [List listeners](/openapi/networking/lb/list-listeners)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiListL7PoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **id** | **string** | L7 정책 ID | 
 **position** | **int32** | 정책 적용 우선순위 (숫자가 작을수록 우선순위 높음) | 
 **action** | [**L7PolicyAction**](L7PolicyAction.md) | 정책 동작 유형 | 
 **provisioningStatus** | [**ProvisioningStatus**](ProvisioningStatus.md) | 프로비저닝 상태 | 
 **operatingStatus** | [**LoadBalancerOperatingStatus**](LoadBalancerOperatingStatus.md) | 운영 상태 | 
 **name** | **string** | L7 정책 이름 | 
 **sortKeys** | **string** | 정렬할 필드를 콤마(,)로 구분 | 
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;) | 
 **offset** | **int32** | 조회 시작 위치 | 
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | 

### Return type

[**ListL7PoliciesResponse**](ListL7PoliciesResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateL7Policy

> UpdateL7PolicyResponse UpdateL7Policy(ctx, l7policyId).UpdateL7PolicyRequest(updateL7PolicyRequest).Execute()

Update L7 policy



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
	l7policyId := "l7policyId_example" // string | 수정할 대상 L7 정책의 ID - [List L7 policies](/openapi/networking/lb/list-l-7-policies)에서 확인
	updateL7PolicyRequest := *openapiclient.NewUpdateL7PolicyRequest(*openapiclient.NewUpdateL7Policy(openapiclient.L7PolicyAction("REDIRECT_PREFIX"))) // UpdateL7PolicyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.L7PolicyAPI.UpdateL7Policy(context.Background(), l7policyId).UpdateL7PolicyRequest(updateL7PolicyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `L7PolicyAPI.UpdateL7Policy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateL7Policy`: UpdateL7PolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `L7PolicyAPI.UpdateL7Policy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**l7policyId** | **string** | 수정할 대상 L7 정책의 ID - [List L7 policies](/openapi/networking/lb/list-l-7-policies)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateL7PolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateL7PolicyRequest** | [**UpdateL7PolicyRequest**](UpdateL7PolicyRequest.md) |  | 

### Return type

[**UpdateL7PolicyResponse**](UpdateL7PolicyResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

