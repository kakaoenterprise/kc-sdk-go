# CreateSecurityGroupModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 생성할 보안 그룹 이름 &lt;br/&gt; - 중복 이름은 허용하지 않음 | 
**Description** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateSecurityGroupModel

`func NewCreateSecurityGroupModel(name string, ) *CreateSecurityGroupModel`

NewCreateSecurityGroupModel instantiates a new CreateSecurityGroupModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSecurityGroupModelWithDefaults

`func NewCreateSecurityGroupModelWithDefaults() *CreateSecurityGroupModel`

NewCreateSecurityGroupModelWithDefaults instantiates a new CreateSecurityGroupModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateSecurityGroupModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSecurityGroupModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSecurityGroupModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateSecurityGroupModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateSecurityGroupModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateSecurityGroupModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateSecurityGroupModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateSecurityGroupModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateSecurityGroupModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


