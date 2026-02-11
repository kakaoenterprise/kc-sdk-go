# BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 대상 그룹 ID | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Protocol** | Pointer to [**NullableTargetGroupProtocol**](TargetGroupProtocol.md) |  | [optional] 
**Listeners** | Pointer to [**[]BnsLoadBalancerV1ApiListTargetGroupsModelListenerModel**](BnsLoadBalancerV1ApiListTargetGroupsModelListenerModel.md) |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**LoadBalancerAlgorithm** | Pointer to [**NullableTargetGroupAlgorithm**](TargetGroupAlgorithm.md) |  | [optional] 
**ProvisioningStatus** | Pointer to [**NullableProvisioningStatus**](ProvisioningStatus.md) |  | [optional] 
**OperatingStatus** | Pointer to [**NullableLoadBalancerOperatingStatus**](LoadBalancerOperatingStatus.md) |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**LoadBalancerId** | Pointer to **NullableString** |  | [optional] 
**LoadBalancerName** | Pointer to **NullableString** |  | [optional] 
**LoadBalancerProvisioningStatus** | Pointer to [**NullableProvisioningStatus**](ProvisioningStatus.md) |  | [optional] 
**LoadBalancerType** | Pointer to [**NullableLoadBalancerType**](LoadBalancerType.md) |  | [optional] 
**AvailabilityZone** | Pointer to [**NullableAvailabilityZone**](AvailabilityZone.md) |  | [optional] 
**SubnetName** | Pointer to **NullableString** |  | [optional] 
**VpcId** | Pointer to **NullableString** |  | [optional] 
**VpcName** | Pointer to **NullableString** |  | [optional] 
**SubnetId** | Pointer to **NullableString** |  | [optional] 
**HealthMonitor** | Pointer to [**NullableBnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel**](BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel.md) |  | [optional] 
**SessionPersistence** | Pointer to [**NullableSessionPersistenceModel**](SessionPersistenceModel.md) |  | [optional] 
**MemberCount** | Pointer to **NullableInt32** |  | [optional] 
**AlpnProtocols** | Pointer to [**[]AlpnProtocol**](AlpnProtocol.md) |  | [optional] 

## Methods

### NewBnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel

`func NewBnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel(id string, ) *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel`

NewBnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel instantiates a new BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModelWithDefaults

`func NewBnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModelWithDefaults() *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel`

NewBnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModelWithDefaults instantiates a new BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProtocol

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetProtocol() TargetGroupProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetProtocolOk() (*TargetGroupProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetProtocol(v TargetGroupProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetListeners

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetListeners() []BnsLoadBalancerV1ApiListTargetGroupsModelListenerModel`

GetListeners returns the Listeners field if non-nil, zero value otherwise.

### GetListenersOk

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetListenersOk() (*[]BnsLoadBalancerV1ApiListTargetGroupsModelListenerModel, bool)`

GetListenersOk returns a tuple with the Listeners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListeners

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetListeners(v []BnsLoadBalancerV1ApiListTargetGroupsModelListenerModel)`

SetListeners sets Listeners field to given value.

### HasListeners

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) HasListeners() bool`

HasListeners returns a boolean if a field has been set.

### SetListenersNil

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetListenersNil(b bool)`

 SetListenersNil sets the value for Listeners to be an explicit nil

### UnsetListeners
`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) UnsetListeners()`

UnsetListeners ensures that no value is present for Listeners, not even an explicit nil
### GetProjectId

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetLoadBalancerAlgorithm

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetLoadBalancerAlgorithm() TargetGroupAlgorithm`

GetLoadBalancerAlgorithm returns the LoadBalancerAlgorithm field if non-nil, zero value otherwise.

### GetLoadBalancerAlgorithmOk

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetLoadBalancerAlgorithmOk() (*TargetGroupAlgorithm, bool)`

GetLoadBalancerAlgorithmOk returns a tuple with the LoadBalancerAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerAlgorithm

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetLoadBalancerAlgorithm(v TargetGroupAlgorithm)`

SetLoadBalancerAlgorithm sets LoadBalancerAlgorithm field to given value.

### HasLoadBalancerAlgorithm

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) HasLoadBalancerAlgorithm() bool`

HasLoadBalancerAlgorithm returns a boolean if a field has been set.

### SetLoadBalancerAlgorithmNil

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetLoadBalancerAlgorithmNil(b bool)`

 SetLoadBalancerAlgorithmNil sets the value for LoadBalancerAlgorithm to be an explicit nil

### UnsetLoadBalancerAlgorithm
`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) UnsetLoadBalancerAlgorithm()`

UnsetLoadBalancerAlgorithm ensures that no value is present for LoadBalancerAlgorithm, not even an explicit nil
### GetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetOperatingStatus

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetOperatingStatus() LoadBalancerOperatingStatus`

GetOperatingStatus returns the OperatingStatus field if non-nil, zero value otherwise.

### GetOperatingStatusOk

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetOperatingStatusOk() (*LoadBalancerOperatingStatus, bool)`

GetOperatingStatusOk returns a tuple with the OperatingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingStatus

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetOperatingStatus(v LoadBalancerOperatingStatus)`

SetOperatingStatus sets OperatingStatus field to given value.

### HasOperatingStatus

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) HasOperatingStatus() bool`

HasOperatingStatus returns a boolean if a field has been set.

### SetOperatingStatusNil

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetOperatingStatusNil(b bool)`

 SetOperatingStatusNil sets the value for OperatingStatus to be an explicit nil

### UnsetOperatingStatus
`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) UnsetOperatingStatus()`

UnsetOperatingStatus ensures that no value is present for OperatingStatus, not even an explicit nil
### GetCreatedAt

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetLoadBalancerId

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetLoadBalancerId() string`

GetLoadBalancerId returns the LoadBalancerId field if non-nil, zero value otherwise.

### GetLoadBalancerIdOk

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetLoadBalancerIdOk() (*string, bool)`

GetLoadBalancerIdOk returns a tuple with the LoadBalancerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerId

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetLoadBalancerId(v string)`

SetLoadBalancerId sets LoadBalancerId field to given value.

### HasLoadBalancerId

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) HasLoadBalancerId() bool`

HasLoadBalancerId returns a boolean if a field has been set.

### SetLoadBalancerIdNil

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetLoadBalancerIdNil(b bool)`

 SetLoadBalancerIdNil sets the value for LoadBalancerId to be an explicit nil

### UnsetLoadBalancerId
`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) UnsetLoadBalancerId()`

UnsetLoadBalancerId ensures that no value is present for LoadBalancerId, not even an explicit nil
### GetLoadBalancerName

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetLoadBalancerName() string`

GetLoadBalancerName returns the LoadBalancerName field if non-nil, zero value otherwise.

### GetLoadBalancerNameOk

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetLoadBalancerNameOk() (*string, bool)`

GetLoadBalancerNameOk returns a tuple with the LoadBalancerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerName

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetLoadBalancerName(v string)`

SetLoadBalancerName sets LoadBalancerName field to given value.

### HasLoadBalancerName

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) HasLoadBalancerName() bool`

HasLoadBalancerName returns a boolean if a field has been set.

### SetLoadBalancerNameNil

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetLoadBalancerNameNil(b bool)`

 SetLoadBalancerNameNil sets the value for LoadBalancerName to be an explicit nil

### UnsetLoadBalancerName
`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) UnsetLoadBalancerName()`

UnsetLoadBalancerName ensures that no value is present for LoadBalancerName, not even an explicit nil
### GetLoadBalancerProvisioningStatus

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetLoadBalancerProvisioningStatus() ProvisioningStatus`

GetLoadBalancerProvisioningStatus returns the LoadBalancerProvisioningStatus field if non-nil, zero value otherwise.

### GetLoadBalancerProvisioningStatusOk

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetLoadBalancerProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetLoadBalancerProvisioningStatusOk returns a tuple with the LoadBalancerProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerProvisioningStatus

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetLoadBalancerProvisioningStatus(v ProvisioningStatus)`

SetLoadBalancerProvisioningStatus sets LoadBalancerProvisioningStatus field to given value.

### HasLoadBalancerProvisioningStatus

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) HasLoadBalancerProvisioningStatus() bool`

HasLoadBalancerProvisioningStatus returns a boolean if a field has been set.

### SetLoadBalancerProvisioningStatusNil

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetLoadBalancerProvisioningStatusNil(b bool)`

 SetLoadBalancerProvisioningStatusNil sets the value for LoadBalancerProvisioningStatus to be an explicit nil

### UnsetLoadBalancerProvisioningStatus
`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) UnsetLoadBalancerProvisioningStatus()`

UnsetLoadBalancerProvisioningStatus ensures that no value is present for LoadBalancerProvisioningStatus, not even an explicit nil
### GetLoadBalancerType

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetLoadBalancerType() LoadBalancerType`

GetLoadBalancerType returns the LoadBalancerType field if non-nil, zero value otherwise.

### GetLoadBalancerTypeOk

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetLoadBalancerTypeOk() (*LoadBalancerType, bool)`

GetLoadBalancerTypeOk returns a tuple with the LoadBalancerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerType

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetLoadBalancerType(v LoadBalancerType)`

SetLoadBalancerType sets LoadBalancerType field to given value.

### HasLoadBalancerType

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) HasLoadBalancerType() bool`

HasLoadBalancerType returns a boolean if a field has been set.

### SetLoadBalancerTypeNil

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetLoadBalancerTypeNil(b bool)`

 SetLoadBalancerTypeNil sets the value for LoadBalancerType to be an explicit nil

### UnsetLoadBalancerType
`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) UnsetLoadBalancerType()`

UnsetLoadBalancerType ensures that no value is present for LoadBalancerType, not even an explicit nil
### GetAvailabilityZone

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### SetAvailabilityZoneNil

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
### GetSubnetName

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetSubnetName() string`

GetSubnetName returns the SubnetName field if non-nil, zero value otherwise.

### GetSubnetNameOk

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetSubnetNameOk() (*string, bool)`

GetSubnetNameOk returns a tuple with the SubnetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetName

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetSubnetName(v string)`

SetSubnetName sets SubnetName field to given value.

### HasSubnetName

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) HasSubnetName() bool`

HasSubnetName returns a boolean if a field has been set.

### SetSubnetNameNil

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetSubnetNameNil(b bool)`

 SetSubnetNameNil sets the value for SubnetName to be an explicit nil

### UnsetSubnetName
`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) UnsetSubnetName()`

UnsetSubnetName ensures that no value is present for SubnetName, not even an explicit nil
### GetVpcId

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### SetVpcIdNil

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetVpcIdNil(b bool)`

 SetVpcIdNil sets the value for VpcId to be an explicit nil

### UnsetVpcId
`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) UnsetVpcId()`

UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil
### GetVpcName

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetVpcName() string`

GetVpcName returns the VpcName field if non-nil, zero value otherwise.

### GetVpcNameOk

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetVpcNameOk() (*string, bool)`

GetVpcNameOk returns a tuple with the VpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcName

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetVpcName(v string)`

SetVpcName sets VpcName field to given value.

### HasVpcName

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) HasVpcName() bool`

HasVpcName returns a boolean if a field has been set.

### SetVpcNameNil

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetVpcNameNil(b bool)`

 SetVpcNameNil sets the value for VpcName to be an explicit nil

### UnsetVpcName
`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) UnsetVpcName()`

UnsetVpcName ensures that no value is present for VpcName, not even an explicit nil
### GetSubnetId

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### SetSubnetIdNil

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetSubnetIdNil(b bool)`

 SetSubnetIdNil sets the value for SubnetId to be an explicit nil

### UnsetSubnetId
`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) UnsetSubnetId()`

UnsetSubnetId ensures that no value is present for SubnetId, not even an explicit nil
### GetHealthMonitor

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetHealthMonitor() BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel`

GetHealthMonitor returns the HealthMonitor field if non-nil, zero value otherwise.

### GetHealthMonitorOk

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetHealthMonitorOk() (*BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel, bool)`

GetHealthMonitorOk returns a tuple with the HealthMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthMonitor

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetHealthMonitor(v BnsLoadBalancerV1ApiListTargetGroupsModelHealthMonitorModel)`

SetHealthMonitor sets HealthMonitor field to given value.

### HasHealthMonitor

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) HasHealthMonitor() bool`

HasHealthMonitor returns a boolean if a field has been set.

### SetHealthMonitorNil

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetHealthMonitorNil(b bool)`

 SetHealthMonitorNil sets the value for HealthMonitor to be an explicit nil

### UnsetHealthMonitor
`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) UnsetHealthMonitor()`

UnsetHealthMonitor ensures that no value is present for HealthMonitor, not even an explicit nil
### GetSessionPersistence

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetSessionPersistence() SessionPersistenceModel`

GetSessionPersistence returns the SessionPersistence field if non-nil, zero value otherwise.

### GetSessionPersistenceOk

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetSessionPersistenceOk() (*SessionPersistenceModel, bool)`

GetSessionPersistenceOk returns a tuple with the SessionPersistence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionPersistence

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetSessionPersistence(v SessionPersistenceModel)`

SetSessionPersistence sets SessionPersistence field to given value.

### HasSessionPersistence

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) HasSessionPersistence() bool`

HasSessionPersistence returns a boolean if a field has been set.

### SetSessionPersistenceNil

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetSessionPersistenceNil(b bool)`

 SetSessionPersistenceNil sets the value for SessionPersistence to be an explicit nil

### UnsetSessionPersistence
`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) UnsetSessionPersistence()`

UnsetSessionPersistence ensures that no value is present for SessionPersistence, not even an explicit nil
### GetMemberCount

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetMemberCount() int32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetMemberCountOk() (*int32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetMemberCount(v int32)`

SetMemberCount sets MemberCount field to given value.

### HasMemberCount

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) HasMemberCount() bool`

HasMemberCount returns a boolean if a field has been set.

### SetMemberCountNil

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetMemberCountNil(b bool)`

 SetMemberCountNil sets the value for MemberCount to be an explicit nil

### UnsetMemberCount
`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) UnsetMemberCount()`

UnsetMemberCount ensures that no value is present for MemberCount, not even an explicit nil
### GetAlpnProtocols

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetAlpnProtocols() []AlpnProtocol`

GetAlpnProtocols returns the AlpnProtocols field if non-nil, zero value otherwise.

### GetAlpnProtocolsOk

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) GetAlpnProtocolsOk() (*[]AlpnProtocol, bool)`

GetAlpnProtocolsOk returns a tuple with the AlpnProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlpnProtocols

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetAlpnProtocols(v []AlpnProtocol)`

SetAlpnProtocols sets AlpnProtocols field to given value.

### HasAlpnProtocols

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) HasAlpnProtocols() bool`

HasAlpnProtocols returns a boolean if a field has been set.

### SetAlpnProtocolsNil

`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) SetAlpnProtocolsNil(b bool)`

 SetAlpnProtocolsNil sets the value for AlpnProtocols to be an explicit nil

### UnsetAlpnProtocols
`func (o *BnsLoadBalancerV1ApiListTargetGroupsModelTargetGroupModel) UnsetAlpnProtocols()`

UnsetAlpnProtocols ensures that no value is present for AlpnProtocols, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


