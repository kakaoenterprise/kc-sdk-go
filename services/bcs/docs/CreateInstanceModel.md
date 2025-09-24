# CreateInstanceModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 생성할 인스턴스의 이름 &lt;br/&gt; - 같은 프로젝트 내 중복된 인스턴스 이름 사용 가능 | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Count** | Pointer to **int32** | 생성할 인스턴스 수 | [optional] [default to 1]
**ImageId** | **string** | 이미지의 고유 ID | 
**FlavorId** | **string** | 인스턴스 유형 ID &lt;br/&gt; - [List instance types](https://docs.kakaocloud.com/openapi/bcs/list-instance-types) API를 통해 조회 가능 | 
**AvailabilityZone** | Pointer to [**NullableAvailabilityZone**](AvailabilityZone.md) |  | [optional] 
**Subnets** | [**[]CreateInstanceSubnetModel**](CreateInstanceSubnetModel.md) | 연결할 서브넷 정보 목록 | 
**Volumes** | [**[]CreateInstanceVolumeModel**](CreateInstanceVolumeModel.md) | 인스턴스에 연결할 볼륨 | 
**KeyName** | Pointer to **NullableString** |  | [optional] 
**SecurityGroups** | [**[]CreateInstanceSecurityGroupModel**](CreateInstanceSecurityGroupModel.md) | 인스턴스에 적용할 보안 그룹 | 
**UserData** | Pointer to **NullableString** |  | [optional] 
**IsDisableHyperThreading** | Pointer to **bool** | 하이퍼스레딩 비활성화 여부 | [optional] [default to true]
**IsBonding** | Pointer to **bool** | 네트워크 인터페이스 본딩 설정 여부 &lt;br/&gt; - 베어메탈(&#x60;bm&#x60;) 유형을 생성할 경우에만 입력 필요 | [optional] 

## Methods

### NewCreateInstanceModel

`func NewCreateInstanceModel(name string, imageId string, flavorId string, subnets []CreateInstanceSubnetModel, volumes []CreateInstanceVolumeModel, securityGroups []CreateInstanceSecurityGroupModel, ) *CreateInstanceModel`

NewCreateInstanceModel instantiates a new CreateInstanceModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInstanceModelWithDefaults

`func NewCreateInstanceModelWithDefaults() *CreateInstanceModel`

NewCreateInstanceModelWithDefaults instantiates a new CreateInstanceModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateInstanceModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateInstanceModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateInstanceModel) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateInstanceModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateInstanceModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateInstanceModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateInstanceModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateInstanceModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateInstanceModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCount

`func (o *CreateInstanceModel) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CreateInstanceModel) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CreateInstanceModel) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *CreateInstanceModel) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetImageId

`func (o *CreateInstanceModel) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *CreateInstanceModel) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *CreateInstanceModel) SetImageId(v string)`

SetImageId sets ImageId field to given value.


### GetFlavorId

`func (o *CreateInstanceModel) GetFlavorId() string`

GetFlavorId returns the FlavorId field if non-nil, zero value otherwise.

### GetFlavorIdOk

`func (o *CreateInstanceModel) GetFlavorIdOk() (*string, bool)`

GetFlavorIdOk returns a tuple with the FlavorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorId

`func (o *CreateInstanceModel) SetFlavorId(v string)`

SetFlavorId sets FlavorId field to given value.


### GetAvailabilityZone

`func (o *CreateInstanceModel) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *CreateInstanceModel) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *CreateInstanceModel) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *CreateInstanceModel) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### SetAvailabilityZoneNil

`func (o *CreateInstanceModel) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *CreateInstanceModel) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
### GetSubnets

`func (o *CreateInstanceModel) GetSubnets() []CreateInstanceSubnetModel`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *CreateInstanceModel) GetSubnetsOk() (*[]CreateInstanceSubnetModel, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *CreateInstanceModel) SetSubnets(v []CreateInstanceSubnetModel)`

SetSubnets sets Subnets field to given value.


### GetVolumes

`func (o *CreateInstanceModel) GetVolumes() []CreateInstanceVolumeModel`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *CreateInstanceModel) GetVolumesOk() (*[]CreateInstanceVolumeModel, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *CreateInstanceModel) SetVolumes(v []CreateInstanceVolumeModel)`

SetVolumes sets Volumes field to given value.


### GetKeyName

`func (o *CreateInstanceModel) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *CreateInstanceModel) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *CreateInstanceModel) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.

### HasKeyName

`func (o *CreateInstanceModel) HasKeyName() bool`

HasKeyName returns a boolean if a field has been set.

### SetKeyNameNil

`func (o *CreateInstanceModel) SetKeyNameNil(b bool)`

 SetKeyNameNil sets the value for KeyName to be an explicit nil

### UnsetKeyName
`func (o *CreateInstanceModel) UnsetKeyName()`

UnsetKeyName ensures that no value is present for KeyName, not even an explicit nil
### GetSecurityGroups

`func (o *CreateInstanceModel) GetSecurityGroups() []CreateInstanceSecurityGroupModel`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *CreateInstanceModel) GetSecurityGroupsOk() (*[]CreateInstanceSecurityGroupModel, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *CreateInstanceModel) SetSecurityGroups(v []CreateInstanceSecurityGroupModel)`

SetSecurityGroups sets SecurityGroups field to given value.


### GetUserData

`func (o *CreateInstanceModel) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *CreateInstanceModel) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *CreateInstanceModel) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *CreateInstanceModel) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### SetUserDataNil

`func (o *CreateInstanceModel) SetUserDataNil(b bool)`

 SetUserDataNil sets the value for UserData to be an explicit nil

### UnsetUserData
`func (o *CreateInstanceModel) UnsetUserData()`

UnsetUserData ensures that no value is present for UserData, not even an explicit nil
### GetIsDisableHyperThreading

`func (o *CreateInstanceModel) GetIsDisableHyperThreading() bool`

GetIsDisableHyperThreading returns the IsDisableHyperThreading field if non-nil, zero value otherwise.

### GetIsDisableHyperThreadingOk

`func (o *CreateInstanceModel) GetIsDisableHyperThreadingOk() (*bool, bool)`

GetIsDisableHyperThreadingOk returns a tuple with the IsDisableHyperThreading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisableHyperThreading

`func (o *CreateInstanceModel) SetIsDisableHyperThreading(v bool)`

SetIsDisableHyperThreading sets IsDisableHyperThreading field to given value.

### HasIsDisableHyperThreading

`func (o *CreateInstanceModel) HasIsDisableHyperThreading() bool`

HasIsDisableHyperThreading returns a boolean if a field has been set.

### GetIsBonding

`func (o *CreateInstanceModel) GetIsBonding() bool`

GetIsBonding returns the IsBonding field if non-nil, zero value otherwise.

### GetIsBondingOk

`func (o *CreateInstanceModel) GetIsBondingOk() (*bool, bool)`

GetIsBondingOk returns a tuple with the IsBonding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBonding

`func (o *CreateInstanceModel) SetIsBonding(v bool)`

SetIsBonding sets IsBonding field to given value.

### HasIsBonding

`func (o *CreateInstanceModel) HasIsBonding() bool`

HasIsBonding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


