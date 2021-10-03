package model

// EGender enum gender
type EGender int8

// enums
const (
	GenderUnknown EGender = iota // Unknown
	GenderFemale  EGender = iota // Female
	GenderMale    EGender = iota // Male
)
