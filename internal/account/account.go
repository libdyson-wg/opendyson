package account

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/libdyson-wg/libdyson-go/cloud"
)

func Login(
	readLine func(prompt string) (string, error),
	readSecret func(prompt string) (string, error),
	beginLogin func(email string) (challengeID uuid.UUID, err error),
	completeLogin func(email, otpCode, password string, challengeID uuid.UUID) (token string, err error),
	setConfigToken func(token string) error,
	setServerToken func(token string),
	setServerRegion func(region cloud.ServerRegion),
) func() error {
	return func() error {
		cns, err := readLine("Use China region for account? (Default is no) [y/n]: ")
		if err != nil {
			return fmt.Errorf("error reading input: %w", err)
		}

		if cns == "y" || cns == "Y" {
			setServerRegion(cloud.RegionChina)
		}

		email, err := readLine("Email Address: ")
		if err != nil {
			return fmt.Errorf("error reading email: %w", err)
		}

		challenge, err := beginLogin(email)
		if err != nil {
			return fmt.Errorf("error starting login: %w", err)
		}

		otp, err := readLine("Code from confirmation email: ")
		if err != nil {
			return fmt.Errorf("error reading OTP Code: %w", err)
		}

		pw, err := readSecret("Password: ")
		if err != nil {
			return fmt.Errorf("error reading password: %w", err)
		}

		tok, err := completeLogin(email, otp, pw, challenge)
		if err != nil {
			return fmt.Errorf("error completing login: %w", err)
		}

		err = setConfigToken(tok)
		if err != nil {
			err = fmt.Errorf("error saving config: %w", err)
		}

		setServerToken(tok)

		return err
	}
}
