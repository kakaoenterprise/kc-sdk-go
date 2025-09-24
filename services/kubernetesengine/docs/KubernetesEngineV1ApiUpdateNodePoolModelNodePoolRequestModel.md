# KubernetesEngineV1ApiUpdateNodePoolModelNodePoolRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**SecurityGroups** | Pointer to **[]string** |  | [optional] 
**NodeCount** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewKubernetesEngineV1ApiUpdateNodePoolModelNodePoolRequestModel

`func NewKubernetesEngineV1ApiUpdateNodePoolModelNodePoolRequestModel() *KubernetesEngineV1ApiUpdateNodePoolModelNodePoolRequestModel`

NewKubernetesEngineV1ApiUpdateNodePoolModelNodePoolRequestModel instantiates a new KubernetesEngineV1ApiUpdateNodePoolModelNodePoolRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKubernetesEngineV1ApiUpdateNodePoolModelNodePoolRequestModelWithDefaults

`func NewKubernetesEngineV1ApiUpdateNodePoolModelNodePoolRequestModelWithDefaults() *KubernetesEngineV1ApiUpdateNodePoolModelNodePoolRequestModel`

NewKubernetesEngineV1ApiUpdateNodePoolModelNodePoolRequestModelWithDefaults instantiates a new KubernetesEngineV1ApiUpdateNodePoolModelNodePoolRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *KubernetesEngineV1ApiUpdateNodePoolModelNodePoolRequestModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KubernetesEngineV1ApiUpdateNodePoolModelNodePoolRequestModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KubernetesEngineV1ApiUpdateNodePoolModelNodePoolRequestModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KubernetesEngineV1ApiUpdateNodePoolModelNodePoolRequestModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *KubernetesEngineV1ApiUpdateNodePoolModelNodePoolRequestModel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *KubernetesEngineV1ApiUpdateNodePoolModelNodePoolRequestModel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSecurityGroups

`func (o *KubernetesEngineV1ApiUpdateNodePoolModelNodePoolRequestModel) GetSecurityGroups() []string`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *KubernetesEngineV1ApiUpdateNodePoolModelNodePoolRequestModel) GetSecurityGroupsOk() (*[]string, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *KubernetesEngineV1ApiUpdateNodePoolModelNodePoolRequestModel) SetSecurityGroups(v []string)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *KubernetesEngineV1ApiUpdateNodePoolModelNodePoolRequestModel) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroupsNil

`func (o *KubernetesEngineV1ApiUpdateNodePoolModelNodePoolRequestModel) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *KubernetesEngineV1ApiUpdateNodePoolModelNodePoolRequestModel) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil
### GetNodeCount

`func (o *KubernetesEngineV1ApiUpdateNodePoolModelNodePoolRequestModel) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *KubernetesEngineV1ApiUpdateNodePoolModelNodePoolRequestModel) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *KubernetesEngineV1ApiUpdateNodePoolModelNodePoolRequestModel) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *KubernetesEngineV1ApiUpdateNodePoolModelNodePoolRequestModel) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### SetNodeCountNil

`func (o *KubernetesEngineV1ApiUpdateNodePoolModelNodePoolRequestModel) SetNodeCountNil(b bool)`

 SetNodeCountNil sets the value for NodeCount to be an explicit nil

### UnsetNodeCount
`func (o *KubernetesEngineV1ApiUpdateNodePoolModelNodePoolRequestModel) UnsetNodeCount()`

UnsetNodeCount ensures that no value is present for NodeCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


