/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests. 
 *
 * API version: v0.6.0-alpha.16
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// SubmitSelfServiceVerificationFlowWithLinkMethod nolint:deadcode,unused
type SubmitSelfServiceVerificationFlowWithLinkMethod struct {
	// Sending the anti-csrf token is only required for browser login flows.
	CsrfToken *string `json:"csrf_token,omitempty"`
	// Email to Verify  Needs to be set when initiating the flow. If the email is a registered verification email, a verification link will be sent. If the email is not known, a email with details on what happened will be sent instead.  format: email in: body
	Email *string `json:"email,omitempty"`
}

// NewSubmitSelfServiceVerificationFlowWithLinkMethod instantiates a new SubmitSelfServiceVerificationFlowWithLinkMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitSelfServiceVerificationFlowWithLinkMethod() *SubmitSelfServiceVerificationFlowWithLinkMethod {
	this := SubmitSelfServiceVerificationFlowWithLinkMethod{}
	return &this
}

// NewSubmitSelfServiceVerificationFlowWithLinkMethodWithDefaults instantiates a new SubmitSelfServiceVerificationFlowWithLinkMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitSelfServiceVerificationFlowWithLinkMethodWithDefaults() *SubmitSelfServiceVerificationFlowWithLinkMethod {
	this := SubmitSelfServiceVerificationFlowWithLinkMethod{}
	return &this
}

// GetCsrfToken returns the CsrfToken field value if set, zero value otherwise.
func (o *SubmitSelfServiceVerificationFlowWithLinkMethod) GetCsrfToken() string {
	if o == nil || o.CsrfToken == nil {
		var ret string
		return ret
	}
	return *o.CsrfToken
}

// GetCsrfTokenOk returns a tuple with the CsrfToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceVerificationFlowWithLinkMethod) GetCsrfTokenOk() (*string, bool) {
	if o == nil || o.CsrfToken == nil {
		return nil, false
	}
	return o.CsrfToken, true
}

// HasCsrfToken returns a boolean if a field has been set.
func (o *SubmitSelfServiceVerificationFlowWithLinkMethod) HasCsrfToken() bool {
	if o != nil && o.CsrfToken != nil {
		return true
	}

	return false
}

// SetCsrfToken gets a reference to the given string and assigns it to the CsrfToken field.
func (o *SubmitSelfServiceVerificationFlowWithLinkMethod) SetCsrfToken(v string) {
	o.CsrfToken = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *SubmitSelfServiceVerificationFlowWithLinkMethod) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitSelfServiceVerificationFlowWithLinkMethod) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *SubmitSelfServiceVerificationFlowWithLinkMethod) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *SubmitSelfServiceVerificationFlowWithLinkMethod) SetEmail(v string) {
	o.Email = &v
}

func (o SubmitSelfServiceVerificationFlowWithLinkMethod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CsrfToken != nil {
		toSerialize["csrf_token"] = o.CsrfToken
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	return json.Marshal(toSerialize)
}

type NullableSubmitSelfServiceVerificationFlowWithLinkMethod struct {
	value *SubmitSelfServiceVerificationFlowWithLinkMethod
	isSet bool
}

func (v NullableSubmitSelfServiceVerificationFlowWithLinkMethod) Get() *SubmitSelfServiceVerificationFlowWithLinkMethod {
	return v.value
}

func (v *NullableSubmitSelfServiceVerificationFlowWithLinkMethod) Set(val *SubmitSelfServiceVerificationFlowWithLinkMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitSelfServiceVerificationFlowWithLinkMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitSelfServiceVerificationFlowWithLinkMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitSelfServiceVerificationFlowWithLinkMethod(val *SubmitSelfServiceVerificationFlowWithLinkMethod) *NullableSubmitSelfServiceVerificationFlowWithLinkMethod {
	return &NullableSubmitSelfServiceVerificationFlowWithLinkMethod{value: val, isSet: true}
}

func (v NullableSubmitSelfServiceVerificationFlowWithLinkMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubmitSelfServiceVerificationFlowWithLinkMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


