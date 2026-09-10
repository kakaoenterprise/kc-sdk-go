# \KeyAPI

All URIs are relative to *https://key-management-service.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUserKeyAccessTargets**](KeyAPI.md#AddUserKeyAccessTargets) | **Post** /api/v1/keys/{key_id}/access-control/targets | Add user key access targets 
[**BulkDeleteUserKeys**](KeyAPI.md#BulkDeleteUserKeys) | **Post** /api/v1/keys/deletions | Bulk delete user keys 
[**CreateUserKey**](KeyAPI.md#CreateUserKey) | **Post** /api/v1/keys | Create user key 
[**DeleteUserKey**](KeyAPI.md#DeleteUserKey) | **Delete** /api/v1/keys/{key_id} | Delete user key 
[**GetKey**](KeyAPI.md#GetKey) | **Get** /api/v1/keys/{key_id} | Get key 
[**GetKeyByName**](KeyAPI.md#GetKeyByName) | **Get** /api/v1/keys/name/{key_name} | Get key by name 
[**GetPublicKey**](KeyAPI.md#GetPublicKey) | **Get** /api/v1/keys/{key_id}/public-key | Get public key 
[**ListKeyAlgorithms**](KeyAPI.md#ListKeyAlgorithms) | **Get** /api/v1/algorithms | List key algorithms 
[**ListServiceKeys**](KeyAPI.md#ListServiceKeys) | **Get** /api/v1/service-keys | List service keys 
[**ListSymmetricKeys**](KeyAPI.md#ListSymmetricKeys) | **Get** /api/v1/symmetric-keys | List symmetric keys 
[**ListUserKeyAccessTargets**](KeyAPI.md#ListUserKeyAccessTargets) | **Get** /api/v1/keys/{key_id}/access-control/targets | List user key access targets 
[**ListUserKeys**](KeyAPI.md#ListUserKeys) | **Get** /api/v1/keys | List user keys 
[**RemoveUserKeyAccessTargets**](KeyAPI.md#RemoveUserKeyAccessTargets) | **Post** /api/v1/keys/{key_id}/access-control/targets/deletions | Remove user key access targets 
[**UpdateUserKeyAccessControl**](KeyAPI.md#UpdateUserKeyAccessControl) | **Put** /api/v1/keys/{key_id}/access-control | Update user key access control 
[**UpdateUserKeyRotation**](KeyAPI.md#UpdateUserKeyRotation) | **Patch** /api/v1/keys/{key_id}/rotation | Update user key rotation 



## AddUserKeyAccessTargets

> AddUserKeyAccessTargetsResponse AddUserKeyAccessTargets(ctx, keyId).XAuthToken(xAuthToken).AddUserKeyAccessTargetsRequest(addUserKeyAccessTargetsRequest).Execute()

Add user key access targets 



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
	addUserKeyAccessTargetsRequest := *openapiclient.NewAddUserKeyAccessTargetsRequest(*openapiclient.NewAddUserKeyAccessTargets([]string{"Ids_example"})) // AddUserKeyAccessTargetsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeyAPI.AddUserKeyAccessTargets(context.Background(), keyId).XAuthToken(xAuthToken).AddUserKeyAccessTargetsRequest(addUserKeyAccessTargetsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyAPI.AddUserKeyAccessTargets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddUserKeyAccessTargets`: AddUserKeyAccessTargetsResponse
	fmt.Fprintf(os.Stdout, "Response from `KeyAPI.AddUserKeyAccessTargets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | KMS 키의 고유 ID &lt;br/&gt;- [List user keys](https://docs.kakaocloud.com/openapi/security/kms/list-user-keys)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddUserKeyAccessTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **addUserKeyAccessTargetsRequest** | [**AddUserKeyAccessTargetsRequest**](AddUserKeyAccessTargetsRequest.md) |  | 

### Return type

[**AddUserKeyAccessTargetsResponse**](AddUserKeyAccessTargetsResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkDeleteUserKeys

> BulkDeleteUserKeysResponse BulkDeleteUserKeys(ctx).XAuthToken(xAuthToken).BulkDeleteUserKeysRequest(bulkDeleteUserKeysRequest).Execute()

Bulk delete user keys 



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
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bulkDeleteUserKeysRequest := *openapiclient.NewBulkDeleteUserKeysRequest(*openapiclient.NewBulkDeleteUserKeys([]string{"Ids_example"})) // BulkDeleteUserKeysRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeyAPI.BulkDeleteUserKeys(context.Background()).XAuthToken(xAuthToken).BulkDeleteUserKeysRequest(bulkDeleteUserKeysRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyAPI.BulkDeleteUserKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkDeleteUserKeys`: BulkDeleteUserKeysResponse
	fmt.Fprintf(os.Stdout, "Response from `KeyAPI.BulkDeleteUserKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkDeleteUserKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bulkDeleteUserKeysRequest** | [**BulkDeleteUserKeysRequest**](BulkDeleteUserKeysRequest.md) |  | 

### Return type

[**BulkDeleteUserKeysResponse**](BulkDeleteUserKeysResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUserKey

> CreateUserKeyResponse CreateUserKey(ctx).XAuthToken(xAuthToken).CreateUserKeyRequest(createUserKeyRequest).Execute()

Create user key 



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
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	createUserKeyRequest := *openapiclient.NewCreateUserKeyRequest(*openapiclient.NewCreateUserKey("Name_example", openapiclient.KeyType("SYMMETRIC"), openapiclient.KeyPurpose("SYMMETRIC_ENCRYPT_DECRYPT"), openapiclient.KeyAlgorithm("AES256-GCM96"), false)) // CreateUserKeyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeyAPI.CreateUserKey(context.Background()).XAuthToken(xAuthToken).CreateUserKeyRequest(createUserKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyAPI.CreateUserKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUserKey`: CreateUserKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `KeyAPI.CreateUserKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **createUserKeyRequest** | [**CreateUserKeyRequest**](CreateUserKeyRequest.md) |  | 

### Return type

[**CreateUserKeyResponse**](CreateUserKeyResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserKey

> DeleteUserKey(ctx, keyId).XAuthToken(xAuthToken).Execute()

Delete user key 



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
	r, err := apiClient.KeyAPI.DeleteUserKey(context.Background(), keyId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyAPI.DeleteUserKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | KMS 키의 고유 ID &lt;br/&gt;- [List user keys](https://docs.kakaocloud.com/openapi/security/kms/list-user-keys)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserKeyRequest struct via the builder pattern


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


## GetKey

> GetKeyResponse GetKey(ctx, keyId).XAuthToken(xAuthToken).Execute()

Get key 



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeyAPI.GetKey(context.Background(), keyId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyAPI.GetKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKey`: GetKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `KeyAPI.GetKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | KMS 키의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**GetKeyResponse**](GetKeyResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKeyByName

> GetKeyByNameResponse GetKeyByName(ctx, keyName).XAuthToken(xAuthToken).Execute()

Get key by name 



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
	keyName := "keyName_example" // string | KMS 키 이름
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeyAPI.GetKeyByName(context.Background(), keyName).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyAPI.GetKeyByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKeyByName`: GetKeyByNameResponse
	fmt.Fprintf(os.Stdout, "Response from `KeyAPI.GetKeyByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyName** | **string** | KMS 키 이름 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKeyByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**GetKeyByNameResponse**](GetKeyByNameResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPublicKey

> GetPublicKeyResponse GetPublicKey(ctx, keyId).XAuthToken(xAuthToken).Execute()

Get public key 



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeyAPI.GetPublicKey(context.Background(), keyId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyAPI.GetPublicKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPublicKey`: GetPublicKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `KeyAPI.GetPublicKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | KMS 키의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**GetPublicKeyResponse**](GetPublicKeyResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKeyAlgorithms

> ListKeyAlgorithmsResponse ListKeyAlgorithms(ctx).XAuthToken(xAuthToken).Type_(type_).Purpose(purpose).Execute()

List key algorithms 



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
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	type_ := openapiclient.KeyType("SYMMETRIC") // KeyType | 조회할 키 유형 (optional)
	purpose := openapiclient.KeyPurpose("SYMMETRIC_ENCRYPT_DECRYPT") // KeyPurpose | 조회할 키 용도 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeyAPI.ListKeyAlgorithms(context.Background()).XAuthToken(xAuthToken).Type_(type_).Purpose(purpose).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyAPI.ListKeyAlgorithms``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListKeyAlgorithms`: ListKeyAlgorithmsResponse
	fmt.Fprintf(os.Stdout, "Response from `KeyAPI.ListKeyAlgorithms`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListKeyAlgorithmsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **type_** | [**KeyType**](KeyType.md) | 조회할 키 유형 | 
 **purpose** | [**KeyPurpose**](KeyPurpose.md) | 조회할 키 용도 | 

### Return type

[**ListKeyAlgorithmsResponse**](ListKeyAlgorithmsResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceKeys

> ListServiceKeysResponse ListServiceKeys(ctx).XAuthToken(xAuthToken).Name(name).Service(service).Offset(offset).Limit(limit).SortKeys(sortKeys).SortDirs(sortDirs).Execute()

List service keys 



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
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	name := "name_example" // string | 조회할 서비스 키 이름 (optional)
	service := "service_example" // string | 조회할 서비스 이름 (optional)
	offset := int32(56) // int32 | 조회할 목록의 시작 위치 (optional)
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional)
	sortKeys := "sortKeys_example" // string | 정렬 기준 필드. 여러 값은 쉼표로 구분 (optional)
	sortDirs := "sortDirs_example" // string | 정렬 방향. 여러 값은 쉼표로 구분 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeyAPI.ListServiceKeys(context.Background()).XAuthToken(xAuthToken).Name(name).Service(service).Offset(offset).Limit(limit).SortKeys(sortKeys).SortDirs(sortDirs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyAPI.ListServiceKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListServiceKeys`: ListServiceKeysResponse
	fmt.Fprintf(os.Stdout, "Response from `KeyAPI.ListServiceKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListServiceKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **name** | **string** | 조회할 서비스 키 이름 | 
 **service** | **string** | 조회할 서비스 이름 | 
 **offset** | **int32** | 조회할 목록의 시작 위치 | 
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | 
 **sortKeys** | **string** | 정렬 기준 필드. 여러 값은 쉼표로 구분 | 
 **sortDirs** | **string** | 정렬 방향. 여러 값은 쉼표로 구분 | 

### Return type

[**ListServiceKeysResponse**](ListServiceKeysResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSymmetricKeys

> ListSymmetricKeysResponse ListSymmetricKeys(ctx).XAuthToken(xAuthToken).Execute()

List symmetric keys 



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
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeyAPI.ListSymmetricKeys(context.Background()).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyAPI.ListSymmetricKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSymmetricKeys`: ListSymmetricKeysResponse
	fmt.Fprintf(os.Stdout, "Response from `KeyAPI.ListSymmetricKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSymmetricKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**ListSymmetricKeysResponse**](ListSymmetricKeysResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserKeyAccessTargets

> ListUserKeyAccessTargetsResponse ListUserKeyAccessTargets(ctx, keyId).XAuthToken(xAuthToken).Execute()

List user key access targets 



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeyAPI.ListUserKeyAccessTargets(context.Background(), keyId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyAPI.ListUserKeyAccessTargets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUserKeyAccessTargets`: ListUserKeyAccessTargetsResponse
	fmt.Fprintf(os.Stdout, "Response from `KeyAPI.ListUserKeyAccessTargets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | KMS 키의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserKeyAccessTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**ListUserKeyAccessTargetsResponse**](ListUserKeyAccessTargetsResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUserKeys

> ListUserKeysResponse ListUserKeys(ctx).XAuthToken(xAuthToken).Name(name).Id(id).Type_(type_).Purpose(purpose).Status(status).Algorithm(algorithm).DefaultVersion(defaultVersion).CreatedBy(createdBy).IsAccessControlEnabled(isAccessControlEnabled).IsAccessible(isAccessible).Offset(offset).Limit(limit).SortKeys(sortKeys).SortDirs(sortDirs).Execute()

List user keys 



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
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	name := "name_example" // string | 조회할 KMS 키 이름 (optional)
	id := "id_example" // string | 조회할 KMS 키의 고유 ID (optional)
	type_ := openapiclient.KeyType("SYMMETRIC") // KeyType | 조회할 키 유형 (optional)
	purpose := openapiclient.KeyPurpose("SYMMETRIC_ENCRYPT_DECRYPT") // KeyPurpose | 조회할 키 용도 (optional)
	status := openapiclient.KeyStatus("PRE_ACTIVATION") // KeyStatus | 조회할 KMS 키 상태 (optional)
	algorithm := openapiclient.KeyAlgorithm("AES256-GCM96") // KeyAlgorithm | 조회할 키 알고리즘 (optional)
	defaultVersion := "defaultVersion_example" // string | 조회할 기본 버전 (optional)
	createdBy := "createdBy_example" // string | KMS 키을 생성한 사용자 또는 주체 (optional)
	isAccessControlEnabled := true // bool | 접근 제어 활성화 여부 (optional)
	isAccessible := true // bool | 요청자의 접근 가능 여부 (optional)
	offset := int32(56) // int32 | 조회할 목록의 시작 위치 (optional)
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional)
	sortKeys := "sortKeys_example" // string | 정렬 기준 필드. 여러 값은 쉼표로 구분 (optional)
	sortDirs := "sortDirs_example" // string | 정렬 방향. 여러 값은 쉼표로 구분 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeyAPI.ListUserKeys(context.Background()).XAuthToken(xAuthToken).Name(name).Id(id).Type_(type_).Purpose(purpose).Status(status).Algorithm(algorithm).DefaultVersion(defaultVersion).CreatedBy(createdBy).IsAccessControlEnabled(isAccessControlEnabled).IsAccessible(isAccessible).Offset(offset).Limit(limit).SortKeys(sortKeys).SortDirs(sortDirs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyAPI.ListUserKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListUserKeys`: ListUserKeysResponse
	fmt.Fprintf(os.Stdout, "Response from `KeyAPI.ListUserKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUserKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **name** | **string** | 조회할 KMS 키 이름 | 
 **id** | **string** | 조회할 KMS 키의 고유 ID | 
 **type_** | [**KeyType**](KeyType.md) | 조회할 키 유형 | 
 **purpose** | [**KeyPurpose**](KeyPurpose.md) | 조회할 키 용도 | 
 **status** | [**KeyStatus**](KeyStatus.md) | 조회할 KMS 키 상태 | 
 **algorithm** | [**KeyAlgorithm**](KeyAlgorithm.md) | 조회할 키 알고리즘 | 
 **defaultVersion** | **string** | 조회할 기본 버전 | 
 **createdBy** | **string** | KMS 키을 생성한 사용자 또는 주체 | 
 **isAccessControlEnabled** | **bool** | 접근 제어 활성화 여부 | 
 **isAccessible** | **bool** | 요청자의 접근 가능 여부 | 
 **offset** | **int32** | 조회할 목록의 시작 위치 | 
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | 
 **sortKeys** | **string** | 정렬 기준 필드. 여러 값은 쉼표로 구분 | 
 **sortDirs** | **string** | 정렬 방향. 여러 값은 쉼표로 구분 | 

### Return type

[**ListUserKeysResponse**](ListUserKeysResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveUserKeyAccessTargets

> RemoveUserKeyAccessTargetsResponse RemoveUserKeyAccessTargets(ctx, keyId).XAuthToken(xAuthToken).RemoveUserKeyAccessTargetsRequest(removeUserKeyAccessTargetsRequest).Execute()

Remove user key access targets 



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
	removeUserKeyAccessTargetsRequest := *openapiclient.NewRemoveUserKeyAccessTargetsRequest(*openapiclient.NewRemoveUserKeyAccessTargets([]string{"Ids_example"})) // RemoveUserKeyAccessTargetsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeyAPI.RemoveUserKeyAccessTargets(context.Background(), keyId).XAuthToken(xAuthToken).RemoveUserKeyAccessTargetsRequest(removeUserKeyAccessTargetsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyAPI.RemoveUserKeyAccessTargets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveUserKeyAccessTargets`: RemoveUserKeyAccessTargetsResponse
	fmt.Fprintf(os.Stdout, "Response from `KeyAPI.RemoveUserKeyAccessTargets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | KMS 키의 고유 ID &lt;br/&gt;- [List user keys](https://docs.kakaocloud.com/openapi/security/kms/list-user-keys)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveUserKeyAccessTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **removeUserKeyAccessTargetsRequest** | [**RemoveUserKeyAccessTargetsRequest**](RemoveUserKeyAccessTargetsRequest.md) |  | 

### Return type

[**RemoveUserKeyAccessTargetsResponse**](RemoveUserKeyAccessTargetsResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserKeyAccessControl

> UpdateUserKeyAccessControlResponse UpdateUserKeyAccessControl(ctx, keyId).XAuthToken(xAuthToken).UpdateUserKeyAccessControlRequest(updateUserKeyAccessControlRequest).Execute()

Update user key access control 



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
	updateUserKeyAccessControlRequest := *openapiclient.NewUpdateUserKeyAccessControlRequest(*openapiclient.NewUpdateUserKeyAccessControl(false)) // UpdateUserKeyAccessControlRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeyAPI.UpdateUserKeyAccessControl(context.Background(), keyId).XAuthToken(xAuthToken).UpdateUserKeyAccessControlRequest(updateUserKeyAccessControlRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyAPI.UpdateUserKeyAccessControl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserKeyAccessControl`: UpdateUserKeyAccessControlResponse
	fmt.Fprintf(os.Stdout, "Response from `KeyAPI.UpdateUserKeyAccessControl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | KMS 키의 고유 ID &lt;br/&gt;- [List user keys](https://docs.kakaocloud.com/openapi/security/kms/list-user-keys)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserKeyAccessControlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **updateUserKeyAccessControlRequest** | [**UpdateUserKeyAccessControlRequest**](UpdateUserKeyAccessControlRequest.md) |  | 

### Return type

[**UpdateUserKeyAccessControlResponse**](UpdateUserKeyAccessControlResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserKeyRotation

> UpdateUserKeyRotationResponse UpdateUserKeyRotation(ctx, keyId).XAuthToken(xAuthToken).UpdateUserKeyRotationRequest(updateUserKeyRotationRequest).Execute()

Update user key rotation 



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
	updateUserKeyRotationRequest := *openapiclient.NewUpdateUserKeyRotationRequest(*openapiclient.NewUpdateUserKeyRotation(int32(123))) // UpdateUserKeyRotationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeyAPI.UpdateUserKeyRotation(context.Background(), keyId).XAuthToken(xAuthToken).UpdateUserKeyRotationRequest(updateUserKeyRotationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeyAPI.UpdateUserKeyRotation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUserKeyRotation`: UpdateUserKeyRotationResponse
	fmt.Fprintf(os.Stdout, "Response from `KeyAPI.UpdateUserKeyRotation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | KMS 키의 고유 ID &lt;br/&gt;- [List user keys](https://docs.kakaocloud.com/openapi/security/kms/list-user-keys)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserKeyRotationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **updateUserKeyRotationRequest** | [**UpdateUserKeyRotationRequest**](UpdateUserKeyRotationRequest.md) |  | 

### Return type

[**UpdateUserKeyRotationResponse**](UpdateUserKeyRotationResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

