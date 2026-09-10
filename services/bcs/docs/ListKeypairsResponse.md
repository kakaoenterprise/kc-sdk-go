# ListKeypairsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keypairs** | [**[]Keypair**](Keypair.md) | 키페어 목록 | 
**Pagination** | [**Pagination**](Pagination.md) | 페이지네이션 정보 | 

## Methods

### NewListKeypairsResponse

`func NewListKeypairsResponse(keypairs []Keypair, pagination Pagination, ) *ListKeypairsResponse`

NewListKeypairsResponse instantiates a new ListKeypairsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListKeypairsResponseWithDefaults

`func NewListKeypairsResponseWithDefaults() *ListKeypairsResponse`

NewListKeypairsResponseWithDefaults instantiates a new ListKeypairsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeypairs

`func (o *ListKeypairsResponse) GetKeypairs() []Keypair`

GetKeypairs returns the Keypairs field if non-nil, zero value otherwise.

### GetKeypairsOk

`func (o *ListKeypairsResponse) GetKeypairsOk() (*[]Keypair, bool)`

GetKeypairsOk returns a tuple with the Keypairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypairs

`func (o *ListKeypairsResponse) SetKeypairs(v []Keypair)`

SetKeypairs sets Keypairs field to given value.


### GetPagination

`func (o *ListKeypairsResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListKeypairsResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListKeypairsResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


