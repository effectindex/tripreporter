// SPDX-FileCopyrightText: 2023 froggie <legal@frogg.ie>
//
// SPDX-License-Identifier: OSL-3.0

package models

type RouteOfAdministration int64

const ( // TODO: check if complete
	RoaUnknown RouteOfAdministration = iota
	RoaOther
	RoaOral
	RoaBuccal
	RoaRectal
	RoaInhaled
	RoaSublabial
	RoaIntranasal
	RoaSublingual
	RoaOtherInjection
	RoaBuccalInjection
	RoaIntravenousInjection
	RoaSubcutaneousInjection
	RoaIntramuscularInjection
)

type DosageUnit int64

const (
	DosageUnknown DosageUnit = iota
	DosageMicrograms
	DosageMilligrams
	DosageGrams
)

type DrugPrescribed int64

const (
	DrugPrescribedUnknown DrugPrescribed = iota
	DrugPrescribedOTC
	DrugPrescribedNo
	DrugPrescribedYes
)

type DisplayUnit int64

const (
	UnitUnknown DisplayUnit = iota
	UnitMetric
	UnitImperial
)

// TODO: Impl
//type UserPreferences struct { // TODO: How do we feel about this being unencrypted?
//	Timezone     time.Location `json:"timezone,omitempty"`      // Default: Europe/London
//	HeightFormat DisplayUnit   `json:"height_format,omitempty"` // Default: UnitMetric // Display height in centimeters or feet + inches
//	WeightFormat DisplayUnit   `json:"weight_format,omitempty"` // Default: UnitMetric // Display weight in kilograms or pounds
//}
