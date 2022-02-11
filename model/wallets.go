package model

type WalletMiniResponse struct {
	Username string `json:"username"`
	Balance  int    `json:"balance"`
}

type WalletUpdateRequest struct {
	Balance int `json:"balance"`
}
type WalletResponse []WalletMiniResponse

type WalletsResponse map[string]int
