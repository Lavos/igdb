// Code generated by "stringer -type=CreditCategory"; DO NOT EDIT.

package igdb

import "strconv"

const _CreditCategory_name = "CreditVoiceActorCreditLanguageCreditCompanyCreditEmployeeCreditMiscCreditSupportCompany"

var _CreditCategory_index = [...]uint8{0, 16, 30, 43, 57, 67, 87}

func (i CreditCategory) String() string {
	i -= 1
	if i < 0 || i >= CreditCategory(len(_CreditCategory_index)-1) {
		return "CreditCategory(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _CreditCategory_name[_CreditCategory_index[i]:_CreditCategory_index[i+1]]
}
