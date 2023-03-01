# MemberSuggestion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Object** | Pointer to [**GraphObject**](GraphObject.md) |  | [optional] 
**Op** | Pointer to **string** | How to modify group membership. | [optional] 

## Methods

### NewMemberSuggestion

`func NewMemberSuggestion() *MemberSuggestion`

NewMemberSuggestion instantiates a new MemberSuggestion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMemberSuggestionWithDefaults

`func NewMemberSuggestionWithDefaults() *MemberSuggestion`

NewMemberSuggestionWithDefaults instantiates a new MemberSuggestion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObject

`func (o *MemberSuggestion) GetObject() GraphObject`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *MemberSuggestion) GetObjectOk() (*GraphObject, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *MemberSuggestion) SetObject(v GraphObject)`

SetObject sets Object field to given value.

### HasObject

`func (o *MemberSuggestion) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetOp

`func (o *MemberSuggestion) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *MemberSuggestion) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *MemberSuggestion) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *MemberSuggestion) HasOp() bool`

HasOp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


