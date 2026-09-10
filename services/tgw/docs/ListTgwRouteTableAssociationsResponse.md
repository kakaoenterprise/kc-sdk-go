# ListTgwRouteTableAssociationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Associations** | [**[]Association**](Association.md) | 조회된 Association 목록 | 
**Pagination** | [**Pagination**](Pagination.md) | 페이지네이션 메타데이터 | 

## Methods

### NewListTgwRouteTableAssociationsResponse

`func NewListTgwRouteTableAssociationsResponse(associations []Association, pagination Pagination, ) *ListTgwRouteTableAssociationsResponse`

NewListTgwRouteTableAssociationsResponse instantiates a new ListTgwRouteTableAssociationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTgwRouteTableAssociationsResponseWithDefaults

`func NewListTgwRouteTableAssociationsResponseWithDefaults() *ListTgwRouteTableAssociationsResponse`

NewListTgwRouteTableAssociationsResponseWithDefaults instantiates a new ListTgwRouteTableAssociationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociations

`func (o *ListTgwRouteTableAssociationsResponse) GetAssociations() []Association`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *ListTgwRouteTableAssociationsResponse) GetAssociationsOk() (*[]Association, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *ListTgwRouteTableAssociationsResponse) SetAssociations(v []Association)`

SetAssociations sets Associations field to given value.


### GetPagination

`func (o *ListTgwRouteTableAssociationsResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListTgwRouteTableAssociationsResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListTgwRouteTableAssociationsResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


