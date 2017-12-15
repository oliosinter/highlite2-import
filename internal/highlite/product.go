package highlite

import "highlite-parser/internal/csv"

const (
	titleProductNo         string = "Product No."
	titleProductName       string = "Product Name"
	titleCountryOfOrigin   string = "Country of Origin"
	titleWeight            string = "Weight (kg)"
	titleLength            string = "Length (m)"
	titleWidth             string = "Width (m)"
	titleHeight            string = "Height (m)"
	titleQuantityInCarton  string = "Qty. in Carton"
	titleEANNo             string = "EAN No."
	titleInternetAdvice    string = "Internet Advice Prijs"
	titleUnitPrice         string = "Unit Price"
	titleMinSalesQuantity  string = "Min. Sales Qty."
	titleCategory          string = "Category"
	titleSubcategory1      string = "Subcategory 1"
	titleSubcategory2      string = "Subcategory 2"
	titleTariffNo          string = "Tariff No."
	titleStatus            string = "Status"
	titleAccessory         string = "Accessory"
	titleSubstitute        string = "Substitute"
	titleCatalogPage       string = "Catalog Page"
	titleInStock           string = "In Stock"
	titleExpWeekOfArrival  string = "\"Exp. Week of Arrival\""
	titleWebshop           string = "Webshop"
	titleSubheadingEN      string = "Subheading EN"
	titleMainDescriptionEN string = "Main Description EN"
	titleSpecsEN           string = "Specs EN"
	titleImagesList        string = "Images List"
	titleBrand             string = "Brand"
	titleSubcategory3      string = "Subcategory 3"
)

// Product is a highlite product.
type Product struct {
	No          string
	Name        string
	SubHeading  string
	Description string
	Specs       string
	Brand       string

	Country string
	Weight  float64
	Length  float64
	Width   float64
	Height  float64
	Price   float64

	Category1 *Category
	Category2 *Category
	Category3 *Category
}

// GetProductFromCSVImport creates product object from csv import data.
func GetProductFromCSVImport(mapper *csv.TitleMap, values []string) Product {
	cat1 := NewCategory(mapper.GetString(titleCategory, values), nil)
	cat2 := NewCategory(mapper.GetString(titleSubcategory1, values), cat1)
	cat3 := NewCategory(mapper.GetString(titleSubcategory2, values), cat2)

	product := Product{
		Category1: cat1,
		Category2: cat2,
		Category3: cat3,

		No:          mapper.GetString(titleProductNo, values),
		Name:        mapper.GetString(titleProductName, values),
		SubHeading:  mapper.GetString(titleSubheadingEN, values),
		Description: mapper.GetString(titleMainDescriptionEN, values),
		Specs:       mapper.GetString(titleSpecsEN, values),
		Brand:       mapper.GetString(titleBrand, values),

		Country: mapper.GetString(titleCountryOfOrigin, values),
		Weight:  mapper.GetFloat(titleWeight, values),
		Length:  mapper.GetFloat(titleLength, values),
		Width:   mapper.GetFloat(titleWidth, values),
		Height:  mapper.GetFloat(titleHeight, values),
		Price:   mapper.GetFloat(titleUnitPrice, values),
	}

	return product
}
