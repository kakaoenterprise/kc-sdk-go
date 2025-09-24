# CreateKeypairModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 생성할 키페어의 이름  &lt;br/&gt;- 같은 프로젝트 내 중복된 키 페어 이름 사용 불가 | 
**PublicKey** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateKeypairModel

`func NewCreateKeypairModel(name string, ) *CreateKeypairModel`

NewCreateKeypairModel instantiates a new CreateKeypairModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateKeypairModelWithDefaults

`func NewCreateKeypairModelWithDefaults() *CreateKeypairModel`

NewCreateKeypairModelWithDefaults instantiates a new CreateKeypairModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateKeypairModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateKeypairModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateKeypairModel) SetName(v string)`

SetName sets Name field to given value.


### GetPublicKey

`func (o *CreateKeypairModel) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *CreateKeypairModel) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *CreateKeypairModel) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *CreateKeypairModel) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### SetPublicKeyNil

`func (o *CreateKeypairModel) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *CreateKeypairModel) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


