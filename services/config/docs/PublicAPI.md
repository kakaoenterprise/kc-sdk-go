# \PublicAPI

All URIs are relative to *https://heimdall.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ResolveAzPolicy**](PublicAPI.md#ResolveAzPolicy) | **Get** /api/v1/az-policies/resolve | 
[**ResolveClientEndpoint**](PublicAPI.md#ResolveClientEndpoint) | **Get** /api/v1/client-endpoints/resolve | 
[**ResolveClientMetadata**](PublicAPI.md#ResolveClientMetadata) | **Get** /api/v1/client-metadata/resolve | 



## ResolveAzPolicy

> AzPolicyResponse ResolveAzPolicy(ctx).XAuthToken(xAuthToken).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/services/config"
)

func main() {
	xAuthToken := "xAuthToken_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.ResolveAzPolicy(context.Background()).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.ResolveAzPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResolveAzPolicy`: AzPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.ResolveAzPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResolveAzPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** |  | 

### Return type

[**AzPolicyResponse**](AzPolicyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResolveClientEndpoint

> ClientEndpointResponse ResolveClientEndpoint(ctx).XAuthToken(xAuthToken).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/services/config"
)

func main() {
	xAuthToken := "xAuthToken_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.ResolveClientEndpoint(context.Background()).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.ResolveClientEndpoint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResolveClientEndpoint`: ClientEndpointResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.ResolveClientEndpoint`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResolveClientEndpointRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** |  | 

### Return type

[**ClientEndpointResponse**](ClientEndpointResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResolveClientMetadata

> ClientMetadataResponse ResolveClientMetadata(ctx).XAuthToken(xAuthToken).Service(service).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/services/config"
)

func main() {
	xAuthToken := "xAuthToken_example" // string | 
	service := "service_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PublicAPI.ResolveClientMetadata(context.Background()).XAuthToken(xAuthToken).Service(service).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PublicAPI.ResolveClientMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResolveClientMetadata`: ClientMetadataResponse
	fmt.Fprintf(os.Stdout, "Response from `PublicAPI.ResolveClientMetadata`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResolveClientMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** |  | 
 **service** | **string** |  | 

### Return type

[**ClientMetadataResponse**](ClientMetadataResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

