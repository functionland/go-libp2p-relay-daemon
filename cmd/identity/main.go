package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"os"

	"github.com/libp2p/go-libp2p/core/crypto"
)

func main() {
	// Define a flag to accept the identity key from the command line
	var encodedKey string
	flag.StringVar(&encodedKey, "identity", "", "Base64 encoded private identity key")
	flag.Parse()

	// Check if the identity key was provided
	if encodedKey == "" {
		fmt.Println("Please provide a base64 encoded identity key using the --identity flag")
		os.Exit(1) // Exit if no identity key is provided
	}

	// Decode the base64 string to get the binary data
	keyBytes, err := base64.StdEncoding.DecodeString(encodedKey)
	if err != nil {
		panic(err)
	}

	// Unmarshal the binary data into a private key
	privKey, err := crypto.UnmarshalPrivateKey(keyBytes)
	if err != nil {
		panic(err)
	}

	// Marshal the private key to get the binary format
	privKeyBytes, err := crypto.MarshalPrivateKey(privKey)
	if err != nil {
		panic(err)
	}

	// Write the binary private key to a file
	err = os.WriteFile("identity.key", privKeyBytes, 0600)
	if err != nil {
		panic(err)
	}

	fmt.Println("Private key saved to 'identity.key'")
}
