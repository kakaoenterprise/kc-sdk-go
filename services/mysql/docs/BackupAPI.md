# \BackupAPI

All URIs are relative to *https://mysql.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMysqlBackup**](BackupAPI.md#CreateMysqlBackup) | **Post** /api/v1/backups | Create MySQL backup
[**DeleteMysqlBackup**](BackupAPI.md#DeleteMysqlBackup) | **Delete** /api/v1/backups/{backup_id} | Delete MySQL backup
[**GetMysqlBackup**](BackupAPI.md#GetMysqlBackup) | **Get** /api/v1/backups/{backup_id} | Get MySQL backup
[**ListMysqlBackups**](BackupAPI.md#ListMysqlBackups) | **Get** /api/v1/backups | List MySQL backups



## CreateMysqlBackup

> CreateMysqlBackupResponse CreateMysqlBackup(ctx).XAuthToken(xAuthToken).CreateMysqlBackupRequest(createMysqlBackupRequest).Execute()

Create MySQL backup



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/mysql"
)

func main() {
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	createMysqlBackupRequest := *openapiclient.NewCreateMysqlBackupRequest(*openapiclient.NewCreateMysqlBackup("Name_example", "InstanceGroupId_example")) // CreateMysqlBackupRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupAPI.CreateMysqlBackup(context.Background()).XAuthToken(xAuthToken).CreateMysqlBackupRequest(createMysqlBackupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAPI.CreateMysqlBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMysqlBackup`: CreateMysqlBackupResponse
	fmt.Fprintf(os.Stdout, "Response from `BackupAPI.CreateMysqlBackup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMysqlBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **createMysqlBackupRequest** | [**CreateMysqlBackupRequest**](CreateMysqlBackupRequest.md) |  | 

### Return type

[**CreateMysqlBackupResponse**](CreateMysqlBackupResponse.md)

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
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/mysql"
)

func main() {
	backupId := "backupId_example" // string | 대상 MySQL 백업 ID <br/>- [List MySQL backups](https://docs.kakaocloud.com/openapi/data-store/mysql/list-mysql-backups)에서 확인
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BackupAPI.DeleteMysqlBackup(context.Background(), backupId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAPI.DeleteMysqlBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**backupId** | **string** | 대상 MySQL 백업 ID &lt;br/&gt;- [List MySQL backups](https://docs.kakaocloud.com/openapi/data-store/mysql/list-mysql-backups)에서 확인 | 

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

> GetMysqlBackupResponse GetMysqlBackup(ctx, backupId).XAuthToken(xAuthToken).Execute()

Get MySQL backup



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/mysql"
)

func main() {
	backupId := "backupId_example" // string | 대상 MySQL 백업 ID
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupAPI.GetMysqlBackup(context.Background(), backupId).XAuthToken(xAuthToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAPI.GetMysqlBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMysqlBackup`: GetMysqlBackupResponse
	fmt.Fprintf(os.Stdout, "Response from `BackupAPI.GetMysqlBackup`: %v\n", resp)
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

[**GetMysqlBackupResponse**](GetMysqlBackupResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMysqlBackups

> ListMysqlBackupsResponse ListMysqlBackups(ctx).XAuthToken(xAuthToken).InstanceGroupId(instanceGroupId).Execute()

List MySQL backups



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kakaoenterprise/kc-sdk-go/v2/services/mysql"
)

func main() {
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	instanceGroupId := "instanceGroupId_example" // string | 조회할 MySQL 인스턴스 그룹 ID (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BackupAPI.ListMysqlBackups(context.Background()).XAuthToken(xAuthToken).InstanceGroupId(instanceGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BackupAPI.ListMysqlBackups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMysqlBackups`: ListMysqlBackupsResponse
	fmt.Fprintf(os.Stdout, "Response from `BackupAPI.ListMysqlBackups`: %v\n", resp)
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

[**ListMysqlBackupsResponse**](ListMysqlBackupsResponse.md)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

