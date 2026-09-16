# EngineVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EngineVersion** | **string** | 사용 가능한 MySQL 엔진 버전 | 
**License** | **string** | 해당 엔진 버전에 적용되는 라이선스 정보 | 

## Methods

### NewEngineVersion

`func NewEngineVersion(engineVersion string, license string, ) *EngineVersion`

NewEngineVersion instantiates a new EngineVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineVersionWithDefaults

`func NewEngineVersionWithDefaults() *EngineVersion`

NewEngineVersionWithDefaults instantiates a new EngineVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngineVersion

`func (o *EngineVersion) GetEngineVersion() string`

GetEngineVersion returns the EngineVersion field if non-nil, zero value otherwise.

### GetEngineVersionOk

`func (o *EngineVersion) GetEngineVersionOk() (*string, bool)`

GetEngineVersionOk returns a tuple with the EngineVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineVersion

`func (o *EngineVersion) SetEngineVersion(v string)`

SetEngineVersion sets EngineVersion field to given value.


### GetLicense

`func (o *EngineVersion) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *EngineVersion) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *EngineVersion) SetLicense(v string)`

SetLicense sets License field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


