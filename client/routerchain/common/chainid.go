package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RPCRequest struct {
	JSONRPC string                 `json:"jsonrpc"`
	Method  string                 `json:"method"`
	Params  map[string]interface{} `json:"params"`
	ID      int                    `json:"id"`
}

type RPCResponse struct {
	Result map[string]interface{} `json:"result"`
}

type NodeInfo struct {
	Network string `json:"network"`
}

func FetchChainID(rpcURL string) (string, error) {

	// Create the request payload
	payload := RPCRequest{
		JSONRPC: "2.0",
		Method:  "status",
		Params:  make(map[string]interface{}),
		ID:      1,
	}

	// Convert payload to JSON
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	// Send the POST request to the Tendermint RPC endpoint
	resp, err := http.Post(rpcURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Parse the response
	var rpcResponse RPCResponse
	err = json.Unmarshal(body, &rpcResponse)
	if err != nil {
		return "", err
	}

	// Extract the chain ID from the response
	nodeInfo := NodeInfo{}
	if info, ok := rpcResponse.Result["node_info"].(map[string]interface{}); ok {
		if network, ok := info["network"].(string); ok {
			nodeInfo.Network = network
		}
	}

	return nodeInfo.Network, nil
}

func main() {
	// Fetch the chain ID
	rpcURL := "http://localhost:26657"
	chainID, err := FetchChainID(rpcURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the chain ID
	fmt.Println("Chain ID:", chainID)
}
