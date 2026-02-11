# BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 대상 그룹 ID | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ProvisioningStatus** | Pointer to [**NullableProvisioningStatus**](ProvisioningStatus.md) |  | [optional] 
**OperatingStatus** | Pointer to [**NullableLoadBalancerOperatingStatus**](LoadBalancerOperatingStatus.md) |  | [optional] 
**Protocol** | Pointer to [**NullableTargetGroupProtocol**](TargetGroupProtocol.md) |  | [optional] 
**LoadBalancerAlgorithm** | Pointer to [**NullableTargetGroupAlgorithm**](TargetGroupAlgorithm.md) |  | [optional] 
**SessionPersistence** | Pointer to [**NullableSessionPersistenceModel**](SessionPersistenceModel.md) |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**LoadBalancers** | Pointer to [**[]BnsLoadBalancerV1ApiCreateTargetGroupModelLoadBalancerModel**](BnsLoadBalancerV1ApiCreateTargetGroupModelLoadBalancerModel.md) |  | [optional] 
**Listeners** | Pointer to [**[]BnsLoadBalancerV1ApiCreateTargetGroupModelListenerModel**](BnsLoadBalancerV1ApiCreateTargetGroupModelListenerModel.md) |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**Members** | Pointer to [**[]BnsLoadBalancerV1ApiCreateTargetGroupModelMemberModel**](BnsLoadBalancerV1ApiCreateTargetGroupModelMemberModel.md) | 연결된 대상 인스턴스 목록 | [optional] 
**AlpnProtocols** | Pointer to [**[]AlpnProtocol**](AlpnProtocol.md) |  | [optional] 

## Methods

### NewBnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel

`func NewBnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel(id string, ) *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel`

NewBnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel instantiates a new BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModelWithDefaults

`func NewBnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModelWithDefaults() *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel`

NewBnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModelWithDefaults instantiates a new BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetOperatingStatus

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) GetOperatingStatus() LoadBalancerOperatingStatus`

GetOperatingStatus returns the OperatingStatus field if non-nil, zero value otherwise.

### GetOperatingStatusOk

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) GetOperatingStatusOk() (*LoadBalancerOperatingStatus, bool)`

GetOperatingStatusOk returns a tuple with the OperatingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingStatus

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) SetOperatingStatus(v LoadBalancerOperatingStatus)`

SetOperatingStatus sets OperatingStatus field to given value.

### HasOperatingStatus

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) HasOperatingStatus() bool`

HasOperatingStatus returns a boolean if a field has been set.

### SetOperatingStatusNil

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) SetOperatingStatusNil(b bool)`

 SetOperatingStatusNil sets the value for OperatingStatus to be an explicit nil

### UnsetOperatingStatus
`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) UnsetOperatingStatus()`

UnsetOperatingStatus ensures that no value is present for OperatingStatus, not even an explicit nil
### GetProtocol

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) GetProtocol() TargetGroupProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) GetProtocolOk() (*TargetGroupProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) SetProtocol(v TargetGroupProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetLoadBalancerAlgorithm

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) GetLoadBalancerAlgorithm() TargetGroupAlgorithm`

GetLoadBalancerAlgorithm returns the LoadBalancerAlgorithm field if non-nil, zero value otherwise.

### GetLoadBalancerAlgorithmOk

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) GetLoadBalancerAlgorithmOk() (*TargetGroupAlgorithm, bool)`

GetLoadBalancerAlgorithmOk returns a tuple with the LoadBalancerAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerAlgorithm

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) SetLoadBalancerAlgorithm(v TargetGroupAlgorithm)`

SetLoadBalancerAlgorithm sets LoadBalancerAlgorithm field to given value.

### HasLoadBalancerAlgorithm

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) HasLoadBalancerAlgorithm() bool`

HasLoadBalancerAlgorithm returns a boolean if a field has been set.

### SetLoadBalancerAlgorithmNil

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) SetLoadBalancerAlgorithmNil(b bool)`

 SetLoadBalancerAlgorithmNil sets the value for LoadBalancerAlgorithm to be an explicit nil

### UnsetLoadBalancerAlgorithm
`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) UnsetLoadBalancerAlgorithm()`

UnsetLoadBalancerAlgorithm ensures that no value is present for LoadBalancerAlgorithm, not even an explicit nil
### GetSessionPersistence

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) GetSessionPersistence() SessionPersistenceModel`

GetSessionPersistence returns the SessionPersistence field if non-nil, zero value otherwise.

### GetSessionPersistenceOk

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) GetSessionPersistenceOk() (*SessionPersistenceModel, bool)`

GetSessionPersistenceOk returns a tuple with the SessionPersistence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionPersistence

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) SetSessionPersistence(v SessionPersistenceModel)`

SetSessionPersistence sets SessionPersistence field to given value.

### HasSessionPersistence

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) HasSessionPersistence() bool`

HasSessionPersistence returns a boolean if a field has been set.

### SetSessionPersistenceNil

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) SetSessionPersistenceNil(b bool)`

 SetSessionPersistenceNil sets the value for SessionPersistence to be an explicit nil

### UnsetSessionPersistence
`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) UnsetSessionPersistence()`

UnsetSessionPersistence ensures that no value is present for SessionPersistence, not even an explicit nil
### GetProjectId

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetLoadBalancers

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) GetLoadBalancers() []BnsLoadBalancerV1ApiCreateTargetGroupModelLoadBalancerModel`

GetLoadBalancers returns the LoadBalancers field if non-nil, zero value otherwise.

### GetLoadBalancersOk

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) GetLoadBalancersOk() (*[]BnsLoadBalancerV1ApiCreateTargetGroupModelLoadBalancerModel, bool)`

GetLoadBalancersOk returns a tuple with the LoadBalancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancers

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) SetLoadBalancers(v []BnsLoadBalancerV1ApiCreateTargetGroupModelLoadBalancerModel)`

SetLoadBalancers sets LoadBalancers field to given value.

### HasLoadBalancers

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) HasLoadBalancers() bool`

HasLoadBalancers returns a boolean if a field has been set.

### SetLoadBalancersNil

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) SetLoadBalancersNil(b bool)`

 SetLoadBalancersNil sets the value for LoadBalancers to be an explicit nil

### UnsetLoadBalancers
`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) UnsetLoadBalancers()`

UnsetLoadBalancers ensures that no value is present for LoadBalancers, not even an explicit nil
### GetListeners

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) GetListeners() []BnsLoadBalancerV1ApiCreateTargetGroupModelListenerModel`

GetListeners returns the Listeners field if non-nil, zero value otherwise.

### GetListenersOk

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) GetListenersOk() (*[]BnsLoadBalancerV1ApiCreateTargetGroupModelListenerModel, bool)`

GetListenersOk returns a tuple with the Listeners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListeners

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) SetListeners(v []BnsLoadBalancerV1ApiCreateTargetGroupModelListenerModel)`

SetListeners sets Listeners field to given value.

### HasListeners

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) HasListeners() bool`

HasListeners returns a boolean if a field has been set.

### SetListenersNil

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) SetListenersNil(b bool)`

 SetListenersNil sets the value for Listeners to be an explicit nil

### UnsetListeners
`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) UnsetListeners()`

UnsetListeners ensures that no value is present for Listeners, not even an explicit nil
### GetCreatedAt

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetMembers

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) GetMembers() []*BnsLoadBalancerV1ApiCreateTargetGroupModelMemberModel`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) GetMembersOk() (*[]*BnsLoadBalancerV1ApiCreateTargetGroupModelMemberModel, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) SetMembers(v []*BnsLoadBalancerV1ApiCreateTargetGroupModelMemberModel)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetAlpnProtocols

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) GetAlpnProtocols() []AlpnProtocol`

GetAlpnProtocols returns the AlpnProtocols field if non-nil, zero value otherwise.

### GetAlpnProtocolsOk

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) GetAlpnProtocolsOk() (*[]AlpnProtocol, bool)`

GetAlpnProtocolsOk returns a tuple with the AlpnProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlpnProtocols

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) SetAlpnProtocols(v []AlpnProtocol)`

SetAlpnProtocols sets AlpnProtocols field to given value.

### HasAlpnProtocols

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) HasAlpnProtocols() bool`

HasAlpnProtocols returns a boolean if a field has been set.

### SetAlpnProtocolsNil

`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) SetAlpnProtocolsNil(b bool)`

 SetAlpnProtocolsNil sets the value for AlpnProtocols to be an explicit nil

### UnsetAlpnProtocols
`func (o *BnsLoadBalancerV1ApiCreateTargetGroupModelTargetGroupModel) UnsetAlpnProtocols()`

UnsetAlpnProtocols ensures that no value is present for AlpnProtocols, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


