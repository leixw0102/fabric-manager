package main

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

//EnrollUser enroll a user have registerd
func EnrollUser(username string, password string) (bool, error) {
	sdk := fabsdk.FabricSDK{}
	mspClient, err := msp.New(sdk.Context(), msp.WithOrg("Org1"))
	if err != nil {
		fmt.Printf("Failed to create msp client: %s\n", err)
		return true, err
	}

	_, err = mspClient.GetSigningIdentity(username)
	if err == msp.ErrUserNotFound {
		fmt.Println("Going to enroll user")
		err = mspClient.Enroll(username, msp.WithSecret(password))
		if err != nil {
			fmt.Printf("Failed to enroll user: %s\n", err)
			return true, err
		}
		fmt.Printf("Success enroll user: %s\n", username)
	} else if err != nil {
		fmt.Printf("Failed to get user: %s\n", err)
		return false, err
	}
	fmt.Printf("User %s already enrolled, skip enrollment.\n", username)
	return true, nil
}

//Register a new user with username , password and department.
func RegisterlUser(username, password, department string) error {
	sdk := fabsdk.FabricSDK{}
	mspClient, err := msp.New(sdk.Context(), msp.WithOrg("Org1"))
	if err != nil {
		fmt.Printf("Failed to create msp client: %s\n", err)
	}
	request := &msp.RegistrationRequest{
		Name:        username,
		Type:        "user",
		Affiliation: department,
		Secret:      password,
	}

	secret, err := mspClient.Register(request)
	if err != nil {
		fmt.Printf("register %s [%s]\n", username, err)
		return err
	}
	fmt.Printf("register %s successfully,with password %s\n", username, secret)
	return nil
}

func main() {
	EnrollUser("user1", "user1pw")
}
