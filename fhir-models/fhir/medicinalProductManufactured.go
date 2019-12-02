// Copyright 2019 The Samply Development Community
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fhir

import "encoding/json"

// THIS FILE IS GENERATED BY https://github.com/samply/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// MedicinalProductManufactured is documented here http://hl7.org/fhir/StructureDefinition/MedicinalProductManufactured
type MedicinalProductManufactured struct {
	Id                      *string             `json:"id,omitempty"`
	Meta                    *Meta               `json:"meta,omitempty"`
	ImplicitRules           *string             `json:"implicitRules,omitempty"`
	Language                *string             `json:"language,omitempty"`
	Text                    *Narrative          `json:"text,omitempty"`
	Extension               []Extension         `json:"extension,omitempty"`
	ModifierExtension       []Extension         `json:"modifierExtension,omitempty"`
	ManufacturedDoseForm    CodeableConcept     `json:"manufacturedDoseForm"`
	UnitOfPresentation      *CodeableConcept    `json:"unitOfPresentation,omitempty"`
	Quantity                Quantity            `json:"quantity"`
	Manufacturer            []Reference         `json:"manufacturer,omitempty"`
	Ingredient              []Reference         `json:"ingredient,omitempty"`
	PhysicalCharacteristics *ProdCharacteristic `json:"physicalCharacteristics,omitempty"`
	OtherCharacteristics    []CodeableConcept   `json:"otherCharacteristics,omitempty"`
}
type OtherMedicinalProductManufactured MedicinalProductManufactured

// MarshalJSON marshals the given MedicinalProductManufactured as JSON into a byte slice
func (r MedicinalProductManufactured) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicinalProductManufactured
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicinalProductManufactured: OtherMedicinalProductManufactured(r),
		ResourceType:                      "MedicinalProductManufactured",
	})
}

// UnmarshalMedicinalProductManufactured unmarshals a MedicinalProductManufactured.
func UnmarshalMedicinalProductManufactured(b []byte) (MedicinalProductManufactured, error) {
	var medicinalProductManufactured MedicinalProductManufactured
	if err := json.Unmarshal(b, &medicinalProductManufactured); err != nil {
		return medicinalProductManufactured, err
	}
	return medicinalProductManufactured, nil
}