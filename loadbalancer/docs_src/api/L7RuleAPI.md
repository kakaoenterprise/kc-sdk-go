# \L7RuleAPI

All URIs are relative to *https://load-balancer.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddL7PolicyRule**](L7RuleAPI.md#AddL7PolicyRule) | **Post** /api/v1/load-balancers/l7policies/{l7policy_id}/rules | Add L7 policy rule
[**DeleteL7PolicyRule**](L7RuleAPI.md#DeleteL7PolicyRule) | **Delete** /api/v1/load-balancers/l7policies/{l7policy_id}/rules/{l7rule_id} | Delete L7 policy rule
[**GetL7PolicyRule**](L7RuleAPI.md#GetL7PolicyRule) | **Get** /api/v1/load-balancers/l7policies/{l7policy_id}/rules/{l7rule_id} | Get L7 policy rule
[**ListL7PolicyRules**](L7RuleAPI.md#ListL7PolicyRules) | **Get** /api/v1/load-balancers/l7policies/{l7policy_id}/rules | List L7 policy rules
[**UpdateL7PolicyRule**](L7RuleAPI.md#UpdateL7PolicyRule) | **Put** /api/v1/load-balancers/l7policies/{l7policy_id}/rules/{l7rule_id} | Update L7 policy rule



## AddL7PolicyRule

> AddL7PolicyRuleResponse AddL7PolicyRule(ctx, l7policyId).AddL7PolicyRuleRequest(addL7PolicyRuleRequest).Execute()

Add L7 policy rule



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
	l7policyId := "l7policyId_example" // string | 규칙을 추가할 대상 L7 정책 ID - [List L7 policies](/openapi/networking/lb/list-l-7-policies)에서 확인
	addL7PolicyRuleRequest := *openapiclient.NewAddL7PolicyRuleRequest(*openapiclient.NewAddL7PolicyRule(openapiclient.L7RuleCompareType("CONTAINS"), openapiclient.L7RuleType("COOKIE"), "Value_example")) // AddL7PolicyRuleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.L7RuleAPI.AddL7PolicyRule(context.Background(), l7policyId).AddL7PolicyRuleRequest(addL7PolicyRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `L7RuleAPI.AddL7PolicyRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddL7PolicyRule`: AddL7PolicyRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `L7RuleAPI.AddL7PolicyRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**l7policyId** | **string** | 규칙을 추가할 대상 L7 정책 ID - [List L7 policies](/openapi/networking/lb/list-l-7-policies)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddL7PolicyRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addL7PolicyRuleRequest** | [**AddL7PolicyRuleRequest**](AddL7PolicyRuleRequest.md) |  | 

### Return type

[**AddL7PolicyRuleResponse**](AddL7PolicyRuleResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteL7PolicyRule

> DeleteL7PolicyRule(ctx, l7policyId, l7ruleId).Execute()

Delete L7 policy rule



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
	l7ruleId := "l7ruleId_example" // string | 삭제할 L7 규칙의 ID - [List L7 policy rules](/openapi/networking/lb/list-l-7-policy-rules)에서 확인

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.L7RuleAPI.DeleteL7PolicyRule(context.Background(), l7policyId, l7ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `L7RuleAPI.DeleteL7PolicyRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**l7policyId** | **string** | 삭제할 L7 정책의 ID - [List L7 policies](/openapi/networking/lb/list-l-7-policies)에서 확인 | 
**l7ruleId** | **string** | 삭제할 L7 규칙의 ID - [List L7 policy rules](/openapi/networking/lb/list-l-7-policy-rules)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteL7PolicyRuleRequest struct via the builder pattern


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


## GetL7PolicyRule

> GetL7PolicyRuleResponse GetL7PolicyRule(ctx, l7policyId, l7ruleId).Execute()

Get L7 policy rule



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
	l7policyId := "l7policyId_example" // string | 조회할 L7 정책 ID - [List L7 policies](/openapi/networking/lb/list-l-7-policies)에서 확인
	l7ruleId := "l7ruleId_example" // string | 조회할 L7 규칙 ID - [List L7 policy rules](/openapi/networking/lb/list-l-7-policy-rules)에서 확인

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.L7RuleAPI.GetL7PolicyRule(context.Background(), l7policyId, l7ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `L7RuleAPI.GetL7PolicyRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetL7PolicyRule`: GetL7PolicyRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `L7RuleAPI.GetL7PolicyRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**l7policyId** | **string** | 조회할 L7 정책 ID - [List L7 policies](/openapi/networking/lb/list-l-7-policies)에서 확인 | 
**l7ruleId** | **string** | 조회할 L7 규칙 ID - [List L7 policy rules](/openapi/networking/lb/list-l-7-policy-rules)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetL7PolicyRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetL7PolicyRuleResponse**](GetL7PolicyRuleResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListL7PolicyRules

> ListL7PolicyRulesResponse ListL7PolicyRules(ctx, l7policyId).Execute()

List L7 policy rules



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
	l7policyId := "l7policyId_example" // string | 조회할 L7 정책 ID - [List L7 policies](/openapi/networking/lb/list-l-7-policies)에서 확인

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.L7RuleAPI.ListL7PolicyRules(context.Background(), l7policyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `L7RuleAPI.ListL7PolicyRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListL7PolicyRules`: ListL7PolicyRulesResponse
	fmt.Fprintf(os.Stdout, "Response from `L7RuleAPI.ListL7PolicyRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**l7policyId** | **string** | 조회할 L7 정책 ID - [List L7 policies](/openapi/networking/lb/list-l-7-policies)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiListL7PolicyRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ListL7PolicyRulesResponse**](ListL7PolicyRulesResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateL7PolicyRule

> UpdateL7PolicyRuleResponse UpdateL7PolicyRule(ctx, l7policyId, l7ruleId).UpdateL7PolicyRuleRequest(updateL7PolicyRuleRequest).Execute()

Update L7 policy rule



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
	l7policyId := "l7policyId_example" // string | 수정할 L7 정책 ID - [List L7 policies](/openapi/networking/lb/list-l-7-policies)에서 확인
	l7ruleId := "l7ruleId_example" // string | 수정할 L7 규칙 ID - [List L7 policy rules](/openapi/networking/lb/list-l-7-policy-rules)에서 확인
	updateL7PolicyRuleRequest := *openapiclient.NewUpdateL7PolicyRuleRequest(*openapiclient.NewUpdateL7PolicyRule(openapiclient.L7RuleCompareType("CONTAINS"), openapiclient.L7RuleType("COOKIE"), "Value_example")) // UpdateL7PolicyRuleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.L7RuleAPI.UpdateL7PolicyRule(context.Background(), l7policyId, l7ruleId).UpdateL7PolicyRuleRequest(updateL7PolicyRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `L7RuleAPI.UpdateL7PolicyRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateL7PolicyRule`: UpdateL7PolicyRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `L7RuleAPI.UpdateL7PolicyRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**l7policyId** | **string** | 수정할 L7 정책 ID - [List L7 policies](/openapi/networking/lb/list-l-7-policies)에서 확인 | 
**l7ruleId** | **string** | 수정할 L7 규칙 ID - [List L7 policy rules](/openapi/networking/lb/list-l-7-policy-rules)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateL7PolicyRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateL7PolicyRuleRequest** | [**UpdateL7PolicyRuleRequest**](UpdateL7PolicyRuleRequest.md) |  | 

### Return type

[**UpdateL7PolicyRuleResponse**](UpdateL7PolicyRuleResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

