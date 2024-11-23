package bots

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type CyberProfileGroup struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Profiles []CyberProfile `json:"profiles"`
}
type CyberProfile struct {
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	Email            string   `json:"email"`
	Phone            string   `json:"phone"`
	BillingDifferent bool     `json:"billingDifferent"`
	Card             CyberCard     `json:"card"`
	Delivery         CyberAddress  `json:"delivery"`
	Billing          CyberAddress  `json:"billing"`
	Properties       struct{} `json:"properties"`
	GooglePay        []string `json:"googlePay"`
}
type CyberCard struct {
	Number   string `json:"number"`
	ExpMonth string `json:"expMonth"`
	ExpYear  string `json:"expYear"`
	Cvv      string `json:"cvv"`
}

type CyberAddress struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Address1  string `json:"address1"`
	Address2  string `json:"address2"`
	City      string `json:"city"`
	Zip       string `json:"zip"`
	Country   string `json:"country"`
	State     string `json:"state"`
}

func Cyber() {
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

	var profiles []CyberProfile

	for i, row := range data {
		if len(row) != len(headers) {
			fmt.Printf("Skipping row %d due to mismatch in column count\n", i+1)
			continue
		}

		profile := createCyberProfile(row)
		profiles = append(profiles, profile)

	}
	profileGroups := []CyberProfileGroup{
		{
			ID:       "",
			Name:     "Profiles",
			Profiles: profiles,
		},
	}

	// Write JSON to a file
	output, err := json.MarshalIndent(profileGroups, "", "  ")
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}

	err = os.WriteFile("cyber.json", output, 0644)
	if err != nil {
		fmt.Println("Error writing JSON file:", err)
		return
	}

	fmt.Println("Conversion successful! JSON saved to cyber.json")
}

func createCyberProfile(row []string) CyberProfile {
	return CyberProfile{
		ID:               "",
		Name:             row[0],
		Email:            row[1],
		Phone:            row[2],
		BillingDifferent: strings.ToLower(row[10]) == "true",
		Card: CyberCard{
			Number:   row[12],
			ExpMonth: row[13],
			ExpYear:  row[14],
			Cvv:      row[15],
		},
		Delivery: CyberAddress{
			FirstName: row[3],
			LastName:  row[4],
			Address1:  row[5],
			Address2:  row[6],
			City:      row[7],
			Zip:       row[8],
			Country:   "United States",
			State:     row[9],
		},
		Billing: CyberAddress{
			FirstName: row[3],
			LastName:  row[4],
			Address1:  row[5],
			Address2:  row[6],
			City:      row[7],
			Zip:       row[8],
			Country:   "United States",
			State:     row[9],
		},
		GooglePay: []string{},
	}
}
