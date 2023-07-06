package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	projectID := "your-project-id"
	datasetID := "your-dataset-id"
	tableID := "your-table-id"

	url := fmt.Sprintf("https://bigquery.googleapis.com/bigquery/v2/projects/%s/datasets/%s/tables/%s/data", projectID, datasetID, tableID)

	requestBody, err := json.Marshal(map[string]interface{}{
		"rows": []map[string]interface{}{
			{
				"json": map[string]interface{}{
					"column1": "value1",
					"column2": "value2",
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
