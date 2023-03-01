# PolicyTemplateWithDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activation** | Pointer to **string** | Requirements before the policy can be activated. | [optional] 
**Behavior** | Pointer to **string** | Specifics about the behavior of the policy. | [optional] 
**ConfigFields** | Pointer to [**[]PolicyTemplateConfigField**](PolicyTemplateConfigField.md) | An unordered list of all the fields that can be configured for this Policy Template. | [optional] 
**Description** | Pointer to **string** | The default description for the Policy. | [optional] 
**DisplayName** | Pointer to **string** | The default display name for the Policy. | [optional] 
**Id** | Pointer to **string** | ObjectId uniquely identifying a Policy Template. | [optional] 
**Name** | Pointer to **string** | The unique name for the Policy Template. | [optional] 
**OsMetaFamily** | Pointer to **string** |  | [optional] 
**OsRestrictions** | Pointer to [**[]OSRestriction**](OSRestriction.md) |  | [optional] 

## Methods

### NewPolicyTemplateWithDetails

`func NewPolicyTemplateWithDetails() *PolicyTemplateWithDetails`

NewPolicyTemplateWithDetails instantiates a new PolicyTemplateWithDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyTemplateWithDetailsWithDefaults

`func NewPolicyTemplateWithDetailsWithDefaults() *PolicyTemplateWithDetails`

NewPolicyTemplateWithDetailsWithDefaults instantiates a new PolicyTemplateWithDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivation

`func (o *PolicyTemplateWithDetails) GetActivation() string`

GetActivation returns the Activation field if non-nil, zero value otherwise.

### GetActivationOk

`func (o *PolicyTemplateWithDetails) GetActivationOk() (*string, bool)`

GetActivationOk returns a tuple with the Activation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivation

`func (o *PolicyTemplateWithDetails) SetActivation(v string)`

SetActivation sets Activation field to given value.

### HasActivation

`func (o *PolicyTemplateWithDetails) HasActivation() bool`

HasActivation returns a boolean if a field has been set.

### GetBehavior

`func (o *PolicyTemplateWithDetails) GetBehavior() string`

GetBehavior returns the Behavior field if non-nil, zero value otherwise.

### GetBehaviorOk

`func (o *PolicyTemplateWithDetails) GetBehaviorOk() (*string, bool)`

GetBehaviorOk returns a tuple with the Behavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehavior

`func (o *PolicyTemplateWithDetails) SetBehavior(v string)`

SetBehavior sets Behavior field to given value.

### HasBehavior

`func (o *PolicyTemplateWithDetails) HasBehavior() bool`

HasBehavior returns a boolean if a field has been set.

### GetConfigFields

`func (o *PolicyTemplateWithDetails) GetConfigFields() []PolicyTemplateConfigField`

GetConfigFields returns the ConfigFields field if non-nil, zero value otherwise.

### GetConfigFieldsOk

`func (o *PolicyTemplateWithDetails) GetConfigFieldsOk() (*[]PolicyTemplateConfigField, bool)`

GetConfigFieldsOk returns a tuple with the ConfigFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigFields

`func (o *PolicyTemplateWithDetails) SetConfigFields(v []PolicyTemplateConfigField)`

SetConfigFields sets ConfigFields field to given value.

### HasConfigFields

`func (o *PolicyTemplateWithDetails) HasConfigFields() bool`

HasConfigFields returns a boolean if a field has been set.

### GetDescription

`func (o *PolicyTemplateWithDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyTemplateWithDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyTemplateWithDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyTemplateWithDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *PolicyTemplateWithDetails) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PolicyTemplateWithDetails) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PolicyTemplateWithDetails) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PolicyTemplateWithDetails) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetId

`func (o *PolicyTemplateWithDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PolicyTemplateWithDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PolicyTemplateWithDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PolicyTemplateWithDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PolicyTemplateWithDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyTemplateWithDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyTemplateWithDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PolicyTemplateWithDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOsMetaFamily

`func (o *PolicyTemplateWithDetails) GetOsMetaFamily() string`

GetOsMetaFamily returns the OsMetaFamily field if non-nil, zero value otherwise.

### GetOsMetaFamilyOk

`func (o *PolicyTemplateWithDetails) GetOsMetaFamilyOk() (*string, bool)`

GetOsMetaFamilyOk returns a tuple with the OsMetaFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsMetaFamily

`func (o *PolicyTemplateWithDetails) SetOsMetaFamily(v string)`

SetOsMetaFamily sets OsMetaFamily field to given value.

### HasOsMetaFamily

`func (o *PolicyTemplateWithDetails) HasOsMetaFamily() bool`

HasOsMetaFamily returns a boolean if a field has been set.

### GetOsRestrictions

`func (o *PolicyTemplateWithDetails) GetOsRestrictions() []OSRestriction`

GetOsRestrictions returns the OsRestrictions field if non-nil, zero value otherwise.

### GetOsRestrictionsOk

`func (o *PolicyTemplateWithDetails) GetOsRestrictionsOk() (*[]OSRestriction, bool)`

GetOsRestrictionsOk returns a tuple with the OsRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsRestrictions

`func (o *PolicyTemplateWithDetails) SetOsRestrictions(v []OSRestriction)`

SetOsRestrictions sets OsRestrictions field to given value.

### HasOsRestrictions

`func (o *PolicyTemplateWithDetails) HasOsRestrictions() bool`

HasOsRestrictions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


