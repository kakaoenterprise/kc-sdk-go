# GetMySQLEngineVersionsResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EngineVersions** | [**[]EngineVersionResponseModel**](EngineVersionResponseModel.md) | 사용 가능한 MySQL 엔진 버전 목록 | 

## Methods

### NewGetMySQLEngineVersionsResponseModel

`func NewGetMySQLEngineVersionsResponseModel(engineVersions []EngineVersionResponseModel, ) *GetMySQLEngineVersionsResponseModel`

NewGetMySQLEngineVersionsResponseModel instantiates a new GetMySQLEngineVersionsResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMySQLEngineVersionsResponseModelWithDefaults

`func NewGetMySQLEngineVersionsResponseModelWithDefaults() *GetMySQLEngineVersionsResponseModel`

NewGetMySQLEngineVersionsResponseModelWithDefaults instantiates a new GetMySQLEngineVersionsResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngineVersions

`func (o *GetMySQLEngineVersionsResponseModel) GetEngineVersions() []EngineVersionResponseModel`

GetEngineVersions returns the EngineVersions field if non-nil, zero value otherwise.

### GetEngineVersionsOk

`func (o *GetMySQLEngineVersionsResponseModel) GetEngineVersionsOk() (*[]EngineVersionResponseModel, bool)`

GetEngineVersionsOk returns a tuple with the EngineVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineVersions

`func (o *GetMySQLEngineVersionsResponseModel) SetEngineVersions(v []EngineVersionResponseModel)`

SetEngineVersions sets EngineVersions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


