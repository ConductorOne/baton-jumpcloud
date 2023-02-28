# WorkdayOutputAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Basic** | Pointer to [**AuthInfo**](AuthInfo.md) |  | [optional] 
**Oauth** | Pointer to [**AuthInfo**](AuthInfo.md) |  | [optional] 

## Methods

### NewWorkdayOutputAuth

`func NewWorkdayOutputAuth() *WorkdayOutputAuth`

NewWorkdayOutputAuth instantiates a new WorkdayOutputAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkdayOutputAuthWithDefaults

`func NewWorkdayOutputAuthWithDefaults() *WorkdayOutputAuth`

NewWorkdayOutputAuthWithDefaults instantiates a new WorkdayOutputAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBasic

`func (o *WorkdayOutputAuth) GetBasic() AuthInfo`

GetBasic returns the Basic field if non-nil, zero value otherwise.

### GetBasicOk

`func (o *WorkdayOutputAuth) GetBasicOk() (*AuthInfo, bool)`

GetBasicOk returns a tuple with the Basic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasic

`func (o *WorkdayOutputAuth) SetBasic(v AuthInfo)`

SetBasic sets Basic field to given value.

### HasBasic

`func (o *WorkdayOutputAuth) HasBasic() bool`

HasBasic returns a boolean if a field has been set.

### GetOauth

`func (o *WorkdayOutputAuth) GetOauth() AuthInfo`

GetOauth returns the Oauth field if non-nil, zero value otherwise.

### GetOauthOk

`func (o *WorkdayOutputAuth) GetOauthOk() (*AuthInfo, bool)`

GetOauthOk returns a tuple with the Oauth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth

`func (o *WorkdayOutputAuth) SetOauth(v AuthInfo)`

SetOauth sets Oauth field to given value.

### HasOauth

`func (o *WorkdayOutputAuth) HasOauth() bool`

HasOauth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


