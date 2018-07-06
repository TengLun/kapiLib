package campaignclient

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func loadCampaignList(name string) (creationList, error) {

	file, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
	}

	csv, err := csv.NewReader(file)
	if err != nil {
		fmt.Println(err)
	}

	data, err := csv.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	for _, datum := range data {
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
	}

	return cList, nil
}
