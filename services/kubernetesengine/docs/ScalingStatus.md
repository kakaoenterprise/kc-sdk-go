# ScalingStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Histories** | Pointer to [**[]ScalingHistory**](ScalingHistory.md) |  | [optional] 

## Methods

### NewScalingStatus

`func NewScalingStatus() *ScalingStatus`

NewScalingStatus instantiates a new ScalingStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScalingStatusWithDefaults

`func NewScalingStatusWithDefaults() *ScalingStatus`

NewScalingStatusWithDefaults instantiates a new ScalingStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHistories

`func (o *ScalingStatus) GetHistories() []ScalingHistory`

GetHistories returns the Histories field if non-nil, zero value otherwise.

### GetHistoriesOk

`func (o *ScalingStatus) GetHistoriesOk() (*[]ScalingHistory, bool)`

GetHistoriesOk returns a tuple with the Histories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistories

`func (o *ScalingStatus) SetHistories(v []ScalingHistory)`

SetHistories sets Histories field to given value.

### HasHistories

`func (o *ScalingStatus) HasHistories() bool`

HasHistories returns a boolean if a field has been set.

### SetHistoriesNil

`func (o *ScalingStatus) SetHistoriesNil(b bool)`

 SetHistoriesNil sets the value for Histories to be an explicit nil

### UnsetHistories
`func (o *ScalingStatus) UnsetHistories()`

UnsetHistories ensures that no value is present for Histories, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


