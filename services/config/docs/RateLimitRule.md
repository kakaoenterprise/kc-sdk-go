# RateLimitRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **string** |  | 
**Limit** | **int32** |  | 
**Action** | **string** |  | 
**Delay** | **string** |  | 

## Methods

### NewRateLimitRule

`func NewRateLimitRule(duration string, limit int32, action string, delay string, ) *RateLimitRule`

NewRateLimitRule instantiates a new RateLimitRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateLimitRuleWithDefaults

`func NewRateLimitRuleWithDefaults() *RateLimitRule`

NewRateLimitRuleWithDefaults instantiates a new RateLimitRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *RateLimitRule) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *RateLimitRule) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *RateLimitRule) SetDuration(v string)`

SetDuration sets Duration field to given value.


### GetLimit

`func (o *RateLimitRule) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *RateLimitRule) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *RateLimitRule) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetAction

`func (o *RateLimitRule) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RateLimitRule) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RateLimitRule) SetAction(v string)`

SetAction sets Action field to given value.


### GetDelay

`func (o *RateLimitRule) GetDelay() string`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *RateLimitRule) GetDelayOk() (*string, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *RateLimitRule) SetDelay(v string)`

SetDelay sets Delay field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


