# BcsInstanceV1ApiCreateKeypairModelKeypairModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 사용자가 지정한 키 페어 이름 | 
**UserId** | **string** | 키 페어를 생성한 사용자의 ID | 
**Fingerprint** | **string** | 퍼블릭 키의 고유 핑거프린트 &lt;br/&gt; - SSH 클라이언트에서 키 검증 시 사용 | 
**PublicKey** | **string** | 생성된 키 페어의 퍼블릭 키 &lt;br/&gt; - 인스턴스에 등록되어 SSH 접속 시 인증에 사용 | 
**PrivateKey** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** | SSH 키 형식 | 

## Methods

### NewBcsInstanceV1ApiCreateKeypairModelKeypairModel

`func NewBcsInstanceV1ApiCreateKeypairModelKeypairModel(name string, userId string, fingerprint string, publicKey string, type_ string, ) *BcsInstanceV1ApiCreateKeypairModelKeypairModel`

NewBcsInstanceV1ApiCreateKeypairModelKeypairModel instantiates a new BcsInstanceV1ApiCreateKeypairModelKeypairModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBcsInstanceV1ApiCreateKeypairModelKeypairModelWithDefaults

`func NewBcsInstanceV1ApiCreateKeypairModelKeypairModelWithDefaults() *BcsInstanceV1ApiCreateKeypairModelKeypairModel`

NewBcsInstanceV1ApiCreateKeypairModelKeypairModelWithDefaults instantiates a new BcsInstanceV1ApiCreateKeypairModelKeypairModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BcsInstanceV1ApiCreateKeypairModelKeypairModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BcsInstanceV1ApiCreateKeypairModelKeypairModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BcsInstanceV1ApiCreateKeypairModelKeypairModel) SetName(v string)`

SetName sets Name field to given value.


### GetUserId

`func (o *BcsInstanceV1ApiCreateKeypairModelKeypairModel) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *BcsInstanceV1ApiCreateKeypairModelKeypairModel) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *BcsInstanceV1ApiCreateKeypairModelKeypairModel) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetFingerprint

`func (o *BcsInstanceV1ApiCreateKeypairModelKeypairModel) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *BcsInstanceV1ApiCreateKeypairModelKeypairModel) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *BcsInstanceV1ApiCreateKeypairModelKeypairModel) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.


### GetPublicKey

`func (o *BcsInstanceV1ApiCreateKeypairModelKeypairModel) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *BcsInstanceV1ApiCreateKeypairModelKeypairModel) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *BcsInstanceV1ApiCreateKeypairModelKeypairModel) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### GetPrivateKey

`func (o *BcsInstanceV1ApiCreateKeypairModelKeypairModel) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *BcsInstanceV1ApiCreateKeypairModelKeypairModel) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *BcsInstanceV1ApiCreateKeypairModelKeypairModel) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *BcsInstanceV1ApiCreateKeypairModelKeypairModel) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### SetPrivateKeyNil

`func (o *BcsInstanceV1ApiCreateKeypairModelKeypairModel) SetPrivateKeyNil(b bool)`

 SetPrivateKeyNil sets the value for PrivateKey to be an explicit nil

### UnsetPrivateKey
`func (o *BcsInstanceV1ApiCreateKeypairModelKeypairModel) UnsetPrivateKey()`

UnsetPrivateKey ensures that no value is present for PrivateKey, not even an explicit nil
### GetType

`func (o *BcsInstanceV1ApiCreateKeypairModelKeypairModel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BcsInstanceV1ApiCreateKeypairModelKeypairModel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BcsInstanceV1ApiCreateKeypairModelKeypairModel) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


