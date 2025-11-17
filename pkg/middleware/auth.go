package middleware

import (
	"context"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	authHeader   = "authorization"
	bearerPrefix = "Bearer "
)

// AuthMiddleware is a middleware function that checks for a valid JWT token in the request header.
func AuthMiddleware() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		// Extract the token from the incoming gRPC metadata
		token, err := extractToken(ctx)
		if err != nil {
			return nil, err
		}

		// Validate the token (placeholder validation)
		if !isValidToken(token) {
			return nil, status.Errorf(codes.Unauthenticated, "invalid token")
		}

		// Continue processing the request
		return handler(ctx, req)
	}
}

// extractToken extracts the JWT token from gRPC metadata (authorization header).
func extractToken(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "missing metadata")
	}

	// metadata keys are lowercase
	if vals := md[authHeader]; len(vals) > 0 {
		val := vals[0]
		// strip "Bearer " prefix if present
		if strings.HasPrefix(val, bearerPrefix) {
			return strings.TrimPrefix(val, bearerPrefix), nil
		}
		return val, nil
	}

	return "", status.Errorf(codes.Unauthenticated, "authorization token required")
}

// isValidToken checks if the provided token is valid.
// Placeholder: replace with real JWT validation.
func isValidToken(token string) bool {
	return token != ""
}
