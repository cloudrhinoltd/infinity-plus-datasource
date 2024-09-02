/**
 * This file is based on the original work from Grafana Labs © 2023.
 * Modifications were made by Syncfish Pty Ltd © 2024
 * The Syncfish Logo and Name are registered Trademarks of Syncfish Pty Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Modifications:
 * - Integrated Azure Managed Identity authentication.
 * - Adjusted data handling logic in building query Authorization header.
 * - Other enhancements and bug fixes as documented in the Syncfish version control system.
 */

package infinity

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/cloudrhinoltd/infinity-plus-datasource/pkg/models"
	"github.com/grafana/grafana-plugin-sdk-go/backend"
)

// Define the default resource for Azure
const defaultResource = "https://management.azure.com/"

const dummyHeader = "xxxxxxxx"

const (
	contentTypeJSON           = "application/json"
	contentTypeFormURLEncoded = "application/x-www-form-urlencoded"
)

const (
	headerKeyAccept        = "Accept"
	headerKeyContentType   = "Content-Type"
	headerKeyAuthorization = "Authorization"
	headerKeyIdToken       = "X-ID-Token"
)

// ApplyAcceptHeader applies the Accept header based on the query type.
func ApplyAcceptHeader(query models.Query, settings models.InfinitySettings, req *http.Request, includeSect bool) *http.Request {
	if query.Type == models.QueryTypeJSON || query.Type == models.QueryTypeGraphQL {
		req.Header.Set(headerKeyAccept, `application/json;q=0.9,text/plain`)
	}
	if query.Type == models.QueryTypeCSV {
		req.Header.Set(headerKeyAccept, `text/csv; charset=utf-8`)
	}
	if query.Type == models.QueryTypeXML {
		req.Header.Set(headerKeyAccept, `text/xml;q=0.9,text/plain`)
	}
	return req
}

// ApplyContentTypeHeader sets the Content-Type header for POST requests.
func ApplyContentTypeHeader(query models.Query, settings models.InfinitySettings, req *http.Request, includeSect bool) *http.Request {
	if strings.ToUpper(query.URLOptions.Method) == http.MethodPost {
		switch query.URLOptions.BodyType {
		case "raw":
			if query.URLOptions.BodyContentType != "" {
				req.Header.Set(headerKeyContentType, query.URLOptions.BodyContentType)
			}
		case "form-data":
			writer := multipart.NewWriter(&bytes.Buffer{})
			for _, f := range query.URLOptions.BodyForm {
				_ = writer.WriteField(f.Key, f.Value)
			}
			if err := writer.Close(); err != nil {
				return req
			}
			req.Header.Set(headerKeyContentType, writer.FormDataContentType())
		case "x-www-form-urlencoded":
			req.Header.Set(headerKeyContentType, contentTypeFormURLEncoded)
		case "graphql":
			req.Header.Set(headerKeyContentType, contentTypeJSON)
		default:
			req.Header.Set(headerKeyContentType, contentTypeJSON)
		}
	}
	return req
}

// ApplyHeadersFromSettings applies custom headers from settings to the request.
func ApplyHeadersFromSettings(settings models.InfinitySettings, req *http.Request, includeSect bool) *http.Request {
	for key, value := range settings.CustomHeaders {
		val := dummyHeader
		if includeSect {
			val = value
		}
		if key != "" {
			req.Header.Add(key, val)
			if strings.EqualFold(key, headerKeyAccept) || strings.EqualFold(key, headerKeyContentType) {
				req.Header.Set(key, val)
			}
		}
	}
	return req
}

// ApplyHeadersFromQuery applies headers from the query to the request.
func ApplyHeadersFromQuery(query models.Query, settings models.InfinitySettings, req *http.Request, includeSect bool) *http.Request {
	for _, header := range query.URLOptions.Headers {
		value := dummyHeader
		if includeSect {
			value = replaceSect(header.Value, settings, includeSect)
		}
		if header.Key != "" {
			req.Header.Add(header.Key, value)
			if strings.EqualFold(header.Key, headerKeyAccept) || strings.EqualFold(header.Key, headerKeyContentType) {
				req.Header.Set(header.Key, value)
			}
		}
	}
	return req
}

// ApplyBasicAuth applies basic authentication using username and password from settings.
func ApplyBasicAuth(settings models.InfinitySettings, req *http.Request, includeSect bool) *http.Request {
	if settings.BasicAuthEnabled && (settings.UserName != "" || settings.Password != "") {
		basicAuthHeader := fmt.Sprintf("Basic %s", dummyHeader)
		if includeSect {
			basicAuthHeader = "Basic " + base64.StdEncoding.EncodeToString([]byte(settings.UserName+":"+settings.Password))
		}
		req.Header.Set(headerKeyAuthorization, basicAuthHeader)
	}
	return req
}

// ApplyBearerToken applies bearer token authentication using settings.
func ApplyBearerToken(settings models.InfinitySettings, req *http.Request, includeSect bool) *http.Request {
	if settings.AuthenticationMethod == models.AuthenticationMethodBearerToken {
		bearerAuthHeader := fmt.Sprintf("Bearer %s", dummyHeader)
		if includeSect {
			bearerAuthHeader = fmt.Sprintf("Bearer %s", settings.BearerToken)
		}
		req.Header.Add(headerKeyAuthorization, bearerAuthHeader)
	}
	return req
}

// ApplyApiKeyAuth applies API key authentication using settings.
func ApplyApiKeyAuth(settings models.InfinitySettings, req *http.Request, includeSect bool) *http.Request {
	if settings.AuthenticationMethod == models.AuthenticationMethodApiKey && settings.ApiKeyType == models.ApiKeyTypeHeader {
		apiKeyHeader := dummyHeader
		if includeSect {
			apiKeyHeader = settings.ApiKeyValue
		}
		if settings.ApiKeyKey != "" {
			req.Header.Add(settings.ApiKeyKey, apiKeyHeader)
		}
	}
	return req
}

// ApplyForwardedOAuthIdentity applies forwarded OAuth identity headers if enabled in settings.
func ApplyForwardedOAuthIdentity(requestHeaders map[string]string, settings models.InfinitySettings, req *http.Request, includeSect bool) *http.Request {
	if settings.ForwardOauthIdentity {
		authHeader := dummyHeader
		token := dummyHeader
		if includeSect {
			authHeader = requestHeaders[headerKeyAuthorization]
			token = requestHeaders[headerKeyIdToken]
		}
		req.Header.Add(headerKeyAuthorization, authHeader)
		if requestHeaders[headerKeyIdToken] != "" {
			req.Header.Add(headerKeyIdToken, token)
		}
	}
	return req
}

// Retrieve Azure token using MSI or fallback to Azure CLI credentials.
// Added by Syncfish Pty Ltd © 2024. Licensed under the Apache License, Version 2.0.
func getToken() (string, error) {
	// Set up a timeout context for MSI token retrieval
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check if the app mode is not 'development' to use MSI
	if os.Getenv("GF_DEFAULT_APP_MODE") != "development" {
		// Try fetching token from the MSI endpoint
		token, err := getMSIToken(timeoutCtx)
		if err == nil {
			backend.Logger.FromContext(timeoutCtx).Debug("Successfully fetched MSI token.")
			return token, nil
		}
		backend.Logger.FromContext(timeoutCtx).Error("Failed to fetch MSI token, falling back to Azure CLI credentials.", "error", err.Error())
	}

	// If MSI token retrieval fails or app is in development mode, fallback to Azure CLI credentials
	cliCred, err := azidentity.NewAzureCLICredential(nil)
	if err != nil {
		return "", fmt.Errorf("failed to create Azure CLI credential: %w", err)
	}

	// Retrieve the token using Azure CLI credentials
	tokenResp, err := cliCred.GetToken(timeoutCtx, policy.TokenRequestOptions{Scopes: []string{defaultResource}})
	if err != nil {
		return "", fmt.Errorf("failed to fetch Azure CLI token: %w", err)
	}

	return tokenResp.Token, nil
}

// Fetch token directly from the MSI endpoint without TenantID.
// Added by Syncfish Pty Ltd © 2024. Licensed under the Apache License, Version 2.0.
func getMSIToken(ctx context.Context) (string, error) {
	msiEndpoint := "http://169.254.169.254/metadata/identity/oauth2/token"
	apiVersion := "2018-02-01"

	// Create HTTP request to fetch MSI token
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, msiEndpoint, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request for MSI token: %w", err)
	}

	// Set query parameters and headers for the MSI request
	q := req.URL.Query()
	q.Add("api-version", apiVersion)
	q.Add("resource", defaultResource)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("Metadata", "true")

	// Execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to get MSI token: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code %d from MSI endpoint", resp.StatusCode)
	}

	// Parse the response to extract the access token
	var result struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode MSI token response: %w", err)
	}

	return result.AccessToken, nil
}

// ApplyAzureManagedIdentity applies the Azure Managed Identity token to the request headers.
// Added by Syncfish Pty Ltd © 2024. Licensed under the Apache License, Version 2.0.
func ApplyAzureManagedIdentity(requestHeaders map[string]string, settings models.InfinitySettings, req *http.Request, includeSect bool) *http.Request {
	// Retrieve token and set Authorization header
	logger := backend.Logger.FromContext(context.Background())
	if settings.AuthenticationMethod == models.AuthenticationMethodManagedIdentity {
		token, err := getToken()
		if err != nil {
			logger.Error("failed to get Azure token", "error", err.Error())
			return req
		}
		// Ensure the Authorization header is set
		req.Header.Add(headerKeyAuthorization, fmt.Sprintf("Bearer %s", token))
		logger.Info("Authorization header set successfully", "Authorization", token)
	}

	return req
}
