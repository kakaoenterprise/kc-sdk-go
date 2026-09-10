# Spec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Taints** | Pointer to [**[]Taint**](Taint.md) |  | [optional] 

## Methods

### NewSpec

`func NewSpec() *Spec`

NewSpec instantiates a new Spec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpecWithDefaults

`func NewSpecWithDefaults() *Spec`

NewSpecWithDefaults instantiates a new Spec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaints

`func (o *Spec) GetTaints() []Taint`

GetTaints returns the Taints field if non-nil, zero value otherwise.

### GetTaintsOk

`func (o *Spec) GetTaintsOk() (*[]Taint, bool)`

GetTaintsOk returns a tuple with the Taints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaints

`func (o *Spec) SetTaints(v []Taint)`

SetTaints sets Taints field to given value.

### HasTaints

`func (o *Spec) HasTaints() bool`

HasTaints returns a boolean if a field has been set.

### SetTaintsNil

`func (o *Spec) SetTaintsNil(b bool)`

 SetTaintsNil sets the value for Taints to be an explicit nil

### UnsetTaints
`func (o *Spec) UnsetTaints()`

UnsetTaints ensures that no value is present for Taints, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


