# \PublicKeyAPI

All URIs are relative to *https://key-management-service.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDataKey**](PublicKeyAPI.md#CreateDataKey) | **Post** /api/v1/keys/{key_id}/data-key | Create data key 
[**DecryptData**](PublicKeyAPI.md#DecryptData) | **Post** /api/v1/keys/{key_id}/decrypt | Decrypt data
[**EncryptData**](PublicKeyAPI.md#EncryptData) | **Post** /api/v1/keys/{key_id}/encrypt | Encrypt data
[**GenerateHmac**](PublicKeyAPI.md#GenerateHmac) | **Post** /api/v1/keys/{key_id}/hmac | Generate HMAC
[**SignData**](PublicKeyAPI.md#SignData) | **Post** /api/v1/keys/{key_id}/sign | Sign data
[**VerifyHmac**](PublicKeyAPI.md#VerifyHmac) | **Post** /api/v1/keys/{key_id}/hmac/verify | Verify HMAC
[**VerifySignature**](PublicKeyAPI.md#VerifySignature) | **Post** /api/v1/keys/{key_id}/sign/verify | Verify signature



## CreateDataKey

> CreateDataKeyResponse CreateDataKey(ctx, keyId).Execute()

Create data key 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/kms"
)

func main() {
	keyId := "keyId_example" // string | KMS 키의 고유 ID - [List user keys](/openapi/security/kms/list-user-keys)에서 확인

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicKeyAPI.CreateDataKey(context.Background(), keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicKeyAPI.CreateDataKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDataKey`: CreateDataKeyResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicKeyAPI.CreateDataKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | KMS 키의 고유 ID - [List user keys](/openapi/security/kms/list-user-keys)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDataKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateDataKeyResponse**](CreateDataKeyResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DecryptData

> DecryptDataResponse DecryptData(ctx, keyId).DecryptDataRequest(decryptDataRequest).Execute()

Decrypt data



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/kms"
)

func main() {
	keyId := "keyId_example" // string | KMS 키의 고유 ID - [List user keys](/openapi/security/kms/list-user-keys)에서 확인
	decryptDataRequest := *openapiclient.NewDecryptDataRequest(*openapiclient.NewDecryptData("CipherText_example")) // DecryptDataRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicKeyAPI.DecryptData(context.Background(), keyId).DecryptDataRequest(decryptDataRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicKeyAPI.DecryptData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DecryptData`: DecryptDataResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicKeyAPI.DecryptData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | KMS 키의 고유 ID - [List user keys](/openapi/security/kms/list-user-keys)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiDecryptDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **decryptDataRequest** | [**DecryptDataRequest**](DecryptDataRequest.md) |  | 

### Return type

[**DecryptDataResponse**](DecryptDataResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EncryptData

> EncryptDataResponse EncryptData(ctx, keyId).EncryptDataRequest(encryptDataRequest).Execute()

Encrypt data



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/kms"
)

func main() {
	keyId := "keyId_example" // string | KMS 키의 고유 ID - [List user keys](/openapi/security/kms/list-user-keys)에서 확인
	encryptDataRequest := *openapiclient.NewEncryptDataRequest(*openapiclient.NewEncryptData("Input_example")) // EncryptDataRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicKeyAPI.EncryptData(context.Background(), keyId).EncryptDataRequest(encryptDataRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicKeyAPI.EncryptData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EncryptData`: EncryptDataResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicKeyAPI.EncryptData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | KMS 키의 고유 ID - [List user keys](/openapi/security/kms/list-user-keys)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiEncryptDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **encryptDataRequest** | [**EncryptDataRequest**](EncryptDataRequest.md) |  | 

### Return type

[**EncryptDataResponse**](EncryptDataResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateHmac

> GenerateHmacResponse GenerateHmac(ctx, keyId).GenerateHmacRequest(generateHmacRequest).Execute()

Generate HMAC



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/kms"
)

func main() {
	keyId := "keyId_example" // string | KMS 키의 고유 ID - [List user keys](/openapi/security/kms/list-user-keys)에서 확인
	generateHmacRequest := *openapiclient.NewGenerateHmacRequest(*openapiclient.NewGenerateHmac("PlainText_example")) // GenerateHmacRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicKeyAPI.GenerateHmac(context.Background(), keyId).GenerateHmacRequest(generateHmacRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicKeyAPI.GenerateHmac``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateHmac`: GenerateHmacResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicKeyAPI.GenerateHmac`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | KMS 키의 고유 ID - [List user keys](/openapi/security/kms/list-user-keys)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateHmacRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **generateHmacRequest** | [**GenerateHmacRequest**](GenerateHmacRequest.md) |  | 

### Return type

[**GenerateHmacResponse**](GenerateHmacResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignData

> SignDataResponse SignData(ctx, keyId).SignDataRequest(signDataRequest).Execute()

Sign data



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/kms"
)

func main() {
	keyId := "keyId_example" // string | KMS 키의 고유 ID - [List user keys](/openapi/security/kms/list-user-keys)에서 확인
	signDataRequest := *openapiclient.NewSignDataRequest(*openapiclient.NewSignData("PlainText_example")) // SignDataRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicKeyAPI.SignData(context.Background(), keyId).SignDataRequest(signDataRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicKeyAPI.SignData``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SignData`: SignDataResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicKeyAPI.SignData`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | KMS 키의 고유 ID - [List user keys](/openapi/security/kms/list-user-keys)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiSignDataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **signDataRequest** | [**SignDataRequest**](SignDataRequest.md) |  | 

### Return type

[**SignDataResponse**](SignDataResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyHmac

> VerifyHmacResponse VerifyHmac(ctx, keyId).VerifyHmacRequest(verifyHmacRequest).Execute()

Verify HMAC



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/kms"
)

func main() {
	keyId := "keyId_example" // string | KMS 키의 고유 ID - [List user keys](/openapi/security/kms/list-user-keys)에서 확인
	verifyHmacRequest := *openapiclient.NewVerifyHmacRequest(*openapiclient.NewVerifyHmac("Input_example", "Output_example")) // VerifyHmacRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicKeyAPI.VerifyHmac(context.Background(), keyId).VerifyHmacRequest(verifyHmacRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicKeyAPI.VerifyHmac``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyHmac`: VerifyHmacResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicKeyAPI.VerifyHmac`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | KMS 키의 고유 ID - [List user keys](/openapi/security/kms/list-user-keys)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifyHmacRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **verifyHmacRequest** | [**VerifyHmacRequest**](VerifyHmacRequest.md) |  | 

### Return type

[**VerifyHmacResponse**](VerifyHmacResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifySignature

> VerifySignatureResponse VerifySignature(ctx, keyId).VerifySignatureRequest(verifySignatureRequest).Execute()

Verify signature



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/kms"
)

func main() {
	keyId := "keyId_example" // string | KMS 키의 고유 ID - [List user keys](/openapi/security/kms/list-user-keys)에서 확인
	verifySignatureRequest := *openapiclient.NewVerifySignatureRequest(*openapiclient.NewVerifySignature("Input_example", "Output_example")) // VerifySignatureRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicKeyAPI.VerifySignature(context.Background(), keyId).VerifySignatureRequest(verifySignatureRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicKeyAPI.VerifySignature``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifySignature`: VerifySignatureResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicKeyAPI.VerifySignature`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keyId** | **string** | KMS 키의 고유 ID - [List user keys](/openapi/security/kms/list-user-keys)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiVerifySignatureRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **verifySignatureRequest** | [**VerifySignatureRequest**](VerifySignatureRequest.md) |  | 

### Return type

[**VerifySignatureResponse**](VerifySignatureResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

