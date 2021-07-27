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
	// create a VirtualNetworks client
	//vnetClient := network.NewVirtualNetworksClient("b3c6d20c-ca4c-4ea7-af76-da58d3e316f6")
	incidentClient := securityinsight.NewIncidentsClient("300f794a-d9b3-4332-a5db-b6c6d6e2a9c7")
	//.incidentClient("b3c6d20c-ca4c-4ea7-af76-da58d3e316f6")
	// create an authorizer from env vars or Azure Managed Service Idenity
	authorizer, err := auth.NewAuthorizerFromCLI()
	if err == nil {
		incidentClient.Authorizer = authorizer
		print("Base URI: " + incidentClient.BaseURI + "\n")
		print("Sub ID: " + incidentClient.SubscriptionID + "\n")
		print("UserAgent: " + incidentClient.UserAgent + "\n")
	}
	const topFiveCnst = 5
	topFive = int32(topFiveCnst)
	fmt.Printf(string(topFive))
	incList, err := incidentClient.List(rootCtx, "SIEM", "sentinel-law", "", "", &topFive, "")
	if err != nil {
		panic(err)
	}
	// s := fmt.Sprintf("%v\n", incList.Values())

	for _, v := range incList.Values() {
		//	fmt.Printf("%v\n", v.Severity)
		//fmt.Printf("%v\n", v.IncidentProperties)
		// fmt.Printf("%v\n", v.UnmarshalJSON(bits))
		fmt.Printf("%v\n", *v.Owner)
	}

	// fmt.Printf("%T\n", incList.Values())

	// io.WriteString(os.Stdout, s)
	fmt.Printf(string(topFive))
}
