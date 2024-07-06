package ui

import (
	"strconv"
	"time"
	"unicode/utf8"

	"fyne.io/fyne/v2/widget"
	"github.com/LobovVit/GophKeeper/internal/client/constants"
	"github.com/LobovVit/GophKeeper/pkg/utils"
)

const userNameMaxLength = 6

func validateLoginForm(usernameLoginEntry *widget.Entry, passwordLoginEntry *widget.Entry) (string, bool) {
	if utf8.RuneCountInString(usernameLoginEntry.Text) < userNameMaxLength {
		return constants.ErrUsernameIncorrect, false
	}
	if !utils.Verify(passwordLoginEntry.Text) {
		return constants.ErrPasswordIncorrect, false
	}
	return "", true
}

func validateRegistrationForm(usernameRegistrationEntry *widget.Entry, passwordRegistrationEntry *widget.Entry,
	passwordConfirmationRegistrationEntry *widget.Entry) (string, bool) {
	if utf8.RuneCountInString(usernameRegistrationEntry.Text) < userNameMaxLength {
		return constants.ErrUsernameIncorrect, false
	}
	if !utils.Verify(passwordRegistrationEntry.Text) {
		return constants.ErrPasswordIncorrect, false
	}
	if passwordRegistrationEntry.Text != passwordConfirmationRegistrationEntry.Text {
		return constants.ErrPasswordDifferent, false
	}
	return "", true
}

func validateLoginPasswordForm(loginPasswordNameEntry *widget.Entry, loginPasswordDescriptionEntry *widget.Entry,
	loginEntry *widget.Entry, passwordEntry *widget.Entry) (string, bool) {
	if loginPasswordNameEntry.Text == "" {
		return constants.ErrNameEmpty, false
	}
	if loginPasswordDescriptionEntry.Text == "" {
		return constants.ErrDescriptionEmpty, false
	}
	if loginEntry.Text == "" {
		return constants.ErrLoginEmpty, false
	}
	if passwordEntry.Text == "" {
		return constants.ErrPasswordEmpty, false
	}
	return "", true
}

func validateTextForm(textNameEntry *widget.Entry, textDescriptionEntry *widget.Entry, textEntry *widget.Entry) (string, bool) {
	if textNameEntry.Text == "" {
		return constants.ErrNameEmpty, false
	}
	if textDescriptionEntry.Text == "" {
		return constants.ErrDescriptionEmpty, false
	}
	if textEntry.Text == "" {
		return constants.ErrTextEmpty, false
	}
	return "", true
}

func validateCardForm(cardNameEntry *widget.Entry, cardDescriptionEntry *widget.Entry, paymentSystemEntry *widget.Entry,
	numberEntry *widget.Entry, holderEntry *widget.Entry, cvcEntry *widget.Entry, endDateEntry *widget.Entry) (string, bool) {
	var err error
	if cardNameEntry.Text == "" {
		return constants.ErrNameEmpty, false
	}
	if cardDescriptionEntry.Text == "" {
		return constants.ErrDescriptionEmpty, false
	}
	if paymentSystemEntry.Text == "" {
		return constants.ErrPaymentSystemEmpty, false
	}
	if numberEntry.Text == "" {
		return constants.ErrNumberEmpty, false
	}
	intNumber, err := strconv.Atoi(numberEntry.Text)
	if err != nil {
		return constants.ErrNumberIncorrect, false
	}
	if !utils.ValidLuhn(intNumber) {
		return constants.ErrNumberIncorrect, false
	}
	if holderEntry.Text == "" {
		return constants.ErrHolderEmpty, false
	}
	if endDateEntry.Text == "" {
		return constants.ErrEndDateEmpty, false
	} else {
		_, err = time.Parse(constants.LayoutDate.ToString(), endDateEntry.Text)
		if err != nil {
			return constants.ErrEndDateIncorrect, false
		}
	}
	if cvcEntry.Text == "" {
		return constants.ErrCvcEmpty, false
	} else {
		_, err = strconv.Atoi(cvcEntry.Text)
		if err != nil {
			return constants.ErrCvcIncorrect, false
		}
	}
	return "", true
}

func searchByColumn(slice [][]string, targetColumn int, targetValue string) bool {
	for i := 1; i < len(slice) && len(slice) > 1; i++ {
		if slice[i][targetColumn] == targetValue {
			return true
		}
	}
	return false
}

func removeRow(slice [][]string, indexRow int) [][]string {
	return append(slice[:indexRow], slice[indexRow+1:]...)
}
