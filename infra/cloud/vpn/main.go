package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	// Initialize a session in the region that you want to use.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"),
	})
	if err != nil {
		log.Fatalf("failed to create session, %v", err)
	}
	// Create an EC2 service client.
	svc := ec2.New(sess)
	// Replace with your Client VPN endpoint ID
	vpnEndpointID := "cvpn-endpoint-082aa292aeb2ccb1e"
	// Get VPN endpoint status
	vpnStatus, err := getVPNStatus(svc, vpnEndpointID)
	if err != nil {
		log.Fatalf("failed to get VPN status, %v", err)
	}
	fmt.Printf("VPN Status: %s\n", vpnStatus)

	// Get VPN client connections
	clientConnections, err := getVPNClientConnections(svc, vpnEndpointID)
	if err != nil {
		log.Fatalf("failed to get VPN client connections, %v", err)
	}
	fmt.Printf("Number of connected clients: %d\n", clientConnections)
}

func getVPNStatus(svc *ec2.EC2, vpnEndpointID string) (string, error) {
	// Describe the Client VPN endpoint
	input := &ec2.DescribeClientVpnEndpointsInput{
		ClientVpnEndpointIds: []*string{aws.String(vpnEndpointID)},
	}

	result, err := svc.DescribeClientVpnEndpoints(input)
	if err != nil {
		return "", err
	}

	if len(result.ClientVpnEndpoints) == 0 {
		return "Unknown", fmt.Errorf("no VPN endpoint found with ID %s", vpnEndpointID)
	}

	// Return the state of the VPN endpoint
	return aws.StringValue(result.ClientVpnEndpoints[0].Status.Code), nil
}

func getVPNClientConnections(svc *ec2.EC2, vpnEndpointID string) (int, error) {
	// Describe the Client VPN connections
	input := &ec2.DescribeClientVpnConnectionsInput{
		ClientVpnEndpointId: aws.String(vpnEndpointID),
	}

	result, err := svc.DescribeClientVpnConnections(input)
	if err != nil {
		return 0, err
	}

	// Count the number of active connections
	activeConnections := 0
	for _, connection := range result.Connections {
		if aws.StringValue(connection.Status.Code) == "active" {
			activeConnections++
		}
	}

	return activeConnections, nil
}
