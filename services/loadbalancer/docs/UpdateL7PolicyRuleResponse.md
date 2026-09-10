# UpdateL7PolicyRuleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**L7Rule** | [**RuleResult**](RuleResult.md) | L7 정책에 연결된 규칙의 규칙 및 상태를 수정하는 요청 본문 | 

## Methods

### NewUpdateL7PolicyRuleResponse

`func NewUpdateL7PolicyRuleResponse(l7Rule RuleResult, ) *UpdateL7PolicyRuleResponse`

NewUpdateL7PolicyRuleResponse instantiates a new UpdateL7PolicyRuleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateL7PolicyRuleResponseWithDefaults

`func NewUpdateL7PolicyRuleResponseWithDefaults() *UpdateL7PolicyRuleResponse`

NewUpdateL7PolicyRuleResponseWithDefaults instantiates a new UpdateL7PolicyRuleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetL7Rule

`func (o *UpdateL7PolicyRuleResponse) GetL7Rule() RuleResult`

GetL7Rule returns the L7Rule field if non-nil, zero value otherwise.

### GetL7RuleOk

`func (o *UpdateL7PolicyRuleResponse) GetL7RuleOk() (*RuleResult, bool)`

GetL7RuleOk returns a tuple with the L7Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL7Rule

`func (o *UpdateL7PolicyRuleResponse) SetL7Rule(v RuleResult)`

SetL7Rule sets L7Rule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


