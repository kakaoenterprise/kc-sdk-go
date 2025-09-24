# \LoadBalancerTargetGroupAPI

All URIs are relative to *https://load-balancer.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTarget**](LoadBalancerTargetGroupAPI.md#AddTarget) | **Post** /api/v1/load-balancers/target-groups/{target_group_id}/members | Add target
[**CreateHealthMonitor**](LoadBalancerTargetGroupAPI.md#CreateHealthMonitor) | **Post** /api/v1/load-balancers/health-monitors | Create health monitor
[**CreateTargetGroup**](LoadBalancerTargetGroupAPI.md#CreateTargetGroup) | **Post** /api/v1/load-balancers/target-groups | Create target group
[**DeleteHealthMonitor**](LoadBalancerTargetGroupAPI.md#DeleteHealthMonitor) | **Delete** /api/v1/load-balancers/health-monitors/{health_monitor_id} | Delete health monitor
[**DeleteTargetGroup**](LoadBalancerTargetGroupAPI.md#DeleteTargetGroup) | **Delete** /api/v1/load-balancers/target-groups/{target_group_id} | Delete target group
[**GetTargetGroup**](LoadBalancerTargetGroupAPI.md#GetTargetGroup) | **Get** /api/v1/load-balancers/target-groups/{target_group_id} | Get target group
[**GetTargetGroupHealthCheckSubnets**](LoadBalancerTargetGroupAPI.md#GetTargetGroupHealthCheckSubnets) | **Get** /api/v1/load-balancers/target-groups/{target_group_id}/health-check-subnets | Get target group health check subnets
[**GetTargetGroupHealthMonitor**](LoadBalancerTargetGroupAPI.md#GetTargetGroupHealthMonitor) | **Get** /api/v1/load-balancers/health-monitors/{health_monitor_id} | Get target group health monitor
[**ListTargetGroups**](LoadBalancerTargetGroupAPI.md#ListTargetGroups) | **Get** /api/v1/load-balancers/target-groups | List target groups
[**ListTargetsInTargetGroup**](LoadBalancerTargetGroupAPI.md#ListTargetsInTargetGroup) | **Get** /api/v1/load-balancers/target-groups/{target_group_id}/members | List targets in target group
[**RemoveTarget**](LoadBalancerTargetGroupAPI.md#RemoveTarget) | **Delete** /api/v1/load-balancers/target-groups/{target_group_id}/members/{member_id} | Remove target
[**UpdateHealthMonitor**](LoadBalancerTargetGroupAPI.md#UpdateHealthMonitor) | **Put** /api/v1/load-balancers/health-monitors/{health_monitor_id} | Update health monitor
[**UpdateTarget**](LoadBalancerTargetGroupAPI.md#UpdateTarget) | **Put** /api/v1/load-balancers/target-groups/{target_group_id}/members/{member_id} | Update target
[**UpdateTargetGroup**](LoadBalancerTargetGroupAPI.md#UpdateTargetGroup) | **Put** /api/v1/load-balancers/target-groups/{target_group_id} | Update target group
[**UpdateTargets**](LoadBalancerTargetGroupAPI.md#UpdateTargets) | **Put** /api/v1/load-balancers/target-groups/{target_group_id}/members | Update targets



## AddTarget

> BnsLoadBalancerV1ApiAddTargetModelResponseTargetGroupMemberModel AddTarget(ctx, targetGroupId).XAuthToken(xAuthToken).BodyAddTarget(bodyAddTarget).Execute()

Add target



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
	targetGroupId := "targetGroupId_example" // string | 대상을 추가할 대상 그룹의 ID 
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bodyAddTarget := *openapiclient.NewBodyAddTarget(*openapiclient.NewCreateTargetGroupMember("Address_example", int32(123), "SubnetId_example")) // BodyAddTarget | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerTargetGroupAPI.AddTarget(context.Background(), targetGroupId).XAuthToken(xAuthToken).BodyAddTarget(bodyAddTarget).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerTargetGroupAPI.AddTarget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddTarget`: BnsLoadBalancerV1ApiAddTargetModelResponseTargetGroupMemberModel
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerTargetGroupAPI.AddTarget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetGroupId** | **string** | 대상을 추가할 대상 그룹의 ID  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyAddTarget** | [**BodyAddTarget**](BodyAddTarget.md) |  | 

### Return type

[**BnsLoadBalancerV1ApiAddTargetModelResponseTargetGroupMemberModel**](BnsLoadBalancerV1ApiAddTargetModelResponseTargetGroupMemberModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateHealthMonitor

> BnsLoadBalancerV1ApiCreateHealthMonitorModelResponseHealthMonitorModel CreateHealthMonitor(ctx).XAuthToken(xAuthToken).BodyCreateHealthMonitor(bodyCreateHealthMonitor).Execute()

Create health monitor



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
	bodyCreateHealthMonitor := *openapiclient.NewBodyCreateHealthMonitor(*openapiclient.NewCreateHealthMonitor(int32(123), int32(123), int32(123), "TargetGroupId_example", int32(123), openapiclient.HealthMonitorType("HTTP"))) // BodyCreateHealthMonitor | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerTargetGroupAPI.CreateHealthMonitor(context.Background()).XAuthToken(xAuthToken).BodyCreateHealthMonitor(bodyCreateHealthMonitor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerTargetGroupAPI.CreateHealthMonitor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateHealthMonitor`: BnsLoadBalancerV1ApiCreateHealthMonitorModelResponseHealthMonitorModel
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerTargetGroupAPI.CreateHealthMonitor`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateHealthMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyCreateHealthMonitor** | [**BodyCreateHealthMonitor**](BodyCreateHealthMonitor.md) |  | 

### Return type

[**BnsLoadBalancerV1ApiCreateHealthMonitorModelResponseHealthMonitorModel**](BnsLoadBalancerV1ApiCreateHealthMonitorModelResponseHealthMonitorModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTargetGroup

> BnsLoadBalancerV1ApiCreateTargetGroupModelResponseTargetGroupModel CreateTargetGroup(ctx).XAuthToken(xAuthToken).BodyCreateTargetGroup(bodyCreateTargetGroup).Execute()

Create target group



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
	bodyCreateTargetGroup := *openapiclient.NewBodyCreateTargetGroup(*openapiclient.NewCreateTargetGroup(openapiclient.TargetGroupAlgorithm("ROUND_ROBIN"), "LoadBalancerId_example", "Name_example", openapiclient.TargetGroupProtocol("HTTP"))) // BodyCreateTargetGroup | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerTargetGroupAPI.CreateTargetGroup(context.Background()).XAuthToken(xAuthToken).BodyCreateTargetGroup(bodyCreateTargetGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerTargetGroupAPI.CreateTargetGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTargetGroup`: BnsLoadBalancerV1ApiCreateTargetGroupModelResponseTargetGroupModel
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerTargetGroupAPI.CreateTargetGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTargetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyCreateTargetGroup** | [**BodyCreateTargetGroup**](BodyCreateTargetGroup.md) |  | 

### Return type

[**BnsLoadBalancerV1ApiCreateTargetGroupModelResponseTargetGroupModel**](BnsLoadBalancerV1ApiCreateTargetGroupModelResponseTargetGroupModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteHealthMonitor

> DeleteHealthMonitor(ctx, healthMonitorId).XAuthToken(xAuthToken).Execute()

Delete health monitor



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
	healthMonitorId := "healthMonitorId_example" // string | 삭제할 헬스 모니터의 ID 
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LoadBalancerTargetGroupAPI.DeleteHealthMonitor(context.Background(), healthMonitorId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerTargetGroupAPI.DeleteHealthMonitor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**healthMonitorId** | **string** | 삭제할 헬스 모니터의 ID  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteHealthMonitorRequest struct via the builder pattern


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


## DeleteTargetGroup

> DeleteTargetGroup(ctx, targetGroupId).XAuthToken(xAuthToken).Execute()

Delete target group



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
	targetGroupId := "targetGroupId_example" // string | 삭제할 대상 그룹의 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LoadBalancerTargetGroupAPI.DeleteTargetGroup(context.Background(), targetGroupId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerTargetGroupAPI.DeleteTargetGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetGroupId** | **string** | 삭제할 대상 그룹의 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTargetGroupRequest struct via the builder pattern


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


## GetTargetGroup

> TargetGroupResponseModel GetTargetGroup(ctx, targetGroupId).XAuthToken(xAuthToken).Execute()

Get target group



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
	targetGroupId := "targetGroupId_example" // string | 조회할 대상 그룹 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerTargetGroupAPI.GetTargetGroup(context.Background(), targetGroupId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerTargetGroupAPI.GetTargetGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTargetGroup`: TargetGroupResponseModel
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerTargetGroupAPI.GetTargetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetGroupId** | **string** | 조회할 대상 그룹 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTargetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**TargetGroupResponseModel**](TargetGroupResponseModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTargetGroupHealthCheckSubnets

> ResponseHealthCheckSubnetsModel GetTargetGroupHealthCheckSubnets(ctx, targetGroupId).XAuthToken(xAuthToken).Execute()

Get target group health check subnets



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
	targetGroupId := "targetGroupId_example" // string | 대상 그룹 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerTargetGroupAPI.GetTargetGroupHealthCheckSubnets(context.Background(), targetGroupId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerTargetGroupAPI.GetTargetGroupHealthCheckSubnets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTargetGroupHealthCheckSubnets`: ResponseHealthCheckSubnetsModel
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerTargetGroupAPI.GetTargetGroupHealthCheckSubnets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetGroupId** | **string** | 대상 그룹 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTargetGroupHealthCheckSubnetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**ResponseHealthCheckSubnetsModel**](ResponseHealthCheckSubnetsModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTargetGroupHealthMonitor

> BnsLoadBalancerV1ApiGetTargetGroupHealthMonitorModelResponseHealthMonitorModel GetTargetGroupHealthMonitor(ctx, healthMonitorId).XAuthToken(xAuthToken).Execute()

Get target group health monitor



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
	healthMonitorId := "healthMonitorId_example" // string | 조회할 헬스 모니터 ID 
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerTargetGroupAPI.GetTargetGroupHealthMonitor(context.Background(), healthMonitorId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerTargetGroupAPI.GetTargetGroupHealthMonitor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTargetGroupHealthMonitor`: BnsLoadBalancerV1ApiGetTargetGroupHealthMonitorModelResponseHealthMonitorModel
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerTargetGroupAPI.GetTargetGroupHealthMonitor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**healthMonitorId** | **string** | 조회할 헬스 모니터 ID  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTargetGroupHealthMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**BnsLoadBalancerV1ApiGetTargetGroupHealthMonitorModelResponseHealthMonitorModel**](BnsLoadBalancerV1ApiGetTargetGroupHealthMonitorModelResponseHealthMonitorModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTargetGroups

> TargetGroupListModel ListTargetGroups(ctx).XAuthToken(xAuthToken).Id(id).Name(name).Protocol(protocol).AvailabilityZone(availabilityZone).LoadBalancerAlgorithm(loadBalancerAlgorithm).LoadBalancerName(loadBalancerName).LoadBalancerId(loadBalancerId).ListenerProtocol(listenerProtocol).VpcName(vpcName).VpcId(vpcId).SubnetName(subnetName).SubnetId(subnetId).HealthMonitorId(healthMonitorId).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Limit(limit).Offset(offset).Execute()

List target groups



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
	id := "id_example" // string | 대상 그룹 ID  (optional)
	name := "name_example" // string | 대상 그룹 이름  (optional)
	protocol := openapiclient.TargetGroupProtocol("HTTP") // TargetGroupProtocol | 대상 그룹의 프로토콜 <br/> - `HTTP`: HTTP <br/> - `HTTPS`: HTTPS <br/> - `TCP`: TCP <br/> - `UDP`: UDP <br/> - `PROXY`: 프록시 프로토콜 (optional)
	availabilityZone := openapiclient.AvailabilityZone("kr-central-2-a") // AvailabilityZone | 대상 그룹이 위치한 가용 영역   (optional)
	loadBalancerAlgorithm := openapiclient.TargetGroupAlgorithm("ROUND_ROBIN") // TargetGroupAlgorithm | 로드 밸런싱 알고리즘 <br/> - `ROUND_ROBIN`: 라운드 로빈 방식 <br/> - `LEAST_CONNECTIONS`: 최소 연결 방식 <br/> - `SOURCE_IP`: 소스 IP 기반 방식 (optional)
	loadBalancerName := "loadBalancerName_example" // string | 연결된 로드 밸런서 이름  (optional)
	loadBalancerId := "loadBalancerId_example" // string | 연결된 로드 밸런서 ID  (optional)
	listenerProtocol := openapiclient.Protocol("HTTP") // Protocol | 리스너 프로토콜 <br/> - `HTTP`: HTTP <br/> - `TCP`: TCP <br/> - `UDP`: UDP <br/> - `TERMINATED_HTTPS`: 로드 밸런서에서 SSL 종료 처리된 HTTPS (optional)
	vpcName := "vpcName_example" // string | 대상 그룹이 속한 VPC 이름  (optional)
	vpcId := "vpcId_example" // string | 대상 그룹이 속한 VPC ID  (optional)
	subnetName := "subnetName_example" // string | 서브넷 이름  (optional)
	subnetId := "subnetId_example" // string | 서브넷 ID  (optional)
	healthMonitorId := "healthMonitorId_example" // string | 연결된 헬스 모니터 ID  (optional)
	createdAt := "createdAt_example" // string | 리소스가 생성된 시간 <br/> - ISO_8601 형식  <br/> - UTC 기준 (optional)
	updatedAt := "updatedAt_example" // string | 리소스가 마지막으로 수정된 시간 <br/> - ISO_8601 형식  <br/> - UTC 기준 (optional)
	sortKeys := "sortKeys_example" // string | 정렬할 필드를 콤마(,)로 구분   (optional) (default to "created_at")
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`)   (optional) (default to "desc")
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional) (default to 20)
	offset := int32(56) // int32 | 조회 시작 위치 (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerTargetGroupAPI.ListTargetGroups(context.Background()).XAuthToken(xAuthToken).Id(id).Name(name).Protocol(protocol).AvailabilityZone(availabilityZone).LoadBalancerAlgorithm(loadBalancerAlgorithm).LoadBalancerName(loadBalancerName).LoadBalancerId(loadBalancerId).ListenerProtocol(listenerProtocol).VpcName(vpcName).VpcId(vpcId).SubnetName(subnetName).SubnetId(subnetId).HealthMonitorId(healthMonitorId).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerTargetGroupAPI.ListTargetGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTargetGroups`: TargetGroupListModel
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerTargetGroupAPI.ListTargetGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTargetGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **id** | **string** | 대상 그룹 ID  | 
 **name** | **string** | 대상 그룹 이름  | 
 **protocol** | [**TargetGroupProtocol**](TargetGroupProtocol.md) | 대상 그룹의 프로토콜 &lt;br/&gt; - &#x60;HTTP&#x60;: HTTP &lt;br/&gt; - &#x60;HTTPS&#x60;: HTTPS &lt;br/&gt; - &#x60;TCP&#x60;: TCP &lt;br/&gt; - &#x60;UDP&#x60;: UDP &lt;br/&gt; - &#x60;PROXY&#x60;: 프록시 프로토콜 | 
 **availabilityZone** | [**AvailabilityZone**](AvailabilityZone.md) | 대상 그룹이 위치한 가용 영역   | 
 **loadBalancerAlgorithm** | [**TargetGroupAlgorithm**](TargetGroupAlgorithm.md) | 로드 밸런싱 알고리즘 &lt;br/&gt; - &#x60;ROUND_ROBIN&#x60;: 라운드 로빈 방식 &lt;br/&gt; - &#x60;LEAST_CONNECTIONS&#x60;: 최소 연결 방식 &lt;br/&gt; - &#x60;SOURCE_IP&#x60;: 소스 IP 기반 방식 | 
 **loadBalancerName** | **string** | 연결된 로드 밸런서 이름  | 
 **loadBalancerId** | **string** | 연결된 로드 밸런서 ID  | 
 **listenerProtocol** | [**Protocol**](Protocol.md) | 리스너 프로토콜 &lt;br/&gt; - &#x60;HTTP&#x60;: HTTP &lt;br/&gt; - &#x60;TCP&#x60;: TCP &lt;br/&gt; - &#x60;UDP&#x60;: UDP &lt;br/&gt; - &#x60;TERMINATED_HTTPS&#x60;: 로드 밸런서에서 SSL 종료 처리된 HTTPS | 
 **vpcName** | **string** | 대상 그룹이 속한 VPC 이름  | 
 **vpcId** | **string** | 대상 그룹이 속한 VPC ID  | 
 **subnetName** | **string** | 서브넷 이름  | 
 **subnetId** | **string** | 서브넷 ID  | 
 **healthMonitorId** | **string** | 연결된 헬스 모니터 ID  | 
 **createdAt** | **string** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
 **updatedAt** | **string** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
 **sortKeys** | **string** | 정렬할 필드를 콤마(,)로 구분   | [default to &quot;created_at&quot;]
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;)   | [default to &quot;desc&quot;]
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | [default to 20]
 **offset** | **int32** | 조회 시작 위치 | [default to 0]

### Return type

[**TargetGroupListModel**](TargetGroupListModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTargetsInTargetGroup

> TargetGroupMemberListModel ListTargetsInTargetGroup(ctx, targetGroupId).XAuthToken(xAuthToken).Ip(ip).ProtocolPort(protocolPort).Weight(weight).ProvisioningStatus(provisioningStatus).OperatingStatus(operatingStatus).InstanceId(instanceId).InstanceName(instanceName).VpcId(vpcId).SubnetId(subnetId).SubnetName(subnetName).SecurityGroupName(securityGroupName).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Limit(limit).Offset(offset).Execute()

List targets in target group



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
	targetGroupId := "targetGroupId_example" // string | 대상 그룹 ID 
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	ip := "ip_example" // string | 대상 인스턴스의 IP 주소  (optional)
	protocolPort := "protocolPort_example" // string | 연결 포트 번호  (optional)
	weight := "weight_example" // string | 트래픽 분산 가중치  (optional)
	provisioningStatus := openapiclient.ProvisioningStatus("ACTIVE") // ProvisioningStatus | 프로비저닝 상태 <br/> - `ACTIVE`: 활성 <br/> - `DELETED`: 삭제됨 <br/> - `ERROR`: 오류 <br/> - `PENDING_CREATE`: 생성 대기 중 <br/> - `PENDING_UPDATE`: 업데이트 대기 중 <br/> - `PENDING_DELETE`: 삭제 대기 중 (optional)
	operatingStatus := openapiclient.LoadBalancerOperatingStatus("ONLINE") // LoadBalancerOperatingStatus | 운영 상태 <br/> - `ONLINE`: 온라인 <br/> - `DRAINING`: 연결 종료 중 <br/> - `OFFLINE`: 오프라인 <br/> - `DEGRADED`: 성능 저하 <br/> - `ERROR`: 오류 <br/> - `NO_MONITOR`: 모니터링 없음 (optional)
	instanceId := "instanceId_example" // string | 인스턴스의 고유 ID (optional)
	instanceName := "instanceName_example" // string | 연결된 인스턴스 이름  (optional)
	vpcId := "vpcId_example" // string | 대상 인스턴스의 VPC ID  (optional)
	subnetId := "subnetId_example" // string | 대상 인스턴스의 서브넷 ID  (optional)
	subnetName := "subnetName_example" // string | 대상 인스턴스의 서브넷 이름  (optional)
	securityGroupName := "securityGroupName_example" // string | 보안 그룹 이름  (optional)
	createdAt := "createdAt_example" // string | 리소스가 생성된 시간 <br/> - ISO_8601 형식  <br/> - UTC 기준 (optional)
	updatedAt := "updatedAt_example" // string | 리소스가 마지막으로 수정된 시간 <br/> - ISO_8601 형식  <br/> - UTC 기준 (optional)
	sortKeys := "sortKeys_example" // string | 정렬할 필드를 콤마(,)로 구분   (optional) (default to "created_at")
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`)  (optional) (default to "desc")
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional) (default to 20)
	offset := int32(56) // int32 | 조회 시작 위치 (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerTargetGroupAPI.ListTargetsInTargetGroup(context.Background(), targetGroupId).XAuthToken(xAuthToken).Ip(ip).ProtocolPort(protocolPort).Weight(weight).ProvisioningStatus(provisioningStatus).OperatingStatus(operatingStatus).InstanceId(instanceId).InstanceName(instanceName).VpcId(vpcId).SubnetId(subnetId).SubnetName(subnetName).SecurityGroupName(securityGroupName).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerTargetGroupAPI.ListTargetsInTargetGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTargetsInTargetGroup`: TargetGroupMemberListModel
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerTargetGroupAPI.ListTargetsInTargetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetGroupId** | **string** | 대상 그룹 ID  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTargetsInTargetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **ip** | **string** | 대상 인스턴스의 IP 주소  | 
 **protocolPort** | **string** | 연결 포트 번호  | 
 **weight** | **string** | 트래픽 분산 가중치  | 
 **provisioningStatus** | [**ProvisioningStatus**](ProvisioningStatus.md) | 프로비저닝 상태 &lt;br/&gt; - &#x60;ACTIVE&#x60;: 활성 &lt;br/&gt; - &#x60;DELETED&#x60;: 삭제됨 &lt;br/&gt; - &#x60;ERROR&#x60;: 오류 &lt;br/&gt; - &#x60;PENDING_CREATE&#x60;: 생성 대기 중 &lt;br/&gt; - &#x60;PENDING_UPDATE&#x60;: 업데이트 대기 중 &lt;br/&gt; - &#x60;PENDING_DELETE&#x60;: 삭제 대기 중 | 
 **operatingStatus** | [**LoadBalancerOperatingStatus**](LoadBalancerOperatingStatus.md) | 운영 상태 &lt;br/&gt; - &#x60;ONLINE&#x60;: 온라인 &lt;br/&gt; - &#x60;DRAINING&#x60;: 연결 종료 중 &lt;br/&gt; - &#x60;OFFLINE&#x60;: 오프라인 &lt;br/&gt; - &#x60;DEGRADED&#x60;: 성능 저하 &lt;br/&gt; - &#x60;ERROR&#x60;: 오류 &lt;br/&gt; - &#x60;NO_MONITOR&#x60;: 모니터링 없음 | 
 **instanceId** | **string** | 인스턴스의 고유 ID | 
 **instanceName** | **string** | 연결된 인스턴스 이름  | 
 **vpcId** | **string** | 대상 인스턴스의 VPC ID  | 
 **subnetId** | **string** | 대상 인스턴스의 서브넷 ID  | 
 **subnetName** | **string** | 대상 인스턴스의 서브넷 이름  | 
 **securityGroupName** | **string** | 보안 그룹 이름  | 
 **createdAt** | **string** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
 **updatedAt** | **string** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
 **sortKeys** | **string** | 정렬할 필드를 콤마(,)로 구분   | [default to &quot;created_at&quot;]
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;)  | [default to &quot;desc&quot;]
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | [default to 20]
 **offset** | **int32** | 조회 시작 위치 | [default to 0]

### Return type

[**TargetGroupMemberListModel**](TargetGroupMemberListModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveTarget

> RemoveTarget(ctx, targetGroupId, memberId).XAuthToken(xAuthToken).Execute()

Remove target



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
	targetGroupId := "targetGroupId_example" // string | 대상 그룹 ID 
	memberId := "memberId_example" // string | 삭제할 대상 인스턴스의 ID 
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LoadBalancerTargetGroupAPI.RemoveTarget(context.Background(), targetGroupId, memberId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerTargetGroupAPI.RemoveTarget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetGroupId** | **string** | 대상 그룹 ID  | 
**memberId** | **string** | 삭제할 대상 인스턴스의 ID  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveTargetRequest struct via the builder pattern


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


## UpdateHealthMonitor

> BnsLoadBalancerV1ApiUpdateHealthMonitorModelResponseHealthMonitorModel UpdateHealthMonitor(ctx, healthMonitorId).XAuthToken(xAuthToken).BodyUpdateHealthMonitor(bodyUpdateHealthMonitor).Execute()

Update health monitor



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
	healthMonitorId := "healthMonitorId_example" // string | 수정 대상 헬스 모니터의 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bodyUpdateHealthMonitor := *openapiclient.NewBodyUpdateHealthMonitor(*openapiclient.NewEditHealthMonitor()) // BodyUpdateHealthMonitor | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerTargetGroupAPI.UpdateHealthMonitor(context.Background(), healthMonitorId).XAuthToken(xAuthToken).BodyUpdateHealthMonitor(bodyUpdateHealthMonitor).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerTargetGroupAPI.UpdateHealthMonitor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateHealthMonitor`: BnsLoadBalancerV1ApiUpdateHealthMonitorModelResponseHealthMonitorModel
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerTargetGroupAPI.UpdateHealthMonitor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**healthMonitorId** | **string** | 수정 대상 헬스 모니터의 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHealthMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyUpdateHealthMonitor** | [**BodyUpdateHealthMonitor**](BodyUpdateHealthMonitor.md) |  | 

### Return type

[**BnsLoadBalancerV1ApiUpdateHealthMonitorModelResponseHealthMonitorModel**](BnsLoadBalancerV1ApiUpdateHealthMonitorModelResponseHealthMonitorModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTarget

> BnsLoadBalancerV1ApiUpdateTargetModelResponseTargetGroupMemberModel UpdateTarget(ctx, targetGroupId, memberId).XAuthToken(xAuthToken).BodyUpdateTarget(bodyUpdateTarget).Execute()

Update target



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
	targetGroupId := "targetGroupId_example" // string | 대상 그룹 ID
	memberId := "memberId_example" // string | 수정할 대상 인스턴스의 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bodyUpdateTarget := *openapiclient.NewBodyUpdateTarget(*openapiclient.NewBnsLoadBalancerV1ApiUpdateTargetModelEditTargetGroupMember()) // BodyUpdateTarget | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerTargetGroupAPI.UpdateTarget(context.Background(), targetGroupId, memberId).XAuthToken(xAuthToken).BodyUpdateTarget(bodyUpdateTarget).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerTargetGroupAPI.UpdateTarget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTarget`: BnsLoadBalancerV1ApiUpdateTargetModelResponseTargetGroupMemberModel
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerTargetGroupAPI.UpdateTarget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetGroupId** | **string** | 대상 그룹 ID | 
**memberId** | **string** | 수정할 대상 인스턴스의 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyUpdateTarget** | [**BodyUpdateTarget**](BodyUpdateTarget.md) |  | 

### Return type

[**BnsLoadBalancerV1ApiUpdateTargetModelResponseTargetGroupMemberModel**](BnsLoadBalancerV1ApiUpdateTargetModelResponseTargetGroupMemberModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTargetGroup

> BnsLoadBalancerV1ApiUpdateTargetGroupModelResponseTargetGroupModel UpdateTargetGroup(ctx, targetGroupId).XAuthToken(xAuthToken).BodyUpdateTargetGroup(bodyUpdateTargetGroup).Execute()

Update target group



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
	targetGroupId := "targetGroupId_example" // string | 수정할 대상 그룹의 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bodyUpdateTargetGroup := *openapiclient.NewBodyUpdateTargetGroup(*openapiclient.NewEditTargetGroup()) // BodyUpdateTargetGroup | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerTargetGroupAPI.UpdateTargetGroup(context.Background(), targetGroupId).XAuthToken(xAuthToken).BodyUpdateTargetGroup(bodyUpdateTargetGroup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerTargetGroupAPI.UpdateTargetGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTargetGroup`: BnsLoadBalancerV1ApiUpdateTargetGroupModelResponseTargetGroupModel
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerTargetGroupAPI.UpdateTargetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetGroupId** | **string** | 수정할 대상 그룹의 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTargetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyUpdateTargetGroup** | [**BodyUpdateTargetGroup**](BodyUpdateTargetGroup.md) |  | 

### Return type

[**BnsLoadBalancerV1ApiUpdateTargetGroupModelResponseTargetGroupModel**](BnsLoadBalancerV1ApiUpdateTargetGroupModelResponseTargetGroupModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTargets

> interface{} UpdateTargets(ctx, targetGroupId).XAuthToken(xAuthToken).BodyUpdateTargets(bodyUpdateTargets).Execute()

Update targets



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
	targetGroupId := "targetGroupId_example" // string | 수정할 대상 인스턴스가 속한 대상 그룹의 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bodyUpdateTargets := *openapiclient.NewBodyUpdateTargets([]openapiclient.BnsLoadBalancerV1ApiUpdateTargetsModelEditTargetGroupMember{*openapiclient.NewBnsLoadBalancerV1ApiUpdateTargetsModelEditTargetGroupMember("Address_example", int32(123), "SubnetId_example")}) // BodyUpdateTargets | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoadBalancerTargetGroupAPI.UpdateTargets(context.Background(), targetGroupId).XAuthToken(xAuthToken).BodyUpdateTargets(bodyUpdateTargets).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoadBalancerTargetGroupAPI.UpdateTargets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTargets`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `LoadBalancerTargetGroupAPI.UpdateTargets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetGroupId** | **string** | 수정할 대상 인스턴스가 속한 대상 그룹의 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyUpdateTargets** | [**BodyUpdateTargets**](BodyUpdateTargets.md) |  | 

### Return type

**interface{}**

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

