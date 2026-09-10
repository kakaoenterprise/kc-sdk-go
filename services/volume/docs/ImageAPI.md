# \ImageAPI

All URIs are relative to *https://volume.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateImage**](ImageAPI.md#CreateImage) | **Post** /api/v1/volumes/{volume_id}/image | Create image



## CreateImage

> CreateImageResponse CreateImage(ctx, volumeId).XAuthToken(xAuthToken).CreateImageRequest(createImageRequest).Execute()

Create image



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/volume"
)

func main() {
	volumeId := "volumeId_example" // string | 볼륨의 고유 ID <br/> - [List volumes](https://docs.kakaocloud.com/openapi/bcs/list-volumes)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	createImageRequest := *openapiclient.NewCreateImageRequest(*openapiclient.NewCreateImage("Name_example")) // CreateImageRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImageAPI.CreateImage(context.Background(), volumeId).XAuthToken(xAuthToken).CreateImageRequest(createImageRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImageAPI.CreateImage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateImage`: CreateImageResponse
	fmt.Fprintf(os.Stdout, "Response from `ImageAPI.CreateImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**volumeId** | **string** | 볼륨의 고유 ID &lt;br/&gt; - [List volumes](https://docs.kakaocloud.com/openapi/bcs/list-volumes)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **createImageRequest** | [**CreateImageRequest**](CreateImageRequest.md) |  | 

### Return type

[**CreateImageResponse**](CreateImageResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

