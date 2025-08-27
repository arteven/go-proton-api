package proton

import (
	"context"
	"fmt"

	"github.com/go-resty/resty/v2"
)

// Pass API endpoints
const (
	passSharesPath     = "/pass/v1/share"
	passSharePath      = "/pass/v1/share/%s"
	passShareKeysPath  = "/pass/v1/share/%s/key"
	passItemsPath      = "/pass/v1/share/%s/item"
)

// GetPassShares retrieves all Pass shares for the authenticated user
func (c *Client) GetPassShares(ctx context.Context) ([]PassShareResponse, error) {
	var res PassSharesResponse
	
	if err := c.do(ctx, func(r *resty.Request) (*resty.Response, error) {
		return r.SetResult(&res).Get(passSharesPath)
	}); err != nil {
		return nil, fmt.Errorf("getting pass shares: %w", err)
	}

	return res.Shares, nil
}

// GetPassShare retrieves a specific Pass share by ID
func (c *Client) GetPassShare(ctx context.Context, shareID string) (*PassShareResponse, error) {
	var res struct {
		Code  int               `json:"Code"`
		Share PassShareResponse `json:"Share"`
	}
	
	path := fmt.Sprintf(passSharePath, shareID)
	if err := c.do(ctx, func(r *resty.Request) (*resty.Response, error) {
		return r.SetResult(&res).Get(path)
	}); err != nil {
		return nil, fmt.Errorf("getting pass share %s: %w", shareID, err)
	}

	return &res.Share, nil
}

// GetPassShareKeys retrieves the encryption keys for a specific Pass share
func (c *Client) GetPassShareKeys(ctx context.Context, shareID string) ([]PassShareKeyResponse, error) {
	var res PassShareKeysResponse
	
	path := fmt.Sprintf(passShareKeysPath, shareID)
	if err := c.do(ctx, func(r *resty.Request) (*resty.Response, error) {
		return r.SetResult(&res).Get(path)
	}); err != nil {
		return nil, fmt.Errorf("getting pass share keys for %s: %w", shareID, err)
	}

	return res.ShareKeys.Keys, nil
}

// GetPassItems retrieves all items for a specific Pass share
func (c *Client) GetPassItems(ctx context.Context, shareID string) ([]PassItemResponse, error) {
	var res PassItemsResponse
	
	path := fmt.Sprintf(passItemsPath, shareID)
	if err := c.do(ctx, func(r *resty.Request) (*resty.Response, error) {
		return r.SetResult(&res).Get(path)
	}); err != nil {
		return nil, fmt.Errorf("getting pass items for share %s: %w", shareID, err)
	}

	// Return the items from RevisionsData
	return res.Items.RevisionsData, nil
}