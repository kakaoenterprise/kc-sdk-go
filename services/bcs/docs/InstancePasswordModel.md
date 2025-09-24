# InstancePasswordModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **string** | 인스턴스에 설정된 관리자 비밀번호 &lt;br/&gt; - Windows: RDP(Remote Desktop Protocol) 로그인에 사용 &lt;br/&gt; - 비밀번호는 SSH 공개키로 암호화된 상태로 제공되며, RDP 접속을 위해서는 사용자의 SSH 개인키로 복호화해야 함 | 

## Methods

### NewInstancePasswordModel

`func NewInstancePasswordModel(password string, ) *InstancePasswordModel`

NewInstancePasswordModel instantiates a new InstancePasswordModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstancePasswordModelWithDefaults

`func NewInstancePasswordModelWithDefaults() *InstancePasswordModel`

NewInstancePasswordModelWithDefaults instantiates a new InstancePasswordModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *InstancePasswordModel) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *InstancePasswordModel) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *InstancePasswordModel) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


