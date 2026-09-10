# \VPCRouteTableAPI

All URIs are relative to *https://vpc.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRoute**](VPCRouteTableAPI.md#AddRoute) | **Post** /api/v1/route-tables/{route_table_id}/routes | Add route
[**CreateRouteTable**](VPCRouteTableAPI.md#CreateRouteTable) | **Post** /api/v1/route-tables | Create route table
[**DeleteRoute**](VPCRouteTableAPI.md#DeleteRoute) | **Delete** /api/v1/route-tables/{route_table_id}/routes/{route_id} | Delete route
[**DeleteRouteTable**](VPCRouteTableAPI.md#DeleteRouteTable) | **Delete** /api/v1/route-tables/{route_table_id} | Delete route table
[**GetRouteTable**](VPCRouteTableAPI.md#GetRouteTable) | **Get** /api/v1/route-tables/{route_table_id} | Get route table
[**ListRouteTableAssociations**](VPCRouteTableAPI.md#ListRouteTableAssociations) | **Get** /api/v1/route-tables/{route_table_id}/associations | List route table associations 
[**ListRouteTables**](VPCRouteTableAPI.md#ListRouteTables) | **Get** /api/v1/route-tables | List route tables
[**SetMainRouteTable**](VPCRouteTableAPI.md#SetMainRouteTable) | **Put** /api/v1/route-tables/{route_table_id}/main | Set main route table
[**UpdateRoute**](VPCRouteTableAPI.md#UpdateRoute) | **Put** /api/v1/route-tables/{route_table_id}/routes/{route_id} | Update route
[**UpdateRouteTable**](VPCRouteTableAPI.md#UpdateRouteTable) | **Put** /api/v1/route-tables/{route_table_id} | Update route table
[**UpdateRouteTableAssociation**](VPCRouteTableAPI.md#UpdateRouteTableAssociation) | **Put** /api/v1/route-tables/{route_table_id}/associations/{association_id} | Update route table association



## AddRoute

> AddRouteResponse AddRoute(ctx, routeTableId).XAuthToken(xAuthToken).AddRouteRequest(addRouteRequest).Execute()

Add route



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/vpc"
)

func main() {
	routeTableId := "routeTableId_example" // string | 라우팅을 추가할 대상 라우팅 테이블의 ID <br/>- [List route tables](https://docs.kakaocloud.com/openapi/networking/vpc/list-route-tables)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	addRouteRequest := *openapiclient.NewAddRouteRequest(*openapiclient.NewAddRoute(openapiclient.RouteType("igw"), "TargetId_example", "Destination_example")) // AddRouteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPCRouteTableAPI.AddRoute(context.Background(), routeTableId).XAuthToken(xAuthToken).AddRouteRequest(addRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCRouteTableAPI.AddRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddRoute`: AddRouteResponse
	fmt.Fprintf(os.Stdout, "Response from `VPCRouteTableAPI.AddRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeTableId** | **string** | 라우팅을 추가할 대상 라우팅 테이블의 ID &lt;br/&gt;- [List route tables](https://docs.kakaocloud.com/openapi/networking/vpc/list-route-tables)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **addRouteRequest** | [**AddRouteRequest**](AddRouteRequest.md) |  | 

### Return type

[**AddRouteResponse**](AddRouteResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRouteTable

> CreateRouteTableResponse CreateRouteTable(ctx).XAuthToken(xAuthToken).CreateRouteTableRequest(createRouteTableRequest).Execute()

Create route table



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/vpc"
)

func main() {
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	createRouteTableRequest := *openapiclient.NewCreateRouteTableRequest(*openapiclient.NewCreateRouteTable("Name_example", "VpcId_example")) // CreateRouteTableRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPCRouteTableAPI.CreateRouteTable(context.Background()).XAuthToken(xAuthToken).CreateRouteTableRequest(createRouteTableRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCRouteTableAPI.CreateRouteTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRouteTable`: CreateRouteTableResponse
	fmt.Fprintf(os.Stdout, "Response from `VPCRouteTableAPI.CreateRouteTable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRouteTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **createRouteTableRequest** | [**CreateRouteTableRequest**](CreateRouteTableRequest.md) |  | 

### Return type

[**CreateRouteTableResponse**](CreateRouteTableResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRoute

> DeleteRoute(ctx, routeTableId, routeId).XAuthToken(xAuthToken).Execute()

Delete route



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/vpc"
)

func main() {
	routeTableId := "routeTableId_example" // string | 삭제할 라우팅 테이블 ID
	routeId := "routeId_example" // string | 삭제할 라우팅 ID <br/>- [List route tables](https://docs.kakaocloud.com/openapi/networking/vpc/list-route-tables)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VPCRouteTableAPI.DeleteRoute(context.Background(), routeTableId, routeId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCRouteTableAPI.DeleteRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeTableId** | **string** | 삭제할 라우팅 테이블 ID | 
**routeId** | **string** | 삭제할 라우팅 ID &lt;br/&gt;- [List route tables](https://docs.kakaocloud.com/openapi/networking/vpc/list-route-tables)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRouteRequest struct via the builder pattern


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
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/vpc"
)

func main() {
	routeTableId := "routeTableId_example" // string | 삭제할 라우팅 테이블 ID <br/>- [List route tables](https://docs.kakaocloud.com/openapi/networking/vpc/list-route-tables)에서 확인
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
**routeTableId** | **string** | 삭제할 라우팅 테이블 ID &lt;br/&gt;- [List route tables](https://docs.kakaocloud.com/openapi/networking/vpc/list-route-tables)에서 확인 | 

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

> GetRouteTableResponse GetRouteTable(ctx, routeTableId).XAuthToken(xAuthToken).Execute()

Get route table



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/vpc"
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
	// response from `GetRouteTable`: GetRouteTableResponse
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

[**GetRouteTableResponse**](GetRouteTableResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRouteTableAssociations

> ListRouteTableAssociationsResponse ListRouteTableAssociations(ctx, routeTableId).XAuthToken(xAuthToken).Execute()

List route table associations 



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/vpc"
)

func main() {
	routeTableId := "routeTableId_example" // string | 라우팅 테이블의 고유 ID <br/>- [List route tables](https://docs.kakaocloud.com/openapi/networking/vpc/list-route-tables)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPCRouteTableAPI.ListRouteTableAssociations(context.Background(), routeTableId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCRouteTableAPI.ListRouteTableAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRouteTableAssociations`: ListRouteTableAssociationsResponse
	fmt.Fprintf(os.Stdout, "Response from `VPCRouteTableAPI.ListRouteTableAssociations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeTableId** | **string** | 라우팅 테이블의 고유 ID &lt;br/&gt;- [List route tables](https://docs.kakaocloud.com/openapi/networking/vpc/list-route-tables)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiListRouteTableAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**ListRouteTableAssociationsResponse**](ListRouteTableAssociationsResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRouteTables

> ListRouteTablesResponse ListRouteTables(ctx).XAuthToken(xAuthToken).Id(id).Name(name).ProvisioningStatus(provisioningStatus).VpcId(vpcId).VpcName(vpcName).VpcProvisioningStatus(vpcProvisioningStatus).SubnetName(subnetName).SubnetId(subnetId).AssociationCount(associationCount).Destination(destination).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Offset(offset).Limit(limit).Execute()

List route tables



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/vpc"
)

func main() {
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	id := "id_example" // string | 라우팅 테이블 ID (optional)
	name := "name_example" // string | 연결된 VPC 이름 (optional)
	provisioningStatus := openapiclient.ProvisioningStatus("ACTIVE") // ProvisioningStatus | 라우팅 테이블의 프로비저닝 상태 (optional)
	vpcId := "vpcId_example" // string | 연결된 VPC ID (optional)
	vpcName := "vpcName_example" // string | 연결된 VPC 이름 (optional)
	vpcProvisioningStatus := openapiclient.ProvisioningStatus("ACTIVE") // ProvisioningStatus | VPC의 프로비저닝 상태 (optional)
	subnetName := "subnetName_example" // string | 연결된 서브넷의 이름 (optional)
	subnetId := "subnetId_example" // string | 연결된 서브넷 ID (optional)
	associationCount := int32(56) // int32 | 연결된 서브넷 수 (optional)
	destination := "destination_example" // string | 라우팅 테이블에 대한 설명 (optional)
	createdAt := "createdAt_example" // string | 리소스가 생성된 시간 <br/> - ISO 8601 형식 <br/> - UTC 기준 (optional)
	updatedAt := "updatedAt_example" // string | 리소스가 마지막으로 수정된 시간 <br/> - ISO 8601 형식 <br/> - UTC 기준 (optional)
	sortKeys := "sortKeys_example" // string | 정렬할 필드를 콤마(,)로 구분 (optional)
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`)  (optional)
	offset := int32(56) // int32 | 조회 시작 위치 (optional)
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPCRouteTableAPI.ListRouteTables(context.Background()).XAuthToken(xAuthToken).Id(id).Name(name).ProvisioningStatus(provisioningStatus).VpcId(vpcId).VpcName(vpcName).VpcProvisioningStatus(vpcProvisioningStatus).SubnetName(subnetName).SubnetId(subnetId).AssociationCount(associationCount).Destination(destination).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCRouteTableAPI.ListRouteTables``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRouteTables`: ListRouteTablesResponse
	fmt.Fprintf(os.Stdout, "Response from `VPCRouteTableAPI.ListRouteTables`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRouteTablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **id** | **string** | 라우팅 테이블 ID | 
 **name** | **string** | 연결된 VPC 이름 | 
 **provisioningStatus** | [**ProvisioningStatus**](ProvisioningStatus.md) | 라우팅 테이블의 프로비저닝 상태 | 
 **vpcId** | **string** | 연결된 VPC ID | 
 **vpcName** | **string** | 연결된 VPC 이름 | 
 **vpcProvisioningStatus** | [**ProvisioningStatus**](ProvisioningStatus.md) | VPC의 프로비저닝 상태 | 
 **subnetName** | **string** | 연결된 서브넷의 이름 | 
 **subnetId** | **string** | 연결된 서브넷 ID | 
 **associationCount** | **int32** | 연결된 서브넷 수 | 
 **destination** | **string** | 라우팅 테이블에 대한 설명 | 
 **createdAt** | **string** | 리소스가 생성된 시간 &lt;br/&gt; - ISO 8601 형식 &lt;br/&gt; - UTC 기준 | 
 **updatedAt** | **string** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO 8601 형식 &lt;br/&gt; - UTC 기준 | 
 **sortKeys** | **string** | 정렬할 필드를 콤마(,)로 구분 | 
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;)  | 
 **offset** | **int32** | 조회 시작 위치 | 
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | 

### Return type

[**ListRouteTablesResponse**](ListRouteTablesResponse.md)

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
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/vpc"
)

func main() {
	routeTableId := "routeTableId_example" // string | 기본(main)으로 설정할 라우팅 테이블 ID <br/>- [List route tables](https://docs.kakaocloud.com/openapi/networking/vpc/list-route-tables)에서 확인
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
**routeTableId** | **string** | 기본(main)으로 설정할 라우팅 테이블 ID &lt;br/&gt;- [List route tables](https://docs.kakaocloud.com/openapi/networking/vpc/list-route-tables)에서 확인 | 

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


## UpdateRoute

> UpdateRouteResponse UpdateRoute(ctx, routeTableId, routeId).XAuthToken(xAuthToken).UpdateRouteRequest(updateRouteRequest).Execute()

Update route



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/vpc"
)

func main() {
	routeTableId := "routeTableId_example" // string | 라우팅 테이블의 고유 ID
	routeId := "routeId_example" // string | 라우팅의 고유 ID <br/>- [List route tables](https://docs.kakaocloud.com/openapi/networking/vpc/list-route-tables)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	updateRouteRequest := *openapiclient.NewUpdateRouteRequest(*openapiclient.NewUpdateRoute(openapiclient.RouteType("igw"), "TargetId_example", "Destination_example")) // UpdateRouteRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPCRouteTableAPI.UpdateRoute(context.Background(), routeTableId, routeId).XAuthToken(xAuthToken).UpdateRouteRequest(updateRouteRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCRouteTableAPI.UpdateRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRoute`: UpdateRouteResponse
	fmt.Fprintf(os.Stdout, "Response from `VPCRouteTableAPI.UpdateRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeTableId** | **string** | 라우팅 테이블의 고유 ID | 
**routeId** | **string** | 라우팅의 고유 ID &lt;br/&gt;- [List route tables](https://docs.kakaocloud.com/openapi/networking/vpc/list-route-tables)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **updateRouteRequest** | [**UpdateRouteRequest**](UpdateRouteRequest.md) |  | 

### Return type

[**UpdateRouteResponse**](UpdateRouteResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRouteTable

> UpdateRouteTableResponse UpdateRouteTable(ctx, routeTableId).XAuthToken(xAuthToken).UpdateRouteTableRequest(updateRouteTableRequest).Execute()

Update route table



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/vpc"
)

func main() {
	routeTableId := "routeTableId_example" // string | 라우팅 테이블의 고유 ID <br/>- [List route tables](https://docs.kakaocloud.com/openapi/networking/vpc/list-route-tables)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	updateRouteTableRequest := *openapiclient.NewUpdateRouteTableRequest(*openapiclient.NewUpdateRouteTable()) // UpdateRouteTableRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPCRouteTableAPI.UpdateRouteTable(context.Background(), routeTableId).XAuthToken(xAuthToken).UpdateRouteTableRequest(updateRouteTableRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCRouteTableAPI.UpdateRouteTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRouteTable`: UpdateRouteTableResponse
	fmt.Fprintf(os.Stdout, "Response from `VPCRouteTableAPI.UpdateRouteTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeTableId** | **string** | 라우팅 테이블의 고유 ID &lt;br/&gt;- [List route tables](https://docs.kakaocloud.com/openapi/networking/vpc/list-route-tables)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRouteTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **updateRouteTableRequest** | [**UpdateRouteTableRequest**](UpdateRouteTableRequest.md) |  | 

### Return type

[**UpdateRouteTableResponse**](UpdateRouteTableResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRouteTableAssociation

> interface{} UpdateRouteTableAssociation(ctx, routeTableId, associationId).XAuthToken(xAuthToken).UpdateRouteTableAssociationRequest(updateRouteTableAssociationRequest).Execute()

Update route table association



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/vpc"
)

func main() {
	routeTableId := "routeTableId_example" // string | 현재 연결에 매핑된 라우팅 테이블 ID
	associationId := "associationId_example" // string | 수정할 라우팅 테이블과 서브넷 연결 정보를 식별하는 ID <br/>- [List route table associations](https://docs.kakaocloud.com/openapi/networking/vpc/list-route-table-associations)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	updateRouteTableAssociationRequest := *openapiclient.NewUpdateRouteTableAssociationRequest(*openapiclient.NewUpdateRouteTableAssociation("TargetRouteTableId_example")) // UpdateRouteTableAssociationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPCRouteTableAPI.UpdateRouteTableAssociation(context.Background(), routeTableId, associationId).XAuthToken(xAuthToken).UpdateRouteTableAssociationRequest(updateRouteTableAssociationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCRouteTableAPI.UpdateRouteTableAssociation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRouteTableAssociation`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `VPCRouteTableAPI.UpdateRouteTableAssociation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routeTableId** | **string** | 현재 연결에 매핑된 라우팅 테이블 ID | 
**associationId** | **string** | 수정할 라우팅 테이블과 서브넷 연결 정보를 식별하는 ID &lt;br/&gt;- [List route table associations](https://docs.kakaocloud.com/openapi/networking/vpc/list-route-table-associations)에서 확인 | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRouteTableAssociationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **updateRouteTableAssociationRequest** | [**UpdateRouteTableAssociationRequest**](UpdateRouteTableAssociationRequest.md) |  | 

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

