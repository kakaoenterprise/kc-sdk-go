# CreateMySQLBackupResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backup** | [**BackupResponseModel**](BackupResponseModel.md) | 생성된 MySQL 백업의 식별자 정보를 포함하는 객체 | 

## Methods

### NewCreateMySQLBackupResponseModel

`func NewCreateMySQLBackupResponseModel(backup BackupResponseModel, ) *CreateMySQLBackupResponseModel`

NewCreateMySQLBackupResponseModel instantiates a new CreateMySQLBackupResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMySQLBackupResponseModelWithDefaults

`func NewCreateMySQLBackupResponseModelWithDefaults() *CreateMySQLBackupResponseModel`

NewCreateMySQLBackupResponseModelWithDefaults instantiates a new CreateMySQLBackupResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackup

`func (o *CreateMySQLBackupResponseModel) GetBackup() BackupResponseModel`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *CreateMySQLBackupResponseModel) GetBackupOk() (*BackupResponseModel, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *CreateMySQLBackupResponseModel) SetBackup(v BackupResponseModel)`

SetBackup sets Backup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


