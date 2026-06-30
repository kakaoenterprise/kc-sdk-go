# GetMySQLBackupResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backup** | [**MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel**](MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel.md) | 조회된 MySQL 백업의 상세 정보 | 

## Methods

### NewGetMySQLBackupResponseModel

`func NewGetMySQLBackupResponseModel(backup MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel, ) *GetMySQLBackupResponseModel`

NewGetMySQLBackupResponseModel instantiates a new GetMySQLBackupResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMySQLBackupResponseModelWithDefaults

`func NewGetMySQLBackupResponseModelWithDefaults() *GetMySQLBackupResponseModel`

NewGetMySQLBackupResponseModelWithDefaults instantiates a new GetMySQLBackupResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackup

`func (o *GetMySQLBackupResponseModel) GetBackup() MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel`

GetBackup returns the Backup field if non-nil, zero value otherwise.

### GetBackupOk

`func (o *GetMySQLBackupResponseModel) GetBackupOk() (*MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel, bool)`

GetBackupOk returns a tuple with the Backup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackup

`func (o *GetMySQLBackupResponseModel) SetBackup(v MysqlV1ApiGetMysqlBackupModelMySQLBackupResponseModel)`

SetBackup sets Backup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


