# AddL7PolicyRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**L7Rule** | [**AddL7PolicyRule**](AddL7PolicyRule.md) | 생성할 L7 규칙 정보 | 

## Methods

### NewAddL7PolicyRuleRequest

`func NewAddL7PolicyRuleRequest(l7Rule AddL7PolicyRule, ) *AddL7PolicyRuleRequest`

NewAddL7PolicyRuleRequest instantiates a new AddL7PolicyRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddL7PolicyRuleRequestWithDefaults

`func NewAddL7PolicyRuleRequestWithDefaults() *AddL7PolicyRuleRequest`

NewAddL7PolicyRuleRequestWithDefaults instantiates a new AddL7PolicyRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetL7Rule

`func (o *AddL7PolicyRuleRequest) GetL7Rule() AddL7PolicyRule`

GetL7Rule returns the L7Rule field if non-nil, zero value otherwise.

### GetL7RuleOk

`func (o *AddL7PolicyRuleRequest) GetL7RuleOk() (*AddL7PolicyRule, bool)`

GetL7RuleOk returns a tuple with the L7Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL7Rule

`func (o *AddL7PolicyRuleRequest) SetL7Rule(v AddL7PolicyRule)`

SetL7Rule sets L7Rule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


