# RequestRestoreVolumeSnapshotModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 볼륨에 대한 설명 | 
**AvailabilityZone** | [**AvailabilityZone**](AvailabilityZone.md) | 신규 볼륨을 생성할 가용 영역&lt;br/&gt; - &#x60;kr-central-2-a&#x60;: kr-central-2-a 가용 영역 &lt;br/&gt; - &#x60;kr-central-2-b&#x60;: kr-central-2-b 가용 영역 &lt;br/&gt; - &#x60;kr-central-2-c&#x60;: kr-central-2-c 가용 영역 | 
**VolumeTypeId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRequestRestoreVolumeSnapshotModel

`func NewRequestRestoreVolumeSnapshotModel(name string, availabilityZone AvailabilityZone, ) *RequestRestoreVolumeSnapshotModel`

NewRequestRestoreVolumeSnapshotModel instantiates a new RequestRestoreVolumeSnapshotModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestRestoreVolumeSnapshotModelWithDefaults

`func NewRequestRestoreVolumeSnapshotModelWithDefaults() *RequestRestoreVolumeSnapshotModel`

NewRequestRestoreVolumeSnapshotModelWithDefaults instantiates a new RequestRestoreVolumeSnapshotModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RequestRestoreVolumeSnapshotModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RequestRestoreVolumeSnapshotModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RequestRestoreVolumeSnapshotModel) SetName(v string)`

SetName sets Name field to given value.


### GetAvailabilityZone

`func (o *RequestRestoreVolumeSnapshotModel) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *RequestRestoreVolumeSnapshotModel) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *RequestRestoreVolumeSnapshotModel) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.


### GetVolumeTypeId

`func (o *RequestRestoreVolumeSnapshotModel) GetVolumeTypeId() string`

GetVolumeTypeId returns the VolumeTypeId field if non-nil, zero value otherwise.

### GetVolumeTypeIdOk

`func (o *RequestRestoreVolumeSnapshotModel) GetVolumeTypeIdOk() (*string, bool)`

GetVolumeTypeIdOk returns a tuple with the VolumeTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeTypeId

`func (o *RequestRestoreVolumeSnapshotModel) SetVolumeTypeId(v string)`

SetVolumeTypeId sets VolumeTypeId field to given value.

### HasVolumeTypeId

`func (o *RequestRestoreVolumeSnapshotModel) HasVolumeTypeId() bool`

HasVolumeTypeId returns a boolean if a field has been set.

### SetVolumeTypeIdNil

`func (o *RequestRestoreVolumeSnapshotModel) SetVolumeTypeIdNil(b bool)`

 SetVolumeTypeIdNil sets the value for VolumeTypeId to be an explicit nil

### UnsetVolumeTypeId
`func (o *RequestRestoreVolumeSnapshotModel) UnsetVolumeTypeId()`

UnsetVolumeTypeId ensures that no value is present for VolumeTypeId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


