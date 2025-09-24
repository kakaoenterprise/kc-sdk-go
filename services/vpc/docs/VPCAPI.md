# \VPCAPI

All URIs are relative to *https://vpc.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVpc**](VPCAPI.md#CreateVpc) | **Post** /api/v1/vpcs | Create VPC 
[**DeleteVpc**](VPCAPI.md#DeleteVpc) | **Delete** /api/v1/vpcs/{vpc_id} | Delete VPC
[**GetVpc**](VPCAPI.md#GetVpc) | **Get** /api/v1/vpcs/{vpc_id} | Get VPC
[**ListVpcs**](VPCAPI.md#ListVpcs) | **Get** /api/v1/vpcs | List VPCs
[**PutBnsVpc**](VPCAPI.md#PutBnsVpc) | **Put** /api/v1/vpcs/{vpc_id} | Update VPC



## CreateVpc

> BnsVpcV1ApiCreateVpcModelResponseVPCModel CreateVpc(ctx).XAuthToken(xAuthToken).BodyCreateVpc(bodyCreateVpc).Execute()

Create VPC 



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
	bodyCreateVpc := *openapiclient.NewBodyCreateVpc(*openapiclient.NewCreateVPCModel("Name_example", "CidrBlock_example")) // BodyCreateVpc | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPCAPI.CreateVpc(context.Background()).XAuthToken(xAuthToken).BodyCreateVpc(bodyCreateVpc).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCAPI.CreateVpc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVpc`: BnsVpcV1ApiCreateVpcModelResponseVPCModel
	fmt.Fprintf(os.Stdout, "Response from `VPCAPI.CreateVpc`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVpcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyCreateVpc** | [**BodyCreateVpc**](BodyCreateVpc.md) |  | 

### Return type

[**BnsVpcV1ApiCreateVpcModelResponseVPCModel**](BnsVpcV1ApiCreateVpcModelResponseVPCModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVpc

> DeleteVpc(ctx, vpcId).XAuthToken(xAuthToken).Execute()

Delete VPC



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
	vpcId := "vpcId_example" // string | 삭제할 VPC ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.VPCAPI.DeleteVpc(context.Background(), vpcId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCAPI.DeleteVpc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpcId** | **string** | 삭제할 VPC ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVpcRequest struct via the builder pattern


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


## GetVpc

> BnsVpcV1ApiGetVpcModelResponseVPCModel GetVpc(ctx, vpcId).XAuthToken(xAuthToken).Execute()

Get VPC



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
	vpcId := "vpcId_example" // string | 조회할 VPC ID <br/>- [List VPCs](https://docs.kakaocloud.com/openapi/bns/vpc/list-vpcs) 참고
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPCAPI.GetVpc(context.Background(), vpcId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCAPI.GetVpc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVpc`: BnsVpcV1ApiGetVpcModelResponseVPCModel
	fmt.Fprintf(os.Stdout, "Response from `VPCAPI.GetVpc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpcId** | **string** | 조회할 VPC ID &lt;br/&gt;- [List VPCs](https://docs.kakaocloud.com/openapi/bns/vpc/list-vpcs) 참고 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVpcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**BnsVpcV1ApiGetVpcModelResponseVPCModel**](BnsVpcV1ApiGetVpcModelResponseVPCModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVpcs

> VPCListModel ListVpcs(ctx).XAuthToken(xAuthToken).Id(id).Name(name).CidrBlock(cidrBlock).ProvisioningStatus(provisioningStatus).IsDefault(isDefault).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Limit(limit).Offset(offset).Execute()

List VPCs



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
	id := "id_example" // string | 조회할 VPC ID <br/>- [List VPCs](https://docs.kakaocloud.com/openapi/bns/vpc/list-vpcs) 참고 (optional)
	name := "name_example" // string | 조회할 VPC 이름 (optional)
	cidrBlock := "cidrBlock_example" // string | VPC의 IPv4 CIDR 블록 (예: `10.0.0.0/16`)  (optional)
	provisioningStatus := openapiclient.ProvisioningStatus("ACTIVE") // ProvisioningStatus | 프로비저닝 상태 <br/> - `ACTIVE`: 활성 상태 <br/> - `DELETED`: 삭제됨 <br/> - `ERROR`: 오류 발생 <br/> - `PENDING_CREATE`: 생성 대기 중 <br/> - `PENDING_UPDATE`: 업데이트 대기 중 <br/> - `PENDING_DELETE`: 삭제 대기 중 (optional)
	isDefault := true // bool | 기본 VPC 여부 (optional)
	createdAt := "createdAt_example" // string | 리소스가 생성된 시간 <br/> - ISO_8601 형식  <br/> - UTC 기준 (optional)
	updatedAt := "updatedAt_example" // string | 리소스가 마지막으로 수정된 시간 <br/> - ISO_8601 형식  <br/> - UTC 기준 (optional)
	sortKeys := "sortKeys_example" // string | 정렬할 필드를 콤마(,)로 구분     (optional) (default to "created_at")
	sortDirs := "sortDirs_example" // string | 정렬 방향 (`asc`, `desc`)   (optional) (default to "desc")
	limit := int32(56) // int32 | 페이지당 최대 반환 항목 수 (optional) (default to 20)
	offset := int32(56) // int32 | 조회 시작 위치 (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPCAPI.ListVpcs(context.Background()).XAuthToken(xAuthToken).Id(id).Name(name).CidrBlock(cidrBlock).ProvisioningStatus(provisioningStatus).IsDefault(isDefault).CreatedAt(createdAt).UpdatedAt(updatedAt).SortKeys(sortKeys).SortDirs(sortDirs).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCAPI.ListVpcs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListVpcs`: VPCListModel
	fmt.Fprintf(os.Stdout, "Response from `VPCAPI.ListVpcs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListVpcsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **id** | **string** | 조회할 VPC ID &lt;br/&gt;- [List VPCs](https://docs.kakaocloud.com/openapi/bns/vpc/list-vpcs) 참고 | 
 **name** | **string** | 조회할 VPC 이름 | 
 **cidrBlock** | **string** | VPC의 IPv4 CIDR 블록 (예: &#x60;10.0.0.0/16&#x60;)  | 
 **provisioningStatus** | [**ProvisioningStatus**](ProvisioningStatus.md) | 프로비저닝 상태 &lt;br/&gt; - &#x60;ACTIVE&#x60;: 활성 상태 &lt;br/&gt; - &#x60;DELETED&#x60;: 삭제됨 &lt;br/&gt; - &#x60;ERROR&#x60;: 오류 발생 &lt;br/&gt; - &#x60;PENDING_CREATE&#x60;: 생성 대기 중 &lt;br/&gt; - &#x60;PENDING_UPDATE&#x60;: 업데이트 대기 중 &lt;br/&gt; - &#x60;PENDING_DELETE&#x60;: 삭제 대기 중 | 
 **isDefault** | **bool** | 기본 VPC 여부 | 
 **createdAt** | **string** | 리소스가 생성된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
 **updatedAt** | **string** | 리소스가 마지막으로 수정된 시간 &lt;br/&gt; - ISO_8601 형식  &lt;br/&gt; - UTC 기준 | 
 **sortKeys** | **string** | 정렬할 필드를 콤마(,)로 구분     | [default to &quot;created_at&quot;]
 **sortDirs** | **string** | 정렬 방향 (&#x60;asc&#x60;, &#x60;desc&#x60;)   | [default to &quot;desc&quot;]
 **limit** | **int32** | 페이지당 최대 반환 항목 수 | [default to 20]
 **offset** | **int32** | 조회 시작 위치 | [default to 0]

### Return type

[**VPCListModel**](VPCListModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutBnsVpc

> BnsVpcV1ApiUpdateVpcModelResponseVPCModel PutBnsVpc(ctx, vpcId).XAuthToken(xAuthToken).BodyPutBnsVpc(bodyPutBnsVpc).Execute()

Update VPC



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
	vpcId := "vpcId_example" // string | 수정할 VPC ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bodyPutBnsVpc := *openapiclient.NewBodyPutBnsVpc(*openapiclient.NewEditVPCModel("Name_example")) // BodyPutBnsVpc | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPCAPI.PutBnsVpc(context.Background(), vpcId).XAuthToken(xAuthToken).BodyPutBnsVpc(bodyPutBnsVpc).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCAPI.PutBnsVpc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutBnsVpc`: BnsVpcV1ApiUpdateVpcModelResponseVPCModel
	fmt.Fprintf(os.Stdout, "Response from `VPCAPI.PutBnsVpc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vpcId** | **string** | 수정할 VPC ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutBnsVpcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyPutBnsVpc** | [**BodyPutBnsVpc**](BodyPutBnsVpc.md) |  | 

### Return type

[**BnsVpcV1ApiUpdateVpcModelResponseVPCModel**](BnsVpcV1ApiUpdateVpcModelResponseVPCModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

