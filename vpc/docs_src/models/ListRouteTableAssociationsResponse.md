# ListRouteTableAssociationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Associations** | [**[]VpcAssociation**](VpcAssociation.md) | 라우팅 테이블 연결 목록 | 
**Pagination** | [**Pagination**](Pagination.md) | 페이지네이션 정보 | 

## Methods

### NewListRouteTableAssociationsResponse

`func NewListRouteTableAssociationsResponse(associations []VpcAssociation, pagination Pagination, ) *ListRouteTableAssociationsResponse`

NewListRouteTableAssociationsResponse instantiates a new ListRouteTableAssociationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRouteTableAssociationsResponseWithDefaults

`func NewListRouteTableAssociationsResponseWithDefaults() *ListRouteTableAssociationsResponse`

NewListRouteTableAssociationsResponseWithDefaults instantiates a new ListRouteTableAssociationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociations

`func (o *ListRouteTableAssociationsResponse) GetAssociations() []VpcAssociation`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *ListRouteTableAssociationsResponse) GetAssociationsOk() (*[]VpcAssociation, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *ListRouteTableAssociationsResponse) SetAssociations(v []VpcAssociation)`

SetAssociations sets Associations field to given value.


### GetPagination

`func (o *ListRouteTableAssociationsResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListRouteTableAssociationsResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListRouteTableAssociationsResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


