# \KeypairAPI

All URIs are relative to *https://bcs.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKeypair**](KeypairAPI.md#CreateKeypair) | **Post** /api/v1/keypairs | Create keypair
[**DeleteKeypair**](KeypairAPI.md#DeleteKeypair) | **Delete** /api/v1/keypairs/{keypair_name} | Delete keypair
[**GetKeypair**](KeypairAPI.md#GetKeypair) | **Get** /api/v1/keypairs/{keypair_name} | Get keypair
[**ListKeypairs**](KeypairAPI.md#ListKeypairs) | **Get** /api/v1/keypairs | List keypairs



## CreateKeypair

> CreateKeypairResponse CreateKeypair(ctx).CreateKeypairRequest(createKeypairRequest).Execute()

Create keypair



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/bcs"
)

func main() {
	createKeypairRequest := *openapiclient.NewCreateKeypairRequest(*openapiclient.NewCreateKeypair("Name_example")) // CreateKeypairRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeypairAPI.CreateKeypair(context.Background()).CreateKeypairRequest(createKeypairRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeypairAPI.CreateKeypair``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateKeypair`: CreateKeypairResponse
	fmt.Fprintf(os.Stdout, "Response from `KeypairAPI.CreateKeypair`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateKeypairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createKeypairRequest** | [**CreateKeypairRequest**](CreateKeypairRequest.md) |  | 

### Return type

[**CreateKeypairResponse**](CreateKeypairResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKeypair

> DeleteKeypair(ctx, keypairName).Execute()

Delete keypair



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/bcs"
)

func main() {
	keypairName := "keypairName_example" // string | 삭제할 키 페어의 이름 - [List keypairs](/openapi/bcs/list-keypairs)에서 확인

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KeypairAPI.DeleteKeypair(context.Background(), keypairName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeypairAPI.DeleteKeypair``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keypairName** | **string** | 삭제할 키 페어의 이름 - [List keypairs](/openapi/bcs/list-keypairs)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKeypairRequest struct via the builder pattern


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


## GetKeypair

> GetKeypairResponse GetKeypair(ctx, keypairName).Execute()

Get keypair



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/bcs"
)

func main() {
	keypairName := "keypairName_example" // string | 키 페어 이름

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeypairAPI.GetKeypair(context.Background(), keypairName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeypairAPI.GetKeypair``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKeypair`: GetKeypairResponse
	fmt.Fprintf(os.Stdout, "Response from `KeypairAPI.GetKeypair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keypairName** | **string** | 키 페어 이름 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKeypairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetKeypairResponse**](GetKeypairResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKeypairs

> ListKeypairsResponse ListKeypairs(ctx).Id(id).Name(name).Type_(type_).Fingerprint(fingerprint).CreatedAt(createdAt).SortKeys(sortKeys).SortDirs(sortDirs).Offset(offset).Limit(limit).Execute()

List keypairs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/bcs"
)

func main() {
	id := "id_example" // string | 키 페어의 고유 ID (optional)
	name := "name_example" // string | 키 페어의 이름 (optional)
	type_ := "type__example" // string | 키 페어의 유형 (optional)
	fingerprint := "fingerprint_example" // string | 퍼블릭 키의 핑거프린트 (optional)
	createdAt := "createdAt_example" // string | 리소스가 생성된 시간 - ISO 8601 형식 - UTC 기준 (optional)
	sortKeys := "sortKeys_example" // string | 정렬할 필드를 콤마(,)로 구분 (optional)
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`) (optional)
	offset := int32(56) // int32 | 조회 시작 위치 (optional)
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeypairAPI.ListKeypairs(context.Background()).Id(id).Name(name).Type_(type_).Fingerprint(fingerprint).CreatedAt(createdAt).SortKeys(sortKeys).SortDirs(sortDirs).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeypairAPI.ListKeypairs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListKeypairs`: ListKeypairsResponse
	fmt.Fprintf(os.Stdout, "Response from `KeypairAPI.ListKeypairs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListKeypairsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | 키 페어의 고유 ID | 
 **name** | **string** | 키 페어의 이름 | 
 **type_** | **string** | 키 페어의 유형 | 
 **fingerprint** | **string** | 퍼블릭 키의 핑거프린트 | 
 **createdAt** | **string** | 리소스가 생성된 시간 - ISO 8601 형식 - UTC 기준 | 
 **sortKeys** | **string** | 정렬할 필드를 콤마(,)로 구분 | 
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;) | 
 **offset** | **int32** | 조회 시작 위치 | 
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | 

### Return type

[**ListKeypairsResponse**](ListKeypairsResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

