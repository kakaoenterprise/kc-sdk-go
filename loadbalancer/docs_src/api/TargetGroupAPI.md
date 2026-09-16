# \TargetGroupAPI

All URIs are relative to *https://load-balancer.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddTarget**](TargetGroupAPI.md#AddTarget) | **Post** /api/v1/load-balancers/target-groups/{target_group_id}/members | Add target
[**CreateTargetGroup**](TargetGroupAPI.md#CreateTargetGroup) | **Post** /api/v1/load-balancers/target-groups | Create target group
[**DeleteTargetGroup**](TargetGroupAPI.md#DeleteTargetGroup) | **Delete** /api/v1/load-balancers/target-groups/{target_group_id} | Delete target group
[**GetTarget**](TargetGroupAPI.md#GetTarget) | **Get** /api/v1/load-balancers/target-groups/{target_group_id}/members/{member_id} | Get target
[**GetTargetGroup**](TargetGroupAPI.md#GetTargetGroup) | **Get** /api/v1/load-balancers/target-groups/{target_group_id} | Get target group
[**ListTargetGroups**](TargetGroupAPI.md#ListTargetGroups) | **Get** /api/v1/load-balancers/target-groups | List target groups
[**ListTargetsInTargetGroup**](TargetGroupAPI.md#ListTargetsInTargetGroup) | **Get** /api/v1/load-balancers/target-groups/{target_group_id}/members | List targets in target group
[**RemoveTarget**](TargetGroupAPI.md#RemoveTarget) | **Delete** /api/v1/load-balancers/target-groups/{target_group_id}/members/{member_id} | Remove target
[**UpdateTarget**](TargetGroupAPI.md#UpdateTarget) | **Put** /api/v1/load-balancers/target-groups/{target_group_id}/members/{member_id} | Update target
[**UpdateTargetGroup**](TargetGroupAPI.md#UpdateTargetGroup) | **Put** /api/v1/load-balancers/target-groups/{target_group_id} | Update target group
[**UpdateTargets**](TargetGroupAPI.md#UpdateTargets) | **Put** /api/v1/load-balancers/target-groups/{target_group_id}/members | Update targets



## AddTarget

> AddTargetResponse AddTarget(ctx, targetGroupId).AddTargetRequest(addTargetRequest).Execute()

Add target



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/loadbalancer"
)

func main() {
	targetGroupId := "targetGroupId_example" // string | 대상을 추가할 대상 그룹의 ID - [List target groups](/openapi/networking/lb/list-target-groups)에서 확인
	addTargetRequest := *openapiclient.NewAddTargetRequest(*openapiclient.NewAddTarget("Address_example", int32(123), "SubnetId_example")) // AddTargetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TargetGroupAPI.AddTarget(context.Background(), targetGroupId).AddTargetRequest(addTargetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TargetGroupAPI.AddTarget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddTarget`: AddTargetResponse
	fmt.Fprintf(os.Stdout, "Response from `TargetGroupAPI.AddTarget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetGroupId** | **string** | 대상을 추가할 대상 그룹의 ID - [List target groups](/openapi/networking/lb/list-target-groups)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addTargetRequest** | [**AddTargetRequest**](AddTargetRequest.md) |  | 

### Return type

[**AddTargetResponse**](AddTargetResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTargetGroup

> CreateTargetGroupResponse CreateTargetGroup(ctx).CreateTargetGroupRequest(createTargetGroupRequest).Execute()

Create target group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/loadbalancer"
)

func main() {
	createTargetGroupRequest := *openapiclient.NewCreateTargetGroupRequest(*openapiclient.NewCreateTargetGroup(openapiclient.LoadBalancerPoolAlgorithm("ROUND_ROBIN"), "Name_example", openapiclient.LoadBalancerPoolProtocol("HTTP"))) // CreateTargetGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TargetGroupAPI.CreateTargetGroup(context.Background()).CreateTargetGroupRequest(createTargetGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TargetGroupAPI.CreateTargetGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTargetGroup`: CreateTargetGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `TargetGroupAPI.CreateTargetGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTargetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTargetGroupRequest** | [**CreateTargetGroupRequest**](CreateTargetGroupRequest.md) |  | 

### Return type

[**CreateTargetGroupResponse**](CreateTargetGroupResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTargetGroup

> DeleteTargetGroup(ctx, targetGroupId).Execute()

Delete target group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/loadbalancer"
)

func main() {
	targetGroupId := "targetGroupId_example" // string | 삭제할 대상 그룹의 ID - [List target groups](/openapi/networking/lb/list-target-groups)에서 확인

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TargetGroupAPI.DeleteTargetGroup(context.Background(), targetGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TargetGroupAPI.DeleteTargetGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetGroupId** | **string** | 삭제할 대상 그룹의 ID - [List target groups](/openapi/networking/lb/list-target-groups)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTargetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetTarget

> GetTargetResponse GetTarget(ctx, targetGroupId, memberId).Execute()

Get target



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/loadbalancer"
)

func main() {
	targetGroupId := "targetGroupId_example" // string | 대상 그룹 ID - [List target groups](/openapi/networking/lb/list-target-groups)에서 확인
	memberId := "memberId_example" // string | 대상 그룹 내에서 조회할 멤버 리소스의 ID - [List targets in target group](/openapi/networking/lb/list-targets-in-target-group)에서 확인

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TargetGroupAPI.GetTarget(context.Background(), targetGroupId, memberId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TargetGroupAPI.GetTarget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTarget`: GetTargetResponse
	fmt.Fprintf(os.Stdout, "Response from `TargetGroupAPI.GetTarget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetGroupId** | **string** | 대상 그룹 ID - [List target groups](/openapi/networking/lb/list-target-groups)에서 확인 | 
**memberId** | **string** | 대상 그룹 내에서 조회할 멤버 리소스의 ID - [List targets in target group](/openapi/networking/lb/list-targets-in-target-group)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetTargetResponse**](GetTargetResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTargetGroup

> GetTargetGroupResponse GetTargetGroup(ctx, targetGroupId).Execute()

Get target group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/loadbalancer"
)

func main() {
	targetGroupId := "targetGroupId_example" // string | 조회할 대상 그룹 ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TargetGroupAPI.GetTargetGroup(context.Background(), targetGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TargetGroupAPI.GetTargetGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTargetGroup`: GetTargetGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `TargetGroupAPI.GetTargetGroup`: %v\n", resp)
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


### Return type

[**GetTargetGroupResponse**](GetTargetGroupResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTargetGroups

> ListTargetGroupsResponse ListTargetGroups(ctx).Id(id).Name(name).Protocol(protocol).AvailabilityZone(availabilityZone).LoadBalancerAlgorithm(loadBalancerAlgorithm).LoadBalancerName(loadBalancerName).LoadBalancerId(loadBalancerId).ListenerProtocol(listenerProtocol).VpcName(vpcName).VpcId(vpcId).SubnetName(subnetName).SubnetId(subnetId).HealthMonitorId(healthMonitorId).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Offset(offset).Limit(limit).Execute()

List target groups



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/loadbalancer"
)

func main() {
	id := "id_example" // string | 대상 그룹 ID (optional)
	name := "name_example" // string | 대상 그룹 이름 (optional)
	protocol := openapiclient.LoadBalancerPoolProtocol("HTTP") // LoadBalancerPoolProtocol | 대상 그룹의 프로토콜 (optional)
	availabilityZone := openapiclient.AvailabilityZone("kr-central-2-a") // AvailabilityZone | 대상 그룹이 위치한 가용 영역 (optional)
	loadBalancerAlgorithm := openapiclient.LoadBalancerPoolAlgorithm("ROUND_ROBIN") // LoadBalancerPoolAlgorithm | 로드 밸런싱 알고리즘 (optional)
	loadBalancerName := "loadBalancerName_example" // string | 연결된 로드 밸런서 이름 (optional)
	loadBalancerId := "loadBalancerId_example" // string | 연결된 로드 밸런서 ID (optional)
	listenerProtocol := openapiclient.Protocol("HTTP") // Protocol | 리스너 프로토콜 (optional)
	vpcName := "vpcName_example" // string | 대상 그룹이 속한 VPC 이름 (optional)
	vpcId := "vpcId_example" // string | 대상 그룹이 속한 VPC ID (optional)
	subnetName := "subnetName_example" // string | 서브넷 이름 (optional)
	subnetId := "subnetId_example" // string | 서브넷 ID (optional)
	healthMonitorId := "healthMonitorId_example" // string | 연결된 헬스 모니터 ID (optional)
	createdAt := "createdAt_example" // string | 리소스가 생성된 시간 - ISO 8601 형식 - UTC 기준 (optional)
	updatedAt := "updatedAt_example" // string | 리소스가 마지막으로 수정된 시간 - ISO 8601 형식 - UTC 기준 (optional)
	sortKeys := "sortKeys_example" // string | 정렬할 필드를 콤마(,)로 구분 (optional)
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`) (optional)
	offset := int32(56) // int32 | 조회 시작 위치 (optional)
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TargetGroupAPI.ListTargetGroups(context.Background()).Id(id).Name(name).Protocol(protocol).AvailabilityZone(availabilityZone).LoadBalancerAlgorithm(loadBalancerAlgorithm).LoadBalancerName(loadBalancerName).LoadBalancerId(loadBalancerId).ListenerProtocol(listenerProtocol).VpcName(vpcName).VpcId(vpcId).SubnetName(subnetName).SubnetId(subnetId).HealthMonitorId(healthMonitorId).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TargetGroupAPI.ListTargetGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTargetGroups`: ListTargetGroupsResponse
	fmt.Fprintf(os.Stdout, "Response from `TargetGroupAPI.ListTargetGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTargetGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **string** | 대상 그룹 ID | 
 **name** | **string** | 대상 그룹 이름 | 
 **protocol** | [**LoadBalancerPoolProtocol**](LoadBalancerPoolProtocol.md) | 대상 그룹의 프로토콜 | 
 **availabilityZone** | [**AvailabilityZone**](AvailabilityZone.md) | 대상 그룹이 위치한 가용 영역 | 
 **loadBalancerAlgorithm** | [**LoadBalancerPoolAlgorithm**](LoadBalancerPoolAlgorithm.md) | 로드 밸런싱 알고리즘 | 
 **loadBalancerName** | **string** | 연결된 로드 밸런서 이름 | 
 **loadBalancerId** | **string** | 연결된 로드 밸런서 ID | 
 **listenerProtocol** | [**Protocol**](Protocol.md) | 리스너 프로토콜 | 
 **vpcName** | **string** | 대상 그룹이 속한 VPC 이름 | 
 **vpcId** | **string** | 대상 그룹이 속한 VPC ID | 
 **subnetName** | **string** | 서브넷 이름 | 
 **subnetId** | **string** | 서브넷 ID | 
 **healthMonitorId** | **string** | 연결된 헬스 모니터 ID | 
 **createdAt** | **string** | 리소스가 생성된 시간 - ISO 8601 형식 - UTC 기준 | 
 **updatedAt** | **string** | 리소스가 마지막으로 수정된 시간 - ISO 8601 형식 - UTC 기준 | 
 **sortKeys** | **string** | 정렬할 필드를 콤마(,)로 구분 | 
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;) | 
 **offset** | **int32** | 조회 시작 위치 | 
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | 

### Return type

[**ListTargetGroupsResponse**](ListTargetGroupsResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTargetsInTargetGroup

> ListTargetsInTargetGroupResponse ListTargetsInTargetGroup(ctx, targetGroupId).Ip(ip).ProtocolPort(protocolPort).Weight(weight).ProvisioningStatus(provisioningStatus).OperatingStatus(operatingStatus).InstanceId(instanceId).InstanceName(instanceName).VpcId(vpcId).SubnetId(subnetId).SubnetName(subnetName).SecurityGroupName(securityGroupName).SortKeys(sortKeys).SortDirs(sortDirs).CreatedAt(createdAt).UpdatedAt(updatedAt).Offset(offset).Limit(limit).Execute()

List targets in target group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/loadbalancer"
)

func main() {
	targetGroupId := "targetGroupId_example" // string | 대상 그룹 ID - [List target groups](/openapi/networking/lb/list-target-groups)에서 확인
	ip := "ip_example" // string | 대상 인스턴스의 IP 주소 (optional)
	protocolPort := int32(56) // int32 | 연결 포트 번호 (optional)
	weight := int32(56) // int32 | 트래픽 분산 가중치 (optional)
	provisioningStatus := openapiclient.ProvisioningStatus("ACTIVE") // ProvisioningStatus | 프로비저닝 상태 (optional)
	operatingStatus := openapiclient.LoadBalancerOperatingStatus("ONLINE") // LoadBalancerOperatingStatus | 운영 상태 (optional)
	instanceId := "instanceId_example" // string | 인스턴스의 고유 ID (optional)
	instanceName := "instanceName_example" // string | 연결된 인스턴스 이름 (optional)
	vpcId := "vpcId_example" // string | 대상 인스턴스의 VPC ID (optional)
	subnetId := "subnetId_example" // string | 대상 인스턴스의 서브넷 ID (optional)
	subnetName := "subnetName_example" // string | 대상 인스턴스의 서브넷 이름 (optional)
	securityGroupName := "securityGroupName_example" // string | 보안 그룹 이름 (optional)
	sortKeys := "sortKeys_example" // string | 정렬할 필드를 콤마(,)로 구분 (optional)
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`) (optional)
	createdAt := "createdAt_example" // string | 리소스가 생성된 시간 - ISO 8601 형식 - UTC 기준 (optional)
	updatedAt := "updatedAt_example" // string | 리소스가 마지막으로 수정된 시간 - ISO 8601 형식 - UTC 기준 (optional)
	offset := int32(56) // int32 | 조회 시작 위치 (optional)
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TargetGroupAPI.ListTargetsInTargetGroup(context.Background(), targetGroupId).Ip(ip).ProtocolPort(protocolPort).Weight(weight).ProvisioningStatus(provisioningStatus).OperatingStatus(operatingStatus).InstanceId(instanceId).InstanceName(instanceName).VpcId(vpcId).SubnetId(subnetId).SubnetName(subnetName).SecurityGroupName(securityGroupName).SortKeys(sortKeys).SortDirs(sortDirs).CreatedAt(createdAt).UpdatedAt(updatedAt).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TargetGroupAPI.ListTargetsInTargetGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTargetsInTargetGroup`: ListTargetsInTargetGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `TargetGroupAPI.ListTargetsInTargetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetGroupId** | **string** | 대상 그룹 ID - [List target groups](/openapi/networking/lb/list-target-groups)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTargetsInTargetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ip** | **string** | 대상 인스턴스의 IP 주소 | 
 **protocolPort** | **int32** | 연결 포트 번호 | 
 **weight** | **int32** | 트래픽 분산 가중치 | 
 **provisioningStatus** | [**ProvisioningStatus**](ProvisioningStatus.md) | 프로비저닝 상태 | 
 **operatingStatus** | [**LoadBalancerOperatingStatus**](LoadBalancerOperatingStatus.md) | 운영 상태 | 
 **instanceId** | **string** | 인스턴스의 고유 ID | 
 **instanceName** | **string** | 연결된 인스턴스 이름 | 
 **vpcId** | **string** | 대상 인스턴스의 VPC ID | 
 **subnetId** | **string** | 대상 인스턴스의 서브넷 ID | 
 **subnetName** | **string** | 대상 인스턴스의 서브넷 이름 | 
 **securityGroupName** | **string** | 보안 그룹 이름 | 
 **sortKeys** | **string** | 정렬할 필드를 콤마(,)로 구분 | 
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;) | 
 **createdAt** | **string** | 리소스가 생성된 시간 - ISO 8601 형식 - UTC 기준 | 
 **updatedAt** | **string** | 리소스가 마지막으로 수정된 시간 - ISO 8601 형식 - UTC 기준 | 
 **offset** | **int32** | 조회 시작 위치 | 
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | 

### Return type

[**ListTargetsInTargetGroupResponse**](ListTargetsInTargetGroupResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveTarget

> RemoveTarget(ctx, targetGroupId, memberId).Execute()

Remove target



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/loadbalancer"
)

func main() {
	targetGroupId := "targetGroupId_example" // string | 대상 그룹 ID - [List target groups](/openapi/networking/lb/list-target-groups)에서 확인
	memberId := "memberId_example" // string | 삭제할 대상 인스턴스의 ID - [List targets in target group](/openapi/networking/lb/list-targets-in-target-group)에서 확인

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TargetGroupAPI.RemoveTarget(context.Background(), targetGroupId, memberId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TargetGroupAPI.RemoveTarget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetGroupId** | **string** | 대상 그룹 ID - [List target groups](/openapi/networking/lb/list-target-groups)에서 확인 | 
**memberId** | **string** | 삭제할 대상 인스턴스의 ID - [List targets in target group](/openapi/networking/lb/list-targets-in-target-group)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## UpdateTarget

> UpdateTargetResponse UpdateTarget(ctx, targetGroupId, memberId).UpdateTargetRequest(updateTargetRequest).Execute()

Update target



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/loadbalancer"
)

func main() {
	targetGroupId := "targetGroupId_example" // string | 대상 그룹 ID - [List target groups](/openapi/networking/lb/list-target-groups)에서 확인
	memberId := "memberId_example" // string | 수정할 대상 인스턴스의 ID - [List targets in target group](/openapi/networking/lb/list-targets-in-target-group)에서 확인
	updateTargetRequest := *openapiclient.NewUpdateTargetRequest(*openapiclient.NewUpdateTarget()) // UpdateTargetRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TargetGroupAPI.UpdateTarget(context.Background(), targetGroupId, memberId).UpdateTargetRequest(updateTargetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TargetGroupAPI.UpdateTarget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTarget`: UpdateTargetResponse
	fmt.Fprintf(os.Stdout, "Response from `TargetGroupAPI.UpdateTarget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetGroupId** | **string** | 대상 그룹 ID - [List target groups](/openapi/networking/lb/list-target-groups)에서 확인 | 
**memberId** | **string** | 수정할 대상 인스턴스의 ID - [List targets in target group](/openapi/networking/lb/list-targets-in-target-group)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTargetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateTargetRequest** | [**UpdateTargetRequest**](UpdateTargetRequest.md) |  | 

### Return type

[**UpdateTargetResponse**](UpdateTargetResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTargetGroup

> UpdateTargetGroupResponse UpdateTargetGroup(ctx, targetGroupId).UpdateTargetGroupRequest(updateTargetGroupRequest).Execute()

Update target group



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/loadbalancer"
)

func main() {
	targetGroupId := "targetGroupId_example" // string | 수정할 대상 그룹의 ID - [List target groups](/openapi/networking/lb/list-target-groups)에서 확인
	updateTargetGroupRequest := *openapiclient.NewUpdateTargetGroupRequest(*openapiclient.NewUpdateTargetGroup()) // UpdateTargetGroupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TargetGroupAPI.UpdateTargetGroup(context.Background(), targetGroupId).UpdateTargetGroupRequest(updateTargetGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TargetGroupAPI.UpdateTargetGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTargetGroup`: UpdateTargetGroupResponse
	fmt.Fprintf(os.Stdout, "Response from `TargetGroupAPI.UpdateTargetGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetGroupId** | **string** | 수정할 대상 그룹의 ID - [List target groups](/openapi/networking/lb/list-target-groups)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTargetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTargetGroupRequest** | [**UpdateTargetGroupRequest**](UpdateTargetGroupRequest.md) |  | 

### Return type

[**UpdateTargetGroupResponse**](UpdateTargetGroupResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTargets

> UpdateTargets(ctx, targetGroupId).UpdateTargetsRequest(updateTargetsRequest).Execute()

Update targets



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/loadbalancer"
)

func main() {
	targetGroupId := "targetGroupId_example" // string | 수정할 대상 인스턴스가 속한 대상 그룹의 ID - [List target groups](/openapi/networking/lb/list-target-groups)에서 확인
	updateTargetsRequest := *openapiclient.NewUpdateTargetsRequest([]openapiclient.UpdateTargets{*openapiclient.NewUpdateTargets("Address_example", int32(123), "SubnetId_example")}) // UpdateTargetsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TargetGroupAPI.UpdateTargets(context.Background(), targetGroupId).UpdateTargetsRequest(updateTargetsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TargetGroupAPI.UpdateTargets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetGroupId** | **string** | 수정할 대상 인스턴스가 속한 대상 그룹의 ID - [List target groups](/openapi/networking/lb/list-target-groups)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTargetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTargetsRequest** | [**UpdateTargetsRequest**](UpdateTargetsRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

