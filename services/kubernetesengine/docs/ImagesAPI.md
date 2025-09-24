# \ImagesAPI

All URIs are relative to *https://kubernetes-engine.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListNodePoolImages**](ImagesAPI.md#ListNodePoolImages) | **Get** /api/v1/images | List node pool images  



## ListNodePoolImages

> GetK8sImagesResponseModel ListNodePoolImages(ctx).XAuthToken(xAuthToken).OsDistro(osDistro).InstanceType(instanceType).IsGpuType(isGpuType).K8sVersion(k8sVersion).Execute()

List node pool images  



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
	osDistro := "osDistro_example" // string | 운영체제 배포판 (optional)
	instanceType := openapiclient.ImageInstanceType("vm") // ImageInstanceType | 인스턴스 유형 <br/> - `vm`: Virtual Machine 유형  <br/> - `bm`: Bare Metal Server 유형 (optional)
	isGpuType := true // bool | GPU 지원 이미지 여부 <br/> - `true`: GPU 장비를 지원하는 이미지<br/> - `false`: CPU 전용 이미지 (optional)
	k8sVersion := "k8sVersion_example" // string | 대상 Kubernetes 버전 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ImagesAPI.ListNodePoolImages(context.Background()).XAuthToken(xAuthToken).OsDistro(osDistro).InstanceType(instanceType).IsGpuType(isGpuType).K8sVersion(k8sVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ImagesAPI.ListNodePoolImages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListNodePoolImages`: GetK8sImagesResponseModel
	fmt.Fprintf(os.Stdout, "Response from `ImagesAPI.ListNodePoolImages`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNodePoolImagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **osDistro** | **string** | 운영체제 배포판 | 
 **instanceType** | [**ImageInstanceType**](ImageInstanceType.md) | 인스턴스 유형 &lt;br/&gt; - &#x60;vm&#x60;: Virtual Machine 유형  &lt;br/&gt; - &#x60;bm&#x60;: Bare Metal Server 유형 | 
 **isGpuType** | **bool** | GPU 지원 이미지 여부 &lt;br/&gt; - &#x60;true&#x60;: GPU 장비를 지원하는 이미지&lt;br/&gt; - &#x60;false&#x60;: CPU 전용 이미지 | 
 **k8sVersion** | **string** | 대상 Kubernetes 버전 | 

### Return type

[**GetK8sImagesResponseModel**](GetK8sImagesResponseModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

