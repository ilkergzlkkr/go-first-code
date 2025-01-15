package helpers

import "os"

func IsLambda() bool {
	if lambdaTaskRoot := os.Getenv("LAMBDA_TASK_ROOT"); lambdaTaskRoot != "" {
		return true
	}

	if sst := os.Getenv("SST_STAGE"); sst != "" {
		return true
	}
	return false
}
