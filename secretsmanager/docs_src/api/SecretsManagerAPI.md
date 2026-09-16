# \SecretsManagerAPI

All URIs are relative to *https://secrets-manager-service.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateSecretsManagerService**](SecretsManagerAPI.md#ActivateSecretsManagerService) | **Post** /api/v1/activation | Activate Secrets Manager service 
[**GetSecretsManagerServiceStatus**](SecretsManagerAPI.md#GetSecretsManagerServiceStatus) | **Get** /api/v1/activation | Get Secrets Manager service status 



## ActivateSecretsManagerService

> ActivateSecretsManagerServiceResponse ActivateSecretsManagerService(ctx).Execute()

Activate Secrets Manager service 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/secretsmanager"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretsManagerAPI.ActivateSecretsManagerService(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsManagerAPI.ActivateSecretsManagerService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActivateSecretsManagerService`: ActivateSecretsManagerServiceResponse
	fmt.Fprintf(os.Stdout, "Response from `SecretsManagerAPI.ActivateSecretsManagerService`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiActivateSecretsManagerServiceRequest struct via the builder pattern


### Return type

[**ActivateSecretsManagerServiceResponse**](ActivateSecretsManagerServiceResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecretsManagerServiceStatus

> GetSecretsManagerServiceStatusResponse GetSecretsManagerServiceStatus(ctx).Execute()

Get Secrets Manager service status 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/secretsmanager"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecretsManagerAPI.GetSecretsManagerServiceStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecretsManagerAPI.GetSecretsManagerServiceStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecretsManagerServiceStatus`: GetSecretsManagerServiceStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `SecretsManagerAPI.GetSecretsManagerServiceStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecretsManagerServiceStatusRequest struct via the builder pattern


### Return type

[**GetSecretsManagerServiceStatusResponse**](GetSecretsManagerServiceStatusResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

