package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Transaction struct represent the data structure for the payload
type Transaction struct {
	symbol    string `json:"symbol"`
	price     uint64 `json:"price"`
	timestamp uint64 `json:"timestamp"`
}

// Transaction status represent status in string
type TxStatus string

// BroadcastTransaction broadcasts a transaction to the api
func BroadcastTransaction(tx Transaction) (string, error) {
	jsonData, err := json.Marshal(tx)
	if err != nil {
		// Handle the error
		return "", fmt.Errorf("failed to marshal data: %w", err)
	}
	url := "https://mock-node-wgqbnxruha-as.a.run.app/broadcast"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		// Handle the error
		return "", fmt.Errorf("failed to create a http request : %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// Handle the error
		return "", fmt.Errorf("failed to make a request to api server : %w", err)
	}
	defer resp.Body.Close()

	// Print response body for debugging
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}
	fmt.Println(string(responseBody))

	// Check response status code
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	return "", nil
}

// MonitorTransactionStatus monitors the status of a transaction by its hash
func MonitorTransactionStatus(txHash string) (TxStatus, error) {
	url := fmt.Sprintf("https://mock-node-wgqbnxruha-as.a.run.app/check/%s", txHash)
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to get transaction status: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	var status TxStatus
	fmt.Fscan(bytes.NewReader(body), &status) // Simpler way to read plaintext response

	return status, nil
}

func main() {
	tx := Transaction{
		symbol:    "ETH",
		price:     4500,
		timestamp: 1678912345,
	}
	txHash, err := BroadcastTransaction(tx)

	if err != nil {
		fmt.Println("Error broadcasting transaction:", err)
	}

	fmt.Println("Transaction hash:", txHash)

	// Monitor transaction status with retry logic
	// var status TxStatus
	// for {
	// 	status, err = MonitorTransactionStatus(txHash)
	// 	if err != nil {
	// 		fmt.Println("Error getting transaction status:", err)
	// 	} else {
	// 		fmt.Println("Transaction status:", status)
	// 		break
	// 	}
	// 	time.Sleep(5 * time.Second) // Wait 5 seconds before retrying
	// }
}
