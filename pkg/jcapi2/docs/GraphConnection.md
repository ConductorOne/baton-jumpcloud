# GraphConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attributes** | Pointer to **map[string]interface{}** | The graph attributes. | [optional] 
**From** | Pointer to [**GraphObject**](GraphObject.md) |  | [optional] 
**To** | [**GraphObject**](GraphObject.md) |  | 

## Methods

### NewGraphConnection

`func NewGraphConnection(to GraphObject, ) *GraphConnection`

NewGraphConnection instantiates a new GraphConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphConnectionWithDefaults

`func NewGraphConnectionWithDefaults() *GraphConnection`

NewGraphConnectionWithDefaults instantiates a new GraphConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributes

`func (o *GraphConnection) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GraphConnection) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GraphConnection) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GraphConnection) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetFrom

`func (o *GraphConnection) GetFrom() GraphObject`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *GraphConnection) GetFromOk() (*GraphObject, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *GraphConnection) SetFrom(v GraphObject)`

SetFrom sets From field to given value.

### HasFrom

`func (o *GraphConnection) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *GraphConnection) GetTo() GraphObject`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *GraphConnection) GetToOk() (*GraphObject, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *GraphConnection) SetTo(v GraphObject)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


