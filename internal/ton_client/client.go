package tonclient

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/Vitokz/ton-nft-searcher/internal/ton_client/dto"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type Client struct {
	baseURL *url.URL
	logger  zerolog.Logger
	client  *http.Client
}

const defaultTimeout = 30 * time.Second

func New(baseURL string, logger zerolog.Logger) (*Client, error) {
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:    1,
			IdleConnTimeout: time.Minute,
		},
		Timeout: defaultTimeout,
	}

	logger = logger.With().Str("module", "ton-client").Logger()

	return &Client{
		logger:  logger,
		client:  client,
		baseURL: parsedURL,
	}, nil
}

func DoRequest[Resp any](client *http.Client, req *http.Request, respEntity Resp) error {
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	switch resp.StatusCode {
	case http.StatusOK:
		err = json.Unmarshal(resBody, respEntity)
		if err != nil {
			return err
		}

		return nil
	case http.StatusBadRequest:
		var errResp dto.ErrorResponse

		err = json.Unmarshal(resBody, &errResp)
		if err != nil {
			return err
		}

		return errors.Wrap(ErrBadRequest, errResp.Message)
	default:
		return ErrFailedAtDuringRequest
	}
}

func (c *Client) constructQueryURL(path string) string {
	baseURL := *c.baseURL

	baseURL.Path = path

	return baseURL.String()
}
