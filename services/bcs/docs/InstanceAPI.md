# \InstanceAPI

All URIs are relative to *https://bcs.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInstance**](InstanceAPI.md#CreateInstance) | **Post** /api/v1/instances | Create instance
[**DeleteInstance**](InstanceAPI.md#DeleteInstance) | **Delete** /api/v1/instances/{instance_id} | Delete instance
[**GetInstance**](InstanceAPI.md#GetInstance) | **Get** /api/v1/instances/{instance_id} | Get instance
[**ListInstances**](InstanceAPI.md#ListInstances) | **Get** /api/v1/instances | List instances
[**UpdateInstance**](InstanceAPI.md#UpdateInstance) | **Put** /api/v1/instances/{instance_id} | Update instance



## CreateInstance

> ResponseCreateInstanceModel CreateInstance(ctx).XAuthToken(xAuthToken).BodyCreateInstance(bodyCreateInstance).Execute()

Create instance



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
	bodyCreateInstance := *openapiclient.NewBodyCreateInstance(*openapiclient.NewCreateInstanceModel("Name_example", "ImageId_example", "FlavorId_example", []openapiclient.CreateInstanceSubnetModel{*openapiclient.NewCreateInstanceSubnetModel("Id_example")}, []openapiclient.CreateInstanceVolumeModel{*openapiclient.NewCreateInstanceVolumeModel(int32(123))}, []openapiclient.CreateInstanceSecurityGroupModel{*openapiclient.NewCreateInstanceSecurityGroupModel("Name_example")})) // BodyCreateInstance | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.CreateInstance(context.Background()).XAuthToken(xAuthToken).BodyCreateInstance(bodyCreateInstance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.CreateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInstance`: ResponseCreateInstanceModel
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.CreateInstance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyCreateInstance** | [**BodyCreateInstance**](BodyCreateInstance.md) |  | 

### Return type

[**ResponseCreateInstanceModel**](ResponseCreateInstanceModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInstance

> DeleteInstance(ctx, instanceId).XAuthToken(xAuthToken).Execute()

Delete instance



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
	instanceId := "instanceId_example" // string | 인스턴스의 고유 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstanceAPI.DeleteInstance(context.Background(), instanceId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.DeleteInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInstanceRequest struct via the builder pattern


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


## GetInstance

> ResponseInstanceModel GetInstance(ctx, instanceId).XAuthToken(xAuthToken).Execute()

Get instance



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
	instanceId := "instanceId_example" // string | 인스턴스의 고유 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.GetInstance(context.Background(), instanceId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.GetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstance`: ResponseInstanceModel
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.GetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**ResponseInstanceModel**](ResponseInstanceModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInstances

> InstanceListModel ListInstances(ctx).XAuthToken(xAuthToken).Id(id).Name(name).VmState(vmState).FlavorName(flavorName).ImageName(imageName).PrivateIp(privateIp).PublicIp(publicIp).AvailabilityZone(availabilityZone).InstanceType(instanceType).Status(status).UserId(userId).Hostname(hostname).OsType(osType).IsHadoop(isHadoop).IsK8se(isK8se).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Limit(limit).Offset(offset).Execute()

List instances



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
	id := "id_example" // string | 조회할 인스턴스의 ID (optional)
	name := "name_example" // string | 인스턴스 이름 <br/> - 특정 이름을 포함하는 인스턴스를 검색 시 사용 (optional)
	vmState := "vmState_example" // string | 인스턴스의 상태 코드 <br/> - 전체 [인스턴스 상태 코드](https://docs.kakaocloud.com/service/bcs/vm/vm-main#instance-state-and-billing) 참고 (optional)
	flavorName := "flavorName_example" // string | 인스턴스 유형 이름  (optional)
	imageName := "imageName_example" // string | 인스턴스를 생성할 때 사용한 이미지 이름 (optional)
	privateIp := "privateIp_example" // string | 인스턴스에 할당된 프라이빗 IP 주소 (IPv4 형식) (optional)
	publicIp := "publicIp_example" // string | 인스턴스에 연결된 퍼블릭 IP 주소 (optional)
	availabilityZone := openapiclient.AvailabilityZone("kr-central-2-a") // AvailabilityZone | 인스턴스가 위치한 가용 영역 <br/> - `kr-central-2-a`: kr-central-2-a 가용 영역 <br/> - `kr-central-2-b`: kr-central-2-b 가용 영역 <br/> - `kr-central-2-c`: kr-central-2-c 가용 영역 (optional)
	instanceType := openapiclient.InstanceType("vm") // InstanceType | 인스턴스 유형 <br/> - `vm`: Virtual Machine 유형  <br/> - `bm`: Bare Metal Server 유형 <br/> - `gpu`: GPU 유형 (optional)
	status := "status_example" // string | 인스턴스의 상세 상태 정보<br/> - 내부적으로 정의된 상태 값 (optional)
	userId := "userId_example" // string | 해당 인스턴스를 생성한 사용자 ID (optional)
	hostname := "hostname_example" // string | 인스턴스의 호스트 이름(내부 DNS 이름 등) (optional)
	osType := "osType_example" // string | 운영체제 유형 (optional)
	isHadoop := true // bool | Hadoop 환경용으로 생성된 인스턴스인지 여부 (optional)
	isK8se := true // bool | Kubernetes Engine 환경용으로 생성된 인스턴스인지 여부 (optional)
	createdAt := "createdAt_example" // string | 리소스가 생성된 시간 <br/> - ISO_8601 형식  <br/> - UTC 기준 (optional)
	updatedAt := "updatedAt_example" // string | 리소스가 마지막으로 수정된 시간 <br/> - ISO_8601 형식  <br/> - UTC 기준 (optional)
	sortKeys := "sortKeys_example" // string | 정렬할 필드를 콤마(,)로 구분   (optional) (default to "created_at")
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`)   (optional) (default to "desc")
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional) (default to 20)
	offset := int32(56) // int32 | 조회 시작 위치 (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.ListInstances(context.Background()).XAuthToken(xAuthToken).Id(id).Name(name).VmState(vmState).FlavorName(flavorName).ImageName(imageName).PrivateIp(privateIp).PublicIp(publicIp).AvailabilityZone(availabilityZone).InstanceType(instanceType).Status(status).UserId(userId).Hostname(hostname).OsType(osType).IsHadoop(isHadoop).IsK8se(isK8se).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.ListInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListInstances`: InstanceListModel
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.ListInstances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **id** | **string** | 조회할 인스턴스의 ID | 
 **name** | **string** | 인스턴스 이름 &lt;br/&gt; - 특정 이름을 포함하는 인스턴스를 검색 시 사용 | 
 **vmState** | **string** | 인스턴스의 상태 코드 &lt;br/&gt; - 전체 [인스턴스 상태 코드](https://docs.kakaocloud.com/service/bcs/vm/vm-main#instance-state-and-billing) 참고 | 
 **flavorName** | **string** | 인스턴스 유형 이름  | 
 **imageName** | **string** | 인스턴스를 생성할 때 사용한 이미지 이름 | 
 **privateIp** | **string** | 인스턴스에 할당된 프라이빗 IP 주소 (IPv4 형식) | 
 **publicIp** | **string** | 인스턴스에 연결된 퍼블릭 IP 주소 | 
 **availabilityZone** | [**AvailabilityZone**](AvailabilityZone.md) | 인스턴스가 위치한 가용 영역 &lt;br/&gt; - &#x60;kr-central-2-a&#x60;: kr-central-2-a 가용 영역 &lt;br/&gt; - &#x60;kr-central-2-b&#x60;: kr-central-2-b 가용 영역 &lt;br/&gt; - &#x60;kr-central-2-c&#x60;: kr-central-2-c 가용 영역 | 
 **instanceType** | [**InstanceType**](InstanceType.md) | 인스턴스 유형 &lt;br/&gt; - &#x60;vm&#x60;: Virtual Machine 유형  &lt;br/&gt; - &#x60;bm&#x60;: Bare Metal Server 유형 &lt;br/&gt; - &#x60;gpu&#x60;: GPU 유형 | 
 **status** | **string** | 인스턴스의 상세 상태 정보&lt;br/&gt; - 내부적으로 정의된 상태 값 | 
 **userId** | **string** | 해당 인스턴스를 생성한 사용자 ID | 
 **hostname** | **string** | 인스턴스의 호스트 이름(내부 DNS 이름 등) | 
 **osType** | **string** | 운영체제 유형 | 
 **isHadoop** | **bool** | Hadoop 환경용으로 생성된 인스턴스인지 여부 | 
 **isK8se** | **bool** | Kubernetes Engine 환경용으로 생성된 인스턴스인지 여부 | 
 **createdAt** | **string** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
 **updatedAt** | **string** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
 **sortKeys** | **string** | 정렬할 필드를 콤마(,)로 구분   | [default to &quot;created_at&quot;]
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;)   | [default to &quot;desc&quot;]
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | [default to 20]
 **offset** | **int32** | 조회 시작 위치 | [default to 0]

### Return type

[**InstanceListModel**](InstanceListModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInstance

> InstanceModelResponse UpdateInstance(ctx, instanceId).XAuthToken(xAuthToken).BodyUpdateInstance(bodyUpdateInstance).Execute()

Update instance



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
	instanceId := "instanceId_example" // string | 인스턴스의 고유 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bodyUpdateInstance := *openapiclient.NewBodyUpdateInstance(*openapiclient.NewEditInstanceModel()) // BodyUpdateInstance | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InstanceAPI.UpdateInstance(context.Background(), instanceId).XAuthToken(xAuthToken).BodyUpdateInstance(bodyUpdateInstance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceAPI.UpdateInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInstance`: InstanceModelResponse
	fmt.Fprintf(os.Stdout, "Response from `InstanceAPI.UpdateInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | 인스턴스의 고유 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyUpdateInstance** | [**BodyUpdateInstance**](BodyUpdateInstance.md) |  | 

### Return type

[**InstanceModelResponse**](InstanceModelResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

