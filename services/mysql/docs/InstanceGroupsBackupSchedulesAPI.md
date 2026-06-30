# \InstanceGroupsBackupSchedulesAPI

All URIs are relative to *https://mysql.kr-central-2.kakaocloud.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateMysqlInstanceGroupBackupSchedule**](InstanceGroupsBackupSchedulesAPI.md#UpdateMysqlInstanceGroupBackupSchedule) | **Post** /api/v1/instance-groups/{instance_group_id}/backup-schedules/{backup_schedule_id} | UpdateMySQLInstanceGroupBackupSchedule



## UpdateMysqlInstanceGroupBackupSchedule

> UpdateMysqlInstanceGroupBackupSchedule(ctx, instanceGroupId, backupScheduleId).XAuthToken(xAuthToken).BodyUpdateMysqlInstanceGroupBackupSchedule(bodyUpdateMysqlInstanceGroupBackupSchedule).Execute()

UpdateMySQLInstanceGroupBackupSchedule



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
	instanceGroupId := "instanceGroupId_example" // string | # instance_group_id  - 설명: (여기에 instance_group_id 에 대한 설명을 작성하세요)
	backupScheduleId := "backupScheduleId_example" // string | # backup_schedule_id  - 설명: (여기에 backup_schedule_id 에 대한 설명을 작성하세요)
	xAuthToken := "xAuthToken_example" // string | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급)
	bodyUpdateMysqlInstanceGroupBackupSchedule := *openapiclient.NewBodyUpdateMysqlInstanceGroupBackupSchedule(*openapiclient.NewMysqlV1ApiUpdateMysqlInstanceGroupBackupScheduleModelBackupScheduleRequestModel(false)) // BodyUpdateMysqlInstanceGroupBackupSchedule | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.InstanceGroupsBackupSchedulesAPI.UpdateMysqlInstanceGroupBackupSchedule(context.Background(), instanceGroupId, backupScheduleId).XAuthToken(xAuthToken).BodyUpdateMysqlInstanceGroupBackupSchedule(bodyUpdateMysqlInstanceGroupBackupSchedule).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InstanceGroupsBackupSchedulesAPI.UpdateMysqlInstanceGroupBackupSchedule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceGroupId** | **string** | # instance_group_id  - 설명: (여기에 instance_group_id 에 대한 설명을 작성하세요) | 
**backupScheduleId** | **string** | # backup_schedule_id  - 설명: (여기에 backup_schedule_id 에 대한 설명을 작성하세요) | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMysqlInstanceGroupBackupScheduleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAuthToken** | **string** | - [API 인증 토큰](https://docs.kakaocloud.com/openapi/start#api-인증-토큰-발급) | 
 **bodyUpdateMysqlInstanceGroupBackupSchedule** | [**BodyUpdateMysqlInstanceGroupBackupSchedule**](BodyUpdateMysqlInstanceGroupBackupSchedule.md) |  | 

### Return type

 (empty response body)

### Authorization

[x-auth-token](../README.md#x-auth-token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

