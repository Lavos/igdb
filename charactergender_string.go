// Code generated by "stringer -type=CharacterGender,CharacterSpecies"; DO NOT EDIT.

package igdb

import "strconv"

const _CharacterGender_name = "GenderMaleGenderFemaleGenderOther"

var _CharacterGender_index = [...]uint8{0, 10, 22, 33}

func (i CharacterGender) String() string {
	i -= 1
	if i < 0 || i >= CharacterGender(len(_CharacterGender_index)-1) {
		return "CharacterGender(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _CharacterGender_name[_CharacterGender_index[i]:_CharacterGender_index[i+1]]
}

const _CharacterSpecies_name = "SpeciesHumanSpeciesAlienSpeciesAnimalSpeciesAndroidSpeciesUnknown"

var _CharacterSpecies_index = [...]uint8{0, 12, 24, 37, 51, 65}

func (i CharacterSpecies) String() string {
	i -= 1
	if i < 0 || i >= CharacterSpecies(len(_CharacterSpecies_index)-1) {
		return "CharacterSpecies(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _CharacterSpecies_name[_CharacterSpecies_index[i]:_CharacterSpecies_index[i+1]]
}
