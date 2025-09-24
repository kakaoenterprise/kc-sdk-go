# BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 네트워크 인터페이스 ID | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Status** | [**NetworkInterfaceStatus**](NetworkInterfaceStatus.md) | 네트워크 인터페이스 상태 | 
**Description** | Pointer to **NullableString** |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**IsNetworkInterfaceSecurityEnabled** | Pointer to **NullableBool** |  | [optional] 
**SecurityGroups** | Pointer to **[]string** |  | [optional] 
**MacAddress** | Pointer to **NullableString** |  | [optional] 
**DeviceId** | Pointer to **NullableString** |  | [optional] 
**SecondaryIps** | Pointer to **[]string** |  | [optional] 
**PrivateIp** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewBcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel

`func NewBcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel(id string, status NetworkInterfaceStatus, ) *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel`

NewBcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel instantiates a new BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModelWithDefaults

`func NewBcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModelWithDefaults() *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel`

NewBcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModelWithDefaults instantiates a new BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStatus

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) GetStatus() NetworkInterfaceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) GetStatusOk() (*NetworkInterfaceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) SetStatus(v NetworkInterfaceStatus)`

SetStatus sets Status field to given value.


### GetDescription

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProjectId

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetIsNetworkInterfaceSecurityEnabled

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) GetIsNetworkInterfaceSecurityEnabled() bool`

GetIsNetworkInterfaceSecurityEnabled returns the IsNetworkInterfaceSecurityEnabled field if non-nil, zero value otherwise.

### GetIsNetworkInterfaceSecurityEnabledOk

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) GetIsNetworkInterfaceSecurityEnabledOk() (*bool, bool)`

GetIsNetworkInterfaceSecurityEnabledOk returns a tuple with the IsNetworkInterfaceSecurityEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNetworkInterfaceSecurityEnabled

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) SetIsNetworkInterfaceSecurityEnabled(v bool)`

SetIsNetworkInterfaceSecurityEnabled sets IsNetworkInterfaceSecurityEnabled field to given value.

### HasIsNetworkInterfaceSecurityEnabled

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) HasIsNetworkInterfaceSecurityEnabled() bool`

HasIsNetworkInterfaceSecurityEnabled returns a boolean if a field has been set.

### SetIsNetworkInterfaceSecurityEnabledNil

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) SetIsNetworkInterfaceSecurityEnabledNil(b bool)`

 SetIsNetworkInterfaceSecurityEnabledNil sets the value for IsNetworkInterfaceSecurityEnabled to be an explicit nil

### UnsetIsNetworkInterfaceSecurityEnabled
`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) UnsetIsNetworkInterfaceSecurityEnabled()`

UnsetIsNetworkInterfaceSecurityEnabled ensures that no value is present for IsNetworkInterfaceSecurityEnabled, not even an explicit nil
### GetSecurityGroups

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) GetSecurityGroups() []string`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) GetSecurityGroupsOk() (*[]string, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) SetSecurityGroups(v []string)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroupsNil

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil
### GetMacAddress

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### SetMacAddressNil

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) SetMacAddressNil(b bool)`

 SetMacAddressNil sets the value for MacAddress to be an explicit nil

### UnsetMacAddress
`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) UnsetMacAddress()`

UnsetMacAddress ensures that no value is present for MacAddress, not even an explicit nil
### GetDeviceId

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### SetDeviceIdNil

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) SetDeviceIdNil(b bool)`

 SetDeviceIdNil sets the value for DeviceId to be an explicit nil

### UnsetDeviceId
`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) UnsetDeviceId()`

UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
### GetSecondaryIps

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) GetSecondaryIps() []string`

GetSecondaryIps returns the SecondaryIps field if non-nil, zero value otherwise.

### GetSecondaryIpsOk

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) GetSecondaryIpsOk() (*[]string, bool)`

GetSecondaryIpsOk returns a tuple with the SecondaryIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryIps

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) SetSecondaryIps(v []string)`

SetSecondaryIps sets SecondaryIps field to given value.

### HasSecondaryIps

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) HasSecondaryIps() bool`

HasSecondaryIps returns a boolean if a field has been set.

### SetSecondaryIpsNil

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) SetSecondaryIpsNil(b bool)`

 SetSecondaryIpsNil sets the value for SecondaryIps to be an explicit nil

### UnsetSecondaryIps
`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) UnsetSecondaryIps()`

UnsetSecondaryIps ensures that no value is present for SecondaryIps, not even an explicit nil
### GetPrivateIp

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### SetPrivateIpNil

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) SetPrivateIpNil(b bool)`

 SetPrivateIpNil sets the value for PrivateIp to be an explicit nil

### UnsetPrivateIp
`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) UnsetPrivateIp()`

UnsetPrivateIp ensures that no value is present for PrivateIp, not even an explicit nil
### GetCreatedAt

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BcsInstanceV1ApiUpdateNetworkInterfaceAllowedAddressesModelInstanceNetworkInterfaceModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


