# Version

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestroyStatus** | **string** | 키 버전 폐기 예약 취소 결과 상태 | 

## Methods

### NewVersion

`func NewVersion(destroyStatus string, ) *Version`

NewVersion instantiates a new Version object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionWithDefaults

`func NewVersionWithDefaults() *Version`

NewVersionWithDefaults instantiates a new Version object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestroyStatus

`func (o *Version) GetDestroyStatus() string`

GetDestroyStatus returns the DestroyStatus field if non-nil, zero value otherwise.

### GetDestroyStatusOk

`func (o *Version) GetDestroyStatusOk() (*string, bool)`

GetDestroyStatusOk returns a tuple with the DestroyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestroyStatus

`func (o *Version) SetDestroyStatus(v string)`

SetDestroyStatus sets DestroyStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


