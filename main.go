/*
 * Adyen API Client
 *
 * Contact: support@adyen.com
 */

package main

import (
    "fmt"

    "github.com/bos-hieu/adyen-go-api-library/v5/src/adyen"
    "github.com/bos-hieu/adyen-go-api-library/v5/src/common"
)

func main() {
    client := adyen.NewClient(&common.Config{Environment: common.TestEnv})
    fmt.Println("Welcome to Adyen API Client. Env: " + client.GetConfig().Environment)
}
