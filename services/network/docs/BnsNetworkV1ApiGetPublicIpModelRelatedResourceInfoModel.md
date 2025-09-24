# BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 연결된 리소스 ID | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**NullableProvisioningStatus**](ProvisioningStatus.md) |  | [optional] 
**Type** | Pointer to **NullableString** |  | [optional] 
**DeviceOwner** | Pointer to **NullableString** |  | [optional] 
**DeviceId** | Pointer to **NullableString** |  | [optional] 
**SubnetId** | Pointer to **NullableString** |  | [optional] 
**SubnetName** | Pointer to **NullableString** |  | [optional] 
**SubnetCidr** | Pointer to **NullableString** |  | [optional] 
**VpcId** | Pointer to **NullableString** |  | [optional] 
**VpcName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewBnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel

`func NewBnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel(id string, ) *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel`

NewBnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel instantiates a new BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModelWithDefaults

`func NewBnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModelWithDefaults() *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel`

NewBnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModelWithDefaults instantiates a new BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStatus

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) GetStatus() ProvisioningStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) GetStatusOk() (*ProvisioningStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) SetStatus(v ProvisioningStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetType

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetDeviceOwner

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) GetDeviceOwner() string`

GetDeviceOwner returns the DeviceOwner field if non-nil, zero value otherwise.

### GetDeviceOwnerOk

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) GetDeviceOwnerOk() (*string, bool)`

GetDeviceOwnerOk returns a tuple with the DeviceOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceOwner

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) SetDeviceOwner(v string)`

SetDeviceOwner sets DeviceOwner field to given value.

### HasDeviceOwner

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) HasDeviceOwner() bool`

HasDeviceOwner returns a boolean if a field has been set.

### SetDeviceOwnerNil

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) SetDeviceOwnerNil(b bool)`

 SetDeviceOwnerNil sets the value for DeviceOwner to be an explicit nil

### UnsetDeviceOwner
`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) UnsetDeviceOwner()`

UnsetDeviceOwner ensures that no value is present for DeviceOwner, not even an explicit nil
### GetDeviceId

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### SetDeviceIdNil

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) SetDeviceIdNil(b bool)`

 SetDeviceIdNil sets the value for DeviceId to be an explicit nil

### UnsetDeviceId
`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) UnsetDeviceId()`

UnsetDeviceId ensures that no value is present for DeviceId, not even an explicit nil
### GetSubnetId

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### SetSubnetIdNil

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) SetSubnetIdNil(b bool)`

 SetSubnetIdNil sets the value for SubnetId to be an explicit nil

### UnsetSubnetId
`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) UnsetSubnetId()`

UnsetSubnetId ensures that no value is present for SubnetId, not even an explicit nil
### GetSubnetName

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) GetSubnetName() string`

GetSubnetName returns the SubnetName field if non-nil, zero value otherwise.

### GetSubnetNameOk

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) GetSubnetNameOk() (*string, bool)`

GetSubnetNameOk returns a tuple with the SubnetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetName

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) SetSubnetName(v string)`

SetSubnetName sets SubnetName field to given value.

### HasSubnetName

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) HasSubnetName() bool`

HasSubnetName returns a boolean if a field has been set.

### SetSubnetNameNil

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) SetSubnetNameNil(b bool)`

 SetSubnetNameNil sets the value for SubnetName to be an explicit nil

### UnsetSubnetName
`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) UnsetSubnetName()`

UnsetSubnetName ensures that no value is present for SubnetName, not even an explicit nil
### GetSubnetCidr

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) GetSubnetCidr() string`

GetSubnetCidr returns the SubnetCidr field if non-nil, zero value otherwise.

### GetSubnetCidrOk

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) GetSubnetCidrOk() (*string, bool)`

GetSubnetCidrOk returns a tuple with the SubnetCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetCidr

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) SetSubnetCidr(v string)`

SetSubnetCidr sets SubnetCidr field to given value.

### HasSubnetCidr

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) HasSubnetCidr() bool`

HasSubnetCidr returns a boolean if a field has been set.

### SetSubnetCidrNil

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) SetSubnetCidrNil(b bool)`

 SetSubnetCidrNil sets the value for SubnetCidr to be an explicit nil

### UnsetSubnetCidr
`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) UnsetSubnetCidr()`

UnsetSubnetCidr ensures that no value is present for SubnetCidr, not even an explicit nil
### GetVpcId

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### SetVpcIdNil

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) SetVpcIdNil(b bool)`

 SetVpcIdNil sets the value for VpcId to be an explicit nil

### UnsetVpcId
`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) UnsetVpcId()`

UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil
### GetVpcName

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) GetVpcName() string`

GetVpcName returns the VpcName field if non-nil, zero value otherwise.

### GetVpcNameOk

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) GetVpcNameOk() (*string, bool)`

GetVpcNameOk returns a tuple with the VpcName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcName

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) SetVpcName(v string)`

SetVpcName sets VpcName field to given value.

### HasVpcName

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) HasVpcName() bool`

HasVpcName returns a boolean if a field has been set.

### SetVpcNameNil

`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) SetVpcNameNil(b bool)`

 SetVpcNameNil sets the value for VpcName to be an explicit nil

### UnsetVpcName
`func (o *BnsNetworkV1ApiGetPublicIpModelRelatedResourceInfoModel) UnsetVpcName()`

UnsetVpcName ensures that no value is present for VpcName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


