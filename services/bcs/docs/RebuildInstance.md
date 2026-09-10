# RebuildInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageId** | **string** | 이미지의 고유 ID &lt;br/&gt; - [List images](https://docs.kakaocloud.com/openapi/bcs/list-images)에서 확인 | 
**KeyName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRebuildInstance

`func NewRebuildInstance(imageId string, ) *RebuildInstance`

NewRebuildInstance instantiates a new RebuildInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRebuildInstanceWithDefaults

`func NewRebuildInstanceWithDefaults() *RebuildInstance`

NewRebuildInstanceWithDefaults instantiates a new RebuildInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageId

`func (o *RebuildInstance) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *RebuildInstance) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *RebuildInstance) SetImageId(v string)`

SetImageId sets ImageId field to given value.


### GetKeyName

`func (o *RebuildInstance) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *RebuildInstance) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *RebuildInstance) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.

### HasKeyName

`func (o *RebuildInstance) HasKeyName() bool`

HasKeyName returns a boolean if a field has been set.

### SetKeyNameNil

`func (o *RebuildInstance) SetKeyNameNil(b bool)`

 SetKeyNameNil sets the value for KeyName to be an explicit nil

### UnsetKeyName
`func (o *RebuildInstance) UnsetKeyName()`

UnsetKeyName ensures that no value is present for KeyName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


