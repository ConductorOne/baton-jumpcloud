# PolicyTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activation** | Pointer to **string** | Requirements before the policy can be activated. | [optional] 
**Alert** | Pointer to **string** | Text to describe any risk associated with this policy. | [optional] 
**Behavior** | Pointer to **string** | Specifics about the behavior of the policy. | [optional] 
**DeliveryTypes** | Pointer to **[]string** | The supported delivery mechanisms for this policy template. | [optional] 
**Description** | Pointer to **string** | The default description for the Policy. | [optional] 
**DisplayName** | Pointer to **string** | The default display name for the Policy. | [optional] 
**Id** | Pointer to **string** | ObjectId uniquely identifying a Policy Template. | [optional] 
**Name** | Pointer to **string** | The unique name for the Policy Template. | [optional] 
**OsMetaFamily** | Pointer to **string** |  | [optional] 
**OsRestrictions** | Pointer to [**[]OSRestriction**](OSRestriction.md) |  | [optional] 
**Reference** | Pointer to **string** | URL to visit for further information. | [optional] 
**State** | Pointer to **string** | String describing the release status of the policy template. | [optional] [default to ""]

## Methods

### NewPolicyTemplate

`func NewPolicyTemplate() *PolicyTemplate`

NewPolicyTemplate instantiates a new PolicyTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyTemplateWithDefaults

`func NewPolicyTemplateWithDefaults() *PolicyTemplate`

NewPolicyTemplateWithDefaults instantiates a new PolicyTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivation

`func (o *PolicyTemplate) GetActivation() string`

GetActivation returns the Activation field if non-nil, zero value otherwise.

### GetActivationOk

`func (o *PolicyTemplate) GetActivationOk() (*string, bool)`

GetActivationOk returns a tuple with the Activation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivation

`func (o *PolicyTemplate) SetActivation(v string)`

SetActivation sets Activation field to given value.

### HasActivation

`func (o *PolicyTemplate) HasActivation() bool`

HasActivation returns a boolean if a field has been set.

### GetAlert

`func (o *PolicyTemplate) GetAlert() string`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *PolicyTemplate) GetAlertOk() (*string, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *PolicyTemplate) SetAlert(v string)`

SetAlert sets Alert field to given value.

### HasAlert

`func (o *PolicyTemplate) HasAlert() bool`

HasAlert returns a boolean if a field has been set.

### GetBehavior

`func (o *PolicyTemplate) GetBehavior() string`

GetBehavior returns the Behavior field if non-nil, zero value otherwise.

### GetBehaviorOk

`func (o *PolicyTemplate) GetBehaviorOk() (*string, bool)`

GetBehaviorOk returns a tuple with the Behavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehavior

`func (o *PolicyTemplate) SetBehavior(v string)`

SetBehavior sets Behavior field to given value.

### HasBehavior

`func (o *PolicyTemplate) HasBehavior() bool`

HasBehavior returns a boolean if a field has been set.

### GetDeliveryTypes

`func (o *PolicyTemplate) GetDeliveryTypes() []string`

GetDeliveryTypes returns the DeliveryTypes field if non-nil, zero value otherwise.

### GetDeliveryTypesOk

`func (o *PolicyTemplate) GetDeliveryTypesOk() (*[]string, bool)`

GetDeliveryTypesOk returns a tuple with the DeliveryTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryTypes

`func (o *PolicyTemplate) SetDeliveryTypes(v []string)`

SetDeliveryTypes sets DeliveryTypes field to given value.

### HasDeliveryTypes

`func (o *PolicyTemplate) HasDeliveryTypes() bool`

HasDeliveryTypes returns a boolean if a field has been set.

### GetDescription

`func (o *PolicyTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *PolicyTemplate) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PolicyTemplate) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PolicyTemplate) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PolicyTemplate) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetId

`func (o *PolicyTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PolicyTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PolicyTemplate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PolicyTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PolicyTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PolicyTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOsMetaFamily

`func (o *PolicyTemplate) GetOsMetaFamily() string`

GetOsMetaFamily returns the OsMetaFamily field if non-nil, zero value otherwise.

### GetOsMetaFamilyOk

`func (o *PolicyTemplate) GetOsMetaFamilyOk() (*string, bool)`

GetOsMetaFamilyOk returns a tuple with the OsMetaFamily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsMetaFamily

`func (o *PolicyTemplate) SetOsMetaFamily(v string)`

SetOsMetaFamily sets OsMetaFamily field to given value.

### HasOsMetaFamily

`func (o *PolicyTemplate) HasOsMetaFamily() bool`

HasOsMetaFamily returns a boolean if a field has been set.

### GetOsRestrictions

`func (o *PolicyTemplate) GetOsRestrictions() []OSRestriction`

GetOsRestrictions returns the OsRestrictions field if non-nil, zero value otherwise.

### GetOsRestrictionsOk

`func (o *PolicyTemplate) GetOsRestrictionsOk() (*[]OSRestriction, bool)`

GetOsRestrictionsOk returns a tuple with the OsRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsRestrictions

`func (o *PolicyTemplate) SetOsRestrictions(v []OSRestriction)`

SetOsRestrictions sets OsRestrictions field to given value.

### HasOsRestrictions

`func (o *PolicyTemplate) HasOsRestrictions() bool`

HasOsRestrictions returns a boolean if a field has been set.

### GetReference

`func (o *PolicyTemplate) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PolicyTemplate) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PolicyTemplate) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *PolicyTemplate) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetState

`func (o *PolicyTemplate) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PolicyTemplate) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PolicyTemplate) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *PolicyTemplate) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


