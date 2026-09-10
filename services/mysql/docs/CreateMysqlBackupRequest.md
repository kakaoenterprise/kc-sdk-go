# CreateMysqlBackupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backup** | [**CreateMysqlBackup**](CreateMysqlBackup.md) | 생성할 MySQL 백업의 이름 및 대상 정보를 정의하는 객체 | 

## Methods

### NewCreateMysqlBackupRequest

`func NewCreateMysqlBackupRequest(backup CreateMysqlBackup, ) *CreateMysqlBackupRequest`

NewCreateMysqlBackupRequest instantiates a new CreateMysqlBackupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMysqlBackupRequestWithDefaults

`func NewCreateMysqlBackupRequestWithDefaults() *CreateMysqlBackupRequest`

NewCreateMysqlBackupRequestWithDefaults instantiates a new CreateMysqlBackupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackup

`func (o *CreateMysqlBackupRequest) GetBackup() CreateMysqlBackup`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *CreateMysqlBackupRequest) GetBackupOk() (*CreateMysqlBackup, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *CreateMysqlBackupRequest) SetBackup(v CreateMysqlBackup)`

SetBackup sets Backup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


