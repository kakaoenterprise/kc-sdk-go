# BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 네트워크 인터페이스 ID | 
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
**AllowedAddressPairs** | Pointer to [**[]BnsVpcV1ApiListNetworkInterfacesModelAllowedAddressPairModel**](BnsVpcV1ApiListNetworkInterfacesModelAllowedAddressPairModel.md) |  | [optional] 
**SecurityGroups** | Pointer to [**[]SecurityGroupModel**](SecurityGroupModel.md) |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewBnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel

`func NewBnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel(id string, ) *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel`

NewBnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel instantiates a new BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModelWithDefaults

`func NewBnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModelWithDefaults() *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel`

NewBnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModelWithDefaults instantiates a new BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStatus

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) GetStatus() NetworkInterfaceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) GetStatusOk() (*NetworkInterfaceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) SetStatus(v NetworkInterfaceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetDescription

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProjectId

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetVpcId

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### SetVpcIdNil

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) SetVpcIdNil(b bool)`

 SetVpcIdNil sets the value for VpcId to be an explicit nil

### UnsetVpcId
`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) UnsetVpcId()`

UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil
### GetSubnetId

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### SetSubnetIdNil

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) SetSubnetIdNil(b bool)`

 SetSubnetIdNil sets the value for SubnetId to be an explicit nil

### UnsetSubnetId
`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) UnsetSubnetId()`

UnsetSubnetId ensures that no value is present for SubnetId, not even an explicit nil
### GetMacAddress

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### SetMacAddressNil

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) SetMacAddressNil(b bool)`

 SetMacAddressNil sets the value for MacAddress to be an explicit nil

### UnsetMacAddress
`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) UnsetMacAddress()`

UnsetMacAddress ensures that no value is present for MacAddress, not even an explicit nil
### GetDeviceId

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### SetDeviceIdNil

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) SetDeviceIdNil(b bool)`

 SetDeviceIdNil sets the value for DeviceId to be an explicit nil

### UnsetDeviceId
`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) UnsetDeviceId()`

UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
### GetDeviceOwner

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) GetDeviceOwner() string`

GetDeviceOwner returns the DeviceOwner field if non-nil, zero value otherwise.

### GetDeviceOwnerOk

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) GetDeviceOwnerOk() (*string, bool)`

GetDeviceOwnerOk returns a tuple with the DeviceOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceOwner

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) SetDeviceOwner(v string)`

SetDeviceOwner sets DeviceOwner field to given value.

### HasDeviceOwner

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) HasDeviceOwner() bool`

HasDeviceOwner returns a boolean if a field has been set.

### SetDeviceOwnerNil

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) SetDeviceOwnerNil(b bool)`

 SetDeviceOwnerNil sets the value for DeviceOwner to be an explicit nil

### UnsetDeviceOwner
`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) UnsetDeviceOwner()`

UnsetDeviceOwner ensures that no value is present for DeviceOwner, not even an explicit nil
### GetProjectName

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### SetProjectNameNil

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) SetProjectNameNil(b bool)`

 SetProjectNameNil sets the value for ProjectName to be an explicit nil

### UnsetProjectName
`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) UnsetProjectName()`

UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
### GetSecondaryIps

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) GetSecondaryIps() []string`

GetSecondaryIps returns the SecondaryIps field if non-nil, zero value otherwise.

### GetSecondaryIpsOk

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) GetSecondaryIpsOk() (*[]string, bool)`

GetSecondaryIpsOk returns a tuple with the SecondaryIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryIps

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) SetSecondaryIps(v []string)`

SetSecondaryIps sets SecondaryIps field to given value.

### HasSecondaryIps

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) HasSecondaryIps() bool`

HasSecondaryIps returns a boolean if a field has been set.

### SetSecondaryIpsNil

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) SetSecondaryIpsNil(b bool)`

 SetSecondaryIpsNil sets the value for SecondaryIps to be an explicit nil

### UnsetSecondaryIps
`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) UnsetSecondaryIps()`

UnsetSecondaryIps ensures that no value is present for SecondaryIps, not even an explicit nil
### GetPublicIp

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### SetPublicIpNil

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) SetPublicIpNil(b bool)`

 SetPublicIpNil sets the value for PublicIp to be an explicit nil

### UnsetPublicIp
`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) UnsetPublicIp()`

UnsetPublicIp ensures that no value is present for PublicIp, not even an explicit nil
### GetPrivateIp

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### SetPrivateIpNil

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) SetPrivateIpNil(b bool)`

 SetPrivateIpNil sets the value for PrivateIp to be an explicit nil

### UnsetPrivateIp
`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) UnsetPrivateIp()`

UnsetPrivateIp ensures that no value is present for PrivateIp, not even an explicit nil
### GetIsNetworkInterfaceSecurityEnabled

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) GetIsNetworkInterfaceSecurityEnabled() bool`

GetIsNetworkInterfaceSecurityEnabled returns the IsNetworkInterfaceSecurityEnabled field if non-nil, zero value otherwise.

### GetIsNetworkInterfaceSecurityEnabledOk

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) GetIsNetworkInterfaceSecurityEnabledOk() (*bool, bool)`

GetIsNetworkInterfaceSecurityEnabledOk returns a tuple with the IsNetworkInterfaceSecurityEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNetworkInterfaceSecurityEnabled

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) SetIsNetworkInterfaceSecurityEnabled(v bool)`

SetIsNetworkInterfaceSecurityEnabled sets IsNetworkInterfaceSecurityEnabled field to given value.

### HasIsNetworkInterfaceSecurityEnabled

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) HasIsNetworkInterfaceSecurityEnabled() bool`

HasIsNetworkInterfaceSecurityEnabled returns a boolean if a field has been set.

### SetIsNetworkInterfaceSecurityEnabledNil

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) SetIsNetworkInterfaceSecurityEnabledNil(b bool)`

 SetIsNetworkInterfaceSecurityEnabledNil sets the value for IsNetworkInterfaceSecurityEnabled to be an explicit nil

### UnsetIsNetworkInterfaceSecurityEnabled
`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) UnsetIsNetworkInterfaceSecurityEnabled()`

UnsetIsNetworkInterfaceSecurityEnabled ensures that no value is present for IsNetworkInterfaceSecurityEnabled, not even an explicit nil
### GetAllowedAddressPairs

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) GetAllowedAddressPairs() []BnsVpcV1ApiListNetworkInterfacesModelAllowedAddressPairModel`

GetAllowedAddressPairs returns the AllowedAddressPairs field if non-nil, zero value otherwise.

### GetAllowedAddressPairsOk

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) GetAllowedAddressPairsOk() (*[]BnsVpcV1ApiListNetworkInterfacesModelAllowedAddressPairModel, bool)`

GetAllowedAddressPairsOk returns a tuple with the AllowedAddressPairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAddressPairs

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) SetAllowedAddressPairs(v []BnsVpcV1ApiListNetworkInterfacesModelAllowedAddressPairModel)`

SetAllowedAddressPairs sets AllowedAddressPairs field to given value.

### HasAllowedAddressPairs

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) HasAllowedAddressPairs() bool`

HasAllowedAddressPairs returns a boolean if a field has been set.

### SetAllowedAddressPairsNil

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) SetAllowedAddressPairsNil(b bool)`

 SetAllowedAddressPairsNil sets the value for AllowedAddressPairs to be an explicit nil

### UnsetAllowedAddressPairs
`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) UnsetAllowedAddressPairs()`

UnsetAllowedAddressPairs ensures that no value is present for AllowedAddressPairs, not even an explicit nil
### GetSecurityGroups

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) GetSecurityGroups() []SecurityGroupModel`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) GetSecurityGroupsOk() (*[]SecurityGroupModel, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) SetSecurityGroups(v []SecurityGroupModel)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroupsNil

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil
### GetCreatedAt

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BnsVpcV1ApiListNetworkInterfacesModelNetworkInterfaceModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


