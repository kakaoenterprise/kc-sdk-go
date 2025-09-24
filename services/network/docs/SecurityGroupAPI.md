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

> BnsNetworkV1ApiCreateSecurityGroupModelResponseSecurityGroupModel CreateSecurityGroup(ctx).XAuthToken(xAuthToken).BodyCreateSecurityGroup(bodyCreateSecurityGroup).Execute()

Create security group



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
	bodyCreateSecurityGroup := *openapiclient.NewBodyCreateSecurityGroup(*openapiclient.NewCreateSecurityGroupModel("Name_example")) // BodyCreateSecurityGroup | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityGroupAPI.CreateSecurityGroup(context.Background()).XAuthToken(xAuthToken).BodyCreateSecurityGroup(bodyCreateSecurityGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupAPI.CreateSecurityGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSecurityGroup`: BnsNetworkV1ApiCreateSecurityGroupModelResponseSecurityGroupModel
	fmt.Fprintf(os.Stdout, "Response from `SecurityGroupAPI.CreateSecurityGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecurityGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyCreateSecurityGroup** | [**BodyCreateSecurityGroup**](BodyCreateSecurityGroup.md) |  | 

### Return type

[**BnsNetworkV1ApiCreateSecurityGroupModelResponseSecurityGroupModel**](BnsNetworkV1ApiCreateSecurityGroupModelResponseSecurityGroupModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSecurityGroupRule

> ResponseSecurityGroupRuleModel CreateSecurityGroupRule(ctx, securityGroupId).XAuthToken(xAuthToken).BodyCreateSecurityGroupRule(bodyCreateSecurityGroupRule).Execute()

Create security group rule



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
	securityGroupId := "securityGroupId_example" // string | 보안 그룹의 고유 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bodyCreateSecurityGroupRule := *openapiclient.NewBodyCreateSecurityGroupRule(*openapiclient.NewCreateSecurityGroupRuleModel(openapiclient.SecurityGroupRuleDirection("ingress"), openapiclient.SecurityGroupRuleProtocol("TCP"))) // BodyCreateSecurityGroupRule | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityGroupAPI.CreateSecurityGroupRule(context.Background(), securityGroupId).XAuthToken(xAuthToken).BodyCreateSecurityGroupRule(bodyCreateSecurityGroupRule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupAPI.CreateSecurityGroupRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSecurityGroupRule`: ResponseSecurityGroupRuleModel
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
 **bodyCreateSecurityGroupRule** | [**BodyCreateSecurityGroupRule**](BodyCreateSecurityGroupRule.md) |  | 

### Return type

[**ResponseSecurityGroupRuleModel**](ResponseSecurityGroupRuleModel.md)

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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	securityGroupId := "securityGroupId_example" // string | 보안 그룹의 고유 ID
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
**securityGroupId** | **string** | 보안 그룹의 고유 ID | 

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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	securityGroupId := "securityGroupId_example" // string | 보안 그룹의 고유 ID
	securityGroupRuleId := "securityGroupRuleId_example" // string | 삭제할 보안 그룹 규칙 ID
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
**securityGroupId** | **string** | 보안 그룹의 고유 ID | 
**securityGroupRuleId** | **string** | 삭제할 보안 그룹 규칙 ID | 

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

> BnsNetworkV1ApiGetSecurityGroupModelResponseSecurityGroupModel GetSecurityGroup(ctx, securityGroupId).XAuthToken(xAuthToken).Execute()

Get security group



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
	securityGroupId := "securityGroupId_example" // string | 보안 그룹의 고유 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityGroupAPI.GetSecurityGroup(context.Background(), securityGroupId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupAPI.GetSecurityGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecurityGroup`: BnsNetworkV1ApiGetSecurityGroupModelResponseSecurityGroupModel
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

[**BnsNetworkV1ApiGetSecurityGroupModelResponseSecurityGroupModel**](BnsNetworkV1ApiGetSecurityGroupModelResponseSecurityGroupModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSecurityGroups

> SecurityGroupListModel ListSecurityGroups(ctx).XAuthToken(xAuthToken).Id(id).Name(name).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Limit(limit).Offset(offset).Execute()

List security groups



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
	id := "id_example" // string | 보안 그룹 ID (optional)
	name := "name_example" // string | 보안 그룹 이름 (optional)
	createdAt := "createdAt_example" // string | 리소스가 생성된 시간 <br/> - ISO_8601 형식  <br/> - UTC 기준 (optional)
	updatedAt := "updatedAt_example" // string | 리소스가 마지막으로 수정된 시간 <br/> - ISO_8601 형식  <br/> - UTC 기준 (optional)
	sortKeys := "sortKeys_example" // string | 정렬할 필드를 콤마(,)로 구분   (optional) (default to "created_at")
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`)   (optional) (default to "desc")
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional) (default to 20)
	offset := int32(56) // int32 | 조회 시작 위치 (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityGroupAPI.ListSecurityGroups(context.Background()).XAuthToken(xAuthToken).Id(id).Name(name).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupAPI.ListSecurityGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSecurityGroups`: SecurityGroupListModel
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
 **createdAt** | **string** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
 **updatedAt** | **string** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
 **sortKeys** | **string** | 정렬할 필드를 콤마(,)로 구분   | [default to &quot;created_at&quot;]
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;)   | [default to &quot;desc&quot;]
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | [default to 20]
 **offset** | **int32** | 조회 시작 위치 | [default to 0]

### Return type

[**SecurityGroupListModel**](SecurityGroupListModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSecurityGroup

> BnsNetworkV1ApiUpdateSecurityGroupModelResponseSecurityGroupModel UpdateSecurityGroup(ctx, securityGroupId).XAuthToken(xAuthToken).BodyUpdateSecurityGroup(bodyUpdateSecurityGroup).Execute()

Update security group



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
	securityGroupId := "securityGroupId_example" // string | 보안 그룹의 고유 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bodyUpdateSecurityGroup := *openapiclient.NewBodyUpdateSecurityGroup(*openapiclient.NewEditSecurityGroupModel()) // BodyUpdateSecurityGroup | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityGroupAPI.UpdateSecurityGroup(context.Background(), securityGroupId).XAuthToken(xAuthToken).BodyUpdateSecurityGroup(bodyUpdateSecurityGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityGroupAPI.UpdateSecurityGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSecurityGroup`: BnsNetworkV1ApiUpdateSecurityGroupModelResponseSecurityGroupModel
	fmt.Fprintf(os.Stdout, "Response from `SecurityGroupAPI.UpdateSecurityGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityGroupId** | **string** | 보안 그룹의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSecurityGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyUpdateSecurityGroup** | [**BodyUpdateSecurityGroup**](BodyUpdateSecurityGroup.md) |  | 

### Return type

[**BnsNetworkV1ApiUpdateSecurityGroupModelResponseSecurityGroupModel**](BnsNetworkV1ApiUpdateSecurityGroupModelResponseSecurityGroupModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

