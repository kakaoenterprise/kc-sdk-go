# \KeypairAPI

All URIs are relative to *https://bcs.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateKeypair**](KeypairAPI.md#CreateKeypair) | **Post** /api/v1/keypairs | Create keypair
[**DeleteKeypair**](KeypairAPI.md#DeleteKeypair) | **Delete** /api/v1/keypairs/{keypair_name} | Delete keypair
[**GetKeypair**](KeypairAPI.md#GetKeypair) | **Get** /api/v1/keypairs/{keypair_name} | Get keypair
[**ListKeypairs**](KeypairAPI.md#ListKeypairs) | **Get** /api/v1/keypairs | List keypairs



## CreateKeypair

> BcsInstanceV1ApiCreateKeypairModelResponseKeypairModel CreateKeypair(ctx).XAuthToken(xAuthToken).BodyCreateKeypair(bodyCreateKeypair).Execute()

Create keypair



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
	bodyCreateKeypair := *openapiclient.NewBodyCreateKeypair(*openapiclient.NewCreateKeypairModel("Name_example")) // BodyCreateKeypair | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeypairAPI.CreateKeypair(context.Background()).XAuthToken(xAuthToken).BodyCreateKeypair(bodyCreateKeypair).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeypairAPI.CreateKeypair``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateKeypair`: BcsInstanceV1ApiCreateKeypairModelResponseKeypairModel
	fmt.Fprintf(os.Stdout, "Response from `KeypairAPI.CreateKeypair`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateKeypairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyCreateKeypair** | [**BodyCreateKeypair**](BodyCreateKeypair.md) |  | 

### Return type

[**BcsInstanceV1ApiCreateKeypairModelResponseKeypairModel**](BcsInstanceV1ApiCreateKeypairModelResponseKeypairModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKeypair

> DeleteKeypair(ctx, keypairName).XAuthToken(xAuthToken).Execute()

Delete keypair



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
	keypairName := "keypairName_example" // string | 삭제할 키 페어의 이름
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.KeypairAPI.DeleteKeypair(context.Background(), keypairName).XAuthToken(xAuthToken).Execute()
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
**keypairName** | **string** | 삭제할 키 페어의 이름 | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKeypairRequest struct via the builder pattern


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


## GetKeypair

> BcsInstanceV1ApiGetKeypairModelResponseKeypairModel GetKeypair(ctx, keypairName).XAuthToken(xAuthToken).Execute()

Get keypair



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
	keypairName := "keypairName_example" // string | 조회할 키 페어의 이름
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeypairAPI.GetKeypair(context.Background(), keypairName).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeypairAPI.GetKeypair``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKeypair`: BcsInstanceV1ApiGetKeypairModelResponseKeypairModel
	fmt.Fprintf(os.Stdout, "Response from `KeypairAPI.GetKeypair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**keypairName** | **string** | 조회할 키 페어의 이름 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKeypairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**BcsInstanceV1ApiGetKeypairModelResponseKeypairModel**](BcsInstanceV1ApiGetKeypairModelResponseKeypairModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListKeypairs

> KeypairListModel ListKeypairs(ctx).XAuthToken(xAuthToken).Id(id).Name(name).Type_(type_).Fingerprint(fingerprint).CreatedAt(createdAt).SortKeys(sortKeys).SortDirs(sortDirs).Limit(limit).Offset(offset).Execute()

List keypairs



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
	id := "id_example" // string | 키 페어의 고유 ID (optional)
	name := "name_example" // string | 키 페어의 이름 (optional)
	type_ := "type__example" // string | 키 페어의 유형 <br/> - 예시: ssh, x509 등 (optional)
	fingerprint := "fingerprint_example" // string | 퍼블릭 키의 핑거프린트 (optional)
	createdAt := "createdAt_example" // string | 리소스가 생성된 시간 <br/> - ISO_8601 형식  <br/> - UTC 기준 (optional)
	sortKeys := "sortKeys_example" // string | 정렬할 필드를 콤마(,)로 구분  (optional) (default to "created_at")
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`)   (optional) (default to "desc")
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional) (default to 20)
	offset := int32(56) // int32 | 조회 시작 위치 (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KeypairAPI.ListKeypairs(context.Background()).XAuthToken(xAuthToken).Id(id).Name(name).Type_(type_).Fingerprint(fingerprint).CreatedAt(createdAt).SortKeys(sortKeys).SortDirs(sortDirs).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KeypairAPI.ListKeypairs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListKeypairs`: KeypairListModel
	fmt.Fprintf(os.Stdout, "Response from `KeypairAPI.ListKeypairs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListKeypairsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **id** | **string** | 키 페어의 고유 ID | 
 **name** | **string** | 키 페어의 이름 | 
 **type_** | **string** | 키 페어의 유형 &lt;br/&gt; - 예시: ssh, x509 등 | 
 **fingerprint** | **string** | 퍼블릭 키의 핑거프린트 | 
 **createdAt** | **string** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
 **sortKeys** | **string** | 정렬할 필드를 콤마(,)로 구분  | [default to &quot;created_at&quot;]
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;)   | [default to &quot;desc&quot;]
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | [default to 20]
 **offset** | **int32** | 조회 시작 위치 | [default to 0]

### Return type

[**KeypairListModel**](KeypairListModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

