# EncryptData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Input** | **string** | 암호화할 데이터 | 
**Aad** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEncryptData

`func NewEncryptData(input string, ) *EncryptData`

NewEncryptData instantiates a new EncryptData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncryptDataWithDefaults

`func NewEncryptDataWithDefaults() *EncryptData`

NewEncryptDataWithDefaults instantiates a new EncryptData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInput

`func (o *EncryptData) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *EncryptData) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *EncryptData) SetInput(v string)`

SetInput sets Input field to given value.


### GetAad

`func (o *EncryptData) GetAad() string`

GetAad returns the Aad field if non-nil, zero value otherwise.

### GetAadOk

`func (o *EncryptData) GetAadOk() (*string, bool)`

GetAadOk returns a tuple with the Aad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAad

`func (o *EncryptData) SetAad(v string)`

SetAad sets Aad field to given value.

### HasAad

`func (o *EncryptData) HasAad() bool`

HasAad returns a boolean if a field has been set.

### SetAadNil

`func (o *EncryptData) SetAadNil(b bool)`

 SetAadNil sets the value for Aad to be an explicit nil

### UnsetAad
`func (o *EncryptData) UnsetAad()`

UnsetAad ensures that no value is present for Aad, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


