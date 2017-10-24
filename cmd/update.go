package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/equinox-io/equinox"
	"os"
)

// assigned when creating a new application in the equinox dashboard
const equinoxAppID = "app_itoJoCoW8dr"

// public portion of signing key generated by `equinox genkey`
var publicKey = []byte(`
-----BEGIN ECDSA PUBLIC KEY-----
MHYwEAYHKoZIzj0CAQYFK4EEACIDYgAEbSwe/sKG3zLxvkpx+sPX3s+QA9s7tMPn
UVQaNmLDPsbHzTZIchnWnla0hzt+LkObK5bxL5QiNdIEfD4ZIELalcY0g3JkVaBA
dDm0JDoKE+xOY9mx3raOGVzXNoxL1EAp
-----END ECDSA PUBLIC KEY-----
`)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update to the latest version",
	Long:  `Check for and automatically update to the latest version of the Cronitor CLI.`,
	Run: func(cmd *cobra.Command, args []string) {
		var opts equinox.Options
		if err := opts.SetPublicKeyPEM(publicKey); err != nil {
			fmt.Println("Update failed:", err)
			os.Exit(1)
		}

		// check for the update
		resp, err := equinox.Check(equinoxAppID, opts)
		switch {
		case err == equinox.NotAvailableErr:
			fmt.Println("No update available.")
			return
		case err != nil:
			fmt.Println("Update failed:", err)
			os.Exit(1)
		}

		// fetch the update and apply it
		err = resp.Apply()
		if err != nil {
			fmt.Println("Could not apply update:", err)
			os.Exit(1)
		}

		fmt.Printf("Updated to version: %s!\n", resp.ReleaseVersion)
	},
}

func init() {
	RootCmd.AddCommand(updateCmd)
}
