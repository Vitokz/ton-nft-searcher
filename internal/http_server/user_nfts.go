package httpserver

import (
	"net/http"
	"strconv"

	"github.com/Vitokz/ton-nft-searcher/internal/http_server/transport"
	"github.com/Vitokz/ton-nft-searcher/internal/http_server/types"

	"github.com/Vitokz/ton-nft-searcher/internal/protocol"
	"github.com/labstack/echo/v4"
	"github.com/tonkeeper/tongo"
)

// UserNFTs gets user nfts.
func (s *Server) UserNFTs(c echo.Context) error {
	// swagger:operation GET /user/nft_items nft_items
	//
	// Get user nft items
	//
	// ---
	// produces:
	// - application/json
	// parameters:
	// - name: owner
	//   in: query
	//   description: |
	//  	NFT owner address in human format
	//   type: string
	// - name: collection
	//   in: query
	//   description: |
	//  	Collection address in human format
	//   type: string
	// - name: limit
	//   in: query
	//   description: |
	//  	Limit of NFTs in request
	//   type: integer
	// - name: offset
	//   in: query
	//   description: |
	//  	Offset of NFTs in request
	//   type: integer
	// tags:
	// - backend
	// responses:
	//   '200':
	//     description: user nfts
	//     schema:
	//          type: object
	//          example:
	//              nft_items:
	//                    - address: "EQDwipfknWmyUJjknngMM_r6DpWaYqC0u3SxuzOnAFQG9wjE"
	//                      approved_by:
	//                           - "GetGems"
	//                      collection:
	//                           address: "EQDwipfknWmyUJjknngMM_r6DpWaYqC0u3SxuzOnAFQG9wjE"
	//                           name: "some_collection_name"
	//                      collection_address: "EQDwipfknWmyUJjknngMM_r6DpWaYqC0u3SxuzOnAFQG9wjE"
	//                      index: 100
	//                      metadata:
	//                             image: "some_image_url.png"
	//                             name:  "some_name"
	//                             attributes:
	//                                  - trait_type: "Rarity"
	//                                    value: "Epic"
	//                      owner:
	//                           address: "EQDwipfknWmyUJjknngMM_r6DpWaYqC0u3SxuzOnAFQG9wjE"
	//                           is_scam: false
	//                      previews:
	//                           - resolution: "100x100"
	//                             url: "some_url.webp"
	//                      verified: true
	//   '400':
	//     description: bad request
	//     schema:
	//          type: object
	//          properties:
	//          	message:
	//          		type: string
	//          example:
	//          	message: invalid address
	//   '500':
	//     description: internal server error
	//     schema:
	//          type: object
	//          properties:
	//          	message:
	//          		type: string
	//          example:
	//          	message: internal server error
	s.logger.Debug().Msg("Got UserNFTs request!")

	var reqParams protocol.NFTSearchReq

	err := parseUserNFTsParameters(c.Request(), &reqParams)
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed to parse input params")

		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	resp, err := s.useCases.SearchUserNFTs(c.Request().Context(), reqParams)
	if err != nil {
		s.logger.Error().Err(err).Msg("Failed at during use-case")

		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	var userNFTs types.NFTItems
	userNFTs.FromDTO(resp.NFTs)

	return c.JSON(http.StatusOK, transport.UserNFTsResponse{
		NFTItems: userNFTs,
	})
}

func parseUserNFTsParameters(r *http.Request, params *protocol.NFTSearchReq) error {
	owner := r.URL.Query().Get("owner")
	if owner != "" {
		ownerAddr, err := tongo.ParseAccountID(owner)
		if err != nil {
			return ErrInvalidUserAddress
		}

		params.OwnerAddr = &ownerAddr
	}

	collection := r.URL.Query().Get("collection")
	if collection != "" {
		collectionAddr, err := tongo.ParseAccountID(owner)
		if err != nil {
			return ErrInvalidCollectionAddress
		}

		params.CollectionAddr = &collectionAddr
	}

	limitStr := r.URL.Query().Get("limit")
	if limitStr != "" {
		limit, err := strconv.ParseInt(limitStr, 10, 64)
		if err != nil {
			return ErrInvalidLimit
		}

		params.Limit = int(limit)
	}

	offsetStr := r.URL.Query().Get("offset")
	if offsetStr != "" {
		offset, err := strconv.ParseUint(offsetStr, 10, 64)
		if err != nil {
			return ErrInvalidOffset
		}

		params.Offset = int(offset)
	}

	return nil
}
