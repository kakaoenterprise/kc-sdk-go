# BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1

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
**Subnet** | [**BnsLoadBalancerV1ApiGetTargetModelSubnetModel**](BnsLoadBalancerV1ApiGetTargetModelSubnetModel.md) | 서브넷 정보 | 
**SecurityGroups** | Pointer to [**[]SecurityGroupModel**](SecurityGroupModel.md) |  | [optional] 

## Methods

### NewBnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1

`func NewBnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1(id string, operatingStatus LoadBalancerOperatingStatus, subnet BnsLoadBalancerV1ApiGetTargetModelSubnetModel, ) *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1`

NewBnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1 instantiates a new BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1WithDefaults

`func NewBnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1WithDefaults() *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1`

NewBnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1WithDefaults instantiates a new BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProjectId

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetTargetGroupId

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) GetTargetGroupId() string`

GetTargetGroupId returns the TargetGroupId field if non-nil, zero value otherwise.

### GetTargetGroupIdOk

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) GetTargetGroupIdOk() (*string, bool)`

GetTargetGroupIdOk returns a tuple with the TargetGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetGroupId

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) SetTargetGroupId(v string)`

SetTargetGroupId sets TargetGroupId field to given value.

### HasTargetGroupId

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) HasTargetGroupId() bool`

HasTargetGroupId returns a boolean if a field has been set.

### SetTargetGroupIdNil

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) SetTargetGroupIdNil(b bool)`

 SetTargetGroupIdNil sets the value for TargetGroupId to be an explicit nil

### UnsetTargetGroupId
`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) UnsetTargetGroupId()`

UnsetTargetGroupId ensures that no value is present for TargetGroupId, not even an explicit nil
### GetIpAddress

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetProtocolPort

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) GetProtocolPort() int32`

GetProtocolPort returns the ProtocolPort field if non-nil, zero value otherwise.

### GetProtocolPortOk

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) GetProtocolPortOk() (*int32, bool)`

GetProtocolPortOk returns a tuple with the ProtocolPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolPort

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) SetProtocolPort(v int32)`

SetProtocolPort sets ProtocolPort field to given value.

### HasProtocolPort

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) HasProtocolPort() bool`

HasProtocolPort returns a boolean if a field has been set.

### SetProtocolPortNil

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) SetProtocolPortNil(b bool)`

 SetProtocolPortNil sets the value for ProtocolPort to be an explicit nil

### UnsetProtocolPort
`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) UnsetProtocolPort()`

UnsetProtocolPort ensures that no value is present for ProtocolPort, not even an explicit nil
### GetWeight

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetOperatingStatus

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) GetOperatingStatus() LoadBalancerOperatingStatus`

GetOperatingStatus returns the OperatingStatus field if non-nil, zero value otherwise.

### GetOperatingStatusOk

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) GetOperatingStatusOk() (*LoadBalancerOperatingStatus, bool)`

GetOperatingStatusOk returns a tuple with the OperatingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingStatus

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) SetOperatingStatus(v LoadBalancerOperatingStatus)`

SetOperatingStatus sets OperatingStatus field to given value.


### GetCreatedAt

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) GetProvisioningStatus() ProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) GetProvisioningStatusOk() (*ProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) SetProvisioningStatus(v ProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetMonitorPort

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) GetMonitorPort() int32`

GetMonitorPort returns the MonitorPort field if non-nil, zero value otherwise.

### GetMonitorPortOk

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) GetMonitorPortOk() (*int32, bool)`

GetMonitorPortOk returns a tuple with the MonitorPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorPort

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) SetMonitorPort(v int32)`

SetMonitorPort sets MonitorPort field to given value.

### HasMonitorPort

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) HasMonitorPort() bool`

HasMonitorPort returns a boolean if a field has been set.

### SetMonitorPortNil

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) SetMonitorPortNil(b bool)`

 SetMonitorPortNil sets the value for MonitorPort to be an explicit nil

### UnsetMonitorPort
`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) UnsetMonitorPort()`

UnsetMonitorPort ensures that no value is present for MonitorPort, not even an explicit nil
### GetNetworkInterfaceId

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) GetNetworkInterfaceId() string`

GetNetworkInterfaceId returns the NetworkInterfaceId field if non-nil, zero value otherwise.

### GetNetworkInterfaceIdOk

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) GetNetworkInterfaceIdOk() (*string, bool)`

GetNetworkInterfaceIdOk returns a tuple with the NetworkInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaceId

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) SetNetworkInterfaceId(v string)`

SetNetworkInterfaceId sets NetworkInterfaceId field to given value.

### HasNetworkInterfaceId

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) HasNetworkInterfaceId() bool`

HasNetworkInterfaceId returns a boolean if a field has been set.

### SetNetworkInterfaceIdNil

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) SetNetworkInterfaceIdNil(b bool)`

 SetNetworkInterfaceIdNil sets the value for NetworkInterfaceId to be an explicit nil

### UnsetNetworkInterfaceId
`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) UnsetNetworkInterfaceId()`

UnsetNetworkInterfaceId ensures that no value is present for NetworkInterfaceId, not even an explicit nil
### GetInstanceId

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### SetInstanceIdNil

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) SetInstanceIdNil(b bool)`

 SetInstanceIdNil sets the value for InstanceId to be an explicit nil

### UnsetInstanceId
`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) UnsetInstanceId()`

UnsetInstanceId ensures that no value is present for InstanceId, not even an explicit nil
### GetInstanceName

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) GetInstanceName() string`

GetInstanceName returns the InstanceName field if non-nil, zero value otherwise.

### GetInstanceNameOk

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) GetInstanceNameOk() (*string, bool)`

GetInstanceNameOk returns a tuple with the InstanceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceName

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) SetInstanceName(v string)`

SetInstanceName sets InstanceName field to given value.

### HasInstanceName

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) HasInstanceName() bool`

HasInstanceName returns a boolean if a field has been set.

### SetInstanceNameNil

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) SetInstanceNameNil(b bool)`

 SetInstanceNameNil sets the value for InstanceName to be an explicit nil

### UnsetInstanceName
`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) UnsetInstanceName()`

UnsetInstanceName ensures that no value is present for InstanceName, not even an explicit nil
### GetVpcId

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### SetVpcIdNil

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) SetVpcIdNil(b bool)`

 SetVpcIdNil sets the value for VpcId to be an explicit nil

### UnsetVpcId
`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) UnsetVpcId()`

UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil
### GetSubnet

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) GetSubnet() BnsLoadBalancerV1ApiGetTargetModelSubnetModel`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) GetSubnetOk() (*BnsLoadBalancerV1ApiGetTargetModelSubnetModel, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) SetSubnet(v BnsLoadBalancerV1ApiGetTargetModelSubnetModel)`

SetSubnet sets Subnet field to given value.


### GetSecurityGroups

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) GetSecurityGroups() []SecurityGroupModel`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) GetSecurityGroupsOk() (*[]SecurityGroupModel, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) SetSecurityGroups(v []SecurityGroupModel)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroupsNil

`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *BnsLoadBalancerV1ApiGetTargetModelTargetGroupMemberModel1) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


