# GetTgwRouteTableAssociationsResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Associations** | [**[]BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel**](BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel.md) | 조회된 Association 목록 | 
**Pagination** | [**PaginationModel**](PaginationModel.md) | 페이지네이션 메타데이터 | 

## Methods

### NewGetTgwRouteTableAssociationsResponseModel

`func NewGetTgwRouteTableAssociationsResponseModel(associations []BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel, pagination PaginationModel, ) *GetTgwRouteTableAssociationsResponseModel`

NewGetTgwRouteTableAssociationsResponseModel instantiates a new GetTgwRouteTableAssociationsResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTgwRouteTableAssociationsResponseModelWithDefaults

`func NewGetTgwRouteTableAssociationsResponseModelWithDefaults() *GetTgwRouteTableAssociationsResponseModel`

NewGetTgwRouteTableAssociationsResponseModelWithDefaults instantiates a new GetTgwRouteTableAssociationsResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociations

`func (o *GetTgwRouteTableAssociationsResponseModel) GetAssociations() []BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel`

GetAssociations returns the Associations field if non-nil, zero value otherwise.

### GetAssociationsOk

`func (o *GetTgwRouteTableAssociationsResponseModel) GetAssociationsOk() (*[]BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel, bool)`

GetAssociationsOk returns a tuple with the Associations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociations

`func (o *GetTgwRouteTableAssociationsResponseModel) SetAssociations(v []BnsTgwV1ApiListTgwRouteTableAssociationsModelAssociationResponseModel)`

SetAssociations sets Associations field to given value.


### GetPagination

`func (o *GetTgwRouteTableAssociationsResponseModel) GetPagination() PaginationModel`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetTgwRouteTableAssociationsResponseModel) GetPaginationOk() (*PaginationModel, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetTgwRouteTableAssociationsResponseModel) SetPagination(v PaginationModel)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


