# \BackupsAPI

All URIs are relative to *https://mysql.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMysqlBackup**](BackupsAPI.md#CreateMysqlBackup) | **Post** /api/v1/backups | Create MySQL backup
[**DeleteMysqlBackup**](BackupsAPI.md#DeleteMysqlBackup) | **Delete** /api/v1/backups/{backup_id} | Delete MySQL backup
[**GetMysqlBackup**](BackupsAPI.md#GetMysqlBackup) | **Get** /api/v1/backups/{backup_id} | Get MySQL backup
[**ListMysqlBackups**](BackupsAPI.md#ListMysqlBackups) | **Get** /api/v1/backups | List MySQL backups



## CreateMysqlBackup

> CreateMySQLBackupResponseModel CreateMysqlBackup(ctx).XAuthToken(xAuthToken).BodyCreateMysqlBackup(bodyCreateMysqlBackup).Execute()

Create MySQL backup



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
	bodyCreateMysqlBackup := *openapiclient.NewBodyCreateMysqlBackup(*openapiclient.NewBackupRequestModel("Name_example", "InstanceGroupId_example")) // BodyCreateMysqlBackup | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupsAPI.CreateMysqlBackup(context.Background()).XAuthToken(xAuthToken).BodyCreateMysqlBackup(bodyCreateMysqlBackup).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupsAPI.CreateMysqlBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMysqlBackup`: CreateMySQLBackupResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BackupsAPI.CreateMysqlBackup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMysqlBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyCreateMysqlBackup** | [**BodyCreateMysqlBackup**](BodyCreateMysqlBackup.md) |  | 

### Return type

[**CreateMySQLBackupResponseModel**](CreateMySQLBackupResponseModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMysqlBackup

> DeleteMysqlBackup(ctx, backupId).XAuthToken(xAuthToken).Execute()

Delete MySQL backup



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
	backupId := "backupId_example" // string | 대상 MySQL 백업 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackupsAPI.DeleteMysqlBackup(context.Background(), backupId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupsAPI.DeleteMysqlBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupId** | **string** | 대상 MySQL 백업 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMysqlBackupRequest struct via the builder pattern


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


## GetMysqlBackup

> GetMySQLBackupResponseModel GetMysqlBackup(ctx, backupId).XAuthToken(xAuthToken).Execute()

Get MySQL backup



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
	backupId := "backupId_example" // string | 대상 MySQL 백업 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupsAPI.GetMysqlBackup(context.Background(), backupId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupsAPI.GetMysqlBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMysqlBackup`: GetMySQLBackupResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BackupsAPI.GetMysqlBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupId** | **string** | 대상 MySQL 백업 ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMysqlBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 

### Return type

[**GetMySQLBackupResponseModel**](GetMySQLBackupResponseModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMysqlBackups

> GetMySQLBackupsResponseModel ListMysqlBackups(ctx).XAuthToken(xAuthToken).InstanceGroupId(instanceGroupId).Execute()

List MySQL backups



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
	instanceGroupId := "instanceGroupId_example" // string | 조회할 MySQL 인스턴스 그룹 ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupsAPI.ListMysqlBackups(context.Background()).XAuthToken(xAuthToken).InstanceGroupId(instanceGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupsAPI.ListMysqlBackups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMysqlBackups`: GetMySQLBackupsResponseModel
	fmt.Fprintf(os.Stdout, "Response from `BackupsAPI.ListMysqlBackups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMysqlBackupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **instanceGroupId** | **string** | 조회할 MySQL 인스턴스 그룹 ID | 

### Return type

[**GetMySQLBackupsResponseModel**](GetMySQLBackupsResponseModel.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

