package cloud

import (
	"context"
	"fmt"
	"time"

	"github.com/libdyson-wg/libdyson-go/internal/generated/oapi"

	"github.com/oapi-codegen/runtime/types"

	"github.com/google/uuid"
)

type httpBeginLoginFunc func(
	ctx context.Context,
	params *oapi.BeginLoginParams,
	body oapi.BeginLoginJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.BeginLoginResponse, error)

func BeginLogin(
	requestBeginLogin httpBeginLoginFunc,
) func(email string) (challengeID uuid.UUID, err error) {
	return func(email string) (challengeID uuid.UUID, err error) {
		body := oapi.BeginLoginJSONRequestBody{
			Email: types.Email(email),
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		resp, err := requestBeginLogin(ctx, nil, body)

		if err != nil || resp.JSON200 == nil {
			return uuid.Nil, fmt.Errorf("couldn't perform login: error %w, http status %s", err, resp.Status())
		}

		return uuid.UUID(resp.JSON200.ChallengeId), nil
	}
}

type httpCompleteLoginFunc func(
	ctx context.Context,
	params *oapi.CompleteLoginParams,
	body oapi.CompleteLoginJSONRequestBody,
	reqEditors ...oapi.RequestEditorFn,
) (*oapi.CompleteLoginResponse, error)

func CompleteLogin(
	requestCompleteLogin httpCompleteLoginFunc,
) func(
	email string,
	otpCode string,
	challengeID uuid.UUID,
	password string,
) (token string, err error) {
	return func(
		email string,
		otpCode string,
		challengeID uuid.UUID,
		password string,
	) (token string, err error) {
		body := oapi.CompleteLoginJSONRequestBody{
			Email:       types.Email(email),
			OtpCode:     otpCode,
			ChallengeId: challengeID,
			Password:    password,
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		resp, err := requestCompleteLogin(ctx, nil, body)

		if err != nil || resp.JSON200 == nil {
			return "", fmt.Errorf("couldn't perform login: error %w, http status %s", err, resp.Status())
		}

		return resp.JSON200.Tokeng, nil

	}
}
