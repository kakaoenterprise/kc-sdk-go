# BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel

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
**AllowedAddressPairs** | Pointer to [**[]AllowedAddressPairModelOutput**](AllowedAddressPairModelOutput.md) |  | [optional] 
**SecurityGroups** | Pointer to [**[]SecurityGroupModel**](SecurityGroupModel.md) |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewBcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel

`func NewBcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel(id string, ) *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel`

NewBcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel instantiates a new BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModelWithDefaults

`func NewBcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModelWithDefaults() *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel`

NewBcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModelWithDefaults instantiates a new BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStatus

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) GetStatus() NetworkInterfaceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) GetStatusOk() (*NetworkInterfaceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) SetStatus(v NetworkInterfaceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetDescription

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProjectId

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetVpcId

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### SetVpcIdNil

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) SetVpcIdNil(b bool)`

 SetVpcIdNil sets the value for VpcId to be an explicit nil

### UnsetVpcId
`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) UnsetVpcId()`

UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil
### GetSubnetId

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### SetSubnetIdNil

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) SetSubnetIdNil(b bool)`

 SetSubnetIdNil sets the value for SubnetId to be an explicit nil

### UnsetSubnetId
`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) UnsetSubnetId()`

UnsetSubnetId ensures that no value is present for SubnetId, not even an explicit nil
### GetMacAddress

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### SetMacAddressNil

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) SetMacAddressNil(b bool)`

 SetMacAddressNil sets the value for MacAddress to be an explicit nil

### UnsetMacAddress
`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) UnsetMacAddress()`

UnsetMacAddress ensures that no value is present for MacAddress, not even an explicit nil
### GetDeviceId

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### SetDeviceIdNil

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) SetDeviceIdNil(b bool)`

 SetDeviceIdNil sets the value for DeviceId to be an explicit nil

### UnsetDeviceId
`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) UnsetDeviceId()`

UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
### GetDeviceOwner

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) GetDeviceOwner() string`

GetDeviceOwner returns the DeviceOwner field if non-nil, zero value otherwise.

### GetDeviceOwnerOk

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) GetDeviceOwnerOk() (*string, bool)`

GetDeviceOwnerOk returns a tuple with the DeviceOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceOwner

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) SetDeviceOwner(v string)`

SetDeviceOwner sets DeviceOwner field to given value.

### HasDeviceOwner

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) HasDeviceOwner() bool`

HasDeviceOwner returns a boolean if a field has been set.

### SetDeviceOwnerNil

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) SetDeviceOwnerNil(b bool)`

 SetDeviceOwnerNil sets the value for DeviceOwner to be an explicit nil

### UnsetDeviceOwner
`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) UnsetDeviceOwner()`

UnsetDeviceOwner ensures that no value is present for DeviceOwner, not even an explicit nil
### GetProjectName

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) GetProjectName() string`

GetProjectName returns the ProjectName field if non-nil, zero value otherwise.

### GetProjectNameOk

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) GetProjectNameOk() (*string, bool)`

GetProjectNameOk returns a tuple with the ProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectName

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) SetProjectName(v string)`

SetProjectName sets ProjectName field to given value.

### HasProjectName

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) HasProjectName() bool`

HasProjectName returns a boolean if a field has been set.

### SetProjectNameNil

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) SetProjectNameNil(b bool)`

 SetProjectNameNil sets the value for ProjectName to be an explicit nil

### UnsetProjectName
`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) UnsetProjectName()`

UnsetProjectName ensures that no value is present for ProjectName, not even an explicit nil
### GetSecondaryIps

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) GetSecondaryIps() []string`

GetSecondaryIps returns the SecondaryIps field if non-nil, zero value otherwise.

### GetSecondaryIpsOk

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) GetSecondaryIpsOk() (*[]string, bool)`

GetSecondaryIpsOk returns a tuple with the SecondaryIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryIps

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) SetSecondaryIps(v []string)`

SetSecondaryIps sets SecondaryIps field to given value.

### HasSecondaryIps

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) HasSecondaryIps() bool`

HasSecondaryIps returns a boolean if a field has been set.

### SetSecondaryIpsNil

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) SetSecondaryIpsNil(b bool)`

 SetSecondaryIpsNil sets the value for SecondaryIps to be an explicit nil

### UnsetSecondaryIps
`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) UnsetSecondaryIps()`

UnsetSecondaryIps ensures that no value is present for SecondaryIps, not even an explicit nil
### GetPublicIp

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### SetPublicIpNil

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) SetPublicIpNil(b bool)`

 SetPublicIpNil sets the value for PublicIp to be an explicit nil

### UnsetPublicIp
`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) UnsetPublicIp()`

UnsetPublicIp ensures that no value is present for PublicIp, not even an explicit nil
### GetPrivateIp

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### SetPrivateIpNil

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) SetPrivateIpNil(b bool)`

 SetPrivateIpNil sets the value for PrivateIp to be an explicit nil

### UnsetPrivateIp
`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) UnsetPrivateIp()`

UnsetPrivateIp ensures that no value is present for PrivateIp, not even an explicit nil
### GetIsNetworkInterfaceSecurityEnabled

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) GetIsNetworkInterfaceSecurityEnabled() bool`

GetIsNetworkInterfaceSecurityEnabled returns the IsNetworkInterfaceSecurityEnabled field if non-nil, zero value otherwise.

### GetIsNetworkInterfaceSecurityEnabledOk

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) GetIsNetworkInterfaceSecurityEnabledOk() (*bool, bool)`

GetIsNetworkInterfaceSecurityEnabledOk returns a tuple with the IsNetworkInterfaceSecurityEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNetworkInterfaceSecurityEnabled

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) SetIsNetworkInterfaceSecurityEnabled(v bool)`

SetIsNetworkInterfaceSecurityEnabled sets IsNetworkInterfaceSecurityEnabled field to given value.

### HasIsNetworkInterfaceSecurityEnabled

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) HasIsNetworkInterfaceSecurityEnabled() bool`

HasIsNetworkInterfaceSecurityEnabled returns a boolean if a field has been set.

### SetIsNetworkInterfaceSecurityEnabledNil

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) SetIsNetworkInterfaceSecurityEnabledNil(b bool)`

 SetIsNetworkInterfaceSecurityEnabledNil sets the value for IsNetworkInterfaceSecurityEnabled to be an explicit nil

### UnsetIsNetworkInterfaceSecurityEnabled
`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) UnsetIsNetworkInterfaceSecurityEnabled()`

UnsetIsNetworkInterfaceSecurityEnabled ensures that no value is present for IsNetworkInterfaceSecurityEnabled, not even an explicit nil
### GetAllowedAddressPairs

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) GetAllowedAddressPairs() []AllowedAddressPairModelOutput`

GetAllowedAddressPairs returns the AllowedAddressPairs field if non-nil, zero value otherwise.

### GetAllowedAddressPairsOk

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) GetAllowedAddressPairsOk() (*[]AllowedAddressPairModelOutput, bool)`

GetAllowedAddressPairsOk returns a tuple with the AllowedAddressPairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAddressPairs

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) SetAllowedAddressPairs(v []AllowedAddressPairModelOutput)`

SetAllowedAddressPairs sets AllowedAddressPairs field to given value.

### HasAllowedAddressPairs

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) HasAllowedAddressPairs() bool`

HasAllowedAddressPairs returns a boolean if a field has been set.

### SetAllowedAddressPairsNil

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) SetAllowedAddressPairsNil(b bool)`

 SetAllowedAddressPairsNil sets the value for AllowedAddressPairs to be an explicit nil

### UnsetAllowedAddressPairs
`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) UnsetAllowedAddressPairs()`

UnsetAllowedAddressPairs ensures that no value is present for AllowedAddressPairs, not even an explicit nil
### GetSecurityGroups

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) GetSecurityGroups() []SecurityGroupModel`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) GetSecurityGroupsOk() (*[]SecurityGroupModel, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) SetSecurityGroups(v []SecurityGroupModel)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroupsNil

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil
### GetCreatedAt

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BcsInstanceV1ApiListInstanceNetworkInterfacesModelInstanceNetworkInterfaceModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


