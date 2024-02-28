package main

import (
	"context"
	"fmt"
	"net/url"
	"os"

	"github.com/vmware/govmomi"
	"github.com/vmware/govmomi/guest"
	"github.com/vmware/govmomi/vim25/types"
)

func main() {
	// Read vCenter login credentials from environment variables
	vcenterURL := os.Getenv("VCENTER_URL")
	username := os.Getenv("VCENTER_USERNAME")
	password := os.Getenv("VCENTER_PASSWORD")

	if vcenterURL == "" || username == "" || password == "" {
		fmt.Println("Please set VCENTER_URL, VCENTER_USERNAME, and VCENTER_PASSWORD environment variables")
		return
	}

	fmt.Println("vCenter URL: ", vcenterURL)
	fmt.Println("Username: ", username)
	fmt.Println("Password: ", password)

	// Parse vCenter URL
	vCenterURL, err := url.Parse(vcenterURL)
	if err != nil {
		fmt.Printf("Failed to parse vCenter URL: %s\n", err)
		return
	}

	// Set username and password
	vCenterURL.User = url.UserPassword(username, password)

	// Set insecure flag
	vCenterURL.RawQuery = url.Values{"insecure": {"true"}}.Encode()

	// Print the vCenter URL
	fmt.Println("vCenter URL: ", vCenterURL.String())

	// Connect to vCenter Server
	client, err := govmomi.NewClient(context.Background(), vCenterURL, true)
	if err != nil {
		fmt.Printf("Failed to connect to vCenter: %s\n", err)
		return
	}

	vmRef := types.ManagedObjectReference{
		Type:  "VirtualMachine",
		Value: "vm-51207",
	}
	// vm := object.NewVirtualMachine(client.Client, vmRef)

	ops := guest.NewOperationsManager(client.Client, vmRef)
	am, err := ops.AuthManager(context.Background())
	if err != nil {
		fmt.Printf("Failed to get AuthManager: %s\n", err)
		return
	}

	authentication := &types.NamePasswordAuthentication{
		Username: "arcvmw\\Administrator",
		Password: "Password~1",
	}

	// Validate credentials
	err = am.ValidateCredentials(context.Background(), authentication)
	if err != nil {
		fmt.Printf("Failed to validate credentials: %s\n", err)
		return
	}

	fmt.Println("Credentials validated successfully.")

}
