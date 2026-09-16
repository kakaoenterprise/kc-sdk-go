# UpdateMysqlInstanceGroupBackupScheduleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupSchedule** | [**UpdateMysqlInstanceGroupBackupSchedule**](UpdateMysqlInstanceGroupBackupSchedule.md) | MySQL 백업 스케줄 설정 정보 | 

## Methods

### NewUpdateMysqlInstanceGroupBackupScheduleRequest

`func NewUpdateMysqlInstanceGroupBackupScheduleRequest(backupSchedule UpdateMysqlInstanceGroupBackupSchedule, ) *UpdateMysqlInstanceGroupBackupScheduleRequest`

NewUpdateMysqlInstanceGroupBackupScheduleRequest instantiates a new UpdateMysqlInstanceGroupBackupScheduleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMysqlInstanceGroupBackupScheduleRequestWithDefaults

`func NewUpdateMysqlInstanceGroupBackupScheduleRequestWithDefaults() *UpdateMysqlInstanceGroupBackupScheduleRequest`

NewUpdateMysqlInstanceGroupBackupScheduleRequestWithDefaults instantiates a new UpdateMysqlInstanceGroupBackupScheduleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupSchedule

`func (o *UpdateMysqlInstanceGroupBackupScheduleRequest) GetBackupSchedule() UpdateMysqlInstanceGroupBackupSchedule`

GetBackupSchedule returns the BackupSchedule field if non-nil, zero value otherwise.

### GetBackupScheduleOk

`func (o *UpdateMysqlInstanceGroupBackupScheduleRequest) GetBackupScheduleOk() (*UpdateMysqlInstanceGroupBackupSchedule, bool)`

GetBackupScheduleOk returns a tuple with the BackupSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupSchedule

`func (o *UpdateMysqlInstanceGroupBackupScheduleRequest) SetBackupSchedule(v UpdateMysqlInstanceGroupBackupSchedule)`

SetBackupSchedule sets BackupSchedule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


