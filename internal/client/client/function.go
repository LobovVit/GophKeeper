package client

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/LobovVit/GophKeeper/internal/client/constants"
	"github.com/LobovVit/GophKeeper/internal/client/model"
	grpc "github.com/LobovVit/GophKeeper/internal/proto"
	"github.com/LobovVit/GophKeeper/pkg/utils"
)

func SearchByColumn(slice [][]string, targetColumn int, targetValue string) bool {
	for i := 1; i < len(slice) && len(slice) > 1; i++ {
		if slice[i][targetColumn] == targetValue {
			return true
		}
	}
	return false
}

func RemoveRow(slice [][]string, indexRow int) [][]string {
	return append(slice[:indexRow], slice[indexRow+1:]...)
}

func AppendTextEntity(node *grpc.Entity, dataTblText *[][]string, plaintext string) error {
	created, err := utils.ConvertTimestampToTime(node.CreatedAt)
	if err != nil {
		return err
	}
	updated, err := utils.ConvertTimestampToTime(node.UpdatedAt)
	if err != nil {
		return err
	}
	var metadata model.MetadataEntity
	err = json.Unmarshal([]byte(node.Metadata), &metadata)
	if err != nil {
		return err
	}
	row := []string{metadata.Name, metadata.Description, plaintext, created.Format(constants.LayoutDateAndTime.ToString()), updated.Format(constants.LayoutDateAndTime.ToString())}
	*dataTblText = append(*dataTblText, row)
	return nil
}

func AppendLoginPasswordEntity(node *grpc.Entity, dataTblLoginPassword *[][]string, jsonLoginPassword model.LoginPassword) error {
	created, err := utils.ConvertTimestampToTime(node.CreatedAt)
	if err != nil {
		return err
	}
	updated, err := utils.ConvertTimestampToTime(node.UpdatedAt)
	if err != nil {
		return err
	}
	var metadata model.MetadataEntity
	err = json.Unmarshal([]byte(node.Metadata), &metadata)
	if err != nil {
		return err
	}
	row := []string{metadata.Name, metadata.Description, jsonLoginPassword.Login, jsonLoginPassword.Password,
		created.Format(constants.LayoutDateAndTime.ToString()), updated.Format(constants.LayoutDateAndTime.ToString())}
	*dataTblLoginPassword = append(*dataTblLoginPassword, row)
	return nil
}

func AppendCardEntity(node *grpc.Entity, dataTblCard *[][]string, jsonCard model.Card) error {
	created, err := utils.ConvertTimestampToTime(node.CreatedAt)
	if err != nil {
		return err
	}
	updated, err := utils.ConvertTimestampToTime(node.UpdatedAt)
	if err != nil {
		return err
	}
	var metadata model.MetadataEntity
	err = json.Unmarshal([]byte(node.Metadata), &metadata)
	if err != nil {
		return err
	}
	row := []string{metadata.Name, metadata.Description, jsonCard.PaymentSystem, jsonCard.Number,
		jsonCard.Holder, strconv.Itoa(jsonCard.CVC), jsonCard.EndDate.Format(constants.LayoutDate.ToString()),
		created.Format(constants.LayoutDateAndTime.ToString()), updated.Format(constants.LayoutDateAndTime.ToString())}
	*dataTblCard = append(*dataTblCard, row)
	return nil
}

func AppendBinary(node *grpc.Binary, dataTblBinary *[][]string) error {
	created, err := utils.ConvertTimestampToTime(node.CreatedAt)
	if err != nil {
		return err
	}
	row := []string{node.Name, created.Format(constants.LayoutDateAndTime.ToString())}
	*dataTblBinary = append(*dataTblBinary, row)
	return nil
}

func UpdateRowLoginPassword(login, password string, slice [][]string, indexRow int) [][]string {
	indexColLogin := 2
	indexColPassword := 3
	indexColUpdateAt := 5
	slice[indexRow][indexColLogin] = login
	slice[indexRow][indexColPassword] = password
	slice[indexRow][indexColUpdateAt] = time.Now().Format(constants.LayoutDateAndTime.ToString())
	return slice
}

func UpdateRowText(text string, slice [][]string, indexRow int) [][]string {
	indexColText := 2
	indexColUpdateAt := 4
	slice[indexRow][indexColText] = text
	slice[indexRow][indexColUpdateAt] = time.Now().Format(constants.LayoutDateAndTime.ToString())
	return slice
}

func UpdateRowCard(paymentSystem, number, holder, cvc, endDate string, slice [][]string, indexRow int) [][]string {
	indexColPaymentSystem := 2
	indexColNumber := 3
	indexColHolder := 4
	indexColCvc := 5
	indexColEndDate := 6
	indexColUpdateAt := 8
	slice[indexRow][indexColPaymentSystem] = paymentSystem
	slice[indexRow][indexColNumber] = number
	slice[indexRow][indexColHolder] = holder
	slice[indexRow][indexColEndDate] = endDate
	slice[indexRow][indexColCvc] = cvc
	slice[indexRow][indexColUpdateAt] = time.Now().Format(constants.LayoutDateAndTime.ToString())
	return slice
}
