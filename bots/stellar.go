package bots

import (
	"fmt"
	"strings"
	"os"
	"encoding/csv"
	"encoding/json"
)

type StellarProfile struct {
	ProfileName      string         `json:"profileName"`
	Email            string         `json:"email"`
	Phone            string         `json:"phone"`
	BillingDifferent bool           `json:"billingAsShipping"`
	Card             StellarCard    `json:"payment"`
	Shipping         StellarAddress `json:"shipping"`
	Billing          StellarAddress `json:"billing"`
}
type StellarCard struct {
	Name     string `json:"cardName"`
	Type     string `json:"cardType"`
	Number   string `json:"cardNumber"`
	ExpMonth string `json:"cardMonth"`
	ExpYear  string `json:"cardYear"`
	Cvv      string `json:"cardCvv"`
}

type StellarAddress struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Address1  string `json:"address1"`
	Address2  string `json:"address2"`
	City      string `json:"city"`
	Zip       string `json:"zip"`
	Country   string `json:"country"`
	State     string `json:"state"`
}

func Stellar() {
	file, err := os.Open("input.csv")
	if err != nil {
		fmt.Println("Error opening CSV file!")
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	if len(records) < 2 {
		fmt.Println("Invalid CSV file!")
		return
	}

	headers := records[0]
	data := records[1:]

	var profiles []StellarProfile

	for i, row := range data {
		if len(row) != len(headers) {
			fmt.Printf("Skipping row %d due to mismatch in column count\n", i+1)
			continue
		}

		profile := createStellarProfile(row)
		profiles = append(profiles, profile)
	}

	output, err := json.MarshalIndent(profiles, "", "  ")
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}

	err = os.WriteFile("stellar.json", output, 0644)
	if err != nil {
		fmt.Println("Error writing JSON file:", err)
		return
	}

	fmt.Println("Conversion successful! JSON saved to stellar.json")
}

func createStellarProfile(row []string) StellarProfile {
	return StellarProfile{
		ProfileName:      row[0],
		Email:            row[1],
		Phone:            row[2],
		BillingDifferent: strings.ToLower(row[10]) != "true",
		Shipping: StellarAddress{
			FirstName: row[3],
			LastName:  row[4],
			Address1:  row[5],
			Address2:  row[6],
			City:      row[7],
			Zip:       row[8],
			Country:   "US",
			State:     GetStateAbbreviation(row[9]),
		},
		Billing: StellarAddress{
			FirstName: row[3],
			LastName:  row[4],
			Address1:  row[5],
			Address2:  row[6],
			City:      row[7],
			Zip:       row[8],
			Country:   "US",
			State:     GetStateAbbreviation(row[9]),
		},
		Card: StellarCard{
			Name: row[3] + " " + row[4],
			Type: row[11],
			Number:   row[12],
			ExpMonth: row[13],
			ExpYear:  row[14],
			Cvv:      row[15],
		},
	}
}
