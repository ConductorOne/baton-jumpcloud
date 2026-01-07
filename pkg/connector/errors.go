package connector

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/conductorone/baton-sdk/pkg/uhttp"
	"google.golang.org/grpc/codes"
)

// httpStatusToGRPCCode maps HTTP status codes to gRPC codes for error classification.
// This helps the Baton SDK determine if errors are retryable and how to handle them.
func httpStatusToGRPCCode(statusCode int) codes.Code {
	switch statusCode {
	case http.StatusBadRequest:
		return codes.InvalidArgument
	case http.StatusUnauthorized:
		return codes.Unauthenticated
	case http.StatusForbidden:
		return codes.PermissionDenied
	case http.StatusNotFound:
		return codes.NotFound
	case http.StatusConflict:
		return codes.AlreadyExists
	case http.StatusTooManyRequests:
		return codes.ResourceExhausted
	case http.StatusInternalServerError, http.StatusBadGateway, http.StatusServiceUnavailable:
		return codes.Unavailable
	case http.StatusGatewayTimeout:
		return codes.DeadlineExceeded
	default:
		if statusCode >= 500 {
			return codes.Unavailable
		}
		if statusCode >= 400 {
			return codes.InvalidArgument
		}
		return codes.Unknown
	}
}

// wrapSDKError wraps errors from the JumpCloud SDK (jcapi1/jcapi2) with appropriate gRPC codes.
// This allows the Baton SDK to properly classify errors for retry logic and alerting.
func wrapSDKError(err error, resp *http.Response, operation string) error {
	if err == nil {
		return nil
	}

	// If we don't have a response, inspect the error to classify it correctly
	if resp == nil {
		// Check for context errors (timeout, cancellation)
		if errors.Is(err, context.DeadlineExceeded) {
			return uhttp.WrapErrors(
				codes.DeadlineExceeded,
				fmt.Sprintf("%s: request timeout", operation),
				err,
			)
		}
		if errors.Is(err, context.Canceled) {
			return uhttp.WrapErrors(
				codes.Canceled,
				fmt.Sprintf("%s: request canceled", operation),
				err,
			)
		}

		// Check for URL errors (configuration issues, network errors, timeouts)
		var urlErr *url.Error
		if errors.As(err, &urlErr) {
			// Timeout on URL operation
			if urlErr.Timeout() {
				return uhttp.WrapErrors(
					codes.DeadlineExceeded,
					fmt.Sprintf("%s: request timeout", operation),
					urlErr,
				)
			}

			// Temporary network errors (DNS, connection issues)
			if urlErr.Temporary() {
				return uhttp.WrapErrors(
					codes.Unavailable,
					fmt.Sprintf("%s: temporary network error", operation),
					urlErr,
				)
			}

			// URL parsing errors indicate configuration issues
			if urlErr.Op == "parse" {
				return uhttp.WrapErrors(
					codes.InvalidArgument,
					fmt.Sprintf("%s: invalid URL configuration", operation),
					urlErr,
				)
			}
		}

		// Default: assume it's a network/connection error (service unavailable)
		return uhttp.WrapErrors(
			codes.Unavailable,
			fmt.Sprintf("%s: network or connection error", operation),
			err,
		)
	}

	// Map HTTP status codes to gRPC codes
	code := httpStatusToGRPCCode(resp.StatusCode)
	var message string

	switch code {
	case codes.InvalidArgument:
		if resp.StatusCode == http.StatusBadRequest {
			message = fmt.Sprintf("%s: invalid request", operation)
		} else {
			message = fmt.Sprintf("%s: client error (status %d)", operation, resp.StatusCode)
		}
	case codes.Unauthenticated:
		message = fmt.Sprintf("%s: authentication failed", operation)
	case codes.PermissionDenied:
		message = fmt.Sprintf("%s: permission denied", operation)
	case codes.NotFound:
		message = fmt.Sprintf("%s: resource not found", operation)
	case codes.AlreadyExists:
		message = fmt.Sprintf("%s: resource already exists", operation)
	case codes.ResourceExhausted:
		message = fmt.Sprintf("%s: rate limit exceeded", operation)
	case codes.Unavailable:
		if resp.StatusCode >= 500 && resp.StatusCode < 600 {
			message = fmt.Sprintf("%s: service unavailable", operation)
		} else {
			message = fmt.Sprintf("%s: server error (status %d)", operation, resp.StatusCode)
		}
	case codes.DeadlineExceeded:
		message = fmt.Sprintf("%s: request timeout", operation)
	default:
		message = fmt.Sprintf("%s: unexpected status code %d", operation, resp.StatusCode)
	}

	return uhttp.WrapErrors(code, message, err)
}
