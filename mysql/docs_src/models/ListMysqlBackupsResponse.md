# ListMysqlBackupsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backups** | [**[]MySQLBackup**](MySQLBackup.md) | 조회된 MySQL 백업 목록 | 

## Methods

### NewListMysqlBackupsResponse

`func NewListMysqlBackupsResponse(backups []MySQLBackup, ) *ListMysqlBackupsResponse`

NewListMysqlBackupsResponse instantiates a new ListMysqlBackupsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMysqlBackupsResponseWithDefaults

`func NewListMysqlBackupsResponseWithDefaults() *ListMysqlBackupsResponse`

NewListMysqlBackupsResponseWithDefaults instantiates a new ListMysqlBackupsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackups

`func (o *ListMysqlBackupsResponse) GetBackups() []MySQLBackup`

GetBackups returns the Backups field if non-nil, zero value otherwise.

### GetBackupsOk

`func (o *ListMysqlBackupsResponse) GetBackupsOk() (*[]MySQLBackup, bool)`

GetBackupsOk returns a tuple with the Backups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackups

`func (o *ListMysqlBackupsResponse) SetBackups(v []MySQLBackup)`

SetBackups sets Backups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


