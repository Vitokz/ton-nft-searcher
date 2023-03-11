package tonclient

import (
	"context"
	"net/http"
	"strconv"
)

const (
	maxLimit = 1000
)

func (c *Client) NFTSearch(args dto.NFTSearchArqs) (protocol.NFTSearchResp, error) {
	req, err := c.buildNFTSearchArgs(args)
	if err != nil {
		return protocol.NFTSearchResp{}, err
	}

	var resp dto.NFTSearchResp

	err = DoRequest(c.client, req, &resp)
	if err != nil {
		c.logger.Error().Err(err).Msg("Failed to at during request")

		return protocol.NFTSearchResp{}, err
	}

	return protocol.NFTSearchResp{
		NFTs: resp.NFTItems.ToDTO(),
	}, nil
}

func (c *Client) buildNFTSearchArgs(args dto.NFTSearchArqs) (*http.Request, error) {
	url := c.constructQueryURL(dto.SearchNFTItems)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		c.logger.Error().Err(err).Msg("Failed to generate request")

		return nil, ErrFailedToCreateRequest
	}

	q := req.URL.Query()

	if args.Limit == 0 || args.Limit > maxLimit {
		args.Limit = maxLimit
	}

	q.Add("limit", strconv.Itoa(args.Limit))

	q.Add("offset", strconv.Itoa(args.Offset))

	if args.IncludeOnSale != nil {
		q.Add("include_on_sale", strconv.FormatBool(*args.IncludeOnSale))
	}

	if args.CollectionAddr != nil {
		collectionAddr := args.CollectionAddr

		q.Add("collection", collectionAddr.String())
	}

	if args.OwnerAddr != nil {
		userAddr := args.OwnerAddr

		q.Add("owner", userAddr.String())
	}

	req.URL.RawQuery = q.Encode()

	return req, nil
}
