package client

import (
	"encoding/json"
	"strconv"

	"github.com/LobovVit/GophKeeper/internal/client/constants"
	"github.com/LobovVit/GophKeeper/internal/client/model"
	grpc "github.com/LobovVit/GophKeeper/internal/proto"
	"github.com/LobovVit/GophKeeper/pkg/utils"
)

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
