# CreateInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | 생성할 인스턴스의 이름 - 같은 프로젝트 내 중복된 인스턴스 이름 사용 가능 | 
**Description** | Pointer to **NullableString** | 인스턴스에 대한 설명 | [optional] 
**Count** | Pointer to **int32** | 생성할 인스턴스 수 | [optional] [default to 1]
**ImageId** | **string** | 이미지의 고유 ID - [List images](/openapi/bcs/list-images)에서 확인 | 
**FlavorId** | **string** | 인스턴스 유형 ID - [List instance types](/openapi/bcs/list-instance-types)에서 확인 | 
**AvailabilityZone** | Pointer to [**NullableAvailabilityZone**](AvailabilityZone.md) | 인스턴스를 생성할 가용 영역 - 미 입력 시 입력한 &#x60;subnets[].id&#x60;에 맞는 가용 영역에 자동 생성 | [optional] 
**Subnets** | [**[]SubnetRequest**](SubnetRequest.md) | 연결할 서브넷 정보 목록 | 
**Volumes** | Pointer to [**[]BlockDeviceMappingRequest**](BlockDeviceMappingRequest.md) | 인스턴스에 연결할 볼륨 | [optional] 
**KeyName** | Pointer to **NullableString** | 인스턴스에 등록할 SSH 키 이름 - [List keypairs](/openapi/bcs/list-keypairs)에서 확인 | [optional] 
**SecurityGroups** | Pointer to [**[]SecurityGroupRequest**](SecurityGroupRequest.md) | 인스턴스에 적용할 보안 그룹 | [optional] 
**UserData** | Pointer to **NullableString** | 인스턴스 시작 시, 사용자 데이터를 실행하여 초기 환경 구성을 자동으로 실행하는 스크립트(인스턴스 최초 시작 시에만 실행됨) - 사용자 데이터 스크립트 및 &#x60;cloud-init&#x60; 명령을 16KB 이내로 입력 - 사용자 스크립트는 Base64 인코딩된 문자열로 입력 - 스크립트는 &#x60;/var/lib/cloud/instances&#x60; 하위에 복사되어 실행  - 스크립트 실행 로그 확인 시 다음 명령어 실행  ㄴ Ubuntu: &#x60;sudo cat /var/log/syslog 또는 sudo journalctl -u cloud-final.service&#x60;  ㄴ CentOS: &#x60;sudo cat /var/log/messages 또는 sudo journalctl -u cloud-final.service&#x60;    ⚠️ 잘못된 사용자 스크립트를 입력하거나 실행이 완료되지 않는 스크립트를 입력할 시, 부팅에 실패할 수 있습니다. | [optional] 
**IsDisableHyperThreading** | Pointer to **NullableBool** | 하이퍼스레딩 비활성화 여부 | [optional] 
**IsBonding** | Pointer to **NullableBool** | 네트워크 인터페이스 본딩 설정 여부 - 베어메탈(&#x60;bm&#x60;) 유형을 생성할 경우에만 입력 필요 | [optional] 

## Methods

### NewCreateInstance

`func NewCreateInstance(name string, imageId string, flavorId string, subnets []SubnetRequest, ) *CreateInstance`

NewCreateInstance instantiates a new CreateInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInstanceWithDefaults

`func NewCreateInstanceWithDefaults() *CreateInstance`

NewCreateInstanceWithDefaults instantiates a new CreateInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateInstance) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateInstance) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateInstance) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateInstance) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateInstance) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateInstance) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateInstance) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCount

`func (o *CreateInstance) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CreateInstance) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CreateInstance) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *CreateInstance) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetImageId

`func (o *CreateInstance) GetImageId() string`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *CreateInstance) GetImageIdOk() (*string, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *CreateInstance) SetImageId(v string)`

SetImageId sets ImageId field to given value.


### GetFlavorId

`func (o *CreateInstance) GetFlavorId() string`

GetFlavorId returns the FlavorId field if non-nil, zero value otherwise.

### GetFlavorIdOk

`func (o *CreateInstance) GetFlavorIdOk() (*string, bool)`

GetFlavorIdOk returns a tuple with the FlavorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorId

`func (o *CreateInstance) SetFlavorId(v string)`

SetFlavorId sets FlavorId field to given value.


### GetAvailabilityZone

`func (o *CreateInstance) GetAvailabilityZone() AvailabilityZone`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *CreateInstance) GetAvailabilityZoneOk() (*AvailabilityZone, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *CreateInstance) SetAvailabilityZone(v AvailabilityZone)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *CreateInstance) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### SetAvailabilityZoneNil

`func (o *CreateInstance) SetAvailabilityZoneNil(b bool)`

 SetAvailabilityZoneNil sets the value for AvailabilityZone to be an explicit nil

### UnsetAvailabilityZone
`func (o *CreateInstance) UnsetAvailabilityZone()`

UnsetAvailabilityZone ensures that no value is present for AvailabilityZone, not even an explicit nil
### GetSubnets

`func (o *CreateInstance) GetSubnets() []SubnetRequest`

GetSubnets returns the Subnets field if non-nil, zero value otherwise.

### GetSubnetsOk

`func (o *CreateInstance) GetSubnetsOk() (*[]SubnetRequest, bool)`

GetSubnetsOk returns a tuple with the Subnets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnets

`func (o *CreateInstance) SetSubnets(v []SubnetRequest)`

SetSubnets sets Subnets field to given value.


### GetVolumes

`func (o *CreateInstance) GetVolumes() []BlockDeviceMappingRequest`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *CreateInstance) GetVolumesOk() (*[]BlockDeviceMappingRequest, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *CreateInstance) SetVolumes(v []BlockDeviceMappingRequest)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *CreateInstance) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### SetVolumesNil

`func (o *CreateInstance) SetVolumesNil(b bool)`

 SetVolumesNil sets the value for Volumes to be an explicit nil

### UnsetVolumes
`func (o *CreateInstance) UnsetVolumes()`

UnsetVolumes ensures that no value is present for Volumes, not even an explicit nil
### GetKeyName

`func (o *CreateInstance) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *CreateInstance) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *CreateInstance) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.

### HasKeyName

`func (o *CreateInstance) HasKeyName() bool`

HasKeyName returns a boolean if a field has been set.

### SetKeyNameNil

`func (o *CreateInstance) SetKeyNameNil(b bool)`

 SetKeyNameNil sets the value for KeyName to be an explicit nil

### UnsetKeyName
`func (o *CreateInstance) UnsetKeyName()`

UnsetKeyName ensures that no value is present for KeyName, not even an explicit nil
### GetSecurityGroups

`func (o *CreateInstance) GetSecurityGroups() []SecurityGroupRequest`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *CreateInstance) GetSecurityGroupsOk() (*[]SecurityGroupRequest, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *CreateInstance) SetSecurityGroups(v []SecurityGroupRequest)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *CreateInstance) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroupsNil

`func (o *CreateInstance) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *CreateInstance) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil
### GetUserData

`func (o *CreateInstance) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *CreateInstance) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *CreateInstance) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *CreateInstance) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### SetUserDataNil

`func (o *CreateInstance) SetUserDataNil(b bool)`

 SetUserDataNil sets the value for UserData to be an explicit nil

### UnsetUserData
`func (o *CreateInstance) UnsetUserData()`

UnsetUserData ensures that no value is present for UserData, not even an explicit nil
### GetIsDisableHyperThreading

`func (o *CreateInstance) GetIsDisableHyperThreading() bool`

GetIsDisableHyperThreading returns the IsDisableHyperThreading field if non-nil, zero value otherwise.

### GetIsDisableHyperThreadingOk

`func (o *CreateInstance) GetIsDisableHyperThreadingOk() (*bool, bool)`

GetIsDisableHyperThreadingOk returns a tuple with the IsDisableHyperThreading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisableHyperThreading

`func (o *CreateInstance) SetIsDisableHyperThreading(v bool)`

SetIsDisableHyperThreading sets IsDisableHyperThreading field to given value.

### HasIsDisableHyperThreading

`func (o *CreateInstance) HasIsDisableHyperThreading() bool`

HasIsDisableHyperThreading returns a boolean if a field has been set.

### SetIsDisableHyperThreadingNil

`func (o *CreateInstance) SetIsDisableHyperThreadingNil(b bool)`

 SetIsDisableHyperThreadingNil sets the value for IsDisableHyperThreading to be an explicit nil

### UnsetIsDisableHyperThreading
`func (o *CreateInstance) UnsetIsDisableHyperThreading()`

UnsetIsDisableHyperThreading ensures that no value is present for IsDisableHyperThreading, not even an explicit nil
### GetIsBonding

`func (o *CreateInstance) GetIsBonding() bool`

GetIsBonding returns the IsBonding field if non-nil, zero value otherwise.

### GetIsBondingOk

`func (o *CreateInstance) GetIsBondingOk() (*bool, bool)`

GetIsBondingOk returns a tuple with the IsBonding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBonding

`func (o *CreateInstance) SetIsBonding(v bool)`

SetIsBonding sets IsBonding field to given value.

### HasIsBonding

`func (o *CreateInstance) HasIsBonding() bool`

HasIsBonding returns a boolean if a field has been set.

### SetIsBondingNil

`func (o *CreateInstance) SetIsBondingNil(b bool)`

 SetIsBondingNil sets the value for IsBonding to be an explicit nil

### UnsetIsBonding
`func (o *CreateInstance) UnsetIsBonding()`

UnsetIsBonding ensures that no value is present for IsBonding, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


