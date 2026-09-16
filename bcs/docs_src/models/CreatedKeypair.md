# CreatedKeypair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrivateKey** | Pointer to **NullableString** | 생성된 키 페어의 프라이빗 키 - 사용자의 로컬 환경에 저장되며, SSH 접속 시 인증에 사용 | [optional] 
**Type** | **string** | SSH 키 형식 | 
**Name** | **string** | 사용자가 지정한 키 페어 이름 | 
**PublicKey** | **string** | 생성된 키 페어의 퍼블릭 키 - 인스턴스에 등록되어 SSH 접속 시 인증에 사용 | 
**Fingerprint** | **string** | 퍼블릭 키의 고유 핑거프린트 - SSH 클라이언트에서 키 검증 시 사용 | 
**UserId** | **string** | 키 페어를 생성한 사용자의 ID | 

## Methods

### NewCreatedKeypair

`func NewCreatedKeypair(type_ string, name string, publicKey string, fingerprint string, userId string, ) *CreatedKeypair`

NewCreatedKeypair instantiates a new CreatedKeypair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatedKeypairWithDefaults

`func NewCreatedKeypairWithDefaults() *CreatedKeypair`

NewCreatedKeypairWithDefaults instantiates a new CreatedKeypair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrivateKey

`func (o *CreatedKeypair) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *CreatedKeypair) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *CreatedKeypair) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *CreatedKeypair) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### SetPrivateKeyNil

`func (o *CreatedKeypair) SetPrivateKeyNil(b bool)`

 SetPrivateKeyNil sets the value for PrivateKey to be an explicit nil

### UnsetPrivateKey
`func (o *CreatedKeypair) UnsetPrivateKey()`

UnsetPrivateKey ensures that no value is present for PrivateKey, not even an explicit nil
### GetType

`func (o *CreatedKeypair) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreatedKeypair) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreatedKeypair) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *CreatedKeypair) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatedKeypair) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatedKeypair) SetName(v string)`

SetName sets Name field to given value.


### GetPublicKey

`func (o *CreatedKeypair) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *CreatedKeypair) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *CreatedKeypair) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### GetFingerprint

`func (o *CreatedKeypair) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *CreatedKeypair) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *CreatedKeypair) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.


### GetUserId

`func (o *CreatedKeypair) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreatedKeypair) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreatedKeypair) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


