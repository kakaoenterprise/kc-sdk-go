# GetMysqlBackupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backup** | [**MySQLBackup**](MySQLBackup.md) | 조회된 MySQL 백업의 상세 정보 | 

## Methods

### NewGetMysqlBackupResponse

`func NewGetMysqlBackupResponse(backup MySQLBackup, ) *GetMysqlBackupResponse`

NewGetMysqlBackupResponse instantiates a new GetMysqlBackupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMysqlBackupResponseWithDefaults

`func NewGetMysqlBackupResponseWithDefaults() *GetMysqlBackupResponse`

NewGetMysqlBackupResponseWithDefaults instantiates a new GetMysqlBackupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackup

`func (o *GetMysqlBackupResponse) GetBackup() MySQLBackup`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *GetMysqlBackupResponse) GetBackupOk() (*MySQLBackup, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *GetMysqlBackupResponse) SetBackup(v MySQLBackup)`

SetBackup sets Backup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


