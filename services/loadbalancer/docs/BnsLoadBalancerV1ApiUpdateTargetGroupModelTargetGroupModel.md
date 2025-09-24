# BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel

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
**LoadBalancers** | Pointer to [**[]BnsLoadBalancerV1ApiUpdateTargetGroupModelLoadBalancerModel**](BnsLoadBalancerV1ApiUpdateTargetGroupModelLoadBalancerModel.md) |  | [optional] 
**Listeners** | Pointer to [**[]BnsLoadBalancerV1ApiUpdateTargetGroupModelListenerModel**](BnsLoadBalancerV1ApiUpdateTargetGroupModelListenerModel.md) |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**Members** | Pointer to [**[]BnsLoadBalancerV1ApiUpdateTargetGroupModelMemberModel**](BnsLoadBalancerV1ApiUpdateTargetGroupModelMemberModel.md) | 대상 그룹에 속한 멤버(서버) 목록 | [optional] 

## Methods

### NewBnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel

`func NewBnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel(id string, ) *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel`

NewBnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel instantiates a new BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModelWithDefaults

`func NewBnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModelWithDefaults() *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel`

NewBnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModelWithDefaults instantiates a new BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetOperatingStatus

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) GetOperatingStatus() LoadBalancerOperatingStatus`

GetOperatingStatus returns the OperatingStatus field if non-nil, zero value otherwise.

### GetOperatingStatusOk

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) GetOperatingStatusOk() (*LoadBalancerOperatingStatus, bool)`

GetOperatingStatusOk returns a tuple with the OperatingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingStatus

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) SetOperatingStatus(v LoadBalancerOperatingStatus)`

SetOperatingStatus sets OperatingStatus field to given value.

### HasOperatingStatus

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) HasOperatingStatus() bool`

HasOperatingStatus returns a boolean if a field has been set.

### SetOperatingStatusNil

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) SetOperatingStatusNil(b bool)`

 SetOperatingStatusNil sets the value for OperatingStatus to be an explicit nil

### UnsetOperatingStatus
`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) UnsetOperatingStatus()`

UnsetOperatingStatus ensures that no value is present for OperatingStatus, not even an explicit nil
### GetProtocol

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) GetProtocol() TargetGroupProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) GetProtocolOk() (*TargetGroupProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) SetProtocol(v TargetGroupProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetLoadBalancerAlgorithm

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) GetLoadBalancerAlgorithm() TargetGroupAlgorithm`

GetLoadBalancerAlgorithm returns the LoadBalancerAlgorithm field if non-nil, zero value otherwise.

### GetLoadBalancerAlgorithmOk

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) GetLoadBalancerAlgorithmOk() (*TargetGroupAlgorithm, bool)`

GetLoadBalancerAlgorithmOk returns a tuple with the LoadBalancerAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerAlgorithm

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) SetLoadBalancerAlgorithm(v TargetGroupAlgorithm)`

SetLoadBalancerAlgorithm sets LoadBalancerAlgorithm field to given value.

### HasLoadBalancerAlgorithm

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) HasLoadBalancerAlgorithm() bool`

HasLoadBalancerAlgorithm returns a boolean if a field has been set.

### SetLoadBalancerAlgorithmNil

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) SetLoadBalancerAlgorithmNil(b bool)`

 SetLoadBalancerAlgorithmNil sets the value for LoadBalancerAlgorithm to be an explicit nil

### UnsetLoadBalancerAlgorithm
`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) UnsetLoadBalancerAlgorithm()`

UnsetLoadBalancerAlgorithm ensures that no value is present for LoadBalancerAlgorithm, not even an explicit nil
### GetSessionPersistence

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) GetSessionPersistence() SessionPersistenceModel`

GetSessionPersistence returns the SessionPersistence field if non-nil, zero value otherwise.

### GetSessionPersistenceOk

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) GetSessionPersistenceOk() (*SessionPersistenceModel, bool)`

GetSessionPersistenceOk returns a tuple with the SessionPersistence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionPersistence

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) SetSessionPersistence(v SessionPersistenceModel)`

SetSessionPersistence sets SessionPersistence field to given value.

### HasSessionPersistence

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) HasSessionPersistence() bool`

HasSessionPersistence returns a boolean if a field has been set.

### SetSessionPersistenceNil

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) SetSessionPersistenceNil(b bool)`

 SetSessionPersistenceNil sets the value for SessionPersistence to be an explicit nil

### UnsetSessionPersistence
`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) UnsetSessionPersistence()`

UnsetSessionPersistence ensures that no value is present for SessionPersistence, not even an explicit nil
### GetProjectId

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetLoadBalancers

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) GetLoadBalancers() []BnsLoadBalancerV1ApiUpdateTargetGroupModelLoadBalancerModel`

GetLoadBalancers returns the LoadBalancers field if non-nil, zero value otherwise.

### GetLoadBalancersOk

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) GetLoadBalancersOk() (*[]BnsLoadBalancerV1ApiUpdateTargetGroupModelLoadBalancerModel, bool)`

GetLoadBalancersOk returns a tuple with the LoadBalancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancers

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) SetLoadBalancers(v []BnsLoadBalancerV1ApiUpdateTargetGroupModelLoadBalancerModel)`

SetLoadBalancers sets LoadBalancers field to given value.

### HasLoadBalancers

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) HasLoadBalancers() bool`

HasLoadBalancers returns a boolean if a field has been set.

### SetLoadBalancersNil

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) SetLoadBalancersNil(b bool)`

 SetLoadBalancersNil sets the value for LoadBalancers to be an explicit nil

### UnsetLoadBalancers
`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) UnsetLoadBalancers()`

UnsetLoadBalancers ensures that no value is present for LoadBalancers, not even an explicit nil
### GetListeners

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) GetListeners() []BnsLoadBalancerV1ApiUpdateTargetGroupModelListenerModel`

GetListeners returns the Listeners field if non-nil, zero value otherwise.

### GetListenersOk

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) GetListenersOk() (*[]BnsLoadBalancerV1ApiUpdateTargetGroupModelListenerModel, bool)`

GetListenersOk returns a tuple with the Listeners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListeners

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) SetListeners(v []BnsLoadBalancerV1ApiUpdateTargetGroupModelListenerModel)`

SetListeners sets Listeners field to given value.

### HasListeners

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) HasListeners() bool`

HasListeners returns a boolean if a field has been set.

### SetListenersNil

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) SetListenersNil(b bool)`

 SetListenersNil sets the value for Listeners to be an explicit nil

### UnsetListeners
`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) UnsetListeners()`

UnsetListeners ensures that no value is present for Listeners, not even an explicit nil
### GetCreatedAt

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetMembers

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) GetMembers() []*BnsLoadBalancerV1ApiUpdateTargetGroupModelMemberModel`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) GetMembersOk() (*[]*BnsLoadBalancerV1ApiUpdateTargetGroupModelMemberModel, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) SetMembers(v []*BnsLoadBalancerV1ApiUpdateTargetGroupModelMemberModel)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *BnsLoadBalancerV1ApiUpdateTargetGroupModelTargetGroupModel) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


