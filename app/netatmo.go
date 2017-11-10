package app

import (
	netatmo "github.com/romainbureau/netatmo-api-go"
	"strings"
)

func getNetatmoClient() (*netatmo.Client, error) {
	netatmoClient, netatmoClientErr := netatmo.NewClient(netatmo.Config{
		ClientID:     Config.Netatmo.ClientID,
		ClientSecret: Config.Netatmo.ClientSecret,
		Username:     Config.Netatmo.Username,
		Password:     Config.Netatmo.Password,
	})

	return netatmoClient, netatmoClientErr
}

func readNetatmoValues(client *netatmo.Client, err error) (*netatmo.DeviceCollection, error) {
	if err != nil {
		return nil, err
	}

	values, err := client.Read()

	return values, err
}

func parseNetatmoValues(values *netatmo.DeviceCollection, err error) (NetatmoValues, error) {
	if err != nil {
		return NetatmoValues{}, err
	}
	var Values []NetatmoValue
	for _, station := range values.Stations() {
		for _, module := range station.Modules() {
			_, data := module.Data()
			for key, value := range data {
				Values = append(
					Values,
					NetatmoValue{
						ModuleName: strings.ToLower(module.ModuleName),
						MetricName: key,
						Value:      value.(float32),
					})
			}
		}
	}

	return NetatmoValues{Values: Values}, err
}

type NetatmoValues struct {
	Values []NetatmoValue
}

type NetatmoValue struct {
	ModuleName string
	MetricName string
	Value      float32
}

func GetNetatmoValues() (NetatmoValues, error) {
	client, err := getNetatmoClient()
	values, err := readNetatmoValues(client, err)
	data, err := parseNetatmoValues(values, err)

	return data, err
}
