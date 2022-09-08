/*
 * Adyen API Client
 *
 * Contact: support@adyen.com
 */

package main

import (
	"fmt"

	"github.com/bos-hieu/adyen-go-api-library/src/adyen"
	"github.com/bos-hieu/adyen-go-api-library/src/common"
)

func main() {
	client := adyen.NewClient(&common.Config{Environment: common.TestEnv})
	fmt.Println("Welcome to Adyen API Client. Env: " + client.GetConfig().Environment)
}
