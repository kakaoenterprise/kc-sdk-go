# GetK8sServiceAgentsResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceAgent** | [**ServiceAgentResponseModel**](ServiceAgentResponseModel.md) | 서비스 에이전트 정보 | 

## Methods

### NewGetK8sServiceAgentsResponseModel

`func NewGetK8sServiceAgentsResponseModel(serviceAgent ServiceAgentResponseModel, ) *GetK8sServiceAgentsResponseModel`

NewGetK8sServiceAgentsResponseModel instantiates a new GetK8sServiceAgentsResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetK8sServiceAgentsResponseModelWithDefaults

`func NewGetK8sServiceAgentsResponseModelWithDefaults() *GetK8sServiceAgentsResponseModel`

NewGetK8sServiceAgentsResponseModelWithDefaults instantiates a new GetK8sServiceAgentsResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceAgent

`func (o *GetK8sServiceAgentsResponseModel) GetServiceAgent() ServiceAgentResponseModel`

GetServiceAgent returns the ServiceAgent field if non-nil, zero value otherwise.

### GetServiceAgentOk

`func (o *GetK8sServiceAgentsResponseModel) GetServiceAgentOk() (*ServiceAgentResponseModel, bool)`

GetServiceAgentOk returns a tuple with the ServiceAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAgent

`func (o *GetK8sServiceAgentsResponseModel) SetServiceAgent(v ServiceAgentResponseModel)`

SetServiceAgent sets ServiceAgent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


