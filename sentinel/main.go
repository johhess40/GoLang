/*
	This code is meant for testing of Azure Sentinel incident creation based on rules in Log Analytics
		-TODO: Make reusable so that sub id's etc. are not hard coded
		-TODO: Add logic for testing all types of Sentinel rules(Fusion,ML,Scheduled, etc)
		-TODO: Add functionality for reading YAML/JSON for Client configuration
  */

package main

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/securityinsight/mgmt/2020-01-01/securityinsight"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

func main() {

	rootCtx := context.Background()

	var topFive int32

	incidentClient := securityinsight.NewIncidentsClient("300f794a-d9b3-4332-a5db-b6c6d6e2a9c7")

	authorizer, err := auth.NewAuthorizerFromCLI()

	if err == nil {
		incidentClient.Authorizer = authorizer
		print("Base URI: " + incidentClient.BaseURI + "\n")
		print("Sub ID: " + incidentClient.SubscriptionID + "\n")
		print("UserAgent: " + incidentClient.UserAgent + "\n")
	}
	const testFive = 5

	topFive = int32(testFive)

	fmt.Printf(string(topFive))

	incList, err := incidentClient.List(rootCtx, "SIEM", "sentinel-law", "", "", &topFive, "")
	if err != nil {
		panic(err)
	}


	for _, v := range incList.Values() {
		fmt.Printf("%v\n", *v.Owner)
	}

}
