# ListL7PolicyRulesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**L7Rules** | [**[]Rule**](Rule.md) | L7 규칙 목록 | 
**Pagination** | [**Pagination**](Pagination.md) | 페이지네이션 정보 | 

## Methods

### NewListL7PolicyRulesResponse

`func NewListL7PolicyRulesResponse(l7Rules []Rule, pagination Pagination, ) *ListL7PolicyRulesResponse`

NewListL7PolicyRulesResponse instantiates a new ListL7PolicyRulesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListL7PolicyRulesResponseWithDefaults

`func NewListL7PolicyRulesResponseWithDefaults() *ListL7PolicyRulesResponse`

NewListL7PolicyRulesResponseWithDefaults instantiates a new ListL7PolicyRulesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetL7Rules

`func (o *ListL7PolicyRulesResponse) GetL7Rules() []Rule`

GetL7Rules returns the L7Rules field if non-nil, zero value otherwise.

### GetL7RulesOk

`func (o *ListL7PolicyRulesResponse) GetL7RulesOk() (*[]Rule, bool)`

GetL7RulesOk returns a tuple with the L7Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetL7Rules

`func (o *ListL7PolicyRulesResponse) SetL7Rules(v []Rule)`

SetL7Rules sets L7Rules field to given value.


### GetPagination

`func (o *ListL7PolicyRulesResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListL7PolicyRulesResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListL7PolicyRulesResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


