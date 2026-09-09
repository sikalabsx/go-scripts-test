package validate_jwt

import (
	"context"

	"github.com/coreos/go-oidc"
	"github.com/golang-jwt/jwt/v4"
)

func ValidateJWT(rawToken string) error {
	ctx := context.Background()

	claims := jwt.MapClaims{}
	_, _, err := jwt.NewParser().ParseUnverified(rawToken, claims)
	if err != nil {
		return err
	}

	issuer, ok := claims["iss"].(string)
	if !ok || issuer == "" {
		return jwt.NewValidationError("issuer claim not found", jwt.ValidationErrorClaimsInvalid)
	}

	provider, err := oidc.NewProvider(ctx, issuer)
	if err != nil {
		return err
	}

	_, err = provider.Verifier(&oidc.Config{SkipClientIDCheck: true}).Verify(ctx, rawToken)
	if err != nil {
		return err
	}

	return nil
}
