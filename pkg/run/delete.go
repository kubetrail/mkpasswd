package run

import (
	"fmt"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"github.com/spf13/cobra"
	secretmanagerpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
	"k8s.io/apimachinery/pkg/util/validation"
)

func Delete(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	persistentFlags := getPersistentFlags(cmd)

	name := args[0]

	if err := setAppCredsEnvVar(persistentFlags.ApplicationCredentials); err != nil {
		err := fmt.Errorf("could not set Google Application credentials env. var: %w", err)
		return err
	}

	if len(name) == 0 {
		return fmt.Errorf("please provide name of the password")
	}

	if errs := validation.IsDNS1123Label(name); len(errs) > 0 {
		return fmt.Errorf("invalid name, need DNS1123Label format: %v", errs)
	}

	// Create the client.
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("failed to create secret manager client: %w", err)
	}
	defer client.Close()

	// Build the request.
	deleteRequest := &secretmanagerpb.DeleteSecretRequest{
		Name: fmt.Sprintf("projects/%s/secrets/%s", persistentFlags.Project, name),
	}

	// Call the API.
	if err := client.DeleteSecret(ctx, deleteRequest); err != nil {
		return fmt.Errorf("failed to delete secret: %w", err)
	}

	return nil
}
