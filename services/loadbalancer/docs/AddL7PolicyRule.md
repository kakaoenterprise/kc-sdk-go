# AddL7PolicyRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompareType** | [**L7RuleCompareType**](L7RuleCompareType.md) | 규칙 값 비교 방식 | 
**IsInverted** | Pointer to **NullableBool** |  | [optional] 
**Key** | Pointer to **NullableString** |  | [optional] 
**Type** | [**L7RuleType**](L7RuleType.md) | 규칙을 검사할 대상 유형 &lt;br/&gt;- &#x60;COOKIE&#x60;: HTTP 요청의 특정 쿠키 값을 기준으로 검사&lt;br/&gt;- &#x60;FILE_TYPE&#x60;: 요청한 리소스의 파일 확장자(예: &#x60;.jpg&#x60;, &#x60;.exe&#x60;)를 기준으로 검사&lt;br/&gt;- &#x60;HEADER&#x60;: HTTP 요청 헤더 값을 기준으로 검사&lt;br/&gt;- &#x60;HOST_NAME&#x60;: 요청의 Host 헤더(도메인 이름)를 기준으로 검사&lt;br/&gt;- &#x60;PATH&#x60;: 요청 URL 경로(예: &#x60;/images/_*&#x60;, &#x60;/api/v1/_*&#x60;)를 기준으로 검사 | 
**Value** | **string** | 비교할 대상 값 | 

## Methods

### NewAddL7PolicyRule

`func NewAddL7PolicyRule(compareType L7RuleCompareType, type_ L7RuleType, value string, ) *AddL7PolicyRule`

NewAddL7PolicyRule instantiates a new AddL7PolicyRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddL7PolicyRuleWithDefaults

`func NewAddL7PolicyRuleWithDefaults() *AddL7PolicyRule`

NewAddL7PolicyRuleWithDefaults instantiates a new AddL7PolicyRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompareType

`func (o *AddL7PolicyRule) GetCompareType() L7RuleCompareType`

GetCompareType returns the CompareType field if non-nil, zero value otherwise.

### GetCompareTypeOk

`func (o *AddL7PolicyRule) GetCompareTypeOk() (*L7RuleCompareType, bool)`

GetCompareTypeOk returns a tuple with the CompareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareType

`func (o *AddL7PolicyRule) SetCompareType(v L7RuleCompareType)`

SetCompareType sets CompareType field to given value.


### GetIsInverted

`func (o *AddL7PolicyRule) GetIsInverted() bool`

GetIsInverted returns the IsInverted field if non-nil, zero value otherwise.

### GetIsInvertedOk

`func (o *AddL7PolicyRule) GetIsInvertedOk() (*bool, bool)`

GetIsInvertedOk returns a tuple with the IsInverted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInverted

`func (o *AddL7PolicyRule) SetIsInverted(v bool)`

SetIsInverted sets IsInverted field to given value.

### HasIsInverted

`func (o *AddL7PolicyRule) HasIsInverted() bool`

HasIsInverted returns a boolean if a field has been set.

### SetIsInvertedNil

`func (o *AddL7PolicyRule) SetIsInvertedNil(b bool)`

 SetIsInvertedNil sets the value for IsInverted to be an explicit nil

### UnsetIsInverted
`func (o *AddL7PolicyRule) UnsetIsInverted()`

UnsetIsInverted ensures that no value is present for IsInverted, not even an explicit nil
### GetKey

`func (o *AddL7PolicyRule) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AddL7PolicyRule) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AddL7PolicyRule) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *AddL7PolicyRule) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *AddL7PolicyRule) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *AddL7PolicyRule) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetType

`func (o *AddL7PolicyRule) GetType() L7RuleType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddL7PolicyRule) GetTypeOk() (*L7RuleType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddL7PolicyRule) SetType(v L7RuleType)`

SetType sets Type field to given value.


### GetValue

`func (o *AddL7PolicyRule) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AddL7PolicyRule) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AddL7PolicyRule) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


