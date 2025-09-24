# BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 대상 그룹 ID | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Protocol** | Pointer to [**NullableTargetGroupProtocol**](TargetGroupProtocol.md) |  | [optional] 
**Listeners** | Pointer to [**[]BnsLoadBalancerV1ApiGetTargetGroupModelListenerModel**](BnsLoadBalancerV1ApiGetTargetGroupModelListenerModel.md) |  | [optional] 
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
**HealthMonitor** | Pointer to [**NullableBnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel**](BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel.md) |  | [optional] 
**SessionPersistence** | Pointer to [**NullableSessionPersistenceModel**](SessionPersistenceModel.md) |  | [optional] 
**MemberCount** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewBnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel

`func NewBnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel(id string, ) *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel`

NewBnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel instantiates a new BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModelWithDefaults

`func NewBnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModelWithDefaults() *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel`

NewBnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModelWithDefaults instantiates a new BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProtocol

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetProtocol() TargetGroupProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetProtocolOk() (*TargetGroupProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) SetProtocol(v TargetGroupProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetListeners

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetListeners() []BnsLoadBalancerV1ApiGetTargetGroupModelListenerModel`

GetListeners returns the Listeners field if non-nil, zero value otherwise.

### GetListenersOk

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetListenersOk() (*[]BnsLoadBalancerV1ApiGetTargetGroupModelListenerModel, bool)`

GetListenersOk returns a tuple with the Listeners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListeners

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) SetListeners(v []BnsLoadBalancerV1ApiGetTargetGroupModelListenerModel)`

SetListeners sets Listeners field to given value.

### HasListeners

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) HasListeners() bool`

HasListeners returns a boolean if a field has been set.

### SetListenersNil

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) SetListenersNil(b bool)`

 SetListenersNil sets the value for Listeners to be an explicit nil

### UnsetListeners
`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) UnsetListeners()`

UnsetListeners ensures that no value is present for Listeners, not even an explicit nil
### GetProjectId

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetLoadBalancerAlgorithm

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetLoadBalancerAlgorithm() TargetGroupAlgorithm`

GetLoadBalancerAlgorithm returns the LoadBalancerAlgorithm field if non-nil, zero value otherwise.

### GetLoadBalancerAlgorithmOk

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetLoadBalancerAlgorithmOk() (*TargetGroupAlgorithm, bool)`

GetLoadBalancerAlgorithmOk returns a tuple with the LoadBalancerAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerAlgorithm

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) SetLoadBalancerAlgorithm(v TargetGroupAlgorithm)`

SetLoadBalancerAlgorithm sets LoadBalancerAlgorithm field to given value.

### HasLoadBalancerAlgorithm

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) HasLoadBalancerAlgorithm() bool`

HasLoadBalancerAlgorithm returns a boolean if a field has been set.

### SetLoadBalancerAlgorithmNil

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) SetLoadBalancerAlgorithmNil(b bool)`

 SetLoadBalancerAlgorithmNil sets the value for LoadBalancerAlgorithm to be an explicit nil

### UnsetLoadBalancerAlgorithm
`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) UnsetLoadBalancerAlgorithm()`

UnsetLoadBalancerAlgorithm ensures that no value is present for LoadBalancerAlgorithm, not even an explicit nil
### GetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetOperatingStatus

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetOperatingStatus() LoadBalancerOperatingStatus`

GetOperatingStatus returns the OperatingStatus field if non-nil, zero value otherwise.

### GetOperatingStatusOk

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetOperatingStatusOk() (*LoadBalancerOperatingStatus, bool)`

GetOperatingStatusOk returns a tuple with the OperatingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingStatus

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) SetOperatingStatus(v LoadBalancerOperatingStatus)`

SetOperatingStatus sets OperatingStatus field to given value.

### HasOperatingStatus

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) HasOperatingStatus() bool`

HasOperatingStatus returns a boolean if a field has been set.

### SetOperatingStatusNil

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) SetOperatingStatusNil(b bool)`

 SetOperatingStatusNil sets the value for OperatingStatus to be an explicit nil

### UnsetOperatingStatus
`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) UnsetOperatingStatus()`

UnsetOperatingStatus ensures that no value is present for OperatingStatus, not even an explicit nil
### GetCreatedAt

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetLoadBalancerId

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetLoadBalancerId() string`

GetLoadBalancerId returns the LoadBalancerId field if non-nil, zero value otherwise.

### GetLoadBalancerIdOk

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetLoadBalancerIdOk() (*string, bool)`

GetLoadBalancerIdOk returns a tuple with the LoadBalancerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerId

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) SetLoadBalancerId(v string)`

SetLoadBalancerId sets LoadBalancerId field to given value.

### HasLoadBalancerId

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) HasLoadBalancerId() bool`

HasLoadBalancerId returns a boolean if a field has been set.

### SetLoadBalancerIdNil

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) SetLoadBalancerIdNil(b bool)`

 SetLoadBalancerIdNil sets the value for LoadBalancerId to be an explicit nil

### UnsetLoadBalancerId
`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) UnsetLoadBalancerId()`

UnsetLoadBalancerId ensures that no value is present for LoadBalancerId, not even an explicit nil
### GetLoadBalancerName

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetLoadBalancerName() string`

GetLoadBalancerName returns the LoadBalancerName field if non-nil, zero value otherwise.

### GetLoadBalancerNameOk

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetLoadBalancerNameOk() (*string, bool)`

GetLoadBalancerNameOk returns a tuple with the LoadBalancerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerName

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) SetLoadBalancerName(v string)`

SetLoadBalancerName sets LoadBalancerName field to given value.

### HasLoadBalancerName

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) HasLoadBalancerName() bool`

HasLoadBalancerName returns a boolean if a field has been set.

### SetLoadBalancerNameNil

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) SetLoadBalancerNameNil(b bool)`

 SetLoadBalancerNameNil sets the value for LoadBalancerName to be an explicit nil

### UnsetLoadBalancerName
`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) UnsetLoadBalancerName()`

UnsetLoadBalancerName ensures that no value is present for LoadBalancerName, not even an explicit nil
### GetLoadBalancerProvisioningStatus

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetLoadBalancerProvisioningStatus() ProvisioningStatus`

GetLoadBalancerProvisioningStatus returns the LoadBalancerProvisioningStatus field if non-nil, zero value otherwise.

### GetLoadBalancerProvisioningStatusOk

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetLoadBalancerProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetLoadBalancerProvisioningStatusOk returns a tuple with the LoadBalancerProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerProvisioningStatus

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) SetLoadBalancerProvisioningStatus(v ProvisioningStatus)`

SetLoadBalancerProvisioningStatus sets LoadBalancerProvisioningStatus field to given value.

### HasLoadBalancerProvisioningStatus

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) HasLoadBalancerProvisioningStatus() bool`

HasLoadBalancerProvisioningStatus returns a boolean if a field has been set.

### SetLoadBalancerProvisioningStatusNil

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) SetLoadBalancerProvisioningStatusNil(b bool)`

 SetLoadBalancerProvisioningStatusNil sets the value for LoadBalancerProvisioningStatus to be an explicit nil

### UnsetLoadBalancerProvisioningStatus
`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) UnsetLoadBalancerProvisioningStatus()`

UnsetLoadBalancerProvisioningStatus ensures that no value is present for LoadBalancerProvisioningStatus, not even an explicit nil
### GetLoadBalancerType

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetLoadBalancerType() LoadBalancerType`

GetLoadBalancerType returns the LoadBalancerType field if non-nil, zero value otherwise.

### GetLoadBalancerTypeOk

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetLoadBalancerTypeOk() (*LoadBalancerType, bool)`

GetLoadBalancerTypeOk returns a tuple with the LoadBalancerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerType

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) SetLoadBalancerType(v LoadBalancerType)`

SetLoadBalancerType sets LoadBalancerType field to given value.

### HasLoadBalancerType

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) HasLoadBalancerType() bool`

HasLoadBalancerType returns a boolean if a field has been set.

### SetLoadBalancerTypeNil

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) SetLoadBalancerTypeNil(b bool)`

 SetLoadBalancerTypeNil sets the value for LoadBalancerType to be an explicit nil

### UnsetLoadBalancerType
`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) UnsetLoadBalancerType()`

UnsetLoadBalancerType ensures that no value is present for LoadBalancerType, not even an explicit nil
### GetAvailabilityZone

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### SetAvailabilityZoneNil

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
### GetSubnetName

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetSubnetName() string`

GetSubnetName returns the SubnetName field if non-nil, zero value otherwise.

### GetSubnetNameOk

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetSubnetNameOk() (*string, bool)`

GetSubnetNameOk returns a tuple with the SubnetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetName

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) SetSubnetName(v string)`

SetSubnetName sets SubnetName field to given value.

### HasSubnetName

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) HasSubnetName() bool`

HasSubnetName returns a boolean if a field has been set.

### SetSubnetNameNil

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) SetSubnetNameNil(b bool)`

 SetSubnetNameNil sets the value for SubnetName to be an explicit nil

### UnsetSubnetName
`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) UnsetSubnetName()`

UnsetSubnetName ensures that no value is present for SubnetName, not even an explicit nil
### GetVpcId

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### SetVpcIdNil

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) SetVpcIdNil(b bool)`

 SetVpcIdNil sets the value for VpcId to be an explicit nil

### UnsetVpcId
`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) UnsetVpcId()`

UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil
### GetVpcName

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetVpcName() string`

GetVpcName returns the VpcName field if non-nil, zero value otherwise.

### GetVpcNameOk

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetVpcNameOk() (*string, bool)`

GetVpcNameOk returns a tuple with the VpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcName

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) SetVpcName(v string)`

SetVpcName sets VpcName field to given value.

### HasVpcName

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) HasVpcName() bool`

HasVpcName returns a boolean if a field has been set.

### SetVpcNameNil

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) SetVpcNameNil(b bool)`

 SetVpcNameNil sets the value for VpcName to be an explicit nil

### UnsetVpcName
`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) UnsetVpcName()`

UnsetVpcName ensures that no value is present for VpcName, not even an explicit nil
### GetSubnetId

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### SetSubnetIdNil

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) SetSubnetIdNil(b bool)`

 SetSubnetIdNil sets the value for SubnetId to be an explicit nil

### UnsetSubnetId
`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) UnsetSubnetId()`

UnsetSubnetId ensures that no value is present for SubnetId, not even an explicit nil
### GetHealthMonitor

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetHealthMonitor() BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel`

GetHealthMonitor returns the HealthMonitor field if non-nil, zero value otherwise.

### GetHealthMonitorOk

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetHealthMonitorOk() (*BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel, bool)`

GetHealthMonitorOk returns a tuple with the HealthMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthMonitor

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) SetHealthMonitor(v BnsLoadBalancerV1ApiGetTargetGroupModelHealthMonitorModel)`

SetHealthMonitor sets HealthMonitor field to given value.

### HasHealthMonitor

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) HasHealthMonitor() bool`

HasHealthMonitor returns a boolean if a field has been set.

### SetHealthMonitorNil

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) SetHealthMonitorNil(b bool)`

 SetHealthMonitorNil sets the value for HealthMonitor to be an explicit nil

### UnsetHealthMonitor
`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) UnsetHealthMonitor()`

UnsetHealthMonitor ensures that no value is present for HealthMonitor, not even an explicit nil
### GetSessionPersistence

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetSessionPersistence() SessionPersistenceModel`

GetSessionPersistence returns the SessionPersistence field if non-nil, zero value otherwise.

### GetSessionPersistenceOk

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetSessionPersistenceOk() (*SessionPersistenceModel, bool)`

GetSessionPersistenceOk returns a tuple with the SessionPersistence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionPersistence

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) SetSessionPersistence(v SessionPersistenceModel)`

SetSessionPersistence sets SessionPersistence field to given value.

### HasSessionPersistence

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) HasSessionPersistence() bool`

HasSessionPersistence returns a boolean if a field has been set.

### SetSessionPersistenceNil

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) SetSessionPersistenceNil(b bool)`

 SetSessionPersistenceNil sets the value for SessionPersistence to be an explicit nil

### UnsetSessionPersistence
`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) UnsetSessionPersistence()`

UnsetSessionPersistence ensures that no value is present for SessionPersistence, not even an explicit nil
### GetMemberCount

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetMemberCount() int32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) GetMemberCountOk() (*int32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) SetMemberCount(v int32)`

SetMemberCount sets MemberCount field to given value.

### HasMemberCount

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) HasMemberCount() bool`

HasMemberCount returns a boolean if a field has been set.

### SetMemberCountNil

`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) SetMemberCountNil(b bool)`

 SetMemberCountNil sets the value for MemberCount to be an explicit nil

### UnsetMemberCount
`func (o *BnsLoadBalancerV1ApiGetTargetGroupModelTargetGroupModel) UnsetMemberCount()`

UnsetMemberCount ensures that no value is present for MemberCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


