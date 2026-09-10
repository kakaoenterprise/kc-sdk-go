# \SecretAPI

All URIs are relative to *https://secrets-manager-service.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSecretsAccessTargets**](SecretAPI.md#AddSecretsAccessTargets) | **Post** /api/v1/secrets/{secret_id}/access-control/targets | Add secret access targets 
[**BulkDeleteSecrets**](SecretAPI.md#BulkDeleteSecrets) | **Post** /api/v1/secrets/deletions | Bulk delete secrets 
[**CreateSecret**](SecretAPI.md#CreateSecret) | **Post** /api/v1/secrets | Create secret 
[**DeleteSecret**](SecretAPI.md#DeleteSecret) | **Delete** /api/v1/secrets/{secret_id} | Delete secret 
[**GetSecret**](SecretAPI.md#GetSecret) | **Get** /api/v1/secrets/{secret_id} | Get secret 
[**ListSecrets**](SecretAPI.md#ListSecrets) | **Get** /api/v1/secrets | List secrets 
[**ListSecretsAccessTargets**](SecretAPI.md#ListSecretsAccessTargets) | **Get** /api/v1/secrets/{secret_id}/access-control/targets | List secret access targets 
[**RemoveSecretsAccessTargets**](SecretAPI.md#RemoveSecretsAccessTargets) | **Post** /api/v1/secrets/{secret_id}/access-control/targets/deletions | Remove secret access targets 
[**UpdateSecretKmsKey**](SecretAPI.md#UpdateSecretKmsKey) | **Patch** /api/v1/secrets/{secret_id}/kms-key | Update secret KMS key 
[**UpdateSecretsAccessControl**](SecretAPI.md#UpdateSecretsAccessControl) | **Put** /api/v1/secrets/{secret_id}/access-control | Update secret access control 



## AddSecretsAccessTargets

> AddSecretsAccessTargetsResponse AddSecretsAccessTargets(ctx, secretId).XAuthToken(xAuthToken).AddSecretsAccessTargetsRequest(addSecretsAccessTargetsRequest).Execute()

Add secret access targets 



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
	addSecretsAccessTargetsRequest := *openapiclient.NewAddSecretsAccessTargetsRequest(*openapiclient.NewAddSecretsAccessTargets([]string{"Ids_example"})) // AddSecretsAccessTargetsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretAPI.AddSecretsAccessTargets(context.Background(), secretId).XAuthToken(xAuthToken).AddSecretsAccessTargetsRequest(addSecretsAccessTargetsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretAPI.AddSecretsAccessTargets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddSecretsAccessTargets`: AddSecretsAccessTargetsResponse
	fmt.Fprintf(os.Stdout, "Response from `SecretAPI.AddSecretsAccessTargets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretId** | **string** | 시크릿의 고유 ID &lt;br/&gt; - [List secrets](https://docs.kakaocloud.com/openapi/security/secrets-manager/list-secrets)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddSecretsAccessTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **addSecretsAccessTargetsRequest** | [**AddSecretsAccessTargetsRequest**](AddSecretsAccessTargetsRequest.md) |  | 

### Return type

[**AddSecretsAccessTargetsResponse**](AddSecretsAccessTargetsResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkDeleteSecrets

> BulkDeleteSecretsResponse BulkDeleteSecrets(ctx).XAuthToken(xAuthToken).BulkDeleteSecretsRequest(bulkDeleteSecretsRequest).Execute()

Bulk delete secrets 



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
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bulkDeleteSecretsRequest := *openapiclient.NewBulkDeleteSecretsRequest(*openapiclient.NewBulkDeleteSecrets([]string{"Ids_example"})) // BulkDeleteSecretsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretAPI.BulkDeleteSecrets(context.Background()).XAuthToken(xAuthToken).BulkDeleteSecretsRequest(bulkDeleteSecretsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretAPI.BulkDeleteSecrets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkDeleteSecrets`: BulkDeleteSecretsResponse
	fmt.Fprintf(os.Stdout, "Response from `SecretAPI.BulkDeleteSecrets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkDeleteSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bulkDeleteSecretsRequest** | [**BulkDeleteSecretsRequest**](BulkDeleteSecretsRequest.md) |  | 

### Return type

[**BulkDeleteSecretsResponse**](BulkDeleteSecretsResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSecret

> CreateSecretResponse CreateSecret(ctx).XAuthToken(xAuthToken).CreateSecretRequest(createSecretRequest).Execute()

Create secret 



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
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	createSecretRequest := *openapiclient.NewCreateSecretRequest(*openapiclient.NewCreateSecret("Name_example", interface{}(123))) // CreateSecretRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretAPI.CreateSecret(context.Background()).XAuthToken(xAuthToken).CreateSecretRequest(createSecretRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretAPI.CreateSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSecret`: CreateSecretResponse
	fmt.Fprintf(os.Stdout, "Response from `SecretAPI.CreateSecret`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **createSecretRequest** | [**CreateSecretRequest**](CreateSecretRequest.md) |  | 

### Return type

[**CreateSecretResponse**](CreateSecretResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSecret

> DeleteSecret(ctx, secretId).XAuthToken(xAuthToken).Execute()

Delete secret 



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecretAPI.DeleteSecret(context.Background(), secretId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretAPI.DeleteSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretId** | **string** | 시크릿의 고유 ID &lt;br/&gt; - [List secrets](https://docs.kakaocloud.com/openapi/security/secrets-manager/list-secrets)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSecretRequest struct via the builder pattern


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


## GetSecret

> GetSecretResponse GetSecret(ctx, secretId).XAuthToken(xAuthToken).Execute()

Get secret 



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
	resp, r, err := apiClient.SecretAPI.GetSecret(context.Background(), secretId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretAPI.GetSecret``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecret`: GetSecretResponse
	fmt.Fprintf(os.Stdout, "Response from `SecretAPI.GetSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretId** | **string** | 시크릿의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**GetSecretResponse**](GetSecretResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSecrets

> ListSecretsResponse ListSecrets(ctx).XAuthToken(xAuthToken).Name(name).Id(id).Status(status).DefaultVersion(defaultVersion).KmsKeyName(kmsKeyName).KmsKeyId(kmsKeyId).CreatedBy(createdBy).Offset(offset).Limit(limit).SortKeys(sortKeys).SortDirs(sortDirs).Execute()

List secrets 



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
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	name := "name_example" // string | 조회할 시크릿 이름 (optional)
	id := "id_example" // string | 조회할 시크릿의 고유 ID (optional)
	status := openapiclient.Status("Active") // Status | 조회할 시크릿 상태 (optional)
	defaultVersion := int32(56) // int32 | 조회할 기본 버전 (optional)
	kmsKeyName := "kmsKeyName_example" // string | 시크릿 보호에 사용된 KMS 키 이름 (optional)
	kmsKeyId := "kmsKeyId_example" // string | 시크릿 보호에 사용된 KMS 키의 고유 ID (optional)
	createdBy := "createdBy_example" // string | 시크릿을 생성한 사용자 또는 주체 (optional)
	offset := int32(56) // int32 | 조회할 목록의 시작 위치 (optional)
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional)
	sortKeys := "sortKeys_example" // string | 정렬 기준 필드. 여러 값은 쉼표로 구분 (optional)
	sortDirs := "sortDirs_example" // string | 정렬 방향. 여러 값은 쉼표로 구분 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretAPI.ListSecrets(context.Background()).XAuthToken(xAuthToken).Name(name).Id(id).Status(status).DefaultVersion(defaultVersion).KmsKeyName(kmsKeyName).KmsKeyId(kmsKeyId).CreatedBy(createdBy).Offset(offset).Limit(limit).SortKeys(sortKeys).SortDirs(sortDirs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretAPI.ListSecrets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSecrets`: ListSecretsResponse
	fmt.Fprintf(os.Stdout, "Response from `SecretAPI.ListSecrets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSecretsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **name** | **string** | 조회할 시크릿 이름 | 
 **id** | **string** | 조회할 시크릿의 고유 ID | 
 **status** | [**Status**](Status.md) | 조회할 시크릿 상태 | 
 **defaultVersion** | **int32** | 조회할 기본 버전 | 
 **kmsKeyName** | **string** | 시크릿 보호에 사용된 KMS 키 이름 | 
 **kmsKeyId** | **string** | 시크릿 보호에 사용된 KMS 키의 고유 ID | 
 **createdBy** | **string** | 시크릿을 생성한 사용자 또는 주체 | 
 **offset** | **int32** | 조회할 목록의 시작 위치 | 
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | 
 **sortKeys** | **string** | 정렬 기준 필드. 여러 값은 쉼표로 구분 | 
 **sortDirs** | **string** | 정렬 방향. 여러 값은 쉼표로 구분 | 

### Return type

[**ListSecretsResponse**](ListSecretsResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSecretsAccessTargets

> ListSecretsAccessTargetsResponse ListSecretsAccessTargets(ctx, secretId).XAuthToken(xAuthToken).Execute()

List secret access targets 



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
	resp, r, err := apiClient.SecretAPI.ListSecretsAccessTargets(context.Background(), secretId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretAPI.ListSecretsAccessTargets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSecretsAccessTargets`: ListSecretsAccessTargetsResponse
	fmt.Fprintf(os.Stdout, "Response from `SecretAPI.ListSecretsAccessTargets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretId** | **string** | 시크릿의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSecretsAccessTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**ListSecretsAccessTargetsResponse**](ListSecretsAccessTargetsResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveSecretsAccessTargets

> RemoveSecretsAccessTargetsResponse RemoveSecretsAccessTargets(ctx, secretId).XAuthToken(xAuthToken).RemoveSecretsAccessTargetsRequest(removeSecretsAccessTargetsRequest).Execute()

Remove secret access targets 



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
	removeSecretsAccessTargetsRequest := *openapiclient.NewRemoveSecretsAccessTargetsRequest(*openapiclient.NewRemoveSecretsAccessTargets([]string{"Ids_example"})) // RemoveSecretsAccessTargetsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretAPI.RemoveSecretsAccessTargets(context.Background(), secretId).XAuthToken(xAuthToken).RemoveSecretsAccessTargetsRequest(removeSecretsAccessTargetsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretAPI.RemoveSecretsAccessTargets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveSecretsAccessTargets`: RemoveSecretsAccessTargetsResponse
	fmt.Fprintf(os.Stdout, "Response from `SecretAPI.RemoveSecretsAccessTargets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretId** | **string** | 시크릿의 고유 ID &lt;br/&gt; - [List secrets](https://docs.kakaocloud.com/openapi/security/secrets-manager/list-secrets)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveSecretsAccessTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **removeSecretsAccessTargetsRequest** | [**RemoveSecretsAccessTargetsRequest**](RemoveSecretsAccessTargetsRequest.md) |  | 

### Return type

[**RemoveSecretsAccessTargetsResponse**](RemoveSecretsAccessTargetsResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSecretKmsKey

> UpdateSecretKmsKeyResponse UpdateSecretKmsKey(ctx, secretId).XAuthToken(xAuthToken).UpdateSecretKmsKeyRequest(updateSecretKmsKeyRequest).Execute()

Update secret KMS key 



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
	updateSecretKmsKeyRequest := *openapiclient.NewUpdateSecretKmsKeyRequest(*openapiclient.NewUpdateSecretKmsKey()) // UpdateSecretKmsKeyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretAPI.UpdateSecretKmsKey(context.Background(), secretId).XAuthToken(xAuthToken).UpdateSecretKmsKeyRequest(updateSecretKmsKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretAPI.UpdateSecretKmsKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSecretKmsKey`: UpdateSecretKmsKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `SecretAPI.UpdateSecretKmsKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretId** | **string** | 시크릿의 고유 ID &lt;br/&gt; - [List secrets](https://docs.kakaocloud.com/openapi/security/secrets-manager/list-secrets)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSecretKmsKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **updateSecretKmsKeyRequest** | [**UpdateSecretKmsKeyRequest**](UpdateSecretKmsKeyRequest.md) |  | 

### Return type

[**UpdateSecretKmsKeyResponse**](UpdateSecretKmsKeyResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSecretsAccessControl

> UpdateSecretsAccessControlResponse UpdateSecretsAccessControl(ctx, secretId).XAuthToken(xAuthToken).UpdateSecretsAccessControlRequest(updateSecretsAccessControlRequest).Execute()

Update secret access control 



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
	updateSecretsAccessControlRequest := *openapiclient.NewUpdateSecretsAccessControlRequest(*openapiclient.NewUpdateSecretsAccessControl(false)) // UpdateSecretsAccessControlRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretAPI.UpdateSecretsAccessControl(context.Background(), secretId).XAuthToken(xAuthToken).UpdateSecretsAccessControlRequest(updateSecretsAccessControlRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretAPI.UpdateSecretsAccessControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSecretsAccessControl`: UpdateSecretsAccessControlResponse
	fmt.Fprintf(os.Stdout, "Response from `SecretAPI.UpdateSecretsAccessControl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**secretId** | **string** | 시크릿의 고유 ID &lt;br/&gt; - [List secrets](https://docs.kakaocloud.com/openapi/security/secrets-manager/list-secrets)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSecretsAccessControlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **updateSecretsAccessControlRequest** | [**UpdateSecretsAccessControlRequest**](UpdateSecretsAccessControlRequest.md) |  | 

### Return type

[**UpdateSecretsAccessControlResponse**](UpdateSecretsAccessControlResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

