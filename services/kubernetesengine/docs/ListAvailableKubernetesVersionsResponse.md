# ListAvailableKubernetesVersionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Versions** | [**[]UpgradableVersion**](UpgradableVersion.md) | Kubernetes 버전 정보 목록 | 

## Methods

### NewListAvailableKubernetesVersionsResponse

`func NewListAvailableKubernetesVersionsResponse(versions []UpgradableVersion, ) *ListAvailableKubernetesVersionsResponse`

NewListAvailableKubernetesVersionsResponse instantiates a new ListAvailableKubernetesVersionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAvailableKubernetesVersionsResponseWithDefaults

`func NewListAvailableKubernetesVersionsResponseWithDefaults() *ListAvailableKubernetesVersionsResponse`

NewListAvailableKubernetesVersionsResponseWithDefaults instantiates a new ListAvailableKubernetesVersionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersions

`func (o *ListAvailableKubernetesVersionsResponse) GetVersions() []UpgradableVersion`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *ListAvailableKubernetesVersionsResponse) GetVersionsOk() (*[]UpgradableVersion, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *ListAvailableKubernetesVersionsResponse) SetVersions(v []UpgradableVersion)`

SetVersions sets Versions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


