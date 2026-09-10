# VolumeAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 볼륨의 고유 ID | 
**InstanceId** | **string** | 인스턴스의 고유 ID | 
**MountPoint** | **string** | 볼륨이 인스턴스에 마운트된 경로 | 
**IsDeleteOnTermination** | **bool** | 인스턴스 삭제 시 해당 볼륨을 자동으로 삭제할지 여부 | 

## Methods

### NewVolumeAttachment

`func NewVolumeAttachment(id string, instanceId string, mountPoint string, isDeleteOnTermination bool, ) *VolumeAttachment`

NewVolumeAttachment instantiates a new VolumeAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeAttachmentWithDefaults

`func NewVolumeAttachmentWithDefaults() *VolumeAttachment`

NewVolumeAttachmentWithDefaults instantiates a new VolumeAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VolumeAttachment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VolumeAttachment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VolumeAttachment) SetId(v string)`

SetId sets Id field to given value.


### GetInstanceId

`func (o *VolumeAttachment) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *VolumeAttachment) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *VolumeAttachment) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.


### GetMountPoint

`func (o *VolumeAttachment) GetMountPoint() string`

GetMountPoint returns the MountPoint field if non-nil, zero value otherwise.

### GetMountPointOk

`func (o *VolumeAttachment) GetMountPointOk() (*string, bool)`

GetMountPointOk returns a tuple with the MountPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMountPoint

`func (o *VolumeAttachment) SetMountPoint(v string)`

SetMountPoint sets MountPoint field to given value.


### GetIsDeleteOnTermination

`func (o *VolumeAttachment) GetIsDeleteOnTermination() bool`

GetIsDeleteOnTermination returns the IsDeleteOnTermination field if non-nil, zero value otherwise.

### GetIsDeleteOnTerminationOk

`func (o *VolumeAttachment) GetIsDeleteOnTerminationOk() (*bool, bool)`

GetIsDeleteOnTerminationOk returns a tuple with the IsDeleteOnTermination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleteOnTermination

`func (o *VolumeAttachment) SetIsDeleteOnTermination(v bool)`

SetIsDeleteOnTermination sets IsDeleteOnTermination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


