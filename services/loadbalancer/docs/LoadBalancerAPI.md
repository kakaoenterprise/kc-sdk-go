# \LoadBalancerAPI

All URIs are relative to *https://load-balancer.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssociateNewPublicIp**](LoadBalancerAPI.md#AssociateNewPublicIp) | **Post** /api/v1/load-balancers/{load_balancer_id}/public-ips | Associate new public IP
[**AssociatePublicIp**](LoadBalancerAPI.md#AssociatePublicIp) | **Put** /api/v1/load-balancers/{load_balancer_id}/public-ips/{public_ip_id} | Associate public IP
[**CreateLoadBalancer**](LoadBalancerAPI.md#CreateLoadBalancer) | **Post** /api/v1/load-balancers | Create load balancer
[**DeleteLoadBalancer**](LoadBalancerAPI.md#DeleteLoadBalancer) | **Delete** /api/v1/load-balancers/{load_balancer_id} | Delete load balancer
[**GetLoadBalancer**](LoadBalancerAPI.md#GetLoadBalancer) | **Get** /api/v1/load-balancers/{load_balancer_id} | Get load balancer
[**ListLoadBalancers**](LoadBalancerAPI.md#ListLoadBalancers) | **Get** /api/v1/load-balancers | List load balancers
[**RemovePublicIp**](LoadBalancerAPI.md#RemovePublicIp) | **Delete** /api/v1/load-balancers/{load_balancer_id}/public-ips | Remove public IP
[**UpdateAccessLog**](LoadBalancerAPI.md#UpdateAccessLog) | **Patch** /api/v1/load-balancers/{load_balancer_id}/access-log | Update access log
[**UpdateLoadBalancer**](LoadBalancerAPI.md#UpdateLoadBalancer) | **Put** /api/v1/load-balancers/{load_balancer_id} | Update load balancer



## AssociateNewPublicIp

> BnsLoadBalancerV1ApiAssociateNewPublicIpModelResponsePublicIpModel AssociateNewPublicIp(ctx, loadBalancerId).XAuthToken(xAuthToken).Execute()

Associate new public IP



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
	loadBalancerId := "loadBalancerId_example" // string | 퍼블릭 IP를 연결할 대상 로드 밸런서 ID 
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerAPI.AssociateNewPublicIp(context.Background(), loadBalancerId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerAPI.AssociateNewPublicIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssociateNewPublicIp`: BnsLoadBalancerV1ApiAssociateNewPublicIpModelResponsePublicIpModel
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerAPI.AssociateNewPublicIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | 퍼블릭 IP를 연결할 대상 로드 밸런서 ID  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateNewPublicIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**BnsLoadBalancerV1ApiAssociateNewPublicIpModelResponsePublicIpModel**](BnsLoadBalancerV1ApiAssociateNewPublicIpModelResponsePublicIpModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssociatePublicIp

> BnsLoadBalancerV1ApiAssociatePublicIpModelResponsePublicIpModel AssociatePublicIp(ctx, loadBalancerId, publicIpId).XAuthToken(xAuthToken).Execute()

Associate public IP



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
	loadBalancerId := "loadBalancerId_example" // string | 퍼블릭 IP를 연결할 대상 로드 밸런서 ID 
	publicIpId := "publicIpId_example" // string | 연결할 퍼블릭 IP의 고유 ID 
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerAPI.AssociatePublicIp(context.Background(), loadBalancerId, publicIpId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerAPI.AssociatePublicIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssociatePublicIp`: BnsLoadBalancerV1ApiAssociatePublicIpModelResponsePublicIpModel
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerAPI.AssociatePublicIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | 퍼블릭 IP를 연결할 대상 로드 밸런서 ID  | 
**publicIpId** | **string** | 연결할 퍼블릭 IP의 고유 ID  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssociatePublicIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**BnsLoadBalancerV1ApiAssociatePublicIpModelResponsePublicIpModel**](BnsLoadBalancerV1ApiAssociatePublicIpModelResponsePublicIpModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLoadBalancer

> BnsLoadBalancerV1ApiCreateLoadBalancerModelResponseLoadBalancerModel CreateLoadBalancer(ctx).XAuthToken(xAuthToken).BodyCreateLoadBalancer(bodyCreateLoadBalancer).Execute()

Create load balancer



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
	bodyCreateLoadBalancer := *openapiclient.NewBodyCreateLoadBalancer(*openapiclient.NewCreateLoadBalancerModel("Name_example", "SubnetId_example", openapiclient.AvailabilityZone("kr-central-2-a"), "FlavorId_example")) // BodyCreateLoadBalancer | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerAPI.CreateLoadBalancer(context.Background()).XAuthToken(xAuthToken).BodyCreateLoadBalancer(bodyCreateLoadBalancer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerAPI.CreateLoadBalancer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLoadBalancer`: BnsLoadBalancerV1ApiCreateLoadBalancerModelResponseLoadBalancerModel
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerAPI.CreateLoadBalancer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLoadBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyCreateLoadBalancer** | [**BodyCreateLoadBalancer**](BodyCreateLoadBalancer.md) |  | 

### Return type

[**BnsLoadBalancerV1ApiCreateLoadBalancerModelResponseLoadBalancerModel**](BnsLoadBalancerV1ApiCreateLoadBalancerModelResponseLoadBalancerModel.md)

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
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	loadBalancerId := "loadBalancerId_example" // string | 삭제할 로드 밸런서의 ID
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
**loadBalancerId** | **string** | 삭제할 로드 밸런서의 ID | 

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

> BnsLoadBalancerV1ApiGetLoadBalancerModelResponseLoadBalancerModel GetLoadBalancer(ctx, loadBalancerId).XAuthToken(xAuthToken).Execute()

Get load balancer



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
	loadBalancerId := "loadBalancerId_example" // string | 조회할 로드 밸런서의 ID 
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerAPI.GetLoadBalancer(context.Background(), loadBalancerId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerAPI.GetLoadBalancer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLoadBalancer`: BnsLoadBalancerV1ApiGetLoadBalancerModelResponseLoadBalancerModel
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerAPI.GetLoadBalancer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | 조회할 로드 밸런서의 ID  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoadBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**BnsLoadBalancerV1ApiGetLoadBalancerModelResponseLoadBalancerModel**](BnsLoadBalancerV1ApiGetLoadBalancerModelResponseLoadBalancerModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLoadBalancers

> LoadBalancerListModel ListLoadBalancers(ctx).XAuthToken(xAuthToken).Id(id).Name(name).Type_(type_).PrivateVip(privateVip).PublicVip(publicVip).ProvisioningStatus(provisioningStatus).OperatingStatus(operatingStatus).SubnetId(subnetId).SubnetCidrBlock(subnetCidrBlock).VpcId(vpcId).VpcName(vpcName).AvailabilityZone(availabilityZone).BeyondLoadBalancerName(beyondLoadBalancerName).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Limit(limit).Offset(offset).Execute()

List load balancers



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
	id := "id_example" // string | 로드 밸런서 ID  (optional)
	name := "name_example" // string | 로드 밸런서 이름  (optional)
	type_ := "type__example" // string | [로드 밸런서 유형](https://docs.kakaocloud.com/service/bns/lb/lb-overview#features)<br/>- `ALB`: Application Load Balancer<br/>- `NLB`: Network Load Balancer<br/>- `NLB_L4_DSR`: Direct Server Return Network Load Balancer (optional)
	privateVip := "privateVip_example" // string | 내부 VIP 주소  (optional)
	publicVip := "publicVip_example" // string | 외부 VIP 주소  (optional)
	provisioningStatus := openapiclient.ProvisioningStatus("ACTIVE") // ProvisioningStatus | 프로비저닝 상태 <br/> - `ACTIVE`: 활성 <br/> - `DELETED`: 삭제됨 <br/> - `ERROR`: 오류 <br/> - `PENDING_CREATE`: 생성 대기 중 <br/> - `PENDING_UPDATE`: 업데이트 대기 중 <br/> - `PENDING_DELETE`: 삭제 대기 중 (optional)
	operatingStatus := openapiclient.LoadBalancerOperatingStatus("ONLINE") // LoadBalancerOperatingStatus | 운영 상태 <br/> - `ONLINE`: 온라인 <br/> - `DRAINING`: 연결 종료 중 <br/> - `OFFLINE`: 오프라인 <br/> - `DEGRADED`: 성능 저하 <br/> - `ERROR`: 오류 <br/> - `NO_MONITOR`: 모니터링 없음 (optional)
	subnetId := "subnetId_example" // string | 서브넷 ID  (optional)
	subnetCidrBlock := "subnetCidrBlock_example" // string | 서브넷의 IPv4 CIDR 블록  (optional)
	vpcId := "vpcId_example" // string | VPC의 고유 ID (optional)
	vpcName := "vpcName_example" // string | VPC 이름  (optional)
	availabilityZone := openapiclient.AvailabilityZone("kr-central-2-a") // AvailabilityZone | 가용 영역  (optional)
	beyondLoadBalancerName := "beyondLoadBalancerName_example" // string | 연결된 고가용성 그룹 이름  (optional)
	createdAt := "createdAt_example" // string | 리소스가 생성된 시간 <br/> - ISO_8601 형식  <br/> - UTC 기준 (optional)
	updatedAt := "updatedAt_example" // string | 리소스가 마지막으로 수정된 시간 <br/> - ISO_8601 형식  <br/> - UTC 기준 (optional)
	sortKeys := "sortKeys_example" // string | 정렬할 필드를 콤마(,)로 구분   (optional) (default to "created_at")
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`)   (optional) (default to "desc")
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional) (default to 20)
	offset := int32(56) // int32 | 조회 시작 위치 (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerAPI.ListLoadBalancers(context.Background()).XAuthToken(xAuthToken).Id(id).Name(name).Type_(type_).PrivateVip(privateVip).PublicVip(publicVip).ProvisioningStatus(provisioningStatus).OperatingStatus(operatingStatus).SubnetId(subnetId).SubnetCidrBlock(subnetCidrBlock).VpcId(vpcId).VpcName(vpcName).AvailabilityZone(availabilityZone).BeyondLoadBalancerName(beyondLoadBalancerName).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerAPI.ListLoadBalancers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLoadBalancers`: LoadBalancerListModel
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerAPI.ListLoadBalancers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLoadBalancersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **id** | **string** | 로드 밸런서 ID  | 
 **name** | **string** | 로드 밸런서 이름  | 
 **type_** | **string** | [로드 밸런서 유형](https://docs.kakaocloud.com/service/bns/lb/lb-overview#features)&lt;br/&gt;- &#x60;ALB&#x60;: Application Load Balancer&lt;br/&gt;- &#x60;NLB&#x60;: Network Load Balancer&lt;br/&gt;- &#x60;NLB_L4_DSR&#x60;: Direct Server Return Network Load Balancer | 
 **privateVip** | **string** | 내부 VIP 주소  | 
 **publicVip** | **string** | 외부 VIP 주소  | 
 **provisioningStatus** | [**ProvisioningStatus**](ProvisioningStatus.md) | 프로비저닝 상태 &lt;br/&gt; - &#x60;ACTIVE&#x60;: 활성 &lt;br/&gt; - &#x60;DELETED&#x60;: 삭제됨 &lt;br/&gt; - &#x60;ERROR&#x60;: 오류 &lt;br/&gt; - &#x60;PENDING_CREATE&#x60;: 생성 대기 중 &lt;br/&gt; - &#x60;PENDING_UPDATE&#x60;: 업데이트 대기 중 &lt;br/&gt; - &#x60;PENDING_DELETE&#x60;: 삭제 대기 중 | 
 **operatingStatus** | [**LoadBalancerOperatingStatus**](LoadBalancerOperatingStatus.md) | 운영 상태 &lt;br/&gt; - &#x60;ONLINE&#x60;: 온라인 &lt;br/&gt; - &#x60;DRAINING&#x60;: 연결 종료 중 &lt;br/&gt; - &#x60;OFFLINE&#x60;: 오프라인 &lt;br/&gt; - &#x60;DEGRADED&#x60;: 성능 저하 &lt;br/&gt; - &#x60;ERROR&#x60;: 오류 &lt;br/&gt; - &#x60;NO_MONITOR&#x60;: 모니터링 없음 | 
 **subnetId** | **string** | 서브넷 ID  | 
 **subnetCidrBlock** | **string** | 서브넷의 IPv4 CIDR 블록  | 
 **vpcId** | **string** | VPC의 고유 ID | 
 **vpcName** | **string** | VPC 이름  | 
 **availabilityZone** | [**AvailabilityZone**](AvailabilityZone.md) | 가용 영역  | 
 **beyondLoadBalancerName** | **string** | 연결된 고가용성 그룹 이름  | 
 **createdAt** | **string** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
 **updatedAt** | **string** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
 **sortKeys** | **string** | 정렬할 필드를 콤마(,)로 구분   | [default to &quot;created_at&quot;]
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;)   | [default to &quot;desc&quot;]
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | [default to 20]
 **offset** | **int32** | 조회 시작 위치 | [default to 0]

### Return type

[**LoadBalancerListModel**](LoadBalancerListModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemovePublicIp

> BnsLoadBalancerV1ApiRemovePublicIpModelResponsePublicIpModel RemovePublicIp(ctx, loadBalancerId).XAuthToken(xAuthToken).IsDelete(isDelete).Execute()

Remove public IP



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
	loadBalancerId := "loadBalancerId_example" // string | 삭제할 퍼블릭 IP ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	isDelete := true // bool | - `true`로 지정하면 퍼블릭 IP를 함께 삭제함 <br/> - `false`인 경우 연결만 해제하고 IP는 유지됨  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerAPI.RemovePublicIp(context.Background(), loadBalancerId).XAuthToken(xAuthToken).IsDelete(isDelete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerAPI.RemovePublicIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemovePublicIp`: BnsLoadBalancerV1ApiRemovePublicIpModelResponsePublicIpModel
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerAPI.RemovePublicIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | 삭제할 퍼블릭 IP ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemovePublicIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **isDelete** | **bool** | - &#x60;true&#x60;로 지정하면 퍼블릭 IP를 함께 삭제함 &lt;br/&gt; - &#x60;false&#x60;인 경우 연결만 해제하고 IP는 유지됨  | [default to false]

### Return type

[**BnsLoadBalancerV1ApiRemovePublicIpModelResponsePublicIpModel**](BnsLoadBalancerV1ApiRemovePublicIpModelResponsePublicIpModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccessLog

> BnsLoadBalancerV1ApiUpdateAccessLogModelResponseLoadBalancerModel UpdateAccessLog(ctx, loadBalancerId).XAuthToken(xAuthToken).BodyUpdateAccessLog(bodyUpdateAccessLog).Execute()

Update access log



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
	loadBalancerId := "loadBalancerId_example" // string | 액세스 로그 설정을 변경할 로드 밸런서 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bodyUpdateAccessLog := *openapiclient.NewBodyUpdateAccessLog(*openapiclient.NewEditLoadBalancerAccessLogModel("Bucket_example", "AccessKey_example", "SecretKey_example")) // BodyUpdateAccessLog | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerAPI.UpdateAccessLog(context.Background(), loadBalancerId).XAuthToken(xAuthToken).BodyUpdateAccessLog(bodyUpdateAccessLog).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerAPI.UpdateAccessLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAccessLog`: BnsLoadBalancerV1ApiUpdateAccessLogModelResponseLoadBalancerModel
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerAPI.UpdateAccessLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | 액세스 로그 설정을 변경할 로드 밸런서 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccessLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyUpdateAccessLog** | [**BodyUpdateAccessLog**](BodyUpdateAccessLog.md) |  | 

### Return type

[**BnsLoadBalancerV1ApiUpdateAccessLogModelResponseLoadBalancerModel**](BnsLoadBalancerV1ApiUpdateAccessLogModelResponseLoadBalancerModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateLoadBalancer

> BnsLoadBalancerV1ApiUpdateLoadBalancerModelResponseLoadBalancerModel UpdateLoadBalancer(ctx, loadBalancerId).XAuthToken(xAuthToken).BodyUpdateLoadBalancer(bodyUpdateLoadBalancer).Execute()

Update load balancer



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
	loadBalancerId := "loadBalancerId_example" // string | 수정할 로드 밸런서의 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bodyUpdateLoadBalancer := *openapiclient.NewBodyUpdateLoadBalancer(*openapiclient.NewEditLoadBalancerModel()) // BodyUpdateLoadBalancer | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerAPI.UpdateLoadBalancer(context.Background(), loadBalancerId).XAuthToken(xAuthToken).BodyUpdateLoadBalancer(bodyUpdateLoadBalancer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerAPI.UpdateLoadBalancer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLoadBalancer`: BnsLoadBalancerV1ApiUpdateLoadBalancerModelResponseLoadBalancerModel
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerAPI.UpdateLoadBalancer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loadBalancerId** | **string** | 수정할 로드 밸런서의 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLoadBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyUpdateLoadBalancer** | [**BodyUpdateLoadBalancer**](BodyUpdateLoadBalancer.md) |  | 

### Return type

[**BnsLoadBalancerV1ApiUpdateLoadBalancerModelResponseLoadBalancerModel**](BnsLoadBalancerV1ApiUpdateLoadBalancerModelResponseLoadBalancerModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

