# PaginationModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** | 전체 항목 수 | [optional] [default to 0]
**Offset** | Pointer to **int32** | 조회 시작 위치 | [optional] [default to 0]
**Limit** | Pointer to **int32** | 페이지당 최대 반환 항목 수 | [optional] [default to 20]

## Methods

### NewPaginationModel

`func NewPaginationModel() *PaginationModel`

NewPaginationModel instantiates a new PaginationModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginationModelWithDefaults

`func NewPaginationModelWithDefaults() *PaginationModel`

NewPaginationModelWithDefaults instantiates a new PaginationModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *PaginationModel) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PaginationModel) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PaginationModel) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PaginationModel) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetOffset

`func (o *PaginationModel) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *PaginationModel) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *PaginationModel) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *PaginationModel) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *PaginationModel) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PaginationModel) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PaginationModel) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PaginationModel) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


