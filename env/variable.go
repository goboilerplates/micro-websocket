package env

import (
	"os"
	"strconv"
)

var (

	// EnableProdMode .
	EnableProdMode = true

	// ServerHost .
	ServerHost = ":9000"

	// EnableCompression .
	EnableCompression = true

	// ReadBufferSize .
	ReadBufferSize = 1024

	// WriteBufferSize .
	WriteBufferSize = 1024
)

// LoadVariables gets environment variables.
func LoadVariables() {
	if os.Getenv("GONI_COMPRESS") == "false" {
		EnableCompression = false
	}

	readBufferSizeEnv, readBufferSizeEnvErr := strconv.Atoi(os.Getenv("GONI_RBSIZE"))
	if readBufferSizeEnvErr == nil && readBufferSizeEnv > 0 {
		ReadBufferSize = readBufferSizeEnv
	}

	writeBufferSizeEnv, writeBufferSizeEnvErr := strconv.Atoi(os.Getenv("GONI_WBSIZE"))
	if writeBufferSizeEnvErr == nil && writeBufferSizeEnv > 0 {
		WriteBufferSize = writeBufferSizeEnv
	}

	if os.Getenv("GBP_PROMODE") == "false" {
		EnableProdMode = false
	}

	host := os.Getenv("GBP_HOST")
	if len(host) > 0 {
		ServerHost = host
	}
}
