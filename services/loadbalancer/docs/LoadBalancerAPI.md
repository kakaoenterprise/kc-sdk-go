# \LoadBalancerAPI

All URIs are relative to *https://load-balancer.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssociateNewPublicIp**](LoadBalancerAPI.md#AssociateNewPublicIp) | **Post** /api/v1/load-balancers/{load_balancer_id}/public-ips | Associate new public IP
[**AssociatePublicIp**](LoadBalancerAPI.md#AssociatePublicIp) | **Put** /api/v1/load-balancers/{load_balancer_id}/public-ips/{public_ip_id} | Associate public IP
[**CreateLoadBalancer**](LoadBalancerAPI.md#CreateLoadBalancer) | **Post** /api/v1/load-balancers | Create load balancer
[**DeleteLoadBalancer**](LoadBalancerAPI.md#DeleteLoadBalancer) | **Delete** /api/v1/load-balancers/{load_balancer_id} | Delete load balancer
[**GetLoadBalancer**](LoadBalancerAPI.md#GetLoadBalancer) | **Get** /api/v1/load-balancers/{load_balancer_id} | Get load balancer
[**ListAvailabilityZones**](LoadBalancerAPI.md#ListAvailabilityZones) | **Get** /api/v1/load-balancers/availability-zones | List availability zones
[**ListLoadBalancerTypes**](LoadBalancerAPI.md#ListLoadBalancerTypes) | **Get** /api/v1/load-balancers/flavors | List load balancer types
[**ListLoadBalancers**](LoadBalancerAPI.md#ListLoadBalancers) | **Get** /api/v1/load-balancers | List load balancers
[**ListTlsCertificates**](LoadBalancerAPI.md#ListTlsCertificates) | **Get** /api/v1/load-balancers/secrets | List TLS certificates
[**RemovePublicIp**](LoadBalancerAPI.md#RemovePublicIp) | **Delete** /api/v1/load-balancers/{load_balancer_id}/public-ips | Remove public IP
[**UpdateAccessLog**](LoadBalancerAPI.md#UpdateAccessLog) | **Patch** /api/v1/load-balancers/{load_balancer_id}/access-log | Update access log
[**UpdateLoadBalancer**](LoadBalancerAPI.md#UpdateLoadBalancer) | **Put** /api/v1/load-balancers/{load_balancer_id} | Update load balancer



## AssociateNewPublicIp

> AssociateNewPublicIpResponse AssociateNewPublicIp(ctx, loadBalancerId).XAuthToken(xAuthToken).Execute()

Associate new public IP



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/loadbalancer"
)

func main() {
	loadBalancerId := "loadBalancerId_example" // string | 퍼블릭 IP를 연결할 대상 로드 밸런서 ID <br/>- [List load balancers](https://docs.kakaocloud.com/openapi/networking/lb/list-load-balancers)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerAPI.AssociateNewPublicIp(context.Background(), loadBalancerId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerAPI.AssociateNewPublicIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssociateNewPublicIp`: AssociateNewPublicIpResponse
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerAPI.AssociateNewPublicIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | 퍼블릭 IP를 연결할 대상 로드 밸런서 ID &lt;br/&gt;- [List load balancers](https://docs.kakaocloud.com/openapi/networking/lb/list-load-balancers)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateNewPublicIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**AssociateNewPublicIpResponse**](AssociateNewPublicIpResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssociatePublicIp

> AssociatePublicIpResponse AssociatePublicIp(ctx, loadBalancerId, publicIpId).XAuthToken(xAuthToken).Execute()

Associate public IP



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/loadbalancer"
)

func main() {
	loadBalancerId := "loadBalancerId_example" // string | 퍼블릭 IP를 연결할 대상 로드 밸런서 ID <br/>- [List load balancers](https://docs.kakaocloud.com/openapi/networking/lb/list-load-balancers)에서 확인
	publicIpId := "publicIpId_example" // string | 연결할 퍼블릭 IP의 고유 ID <br/>- [List public IPs](https://docs.kakaocloud.com/openapi/networking/vpc/list-public-ips)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerAPI.AssociatePublicIp(context.Background(), loadBalancerId, publicIpId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerAPI.AssociatePublicIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssociatePublicIp`: AssociatePublicIpResponse
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerAPI.AssociatePublicIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | 퍼블릭 IP를 연결할 대상 로드 밸런서 ID &lt;br/&gt;- [List load balancers](https://docs.kakaocloud.com/openapi/networking/lb/list-load-balancers)에서 확인 | 
**publicIpId** | **string** | 연결할 퍼블릭 IP의 고유 ID &lt;br/&gt;- [List public IPs](https://docs.kakaocloud.com/openapi/networking/vpc/list-public-ips)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssociatePublicIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**AssociatePublicIpResponse**](AssociatePublicIpResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLoadBalancer

> CreateLoadBalancerResponse CreateLoadBalancer(ctx).XAuthToken(xAuthToken).CreateLoadBalancerRequest(createLoadBalancerRequest).Execute()

Create load balancer



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/loadbalancer"
)

func main() {
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	createLoadBalancerRequest := *openapiclient.NewCreateLoadBalancerRequest(*openapiclient.NewCreateLoadBalancer("Name_example", "SubnetId_example", openapiclient.AvailabilityZone("kr-central-2-a"), "FlavorId_example")) // CreateLoadBalancerRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerAPI.CreateLoadBalancer(context.Background()).XAuthToken(xAuthToken).CreateLoadBalancerRequest(createLoadBalancerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerAPI.CreateLoadBalancer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLoadBalancer`: CreateLoadBalancerResponse
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerAPI.CreateLoadBalancer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLoadBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **createLoadBalancerRequest** | [**CreateLoadBalancerRequest**](CreateLoadBalancerRequest.md) |  | 

### Return type

[**CreateLoadBalancerResponse**](CreateLoadBalancerResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLoadBalancer

> DeleteLoadBalancer(ctx, loadBalancerId).XAuthToken(xAuthToken).Execute()

Delete load balancer



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/loadbalancer"
)

func main() {
	loadBalancerId := "loadBalancerId_example" // string | 삭제할 로드 밸런서의 ID <br/>- [List load balancers](https://docs.kakaocloud.com/openapi/networking/lb/list-load-balancers)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LoadBalancerAPI.DeleteLoadBalancer(context.Background(), loadBalancerId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerAPI.DeleteLoadBalancer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | 삭제할 로드 밸런서의 ID &lt;br/&gt;- [List load balancers](https://docs.kakaocloud.com/openapi/networking/lb/list-load-balancers)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLoadBalancerRequest struct via the builder pattern


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


## GetLoadBalancer

> GetLoadBalancerResponse GetLoadBalancer(ctx, loadBalancerId).XAuthToken(xAuthToken).Execute()

Get load balancer



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/loadbalancer"
)

func main() {
	loadBalancerId := "loadBalancerId_example" // string | 조회할 로드 밸런서의 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerAPI.GetLoadBalancer(context.Background(), loadBalancerId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerAPI.GetLoadBalancer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoadBalancer`: GetLoadBalancerResponse
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerAPI.GetLoadBalancer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | 조회할 로드 밸런서의 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoadBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**GetLoadBalancerResponse**](GetLoadBalancerResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAvailabilityZones

> ListAvailabilityZonesResponse ListAvailabilityZones(ctx).XAuthToken(xAuthToken).Execute()

List availability zones



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/loadbalancer"
)

func main() {
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerAPI.ListAvailabilityZones(context.Background()).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerAPI.ListAvailabilityZones``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAvailabilityZones`: ListAvailabilityZonesResponse
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerAPI.ListAvailabilityZones`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAvailabilityZonesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**ListAvailabilityZonesResponse**](ListAvailabilityZonesResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLoadBalancerTypes

> ListLoadBalancerTypesResponse ListLoadBalancerTypes(ctx).XAuthToken(xAuthToken).Execute()

List load balancer types



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/loadbalancer"
)

func main() {
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerAPI.ListLoadBalancerTypes(context.Background()).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerAPI.ListLoadBalancerTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLoadBalancerTypes`: ListLoadBalancerTypesResponse
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerAPI.ListLoadBalancerTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLoadBalancerTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**ListLoadBalancerTypesResponse**](ListLoadBalancerTypesResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLoadBalancers

> ListLoadBalancersResponse ListLoadBalancers(ctx).XAuthToken(xAuthToken).Id(id).Name(name).Type_(type_).PrivateVip(privateVip).PublicVip(publicVip).ProvisioningStatus(provisioningStatus).OperatingStatus(operatingStatus).SubnetId(subnetId).SubnetCidrBlock(subnetCidrBlock).VpcName(vpcName).VpcId(vpcId).AvailabilityZone(availabilityZone).BeyondLoadBalancerName(beyondLoadBalancerName).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Offset(offset).Limit(limit).Execute()

List load balancers



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/loadbalancer"
)

func main() {
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	id := "id_example" // string | 로드 밸런서 ID (optional)
	name := "name_example" // string | 로드 밸런서 이름 (optional)
	type_ := openapiclient.LoadBalancerType("ALB") // LoadBalancerType | 로드 밸런서 유형 (optional)
	privateVip := "privateVip_example" // string | 내부 VIP 주소 (optional)
	publicVip := "publicVip_example" // string | 외부 VIP 주소 (optional)
	provisioningStatus := openapiclient.ProvisioningStatus("ACTIVE") // ProvisioningStatus | 프로비저닝 상태 (optional)
	operatingStatus := openapiclient.LoadBalancerOperatingStatus("ONLINE") // LoadBalancerOperatingStatus | 운영 상태 (optional)
	subnetId := "subnetId_example" // string | 서브넷 ID (optional)
	subnetCidrBlock := "subnetCidrBlock_example" // string | 서브넷의 IPv4 CIDR 블록 (optional)
	vpcName := "vpcName_example" // string | VPC 이름 (optional)
	vpcId := "vpcId_example" // string | VPC의 고유 ID (optional)
	availabilityZone := openapiclient.AvailabilityZone("kr-central-2-a") // AvailabilityZone | 가용 영역 (optional)
	beyondLoadBalancerName := "beyondLoadBalancerName_example" // string | 연결된 고가용성 그룹 이름 (optional)
	createdAt := "createdAt_example" // string | 리소스가 생성된 시간 <br/> - ISO 8601 형식 <br/> - UTC 기준 (optional)
	updatedAt := "updatedAt_example" // string | 리소스가 마지막으로 수정된 시간 <br/> - ISO 8601 형식 <br/> - UTC 기준 (optional)
	sortKeys := "sortKeys_example" // string | 정렬할 필드를 콤마(,)로 구분   (optional)
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`)  (optional)
	offset := int32(56) // int32 | 조회 시작 위치 (optional)
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerAPI.ListLoadBalancers(context.Background()).XAuthToken(xAuthToken).Id(id).Name(name).Type_(type_).PrivateVip(privateVip).PublicVip(publicVip).ProvisioningStatus(provisioningStatus).OperatingStatus(operatingStatus).SubnetId(subnetId).SubnetCidrBlock(subnetCidrBlock).VpcName(vpcName).VpcId(vpcId).AvailabilityZone(availabilityZone).BeyondLoadBalancerName(beyondLoadBalancerName).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerAPI.ListLoadBalancers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLoadBalancers`: ListLoadBalancersResponse
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerAPI.ListLoadBalancers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLoadBalancersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **id** | **string** | 로드 밸런서 ID | 
 **name** | **string** | 로드 밸런서 이름 | 
 **type_** | [**LoadBalancerType**](LoadBalancerType.md) | 로드 밸런서 유형 | 
 **privateVip** | **string** | 내부 VIP 주소 | 
 **publicVip** | **string** | 외부 VIP 주소 | 
 **provisioningStatus** | [**ProvisioningStatus**](ProvisioningStatus.md) | 프로비저닝 상태 | 
 **operatingStatus** | [**LoadBalancerOperatingStatus**](LoadBalancerOperatingStatus.md) | 운영 상태 | 
 **subnetId** | **string** | 서브넷 ID | 
 **subnetCidrBlock** | **string** | 서브넷의 IPv4 CIDR 블록 | 
 **vpcName** | **string** | VPC 이름 | 
 **vpcId** | **string** | VPC의 고유 ID | 
 **availabilityZone** | [**AvailabilityZone**](AvailabilityZone.md) | 가용 영역 | 
 **beyondLoadBalancerName** | **string** | 연결된 고가용성 그룹 이름 | 
 **createdAt** | **string** | 리소스가 생성된 시간 &lt;br/&gt; - ISO 8601 형식 &lt;br/&gt; - UTC 기준 | 
 **updatedAt** | **string** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO 8601 형식 &lt;br/&gt; - UTC 기준 | 
 **sortKeys** | **string** | 정렬할 필드를 콤마(,)로 구분   | 
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;)  | 
 **offset** | **int32** | 조회 시작 위치 | 
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | 

### Return type

[**ListLoadBalancersResponse**](ListLoadBalancersResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTlsCertificates

> ListTlsCertificatesResponse ListTlsCertificates(ctx).XAuthToken(xAuthToken).Offset(offset).Limit(limit).Name(name).CreatedAt(createdAt).UpdatedAt(updatedAt).Expiration(expiration).Execute()

List TLS certificates



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/loadbalancer"
)

func main() {
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	offset := int32(56) // int32 | 조회 시작 위치 (optional)
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional)
	name := "name_example" // string | 인증서 이름 (optional)
	createdAt := "createdAt_example" // string | 리소스가 생성된 시간 <br/> - ISO 8601 형식 <br/> - UTC 기준 (optional)
	updatedAt := "updatedAt_example" // string | 리소스가 마지막으로 수정된 시간 <br/> - ISO 8601 형식 <br/> - UTC 기준 (optional)
	expiration := "expiration_example" // string | 만료일 <br/> - ISO 8601 형식 <br/> - UTC 기준 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerAPI.ListTlsCertificates(context.Background()).XAuthToken(xAuthToken).Offset(offset).Limit(limit).Name(name).CreatedAt(createdAt).UpdatedAt(updatedAt).Expiration(expiration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerAPI.ListTlsCertificates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTlsCertificates`: ListTlsCertificatesResponse
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerAPI.ListTlsCertificates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTlsCertificatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **offset** | **int32** | 조회 시작 위치 | 
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | 
 **name** | **string** | 인증서 이름 | 
 **createdAt** | **string** | 리소스가 생성된 시간 &lt;br/&gt; - ISO 8601 형식 &lt;br/&gt; - UTC 기준 | 
 **updatedAt** | **string** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO 8601 형식 &lt;br/&gt; - UTC 기준 | 
 **expiration** | **string** | 만료일 &lt;br/&gt; - ISO 8601 형식 &lt;br/&gt; - UTC 기준 | 

### Return type

[**ListTlsCertificatesResponse**](ListTlsCertificatesResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemovePublicIp

> RemovePublicIpResponse RemovePublicIp(ctx, loadBalancerId).XAuthToken(xAuthToken).IsDelete(isDelete).Execute()

Remove public IP



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/loadbalancer"
)

func main() {
	loadBalancerId := "loadBalancerId_example" // string | 퍼블릭 IP를 제거할 대상 로드 밸런서 ID <br/>- [List load balancers](https://docs.kakaocloud.com/openapi/networking/lb/list-load-balancers)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	isDelete := true // bool | - `true`로 지정하면 퍼블릭 IP를 함께 삭제함 <br/> - `false`인 경우 연결만 해제하고 IP는 유지됨 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerAPI.RemovePublicIp(context.Background(), loadBalancerId).XAuthToken(xAuthToken).IsDelete(isDelete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerAPI.RemovePublicIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemovePublicIp`: RemovePublicIpResponse
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerAPI.RemovePublicIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | 퍼블릭 IP를 제거할 대상 로드 밸런서 ID &lt;br/&gt;- [List load balancers](https://docs.kakaocloud.com/openapi/networking/lb/list-load-balancers)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemovePublicIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **isDelete** | **bool** | - &#x60;true&#x60;로 지정하면 퍼블릭 IP를 함께 삭제함 &lt;br/&gt; - &#x60;false&#x60;인 경우 연결만 해제하고 IP는 유지됨 | 

### Return type

[**RemovePublicIpResponse**](RemovePublicIpResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccessLog

> UpdateAccessLogResponse UpdateAccessLog(ctx, loadBalancerId).XAuthToken(xAuthToken).UpdateAccessLogRequest(updateAccessLogRequest).Execute()

Update access log



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/loadbalancer"
)

func main() {
	loadBalancerId := "loadBalancerId_example" // string | 액세스 로그 설정을 변경할 로드 밸런서 ID <br/>- [List load balancers](https://docs.kakaocloud.com/openapi/networking/lb/list-load-balancers)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	updateAccessLogRequest := *openapiclient.NewUpdateAccessLogRequest() // UpdateAccessLogRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerAPI.UpdateAccessLog(context.Background(), loadBalancerId).XAuthToken(xAuthToken).UpdateAccessLogRequest(updateAccessLogRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerAPI.UpdateAccessLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAccessLog`: UpdateAccessLogResponse
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerAPI.UpdateAccessLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | 액세스 로그 설정을 변경할 로드 밸런서 ID &lt;br/&gt;- [List load balancers](https://docs.kakaocloud.com/openapi/networking/lb/list-load-balancers)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccessLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **updateAccessLogRequest** | [**UpdateAccessLogRequest**](UpdateAccessLogRequest.md) |  | 

### Return type

[**UpdateAccessLogResponse**](UpdateAccessLogResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLoadBalancer

> UpdateLoadBalancerResponse UpdateLoadBalancer(ctx, loadBalancerId).XAuthToken(xAuthToken).UpdateLoadBalancerRequest(updateLoadBalancerRequest).Execute()

Update load balancer



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/loadbalancer"
)

func main() {
	loadBalancerId := "loadBalancerId_example" // string | 수정할 로드 밸런서의 ID <br/>- [List load balancers](https://docs.kakaocloud.com/openapi/networking/lb/list-load-balancers)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	updateLoadBalancerRequest := *openapiclient.NewUpdateLoadBalancerRequest(*openapiclient.NewUpdateLoadBalancer()) // UpdateLoadBalancerRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerAPI.UpdateLoadBalancer(context.Background(), loadBalancerId).XAuthToken(xAuthToken).UpdateLoadBalancerRequest(updateLoadBalancerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerAPI.UpdateLoadBalancer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLoadBalancer`: UpdateLoadBalancerResponse
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerAPI.UpdateLoadBalancer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | 수정할 로드 밸런서의 ID &lt;br/&gt;- [List load balancers](https://docs.kakaocloud.com/openapi/networking/lb/list-load-balancers)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLoadBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **updateLoadBalancerRequest** | [**UpdateLoadBalancerRequest**](UpdateLoadBalancerRequest.md) |  | 

### Return type

[**UpdateLoadBalancerResponse**](UpdateLoadBalancerResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

