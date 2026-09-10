# ListSecretsAccessTargetsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Targets** | Pointer to [**[]Target**](Target.md) |  | [optional] 

## Methods

### NewListSecretsAccessTargetsResponse

`func NewListSecretsAccessTargetsResponse() *ListSecretsAccessTargetsResponse`

NewListSecretsAccessTargetsResponse instantiates a new ListSecretsAccessTargetsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSecretsAccessTargetsResponseWithDefaults

`func NewListSecretsAccessTargetsResponseWithDefaults() *ListSecretsAccessTargetsResponse`

NewListSecretsAccessTargetsResponseWithDefaults instantiates a new ListSecretsAccessTargetsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargets

`func (o *ListSecretsAccessTargetsResponse) GetTargets() []Target`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *ListSecretsAccessTargetsResponse) GetTargetsOk() (*[]Target, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *ListSecretsAccessTargetsResponse) SetTargets(v []Target)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *ListSecretsAccessTargetsResponse) HasTargets() bool`

HasTargets returns a boolean if a field has been set.

### SetTargetsNil

`func (o *ListSecretsAccessTargetsResponse) SetTargetsNil(b bool)`

 SetTargetsNil sets the value for Targets to be an explicit nil

### UnsetTargets
`func (o *ListSecretsAccessTargetsResponse) UnsetTargets()`

UnsetTargets ensures that no value is present for Targets, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


