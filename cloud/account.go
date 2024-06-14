package cloud

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/libdyson-wg/libdyson-go/internal/generated/oapi"

	"github.com/oapi-codegen/runtime/types"

	"github.com/google/uuid"
)

func BeginLogin(email string) (challengeID uuid.UUID, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	{
		body := oapi.GetUserStatusJSONRequestBody{
			Email: types.Email(email),
		}
		resp, err := client.GetUserStatusWithResponse(ctx, nil, body)
		if err != nil {
			return uuid.Nil, fmt.Errorf("couldn't get user status: %w", err)
		}

		if resp.StatusCode() != http.StatusOK {
			return uuid.Nil, fmt.Errorf("couldn't get user status: http status code %d", resp.StatusCode())
		}
	}

	body := oapi.BeginLoginJSONRequestBody{
		Email: types.Email(email),
	}
	resp, err := client.BeginLoginWithResponse(ctx, nil, body)

	if err != nil {
		return uuid.Nil, fmt.Errorf("couldn't begin login: %w", err)
	}

	if resp.StatusCode() != http.StatusOK {
		return uuid.Nil, fmt.Errorf("couldn't begin login: http status code %d", resp.StatusCode())
	}

	return resp.JSON200.ChallengeId, nil
}

func CompleteLogin(email, otpCode, password string, challengeID uuid.UUID) (token string, err error) {
	body := oapi.CompleteLoginJSONRequestBody{
		Email:       types.Email(email),
		OtpCode:     otpCode,
		ChallengeId: challengeID,
		Password:    password,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	resp, err := client.CompleteLoginWithResponse(ctx, nil, body)

	if err != nil {
		return "", fmt.Errorf("couldn't complete login: %w", err)
	}

	if resp == nil || resp.JSON200 == nil {
		return "", fmt.Errorf("couldn't complete login: nil response")
	}

	return resp.JSON200.Token, nil
}
