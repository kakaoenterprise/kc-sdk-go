# BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 네트워크 인터페이스의 ID | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**NullableNetworkInterfaceStatus**](NetworkInterfaceStatus.md) |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**VpcId** | Pointer to **NullableString** |  | [optional] 
**SubnetId** | Pointer to **NullableString** |  | [optional] 
**MacAddress** | Pointer to **NullableString** |  | [optional] 
**DeviceId** | Pointer to **NullableString** |  | [optional] 
**DeviceOwner** | Pointer to **NullableString** |  | [optional] 
**ProjectName** | Pointer to **NullableString** |  | [optional] 
**SecondaryIps** | Pointer to **[]string** |  | [optional] 
**PublicIp** | Pointer to **NullableString** |  | [optional] 
**PrivateIp** | Pointer to **NullableString** |  | [optional] 
**IsNetworkInterfaceSecurityEnabled** | Pointer to **NullableBool** |  | [optional] 
**AllowedAddressPairs** | Pointer to [**[]BnsVpcV1ApiGetNetworkInterfaceModelAllowedAddressPairModel**](BnsVpcV1ApiGetNetworkInterfaceModelAllowedAddressPairModel.md) |  | [optional] 
**SecurityGroups** | Pointer to [**[]SecurityGroupModel**](SecurityGroupModel.md) |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewBnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel

`func NewBnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel(id string, ) *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel`

NewBnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel instantiates a new BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModelWithDefaults

`func NewBnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModelWithDefaults() *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel`

NewBnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModelWithDefaults instantiates a new BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStatus

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) GetStatus() NetworkInterfaceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) GetStatusOk() (*NetworkInterfaceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) SetStatus(v NetworkInterfaceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetDescription

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProjectId

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetVpcId

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### SetVpcIdNil

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) SetVpcIdNil(b bool)`

 SetVpcIdNil sets the value for VpcId to be an explicit nil

### UnsetVpcId
`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) UnsetVpcId()`

UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil
### GetSubnetId

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### SetSubnetIdNil

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) SetSubnetIdNil(b bool)`

 SetSubnetIdNil sets the value for SubnetId to be an explicit nil

### UnsetSubnetId
`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) UnsetSubnetId()`

UnsetSubnetId ensures that no value is present for SubnetId, not even an explicit nil
### GetMacAddress

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### SetMacAddressNil

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) SetMacAddressNil(b bool)`

 SetMacAddressNil sets the value for MacAddress to be an explicit nil

### UnsetMacAddress
`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) UnsetMacAddress()`

UnsetMacAddress ensures that no value is present for MacAddress, not even an explicit nil
### GetDeviceId

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### SetDeviceIdNil

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) SetDeviceIdNil(b bool)`

 SetDeviceIdNil sets the value for DeviceId to be an explicit nil

### UnsetDeviceId
`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) UnsetDeviceId()`

UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
### GetDeviceOwner

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) GetDeviceOwner() string`

GetDeviceOwner returns the DeviceOwner field if non-nil, zero value otherwise.

### GetDeviceOwnerOk

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) GetDeviceOwnerOk() (*string, bool)`

GetDeviceOwnerOk returns a tuple with the DeviceOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceOwner

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) SetDeviceOwner(v string)`

SetDeviceOwner sets DeviceOwner field to given value.

### HasDeviceOwner

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) HasDeviceOwner() bool`

HasDeviceOwner returns a boolean if a field has been set.

### SetDeviceOwnerNil

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) SetDeviceOwnerNil(b bool)`

 SetDeviceOwnerNil sets the value for DeviceOwner to be an explicit nil

### UnsetDeviceOwner
`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) UnsetDeviceOwner()`

UnsetDeviceOwner ensures that no value is present for DeviceOwner, not even an explicit nil
### GetProjectName

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### SetProjectNameNil

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) SetProjectNameNil(b bool)`

 SetProjectNameNil sets the value for ProjectName to be an explicit nil

### UnsetProjectName
`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) UnsetProjectName()`

UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
### GetSecondaryIps

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) GetSecondaryIps() []string`

GetSecondaryIps returns the SecondaryIps field if non-nil, zero value otherwise.

### GetSecondaryIpsOk

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) GetSecondaryIpsOk() (*[]string, bool)`

GetSecondaryIpsOk returns a tuple with the SecondaryIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryIps

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) SetSecondaryIps(v []string)`

SetSecondaryIps sets SecondaryIps field to given value.

### HasSecondaryIps

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) HasSecondaryIps() bool`

HasSecondaryIps returns a boolean if a field has been set.

### SetSecondaryIpsNil

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) SetSecondaryIpsNil(b bool)`

 SetSecondaryIpsNil sets the value for SecondaryIps to be an explicit nil

### UnsetSecondaryIps
`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) UnsetSecondaryIps()`

UnsetSecondaryIps ensures that no value is present for SecondaryIps, not even an explicit nil
### GetPublicIp

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### SetPublicIpNil

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) SetPublicIpNil(b bool)`

 SetPublicIpNil sets the value for PublicIp to be an explicit nil

### UnsetPublicIp
`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) UnsetPublicIp()`

UnsetPublicIp ensures that no value is present for PublicIp, not even an explicit nil
### GetPrivateIp

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### SetPrivateIpNil

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) SetPrivateIpNil(b bool)`

 SetPrivateIpNil sets the value for PrivateIp to be an explicit nil

### UnsetPrivateIp
`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) UnsetPrivateIp()`

UnsetPrivateIp ensures that no value is present for PrivateIp, not even an explicit nil
### GetIsNetworkInterfaceSecurityEnabled

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) GetIsNetworkInterfaceSecurityEnabled() bool`

GetIsNetworkInterfaceSecurityEnabled returns the IsNetworkInterfaceSecurityEnabled field if non-nil, zero value otherwise.

### GetIsNetworkInterfaceSecurityEnabledOk

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) GetIsNetworkInterfaceSecurityEnabledOk() (*bool, bool)`

GetIsNetworkInterfaceSecurityEnabledOk returns a tuple with the IsNetworkInterfaceSecurityEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNetworkInterfaceSecurityEnabled

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) SetIsNetworkInterfaceSecurityEnabled(v bool)`

SetIsNetworkInterfaceSecurityEnabled sets IsNetworkInterfaceSecurityEnabled field to given value.

### HasIsNetworkInterfaceSecurityEnabled

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) HasIsNetworkInterfaceSecurityEnabled() bool`

HasIsNetworkInterfaceSecurityEnabled returns a boolean if a field has been set.

### SetIsNetworkInterfaceSecurityEnabledNil

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) SetIsNetworkInterfaceSecurityEnabledNil(b bool)`

 SetIsNetworkInterfaceSecurityEnabledNil sets the value for IsNetworkInterfaceSecurityEnabled to be an explicit nil

### UnsetIsNetworkInterfaceSecurityEnabled
`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) UnsetIsNetworkInterfaceSecurityEnabled()`

UnsetIsNetworkInterfaceSecurityEnabled ensures that no value is present for IsNetworkInterfaceSecurityEnabled, not even an explicit nil
### GetAllowedAddressPairs

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) GetAllowedAddressPairs() []BnsVpcV1ApiGetNetworkInterfaceModelAllowedAddressPairModel`

GetAllowedAddressPairs returns the AllowedAddressPairs field if non-nil, zero value otherwise.

### GetAllowedAddressPairsOk

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) GetAllowedAddressPairsOk() (*[]BnsVpcV1ApiGetNetworkInterfaceModelAllowedAddressPairModel, bool)`

GetAllowedAddressPairsOk returns a tuple with the AllowedAddressPairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAddressPairs

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) SetAllowedAddressPairs(v []BnsVpcV1ApiGetNetworkInterfaceModelAllowedAddressPairModel)`

SetAllowedAddressPairs sets AllowedAddressPairs field to given value.

### HasAllowedAddressPairs

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) HasAllowedAddressPairs() bool`

HasAllowedAddressPairs returns a boolean if a field has been set.

### SetAllowedAddressPairsNil

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) SetAllowedAddressPairsNil(b bool)`

 SetAllowedAddressPairsNil sets the value for AllowedAddressPairs to be an explicit nil

### UnsetAllowedAddressPairs
`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) UnsetAllowedAddressPairs()`

UnsetAllowedAddressPairs ensures that no value is present for AllowedAddressPairs, not even an explicit nil
### GetSecurityGroups

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) GetSecurityGroups() []SecurityGroupModel`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) GetSecurityGroupsOk() (*[]SecurityGroupModel, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) SetSecurityGroups(v []SecurityGroupModel)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroupsNil

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil
### GetCreatedAt

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BnsVpcV1ApiGetNetworkInterfaceModelNetworkInterfaceModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


