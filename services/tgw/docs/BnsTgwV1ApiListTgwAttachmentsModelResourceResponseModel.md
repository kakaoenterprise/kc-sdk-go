# BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**AvailabilityZone** | Pointer to [**NullableAvailabilityZone**](AvailabilityZone.md) |  | [optional] 
**CidrBlock** | Pointer to **NullableString** |  | [optional] 
**OperatingStatus** | Pointer to **NullableString** |  | [optional] 
**ProvisioningStatus** | Pointer to [**NullableTGWProvisioningStatus**](TGWProvisioningStatus.md) |  | [optional] 
**VpcId** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewBnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel

`func NewBnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel() *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel`

NewBnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel instantiates a new BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsTgwV1ApiListTgwAttachmentsModelResourceResponseModelWithDefaults

`func NewBnsTgwV1ApiListTgwAttachmentsModelResourceResponseModelWithDefaults() *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel`

NewBnsTgwV1ApiListTgwAttachmentsModelResourceResponseModelWithDefaults instantiates a new BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetAvailabilityZone

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### SetAvailabilityZoneNil

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
### GetCidrBlock

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.

### HasCidrBlock

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) HasCidrBlock() bool`

HasCidrBlock returns a boolean if a field has been set.

### SetCidrBlockNil

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) SetCidrBlockNil(b bool)`

 SetCidrBlockNil sets the value for CidrBlock to be an explicit nil

### UnsetCidrBlock
`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) UnsetCidrBlock()`

UnsetCidrBlock ensures that no value is present for CidrBlock, not even an explicit nil
### GetOperatingStatus

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) GetOperatingStatus() string`

GetOperatingStatus returns the OperatingStatus field if non-nil, zero value otherwise.

### GetOperatingStatusOk

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) GetOperatingStatusOk() (*string, bool)`

GetOperatingStatusOk returns a tuple with the OperatingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingStatus

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) SetOperatingStatus(v string)`

SetOperatingStatus sets OperatingStatus field to given value.

### HasOperatingStatus

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) HasOperatingStatus() bool`

HasOperatingStatus returns a boolean if a field has been set.

### SetOperatingStatusNil

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) SetOperatingStatusNil(b bool)`

 SetOperatingStatusNil sets the value for OperatingStatus to be an explicit nil

### UnsetOperatingStatus
`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) UnsetOperatingStatus()`

UnsetOperatingStatus ensures that no value is present for OperatingStatus, not even an explicit nil
### GetProvisioningStatus

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) GetProvisioningStatus() TGWProvisioningStatus`

GetProvisioningStatus returns the ProvisioningStatus field if non-nil, zero value otherwise.

### GetProvisioningStatusOk

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) GetProvisioningStatusOk() (*TGWProvisioningStatus, bool)`

GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatus

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) SetProvisioningStatus(v TGWProvisioningStatus)`

SetProvisioningStatus sets ProvisioningStatus field to given value.

### HasProvisioningStatus

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) HasProvisioningStatus() bool`

HasProvisioningStatus returns a boolean if a field has been set.

### SetProvisioningStatusNil

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) SetProvisioningStatusNil(b bool)`

 SetProvisioningStatusNil sets the value for ProvisioningStatus to be an explicit nil

### UnsetProvisioningStatus
`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) UnsetProvisioningStatus()`

UnsetProvisioningStatus ensures that no value is present for ProvisioningStatus, not even an explicit nil
### GetVpcId

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) GetVpcId() string`

GetVpcId returns the VpcId field if non-nil, zero value otherwise.

### GetVpcIdOk

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) GetVpcIdOk() (*string, bool)`

GetVpcIdOk returns a tuple with the VpcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcId

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) SetVpcId(v string)`

SetVpcId sets VpcId field to given value.

### HasVpcId

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) HasVpcId() bool`

HasVpcId returns a boolean if a field has been set.

### SetVpcIdNil

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) SetVpcIdNil(b bool)`

 SetVpcIdNil sets the value for VpcId to be an explicit nil

### UnsetVpcId
`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) UnsetVpcId()`

UnsetVpcId ensures that no value is present for VpcId, not even an explicit nil
### GetCreatedAt

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BnsTgwV1ApiListTgwAttachmentsModelResourceResponseModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


