# \BeyondLoadBalancerAPI

All URIs are relative to *https://load-balancer.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateHaGroup**](BeyondLoadBalancerAPI.md#CreateHaGroup) | **Post** /api/v1/beyond-load-balancers | Create HA group
[**DeleteHaGroup**](BeyondLoadBalancerAPI.md#DeleteHaGroup) | **Delete** /api/v1/beyond-load-balancers/{beyond_load_balancer_id} | Delete HA group
[**DetachHaGroupLoadBalancer**](BeyondLoadBalancerAPI.md#DetachHaGroupLoadBalancer) | **Delete** /api/v1/beyond-load-balancers/{beyond_load_balancer_id}/subnets/{load_balancer_id} | Detach HA group load balancer
[**GetHaGroup**](BeyondLoadBalancerAPI.md#GetHaGroup) | **Get** /api/v1/beyond-load-balancers/{beyond_load_balancer_id} | Get HA group
[**ListHaGroups**](BeyondLoadBalancerAPI.md#ListHaGroups) | **Get** /api/v1/beyond-load-balancers | List HA groups
[**UpdateHaGroup**](BeyondLoadBalancerAPI.md#UpdateHaGroup) | **Put** /api/v1/beyond-load-balancers/{beyond_load_balancer_id} | Update HA group
[**UpdateHaGroupLoadBalancer**](BeyondLoadBalancerAPI.md#UpdateHaGroupLoadBalancer) | **Post** /api/v1/beyond-load-balancers/{beyond_load_balancer_id}/subnets | Update HA group load balancer



## CreateHaGroup

> BnsLoadBalancerV1ApiCreateHaGroupModelResponseBeyondLoadBalancerModel CreateHaGroup(ctx).XAuthToken(xAuthToken).BodyCreateHaGroup(bodyCreateHaGroup).Execute()

Create HA group



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
	bodyCreateHaGroup := *openapiclient.NewBodyCreateHaGroup(*openapiclient.NewBnsLoadBalancerV1ApiCreateHaGroupModelCreateBeyondLoadBalancerModel("Name_example", "TypeId_example", openapiclient.Scheme("internet-facing"), "VpcId_example", []openapiclient.BnsLoadBalancerV1ApiCreateHaGroupModelSubnetModel{*openapiclient.NewBnsLoadBalancerV1ApiCreateHaGroupModelSubnetModel("LoadBalancerId_example", openapiclient.AvailabilityZone("kr-central-2-a"))})) // BodyCreateHaGroup | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BeyondLoadBalancerAPI.CreateHaGroup(context.Background()).XAuthToken(xAuthToken).BodyCreateHaGroup(bodyCreateHaGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BeyondLoadBalancerAPI.CreateHaGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateHaGroup`: BnsLoadBalancerV1ApiCreateHaGroupModelResponseBeyondLoadBalancerModel
	fmt.Fprintf(os.Stdout, "Response from `BeyondLoadBalancerAPI.CreateHaGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateHaGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyCreateHaGroup** | [**BodyCreateHaGroup**](BodyCreateHaGroup.md) |  | 

### Return type

[**BnsLoadBalancerV1ApiCreateHaGroupModelResponseBeyondLoadBalancerModel**](BnsLoadBalancerV1ApiCreateHaGroupModelResponseBeyondLoadBalancerModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteHaGroup

> DeleteHaGroup(ctx, beyondLoadBalancerId).XAuthToken(xAuthToken).Execute()

Delete HA group



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
	beyondLoadBalancerId := "beyondLoadBalancerId_example" // string | 삭제할 고가용성 그룹의 ID 
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BeyondLoadBalancerAPI.DeleteHaGroup(context.Background(), beyondLoadBalancerId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BeyondLoadBalancerAPI.DeleteHaGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**beyondLoadBalancerId** | **string** | 삭제할 고가용성 그룹의 ID  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHaGroupRequest struct via the builder pattern


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


## DetachHaGroupLoadBalancer

> DetachHaGroupLoadBalancer(ctx, beyondLoadBalancerId, loadBalancerId).XAuthToken(xAuthToken).Execute()

Detach HA group load balancer



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
	beyondLoadBalancerId := "beyondLoadBalancerId_example" // string | 고가용성 그룹의 ID 
	loadBalancerId := "loadBalancerId_example" // string | 연결 해제할 로드 밸런서의 ID 
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BeyondLoadBalancerAPI.DetachHaGroupLoadBalancer(context.Background(), beyondLoadBalancerId, loadBalancerId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BeyondLoadBalancerAPI.DetachHaGroupLoadBalancer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**beyondLoadBalancerId** | **string** | 고가용성 그룹의 ID  | 
**loadBalancerId** | **string** | 연결 해제할 로드 밸런서의 ID  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachHaGroupLoadBalancerRequest struct via the builder pattern


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


## GetHaGroup

> BnsLoadBalancerV1ApiGetHaGroupModelResponseBeyondLoadBalancerModel GetHaGroup(ctx, beyondLoadBalancerId).XAuthToken(xAuthToken).Execute()

Get HA group



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
	beyondLoadBalancerId := "beyondLoadBalancerId_example" // string | 조회할 고가용성 그룹의 ID 
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BeyondLoadBalancerAPI.GetHaGroup(context.Background(), beyondLoadBalancerId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BeyondLoadBalancerAPI.GetHaGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetHaGroup`: BnsLoadBalancerV1ApiGetHaGroupModelResponseBeyondLoadBalancerModel
	fmt.Fprintf(os.Stdout, "Response from `BeyondLoadBalancerAPI.GetHaGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**beyondLoadBalancerId** | **string** | 조회할 고가용성 그룹의 ID  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHaGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**BnsLoadBalancerV1ApiGetHaGroupModelResponseBeyondLoadBalancerModel**](BnsLoadBalancerV1ApiGetHaGroupModelResponseBeyondLoadBalancerModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListHaGroups

> BeyondLoadBalancerListModel ListHaGroups(ctx).XAuthToken(xAuthToken).Id(id).Name(name).DnsName(dnsName).Scheme(scheme).ProvisioningStatus(provisioningStatus).OperatingStatus(operatingStatus).Type_(type_).VpcName(vpcName).VpcId(vpcId).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Limit(limit).Offset(offset).Execute()

List HA groups



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
	id := "id_example" // string | 조회할 고가용성 그룹 ID (optional)
	name := "name_example" // string | 고가용성 그룹 이름  (optional)
	dnsName := "dnsName_example" // string | DNS 이름 (optional)
	scheme := openapiclient.Scheme("internet-facing") // Scheme | 접근 방식 <br/> - `internet-facing`: 인터넷과 연결됨 <br/> - `internal`: 내부 전용  (optional)
	provisioningStatus := openapiclient.ProvisioningStatus("ACTIVE") // ProvisioningStatus | 프로비저닝 상태 <br/> - `ACTIVE`: 활성 상태 <br/> - `DELETED`: 삭제됨 <br/> - `ERROR`: 오류 발생 <br/> - `PENDING_CREATE`: 생성 대기 중 <br/> - `PENDING_UPDATE`: 업데이트 대기 중 <br/> - `PENDING_DELETE`: 삭제 대기 중 (optional)
	operatingStatus := openapiclient.LoadBalancerOperatingStatus("ONLINE") // LoadBalancerOperatingStatus | 운영 상태 <br/> - `ONLINE`: 온라인 상태 <br/> - `DRAINING`: 연결 종료 중 <br/> - `OFFLINE`: 오프라인 상태 <br/> - `DEGRADED`: 성능 저하 상태 <br/> - `ERROR`: 오류 발생 <br/> - `NO_MONITOR`: 모니터링 없음 (optional)
	type_ := openapiclient.LoadBalancerType("ALB") // LoadBalancerType | [로드 밸런서 유형](https://docs.kakaocloud.com/service/bns/lb/lb-overview#사용-목적에-따른-로드-밸런서-유형-제공) <br/> - `ALB`: Application Load Balancer <br/> - `NLB`: Network Load Balancer <br/> - `NLB_L4_DSR`: Network Load Balancer (L4 Direct Server Return) (optional)
	vpcName := "vpcName_example" // string | 연결된 VPC 이름  (optional)
	vpcId := "vpcId_example" // string | 연결된 VPC ID (optional)
	createdAt := "createdAt_example" // string | 리소스가 생성된 시간 <br/> - ISO_8601 형식  <br/> - UTC 기준 (optional)
	updatedAt := "updatedAt_example" // string | 리소스가 마지막으로 수정된 시간 <br/> - ISO_8601 형식  <br/> - UTC 기준 (optional)
	sortKeys := "sortKeys_example" // string | 정렬할 필드를 콤마(,)로 구분   (optional) (default to "created_at")
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`)   (optional) (default to "desc")
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional) (default to 20)
	offset := int32(56) // int32 | 조회 시작 위치 (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BeyondLoadBalancerAPI.ListHaGroups(context.Background()).XAuthToken(xAuthToken).Id(id).Name(name).DnsName(dnsName).Scheme(scheme).ProvisioningStatus(provisioningStatus).OperatingStatus(operatingStatus).Type_(type_).VpcName(vpcName).VpcId(vpcId).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BeyondLoadBalancerAPI.ListHaGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListHaGroups`: BeyondLoadBalancerListModel
	fmt.Fprintf(os.Stdout, "Response from `BeyondLoadBalancerAPI.ListHaGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListHaGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **id** | **string** | 조회할 고가용성 그룹 ID | 
 **name** | **string** | 고가용성 그룹 이름  | 
 **dnsName** | **string** | DNS 이름 | 
 **scheme** | [**Scheme**](Scheme.md) | 접근 방식 &lt;br/&gt; - &#x60;internet-facing&#x60;: 인터넷과 연결됨 &lt;br/&gt; - &#x60;internal&#x60;: 내부 전용  | 
 **provisioningStatus** | [**ProvisioningStatus**](ProvisioningStatus.md) | 프로비저닝 상태 &lt;br/&gt; - &#x60;ACTIVE&#x60;: 활성 상태 &lt;br/&gt; - &#x60;DELETED&#x60;: 삭제됨 &lt;br/&gt; - &#x60;ERROR&#x60;: 오류 발생 &lt;br/&gt; - &#x60;PENDING_CREATE&#x60;: 생성 대기 중 &lt;br/&gt; - &#x60;PENDING_UPDATE&#x60;: 업데이트 대기 중 &lt;br/&gt; - &#x60;PENDING_DELETE&#x60;: 삭제 대기 중 | 
 **operatingStatus** | [**LoadBalancerOperatingStatus**](LoadBalancerOperatingStatus.md) | 운영 상태 &lt;br/&gt; - &#x60;ONLINE&#x60;: 온라인 상태 &lt;br/&gt; - &#x60;DRAINING&#x60;: 연결 종료 중 &lt;br/&gt; - &#x60;OFFLINE&#x60;: 오프라인 상태 &lt;br/&gt; - &#x60;DEGRADED&#x60;: 성능 저하 상태 &lt;br/&gt; - &#x60;ERROR&#x60;: 오류 발생 &lt;br/&gt; - &#x60;NO_MONITOR&#x60;: 모니터링 없음 | 
 **type_** | [**LoadBalancerType**](LoadBalancerType.md) | [로드 밸런서 유형](https://docs.kakaocloud.com/service/bns/lb/lb-overview#사용-목적에-따른-로드-밸런서-유형-제공) &lt;br/&gt; - &#x60;ALB&#x60;: Application Load Balancer &lt;br/&gt; - &#x60;NLB&#x60;: Network Load Balancer &lt;br/&gt; - &#x60;NLB_L4_DSR&#x60;: Network Load Balancer (L4 Direct Server Return) | 
 **vpcName** | **string** | 연결된 VPC 이름  | 
 **vpcId** | **string** | 연결된 VPC ID | 
 **createdAt** | **string** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
 **updatedAt** | **string** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
 **sortKeys** | **string** | 정렬할 필드를 콤마(,)로 구분   | [default to &quot;created_at&quot;]
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;)   | [default to &quot;desc&quot;]
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | [default to 20]
 **offset** | **int32** | 조회 시작 위치 | [default to 0]

### Return type

[**BeyondLoadBalancerListModel**](BeyondLoadBalancerListModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHaGroup

> BnsLoadBalancerV1ApiUpdateHaGroupModelResponseBeyondLoadBalancerModel UpdateHaGroup(ctx, beyondLoadBalancerId).XAuthToken(xAuthToken).BodyUpdateHaGroup(bodyUpdateHaGroup).Execute()

Update HA group



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
	beyondLoadBalancerId := "beyondLoadBalancerId_example" // string | 수정할 고가용성 로드 밸런서 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bodyUpdateHaGroup := *openapiclient.NewBodyUpdateHaGroup(*openapiclient.NewEditBeyondLoadBalancerModel()) // BodyUpdateHaGroup | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BeyondLoadBalancerAPI.UpdateHaGroup(context.Background(), beyondLoadBalancerId).XAuthToken(xAuthToken).BodyUpdateHaGroup(bodyUpdateHaGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BeyondLoadBalancerAPI.UpdateHaGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateHaGroup`: BnsLoadBalancerV1ApiUpdateHaGroupModelResponseBeyondLoadBalancerModel
	fmt.Fprintf(os.Stdout, "Response from `BeyondLoadBalancerAPI.UpdateHaGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**beyondLoadBalancerId** | **string** | 수정할 고가용성 로드 밸런서 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHaGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyUpdateHaGroup** | [**BodyUpdateHaGroup**](BodyUpdateHaGroup.md) |  | 

### Return type

[**BnsLoadBalancerV1ApiUpdateHaGroupModelResponseBeyondLoadBalancerModel**](BnsLoadBalancerV1ApiUpdateHaGroupModelResponseBeyondLoadBalancerModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateHaGroupLoadBalancer

> BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelResponseBeyondLoadBalancerModel UpdateHaGroupLoadBalancer(ctx, beyondLoadBalancerId).XAuthToken(xAuthToken).BodyUpdateHaGroupLoadBalancer(bodyUpdateHaGroupLoadBalancer).Execute()

Update HA group load balancer



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
	beyondLoadBalancerId := "beyondLoadBalancerId_example" // string | 연결 대상 고가용성 로드 밸런서 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bodyUpdateHaGroupLoadBalancer := *openapiclient.NewBodyUpdateHaGroupLoadBalancer(*openapiclient.NewBnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelCreateBeyondLoadBalancerModel([]openapiclient.BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelSubnetModel{*openapiclient.NewBnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelSubnetModel("LoadBalancerId_example", openapiclient.AvailabilityZone("kr-central-2-a"))})) // BodyUpdateHaGroupLoadBalancer | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BeyondLoadBalancerAPI.UpdateHaGroupLoadBalancer(context.Background(), beyondLoadBalancerId).XAuthToken(xAuthToken).BodyUpdateHaGroupLoadBalancer(bodyUpdateHaGroupLoadBalancer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BeyondLoadBalancerAPI.UpdateHaGroupLoadBalancer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateHaGroupLoadBalancer`: BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelResponseBeyondLoadBalancerModel
	fmt.Fprintf(os.Stdout, "Response from `BeyondLoadBalancerAPI.UpdateHaGroupLoadBalancer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**beyondLoadBalancerId** | **string** | 연결 대상 고가용성 로드 밸런서 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHaGroupLoadBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyUpdateHaGroupLoadBalancer** | [**BodyUpdateHaGroupLoadBalancer**](BodyUpdateHaGroupLoadBalancer.md) |  | 

### Return type

[**BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelResponseBeyondLoadBalancerModel**](BnsLoadBalancerV1ApiUpdateHaGroupLoadBalancerModelResponseBeyondLoadBalancerModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

