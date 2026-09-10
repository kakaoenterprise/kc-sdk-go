# CreatedPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 대상 그룹 ID | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ProvisioningStatus** | Pointer to [**NullableProvisioningStatus**](ProvisioningStatus.md) |  | [optional] 
**OperatingStatus** | Pointer to [**NullableLoadBalancerOperatingStatus**](LoadBalancerOperatingStatus.md) |  | [optional] 
**Protocol** | [**LoadBalancerPoolProtocol**](LoadBalancerPoolProtocol.md) | 적용된 프로토콜 | 
**LoadBalancerAlgorithm** | [**LoadBalancerPoolAlgorithm**](LoadBalancerPoolAlgorithm.md) | 적용된 분산 알고리즘 | 
**SessionPersistence** | Pointer to [**NullableSessionPersistence**](SessionPersistence.md) |  | [optional] 
**Members** | [**[]CreatedMember**](CreatedMember.md) | 현재 등록된 서버 목록 | 
**AlpnProtocols** | Pointer to [**[]LoadBalancerAlpnProtocol**](LoadBalancerAlpnProtocol.md) |  | [optional] 

## Methods

### NewCreatedPool

`func NewCreatedPool(id string, protocol LoadBalancerPoolProtocol, loadBalancerAlgorithm LoadBalancerPoolAlgorithm, members []CreatedMember, ) *CreatedPool`

NewCreatedPool instantiates a new CreatedPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatedPoolWithDefaults

`func NewCreatedPoolWithDefaults() *CreatedPool`

NewCreatedPoolWithDefaults instantiates a new CreatedPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreatedPool) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreatedPool) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreatedPool) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CreatedPool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatedPool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatedPool) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreatedPool) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreatedPool) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreatedPool) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *CreatedPool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreatedPool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreatedPool) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreatedPool) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreatedPool) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreatedPool) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProvisioningStatus

`func (o *CreatedPool) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *CreatedPool) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *CreatedPool) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *CreatedPool) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *CreatedPool) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *CreatedPool) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetOperatingStatus

`func (o *CreatedPool) GetOperatingStatus() LoadBalancerOperatingStatus`

GetOperatingStatus returns the OperatingStatus field if non-nil, zero value otherwise.

### GetOperatingStatusOk

`func (o *CreatedPool) GetOperatingStatusOk() (*LoadBalancerOperatingStatus, bool)`

GetOperatingStatusOk returns a tuple with the OperatingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingStatus

`func (o *CreatedPool) SetOperatingStatus(v LoadBalancerOperatingStatus)`

SetOperatingStatus sets OperatingStatus field to given value.

### HasOperatingStatus

`func (o *CreatedPool) HasOperatingStatus() bool`

HasOperatingStatus returns a boolean if a field has been set.

### SetOperatingStatusNil

`func (o *CreatedPool) SetOperatingStatusNil(b bool)`

 SetOperatingStatusNil sets the value for OperatingStatus to be an explicit nil

### UnsetOperatingStatus
`func (o *CreatedPool) UnsetOperatingStatus()`

UnsetOperatingStatus ensures that no value is present for OperatingStatus, not even an explicit nil
### GetProtocol

`func (o *CreatedPool) GetProtocol() LoadBalancerPoolProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *CreatedPool) GetProtocolOk() (*LoadBalancerPoolProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *CreatedPool) SetProtocol(v LoadBalancerPoolProtocol)`

SetProtocol sets Protocol field to given value.


### GetLoadBalancerAlgorithm

`func (o *CreatedPool) GetLoadBalancerAlgorithm() LoadBalancerPoolAlgorithm`

GetLoadBalancerAlgorithm returns the LoadBalancerAlgorithm field if non-nil, zero value otherwise.

### GetLoadBalancerAlgorithmOk

`func (o *CreatedPool) GetLoadBalancerAlgorithmOk() (*LoadBalancerPoolAlgorithm, bool)`

GetLoadBalancerAlgorithmOk returns a tuple with the LoadBalancerAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerAlgorithm

`func (o *CreatedPool) SetLoadBalancerAlgorithm(v LoadBalancerPoolAlgorithm)`

SetLoadBalancerAlgorithm sets LoadBalancerAlgorithm field to given value.


### GetSessionPersistence

`func (o *CreatedPool) GetSessionPersistence() SessionPersistence`

GetSessionPersistence returns the SessionPersistence field if non-nil, zero value otherwise.

### GetSessionPersistenceOk

`func (o *CreatedPool) GetSessionPersistenceOk() (*SessionPersistence, bool)`

GetSessionPersistenceOk returns a tuple with the SessionPersistence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionPersistence

`func (o *CreatedPool) SetSessionPersistence(v SessionPersistence)`

SetSessionPersistence sets SessionPersistence field to given value.

### HasSessionPersistence

`func (o *CreatedPool) HasSessionPersistence() bool`

HasSessionPersistence returns a boolean if a field has been set.

### SetSessionPersistenceNil

`func (o *CreatedPool) SetSessionPersistenceNil(b bool)`

 SetSessionPersistenceNil sets the value for SessionPersistence to be an explicit nil

### UnsetSessionPersistence
`func (o *CreatedPool) UnsetSessionPersistence()`

UnsetSessionPersistence ensures that no value is present for SessionPersistence, not even an explicit nil
### GetMembers

`func (o *CreatedPool) GetMembers() []CreatedMember`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *CreatedPool) GetMembersOk() (*[]CreatedMember, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *CreatedPool) SetMembers(v []CreatedMember)`

SetMembers sets Members field to given value.


### GetAlpnProtocols

`func (o *CreatedPool) GetAlpnProtocols() []*LoadBalancerAlpnProtocol`

GetAlpnProtocols returns the AlpnProtocols field if non-nil, zero value otherwise.

### GetAlpnProtocolsOk

`func (o *CreatedPool) GetAlpnProtocolsOk() (*[]*LoadBalancerAlpnProtocol, bool)`

GetAlpnProtocolsOk returns a tuple with the AlpnProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlpnProtocols

`func (o *CreatedPool) SetAlpnProtocols(v []*LoadBalancerAlpnProtocol)`

SetAlpnProtocols sets AlpnProtocols field to given value.

### HasAlpnProtocols

`func (o *CreatedPool) HasAlpnProtocols() bool`

HasAlpnProtocols returns a boolean if a field has been set.

### SetAlpnProtocolsNil

`func (o *CreatedPool) SetAlpnProtocolsNil(b bool)`

 SetAlpnProtocolsNil sets the value for AlpnProtocols to be an explicit nil

### UnsetAlpnProtocols
`func (o *CreatedPool) UnsetAlpnProtocols()`

UnsetAlpnProtocols ensures that no value is present for AlpnProtocols, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


