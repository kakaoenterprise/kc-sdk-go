# \VPCRouteTableAPI

All URIs are relative to *https://vpc.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRouteTable**](VPCRouteTableAPI.md#CreateRouteTable) | **Post** /api/v1/route-tables | Create route table
[**DeleteRouteTable**](VPCRouteTableAPI.md#DeleteRouteTable) | **Delete** /api/v1/route-tables/{route_table_id} | Delete route table
[**GetRouteTable**](VPCRouteTableAPI.md#GetRouteTable) | **Get** /api/v1/route-tables/{route_table_id} | Get route table
[**ListRouteTables**](VPCRouteTableAPI.md#ListRouteTables) | **Get** /api/v1/route-tables | List route tables
[**SetMainRouteTable**](VPCRouteTableAPI.md#SetMainRouteTable) | **Put** /api/v1/route-tables/{route_table_id}/main | Set main route table



## CreateRouteTable

> BnsVpcV1ApiCreateRouteTableModelResponseRouteTableModel CreateRouteTable(ctx).XAuthToken(xAuthToken).BodyCreateRouteTable(bodyCreateRouteTable).Execute()

Create route table



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
	bodyCreateRouteTable := *openapiclient.NewBodyCreateRouteTable(*openapiclient.NewCreateRouteTableModel("Name_example", "VpcId_example")) // BodyCreateRouteTable | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPCRouteTableAPI.CreateRouteTable(context.Background()).XAuthToken(xAuthToken).BodyCreateRouteTable(bodyCreateRouteTable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCRouteTableAPI.CreateRouteTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRouteTable`: BnsVpcV1ApiCreateRouteTableModelResponseRouteTableModel
	fmt.Fprintf(os.Stdout, "Response from `VPCRouteTableAPI.CreateRouteTable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRouteTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyCreateRouteTable** | [**BodyCreateRouteTable**](BodyCreateRouteTable.md) |  | 

### Return type

[**BnsVpcV1ApiCreateRouteTableModelResponseRouteTableModel**](BnsVpcV1ApiCreateRouteTableModelResponseRouteTableModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRouteTable

> DeleteRouteTable(ctx, routeTableId).XAuthToken(xAuthToken).Execute()

Delete route table



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
	routeTableId := "routeTableId_example" // string | 삭제할 라우팅 테이블 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VPCRouteTableAPI.DeleteRouteTable(context.Background(), routeTableId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCRouteTableAPI.DeleteRouteTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeTableId** | **string** | 삭제할 라우팅 테이블 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRouteTableRequest struct via the builder pattern


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


## GetRouteTable

> BnsVpcV1ApiGetRouteTableModelResponseRouteTableModel GetRouteTable(ctx, routeTableId).XAuthToken(xAuthToken).Execute()

Get route table



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
	routeTableId := "routeTableId_example" // string | 조회할 라우팅 테이블 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPCRouteTableAPI.GetRouteTable(context.Background(), routeTableId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCRouteTableAPI.GetRouteTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRouteTable`: BnsVpcV1ApiGetRouteTableModelResponseRouteTableModel
	fmt.Fprintf(os.Stdout, "Response from `VPCRouteTableAPI.GetRouteTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeTableId** | **string** | 조회할 라우팅 테이블 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**BnsVpcV1ApiGetRouteTableModelResponseRouteTableModel**](BnsVpcV1ApiGetRouteTableModelResponseRouteTableModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRouteTables

> RouteTableListModel ListRouteTables(ctx).XAuthToken(xAuthToken).Id(id).Name(name).ProvisioningStatus(provisioningStatus).VpcId(vpcId).VpcName(vpcName).VpcProvisioningStatus(vpcProvisioningStatus).SubnetName(subnetName).SubnetId(subnetId).AssociationCount(associationCount).Destination(destination).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Limit(limit).Offset(offset).Execute()

List route tables



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
	id := "id_example" // string | 라우팅 테이블 ID  (optional)
	name := "name_example" // string | 연결된 VPC 이름 (optional)
	provisioningStatus := openapiclient.ProvisioningStatus("ACTIVE") // ProvisioningStatus | 라우팅 테이블의 프로비저닝 상태 <br/> - `ACTIVE`: 활성 상태 <br/> - `DELETED`: 삭제됨 <br/> - `ERROR`: 오류 발생 <br/> - `PENDING_CREATE`: 생성 대기 중 <br/> - `PENDING_UPDATE`: 업데이트 대기 중 <br/> - `PENDING_DELETE`: 삭제 대기 중 (optional)
	vpcId := "vpcId_example" // string | 연결된 VPC ID (optional)
	vpcName := "vpcName_example" // string | 연결된 VPC 이름 (optional)
	vpcProvisioningStatus := openapiclient.ProvisioningStatus("ACTIVE") // ProvisioningStatus | VPC의 프로비저닝 상태 <br/> - `ACTIVE`: 활성 상태 <br/> - `DELETED`: 삭제됨 <br/> - `ERROR`: 오류 발생 <br/> - `PENDING_CREATE`: 생성 대기 중 <br/> - `PENDING_UPDATE`: 업데이트 대기 중 <br/> - `PENDING_DELETE`: 삭제 대기 중 (optional)
	subnetName := "subnetName_example" // string | 연결된 서브넷의 이름 (optional)
	subnetId := "subnetId_example" // string | 연결된 서브넷 ID (optional)
	associationCount := "associationCount_example" // string | 연결된 서브넷 수 (optional)
	destination := "destination_example" // string | 라우팅 경로  (optional)
	createdAt := "createdAt_example" // string | 리소스가 생성된 시간 <br/> - ISO_8601 형식  <br/> - UTC 기준 (optional)
	updatedAt := "updatedAt_example" // string | 리소스가 마지막으로 수정된 시간 <br/> - ISO_8601 형식  <br/> - UTC 기준 (optional)
	sortKeys := "sortKeys_example" // string | 정렬할 필드를 콤마(,)로 구분     (optional) (default to "created_at")
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`)   (optional) (default to "desc")
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional) (default to 20)
	offset := int32(56) // int32 | 조회 시작 위치 (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPCRouteTableAPI.ListRouteTables(context.Background()).XAuthToken(xAuthToken).Id(id).Name(name).ProvisioningStatus(provisioningStatus).VpcId(vpcId).VpcName(vpcName).VpcProvisioningStatus(vpcProvisioningStatus).SubnetName(subnetName).SubnetId(subnetId).AssociationCount(associationCount).Destination(destination).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCRouteTableAPI.ListRouteTables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRouteTables`: RouteTableListModel
	fmt.Fprintf(os.Stdout, "Response from `VPCRouteTableAPI.ListRouteTables`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRouteTablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **id** | **string** | 라우팅 테이블 ID  | 
 **name** | **string** | 연결된 VPC 이름 | 
 **provisioningStatus** | [**ProvisioningStatus**](ProvisioningStatus.md) | 라우팅 테이블의 프로비저닝 상태 &lt;br/&gt; - &#x60;ACTIVE&#x60;: 활성 상태 &lt;br/&gt; - &#x60;DELETED&#x60;: 삭제됨 &lt;br/&gt; - &#x60;ERROR&#x60;: 오류 발생 &lt;br/&gt; - &#x60;PENDING_CREATE&#x60;: 생성 대기 중 &lt;br/&gt; - &#x60;PENDING_UPDATE&#x60;: 업데이트 대기 중 &lt;br/&gt; - &#x60;PENDING_DELETE&#x60;: 삭제 대기 중 | 
 **vpcId** | **string** | 연결된 VPC ID | 
 **vpcName** | **string** | 연결된 VPC 이름 | 
 **vpcProvisioningStatus** | [**ProvisioningStatus**](ProvisioningStatus.md) | VPC의 프로비저닝 상태 &lt;br/&gt; - &#x60;ACTIVE&#x60;: 활성 상태 &lt;br/&gt; - &#x60;DELETED&#x60;: 삭제됨 &lt;br/&gt; - &#x60;ERROR&#x60;: 오류 발생 &lt;br/&gt; - &#x60;PENDING_CREATE&#x60;: 생성 대기 중 &lt;br/&gt; - &#x60;PENDING_UPDATE&#x60;: 업데이트 대기 중 &lt;br/&gt; - &#x60;PENDING_DELETE&#x60;: 삭제 대기 중 | 
 **subnetName** | **string** | 연결된 서브넷의 이름 | 
 **subnetId** | **string** | 연결된 서브넷 ID | 
 **associationCount** | **string** | 연결된 서브넷 수 | 
 **destination** | **string** | 라우팅 경로  | 
 **createdAt** | **string** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
 **updatedAt** | **string** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
 **sortKeys** | **string** | 정렬할 필드를 콤마(,)로 구분     | [default to &quot;created_at&quot;]
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;)   | [default to &quot;desc&quot;]
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | [default to 20]
 **offset** | **int32** | 조회 시작 위치 | [default to 0]

### Return type

[**RouteTableListModel**](RouteTableListModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetMainRouteTable

> interface{} SetMainRouteTable(ctx, routeTableId).XAuthToken(xAuthToken).Execute()

Set main route table



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
	routeTableId := "routeTableId_example" // string | 기본(main)으로 설정할 라우팅 테이블 ID 
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPCRouteTableAPI.SetMainRouteTable(context.Background(), routeTableId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCRouteTableAPI.SetMainRouteTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetMainRouteTable`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `VPCRouteTableAPI.SetMainRouteTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeTableId** | **string** | 기본(main)으로 설정할 라우팅 테이블 ID  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetMainRouteTableRequest struct via the builder pattern


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

