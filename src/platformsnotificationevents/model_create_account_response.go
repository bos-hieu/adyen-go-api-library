/*
 * Adyen for Platforms: Notifications
 *
 * The Notification API sends notifications to the endpoints specified in a given subscription. Subscriptions are managed through the Notification Configuration API. The API specifications listed here detail the format of each notification.  For more information, refer to our [documentation](https://docs.adyen.com/platforms/notifications).
 *
 * API version: 6
 * Contact: support@adyen.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package platformsnotificationevents

import (
	"encoding/json"
)

// CreateAccountResponse struct for CreateAccountResponse
type CreateAccountResponse struct {
	// The code of the new account.
	AccountCode string `json:"accountCode"`
	// The code of the account holder.
	AccountHolderCode string `json:"accountHolderCode"`
	// The bankAccountUUID of the bank account held by the account holder to couple the account with. Scheduled payouts in currencies matching the currency of this bank account will be sent to this bank account. Payouts in different currencies will be sent to a matching bank account of the account holder.
	BankAccountUUID *string `json:"bankAccountUUID,omitempty"`
	// The description of the account.
	Description *string `json:"description,omitempty"`
	// A list of fields that caused the `/createAccount` request to fail.
	InvalidFields *[]ErrorFieldType `json:"invalidFields,omitempty"`
	Metadata *map[string]string `json:"metadata,omitempty"`
	// The payout method code held by the account holder to couple the account with. Scheduled card payouts will be sent using this payout method code.
	PayoutMethodCode *string `json:"payoutMethodCode,omitempty"`
	PayoutSchedule *PayoutScheduleResponse `json:"payoutSchedule,omitempty"`
	// Speed with which payouts for this account are processed. Permitted values: `STANDARD`, `SAME_DAY`.
	PayoutSpeed *string `json:"payoutSpeed,omitempty"`
	// The reference of a request. Can be used to uniquely identify the request.
	PspReference *string `json:"pspReference,omitempty"`
	// The result code.
	ResultCode *string `json:"resultCode,omitempty"`
	// The status of the account. >Permitted values: `Active`.
	Status string `json:"status"`
}

// NewCreateAccountResponse instantiates a new CreateAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAccountResponse(accountCode string, accountHolderCode string, status string, ) *CreateAccountResponse {
	this := CreateAccountResponse{}
	this.AccountCode = accountCode
	this.AccountHolderCode = accountHolderCode
	this.Status = status
	return &this
}

// NewCreateAccountResponseWithDefaults instantiates a new CreateAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAccountResponseWithDefaults() *CreateAccountResponse {
	this := CreateAccountResponse{}
	return &this
}

// GetAccountCode returns the AccountCode field value
func (o *CreateAccountResponse) GetAccountCode() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.AccountCode
}

// GetAccountCodeOk returns a tuple with the AccountCode field value
// and a boolean to check if the value has been set.
func (o *CreateAccountResponse) GetAccountCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountCode, true
}

// SetAccountCode sets field value
func (o *CreateAccountResponse) SetAccountCode(v string) {
	o.AccountCode = v
}

// GetAccountHolderCode returns the AccountHolderCode field value
func (o *CreateAccountResponse) GetAccountHolderCode() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.AccountHolderCode
}

// GetAccountHolderCodeOk returns a tuple with the AccountHolderCode field value
// and a boolean to check if the value has been set.
func (o *CreateAccountResponse) GetAccountHolderCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountHolderCode, true
}

// SetAccountHolderCode sets field value
func (o *CreateAccountResponse) SetAccountHolderCode(v string) {
	o.AccountHolderCode = v
}

// GetBankAccountUUID returns the BankAccountUUID field value if set, zero value otherwise.
func (o *CreateAccountResponse) GetBankAccountUUID() string {
	if o == nil || o.BankAccountUUID == nil {
		var ret string
		return ret
	}
	return *o.BankAccountUUID
}

// GetBankAccountUUIDOk returns a tuple with the BankAccountUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountResponse) GetBankAccountUUIDOk() (*string, bool) {
	if o == nil || o.BankAccountUUID == nil {
		return nil, false
	}
	return o.BankAccountUUID, true
}

// HasBankAccountUUID returns a boolean if a field has been set.
func (o *CreateAccountResponse) HasBankAccountUUID() bool {
	if o != nil && o.BankAccountUUID != nil {
		return true
	}

	return false
}

// SetBankAccountUUID gets a reference to the given string and assigns it to the BankAccountUUID field.
func (o *CreateAccountResponse) SetBankAccountUUID(v string) {
	o.BankAccountUUID = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateAccountResponse) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateAccountResponse) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateAccountResponse) SetDescription(v string) {
	o.Description = &v
}

// GetInvalidFields returns the InvalidFields field value if set, zero value otherwise.
func (o *CreateAccountResponse) GetInvalidFields() []ErrorFieldType {
	if o == nil || o.InvalidFields == nil {
		var ret []ErrorFieldType
		return ret
	}
	return *o.InvalidFields
}

// GetInvalidFieldsOk returns a tuple with the InvalidFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountResponse) GetInvalidFieldsOk() (*[]ErrorFieldType, bool) {
	if o == nil || o.InvalidFields == nil {
		return nil, false
	}
	return o.InvalidFields, true
}

// HasInvalidFields returns a boolean if a field has been set.
func (o *CreateAccountResponse) HasInvalidFields() bool {
	if o != nil && o.InvalidFields != nil {
		return true
	}

	return false
}

// SetInvalidFields gets a reference to the given []ErrorFieldType and assigns it to the InvalidFields field.
func (o *CreateAccountResponse) SetInvalidFields(v []ErrorFieldType) {
	o.InvalidFields = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *CreateAccountResponse) GetMetadata() map[string]string {
	if o == nil || o.Metadata == nil {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountResponse) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *CreateAccountResponse) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *CreateAccountResponse) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetPayoutMethodCode returns the PayoutMethodCode field value if set, zero value otherwise.
func (o *CreateAccountResponse) GetPayoutMethodCode() string {
	if o == nil || o.PayoutMethodCode == nil {
		var ret string
		return ret
	}
	return *o.PayoutMethodCode
}

// GetPayoutMethodCodeOk returns a tuple with the PayoutMethodCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountResponse) GetPayoutMethodCodeOk() (*string, bool) {
	if o == nil || o.PayoutMethodCode == nil {
		return nil, false
	}
	return o.PayoutMethodCode, true
}

// HasPayoutMethodCode returns a boolean if a field has been set.
func (o *CreateAccountResponse) HasPayoutMethodCode() bool {
	if o != nil && o.PayoutMethodCode != nil {
		return true
	}

	return false
}

// SetPayoutMethodCode gets a reference to the given string and assigns it to the PayoutMethodCode field.
func (o *CreateAccountResponse) SetPayoutMethodCode(v string) {
	o.PayoutMethodCode = &v
}

// GetPayoutSchedule returns the PayoutSchedule field value if set, zero value otherwise.
func (o *CreateAccountResponse) GetPayoutSchedule() PayoutScheduleResponse {
	if o == nil || o.PayoutSchedule == nil {
		var ret PayoutScheduleResponse
		return ret
	}
	return *o.PayoutSchedule
}

// GetPayoutScheduleOk returns a tuple with the PayoutSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountResponse) GetPayoutScheduleOk() (*PayoutScheduleResponse, bool) {
	if o == nil || o.PayoutSchedule == nil {
		return nil, false
	}
	return o.PayoutSchedule, true
}

// HasPayoutSchedule returns a boolean if a field has been set.
func (o *CreateAccountResponse) HasPayoutSchedule() bool {
	if o != nil && o.PayoutSchedule != nil {
		return true
	}

	return false
}

// SetPayoutSchedule gets a reference to the given PayoutScheduleResponse and assigns it to the PayoutSchedule field.
func (o *CreateAccountResponse) SetPayoutSchedule(v PayoutScheduleResponse) {
	o.PayoutSchedule = &v
}

// GetPayoutSpeed returns the PayoutSpeed field value if set, zero value otherwise.
func (o *CreateAccountResponse) GetPayoutSpeed() string {
	if o == nil || o.PayoutSpeed == nil {
		var ret string
		return ret
	}
	return *o.PayoutSpeed
}

// GetPayoutSpeedOk returns a tuple with the PayoutSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountResponse) GetPayoutSpeedOk() (*string, bool) {
	if o == nil || o.PayoutSpeed == nil {
		return nil, false
	}
	return o.PayoutSpeed, true
}

// HasPayoutSpeed returns a boolean if a field has been set.
func (o *CreateAccountResponse) HasPayoutSpeed() bool {
	if o != nil && o.PayoutSpeed != nil {
		return true
	}

	return false
}

// SetPayoutSpeed gets a reference to the given string and assigns it to the PayoutSpeed field.
func (o *CreateAccountResponse) SetPayoutSpeed(v string) {
	o.PayoutSpeed = &v
}

// GetPspReference returns the PspReference field value if set, zero value otherwise.
func (o *CreateAccountResponse) GetPspReference() string {
	if o == nil || o.PspReference == nil {
		var ret string
		return ret
	}
	return *o.PspReference
}

// GetPspReferenceOk returns a tuple with the PspReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountResponse) GetPspReferenceOk() (*string, bool) {
	if o == nil || o.PspReference == nil {
		return nil, false
	}
	return o.PspReference, true
}

// HasPspReference returns a boolean if a field has been set.
func (o *CreateAccountResponse) HasPspReference() bool {
	if o != nil && o.PspReference != nil {
		return true
	}

	return false
}

// SetPspReference gets a reference to the given string and assigns it to the PspReference field.
func (o *CreateAccountResponse) SetPspReference(v string) {
	o.PspReference = &v
}

// GetResultCode returns the ResultCode field value if set, zero value otherwise.
func (o *CreateAccountResponse) GetResultCode() string {
	if o == nil || o.ResultCode == nil {
		var ret string
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccountResponse) GetResultCodeOk() (*string, bool) {
	if o == nil || o.ResultCode == nil {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *CreateAccountResponse) HasResultCode() bool {
	if o != nil && o.ResultCode != nil {
		return true
	}

	return false
}

// SetResultCode gets a reference to the given string and assigns it to the ResultCode field.
func (o *CreateAccountResponse) SetResultCode(v string) {
	o.ResultCode = &v
}

// GetStatus returns the Status field value
func (o *CreateAccountResponse) GetStatus() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CreateAccountResponse) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CreateAccountResponse) SetStatus(v string) {
	o.Status = v
}

func (o CreateAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accountCode"] = o.AccountCode
	}
	if true {
		toSerialize["accountHolderCode"] = o.AccountHolderCode
	}
	if o.BankAccountUUID != nil {
		toSerialize["bankAccountUUID"] = o.BankAccountUUID
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.InvalidFields != nil {
		toSerialize["invalidFields"] = o.InvalidFields
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.PayoutMethodCode != nil {
		toSerialize["payoutMethodCode"] = o.PayoutMethodCode
	}
	if o.PayoutSchedule != nil {
		toSerialize["payoutSchedule"] = o.PayoutSchedule
	}
	if o.PayoutSpeed != nil {
		toSerialize["payoutSpeed"] = o.PayoutSpeed
	}
	if o.PspReference != nil {
		toSerialize["pspReference"] = o.PspReference
	}
	if o.ResultCode != nil {
		toSerialize["resultCode"] = o.ResultCode
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableCreateAccountResponse struct {
	value *CreateAccountResponse
	isSet bool
}

func (v NullableCreateAccountResponse) Get() *CreateAccountResponse {
	return v.value
}

func (v *NullableCreateAccountResponse) Set(val *CreateAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAccountResponse(val *CreateAccountResponse) *NullableCreateAccountResponse {
	return &NullableCreateAccountResponse{value: val, isSet: true}
}

func (v NullableCreateAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

