// Code generated by "stringer -type=AgeRatingCategory,AgeRatingEnum"; DO NOT EDIT.

package igdb

import "strconv"

const _AgeRatingCategory_name = "ESRBPEGI"

var _AgeRatingCategory_index = [...]uint8{0, 4, 8}

func (i AgeRatingCategory) String() string {
	i -= 1
	if i < 0 || i >= AgeRatingCategory(len(_AgeRatingCategory_index)-1) {
		return "AgeRatingCategory(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _AgeRatingCategory_name[_AgeRatingCategory_index[i]:_AgeRatingCategory_index[i+1]]
}

const _AgeRatingEnum_name = "ThreeSevenTwelveSixteenEighteenRPECEE10TMAO"

var _AgeRatingEnum_index = [...]uint8{0, 5, 10, 16, 23, 31, 33, 35, 36, 39, 40, 41, 43}

func (i AgeRatingEnum) String() string {
	i -= 1
	if i < 0 || i >= AgeRatingEnum(len(_AgeRatingEnum_index)-1) {
		return "AgeRatingEnum(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _AgeRatingEnum_name[_AgeRatingEnum_index[i]:_AgeRatingEnum_index[i+1]]
}
