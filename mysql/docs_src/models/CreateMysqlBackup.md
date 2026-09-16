# CreateMysqlBackup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 생성할 MySQL 백업 이름 | 
**InstanceGroupId** | **string** | 백업을 생성할 대상 MySQL 인스턴스 그룹 ID - [List MySQL instance groups](/openapi/data-store/mysql/list-mysql-instance-groups)에서 확인 | 

## Methods

### NewCreateMysqlBackup

`func NewCreateMysqlBackup(name string, instanceGroupId string, ) *CreateMysqlBackup`

NewCreateMysqlBackup instantiates a new CreateMysqlBackup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMysqlBackupWithDefaults

`func NewCreateMysqlBackupWithDefaults() *CreateMysqlBackup`

NewCreateMysqlBackupWithDefaults instantiates a new CreateMysqlBackup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateMysqlBackup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateMysqlBackup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateMysqlBackup) SetName(v string)`

SetName sets Name field to given value.


### GetInstanceGroupId

`func (o *CreateMysqlBackup) GetInstanceGroupId() string`

GetInstanceGroupId returns the InstanceGroupId field if non-nil, zero value otherwise.

### GetInstanceGroupIdOk

`func (o *CreateMysqlBackup) GetInstanceGroupIdOk() (*string, bool)`

GetInstanceGroupIdOk returns a tuple with the InstanceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGroupId

`func (o *CreateMysqlBackup) SetInstanceGroupId(v string)`

SetInstanceGroupId sets InstanceGroupId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


