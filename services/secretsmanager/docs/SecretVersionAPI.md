# \SecretVersionAPI

All URIs are relative to *https://secrets-manager-service.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateSecretVersion**](SecretVersionAPI.md#ActivateSecretVersion) | **Patch** /api/v1/secrets/{secret_id}/versions/{version}/activate | Activate secret version 
[**BulkCancelScheduledSecretVersionDestruction**](SecretVersionAPI.md#BulkCancelScheduledSecretVersionDestruction) | **Patch** /api/v1/secrets/{secret_id}/versions/cancel-schedule-destroy | Bulk cancel scheduled secret version destruction 
[**BulkDeactivateSecretVersions**](SecretVersionAPI.md#BulkDeactivateSecretVersions) | **Patch** /api/v1/secrets/{secret_id}/versions/deactivate | Bulk deactivate secret versions 
[**BulkDestroySecretVersions**](SecretVersionAPI.md#BulkDestroySecretVersions) | **Patch** /api/v1/secrets/{secret_id}/versions/destroy | Bulk destroy secret versions 
[**BulkScheduleSecretVersionDestruction**](SecretVersionAPI.md#BulkScheduleSecretVersionDestruction) | **Patch** /api/v1/secrets/{secret_id}/versions/schedule-destroy | Bulk schedule secret version destruction 
[**CancelScheduledSecretVersionDestruction**](SecretVersionAPI.md#CancelScheduledSecretVersionDestruction) | **Patch** /api/v1/secrets/{secret_id}/versions/{version}/cancel-schedule-destroy | Cancel scheduled secret version destruction 
[**CreateSecretVersion**](SecretVersionAPI.md#CreateSecretVersion) | **Post** /api/v1/secrets/{secret_id}/versions | Create secret version 
[**DeactivateSecretVersion**](SecretVersionAPI.md#DeactivateSecretVersion) | **Patch** /api/v1/secrets/{secret_id}/versions/{version}/deactivate | Deactivate secret version 
[**DestroySecretVersion**](SecretVersionAPI.md#DestroySecretVersion) | **Patch** /api/v1/secrets/{secret_id}/versions/{version}/destroy | Destroy secret version 
[**GetSecretDefaultVersion**](SecretVersionAPI.md#GetSecretDefaultVersion) | **Get** /api/v1/secrets/{secret_id}/versions/value | Get secret default version 
[**GetSecretVersion**](SecretVersionAPI.md#GetSecretVersion) | **Get** /api/v1/secrets/{secret_id}/versions/{version}/value | Get secret version 
[**ListSecretVersions**](SecretVersionAPI.md#ListSecretVersions) | **Get** /api/v1/secrets/{secret_id}/versions | List secret versions 
[**ScheduleSecretVersionDestruction**](SecretVersionAPI.md#ScheduleSecretVersionDestruction) | **Patch** /api/v1/secrets/{secret_id}/versions/{version}/schedule-destroy | Schedule secret version destruction 



## ActivateSecretVersion

> ActivateSecretVersionResponse ActivateSecretVersion(ctx, secretId, version).XAuthToken(xAuthToken).Execute()

Activate secret version 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/secretsmanager"
)

func main() {
	secretId := "secretId_example" // string | 시크릿의 고유 ID <br/> - [List secrets](https://docs.kakaocloud.com/openapi/security/secrets-manager/list-secrets)에서 확인
	version := "version_example" // string | 대상 시크릿 버전 <br/>- [List secret versions](https://docs.kakaocloud.com/openapi/security/secrets-manager/list-secret-versions)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretVersionAPI.ActivateSecretVersion(context.Background(), secretId, version).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretVersionAPI.ActivateSecretVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActivateSecretVersion`: ActivateSecretVersionResponse
	fmt.Fprintf(os.Stdout, "Response from `SecretVersionAPI.ActivateSecretVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretId** | **string** | 시크릿의 고유 ID &lt;br/&gt; - [List secrets](https://docs.kakaocloud.com/openapi/security/secrets-manager/list-secrets)에서 확인 | 
**version** | **string** | 대상 시크릿 버전 &lt;br/&gt;- [List secret versions](https://docs.kakaocloud.com/openapi/security/secrets-manager/list-secret-versions)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateSecretVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**ActivateSecretVersionResponse**](ActivateSecretVersionResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkCancelScheduledSecretVersionDestruction

> BulkCancelScheduledSecretVersionDestructionResponse BulkCancelScheduledSecretVersionDestruction(ctx, secretId).XAuthToken(xAuthToken).BulkCancelScheduledSecretVersionDestructionRequest(bulkCancelScheduledSecretVersionDestructionRequest).Execute()

Bulk cancel scheduled secret version destruction 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/secretsmanager"
)

func main() {
	secretId := "secretId_example" // string | 시크릿의 고유 ID <br/> - [List secrets](https://docs.kakaocloud.com/openapi/security/secrets-manager/list-secrets)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bulkCancelScheduledSecretVersionDestructionRequest := *openapiclient.NewBulkCancelScheduledSecretVersionDestructionRequest(*openapiclient.NewBulkCancelScheduledSecretVersionDestruction([]int32{int32(123)})) // BulkCancelScheduledSecretVersionDestructionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretVersionAPI.BulkCancelScheduledSecretVersionDestruction(context.Background(), secretId).XAuthToken(xAuthToken).BulkCancelScheduledSecretVersionDestructionRequest(bulkCancelScheduledSecretVersionDestructionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretVersionAPI.BulkCancelScheduledSecretVersionDestruction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkCancelScheduledSecretVersionDestruction`: BulkCancelScheduledSecretVersionDestructionResponse
	fmt.Fprintf(os.Stdout, "Response from `SecretVersionAPI.BulkCancelScheduledSecretVersionDestruction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretId** | **string** | 시크릿의 고유 ID &lt;br/&gt; - [List secrets](https://docs.kakaocloud.com/openapi/security/secrets-manager/list-secrets)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkCancelScheduledSecretVersionDestructionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bulkCancelScheduledSecretVersionDestructionRequest** | [**BulkCancelScheduledSecretVersionDestructionRequest**](BulkCancelScheduledSecretVersionDestructionRequest.md) |  | 

### Return type

[**BulkCancelScheduledSecretVersionDestructionResponse**](BulkCancelScheduledSecretVersionDestructionResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkDeactivateSecretVersions

> BulkDeactivateSecretVersionsResponse BulkDeactivateSecretVersions(ctx, secretId).XAuthToken(xAuthToken).BulkDeactivateSecretVersionsRequest(bulkDeactivateSecretVersionsRequest).Execute()

Bulk deactivate secret versions 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/secretsmanager"
)

func main() {
	secretId := "secretId_example" // string | 시크릿의 고유 ID <br/> - [List secrets](https://docs.kakaocloud.com/openapi/security/secrets-manager/list-secrets)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bulkDeactivateSecretVersionsRequest := *openapiclient.NewBulkDeactivateSecretVersionsRequest(*openapiclient.NewBulkDeactivateSecretVersions([]int32{int32(123)})) // BulkDeactivateSecretVersionsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretVersionAPI.BulkDeactivateSecretVersions(context.Background(), secretId).XAuthToken(xAuthToken).BulkDeactivateSecretVersionsRequest(bulkDeactivateSecretVersionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretVersionAPI.BulkDeactivateSecretVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkDeactivateSecretVersions`: BulkDeactivateSecretVersionsResponse
	fmt.Fprintf(os.Stdout, "Response from `SecretVersionAPI.BulkDeactivateSecretVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretId** | **string** | 시크릿의 고유 ID &lt;br/&gt; - [List secrets](https://docs.kakaocloud.com/openapi/security/secrets-manager/list-secrets)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkDeactivateSecretVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bulkDeactivateSecretVersionsRequest** | [**BulkDeactivateSecretVersionsRequest**](BulkDeactivateSecretVersionsRequest.md) |  | 

### Return type

[**BulkDeactivateSecretVersionsResponse**](BulkDeactivateSecretVersionsResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkDestroySecretVersions

> BulkDestroySecretVersionsResponse BulkDestroySecretVersions(ctx, secretId).XAuthToken(xAuthToken).BulkDestroySecretVersionsRequest(bulkDestroySecretVersionsRequest).Execute()

Bulk destroy secret versions 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/secretsmanager"
)

func main() {
	secretId := "secretId_example" // string | 시크릿의 고유 ID <br/> - [List secrets](https://docs.kakaocloud.com/openapi/security/secrets-manager/list-secrets)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bulkDestroySecretVersionsRequest := *openapiclient.NewBulkDestroySecretVersionsRequest(*openapiclient.NewBulkDestroySecretVersions([]int32{int32(123)})) // BulkDestroySecretVersionsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretVersionAPI.BulkDestroySecretVersions(context.Background(), secretId).XAuthToken(xAuthToken).BulkDestroySecretVersionsRequest(bulkDestroySecretVersionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretVersionAPI.BulkDestroySecretVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkDestroySecretVersions`: BulkDestroySecretVersionsResponse
	fmt.Fprintf(os.Stdout, "Response from `SecretVersionAPI.BulkDestroySecretVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretId** | **string** | 시크릿의 고유 ID &lt;br/&gt; - [List secrets](https://docs.kakaocloud.com/openapi/security/secrets-manager/list-secrets)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkDestroySecretVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bulkDestroySecretVersionsRequest** | [**BulkDestroySecretVersionsRequest**](BulkDestroySecretVersionsRequest.md) |  | 

### Return type

[**BulkDestroySecretVersionsResponse**](BulkDestroySecretVersionsResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkScheduleSecretVersionDestruction

> BulkScheduleSecretVersionDestructionResponse BulkScheduleSecretVersionDestruction(ctx, secretId).XAuthToken(xAuthToken).BulkScheduleSecretVersionDestructionRequest(bulkScheduleSecretVersionDestructionRequest).Execute()

Bulk schedule secret version destruction 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/secretsmanager"
)

func main() {
	secretId := "secretId_example" // string | 시크릿의 고유 ID <br/> - [List secrets](https://docs.kakaocloud.com/openapi/security/secrets-manager/list-secrets)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bulkScheduleSecretVersionDestructionRequest := *openapiclient.NewBulkScheduleSecretVersionDestructionRequest(*openapiclient.NewBulkScheduleSecretVersionDestruction([]int32{int32(123)}, time.Now())) // BulkScheduleSecretVersionDestructionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretVersionAPI.BulkScheduleSecretVersionDestruction(context.Background(), secretId).XAuthToken(xAuthToken).BulkScheduleSecretVersionDestructionRequest(bulkScheduleSecretVersionDestructionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretVersionAPI.BulkScheduleSecretVersionDestruction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkScheduleSecretVersionDestruction`: BulkScheduleSecretVersionDestructionResponse
	fmt.Fprintf(os.Stdout, "Response from `SecretVersionAPI.BulkScheduleSecretVersionDestruction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretId** | **string** | 시크릿의 고유 ID &lt;br/&gt; - [List secrets](https://docs.kakaocloud.com/openapi/security/secrets-manager/list-secrets)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkScheduleSecretVersionDestructionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bulkScheduleSecretVersionDestructionRequest** | [**BulkScheduleSecretVersionDestructionRequest**](BulkScheduleSecretVersionDestructionRequest.md) |  | 

### Return type

[**BulkScheduleSecretVersionDestructionResponse**](BulkScheduleSecretVersionDestructionResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelScheduledSecretVersionDestruction

> CancelScheduledSecretVersionDestructionResponse CancelScheduledSecretVersionDestruction(ctx, secretId, version).XAuthToken(xAuthToken).Execute()

Cancel scheduled secret version destruction 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/secretsmanager"
)

func main() {
	secretId := "secretId_example" // string | 시크릿의 고유 ID <br/> - [List secrets](https://docs.kakaocloud.com/openapi/security/secrets-manager/list-secrets)에서 확인
	version := "version_example" // string | 대상 시크릿 버전 <br/>- [List secret versions](https://docs.kakaocloud.com/openapi/security/secrets-manager/list-secret-versions)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretVersionAPI.CancelScheduledSecretVersionDestruction(context.Background(), secretId, version).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretVersionAPI.CancelScheduledSecretVersionDestruction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelScheduledSecretVersionDestruction`: CancelScheduledSecretVersionDestructionResponse
	fmt.Fprintf(os.Stdout, "Response from `SecretVersionAPI.CancelScheduledSecretVersionDestruction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretId** | **string** | 시크릿의 고유 ID &lt;br/&gt; - [List secrets](https://docs.kakaocloud.com/openapi/security/secrets-manager/list-secrets)에서 확인 | 
**version** | **string** | 대상 시크릿 버전 &lt;br/&gt;- [List secret versions](https://docs.kakaocloud.com/openapi/security/secrets-manager/list-secret-versions)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelScheduledSecretVersionDestructionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**CancelScheduledSecretVersionDestructionResponse**](CancelScheduledSecretVersionDestructionResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSecretVersion

> CreateSecretVersionResponse CreateSecretVersion(ctx, secretId).XAuthToken(xAuthToken).CreateSecretVersionRequest(createSecretVersionRequest).Execute()

Create secret version 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/secretsmanager"
)

func main() {
	secretId := "secretId_example" // string | 시크릿의 고유 ID <br/> - [List secrets](https://docs.kakaocloud.com/openapi/security/secrets-manager/list-secrets)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	createSecretVersionRequest := *openapiclient.NewCreateSecretVersionRequest(*openapiclient.NewCreateSecretVersion(interface{}(123))) // CreateSecretVersionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretVersionAPI.CreateSecretVersion(context.Background(), secretId).XAuthToken(xAuthToken).CreateSecretVersionRequest(createSecretVersionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretVersionAPI.CreateSecretVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSecretVersion`: CreateSecretVersionResponse
	fmt.Fprintf(os.Stdout, "Response from `SecretVersionAPI.CreateSecretVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretId** | **string** | 시크릿의 고유 ID &lt;br/&gt; - [List secrets](https://docs.kakaocloud.com/openapi/security/secrets-manager/list-secrets)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecretVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **createSecretVersionRequest** | [**CreateSecretVersionRequest**](CreateSecretVersionRequest.md) |  | 

### Return type

[**CreateSecretVersionResponse**](CreateSecretVersionResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateSecretVersion

> DeactivateSecretVersionResponse DeactivateSecretVersion(ctx, secretId, version).XAuthToken(xAuthToken).Execute()

Deactivate secret version 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/secretsmanager"
)

func main() {
	secretId := "secretId_example" // string | 시크릿의 고유 ID <br/> - [List secrets](https://docs.kakaocloud.com/openapi/security/secrets-manager/list-secrets)에서 확인
	version := "version_example" // string | 대상 시크릿 버전 <br/>- [List secret versions](https://docs.kakaocloud.com/openapi/security/secrets-manager/list-secret-versions)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretVersionAPI.DeactivateSecretVersion(context.Background(), secretId, version).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretVersionAPI.DeactivateSecretVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeactivateSecretVersion`: DeactivateSecretVersionResponse
	fmt.Fprintf(os.Stdout, "Response from `SecretVersionAPI.DeactivateSecretVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretId** | **string** | 시크릿의 고유 ID &lt;br/&gt; - [List secrets](https://docs.kakaocloud.com/openapi/security/secrets-manager/list-secrets)에서 확인 | 
**version** | **string** | 대상 시크릿 버전 &lt;br/&gt;- [List secret versions](https://docs.kakaocloud.com/openapi/security/secrets-manager/list-secret-versions)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateSecretVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**DeactivateSecretVersionResponse**](DeactivateSecretVersionResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DestroySecretVersion

> DestroySecretVersion(ctx, secretId, version).XAuthToken(xAuthToken).Execute()

Destroy secret version 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/secretsmanager"
)

func main() {
	secretId := "secretId_example" // string | 시크릿의 고유 ID <br/> - [List secrets](https://docs.kakaocloud.com/openapi/security/secrets-manager/list-secrets)에서 확인
	version := "version_example" // string | 대상 시크릿 버전 <br/>- [List secret versions](https://docs.kakaocloud.com/openapi/security/secrets-manager/list-secret-versions)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecretVersionAPI.DestroySecretVersion(context.Background(), secretId, version).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretVersionAPI.DestroySecretVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretId** | **string** | 시크릿의 고유 ID &lt;br/&gt; - [List secrets](https://docs.kakaocloud.com/openapi/security/secrets-manager/list-secrets)에서 확인 | 
**version** | **string** | 대상 시크릿 버전 &lt;br/&gt;- [List secret versions](https://docs.kakaocloud.com/openapi/security/secrets-manager/list-secret-versions)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiDestroySecretVersionRequest struct via the builder pattern


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


## GetSecretDefaultVersion

> GetSecretDefaultVersionResponse GetSecretDefaultVersion(ctx, secretId).XAuthToken(xAuthToken).Execute()

Get secret default version 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/secretsmanager"
)

func main() {
	secretId := "secretId_example" // string | 시크릿의 고유 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretVersionAPI.GetSecretDefaultVersion(context.Background(), secretId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretVersionAPI.GetSecretDefaultVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecretDefaultVersion`: GetSecretDefaultVersionResponse
	fmt.Fprintf(os.Stdout, "Response from `SecretVersionAPI.GetSecretDefaultVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretId** | **string** | 시크릿의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretDefaultVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**GetSecretDefaultVersionResponse**](GetSecretDefaultVersionResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecretVersion

> GetSecretVersionResponse GetSecretVersion(ctx, secretId, version).XAuthToken(xAuthToken).Execute()

Get secret version 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/secretsmanager"
)

func main() {
	secretId := "secretId_example" // string | 시크릿의 고유 ID
	version := "version_example" // string | 대상 시크릿 버전 <br/>- [List secret versions](https://docs.kakaocloud.com/openapi/security/secrets-manager/list-secret-versions)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretVersionAPI.GetSecretVersion(context.Background(), secretId, version).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretVersionAPI.GetSecretVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecretVersion`: GetSecretVersionResponse
	fmt.Fprintf(os.Stdout, "Response from `SecretVersionAPI.GetSecretVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretId** | **string** | 시크릿의 고유 ID | 
**version** | **string** | 대상 시크릿 버전 &lt;br/&gt;- [List secret versions](https://docs.kakaocloud.com/openapi/security/secrets-manager/list-secret-versions)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**GetSecretVersionResponse**](GetSecretVersionResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSecretVersions

> ListSecretVersionsResponse ListSecretVersions(ctx, secretId).XAuthToken(xAuthToken).SecretId2(secretId2).Version(version).Status(status).KmsKeyName(kmsKeyName).KmsKeyId(kmsKeyId).KmsKeyVersion(kmsKeyVersion).Offset(offset).Limit(limit).SortKeys(sortKeys).SortDirs(sortDirs).Execute()

List secret versions 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/secretsmanager"
)

func main() {
	secretId := "secretId_example" // string | 시크릿의 고유 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	secretId2 := "secretId_example" // string | 조회할 시크릿의 고유 ID (optional)
	version := int32(56) // int32 | 조회할 시크릿 버전 (optional)
	status := openapiclient.Status("Active") // Status | 조회할 시크릿 버전 상태 (optional)
	kmsKeyName := "kmsKeyName_example" // string | 시크릿 보호에 사용된 KMS 키 이름 (optional)
	kmsKeyId := "kmsKeyId_example" // string | 시크릿 보호에 사용된 KMS 키의 고유 ID (optional)
	kmsKeyVersion := int32(56) // int32 | 시크릿 암호화에 사용된 KMS 키 버전 (optional)
	offset := int32(56) // int32 | 조회할 목록의 시작 위치 (optional)
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional)
	sortKeys := "sortKeys_example" // string | 정렬 기준 필드. 여러 값은 쉼표로 구분 (optional)
	sortDirs := "sortDirs_example" // string | 정렬 방향. 여러 값은 쉼표로 구분 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretVersionAPI.ListSecretVersions(context.Background(), secretId).XAuthToken(xAuthToken).SecretId2(secretId2).Version(version).Status(status).KmsKeyName(kmsKeyName).KmsKeyId(kmsKeyId).KmsKeyVersion(kmsKeyVersion).Offset(offset).Limit(limit).SortKeys(sortKeys).SortDirs(sortDirs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretVersionAPI.ListSecretVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSecretVersions`: ListSecretVersionsResponse
	fmt.Fprintf(os.Stdout, "Response from `SecretVersionAPI.ListSecretVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretId** | **string** | 시크릿의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSecretVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **secretId2** | **string** | 조회할 시크릿의 고유 ID | 
 **version** | **int32** | 조회할 시크릿 버전 | 
 **status** | [**Status**](Status.md) | 조회할 시크릿 버전 상태 | 
 **kmsKeyName** | **string** | 시크릿 보호에 사용된 KMS 키 이름 | 
 **kmsKeyId** | **string** | 시크릿 보호에 사용된 KMS 키의 고유 ID | 
 **kmsKeyVersion** | **int32** | 시크릿 암호화에 사용된 KMS 키 버전 | 
 **offset** | **int32** | 조회할 목록의 시작 위치 | 
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | 
 **sortKeys** | **string** | 정렬 기준 필드. 여러 값은 쉼표로 구분 | 
 **sortDirs** | **string** | 정렬 방향. 여러 값은 쉼표로 구분 | 

### Return type

[**ListSecretVersionsResponse**](ListSecretVersionsResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduleSecretVersionDestruction

> ScheduleSecretVersionDestructionResponse ScheduleSecretVersionDestruction(ctx, secretId, version).XAuthToken(xAuthToken).ScheduleSecretVersionDestructionRequest(scheduleSecretVersionDestructionRequest).Execute()

Schedule secret version destruction 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/secretsmanager"
)

func main() {
	secretId := "secretId_example" // string | 시크릿의 고유 ID <br/> - [List secrets](https://docs.kakaocloud.com/openapi/security/secrets-manager/list-secrets)에서 확인
	version := "version_example" // string | 대상 시크릿 버전 <br/>- [List secret versions](https://docs.kakaocloud.com/openapi/security/secrets-manager/list-secret-versions)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	scheduleSecretVersionDestructionRequest := *openapiclient.NewScheduleSecretVersionDestructionRequest(*openapiclient.NewScheduleSecretVersionDestruction(time.Now())) // ScheduleSecretVersionDestructionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretVersionAPI.ScheduleSecretVersionDestruction(context.Background(), secretId, version).XAuthToken(xAuthToken).ScheduleSecretVersionDestructionRequest(scheduleSecretVersionDestructionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretVersionAPI.ScheduleSecretVersionDestruction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScheduleSecretVersionDestruction`: ScheduleSecretVersionDestructionResponse
	fmt.Fprintf(os.Stdout, "Response from `SecretVersionAPI.ScheduleSecretVersionDestruction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretId** | **string** | 시크릿의 고유 ID &lt;br/&gt; - [List secrets](https://docs.kakaocloud.com/openapi/security/secrets-manager/list-secrets)에서 확인 | 
**version** | **string** | 대상 시크릿 버전 &lt;br/&gt;- [List secret versions](https://docs.kakaocloud.com/openapi/security/secrets-manager/list-secret-versions)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiScheduleSecretVersionDestructionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **scheduleSecretVersionDestructionRequest** | [**ScheduleSecretVersionDestructionRequest**](ScheduleSecretVersionDestructionRequest.md) |  | 

### Return type

[**ScheduleSecretVersionDestructionResponse**](ScheduleSecretVersionDestructionResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

