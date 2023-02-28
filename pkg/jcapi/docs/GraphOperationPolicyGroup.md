# GraphOperationPolicyGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ObjectID of graph object being added or removed as an association. | 
**Op** | **string** | How to modify the graph connection. | 
**Attributes** | Pointer to **map[string]interface{}** | The graph attributes. | [optional] 
**Type** | **string** | Targets which a \&quot;policy_group\&quot; can be associated to. | 

## Methods

### NewGraphOperationPolicyGroup

`func NewGraphOperationPolicyGroup(id string, op string, type_ string, ) *GraphOperationPolicyGroup`

NewGraphOperationPolicyGroup instantiates a new GraphOperationPolicyGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphOperationPolicyGroupWithDefaults

`func NewGraphOperationPolicyGroupWithDefaults() *GraphOperationPolicyGroup`

NewGraphOperationPolicyGroupWithDefaults instantiates a new GraphOperationPolicyGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GraphOperationPolicyGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GraphOperationPolicyGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GraphOperationPolicyGroup) SetId(v string)`

SetId sets Id field to given value.


### GetOp

`func (o *GraphOperationPolicyGroup) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *GraphOperationPolicyGroup) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *GraphOperationPolicyGroup) SetOp(v string)`

SetOp sets Op field to given value.


### GetAttributes

`func (o *GraphOperationPolicyGroup) GetAttributes() map[string]interface{}`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GraphOperationPolicyGroup) GetAttributesOk() (*map[string]interface{}, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GraphOperationPolicyGroup) SetAttributes(v map[string]interface{})`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GraphOperationPolicyGroup) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetType

`func (o *GraphOperationPolicyGroup) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GraphOperationPolicyGroup) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GraphOperationPolicyGroup) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


