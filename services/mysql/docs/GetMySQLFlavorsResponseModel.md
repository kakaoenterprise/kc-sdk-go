# GetMySQLFlavorsResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flavors** | [**[]FlavorResponseModel**](FlavorResponseModel.md) | 선택 가능한 인스턴스 유형(Flavor) 목록 | 

## Methods

### NewGetMySQLFlavorsResponseModel

`func NewGetMySQLFlavorsResponseModel(flavors []FlavorResponseModel, ) *GetMySQLFlavorsResponseModel`

NewGetMySQLFlavorsResponseModel instantiates a new GetMySQLFlavorsResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMySQLFlavorsResponseModelWithDefaults

`func NewGetMySQLFlavorsResponseModelWithDefaults() *GetMySQLFlavorsResponseModel`

NewGetMySQLFlavorsResponseModelWithDefaults instantiates a new GetMySQLFlavorsResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlavors

`func (o *GetMySQLFlavorsResponseModel) GetFlavors() []FlavorResponseModel`

GetFlavors returns the Flavors field if non-nil, zero value otherwise.

### GetFlavorsOk

`func (o *GetMySQLFlavorsResponseModel) GetFlavorsOk() (*[]FlavorResponseModel, bool)`

GetFlavorsOk returns a tuple with the Flavors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavors

`func (o *GetMySQLFlavorsResponseModel) SetFlavors(v []FlavorResponseModel)`

SetFlavors sets Flavors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


