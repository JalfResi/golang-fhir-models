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

// MedicinalProductContraindication is documented here http://hl7.org/fhir/StructureDefinition/MedicinalProductContraindication
type MedicinalProductContraindication struct {
	Id                    *string                                        `json:"id,omitempty"`
	Meta                  *Meta                                          `json:"meta,omitempty"`
	ImplicitRules         *string                                        `json:"implicitRules,omitempty"`
	Language              *string                                        `json:"language,omitempty"`
	Text                  *Narrative                                     `json:"text,omitempty"`
	Extension             []Extension                                    `json:"extension,omitempty"`
	ModifierExtension     []Extension                                    `json:"modifierExtension,omitempty"`
	Subject               []Reference                                    `json:"subject,omitempty"`
	Disease               *CodeableConcept                               `json:"disease,omitempty"`
	DiseaseStatus         *CodeableConcept                               `json:"diseaseStatus,omitempty"`
	Comorbidity           []CodeableConcept                              `json:"comorbidity,omitempty"`
	TherapeuticIndication []Reference                                    `json:"therapeuticIndication,omitempty"`
	OtherTherapy          []MedicinalProductContraindicationOtherTherapy `json:"otherTherapy,omitempty"`
	Population            []Population                                   `json:"population,omitempty"`
}
type MedicinalProductContraindicationOtherTherapy struct {
	Id                      *string         `json:"id,omitempty"`
	Extension               []Extension     `json:"extension,omitempty"`
	ModifierExtension       []Extension     `json:"modifierExtension,omitempty"`
	TherapyRelationshipType CodeableConcept `json:"therapyRelationshipType"`
}
type OtherMedicinalProductContraindication MedicinalProductContraindication

// MarshalJSON marshals the given MedicinalProductContraindication as JSON into a byte slice
func (r MedicinalProductContraindication) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicinalProductContraindication
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicinalProductContraindication: OtherMedicinalProductContraindication(r),
		ResourceType:                          "MedicinalProductContraindication",
	})
}

// UnmarshalMedicinalProductContraindication unmarshals a MedicinalProductContraindication.
func UnmarshalMedicinalProductContraindication(b []byte) (MedicinalProductContraindication, error) {
	var medicinalProductContraindication MedicinalProductContraindication
	if err := json.Unmarshal(b, &medicinalProductContraindication); err != nil {
		return medicinalProductContraindication, err
	}
	return medicinalProductContraindication, nil
}