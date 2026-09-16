# UpdateTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | 인스턴스의 이름 | [optional] 
**Weight** | Pointer to **NullableInt32** | 트래픽 분산 가중치 | [optional] 
**MonitorPort** | Pointer to **NullableInt32** | 헬스 체크에 사용할 포트 번호 | [optional] 

## Methods

### NewUpdateTarget

`func NewUpdateTarget() *UpdateTarget`

NewUpdateTarget instantiates a new UpdateTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTargetWithDefaults

`func NewUpdateTargetWithDefaults() *UpdateTarget`

NewUpdateTargetWithDefaults instantiates a new UpdateTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateTarget) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateTarget) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateTarget) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateTarget) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetWeight

`func (o *UpdateTarget) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *UpdateTarget) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *UpdateTarget) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *UpdateTarget) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *UpdateTarget) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *UpdateTarget) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetMonitorPort

`func (o *UpdateTarget) GetMonitorPort() int32`

GetMonitorPort returns the MonitorPort field if non-nil, zero value otherwise.

### GetMonitorPortOk

`func (o *UpdateTarget) GetMonitorPortOk() (*int32, bool)`

GetMonitorPortOk returns a tuple with the MonitorPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorPort

`func (o *UpdateTarget) SetMonitorPort(v int32)`

SetMonitorPort sets MonitorPort field to given value.

### HasMonitorPort

`func (o *UpdateTarget) HasMonitorPort() bool`

HasMonitorPort returns a boolean if a field has been set.

### SetMonitorPortNil

`func (o *UpdateTarget) SetMonitorPortNil(b bool)`

 SetMonitorPortNil sets the value for MonitorPort to be an explicit nil

### UnsetMonitorPort
`func (o *UpdateTarget) UnsetMonitorPort()`

UnsetMonitorPort ensures that no value is present for MonitorPort, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


