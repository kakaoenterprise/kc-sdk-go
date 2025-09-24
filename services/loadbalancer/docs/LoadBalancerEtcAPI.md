# \LoadBalancerEtcAPI

All URIs are relative to *https://load-balancer.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAvailabilityZones**](LoadBalancerEtcAPI.md#ListAvailabilityZones) | **Get** /api/v1/load-balancers/availability-zones | List availability zones
[**ListLoadBalancerTypes**](LoadBalancerEtcAPI.md#ListLoadBalancerTypes) | **Get** /api/v1/load-balancers/flavors | List load balancer types
[**ListTlsCertificates**](LoadBalancerEtcAPI.md#ListTlsCertificates) | **Get** /api/v1/load-balancers/secrets | List TLS certificates



## ListAvailabilityZones

> AvailabilityZoneListModel ListAvailabilityZones(ctx).XAuthToken(xAuthToken).Execute()

List availability zones



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerEtcAPI.ListAvailabilityZones(context.Background()).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerEtcAPI.ListAvailabilityZones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAvailabilityZones`: AvailabilityZoneListModel
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerEtcAPI.ListAvailabilityZones`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAvailabilityZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**AvailabilityZoneListModel**](AvailabilityZoneListModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLoadBalancerTypes

> FlavorListModel ListLoadBalancerTypes(ctx).XAuthToken(xAuthToken).Execute()

List load balancer types



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerEtcAPI.ListLoadBalancerTypes(context.Background()).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerEtcAPI.ListLoadBalancerTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLoadBalancerTypes`: FlavorListModel
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerEtcAPI.ListLoadBalancerTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLoadBalancerTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**FlavorListModel**](FlavorListModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTlsCertificates

> SecretListModel ListTlsCertificates(ctx).XAuthToken(xAuthToken).Name(name).CreatedAt(createdAt).UpdatedAt(updatedAt).Expiration(expiration).Limit(limit).Offset(offset).Execute()

List TLS certificates



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
	name := "name_example" // string | 인증서 이름  (optional)
	createdAt := "createdAt_example" // string | 리소스가 생성된 시간 <br/> - ISO_8601 형식  <br/> - UTC 기준 (optional)
	updatedAt := "updatedAt_example" // string | 리소스가 마지막으로 수정된 시간 <br/> - ISO_8601 형식  <br/> - UTC 기준 (optional)
	expiration := "expiration_example" // string | 만료일 <br/> - ISO_8601 형식  <br/> - UTC 기준  (optional)
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional) (default to 20)
	offset := int32(56) // int32 | 조회 시작 위치 (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerEtcAPI.ListTlsCertificates(context.Background()).XAuthToken(xAuthToken).Name(name).CreatedAt(createdAt).UpdatedAt(updatedAt).Expiration(expiration).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerEtcAPI.ListTlsCertificates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTlsCertificates`: SecretListModel
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerEtcAPI.ListTlsCertificates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTlsCertificatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **name** | **string** | 인증서 이름  | 
 **createdAt** | **string** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
 **updatedAt** | **string** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
 **expiration** | **string** | 만료일 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준  | 
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | [default to 20]
 **offset** | **int32** | 조회 시작 위치 | [default to 0]

### Return type

[**SecretListModel**](SecretListModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

