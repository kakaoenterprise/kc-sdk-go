# ContentTypesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | **string** | 인증서의 기본 MIME 유형 - 예시: &#x60;application/x-pem-file&#x60;, &#x60;application/octet-stream&#x60; 등 | 

## Methods

### NewContentTypesDto

`func NewContentTypesDto(default_ string, ) *ContentTypesDto`

NewContentTypesDto instantiates a new ContentTypesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentTypesDtoWithDefaults

`func NewContentTypesDtoWithDefaults() *ContentTypesDto`

NewContentTypesDtoWithDefaults instantiates a new ContentTypesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefault

`func (o *ContentTypesDto) GetDefault() string`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *ContentTypesDto) GetDefaultOk() (*string, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *ContentTypesDto) SetDefault(v string)`

SetDefault sets Default field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


