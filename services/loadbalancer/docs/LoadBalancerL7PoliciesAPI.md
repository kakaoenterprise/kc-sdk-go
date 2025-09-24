# \LoadBalancerL7PoliciesAPI

All URIs are relative to *https://load-balancer.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddL7PolicyRule**](LoadBalancerL7PoliciesAPI.md#AddL7PolicyRule) | **Post** /api/v1/load-balancers/l7policies/{l7_policy_id}/rules | Add L7 policy rule
[**CreateL7Policy**](LoadBalancerL7PoliciesAPI.md#CreateL7Policy) | **Post** /api/v1/load-balancers/l7policies | Create L7 policy
[**DeleteL7Policy**](LoadBalancerL7PoliciesAPI.md#DeleteL7Policy) | **Delete** /api/v1/load-balancers/l7policies/{l7_policy_id} | Delete L7 policy
[**DeleteL7PolicyRule**](LoadBalancerL7PoliciesAPI.md#DeleteL7PolicyRule) | **Delete** /api/v1/load-balancers/l7policies/{l7_policy_id}/rules/{l7_rule_id} | Delete L7 policy rule
[**GetL7Policy**](LoadBalancerL7PoliciesAPI.md#GetL7Policy) | **Get** /api/v1/load-balancers/l7policies/{l7_policy_id} | Get L7 policy
[**GetL7PolicyRule**](LoadBalancerL7PoliciesAPI.md#GetL7PolicyRule) | **Get** /api/v1/load-balancers/l7policies/{l7_policy_id}/rules/{l7_rule_id} | Get L7 policy rule
[**ListL7Policies**](LoadBalancerL7PoliciesAPI.md#ListL7Policies) | **Get** /api/v1/load-balancers/{load_balancer_id}/listeners/{listener_id}/l7policies | List L7 policies
[**ListL7PolicyRules**](LoadBalancerL7PoliciesAPI.md#ListL7PolicyRules) | **Get** /api/v1/load-balancers/l7policies/{l7_policy_id}/rules | List L7 policy rules
[**UpdateL7Policy**](LoadBalancerL7PoliciesAPI.md#UpdateL7Policy) | **Put** /api/v1/load-balancers/l7policies/{l7_policy_id} | Update L7 policy
[**UpdateL7PolicyRule**](LoadBalancerL7PoliciesAPI.md#UpdateL7PolicyRule) | **Put** /api/v1/load-balancers/l7policies/{l7_policy_id}/rules/{l7_rule_id} | Update L7 policy rule



## AddL7PolicyRule

> BnsLoadBalancerV1ApiAddL7PolicyRuleModelResponseL7PolicyRuleModel AddL7PolicyRule(ctx, l7PolicyId).XAuthToken(xAuthToken).BodyAddL7PolicyRule(bodyAddL7PolicyRule).Execute()

Add L7 policy rule



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
	l7PolicyId := "l7PolicyId_example" // string | 규칙을 추가할 대상 L7 정책 ID 
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bodyAddL7PolicyRule := *openapiclient.NewBodyAddL7PolicyRule(*openapiclient.NewCreateL7PolicyRuleModel(openapiclient.L7RuleCompareType("CONTAINS"), openapiclient.L7RuleType("COOKIE"), "Value_example")) // BodyAddL7PolicyRule | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerL7PoliciesAPI.AddL7PolicyRule(context.Background(), l7PolicyId).XAuthToken(xAuthToken).BodyAddL7PolicyRule(bodyAddL7PolicyRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerL7PoliciesAPI.AddL7PolicyRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddL7PolicyRule`: BnsLoadBalancerV1ApiAddL7PolicyRuleModelResponseL7PolicyRuleModel
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerL7PoliciesAPI.AddL7PolicyRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**l7PolicyId** | **string** | 규칙을 추가할 대상 L7 정책 ID  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddL7PolicyRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyAddL7PolicyRule** | [**BodyAddL7PolicyRule**](BodyAddL7PolicyRule.md) |  | 

### Return type

[**BnsLoadBalancerV1ApiAddL7PolicyRuleModelResponseL7PolicyRuleModel**](BnsLoadBalancerV1ApiAddL7PolicyRuleModelResponseL7PolicyRuleModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateL7Policy

> BnsLoadBalancerV1ApiCreateL7PolicyModelResponseL7PolicyModel CreateL7Policy(ctx).XAuthToken(xAuthToken).BodyCreateL7Policy(bodyCreateL7Policy).Execute()

Create L7 policy



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
	bodyCreateL7Policy := *openapiclient.NewBodyCreateL7Policy(*openapiclient.NewCreateL7PolicyModel(openapiclient.L7PolicyAction("REDIRECT_PREFIX"), "ListenerId_example")) // BodyCreateL7Policy | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerL7PoliciesAPI.CreateL7Policy(context.Background()).XAuthToken(xAuthToken).BodyCreateL7Policy(bodyCreateL7Policy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerL7PoliciesAPI.CreateL7Policy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateL7Policy`: BnsLoadBalancerV1ApiCreateL7PolicyModelResponseL7PolicyModel
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerL7PoliciesAPI.CreateL7Policy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateL7PolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyCreateL7Policy** | [**BodyCreateL7Policy**](BodyCreateL7Policy.md) |  | 

### Return type

[**BnsLoadBalancerV1ApiCreateL7PolicyModelResponseL7PolicyModel**](BnsLoadBalancerV1ApiCreateL7PolicyModelResponseL7PolicyModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteL7Policy

> DeleteL7Policy(ctx, l7PolicyId).XAuthToken(xAuthToken).Execute()

Delete L7 policy



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
	l7PolicyId := "l7PolicyId_example" // string | 삭제할 L7 정책의 ID 
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LoadBalancerL7PoliciesAPI.DeleteL7Policy(context.Background(), l7PolicyId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerL7PoliciesAPI.DeleteL7Policy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**l7PolicyId** | **string** | 삭제할 L7 정책의 ID  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteL7PolicyRequest struct via the builder pattern


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


## DeleteL7PolicyRule

> DeleteL7PolicyRule(ctx, l7PolicyId, l7RuleId).XAuthToken(xAuthToken).Execute()

Delete L7 policy rule



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
	l7PolicyId := "l7PolicyId_example" // string | 삭제할 L7 정책의 ID 
	l7RuleId := "l7RuleId_example" // string | 삭제할 L7 규칙의 ID 
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LoadBalancerL7PoliciesAPI.DeleteL7PolicyRule(context.Background(), l7PolicyId, l7RuleId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerL7PoliciesAPI.DeleteL7PolicyRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**l7PolicyId** | **string** | 삭제할 L7 정책의 ID  | 
**l7RuleId** | **string** | 삭제할 L7 규칙의 ID  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteL7PolicyRuleRequest struct via the builder pattern


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


## GetL7Policy

> BnsLoadBalancerV1ApiGetL7PolicyModelResponseL7PolicyModel GetL7Policy(ctx, l7PolicyId).XAuthToken(xAuthToken).Execute()

Get L7 policy



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
	l7PolicyId := "l7PolicyId_example" // string | 조회할 L7 정책 ID 
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerL7PoliciesAPI.GetL7Policy(context.Background(), l7PolicyId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerL7PoliciesAPI.GetL7Policy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetL7Policy`: BnsLoadBalancerV1ApiGetL7PolicyModelResponseL7PolicyModel
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerL7PoliciesAPI.GetL7Policy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**l7PolicyId** | **string** | 조회할 L7 정책 ID  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetL7PolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**BnsLoadBalancerV1ApiGetL7PolicyModelResponseL7PolicyModel**](BnsLoadBalancerV1ApiGetL7PolicyModelResponseL7PolicyModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetL7PolicyRule

> Responsel7PolicyRuleModel GetL7PolicyRule(ctx, l7PolicyId, l7RuleId).XAuthToken(xAuthToken).Execute()

Get L7 policy rule



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
	l7PolicyId := "l7PolicyId_example" // string | 조회할 L7 정책 ID 
	l7RuleId := "l7RuleId_example" // string | 조회할 L7 규칙 ID 
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerL7PoliciesAPI.GetL7PolicyRule(context.Background(), l7PolicyId, l7RuleId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerL7PoliciesAPI.GetL7PolicyRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetL7PolicyRule`: Responsel7PolicyRuleModel
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerL7PoliciesAPI.GetL7PolicyRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**l7PolicyId** | **string** | 조회할 L7 정책 ID  | 
**l7RuleId** | **string** | 조회할 L7 규칙 ID  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetL7PolicyRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**Responsel7PolicyRuleModel**](Responsel7PolicyRuleModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListL7Policies

> L7PolicyListModel ListL7Policies(ctx, loadBalancerId, listenerId).XAuthToken(xAuthToken).Position(position).Action(action).ProvisioningStatus(provisioningStatus).OperatingStatus(operatingStatus).Name(name).SortKeys(sortKeys).SortDirs(sortDirs).Limit(limit).Offset(offset).Execute()

List L7 policies



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
	loadBalancerId := "loadBalancerId_example" // string | 조회할 로드 밸런서 ID 
	listenerId := "listenerId_example" // string | 조회할 리스너 ID 
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	position := int32(56) // int32 | 정책 적용 우선순위 (숫자가 작을수록 우선순위 높음)  (optional)
	action := openapiclient.L7PolicyAction("REDIRECT_PREFIX") // L7PolicyAction | 정책 동작 유형 <br/> - `REDIRECT_PREFIX`: Prefix로 리디렉션 <br/> - `REDIRECT_TO_POOL`: 대상 풀로 리디렉션 <br/> - `REDIRECT_TO_URL`: URL로 리디렉션 (optional)
	provisioningStatus := openapiclient.ProvisioningStatus("ACTIVE") // ProvisioningStatus | 프로비저닝 상태 <br/> - `ACTIVE`: 활성 <br/> - `DELETED`: 삭제됨 <br/> - `ERROR`: 오류 <br/> - `PENDING_CREATE`: 생성 대기 중 <br/> - `PENDING_UPDATE`: 업데이트 대기 중 <br/> - `PENDING_DELETE`: 삭제 대기 중 (optional)
	operatingStatus := openapiclient.LoadBalancerOperatingStatus("ONLINE") // LoadBalancerOperatingStatus | 운영 상태 <br/> - `ONLINE`: 온라인 <br/> - `DRAINING`: 연결 종료 중 <br/> - `OFFLINE`: 오프라인 <br/> - `DEGRADED`: 성능 저하 <br/> - `ERROR`: 오류 <br/> - `NO_MONITOR`: 모니터링 없음 (optional)
	name := "name_example" // string | L7 정책 이름  (optional)
	sortKeys := "sortKeys_example" // string | 정렬할 필드를 콤마(,)로 구분   (optional) (default to "created_at")
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`)   (optional) (default to "desc")
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional) (default to 20)
	offset := int32(56) // int32 | 조회 시작 위치 (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerL7PoliciesAPI.ListL7Policies(context.Background(), loadBalancerId, listenerId).XAuthToken(xAuthToken).Position(position).Action(action).ProvisioningStatus(provisioningStatus).OperatingStatus(operatingStatus).Name(name).SortKeys(sortKeys).SortDirs(sortDirs).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerL7PoliciesAPI.ListL7Policies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListL7Policies`: L7PolicyListModel
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerL7PoliciesAPI.ListL7Policies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | 조회할 로드 밸런서 ID  | 
**listenerId** | **string** | 조회할 리스너 ID  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListL7PoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **position** | **int32** | 정책 적용 우선순위 (숫자가 작을수록 우선순위 높음)  | 
 **action** | [**L7PolicyAction**](L7PolicyAction.md) | 정책 동작 유형 &lt;br/&gt; - &#x60;REDIRECT_PREFIX&#x60;: Prefix로 리디렉션 &lt;br/&gt; - &#x60;REDIRECT_TO_POOL&#x60;: 대상 풀로 리디렉션 &lt;br/&gt; - &#x60;REDIRECT_TO_URL&#x60;: URL로 리디렉션 | 
 **provisioningStatus** | [**ProvisioningStatus**](ProvisioningStatus.md) | 프로비저닝 상태 &lt;br/&gt; - &#x60;ACTIVE&#x60;: 활성 &lt;br/&gt; - &#x60;DELETED&#x60;: 삭제됨 &lt;br/&gt; - &#x60;ERROR&#x60;: 오류 &lt;br/&gt; - &#x60;PENDING_CREATE&#x60;: 생성 대기 중 &lt;br/&gt; - &#x60;PENDING_UPDATE&#x60;: 업데이트 대기 중 &lt;br/&gt; - &#x60;PENDING_DELETE&#x60;: 삭제 대기 중 | 
 **operatingStatus** | [**LoadBalancerOperatingStatus**](LoadBalancerOperatingStatus.md) | 운영 상태 &lt;br/&gt; - &#x60;ONLINE&#x60;: 온라인 &lt;br/&gt; - &#x60;DRAINING&#x60;: 연결 종료 중 &lt;br/&gt; - &#x60;OFFLINE&#x60;: 오프라인 &lt;br/&gt; - &#x60;DEGRADED&#x60;: 성능 저하 &lt;br/&gt; - &#x60;ERROR&#x60;: 오류 &lt;br/&gt; - &#x60;NO_MONITOR&#x60;: 모니터링 없음 | 
 **name** | **string** | L7 정책 이름  | 
 **sortKeys** | **string** | 정렬할 필드를 콤마(,)로 구분   | [default to &quot;created_at&quot;]
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;)   | [default to &quot;desc&quot;]
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | [default to 20]
 **offset** | **int32** | 조회 시작 위치 | [default to 0]

### Return type

[**L7PolicyListModel**](L7PolicyListModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListL7PolicyRules

> L7PolicyRuleListModel ListL7PolicyRules(ctx, l7PolicyId).XAuthToken(xAuthToken).Execute()

List L7 policy rules



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
	l7PolicyId := "l7PolicyId_example" // string | 조회할 L7 정책 ID 
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerL7PoliciesAPI.ListL7PolicyRules(context.Background(), l7PolicyId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerL7PoliciesAPI.ListL7PolicyRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListL7PolicyRules`: L7PolicyRuleListModel
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerL7PoliciesAPI.ListL7PolicyRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**l7PolicyId** | **string** | 조회할 L7 정책 ID  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListL7PolicyRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**L7PolicyRuleListModel**](L7PolicyRuleListModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateL7Policy

> BnsLoadBalancerV1ApiUpdateL7PolicyModelResponseL7PolicyModel UpdateL7Policy(ctx, l7PolicyId).XAuthToken(xAuthToken).BodyUpdateL7Policy(bodyUpdateL7Policy).Execute()

Update L7 policy



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
	l7PolicyId := "l7PolicyId_example" // string | 수정할 대상 L7 정책의 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bodyUpdateL7Policy := *openapiclient.NewBodyUpdateL7Policy(*openapiclient.NewEditL7PolicyModel(openapiclient.L7PolicyAction("REDIRECT_PREFIX"))) // BodyUpdateL7Policy | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerL7PoliciesAPI.UpdateL7Policy(context.Background(), l7PolicyId).XAuthToken(xAuthToken).BodyUpdateL7Policy(bodyUpdateL7Policy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerL7PoliciesAPI.UpdateL7Policy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateL7Policy`: BnsLoadBalancerV1ApiUpdateL7PolicyModelResponseL7PolicyModel
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerL7PoliciesAPI.UpdateL7Policy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**l7PolicyId** | **string** | 수정할 대상 L7 정책의 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateL7PolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyUpdateL7Policy** | [**BodyUpdateL7Policy**](BodyUpdateL7Policy.md) |  | 

### Return type

[**BnsLoadBalancerV1ApiUpdateL7PolicyModelResponseL7PolicyModel**](BnsLoadBalancerV1ApiUpdateL7PolicyModelResponseL7PolicyModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateL7PolicyRule

> BnsLoadBalancerV1ApiUpdateL7PolicyRuleModelResponseL7PolicyRuleModel UpdateL7PolicyRule(ctx, l7PolicyId, l7RuleId).XAuthToken(xAuthToken).BodyUpdateL7PolicyRule(bodyUpdateL7PolicyRule).Execute()

Update L7 policy rule



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
	l7PolicyId := "l7PolicyId_example" // string | 수정할 L7 정책 ID
	l7RuleId := "l7RuleId_example" // string | 수정할 L7 규칙 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bodyUpdateL7PolicyRule := *openapiclient.NewBodyUpdateL7PolicyRule(*openapiclient.NewEditL7PolicyRuleModel(openapiclient.L7RuleCompareType("CONTAINS"), openapiclient.L7RuleType("COOKIE"), "Value_example")) // BodyUpdateL7PolicyRule | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerL7PoliciesAPI.UpdateL7PolicyRule(context.Background(), l7PolicyId, l7RuleId).XAuthToken(xAuthToken).BodyUpdateL7PolicyRule(bodyUpdateL7PolicyRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerL7PoliciesAPI.UpdateL7PolicyRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateL7PolicyRule`: BnsLoadBalancerV1ApiUpdateL7PolicyRuleModelResponseL7PolicyRuleModel
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerL7PoliciesAPI.UpdateL7PolicyRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**l7PolicyId** | **string** | 수정할 L7 정책 ID | 
**l7RuleId** | **string** | 수정할 L7 규칙 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateL7PolicyRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyUpdateL7PolicyRule** | [**BodyUpdateL7PolicyRule**](BodyUpdateL7PolicyRule.md) |  | 

### Return type

[**BnsLoadBalancerV1ApiUpdateL7PolicyRuleModelResponseL7PolicyRuleModel**](BnsLoadBalancerV1ApiUpdateL7PolicyRuleModelResponseL7PolicyRuleModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

