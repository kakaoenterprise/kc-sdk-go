# \SecurityGroupAPI

All URIs are relative to *https://network.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSecurityGroup**](SecurityGroupAPI.md#CreateSecurityGroup) | **Post** /api/v1/security-groups | Create security group
[**CreateSecurityGroupRule**](SecurityGroupAPI.md#CreateSecurityGroupRule) | **Post** /api/v1/security-groups/{security_group_id}/rules | Create security group rule
[**DeleteSecurityGroup**](SecurityGroupAPI.md#DeleteSecurityGroup) | **Delete** /api/v1/security-groups/{security_group_id} | Delete security group
[**DeleteSecurityGroupRule**](SecurityGroupAPI.md#DeleteSecurityGroupRule) | **Delete** /api/v1/security-groups/{security_group_id}/rules/{security_group_rule_id} | Delete security group rule
[**GetSecurityGroup**](SecurityGroupAPI.md#GetSecurityGroup) | **Get** /api/v1/security-groups/{security_group_id} | Get security group
[**ListSecurityGroups**](SecurityGroupAPI.md#ListSecurityGroups) | **Get** /api/v1/security-groups | List security groups
[**UpdateSecurityGroup**](SecurityGroupAPI.md#UpdateSecurityGroup) | **Put** /api/v1/security-groups/{security_group_id} | Update security group



## CreateSecurityGroup

> CreateSecurityGroupResponse CreateSecurityGroup(ctx).XAuthToken(xAuthToken).CreateSecurityGroupRequest(createSecurityGroupRequest).Execute()

Create security group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/network"
)

func main() {
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	createSecurityGroupRequest := *openapiclient.NewCreateSecurityGroupRequest(*openapiclient.NewCreateSecurityGroup("Name_example")) // CreateSecurityGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityGroupAPI.CreateSecurityGroup(context.Background()).XAuthToken(xAuthToken).CreateSecurityGroupRequest(createSecurityGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupAPI.CreateSecurityGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSecurityGroup`: CreateSecurityGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityGroupAPI.CreateSecurityGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecurityGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **createSecurityGroupRequest** | [**CreateSecurityGroupRequest**](CreateSecurityGroupRequest.md) |  | 

### Return type

[**CreateSecurityGroupResponse**](CreateSecurityGroupResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSecurityGroupRule

> CreateSecurityGroupRuleResponse CreateSecurityGroupRule(ctx, securityGroupId).XAuthToken(xAuthToken).CreateSecurityGroupRuleRequest(createSecurityGroupRuleRequest).Execute()

Create security group rule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/network"
)

func main() {
	securityGroupId := "securityGroupId_example" // string | 보안 그룹의 고유 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	createSecurityGroupRuleRequest := *openapiclient.NewCreateSecurityGroupRuleRequest(*openapiclient.NewCreateSecurityGroupRule(openapiclient.RuleDirection("ingress"), openapiclient.RuleProtocol("TCP"))) // CreateSecurityGroupRuleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityGroupAPI.CreateSecurityGroupRule(context.Background(), securityGroupId).XAuthToken(xAuthToken).CreateSecurityGroupRuleRequest(createSecurityGroupRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupAPI.CreateSecurityGroupRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSecurityGroupRule`: CreateSecurityGroupRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityGroupAPI.CreateSecurityGroupRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityGroupId** | **string** | 보안 그룹의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecurityGroupRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **createSecurityGroupRuleRequest** | [**CreateSecurityGroupRuleRequest**](CreateSecurityGroupRuleRequest.md) |  | 

### Return type

[**CreateSecurityGroupRuleResponse**](CreateSecurityGroupRuleResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSecurityGroup

> DeleteSecurityGroup(ctx, securityGroupId).XAuthToken(xAuthToken).Execute()

Delete security group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/network"
)

func main() {
	securityGroupId := "securityGroupId_example" // string | 보안 그룹의 고유 ID <br/>- [List security groups](https://docs.kakaocloud.com/openapi/networking/vpc/list-security-groups)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityGroupAPI.DeleteSecurityGroup(context.Background(), securityGroupId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupAPI.DeleteSecurityGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityGroupId** | **string** | 보안 그룹의 고유 ID &lt;br/&gt;- [List security groups](https://docs.kakaocloud.com/openapi/networking/vpc/list-security-groups)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSecurityGroupRequest struct via the builder pattern


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


## DeleteSecurityGroupRule

> DeleteSecurityGroupRule(ctx, securityGroupId, securityGroupRuleId).XAuthToken(xAuthToken).Execute()

Delete security group rule



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/network"
)

func main() {
	securityGroupId := "securityGroupId_example" // string | 보안 그룹의 고유 ID <br/>- [List security groups](https://docs.kakaocloud.com/openapi/networking/vpc/list-security-groups)에서 확인
	securityGroupRuleId := "securityGroupRuleId_example" // string | 삭제할 보안 그룹 규칙 ID <br/>- [Get security group](https://docs.kakaocloud.com/openapi/networking/vpc/get-security-group)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityGroupAPI.DeleteSecurityGroupRule(context.Background(), securityGroupId, securityGroupRuleId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupAPI.DeleteSecurityGroupRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityGroupId** | **string** | 보안 그룹의 고유 ID &lt;br/&gt;- [List security groups](https://docs.kakaocloud.com/openapi/networking/vpc/list-security-groups)에서 확인 | 
**securityGroupRuleId** | **string** | 삭제할 보안 그룹 규칙 ID &lt;br/&gt;- [Get security group](https://docs.kakaocloud.com/openapi/networking/vpc/get-security-group)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSecurityGroupRuleRequest struct via the builder pattern


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


## GetSecurityGroup

> GetSecurityGroupResponse GetSecurityGroup(ctx, securityGroupId).XAuthToken(xAuthToken).Execute()

Get security group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/network"
)

func main() {
	securityGroupId := "securityGroupId_example" // string | 보안 그룹의 고유 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityGroupAPI.GetSecurityGroup(context.Background(), securityGroupId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupAPI.GetSecurityGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecurityGroup`: GetSecurityGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityGroupAPI.GetSecurityGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityGroupId** | **string** | 보안 그룹의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**GetSecurityGroupResponse**](GetSecurityGroupResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSecurityGroups

> ListSecurityGroupsResponse ListSecurityGroups(ctx).XAuthToken(xAuthToken).Id(id).Name(name).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Offset(offset).Limit(limit).Execute()

List security groups



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/network"
)

func main() {
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	id := "id_example" // string | 보안 그룹 ID (optional)
	name := "name_example" // string | 보안 그룹 이름 (optional)
	createdAt := "createdAt_example" // string | 리소스가 생성된 시간 <br/> - ISO 8601 형식 <br/> - UTC 기준 (optional)
	updatedAt := "updatedAt_example" // string | 리소스가 마지막으로 수정된 시간 <br/> - ISO 8601 형식 <br/> - UTC 기준 (optional)
	sortKeys := "sortKeys_example" // string | 정렬할 필드를 콤마(,)로 구분   (optional)
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`)  (optional)
	offset := int32(56) // int32 | 조회 시작 위치 (optional)
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityGroupAPI.ListSecurityGroups(context.Background()).XAuthToken(xAuthToken).Id(id).Name(name).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupAPI.ListSecurityGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSecurityGroups`: ListSecurityGroupsResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityGroupAPI.ListSecurityGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSecurityGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **id** | **string** | 보안 그룹 ID | 
 **name** | **string** | 보안 그룹 이름 | 
 **createdAt** | **string** | 리소스가 생성된 시간 &lt;br/&gt; - ISO 8601 형식 &lt;br/&gt; - UTC 기준 | 
 **updatedAt** | **string** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO 8601 형식 &lt;br/&gt; - UTC 기준 | 
 **sortKeys** | **string** | 정렬할 필드를 콤마(,)로 구분   | 
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;)  | 
 **offset** | **int32** | 조회 시작 위치 | 
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | 

### Return type

[**ListSecurityGroupsResponse**](ListSecurityGroupsResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSecurityGroup

> UpdateSecurityGroupResponse UpdateSecurityGroup(ctx, securityGroupId).XAuthToken(xAuthToken).UpdateSecurityGroupRequest(updateSecurityGroupRequest).Execute()

Update security group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/network"
)

func main() {
	securityGroupId := "securityGroupId_example" // string | 보안 그룹의 고유 ID <br/>- [List security groups](https://docs.kakaocloud.com/openapi/networking/vpc/list-security-groups)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	updateSecurityGroupRequest := *openapiclient.NewUpdateSecurityGroupRequest(*openapiclient.NewUpdateSecurityGroup()) // UpdateSecurityGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityGroupAPI.UpdateSecurityGroup(context.Background(), securityGroupId).XAuthToken(xAuthToken).UpdateSecurityGroupRequest(updateSecurityGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupAPI.UpdateSecurityGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSecurityGroup`: UpdateSecurityGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityGroupAPI.UpdateSecurityGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityGroupId** | **string** | 보안 그룹의 고유 ID &lt;br/&gt;- [List security groups](https://docs.kakaocloud.com/openapi/networking/vpc/list-security-groups)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSecurityGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **updateSecurityGroupRequest** | [**UpdateSecurityGroupRequest**](UpdateSecurityGroupRequest.md) |  | 

### Return type

[**UpdateSecurityGroupResponse**](UpdateSecurityGroupResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

