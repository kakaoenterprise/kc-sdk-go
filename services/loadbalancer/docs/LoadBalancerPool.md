# LoadBalancerPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 대상 그룹 ID | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Protocol** | Pointer to [**NullableLoadBalancerPoolProtocol**](LoadBalancerPoolProtocol.md) |  | [optional] 
**Listeners** | Pointer to [**[]Listener**](Listener.md) |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**LoadBalancerAlgorithm** | Pointer to [**NullableLoadBalancerPoolAlgorithm**](LoadBalancerPoolAlgorithm.md) |  | [optional] 
**OperatingStatus** | Pointer to [**NullableLoadBalancerOperatingStatus**](LoadBalancerOperatingStatus.md) |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**ProvisioningStatus** | Pointer to [**NullableProvisioningStatus**](ProvisioningStatus.md) |  | [optional] 
**LoadBalancerId** | Pointer to **NullableString** |  | [optional] 
**LoadBalancerName** | Pointer to **NullableString** |  | [optional] 
**LoadBalancerProvisioningStatus** | Pointer to [**NullableProvisioningStatus**](ProvisioningStatus.md) |  | [optional] 
**LoadBalancerType** | Pointer to [**NullableLoadBalancerType**](LoadBalancerType.md) |  | [optional] 
**AvailabilityZone** | Pointer to [**NullableAvailabilityZone**](AvailabilityZone.md) |  | [optional] 
**SubnetName** | Pointer to **NullableString** |  | [optional] 
**VpcId** | Pointer to **NullableString** |  | [optional] 
**VpcName** | Pointer to **NullableString** |  | [optional] 
**SubnetId** | Pointer to **NullableString** |  | [optional] 
**HealthMonitor** | Pointer to [**NullableTargetGroupHealthMonitor**](TargetGroupHealthMonitor.md) |  | [optional] 
**SessionPersistence** | Pointer to [**NullableSessionPersistence**](SessionPersistence.md) |  | [optional] 
**MemberCount** | Pointer to **NullableInt32** |  | [optional] 
**AlpnProtocols** | Pointer to [**[]LoadBalancerAlpnProtocol**](LoadBalancerAlpnProtocol.md) |  | [optional] 

## Methods

### NewLoadBalancerPool

`func NewLoadBalancerPool(id string, ) *LoadBalancerPool`

NewLoadBalancerPool instantiates a new LoadBalancerPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerPoolWithDefaults

`func NewLoadBalancerPoolWithDefaults() *LoadBalancerPool`

NewLoadBalancerPoolWithDefaults instantiates a new LoadBalancerPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LoadBalancerPool) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoadBalancerPool) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoadBalancerPool) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *LoadBalancerPool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoadBalancerPool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoadBalancerPool) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LoadBalancerPool) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *LoadBalancerPool) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *LoadBalancerPool) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *LoadBalancerPool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LoadBalancerPool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LoadBalancerPool) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LoadBalancerPool) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *LoadBalancerPool) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *LoadBalancerPool) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProtocol

`func (o *LoadBalancerPool) GetProtocol() LoadBalancerPoolProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *LoadBalancerPool) GetProtocolOk() (*LoadBalancerPoolProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *LoadBalancerPool) SetProtocol(v LoadBalancerPoolProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *LoadBalancerPool) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### SetProtocolNil

`func (o *LoadBalancerPool) SetProtocolNil(b bool)`

 SetProtocolNil sets the value for Protocol to be an explicit nil

### UnsetProtocol
`func (o *LoadBalancerPool) UnsetProtocol()`

UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
### GetListeners

`func (o *LoadBalancerPool) GetListeners() []Listener`

GetListeners returns the Listeners field if non-nil, zero value otherwise.

### GetListenersOk

`func (o *LoadBalancerPool) GetListenersOk() (*[]Listener, bool)`

GetListenersOk returns a tuple with the Listeners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListeners

`func (o *LoadBalancerPool) SetListeners(v []Listener)`

SetListeners sets Listeners field to given value.

### HasListeners

`func (o *LoadBalancerPool) HasListeners() bool`

HasListeners returns a boolean if a field has been set.

### SetListenersNil

`func (o *LoadBalancerPool) SetListenersNil(b bool)`

 SetListenersNil sets the value for Listeners to be an explicit nil

### UnsetListeners
`func (o *LoadBalancerPool) UnsetListeners()`

UnsetListeners ensures that no value is present for Listeners, not even an explicit nil
### GetProjectId

`func (o *LoadBalancerPool) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *LoadBalancerPool) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *LoadBalancerPool) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *LoadBalancerPool) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *LoadBalancerPool) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *LoadBalancerPool) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetLoadBalancerAlgorithm

`func (o *LoadBalancerPool) GetLoadBalancerAlgorithm() LoadBalancerPoolAlgorithm`

GetLoadBalancerAlgorithm returns the LoadBalancerAlgorithm field if non-nil, zero value otherwise.

### GetLoadBalancerAlgorithmOk

`func (o *LoadBalancerPool) GetLoadBalancerAlgorithmOk() (*LoadBalancerPoolAlgorithm, bool)`

GetLoadBalancerAlgorithmOk returns a tuple with the LoadBalancerAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerAlgorithm

`func (o *LoadBalancerPool) SetLoadBalancerAlgorithm(v LoadBalancerPoolAlgorithm)`

SetLoadBalancerAlgorithm sets LoadBalancerAlgorithm field to given value.

### HasLoadBalancerAlgorithm

`func (o *LoadBalancerPool) HasLoadBalancerAlgorithm() bool`

HasLoadBalancerAlgorithm returns a boolean if a field has been set.

### SetLoadBalancerAlgorithmNil

`func (o *LoadBalancerPool) SetLoadBalancerAlgorithmNil(b bool)`

 SetLoadBalancerAlgorithmNil sets the value for LoadBalancerAlgorithm to be an explicit nil

### UnsetLoadBalancerAlgorithm
`func (o *LoadBalancerPool) UnsetLoadBalancerAlgorithm()`

UnsetLoadBalancerAlgorithm ensures that no value is present for LoadBalancerAlgorithm, not even an explicit nil
### GetOperatingStatus

`func (o *LoadBalancerPool) GetOperatingStatus() LoadBalancerOperatingStatus`

GetOperatingStatus returns the OperatingStatus field if non-nil, zero value otherwise.

### GetOperatingStatusOk

`func (o *LoadBalancerPool) GetOperatingStatusOk() (*LoadBalancerOperatingStatus, bool)`

GetOperatingStatusOk returns a tuple with the OperatingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingStatus

`func (o *LoadBalancerPool) SetOperatingStatus(v LoadBalancerOperatingStatus)`

SetOperatingStatus sets OperatingStatus field to given value.

### HasOperatingStatus

`func (o *LoadBalancerPool) HasOperatingStatus() bool`

HasOperatingStatus returns a boolean if a field has been set.

### SetOperatingStatusNil

`func (o *LoadBalancerPool) SetOperatingStatusNil(b bool)`

 SetOperatingStatusNil sets the value for OperatingStatus to be an explicit nil

### UnsetOperatingStatus
`func (o *LoadBalancerPool) UnsetOperatingStatus()`

UnsetOperatingStatus ensures that no value is present for OperatingStatus, not even an explicit nil
### GetCreatedAt

`func (o *LoadBalancerPool) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LoadBalancerPool) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LoadBalancerPool) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LoadBalancerPool) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *LoadBalancerPool) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *LoadBalancerPool) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *LoadBalancerPool) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LoadBalancerPool) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LoadBalancerPool) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *LoadBalancerPool) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *LoadBalancerPool) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *LoadBalancerPool) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetProvisioningStatus

`func (o *LoadBalancerPool) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *LoadBalancerPool) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *LoadBalancerPool) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *LoadBalancerPool) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *LoadBalancerPool) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *LoadBalancerPool) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetLoadBalancerId

`func (o *LoadBalancerPool) GetLoadBalancerId() string`

GetLoadBalancerId returns the LoadBalancerId field if non-nil, zero value otherwise.

### GetLoadBalancerIdOk

`func (o *LoadBalancerPool) GetLoadBalancerIdOk() (*string, bool)`

GetLoadBalancerIdOk returns a tuple with the LoadBalancerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerId

`func (o *LoadBalancerPool) SetLoadBalancerId(v string)`

SetLoadBalancerId sets LoadBalancerId field to given value.

### HasLoadBalancerId

`func (o *LoadBalancerPool) HasLoadBalancerId() bool`

HasLoadBalancerId returns a boolean if a field has been set.

### SetLoadBalancerIdNil

`func (o *LoadBalancerPool) SetLoadBalancerIdNil(b bool)`

 SetLoadBalancerIdNil sets the value for LoadBalancerId to be an explicit nil

### UnsetLoadBalancerId
`func (o *LoadBalancerPool) UnsetLoadBalancerId()`

UnsetLoadBalancerId ensures that no value is present for LoadBalancerId, not even an explicit nil
### GetLoadBalancerName

`func (o *LoadBalancerPool) GetLoadBalancerName() string`

GetLoadBalancerName returns the LoadBalancerName field if non-nil, zero value otherwise.

### GetLoadBalancerNameOk

`func (o *LoadBalancerPool) GetLoadBalancerNameOk() (*string, bool)`

GetLoadBalancerNameOk returns a tuple with the LoadBalancerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerName

`func (o *LoadBalancerPool) SetLoadBalancerName(v string)`

SetLoadBalancerName sets LoadBalancerName field to given value.

### HasLoadBalancerName

`func (o *LoadBalancerPool) HasLoadBalancerName() bool`

HasLoadBalancerName returns a boolean if a field has been set.

### SetLoadBalancerNameNil

`func (o *LoadBalancerPool) SetLoadBalancerNameNil(b bool)`

 SetLoadBalancerNameNil sets the value for LoadBalancerName to be an explicit nil

### UnsetLoadBalancerName
`func (o *LoadBalancerPool) UnsetLoadBalancerName()`

UnsetLoadBalancerName ensures that no value is present for LoadBalancerName, not even an explicit nil
### GetLoadBalancerProvisioningStatus

`func (o *LoadBalancerPool) GetLoadBalancerProvisioningStatus() ProvisioningStatus`

GetLoadBalancerProvisioningStatus returns the LoadBalancerProvisioningStatus field if non-nil, zero value otherwise.

### GetLoadBalancerProvisioningStatusOk

`func (o *LoadBalancerPool) GetLoadBalancerProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetLoadBalancerProvisioningStatusOk returns a tuple with the LoadBalancerProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerProvisioningStatus

`func (o *LoadBalancerPool) SetLoadBalancerProvisioningStatus(v ProvisioningStatus)`

SetLoadBalancerProvisioningStatus sets LoadBalancerProvisioningStatus field to given value.

### HasLoadBalancerProvisioningStatus

`func (o *LoadBalancerPool) HasLoadBalancerProvisioningStatus() bool`

HasLoadBalancerProvisioningStatus returns a boolean if a field has been set.

### SetLoadBalancerProvisioningStatusNil

`func (o *LoadBalancerPool) SetLoadBalancerProvisioningStatusNil(b bool)`

 SetLoadBalancerProvisioningStatusNil sets the value for LoadBalancerProvisioningStatus to be an explicit nil

### UnsetLoadBalancerProvisioningStatus
`func (o *LoadBalancerPool) UnsetLoadBalancerProvisioningStatus()`

UnsetLoadBalancerProvisioningStatus ensures that no value is present for LoadBalancerProvisioningStatus, not even an explicit nil
### GetLoadBalancerType

`func (o *LoadBalancerPool) GetLoadBalancerType() LoadBalancerType`

GetLoadBalancerType returns the LoadBalancerType field if non-nil, zero value otherwise.

### GetLoadBalancerTypeOk

`func (o *LoadBalancerPool) GetLoadBalancerTypeOk() (*LoadBalancerType, bool)`

GetLoadBalancerTypeOk returns a tuple with the LoadBalancerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerType

`func (o *LoadBalancerPool) SetLoadBalancerType(v LoadBalancerType)`

SetLoadBalancerType sets LoadBalancerType field to given value.

### HasLoadBalancerType

`func (o *LoadBalancerPool) HasLoadBalancerType() bool`

HasLoadBalancerType returns a boolean if a field has been set.

### SetLoadBalancerTypeNil

`func (o *LoadBalancerPool) SetLoadBalancerTypeNil(b bool)`

 SetLoadBalancerTypeNil sets the value for LoadBalancerType to be an explicit nil

### UnsetLoadBalancerType
`func (o *LoadBalancerPool) UnsetLoadBalancerType()`

UnsetLoadBalancerType ensures that no value is present for LoadBalancerType, not even an explicit nil
### GetAvailabilityZone

`func (o *LoadBalancerPool) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *LoadBalancerPool) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *LoadBalancerPool) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *LoadBalancerPool) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### SetAvailabilityZoneNil

`func (o *LoadBalancerPool) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *LoadBalancerPool) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
### GetSubnetName

`func (o *LoadBalancerPool) GetSubnetName() string`

GetSubnetName returns the SubnetName field if non-nil, zero value otherwise.

### GetSubnetNameOk

`func (o *LoadBalancerPool) GetSubnetNameOk() (*string, bool)`

GetSubnetNameOk returns a tuple with the SubnetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetName

`func (o *LoadBalancerPool) SetSubnetName(v string)`

SetSubnetName sets SubnetName field to given value.

### HasSubnetName

`func (o *LoadBalancerPool) HasSubnetName() bool`

HasSubnetName returns a boolean if a field has been set.

### SetSubnetNameNil

`func (o *LoadBalancerPool) SetSubnetNameNil(b bool)`

 SetSubnetNameNil sets the value for SubnetName to be an explicit nil

### UnsetSubnetName
`func (o *LoadBalancerPool) UnsetSubnetName()`

UnsetSubnetName ensures that no value is present for SubnetName, not even an explicit nil
### GetVpcId

`func (o *LoadBalancerPool) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *LoadBalancerPool) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *LoadBalancerPool) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *LoadBalancerPool) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### SetVpcIdNil

`func (o *LoadBalancerPool) SetVpcIdNil(b bool)`

 SetVpcIdNil sets the value for VpcId to be an explicit nil

### UnsetVpcId
`func (o *LoadBalancerPool) UnsetVpcId()`

UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil
### GetVpcName

`func (o *LoadBalancerPool) GetVpcName() string`

GetVpcName returns the VpcName field if non-nil, zero value otherwise.

### GetVpcNameOk

`func (o *LoadBalancerPool) GetVpcNameOk() (*string, bool)`

GetVpcNameOk returns a tuple with the VpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcName

`func (o *LoadBalancerPool) SetVpcName(v string)`

SetVpcName sets VpcName field to given value.

### HasVpcName

`func (o *LoadBalancerPool) HasVpcName() bool`

HasVpcName returns a boolean if a field has been set.

### SetVpcNameNil

`func (o *LoadBalancerPool) SetVpcNameNil(b bool)`

 SetVpcNameNil sets the value for VpcName to be an explicit nil

### UnsetVpcName
`func (o *LoadBalancerPool) UnsetVpcName()`

UnsetVpcName ensures that no value is present for VpcName, not even an explicit nil
### GetSubnetId

`func (o *LoadBalancerPool) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *LoadBalancerPool) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *LoadBalancerPool) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *LoadBalancerPool) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### SetSubnetIdNil

`func (o *LoadBalancerPool) SetSubnetIdNil(b bool)`

 SetSubnetIdNil sets the value for SubnetId to be an explicit nil

### UnsetSubnetId
`func (o *LoadBalancerPool) UnsetSubnetId()`

UnsetSubnetId ensures that no value is present for SubnetId, not even an explicit nil
### GetHealthMonitor

`func (o *LoadBalancerPool) GetHealthMonitor() TargetGroupHealthMonitor`

GetHealthMonitor returns the HealthMonitor field if non-nil, zero value otherwise.

### GetHealthMonitorOk

`func (o *LoadBalancerPool) GetHealthMonitorOk() (*TargetGroupHealthMonitor, bool)`

GetHealthMonitorOk returns a tuple with the HealthMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthMonitor

`func (o *LoadBalancerPool) SetHealthMonitor(v TargetGroupHealthMonitor)`

SetHealthMonitor sets HealthMonitor field to given value.

### HasHealthMonitor

`func (o *LoadBalancerPool) HasHealthMonitor() bool`

HasHealthMonitor returns a boolean if a field has been set.

### SetHealthMonitorNil

`func (o *LoadBalancerPool) SetHealthMonitorNil(b bool)`

 SetHealthMonitorNil sets the value for HealthMonitor to be an explicit nil

### UnsetHealthMonitor
`func (o *LoadBalancerPool) UnsetHealthMonitor()`

UnsetHealthMonitor ensures that no value is present for HealthMonitor, not even an explicit nil
### GetSessionPersistence

`func (o *LoadBalancerPool) GetSessionPersistence() SessionPersistence`

GetSessionPersistence returns the SessionPersistence field if non-nil, zero value otherwise.

### GetSessionPersistenceOk

`func (o *LoadBalancerPool) GetSessionPersistenceOk() (*SessionPersistence, bool)`

GetSessionPersistenceOk returns a tuple with the SessionPersistence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionPersistence

`func (o *LoadBalancerPool) SetSessionPersistence(v SessionPersistence)`

SetSessionPersistence sets SessionPersistence field to given value.

### HasSessionPersistence

`func (o *LoadBalancerPool) HasSessionPersistence() bool`

HasSessionPersistence returns a boolean if a field has been set.

### SetSessionPersistenceNil

`func (o *LoadBalancerPool) SetSessionPersistenceNil(b bool)`

 SetSessionPersistenceNil sets the value for SessionPersistence to be an explicit nil

### UnsetSessionPersistence
`func (o *LoadBalancerPool) UnsetSessionPersistence()`

UnsetSessionPersistence ensures that no value is present for SessionPersistence, not even an explicit nil
### GetMemberCount

`func (o *LoadBalancerPool) GetMemberCount() int32`

GetMemberCount returns the MemberCount field if non-nil, zero value otherwise.

### GetMemberCountOk

`func (o *LoadBalancerPool) GetMemberCountOk() (*int32, bool)`

GetMemberCountOk returns a tuple with the MemberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberCount

`func (o *LoadBalancerPool) SetMemberCount(v int32)`

SetMemberCount sets MemberCount field to given value.

### HasMemberCount

`func (o *LoadBalancerPool) HasMemberCount() bool`

HasMemberCount returns a boolean if a field has been set.

### SetMemberCountNil

`func (o *LoadBalancerPool) SetMemberCountNil(b bool)`

 SetMemberCountNil sets the value for MemberCount to be an explicit nil

### UnsetMemberCount
`func (o *LoadBalancerPool) UnsetMemberCount()`

UnsetMemberCount ensures that no value is present for MemberCount, not even an explicit nil
### GetAlpnProtocols

`func (o *LoadBalancerPool) GetAlpnProtocols() []*LoadBalancerAlpnProtocol`

GetAlpnProtocols returns the AlpnProtocols field if non-nil, zero value otherwise.

### GetAlpnProtocolsOk

`func (o *LoadBalancerPool) GetAlpnProtocolsOk() (*[]*LoadBalancerAlpnProtocol, bool)`

GetAlpnProtocolsOk returns a tuple with the AlpnProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlpnProtocols

`func (o *LoadBalancerPool) SetAlpnProtocols(v []*LoadBalancerAlpnProtocol)`

SetAlpnProtocols sets AlpnProtocols field to given value.

### HasAlpnProtocols

`func (o *LoadBalancerPool) HasAlpnProtocols() bool`

HasAlpnProtocols returns a boolean if a field has been set.

### SetAlpnProtocolsNil

`func (o *LoadBalancerPool) SetAlpnProtocolsNil(b bool)`

 SetAlpnProtocolsNil sets the value for AlpnProtocols to be an explicit nil

### UnsetAlpnProtocols
`func (o *LoadBalancerPool) UnsetAlpnProtocols()`

UnsetAlpnProtocols ensures that no value is present for AlpnProtocols, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


