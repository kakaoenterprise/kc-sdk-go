# BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 대상 ID | 
**Name** | Pointer to **NullableString** |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**TargetGroupId** | Pointer to **NullableString** |  | [optional] 
**IpAddress** | Pointer to **NullableString** |  | [optional] 
**ProtocolPort** | Pointer to **NullableInt32** |  | [optional] 
**Weight** | Pointer to **NullableInt32** |  | [optional] 
**OperatingStatus** | [**LoadBalancerOperatingStatus**](LoadBalancerOperatingStatus.md) | 운영 상태 | 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**ProvisioningStatus** | Pointer to [**NullableProvisioningStatus**](ProvisioningStatus.md) |  | [optional] 
**MonitorPort** | Pointer to **NullableInt32** |  | [optional] 
**NetworkInterfaceId** | Pointer to **NullableString** |  | [optional] 
**InstanceId** | Pointer to **NullableString** |  | [optional] 
**InstanceName** | Pointer to **NullableString** |  | [optional] 
**VpcId** | Pointer to **NullableString** |  | [optional] 
**Subnet** | [**BnsLoadBalancerV1ApiListTargetsInTargetGroupModelSubnetModel**](BnsLoadBalancerV1ApiListTargetsInTargetGroupModelSubnetModel.md) | 서브넷 정보 | 
**SecurityGroups** | Pointer to [**[]SecurityGroupModel**](SecurityGroupModel.md) |  | [optional] 

## Methods

### NewBnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel

`func NewBnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel(id string, operatingStatus LoadBalancerOperatingStatus, subnet BnsLoadBalancerV1ApiListTargetsInTargetGroupModelSubnetModel, ) *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel`

NewBnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel instantiates a new BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModelWithDefaults

`func NewBnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModelWithDefaults() *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel`

NewBnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModelWithDefaults instantiates a new BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProjectId

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetTargetGroupId

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) GetTargetGroupId() string`

GetTargetGroupId returns the TargetGroupId field if non-nil, zero value otherwise.

### GetTargetGroupIdOk

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) GetTargetGroupIdOk() (*string, bool)`

GetTargetGroupIdOk returns a tuple with the TargetGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroupId

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) SetTargetGroupId(v string)`

SetTargetGroupId sets TargetGroupId field to given value.

### HasTargetGroupId

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) HasTargetGroupId() bool`

HasTargetGroupId returns a boolean if a field has been set.

### SetTargetGroupIdNil

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) SetTargetGroupIdNil(b bool)`

 SetTargetGroupIdNil sets the value for TargetGroupId to be an explicit nil

### UnsetTargetGroupId
`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) UnsetTargetGroupId()`

UnsetTargetGroupId ensures that no value is present for TargetGroupId, not even an explicit nil
### GetIpAddress

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetProtocolPort

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) GetProtocolPort() int32`

GetProtocolPort returns the ProtocolPort field if non-nil, zero value otherwise.

### GetProtocolPortOk

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) GetProtocolPortOk() (*int32, bool)`

GetProtocolPortOk returns a tuple with the ProtocolPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolPort

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) SetProtocolPort(v int32)`

SetProtocolPort sets ProtocolPort field to given value.

### HasProtocolPort

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) HasProtocolPort() bool`

HasProtocolPort returns a boolean if a field has been set.

### SetProtocolPortNil

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) SetProtocolPortNil(b bool)`

 SetProtocolPortNil sets the value for ProtocolPort to be an explicit nil

### UnsetProtocolPort
`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) UnsetProtocolPort()`

UnsetProtocolPort ensures that no value is present for ProtocolPort, not even an explicit nil
### GetWeight

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetOperatingStatus

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) GetOperatingStatus() LoadBalancerOperatingStatus`

GetOperatingStatus returns the OperatingStatus field if non-nil, zero value otherwise.

### GetOperatingStatusOk

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) GetOperatingStatusOk() (*LoadBalancerOperatingStatus, bool)`

GetOperatingStatusOk returns a tuple with the OperatingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingStatus

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) SetOperatingStatus(v LoadBalancerOperatingStatus)`

SetOperatingStatus sets OperatingStatus field to given value.


### GetCreatedAt

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetMonitorPort

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) GetMonitorPort() int32`

GetMonitorPort returns the MonitorPort field if non-nil, zero value otherwise.

### GetMonitorPortOk

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) GetMonitorPortOk() (*int32, bool)`

GetMonitorPortOk returns a tuple with the MonitorPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorPort

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) SetMonitorPort(v int32)`

SetMonitorPort sets MonitorPort field to given value.

### HasMonitorPort

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) HasMonitorPort() bool`

HasMonitorPort returns a boolean if a field has been set.

### SetMonitorPortNil

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) SetMonitorPortNil(b bool)`

 SetMonitorPortNil sets the value for MonitorPort to be an explicit nil

### UnsetMonitorPort
`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) UnsetMonitorPort()`

UnsetMonitorPort ensures that no value is present for MonitorPort, not even an explicit nil
### GetNetworkInterfaceId

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) GetNetworkInterfaceId() string`

GetNetworkInterfaceId returns the NetworkInterfaceId field if non-nil, zero value otherwise.

### GetNetworkInterfaceIdOk

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) GetNetworkInterfaceIdOk() (*string, bool)`

GetNetworkInterfaceIdOk returns a tuple with the NetworkInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaceId

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) SetNetworkInterfaceId(v string)`

SetNetworkInterfaceId sets NetworkInterfaceId field to given value.

### HasNetworkInterfaceId

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) HasNetworkInterfaceId() bool`

HasNetworkInterfaceId returns a boolean if a field has been set.

### SetNetworkInterfaceIdNil

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) SetNetworkInterfaceIdNil(b bool)`

 SetNetworkInterfaceIdNil sets the value for NetworkInterfaceId to be an explicit nil

### UnsetNetworkInterfaceId
`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) UnsetNetworkInterfaceId()`

UnsetNetworkInterfaceId ensures that no value is present for NetworkInterfaceId, not even an explicit nil
### GetInstanceId

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### SetInstanceIdNil

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) SetInstanceIdNil(b bool)`

 SetInstanceIdNil sets the value for InstanceId to be an explicit nil

### UnsetInstanceId
`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) UnsetInstanceId()`

UnsetInstanceId ensures that no value is present for InstanceId, not even an explicit nil
### GetInstanceName

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.

### HasInstanceName

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) HasInstanceName() bool`

HasInstanceName returns a boolean if a field has been set.

### SetInstanceNameNil

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) SetInstanceNameNil(b bool)`

 SetInstanceNameNil sets the value for InstanceName to be an explicit nil

### UnsetInstanceName
`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) UnsetInstanceName()`

UnsetInstanceName ensures that no value is present for InstanceName, not even an explicit nil
### GetVpcId

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### SetVpcIdNil

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) SetVpcIdNil(b bool)`

 SetVpcIdNil sets the value for VpcId to be an explicit nil

### UnsetVpcId
`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) UnsetVpcId()`

UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil
### GetSubnet

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) GetSubnet() BnsLoadBalancerV1ApiListTargetsInTargetGroupModelSubnetModel`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) GetSubnetOk() (*BnsLoadBalancerV1ApiListTargetsInTargetGroupModelSubnetModel, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) SetSubnet(v BnsLoadBalancerV1ApiListTargetsInTargetGroupModelSubnetModel)`

SetSubnet sets Subnet field to given value.


### GetSecurityGroups

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) GetSecurityGroups() []SecurityGroupModel`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) GetSecurityGroupsOk() (*[]SecurityGroupModel, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) SetSecurityGroups(v []SecurityGroupModel)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroupsNil

`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *BnsLoadBalancerV1ApiListTargetsInTargetGroupModelTargetGroupMemberModel) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


