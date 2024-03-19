/*
Ory Identities API

This is the API specification for Ory Identities with features such as registration, login, recovery, account verification, profile settings, password reset, identity management, session management, email and sms delivery, and more. 

API version: v1.1.0
Contact: office@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// UpdateRegistrationFlowBody - Update Registration Request Body
type UpdateRegistrationFlowBody struct {
	UpdateRegistrationFlowWithCodeMethod *UpdateRegistrationFlowWithCodeMethod
	UpdateRegistrationFlowWithOidcMethod *UpdateRegistrationFlowWithOidcMethod
	UpdateRegistrationFlowWithPasswordMethod *UpdateRegistrationFlowWithPasswordMethod
	UpdateRegistrationFlowWithWebAuthnMethod *UpdateRegistrationFlowWithWebAuthnMethod
}

// UpdateRegistrationFlowWithCodeMethodAsUpdateRegistrationFlowBody is a convenience function that returns UpdateRegistrationFlowWithCodeMethod wrapped in UpdateRegistrationFlowBody
func UpdateRegistrationFlowWithCodeMethodAsUpdateRegistrationFlowBody(v *UpdateRegistrationFlowWithCodeMethod) UpdateRegistrationFlowBody {
	return UpdateRegistrationFlowBody{
		UpdateRegistrationFlowWithCodeMethod: v,
	}
}

// UpdateRegistrationFlowWithOidcMethodAsUpdateRegistrationFlowBody is a convenience function that returns UpdateRegistrationFlowWithOidcMethod wrapped in UpdateRegistrationFlowBody
func UpdateRegistrationFlowWithOidcMethodAsUpdateRegistrationFlowBody(v *UpdateRegistrationFlowWithOidcMethod) UpdateRegistrationFlowBody {
	return UpdateRegistrationFlowBody{
		UpdateRegistrationFlowWithOidcMethod: v,
	}
}

// UpdateRegistrationFlowWithPasswordMethodAsUpdateRegistrationFlowBody is a convenience function that returns UpdateRegistrationFlowWithPasswordMethod wrapped in UpdateRegistrationFlowBody
func UpdateRegistrationFlowWithPasswordMethodAsUpdateRegistrationFlowBody(v *UpdateRegistrationFlowWithPasswordMethod) UpdateRegistrationFlowBody {
	return UpdateRegistrationFlowBody{
		UpdateRegistrationFlowWithPasswordMethod: v,
	}
}

// UpdateRegistrationFlowWithWebAuthnMethodAsUpdateRegistrationFlowBody is a convenience function that returns UpdateRegistrationFlowWithWebAuthnMethod wrapped in UpdateRegistrationFlowBody
func UpdateRegistrationFlowWithWebAuthnMethodAsUpdateRegistrationFlowBody(v *UpdateRegistrationFlowWithWebAuthnMethod) UpdateRegistrationFlowBody {
	return UpdateRegistrationFlowBody{
		UpdateRegistrationFlowWithWebAuthnMethod: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *UpdateRegistrationFlowBody) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'code'
	if jsonDict["method"] == "code" {
		// try to unmarshal JSON data into UpdateRegistrationFlowWithCodeMethod
		err = json.Unmarshal(data, &dst.UpdateRegistrationFlowWithCodeMethod)
		if err == nil {
			return nil // data stored in dst.UpdateRegistrationFlowWithCodeMethod, return on the first match
		} else {
			dst.UpdateRegistrationFlowWithCodeMethod = nil
			return fmt.Errorf("failed to unmarshal UpdateRegistrationFlowBody as UpdateRegistrationFlowWithCodeMethod: %s", err.Error())
		}
	}

	// check if the discriminator value is 'oidc'
	if jsonDict["method"] == "oidc" {
		// try to unmarshal JSON data into UpdateRegistrationFlowWithOidcMethod
		err = json.Unmarshal(data, &dst.UpdateRegistrationFlowWithOidcMethod)
		if err == nil {
			return nil // data stored in dst.UpdateRegistrationFlowWithOidcMethod, return on the first match
		} else {
			dst.UpdateRegistrationFlowWithOidcMethod = nil
			return fmt.Errorf("failed to unmarshal UpdateRegistrationFlowBody as UpdateRegistrationFlowWithOidcMethod: %s", err.Error())
		}
	}

	// check if the discriminator value is 'password'
	if jsonDict["method"] == "password" {
		// try to unmarshal JSON data into UpdateRegistrationFlowWithPasswordMethod
		err = json.Unmarshal(data, &dst.UpdateRegistrationFlowWithPasswordMethod)
		if err == nil {
			return nil // data stored in dst.UpdateRegistrationFlowWithPasswordMethod, return on the first match
		} else {
			dst.UpdateRegistrationFlowWithPasswordMethod = nil
			return fmt.Errorf("failed to unmarshal UpdateRegistrationFlowBody as UpdateRegistrationFlowWithPasswordMethod: %s", err.Error())
		}
	}

	// check if the discriminator value is 'webauthn'
	if jsonDict["method"] == "webauthn" {
		// try to unmarshal JSON data into UpdateRegistrationFlowWithWebAuthnMethod
		err = json.Unmarshal(data, &dst.UpdateRegistrationFlowWithWebAuthnMethod)
		if err == nil {
			return nil // data stored in dst.UpdateRegistrationFlowWithWebAuthnMethod, return on the first match
		} else {
			dst.UpdateRegistrationFlowWithWebAuthnMethod = nil
			return fmt.Errorf("failed to unmarshal UpdateRegistrationFlowBody as UpdateRegistrationFlowWithWebAuthnMethod: %s", err.Error())
		}
	}

	// check if the discriminator value is 'updateRegistrationFlowWithCodeMethod'
	if jsonDict["method"] == "updateRegistrationFlowWithCodeMethod" {
		// try to unmarshal JSON data into UpdateRegistrationFlowWithCodeMethod
		err = json.Unmarshal(data, &dst.UpdateRegistrationFlowWithCodeMethod)
		if err == nil {
			return nil // data stored in dst.UpdateRegistrationFlowWithCodeMethod, return on the first match
		} else {
			dst.UpdateRegistrationFlowWithCodeMethod = nil
			return fmt.Errorf("failed to unmarshal UpdateRegistrationFlowBody as UpdateRegistrationFlowWithCodeMethod: %s", err.Error())
		}
	}

	// check if the discriminator value is 'updateRegistrationFlowWithOidcMethod'
	if jsonDict["method"] == "updateRegistrationFlowWithOidcMethod" {
		// try to unmarshal JSON data into UpdateRegistrationFlowWithOidcMethod
		err = json.Unmarshal(data, &dst.UpdateRegistrationFlowWithOidcMethod)
		if err == nil {
			return nil // data stored in dst.UpdateRegistrationFlowWithOidcMethod, return on the first match
		} else {
			dst.UpdateRegistrationFlowWithOidcMethod = nil
			return fmt.Errorf("failed to unmarshal UpdateRegistrationFlowBody as UpdateRegistrationFlowWithOidcMethod: %s", err.Error())
		}
	}

	// check if the discriminator value is 'updateRegistrationFlowWithPasswordMethod'
	if jsonDict["method"] == "updateRegistrationFlowWithPasswordMethod" {
		// try to unmarshal JSON data into UpdateRegistrationFlowWithPasswordMethod
		err = json.Unmarshal(data, &dst.UpdateRegistrationFlowWithPasswordMethod)
		if err == nil {
			return nil // data stored in dst.UpdateRegistrationFlowWithPasswordMethod, return on the first match
		} else {
			dst.UpdateRegistrationFlowWithPasswordMethod = nil
			return fmt.Errorf("failed to unmarshal UpdateRegistrationFlowBody as UpdateRegistrationFlowWithPasswordMethod: %s", err.Error())
		}
	}

	// check if the discriminator value is 'updateRegistrationFlowWithWebAuthnMethod'
	if jsonDict["method"] == "updateRegistrationFlowWithWebAuthnMethod" {
		// try to unmarshal JSON data into UpdateRegistrationFlowWithWebAuthnMethod
		err = json.Unmarshal(data, &dst.UpdateRegistrationFlowWithWebAuthnMethod)
		if err == nil {
			return nil // data stored in dst.UpdateRegistrationFlowWithWebAuthnMethod, return on the first match
		} else {
			dst.UpdateRegistrationFlowWithWebAuthnMethod = nil
			return fmt.Errorf("failed to unmarshal UpdateRegistrationFlowBody as UpdateRegistrationFlowWithWebAuthnMethod: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UpdateRegistrationFlowBody) MarshalJSON() ([]byte, error) {
	if src.UpdateRegistrationFlowWithCodeMethod != nil {
		return json.Marshal(&src.UpdateRegistrationFlowWithCodeMethod)
	}

	if src.UpdateRegistrationFlowWithOidcMethod != nil {
		return json.Marshal(&src.UpdateRegistrationFlowWithOidcMethod)
	}

	if src.UpdateRegistrationFlowWithPasswordMethod != nil {
		return json.Marshal(&src.UpdateRegistrationFlowWithPasswordMethod)
	}

	if src.UpdateRegistrationFlowWithWebAuthnMethod != nil {
		return json.Marshal(&src.UpdateRegistrationFlowWithWebAuthnMethod)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UpdateRegistrationFlowBody) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.UpdateRegistrationFlowWithCodeMethod != nil {
		return obj.UpdateRegistrationFlowWithCodeMethod
	}

	if obj.UpdateRegistrationFlowWithOidcMethod != nil {
		return obj.UpdateRegistrationFlowWithOidcMethod
	}

	if obj.UpdateRegistrationFlowWithPasswordMethod != nil {
		return obj.UpdateRegistrationFlowWithPasswordMethod
	}

	if obj.UpdateRegistrationFlowWithWebAuthnMethod != nil {
		return obj.UpdateRegistrationFlowWithWebAuthnMethod
	}

	// all schemas are nil
	return nil
}

type NullableUpdateRegistrationFlowBody struct {
	value *UpdateRegistrationFlowBody
	isSet bool
}

func (v NullableUpdateRegistrationFlowBody) Get() *UpdateRegistrationFlowBody {
	return v.value
}

func (v *NullableUpdateRegistrationFlowBody) Set(val *UpdateRegistrationFlowBody) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateRegistrationFlowBody) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateRegistrationFlowBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateRegistrationFlowBody(val *UpdateRegistrationFlowBody) *NullableUpdateRegistrationFlowBody {
	return &NullableUpdateRegistrationFlowBody{value: val, isSet: true}
}

func (v NullableUpdateRegistrationFlowBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateRegistrationFlowBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


