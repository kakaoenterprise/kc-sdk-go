# ListAvailableMysqlEngineVersionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EngineVersions** | [**[]EngineVersion**](EngineVersion.md) | 사용 가능한 MySQL 엔진 버전 목록 | 

## Methods

### NewListAvailableMysqlEngineVersionsResponse

`func NewListAvailableMysqlEngineVersionsResponse(engineVersions []EngineVersion, ) *ListAvailableMysqlEngineVersionsResponse`

NewListAvailableMysqlEngineVersionsResponse instantiates a new ListAvailableMysqlEngineVersionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAvailableMysqlEngineVersionsResponseWithDefaults

`func NewListAvailableMysqlEngineVersionsResponseWithDefaults() *ListAvailableMysqlEngineVersionsResponse`

NewListAvailableMysqlEngineVersionsResponseWithDefaults instantiates a new ListAvailableMysqlEngineVersionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngineVersions

`func (o *ListAvailableMysqlEngineVersionsResponse) GetEngineVersions() []EngineVersion`

GetEngineVersions returns the EngineVersions field if non-nil, zero value otherwise.

### GetEngineVersionsOk

`func (o *ListAvailableMysqlEngineVersionsResponse) GetEngineVersionsOk() (*[]EngineVersion, bool)`

GetEngineVersionsOk returns a tuple with the EngineVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineVersions

`func (o *ListAvailableMysqlEngineVersionsResponse) SetEngineVersions(v []EngineVersion)`

SetEngineVersions sets EngineVersions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


