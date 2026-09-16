# UpdateImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | 변경할 이미지의 이름 | [optional] 
**Description** | Pointer to **NullableString** | 이미지에 대한 설명 - 필드를 사용하지 않거나 &#x60;null&#x60;이면 기존값 유지 | [optional] 

## Methods

### NewUpdateImage

`func NewUpdateImage() *UpdateImage`

NewUpdateImage instantiates a new UpdateImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateImageWithDefaults

`func NewUpdateImageWithDefaults() *UpdateImage`

NewUpdateImageWithDefaults instantiates a new UpdateImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateImage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateImage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateImage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateImage) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateImage) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateImage) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *UpdateImage) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateImage) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateImage) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateImage) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateImage) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateImage) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


