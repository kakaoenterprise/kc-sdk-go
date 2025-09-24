# \VPCSubnetAPI

All URIs are relative to *https://vpc.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSubnet**](VPCSubnetAPI.md#CreateSubnet) | **Post** /api/v1/subnets | Create subnet
[**DeleteSubnet**](VPCSubnetAPI.md#DeleteSubnet) | **Delete** /api/v1/subnets/{subnet_id} | Delete subnet
[**GetSubnet**](VPCSubnetAPI.md#GetSubnet) | **Get** /api/v1/subnets/{subnet_id} | Get subnet
[**ListSubnetSharedProjects**](VPCSubnetAPI.md#ListSubnetSharedProjects) | **Get** /api/v1/subnets/{subnet_id}/projects | List subnet shared projects
[**ListSubnets**](VPCSubnetAPI.md#ListSubnets) | **Get** /api/v1/subnets | List subnets
[**ShareSubnet**](VPCSubnetAPI.md#ShareSubnet) | **Post** /api/v1/subnets/{subnet_id}/projects/{project_id} | Share subnet
[**UnshareSubnet**](VPCSubnetAPI.md#UnshareSubnet) | **Delete** /api/v1/subnets/{subnet_id}/projects/{project_id} | Unshare subnet
[**UpdateSubnet**](VPCSubnetAPI.md#UpdateSubnet) | **Put** /api/v1/subnets/{subnet_id} | Update subnet



## CreateSubnet

> BnsVpcV1ApiCreateSubnetModelResponseSubnetModel CreateSubnet(ctx).XAuthToken(xAuthToken).BodyCreateSubnet(bodyCreateSubnet).Execute()

Create subnet



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
	bodyCreateSubnet := *openapiclient.NewBodyCreateSubnet(*openapiclient.NewCreateSubnetModel("Name_example", "VpcId_example", "CidrBlock_example", openapiclient.AvailabilityZone("kr-central-2-a"))) // BodyCreateSubnet | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPCSubnetAPI.CreateSubnet(context.Background()).XAuthToken(xAuthToken).BodyCreateSubnet(bodyCreateSubnet).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCSubnetAPI.CreateSubnet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSubnet`: BnsVpcV1ApiCreateSubnetModelResponseSubnetModel
	fmt.Fprintf(os.Stdout, "Response from `VPCSubnetAPI.CreateSubnet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubnetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyCreateSubnet** | [**BodyCreateSubnet**](BodyCreateSubnet.md) |  | 

### Return type

[**BnsVpcV1ApiCreateSubnetModelResponseSubnetModel**](BnsVpcV1ApiCreateSubnetModelResponseSubnetModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSubnet

> DeleteSubnet(ctx, subnetId).XAuthToken(xAuthToken).Execute()

Delete subnet



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
	subnetId := "subnetId_example" // string | 삭제할 서브넷 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VPCSubnetAPI.DeleteSubnet(context.Background(), subnetId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCSubnetAPI.DeleteSubnet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subnetId** | **string** | 삭제할 서브넷 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubnetRequest struct via the builder pattern


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


## GetSubnet

> BnsVpcV1ApiGetSubnetModelResponseSubnetModel GetSubnet(ctx, subnetId).XAuthToken(xAuthToken).Execute()

Get subnet



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
	subnetId := "subnetId_example" // string | 조회할 서브넷 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPCSubnetAPI.GetSubnet(context.Background(), subnetId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCSubnetAPI.GetSubnet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSubnet`: BnsVpcV1ApiGetSubnetModelResponseSubnetModel
	fmt.Fprintf(os.Stdout, "Response from `VPCSubnetAPI.GetSubnet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subnetId** | **string** | 조회할 서브넷 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubnetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**BnsVpcV1ApiGetSubnetModelResponseSubnetModel**](BnsVpcV1ApiGetSubnetModelResponseSubnetModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubnetSharedProjects

> ResponseSubnetSharedProjectListModel ListSubnetSharedProjects(ctx, subnetId).XAuthToken(xAuthToken).Execute()

List subnet shared projects



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
	subnetId := "subnetId_example" // string | 공유받은 프로젝트를 조회할 대상 서브넷 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPCSubnetAPI.ListSubnetSharedProjects(context.Background(), subnetId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCSubnetAPI.ListSubnetSharedProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSubnetSharedProjects`: ResponseSubnetSharedProjectListModel
	fmt.Fprintf(os.Stdout, "Response from `VPCSubnetAPI.ListSubnetSharedProjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subnetId** | **string** | 공유받은 프로젝트를 조회할 대상 서브넷 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSubnetSharedProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**ResponseSubnetSharedProjectListModel**](ResponseSubnetSharedProjectListModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSubnets

> SubnetListModel ListSubnets(ctx).XAuthToken(xAuthToken).Id(id).Name(name).AvailabilityZone(availabilityZone).ProvisioningStatus(provisioningStatus).OperatingStatus(operatingStatus).CidrBlock(cidrBlock).VpcId(vpcId).VpcName(vpcName).RouteTableId(routeTableId).RouteTableName(routeTableName).IsShared(isShared).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Limit(limit).Offset(offset).Execute()

List subnets



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
	id := "id_example" // string | 조회할 서브넷 ID (optional)
	name := "name_example" // string | 조회할 서브넷 이름 (optional)
	availabilityZone := openapiclient.AvailabilityZone("kr-central-2-a") // AvailabilityZone | 서브넷이 위치한 가용 영역<br/> - `kr-central-2-a`: kr-central-2-a 가용 영역 <br/> - `kr-central-2-b`: kr-central-2-b 가용 영역 <br/> - `kr-central-2-c`: kr-central-2-c 가용 영역 (optional)
	provisioningStatus := openapiclient.ProvisioningStatus("ACTIVE") // ProvisioningStatus | 프로비저닝 상태 <br/> - `ACTIVE`: 활성 상태 <br/> - `DELETED`: 삭제됨 <br/> - `ERROR`: 오류 발생 <br/> - `PENDING_CREATE`: 생성 대기 중 <br/> - `PENDING_UPDATE`: 업데이트 대기 중 <br/> - `PENDING_DELETE`: 삭제 대기 중 (optional)
	operatingStatus := openapiclient.SubnetOperatingStatus("ONLINE") // SubnetOperatingStatus | 운영 상태 <br/> - `ONLINE`: 온라인 상태 <br/> - `OFFLINE`: 오프라인 상태 <br/> - `IN_MAINTENANCE`: 점검 중 상태 <br/> - `ERROR`: 오류 발생 (optional)
	cidrBlock := "cidrBlock_example" // string | 서브넷의 IPv4 CIDR 블록 (예: `10.0.1.0/24`)  (optional)
	vpcId := "vpcId_example" // string | 연결된 VPC ID (optional)
	vpcName := "vpcName_example" // string | 연결된 VPC 이름 (optional)
	routeTableId := "routeTableId_example" // string | 연결된 라우팅 테이블 ID  (optional)
	routeTableName := "routeTableName_example" // string | 연결된 라우팅 테이블 이름  (optional)
	isShared := true // bool | 공유 여부 (optional)
	createdAt := "createdAt_example" // string | 리소스가 생성된 시간 <br/> - ISO_8601 형식  <br/> - UTC 기준 (optional)
	updatedAt := "updatedAt_example" // string | 리소스가 마지막으로 수정된 시간 <br/> - ISO_8601 형식  <br/> - UTC 기준 (optional)
	sortKeys := "sortKeys_example" // string | 정렬할 필드를 콤마(,)로 구분     (optional) (default to "created_at")
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`)   (optional) (default to "desc")
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional) (default to 20)
	offset := int32(56) // int32 | 조회 시작 위치 (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPCSubnetAPI.ListSubnets(context.Background()).XAuthToken(xAuthToken).Id(id).Name(name).AvailabilityZone(availabilityZone).ProvisioningStatus(provisioningStatus).OperatingStatus(operatingStatus).CidrBlock(cidrBlock).VpcId(vpcId).VpcName(vpcName).RouteTableId(routeTableId).RouteTableName(routeTableName).IsShared(isShared).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCSubnetAPI.ListSubnets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSubnets`: SubnetListModel
	fmt.Fprintf(os.Stdout, "Response from `VPCSubnetAPI.ListSubnets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSubnetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **id** | **string** | 조회할 서브넷 ID | 
 **name** | **string** | 조회할 서브넷 이름 | 
 **availabilityZone** | [**AvailabilityZone**](AvailabilityZone.md) | 서브넷이 위치한 가용 영역&lt;br/&gt; - &#x60;kr-central-2-a&#x60;: kr-central-2-a 가용 영역 &lt;br/&gt; - &#x60;kr-central-2-b&#x60;: kr-central-2-b 가용 영역 &lt;br/&gt; - &#x60;kr-central-2-c&#x60;: kr-central-2-c 가용 영역 | 
 **provisioningStatus** | [**ProvisioningStatus**](ProvisioningStatus.md) | 프로비저닝 상태 &lt;br/&gt; - &#x60;ACTIVE&#x60;: 활성 상태 &lt;br/&gt; - &#x60;DELETED&#x60;: 삭제됨 &lt;br/&gt; - &#x60;ERROR&#x60;: 오류 발생 &lt;br/&gt; - &#x60;PENDING_CREATE&#x60;: 생성 대기 중 &lt;br/&gt; - &#x60;PENDING_UPDATE&#x60;: 업데이트 대기 중 &lt;br/&gt; - &#x60;PENDING_DELETE&#x60;: 삭제 대기 중 | 
 **operatingStatus** | [**SubnetOperatingStatus**](SubnetOperatingStatus.md) | 운영 상태 &lt;br/&gt; - &#x60;ONLINE&#x60;: 온라인 상태 &lt;br/&gt; - &#x60;OFFLINE&#x60;: 오프라인 상태 &lt;br/&gt; - &#x60;IN_MAINTENANCE&#x60;: 점검 중 상태 &lt;br/&gt; - &#x60;ERROR&#x60;: 오류 발생 | 
 **cidrBlock** | **string** | 서브넷의 IPv4 CIDR 블록 (예: &#x60;10.0.1.0/24&#x60;)  | 
 **vpcId** | **string** | 연결된 VPC ID | 
 **vpcName** | **string** | 연결된 VPC 이름 | 
 **routeTableId** | **string** | 연결된 라우팅 테이블 ID  | 
 **routeTableName** | **string** | 연결된 라우팅 테이블 이름  | 
 **isShared** | **bool** | 공유 여부 | 
 **createdAt** | **string** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
 **updatedAt** | **string** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
 **sortKeys** | **string** | 정렬할 필드를 콤마(,)로 구분     | [default to &quot;created_at&quot;]
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;)   | [default to &quot;desc&quot;]
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | [default to 20]
 **offset** | **int32** | 조회 시작 위치 | [default to 0]

### Return type

[**SubnetListModel**](SubnetListModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShareSubnet

> interface{} ShareSubnet(ctx, subnetId, projectId).XAuthToken(xAuthToken).Execute()

Share subnet



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
	subnetId := "subnetId_example" // string | 공유할 서브넷 ID
	projectId := "projectId_example" // string | 서브넷을 공유받을 대상 프로젝트 ID 
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPCSubnetAPI.ShareSubnet(context.Background(), subnetId, projectId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCSubnetAPI.ShareSubnet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ShareSubnet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `VPCSubnetAPI.ShareSubnet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subnetId** | **string** | 공유할 서브넷 ID | 
**projectId** | **string** | 서브넷을 공유받을 대상 프로젝트 ID  | 

### Other Parameters

Other parameters are passed through a pointer to a apiShareSubnetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

**interface{}**

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnshareSubnet

> UnshareSubnet(ctx, subnetId, projectId).XAuthToken(xAuthToken).Execute()

Unshare subnet



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
	subnetId := "subnetId_example" // string | 공유 해제할 서브넷 ID
	projectId := "projectId_example" // string | 공유를 해제할 대상 프로젝트 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VPCSubnetAPI.UnshareSubnet(context.Background(), subnetId, projectId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCSubnetAPI.UnshareSubnet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subnetId** | **string** | 공유 해제할 서브넷 ID | 
**projectId** | **string** | 공유를 해제할 대상 프로젝트 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnshareSubnetRequest struct via the builder pattern


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


## UpdateSubnet

> BnsVpcV1ApiUpdateSubnetModelResponseSubnetModel UpdateSubnet(ctx, subnetId).XAuthToken(xAuthToken).BodyUpdateSubnet(bodyUpdateSubnet).Execute()

Update subnet



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
	subnetId := "subnetId_example" // string | 이름을 수정할 서브넷의 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bodyUpdateSubnet := *openapiclient.NewBodyUpdateSubnet(*openapiclient.NewEditSubnetModel("Name_example")) // BodyUpdateSubnet | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPCSubnetAPI.UpdateSubnet(context.Background(), subnetId).XAuthToken(xAuthToken).BodyUpdateSubnet(bodyUpdateSubnet).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCSubnetAPI.UpdateSubnet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSubnet`: BnsVpcV1ApiUpdateSubnetModelResponseSubnetModel
	fmt.Fprintf(os.Stdout, "Response from `VPCSubnetAPI.UpdateSubnet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subnetId** | **string** | 이름을 수정할 서브넷의 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSubnetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyUpdateSubnet** | [**BodyUpdateSubnet**](BodyUpdateSubnet.md) |  | 

### Return type

[**BnsVpcV1ApiUpdateSubnetModelResponseSubnetModel**](BnsVpcV1ApiUpdateSubnetModelResponseSubnetModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

