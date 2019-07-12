package blockatlas

type Collection struct {
	Name            string `json:"name"`
	Symbol          string `json:"symbol"`
	ImageUrl        string `json:"image_url"`
	Description     string `json:"description"`
	ExternalLink    string `json:"external_link"`
	Total           string `json:"total"`
	CategoryAddress string `json:"category_address"`
	Address         string `json:"address"`
	Version         string `json:"version"`
	Coin            uint   `json:"coin"`
	Type            string `json:"type"`
}

type CollectionPage []Collection

type Collectible struct {
	TokenID         string `json:"token_id"`
	ContractAddress string `json:"contract_address"`
	Category        string `json:"category"`
	ImageUrl        string `json:"image_url"`
	ExternalLink    string `json:"external_link"`
	Type            string `json:"type"`
	Description     string `json:"description"`
	Coin            uint   `json:"coin"`
	Name            string `json:"name"`
}

type CollectiblePage []Collectible