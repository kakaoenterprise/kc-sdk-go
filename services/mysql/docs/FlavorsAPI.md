# \FlavorsAPI

All URIs are relative to *https://mysql.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListMysqlInstanceTypesFlavors**](FlavorsAPI.md#ListMysqlInstanceTypesFlavors) | **Get** /api/v1/flavors | List MySQL instance types (flavors)



## ListMysqlInstanceTypesFlavors

> GetMySQLFlavorsResponseModel ListMysqlInstanceTypesFlavors(ctx).XAuthToken(xAuthToken).ShowAll(showAll).Execute()

List MySQL instance types (flavors)



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/services/mysql"
)

func main() {
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	showAll := true // bool | 모든 인스턴스 유형(Flavor)을 포함할지 여부 <br/> - `false`로 설정 시, 사용 중단된 인스턴스 유형을 모두 제외하여 반환 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FlavorsAPI.ListMysqlInstanceTypesFlavors(context.Background()).XAuthToken(xAuthToken).ShowAll(showAll).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FlavorsAPI.ListMysqlInstanceTypesFlavors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMysqlInstanceTypesFlavors`: GetMySQLFlavorsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `FlavorsAPI.ListMysqlInstanceTypesFlavors`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMysqlInstanceTypesFlavorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **showAll** | **bool** | 모든 인스턴스 유형(Flavor)을 포함할지 여부 &lt;br/&gt; - &#x60;false&#x60;로 설정 시, 사용 중단된 인스턴스 유형을 모두 제외하여 반환 | 

### Return type

[**GetMySQLFlavorsResponseModel**](GetMySQLFlavorsResponseModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

