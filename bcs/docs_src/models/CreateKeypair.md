# CreateKeypair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 생성할 키페어의 이름 - 같은 프로젝트 내 중복된 키 페어 이름 사용 불가 | 
**PublicKey** | Pointer to **NullableString** | 인스턴스 접근 보안을 위해 인스턴스에 등록하고 관리하는 퍼블릭 키 | [optional] 

## Methods

### NewCreateKeypair

`func NewCreateKeypair(name string, ) *CreateKeypair`

NewCreateKeypair instantiates a new CreateKeypair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateKeypairWithDefaults

`func NewCreateKeypairWithDefaults() *CreateKeypair`

NewCreateKeypairWithDefaults instantiates a new CreateKeypair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateKeypair) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateKeypair) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateKeypair) SetName(v string)`

SetName sets Name field to given value.


### GetPublicKey

`func (o *CreateKeypair) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *CreateKeypair) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *CreateKeypair) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *CreateKeypair) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### SetPublicKeyNil

`func (o *CreateKeypair) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *CreateKeypair) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


