# \KMSAPI

All URIs are relative to *https://key-management-service.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateKmsService**](KMSAPI.md#ActivateKmsService) | **Post** /api/v1/activation | Activate KMS service 
[**GetKmsServiceStatus**](KMSAPI.md#GetKmsServiceStatus) | **Get** /api/v1/activation | Get KMS service status 



## ActivateKmsService

> ActivateKmsServiceResponse ActivateKmsService(ctx).Execute()

Activate KMS service 



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KMSAPI.ActivateKmsService(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KMSAPI.ActivateKmsService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActivateKmsService`: ActivateKmsServiceResponse
	fmt.Fprintf(os.Stdout, "Response from `KMSAPI.ActivateKmsService`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiActivateKmsServiceRequest struct via the builder pattern


### Return type

[**ActivateKmsServiceResponse**](ActivateKmsServiceResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKmsServiceStatus

> GetKmsServiceStatusResponse GetKmsServiceStatus(ctx).Execute()

Get KMS service status 



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.KMSAPI.GetKmsServiceStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `KMSAPI.GetKmsServiceStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKmsServiceStatus`: GetKmsServiceStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `KMSAPI.GetKmsServiceStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetKmsServiceStatusRequest struct via the builder pattern


### Return type

[**GetKmsServiceStatusResponse**](GetKmsServiceStatusResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

