package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Transaction struct represent the data structure for the payload
type Transaction struct {
	Symbol    string `json:"symbol"`
	Price     uint64 `json:"price"`
	Timestamp uint64 `json:"timestamp"`
}

// Transaction status represent status in string
type TxStatus string

// BroadcastTransaction broadcasts a transaction to the api
func BroadcastTransaction(tx Transaction) (string, error) {
	jsonData, err := json.Marshal(tx)

	if err != nil {
		return "", fmt.Errorf("failed to marshal data: %w", err)
	}

	url := "https://mock-node-wgqbnxruha-as.a.run.app/broadcast"

	res, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		return "", fmt.Errorf("failed to make a post request %w", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read the response body %w", err)
	}
	var result map[string]string

	err = json.Unmarshal([]byte(body), &result)

	if err != nil {
		return "", fmt.Errorf("error unmarshaling data from request %w", err)
	}

	fmt.Println(42, result)
	return result["tx_hash"], nil
}

// MonitorTransactionStatus monitors the status of a transaction by its hash
func MonitorTransactionStatus(txHash string) (TxStatus, error) {
	url := fmt.Sprintf("https://mock-node-wgqbnxruha-as.a.run.app/check/%s", txHash)
	res, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to get transaction status: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	var status TxStatus
	fmt.Fscan(bytes.NewReader(body), &status) // Simpler way to read plaintext response

	return status, nil
}

func main() {
	tx := Transaction{
		Symbol:    "ETH",
		Price:     4500,
		Timestamp: 1678912345,
	}
	txHash, err := BroadcastTransaction(tx)

	if err != nil {
		fmt.Println("Error broadcasting transaction:", err)
	}

	fmt.Println("Transaction hash:", txHash)

	// Monitor transaction status with retry logic
	var status TxStatus
	for {
		status, err = MonitorTransactionStatus(txHash)
		if err != nil {
			fmt.Println("Error getting transaction status:", err)
		} else {
			fmt.Println("Transaction status:", status)
			break
		}
		time.Sleep(5 * time.Second) // Wait 5 seconds before retrying
	}
}
