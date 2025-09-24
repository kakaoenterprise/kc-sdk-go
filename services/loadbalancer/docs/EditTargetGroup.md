# EditTargetGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**LoadBalancerAlgorithm** | Pointer to [**NullableTargetGroupAlgorithm**](TargetGroupAlgorithm.md) |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**SessionPersistence** | Pointer to [**NullableSessionPersistenceModel**](SessionPersistenceModel.md) |  | [optional] 

## Methods

### NewEditTargetGroup

`func NewEditTargetGroup() *EditTargetGroup`

NewEditTargetGroup instantiates a new EditTargetGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditTargetGroupWithDefaults

`func NewEditTargetGroupWithDefaults() *EditTargetGroup`

NewEditTargetGroupWithDefaults instantiates a new EditTargetGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *EditTargetGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EditTargetGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EditTargetGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EditTargetGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *EditTargetGroup) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *EditTargetGroup) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLoadBalancerAlgorithm

`func (o *EditTargetGroup) GetLoadBalancerAlgorithm() TargetGroupAlgorithm`

GetLoadBalancerAlgorithm returns the LoadBalancerAlgorithm field if non-nil, zero value otherwise.

### GetLoadBalancerAlgorithmOk

`func (o *EditTargetGroup) GetLoadBalancerAlgorithmOk() (*TargetGroupAlgorithm, bool)`

GetLoadBalancerAlgorithmOk returns a tuple with the LoadBalancerAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerAlgorithm

`func (o *EditTargetGroup) SetLoadBalancerAlgorithm(v TargetGroupAlgorithm)`

SetLoadBalancerAlgorithm sets LoadBalancerAlgorithm field to given value.

### HasLoadBalancerAlgorithm

`func (o *EditTargetGroup) HasLoadBalancerAlgorithm() bool`

HasLoadBalancerAlgorithm returns a boolean if a field has been set.

### SetLoadBalancerAlgorithmNil

`func (o *EditTargetGroup) SetLoadBalancerAlgorithmNil(b bool)`

 SetLoadBalancerAlgorithmNil sets the value for LoadBalancerAlgorithm to be an explicit nil

### UnsetLoadBalancerAlgorithm
`func (o *EditTargetGroup) UnsetLoadBalancerAlgorithm()`

UnsetLoadBalancerAlgorithm ensures that no value is present for LoadBalancerAlgorithm, not even an explicit nil
### GetName

`func (o *EditTargetGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditTargetGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditTargetGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EditTargetGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *EditTargetGroup) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *EditTargetGroup) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSessionPersistence

`func (o *EditTargetGroup) GetSessionPersistence() SessionPersistenceModel`

GetSessionPersistence returns the SessionPersistence field if non-nil, zero value otherwise.

### GetSessionPersistenceOk

`func (o *EditTargetGroup) GetSessionPersistenceOk() (*SessionPersistenceModel, bool)`

GetSessionPersistenceOk returns a tuple with the SessionPersistence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionPersistence

`func (o *EditTargetGroup) SetSessionPersistence(v SessionPersistenceModel)`

SetSessionPersistence sets SessionPersistence field to given value.

### HasSessionPersistence

`func (o *EditTargetGroup) HasSessionPersistence() bool`

HasSessionPersistence returns a boolean if a field has been set.

### SetSessionPersistenceNil

`func (o *EditTargetGroup) SetSessionPersistenceNil(b bool)`

 SetSessionPersistenceNil sets the value for SessionPersistence to be an explicit nil

### UnsetSessionPersistence
`func (o *EditTargetGroup) UnsetSessionPersistence()`

UnsetSessionPersistence ensures that no value is present for SessionPersistence, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


