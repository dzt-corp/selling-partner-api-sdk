package catalog

type V2022GetCatalogItemParams struct {

	// A marketplace identifier. Specifies the marketplace for the item.
	MarketplaceIds []string `json:"MarketplaceId"`
	IncludedData   []string `json:"IncludedData,omitempty"`
}

type ProductTypes []ProductType
type ProductType struct {
	ProductType   string `json:"productType"`
	MarketplaceId string `json:"marketplaceId"`
}

type ProductImages []ProductImage

type ProductImage struct {
	MarketplaceId string   `json:"marketplaceId"`
	Images        []PImage `json:"images"`
}
type PImage struct {
	Variant string `json:"variant"`
	Link    string `json:"link"`
	Height  int    `json:"height"`
	Width   int    `json:"width"`
}

type HeaderWearSize struct {
	Size      string `json:"size_to"`
	SizeClass string `json:"size_class"`
}

type HeaderWearSizes []HeaderWearSize
type Brand struct {
	Value         string `json:"value"`
	MarketplaceId string `json:"marketplace_id"`
}

type Brands []Brand
type ProductDescription struct {
	Language string `json:"language_tag"`
	Value    string `json:"value"`
}
type ProductAttributes struct {
	ProductDescription []*ProductDescription `json:"product_description,omitempty"`
	ProductTypes       *ProductTypes         `json:"productTypes,omitempty"`
	Brands             *Brands               `json:"brand,omitempty"`
	HeaderWearSizes    *HeaderWearSizes      `json:"headwear_size,omitempty"`
	ProductImages      *ProductImages        `json:"images,omitempty"`
}

type V2022GetCatalogItemResponse struct {
	Attributes   ProductAttributes `json:"attributes"`
	Images       *ProductImages    `json:"images,omitempty"`
	ProductTypes *ProductTypes     `json:"productTypes,omitempty"`
}
