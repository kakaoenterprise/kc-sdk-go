# \FlavorAPI

All URIs are relative to *https://bcs.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInstanceType**](FlavorAPI.md#GetInstanceType) | **Get** /api/v1/flavors/{flavor_id} | Get instance type
[**ListInstanceTypes**](FlavorAPI.md#ListInstanceTypes) | **Get** /api/v1/flavors | List instance types (flavors) 



## GetInstanceType

> ResponseFlavorModel GetInstanceType(ctx, flavorId).XAuthToken(xAuthToken).Execute()

Get instance type



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
	flavorId := "flavorId_example" // string | 조회할 인스턴스 유형(Flavor)의 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlavorAPI.GetInstanceType(context.Background(), flavorId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlavorAPI.GetInstanceType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstanceType`: ResponseFlavorModel
	fmt.Fprintf(os.Stdout, "Response from `FlavorAPI.GetInstanceType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**flavorId** | **string** | 조회할 인스턴스 유형(Flavor)의 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**ResponseFlavorModel**](ResponseFlavorModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInstanceTypes

> FlavorListModel ListInstanceTypes(ctx).XAuthToken(xAuthToken).Id(id).Name(name).IsBurstable(isBurstable).Vcpus(vcpus).Architecture(architecture).MemoryMb(memoryMb).InstanceType(instanceType).InstanceFamily(instanceFamily).InstanceSize(instanceSize).Manufacturer(manufacturer).MaximumNetworkInterfaces(maximumNetworkInterfaces).Processor(processor).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Limit(limit).Offset(offset).Execute()

List instance types (flavors) 



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
	id := "id_example" // string | 조회할 인스턴스 유형(Flavor) ID (optional)
	name := "name_example" // string | 인스턴스 유형의 이름 (optional)
	isBurstable := true // bool | [버스터블 인스턴스](https://docs.kakaocloud.com/service/bcs/bcs-instance/bcs-type/general-purpose/burstable-main) 여부 (optional)
	vcpus := int32(56) // int32 | 가상 CPU 개수 (optional)
	architecture := "architecture_example" // string | 인스턴스 유형의 아키텍처(CPU 구조) (optional)
	memoryMb := int32(56) // int32 | 메모리 크기 (MB 단위) (optional)
	instanceType := openapiclient.InstanceType("vm") // InstanceType | 인스턴스 유형 <br/> - `vm`: Virtual Machine 유형  <br/> - `bm`: Bare Metal Server 유형 <br/> - `gpu`: GPU 유형 (optional)
	instanceFamily := "instanceFamily_example" // string | [인스턴스 패밀리](https://docs.kakaocloud.com/service/bcs/bcs-instance/bcs-instance-overview#instance-family) <br/> - 예시:  `r2a`, `c2a`  등 (optional)
	instanceSize := "instanceSize_example" // string | 인스턴스 크기 (optional)
	manufacturer := "manufacturer_example" // string | 제조사 정보   (optional)
	maximumNetworkInterfaces := int32(56) // int32 | 최대 네트워크 인터페이스 개수 (optional)
	processor := "processor_example" // string | 프로세서 이름 (optional)
	createdAt := "createdAt_example" // string | 리소스가 생성된 시간 <br/> - ISO_8601 형식  <br/> - UTC 기준 (optional)
	updatedAt := "updatedAt_example" // string | 리소스가 마지막으로 수정된 시간 <br/> - ISO_8601 형식  <br/> - UTC 기준 (optional)
	sortKeys := "sortKeys_example" // string |  정렬할 필드를 콤마(,)로 구분   (optional) (default to "created_at")
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`)   (optional) (default to "desc")
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional) (default to 20)
	offset := int32(56) // int32 | 조회 시작 위치 (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlavorAPI.ListInstanceTypes(context.Background()).XAuthToken(xAuthToken).Id(id).Name(name).IsBurstable(isBurstable).Vcpus(vcpus).Architecture(architecture).MemoryMb(memoryMb).InstanceType(instanceType).InstanceFamily(instanceFamily).InstanceSize(instanceSize).Manufacturer(manufacturer).MaximumNetworkInterfaces(maximumNetworkInterfaces).Processor(processor).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlavorAPI.ListInstanceTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInstanceTypes`: FlavorListModel
	fmt.Fprintf(os.Stdout, "Response from `FlavorAPI.ListInstanceTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInstanceTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **id** | **string** | 조회할 인스턴스 유형(Flavor) ID | 
 **name** | **string** | 인스턴스 유형의 이름 | 
 **isBurstable** | **bool** | [버스터블 인스턴스](https://docs.kakaocloud.com/service/bcs/bcs-instance/bcs-type/general-purpose/burstable-main) 여부 | 
 **vcpus** | **int32** | 가상 CPU 개수 | 
 **architecture** | **string** | 인스턴스 유형의 아키텍처(CPU 구조) | 
 **memoryMb** | **int32** | 메모리 크기 (MB 단위) | 
 **instanceType** | [**InstanceType**](InstanceType.md) | 인스턴스 유형 &lt;br/&gt; - &#x60;vm&#x60;: Virtual Machine 유형  &lt;br/&gt; - &#x60;bm&#x60;: Bare Metal Server 유형 &lt;br/&gt; - &#x60;gpu&#x60;: GPU 유형 | 
 **instanceFamily** | **string** | [인스턴스 패밀리](https://docs.kakaocloud.com/service/bcs/bcs-instance/bcs-instance-overview#instance-family) &lt;br/&gt; - 예시:  &#x60;r2a&#x60;, &#x60;c2a&#x60;  등 | 
 **instanceSize** | **string** | 인스턴스 크기 | 
 **manufacturer** | **string** | 제조사 정보   | 
 **maximumNetworkInterfaces** | **int32** | 최대 네트워크 인터페이스 개수 | 
 **processor** | **string** | 프로세서 이름 | 
 **createdAt** | **string** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
 **updatedAt** | **string** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
 **sortKeys** | **string** |  정렬할 필드를 콤마(,)로 구분   | [default to &quot;created_at&quot;]
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;)   | [default to &quot;desc&quot;]
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | [default to 20]
 **offset** | **int32** | 조회 시작 위치 | [default to 0]

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

