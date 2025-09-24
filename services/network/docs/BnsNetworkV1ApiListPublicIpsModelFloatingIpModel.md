# BnsNetworkV1ApiListPublicIpsModelFloatingIpModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 퍼블릭 IP 자원의 ID | 
**Status** | Pointer to [**NullablePublicIpStatus**](PublicIpStatus.md) |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ProjectId** | Pointer to **NullableString** |  | [optional] 
**PublicIp** | Pointer to **NullableString** |  | [optional] 
**PrivateIp** | Pointer to **NullableString** |  | [optional] 
**RelatedResource** | Pointer to [**NullableBnsNetworkV1ApiListPublicIpsModelRelatedResourceInfoModel**](BnsNetworkV1ApiListPublicIpsModelRelatedResourceInfoModel.md) |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewBnsNetworkV1ApiListPublicIpsModelFloatingIpModel

`func NewBnsNetworkV1ApiListPublicIpsModelFloatingIpModel(id string, ) *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel`

NewBnsNetworkV1ApiListPublicIpsModelFloatingIpModel instantiates a new BnsNetworkV1ApiListPublicIpsModelFloatingIpModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBnsNetworkV1ApiListPublicIpsModelFloatingIpModelWithDefaults

`func NewBnsNetworkV1ApiListPublicIpsModelFloatingIpModelWithDefaults() *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel`

NewBnsNetworkV1ApiListPublicIpsModelFloatingIpModelWithDefaults instantiates a new BnsNetworkV1ApiListPublicIpsModelFloatingIpModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) GetStatus() PublicIpStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) GetStatusOk() (*PublicIpStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) SetStatus(v PublicIpStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetDescription

`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetProjectId

`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### SetProjectIdNil

`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) SetProjectIdNil(b bool)`

 SetProjectIdNil sets the value for ProjectId to be an explicit nil

### UnsetProjectId
`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) UnsetProjectId()`

UnsetProjectId ensures that no value is present for ProjectId, not even an explicit nil
### GetPublicIp

`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### SetPublicIpNil

`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) SetPublicIpNil(b bool)`

 SetPublicIpNil sets the value for PublicIp to be an explicit nil

### UnsetPublicIp
`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) UnsetPublicIp()`

UnsetPublicIp ensures that no value is present for PublicIp, not even an explicit nil
### GetPrivateIp

`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) GetPrivateIp() string`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) GetPrivateIpOk() (*string, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) SetPrivateIp(v string)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### SetPrivateIpNil

`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) SetPrivateIpNil(b bool)`

 SetPrivateIpNil sets the value for PrivateIp to be an explicit nil

### UnsetPrivateIp
`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) UnsetPrivateIp()`

UnsetPrivateIp ensures that no value is present for PrivateIp, not even an explicit nil
### GetRelatedResource

`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) GetRelatedResource() BnsNetworkV1ApiListPublicIpsModelRelatedResourceInfoModel`

GetRelatedResource returns the RelatedResource field if non-nil, zero value otherwise.

### GetRelatedResourceOk

`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) GetRelatedResourceOk() (*BnsNetworkV1ApiListPublicIpsModelRelatedResourceInfoModel, bool)`

GetRelatedResourceOk returns a tuple with the RelatedResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelatedResource

`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) SetRelatedResource(v BnsNetworkV1ApiListPublicIpsModelRelatedResourceInfoModel)`

SetRelatedResource sets RelatedResource field to given value.

### HasRelatedResource

`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) HasRelatedResource() bool`

HasRelatedResource returns a boolean if a field has been set.

### SetRelatedResourceNil

`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) SetRelatedResourceNil(b bool)`

 SetRelatedResourceNil sets the value for RelatedResource to be an explicit nil

### UnsetRelatedResource
`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) UnsetRelatedResource()`

UnsetRelatedResource ensures that no value is present for RelatedResource, not even an explicit nil
### GetCreatedAt

`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetUpdatedAt

`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *BnsNetworkV1ApiListPublicIpsModelFloatingIpModel) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


