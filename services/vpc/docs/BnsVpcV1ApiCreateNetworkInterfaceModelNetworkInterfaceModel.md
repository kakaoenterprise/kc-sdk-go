# BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 네트워크 인터페이스 ID | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Status** | [**NetworkInterfaceStatus**](NetworkInterfaceStatus.md) | 네트워크 인터페이스의 상태 | 
**Description** | Pointer to **NullableString** |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**IsNetworkInterfaceSecurityEnabled** | Pointer to **NullableBool** |  | [optional] 
**SecurityGroups** | Pointer to **[]string** |  | [optional] 
**MacAddress** | Pointer to **NullableString** |  | [optional] 
**DeviceId** | Pointer to **NullableString** |  | [optional] 
**PrivateIp** | Pointer to **NullableString** |  | [optional] 
**SecondaryIps** | Pointer to **[]string** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewBnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel

`func NewBnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel(id string, status NetworkInterfaceStatus, ) *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel`

NewBnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel instantiates a new BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModelWithDefaults

`func NewBnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModelWithDefaults() *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel`

NewBnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModelWithDefaults instantiates a new BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStatus

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) GetStatus() NetworkInterfaceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) GetStatusOk() (*NetworkInterfaceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) SetStatus(v NetworkInterfaceStatus)`

SetStatus sets Status field to given value.


### GetDescription

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProjectId

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetIsNetworkInterfaceSecurityEnabled

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) GetIsNetworkInterfaceSecurityEnabled() bool`

GetIsNetworkInterfaceSecurityEnabled returns the IsNetworkInterfaceSecurityEnabled field if non-nil, zero value otherwise.

### GetIsNetworkInterfaceSecurityEnabledOk

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) GetIsNetworkInterfaceSecurityEnabledOk() (*bool, bool)`

GetIsNetworkInterfaceSecurityEnabledOk returns a tuple with the IsNetworkInterfaceSecurityEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNetworkInterfaceSecurityEnabled

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) SetIsNetworkInterfaceSecurityEnabled(v bool)`

SetIsNetworkInterfaceSecurityEnabled sets IsNetworkInterfaceSecurityEnabled field to given value.

### HasIsNetworkInterfaceSecurityEnabled

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) HasIsNetworkInterfaceSecurityEnabled() bool`

HasIsNetworkInterfaceSecurityEnabled returns a boolean if a field has been set.

### SetIsNetworkInterfaceSecurityEnabledNil

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) SetIsNetworkInterfaceSecurityEnabledNil(b bool)`

 SetIsNetworkInterfaceSecurityEnabledNil sets the value for IsNetworkInterfaceSecurityEnabled to be an explicit nil

### UnsetIsNetworkInterfaceSecurityEnabled
`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) UnsetIsNetworkInterfaceSecurityEnabled()`

UnsetIsNetworkInterfaceSecurityEnabled ensures that no value is present for IsNetworkInterfaceSecurityEnabled, not even an explicit nil
### GetSecurityGroups

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) GetSecurityGroups() []string`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) GetSecurityGroupsOk() (*[]string, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) SetSecurityGroups(v []string)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroupsNil

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil
### GetMacAddress

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### SetMacAddressNil

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) SetMacAddressNil(b bool)`

 SetMacAddressNil sets the value for MacAddress to be an explicit nil

### UnsetMacAddress
`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) UnsetMacAddress()`

UnsetMacAddress ensures that no value is present for MacAddress, not even an explicit nil
### GetDeviceId

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### SetDeviceIdNil

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) SetDeviceIdNil(b bool)`

 SetDeviceIdNil sets the value for DeviceId to be an explicit nil

### UnsetDeviceId
`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) UnsetDeviceId()`

UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
### GetPrivateIp

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### SetPrivateIpNil

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) SetPrivateIpNil(b bool)`

 SetPrivateIpNil sets the value for PrivateIp to be an explicit nil

### UnsetPrivateIp
`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) UnsetPrivateIp()`

UnsetPrivateIp ensures that no value is present for PrivateIp, not even an explicit nil
### GetSecondaryIps

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) GetSecondaryIps() []string`

GetSecondaryIps returns the SecondaryIps field if non-nil, zero value otherwise.

### GetSecondaryIpsOk

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) GetSecondaryIpsOk() (*[]string, bool)`

GetSecondaryIpsOk returns a tuple with the SecondaryIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryIps

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) SetSecondaryIps(v []string)`

SetSecondaryIps sets SecondaryIps field to given value.

### HasSecondaryIps

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) HasSecondaryIps() bool`

HasSecondaryIps returns a boolean if a field has been set.

### SetSecondaryIpsNil

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) SetSecondaryIpsNil(b bool)`

 SetSecondaryIpsNil sets the value for SecondaryIps to be an explicit nil

### UnsetSecondaryIps
`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) UnsetSecondaryIps()`

UnsetSecondaryIps ensures that no value is present for SecondaryIps, not even an explicit nil
### GetCreatedAt

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BnsVpcV1ApiCreateNetworkInterfaceModelNetworkInterfaceModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


