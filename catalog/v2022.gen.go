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

type ShirtSize struct {
	Size      string `json:"size"`
	SizeClass string `json:"size_class"`
}

type ShirtSizes []ShirtSize

type Brand struct {
	Value         string `json:"value"`
	MarketplaceId string `json:"marketplace_id"`
}

type Brands []Brand
type ProductDescription struct {
	Language string `json:"language_tag"`
	Value    string `json:"value"`
}

type ProductColor struct {
	Language          string    `json:"language_tag"`
	StandardizedColor []*string `json:"standardized_values,omitempty"`
	Value             string    `json:"value"`
}
type Size struct {
	Unit  string  `json:"unit"`
	Value float64 `json:"value"`
}

type Capacities []Size

type PItemSize struct {
	Width  Size `json:"width"`
	Height Size `json:"height"`
}

type UnitType struct {
	Value string `json:"value"`
}
type Unit struct {
	Type  UnitType `json:"type"`
	Value float64  `json:"value"`
}

type ProductSize []PItemSize
type ProductAttributes struct {
	ProductDescription []*ProductDescription `json:"product_description,omitempty"`
	ProductTypes       *ProductTypes         `json:"productTypes,omitempty"`
	ProductColor       []*ProductColor       `json:"color,omitempty"`
	Brands             *Brands               `json:"brand,omitempty"`
	HeaderWearSizes    *HeaderWearSizes      `json:"headwear_size,omitempty"`
	ShirtSizes         *ShirtSizes           `json:"shirt_size,omitempty"`
	Capacities         *Capacities           `json:"capacity,omitempty"`
	ProductImages      *ProductImages        `json:"images,omitempty"`
	Unit               []*Unit               `json:"unit_count,omitempty"`
	ProductSize        *ProductSize          `json:"item_width_height,omitempty"`
}

type BrowseClassification struct {
	DisplayName      string `json:"displayName"`
	ClassificationId string `json:"classificationId,omitempty"`
}
type ProductSummary struct {
	MarketplaceId        string                `json:"marketplaceId"`
	AdultProduct         *bool                 `json:"adultProduct,omitempty"`
	Autographed          *bool                 `json:"autographed,omitempty"`
	BrowseClassification *BrowseClassification `json:"browseClassification,omitempty"`
	Color                string                `json:"color,omitempty"`
	ItemClassification   *string               `json:"itemClassification,omitempty"`
	ItemName             *string               `json:"itemName,omitempty"`
	Memorabilia          *bool                 `json:"memorabilia,omitempty"`
	Size                 *string               `json:"size,omitempty"`
	Style                *string               `json:"style,omitempty"`
}
type V2022GetCatalogItemResponse struct {
	Attributes     ProductAttributes `json:"attributes"`
	Images         *ProductImages    `json:"images,omitempty"`
	ProductTypes   *ProductTypes     `json:"productTypes,omitempty"`
	ProductSummary *[]ProductSummary `json:"summaries,omitempty"`
}
