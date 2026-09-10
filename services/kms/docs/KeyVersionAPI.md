# \KeyVersionAPI

All URIs are relative to *https://key-management-service.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkCancelScheduledKeyVersionDestruction**](KeyVersionAPI.md#BulkCancelScheduledKeyVersionDestruction) | **Patch** /api/v1/keys/{key_id}/versions/cancel-schedule-destroy | Bulk cancel scheduled key version destruction 
[**BulkScheduleKeyVersionDestruction**](KeyVersionAPI.md#BulkScheduleKeyVersionDestruction) | **Patch** /api/v1/keys/{key_id}/versions/schedule-destroy | Bulk schedule key version destruction 
[**CancelScheduledKeyVersionDestruction**](KeyVersionAPI.md#CancelScheduledKeyVersionDestruction) | **Patch** /api/v1/keys/{key_id}/versions/{version}/cancel-schedule-destroy | Cancel scheduled key version destruction 
[**CreateKeyVersion**](KeyVersionAPI.md#CreateKeyVersion) | **Post** /api/v1/keys/{key_id}/versions | Create key version 
[**DeactivateKeyVersion**](KeyVersionAPI.md#DeactivateKeyVersion) | **Patch** /api/v1/keys/{key_id}/versions/{version}/deactivate | Deactivate key version 
[**ListKeyVersions**](KeyVersionAPI.md#ListKeyVersions) | **Get** /api/v1/keys/{key_id}/versions | List key versions 
[**ScheduleKeyVersionDestruction**](KeyVersionAPI.md#ScheduleKeyVersionDestruction) | **Patch** /api/v1/keys/{key_id}/versions/{version}/schedule-destroy | Schedule key version destruction 



## BulkCancelScheduledKeyVersionDestruction

> BulkCancelScheduledKeyVersionDestructionResponse BulkCancelScheduledKeyVersionDestruction(ctx, keyId).XAuthToken(xAuthToken).BulkCancelScheduledKeyVersionDestructionRequest(bulkCancelScheduledKeyVersionDestructionRequest).Execute()

Bulk cancel scheduled key version destruction 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/kms"
)

func main() {
	keyId := "keyId_example" // string | KMS 키의 고유 ID <br/>- [List user keys](https://docs.kakaocloud.com/openapi/security/kms/list-user-keys)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bulkCancelScheduledKeyVersionDestructionRequest := *openapiclient.NewBulkCancelScheduledKeyVersionDestructionRequest(*openapiclient.NewBulkCancelScheduledKeyVersionDestruction([]int32{int32(123)})) // BulkCancelScheduledKeyVersionDestructionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeyVersionAPI.BulkCancelScheduledKeyVersionDestruction(context.Background(), keyId).XAuthToken(xAuthToken).BulkCancelScheduledKeyVersionDestructionRequest(bulkCancelScheduledKeyVersionDestructionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyVersionAPI.BulkCancelScheduledKeyVersionDestruction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkCancelScheduledKeyVersionDestruction`: BulkCancelScheduledKeyVersionDestructionResponse
	fmt.Fprintf(os.Stdout, "Response from `KeyVersionAPI.BulkCancelScheduledKeyVersionDestruction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | KMS 키의 고유 ID &lt;br/&gt;- [List user keys](https://docs.kakaocloud.com/openapi/security/kms/list-user-keys)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkCancelScheduledKeyVersionDestructionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bulkCancelScheduledKeyVersionDestructionRequest** | [**BulkCancelScheduledKeyVersionDestructionRequest**](BulkCancelScheduledKeyVersionDestructionRequest.md) |  | 

### Return type

[**BulkCancelScheduledKeyVersionDestructionResponse**](BulkCancelScheduledKeyVersionDestructionResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkScheduleKeyVersionDestruction

> BulkScheduleKeyVersionDestructionResponse BulkScheduleKeyVersionDestruction(ctx, keyId).XAuthToken(xAuthToken).BulkScheduleKeyVersionDestructionRequest(bulkScheduleKeyVersionDestructionRequest).Execute()

Bulk schedule key version destruction 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/kms"
)

func main() {
	keyId := "keyId_example" // string | KMS 키의 고유 ID <br/>- [List user keys](https://docs.kakaocloud.com/openapi/security/kms/list-user-keys)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bulkScheduleKeyVersionDestructionRequest := *openapiclient.NewBulkScheduleKeyVersionDestructionRequest(*openapiclient.NewBulkScheduleKeyVersionDestruction([]int32{int32(123)}, "DestructionScheduledAt_example")) // BulkScheduleKeyVersionDestructionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeyVersionAPI.BulkScheduleKeyVersionDestruction(context.Background(), keyId).XAuthToken(xAuthToken).BulkScheduleKeyVersionDestructionRequest(bulkScheduleKeyVersionDestructionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyVersionAPI.BulkScheduleKeyVersionDestruction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkScheduleKeyVersionDestruction`: BulkScheduleKeyVersionDestructionResponse
	fmt.Fprintf(os.Stdout, "Response from `KeyVersionAPI.BulkScheduleKeyVersionDestruction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | KMS 키의 고유 ID &lt;br/&gt;- [List user keys](https://docs.kakaocloud.com/openapi/security/kms/list-user-keys)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkScheduleKeyVersionDestructionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bulkScheduleKeyVersionDestructionRequest** | [**BulkScheduleKeyVersionDestructionRequest**](BulkScheduleKeyVersionDestructionRequest.md) |  | 

### Return type

[**BulkScheduleKeyVersionDestructionResponse**](BulkScheduleKeyVersionDestructionResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelScheduledKeyVersionDestruction

> CancelScheduledKeyVersionDestructionResponse CancelScheduledKeyVersionDestruction(ctx, keyId, version).XAuthToken(xAuthToken).Execute()

Cancel scheduled key version destruction 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/kms"
)

func main() {
	keyId := "keyId_example" // string | KMS 키의 고유 ID <br/>- [List user keys](https://docs.kakaocloud.com/openapi/security/kms/list-user-keys)에서 확인
	version := "version_example" // string | 대상 키 버전 <br/>- [List key versions](https://docs.kakaocloud.com/openapi/security/kms/list-key-versions)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeyVersionAPI.CancelScheduledKeyVersionDestruction(context.Background(), keyId, version).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyVersionAPI.CancelScheduledKeyVersionDestruction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelScheduledKeyVersionDestruction`: CancelScheduledKeyVersionDestructionResponse
	fmt.Fprintf(os.Stdout, "Response from `KeyVersionAPI.CancelScheduledKeyVersionDestruction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | KMS 키의 고유 ID &lt;br/&gt;- [List user keys](https://docs.kakaocloud.com/openapi/security/kms/list-user-keys)에서 확인 | 
**version** | **string** | 대상 키 버전 &lt;br/&gt;- [List key versions](https://docs.kakaocloud.com/openapi/security/kms/list-key-versions)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelScheduledKeyVersionDestructionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**CancelScheduledKeyVersionDestructionResponse**](CancelScheduledKeyVersionDestructionResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateKeyVersion

> CreateKeyVersionResponse CreateKeyVersion(ctx, keyId).XAuthToken(xAuthToken).Execute()

Create key version 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/kms"
)

func main() {
	keyId := "keyId_example" // string | KMS 키의 고유 ID <br/>- [List user keys](https://docs.kakaocloud.com/openapi/security/kms/list-user-keys)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeyVersionAPI.CreateKeyVersion(context.Background(), keyId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyVersionAPI.CreateKeyVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateKeyVersion`: CreateKeyVersionResponse
	fmt.Fprintf(os.Stdout, "Response from `KeyVersionAPI.CreateKeyVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | KMS 키의 고유 ID &lt;br/&gt;- [List user keys](https://docs.kakaocloud.com/openapi/security/kms/list-user-keys)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateKeyVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**CreateKeyVersionResponse**](CreateKeyVersionResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateKeyVersion

> DeactivateKeyVersionResponse DeactivateKeyVersion(ctx, keyId, version).XAuthToken(xAuthToken).Execute()

Deactivate key version 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/kms"
)

func main() {
	keyId := "keyId_example" // string | KMS 키의 고유 ID <br/>- [List user keys](https://docs.kakaocloud.com/openapi/security/kms/list-user-keys)에서 확인
	version := "version_example" // string | 대상 키 버전 <br/>- [List key versions](https://docs.kakaocloud.com/openapi/security/kms/list-key-versions)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeyVersionAPI.DeactivateKeyVersion(context.Background(), keyId, version).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyVersionAPI.DeactivateKeyVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeactivateKeyVersion`: DeactivateKeyVersionResponse
	fmt.Fprintf(os.Stdout, "Response from `KeyVersionAPI.DeactivateKeyVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | KMS 키의 고유 ID &lt;br/&gt;- [List user keys](https://docs.kakaocloud.com/openapi/security/kms/list-user-keys)에서 확인 | 
**version** | **string** | 대상 키 버전 &lt;br/&gt;- [List key versions](https://docs.kakaocloud.com/openapi/security/kms/list-key-versions)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateKeyVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**DeactivateKeyVersionResponse**](DeactivateKeyVersionResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKeyVersions

> ListKeyVersionsResponse ListKeyVersions(ctx, keyId).XAuthToken(xAuthToken).Version(version).Status(status).Offset(offset).Limit(limit).SortKeys(sortKeys).SortDirs(sortDirs).Execute()

List key versions 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/kms"
)

func main() {
	keyId := "keyId_example" // string | KMS 키의 고유 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	version := "version_example" // string | 조회할 KMS 키 버전 (optional)
	status := openapiclient.KeyVersionStatus("ACTIVE") // KeyVersionStatus | 조회할 KMS 키 상태 (optional)
	offset := int32(56) // int32 | 조회할 목록의 시작 위치 (optional)
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional)
	sortKeys := "sortKeys_example" // string | 정렬 기준 필드. 여러 값은 쉼표로 구분 (optional)
	sortDirs := "sortDirs_example" // string | 정렬 방향. 여러 값은 쉼표로 구분 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeyVersionAPI.ListKeyVersions(context.Background(), keyId).XAuthToken(xAuthToken).Version(version).Status(status).Offset(offset).Limit(limit).SortKeys(sortKeys).SortDirs(sortDirs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyVersionAPI.ListKeyVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListKeyVersions`: ListKeyVersionsResponse
	fmt.Fprintf(os.Stdout, "Response from `KeyVersionAPI.ListKeyVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | KMS 키의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListKeyVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **version** | **string** | 조회할 KMS 키 버전 | 
 **status** | [**KeyVersionStatus**](KeyVersionStatus.md) | 조회할 KMS 키 상태 | 
 **offset** | **int32** | 조회할 목록의 시작 위치 | 
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | 
 **sortKeys** | **string** | 정렬 기준 필드. 여러 값은 쉼표로 구분 | 
 **sortDirs** | **string** | 정렬 방향. 여러 값은 쉼표로 구분 | 

### Return type

[**ListKeyVersionsResponse**](ListKeyVersionsResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduleKeyVersionDestruction

> ScheduleKeyVersionDestructionResponse ScheduleKeyVersionDestruction(ctx, keyId, version).XAuthToken(xAuthToken).ScheduleKeyVersionDestructionRequest(scheduleKeyVersionDestructionRequest).Execute()

Schedule key version destruction 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/kms"
)

func main() {
	keyId := "keyId_example" // string | KMS 키의 고유 ID <br/>- [List user keys](https://docs.kakaocloud.com/openapi/security/kms/list-user-keys)에서 확인
	version := "version_example" // string | 대상 키 버전 <br/>- [List key versions](https://docs.kakaocloud.com/openapi/security/kms/list-key-versions)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	scheduleKeyVersionDestructionRequest := *openapiclient.NewScheduleKeyVersionDestructionRequest(*openapiclient.NewScheduleKeyVersionDestruction("DestructionScheduledAt_example")) // ScheduleKeyVersionDestructionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeyVersionAPI.ScheduleKeyVersionDestruction(context.Background(), keyId, version).XAuthToken(xAuthToken).ScheduleKeyVersionDestructionRequest(scheduleKeyVersionDestructionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyVersionAPI.ScheduleKeyVersionDestruction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ScheduleKeyVersionDestruction`: ScheduleKeyVersionDestructionResponse
	fmt.Fprintf(os.Stdout, "Response from `KeyVersionAPI.ScheduleKeyVersionDestruction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | KMS 키의 고유 ID &lt;br/&gt;- [List user keys](https://docs.kakaocloud.com/openapi/security/kms/list-user-keys)에서 확인 | 
**version** | **string** | 대상 키 버전 &lt;br/&gt;- [List key versions](https://docs.kakaocloud.com/openapi/security/kms/list-key-versions)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiScheduleKeyVersionDestructionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **scheduleKeyVersionDestructionRequest** | [**ScheduleKeyVersionDestructionRequest**](ScheduleKeyVersionDestructionRequest.md) |  | 

### Return type

[**ScheduleKeyVersionDestructionResponse**](ScheduleKeyVersionDestructionResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

