package tiktok

import (
	"context"
	"fmt"
	"net/http"
)

// GenerateAuthURL generate auth url for user to login.
// Doc: https://bytedance.feishu.cn/docs/doccnROmkE6WI9zFeJuT3DQ3YOg#ckvNFO
func (c *Client) GenerateAuthURL(state string) string {
	return fmt.Sprintf(
		AuthBaseURL+"/oauth/authorize?app_key=%s&state=%s",
		c.appKey, state,
	)
}

/**
{
    "code": 0,
    "message": "success",
    "data": {
        "access_token": "TTP_F9t7wQAAAAAqhavs-UCqgsvpfMbIeHkgwjhz2JUhelGklA6pyvNhPVwHj71v__F-4IefdT09TAOGrHU03LeN5KTM7qJS3d8U2Qq1LU6bJX_IObRgyyGryge7l3rmiduQmBv7hMbZE5mOw1_5ot0xiIhncRC3eCF9h5bqq6zxpPsVtrGjig4tuQ",
        "access_token_expire_in": 1705073850,
        "refresh_token": "TTP_lYnukAAAAADclTdpedWBWCOfn95DH4j49d9SSmMkEDvgXMGcHhB7HtnXzjqPmrKKddFx4eVQZuY",
        "refresh_token_expire_in": 1710994818,
        "open_id": "mRd0RgAAAAAwOD86290YUwIp0iecIUEKfZsTlfulY3QHZ4s-ShhS4w",
        "seller_name": "TestMaker",
        "seller_base_region": "US",
        "user_type": 0
    },
    "request_id": "2024010515382517EBC389FB2D9801AF2B"
}
**/
// GetAccessToken get access token from tiktok server.
// Doc: https://bytedance.feishu.cn/docs/doccnROmkE6WI9zFeJuT3DQ3YOg#qYtWHF
func (c *Client) GetAccessToken(ctx context.Context, code string) (resp AccessTokenResponse, err error) {
	grantType := "authorized_code"
	err = c.request(
		ctx, http.MethodGet, AuthBaseURL,
		fmt.Sprintf("/api/v2/token/get?app_key=%s&app_secret=%s&auth_code=%s&grant_type=%s",
			c.appKey,
			c.appSecret,
			code,
			grantType),
		nil, nil, &resp)
	return
}

// RefreshToken refresh access token.
// Doc: https://bytedance.feishu.cn/docs/doccnROmkE6WI9zFeJuT3DQ3YOg#bG2h09
func (c *Client) RefreshToken(ctx context.Context, rk string) (resp AccessTokenResponse, err error) {
	var req RefreshTokenRequest
	req.AppKey = c.appKey
	req.AppSecret = c.appSecret
	req.RefreshToken = rk
	req.GrantType = "refresh_token"
	r := c.prepareBody(req)
	err = c.request(
		ctx, http.MethodGet, AuthBaseURL,
		"/api/v2/token/refresh",
		nil, r, &resp)
	return
}
