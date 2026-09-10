# AddL7PolicyRuleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**L7Rule** | [**RuleResult**](RuleResult.md) | 생성된 L7 규칙 정보 | 

## Methods

### NewAddL7PolicyRuleResponse

`func NewAddL7PolicyRuleResponse(l7Rule RuleResult, ) *AddL7PolicyRuleResponse`

NewAddL7PolicyRuleResponse instantiates a new AddL7PolicyRuleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddL7PolicyRuleResponseWithDefaults

`func NewAddL7PolicyRuleResponseWithDefaults() *AddL7PolicyRuleResponse`

NewAddL7PolicyRuleResponseWithDefaults instantiates a new AddL7PolicyRuleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetL7Rule

`func (o *AddL7PolicyRuleResponse) GetL7Rule() RuleResult`

GetL7Rule returns the L7Rule field if non-nil, zero value otherwise.

### GetL7RuleOk

`func (o *AddL7PolicyRuleResponse) GetL7RuleOk() (*RuleResult, bool)`

GetL7RuleOk returns a tuple with the L7Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL7Rule

`func (o *AddL7PolicyRuleResponse) SetL7Rule(v RuleResult)`

SetL7Rule sets L7Rule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


