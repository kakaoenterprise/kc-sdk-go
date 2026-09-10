# InstanceSg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | 보안 그룹의 ID | 
**Name** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewInstanceSg

`func NewInstanceSg(id string, ) *InstanceSg`

NewInstanceSg instantiates a new InstanceSg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceSgWithDefaults

`func NewInstanceSgWithDefaults() *InstanceSg`

NewInstanceSgWithDefaults instantiates a new InstanceSg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceSg) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceSg) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceSg) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *InstanceSg) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceSg) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceSg) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceSg) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *InstanceSg) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *InstanceSg) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


