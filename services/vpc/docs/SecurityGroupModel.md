# SecurityGroupModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 보안 그룹의 ID | 
**Name** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSecurityGroupModel

`func NewSecurityGroupModel(id string, ) *SecurityGroupModel`

NewSecurityGroupModel instantiates a new SecurityGroupModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGroupModelWithDefaults

`func NewSecurityGroupModelWithDefaults() *SecurityGroupModel`

NewSecurityGroupModelWithDefaults instantiates a new SecurityGroupModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecurityGroupModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityGroupModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityGroupModel) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SecurityGroupModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityGroupModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityGroupModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecurityGroupModel) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SecurityGroupModel) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SecurityGroupModel) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


