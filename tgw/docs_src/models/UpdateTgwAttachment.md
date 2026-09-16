# UpdateTgwAttachment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | 변경할 Transit Gateway Attachment 이름 | [optional] 
**SubnetIds** | **[]string** | Attachment에 연결할 서브넷 ID 목록 - [List subnets](/openapi/networking/vpc/list-subnets)에서 확인 | 

## Methods

### NewUpdateTgwAttachment

`func NewUpdateTgwAttachment(subnetIds []string, ) *UpdateTgwAttachment`

NewUpdateTgwAttachment instantiates a new UpdateTgwAttachment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTgwAttachmentWithDefaults

`func NewUpdateTgwAttachmentWithDefaults() *UpdateTgwAttachment`

NewUpdateTgwAttachmentWithDefaults instantiates a new UpdateTgwAttachment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateTgwAttachment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateTgwAttachment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateTgwAttachment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateTgwAttachment) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateTgwAttachment) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateTgwAttachment) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSubnetIds

`func (o *UpdateTgwAttachment) GetSubnetIds() []string`

GetSubnetIds returns the SubnetIds field if non-nil, zero value otherwise.

### GetSubnetIdsOk

`func (o *UpdateTgwAttachment) GetSubnetIdsOk() (*[]string, bool)`

GetSubnetIdsOk returns a tuple with the SubnetIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetIds

`func (o *UpdateTgwAttachment) SetSubnetIds(v []string)`

SetSubnetIds sets SubnetIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


