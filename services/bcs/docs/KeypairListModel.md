# KeypairListModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keypairs** | Pointer to [**[]BcsInstanceV1ApiListKeypairsModelKeypairModel**](BcsInstanceV1ApiListKeypairsModelKeypairModel.md) |  | [optional] 
**Pagination** | [**PaginationModel**](PaginationModel.md) |  | 

## Methods

### NewKeypairListModel

`func NewKeypairListModel(pagination PaginationModel, ) *KeypairListModel`

NewKeypairListModel instantiates a new KeypairListModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeypairListModelWithDefaults

`func NewKeypairListModelWithDefaults() *KeypairListModel`

NewKeypairListModelWithDefaults instantiates a new KeypairListModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeypairs

`func (o *KeypairListModel) GetKeypairs() []BcsInstanceV1ApiListKeypairsModelKeypairModel`

GetKeypairs returns the Keypairs field if non-nil, zero value otherwise.

### GetKeypairsOk

`func (o *KeypairListModel) GetKeypairsOk() (*[]BcsInstanceV1ApiListKeypairsModelKeypairModel, bool)`

GetKeypairsOk returns a tuple with the Keypairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypairs

`func (o *KeypairListModel) SetKeypairs(v []BcsInstanceV1ApiListKeypairsModelKeypairModel)`

SetKeypairs sets Keypairs field to given value.

### HasKeypairs

`func (o *KeypairListModel) HasKeypairs() bool`

HasKeypairs returns a boolean if a field has been set.

### GetPagination

`func (o *KeypairListModel) GetPagination() PaginationModel`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *KeypairListModel) GetPaginationOk() (*PaginationModel, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *KeypairListModel) SetPagination(v PaginationModel)`

SetPagination sets Pagination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


