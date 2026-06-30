# BackupRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 생성할 MySQL 백업 이름 | 
**InstanceGroupId** | **string** | 백업을 생성할 대상 MySQL 인스턴스 그룹 ID | 

## Methods

### NewBackupRequestModel

`func NewBackupRequestModel(name string, instanceGroupId string, ) *BackupRequestModel`

NewBackupRequestModel instantiates a new BackupRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackupRequestModelWithDefaults

`func NewBackupRequestModelWithDefaults() *BackupRequestModel`

NewBackupRequestModelWithDefaults instantiates a new BackupRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BackupRequestModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BackupRequestModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BackupRequestModel) SetName(v string)`

SetName sets Name field to given value.


### GetInstanceGroupId

`func (o *BackupRequestModel) GetInstanceGroupId() string`

GetInstanceGroupId returns the InstanceGroupId field if non-nil, zero value otherwise.

### GetInstanceGroupIdOk

`func (o *BackupRequestModel) GetInstanceGroupIdOk() (*string, bool)`

GetInstanceGroupIdOk returns a tuple with the InstanceGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceGroupId

`func (o *BackupRequestModel) SetInstanceGroupId(v string)`

SetInstanceGroupId sets InstanceGroupId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


