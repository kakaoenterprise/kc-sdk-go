# UpdateTargetGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | 대상 그룹에 대한 설명 | [optional] 
**LoadBalancerAlgorithm** | Pointer to [**NullableLoadBalancerPoolAlgorithm**](LoadBalancerPoolAlgorithm.md) | 로드 밸런싱 알고리즘 | [optional] 
**Name** | Pointer to **NullableString** | 대상 그룹 이름 | [optional] 
**SessionPersistence** | Pointer to [**NullableSessionPersistenceRequest**](SessionPersistenceRequest.md) | 세션 설정 | [optional] 

## Methods

### NewUpdateTargetGroup

`func NewUpdateTargetGroup() *UpdateTargetGroup`

NewUpdateTargetGroup instantiates a new UpdateTargetGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTargetGroupWithDefaults

`func NewUpdateTargetGroupWithDefaults() *UpdateTargetGroup`

NewUpdateTargetGroupWithDefaults instantiates a new UpdateTargetGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *UpdateTargetGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateTargetGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateTargetGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateTargetGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateTargetGroup) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateTargetGroup) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLoadBalancerAlgorithm

`func (o *UpdateTargetGroup) GetLoadBalancerAlgorithm() LoadBalancerPoolAlgorithm`

GetLoadBalancerAlgorithm returns the LoadBalancerAlgorithm field if non-nil, zero value otherwise.

### GetLoadBalancerAlgorithmOk

`func (o *UpdateTargetGroup) GetLoadBalancerAlgorithmOk() (*LoadBalancerPoolAlgorithm, bool)`

GetLoadBalancerAlgorithmOk returns a tuple with the LoadBalancerAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerAlgorithm

`func (o *UpdateTargetGroup) SetLoadBalancerAlgorithm(v LoadBalancerPoolAlgorithm)`

SetLoadBalancerAlgorithm sets LoadBalancerAlgorithm field to given value.

### HasLoadBalancerAlgorithm

`func (o *UpdateTargetGroup) HasLoadBalancerAlgorithm() bool`

HasLoadBalancerAlgorithm returns a boolean if a field has been set.

### SetLoadBalancerAlgorithmNil

`func (o *UpdateTargetGroup) SetLoadBalancerAlgorithmNil(b bool)`

 SetLoadBalancerAlgorithmNil sets the value for LoadBalancerAlgorithm to be an explicit nil

### UnsetLoadBalancerAlgorithm
`func (o *UpdateTargetGroup) UnsetLoadBalancerAlgorithm()`

UnsetLoadBalancerAlgorithm ensures that no value is present for LoadBalancerAlgorithm, not even an explicit nil
### GetName

`func (o *UpdateTargetGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateTargetGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateTargetGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateTargetGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateTargetGroup) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateTargetGroup) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSessionPersistence

`func (o *UpdateTargetGroup) GetSessionPersistence() SessionPersistenceRequest`

GetSessionPersistence returns the SessionPersistence field if non-nil, zero value otherwise.

### GetSessionPersistenceOk

`func (o *UpdateTargetGroup) GetSessionPersistenceOk() (*SessionPersistenceRequest, bool)`

GetSessionPersistenceOk returns a tuple with the SessionPersistence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionPersistence

`func (o *UpdateTargetGroup) SetSessionPersistence(v SessionPersistenceRequest)`

SetSessionPersistence sets SessionPersistence field to given value.

### HasSessionPersistence

`func (o *UpdateTargetGroup) HasSessionPersistence() bool`

HasSessionPersistence returns a boolean if a field has been set.

### SetSessionPersistenceNil

`func (o *UpdateTargetGroup) SetSessionPersistenceNil(b bool)`

 SetSessionPersistenceNil sets the value for SessionPersistence to be an explicit nil

### UnsetSessionPersistence
`func (o *UpdateTargetGroup) UnsetSessionPersistence()`

UnsetSessionPersistence ensures that no value is present for SessionPersistence, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


