package models

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

type Category struct {
	ID      int    `db:"id"`
	Label   string `db:"label"`
	LogoUrl string `db:"logo_url"`
}

type Tag struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type Editor struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type ProductCreator struct {
	ProductId int `db:"id"`
	CreatorId int `db:"name"`
}

type Creator struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type Product struct {
	ID           int    `db:"id"`
	Name         string `db:"name"`
	Reference    string `db:"reference"`
	Price        int    `db:"price"`
	Description  string `db:"description"`
	Stock        int    `db:"stock"`
	Length       int    `db:"length"`
	Height       int    `db:"height"`
	Width        int    `db:"width"`
	Weight       int    `db:"weight"`
	CreationDate string `db:"creation_date"`
	IsArchived   bool   `db:"is_archived"`
	IsCollector  bool   `db:"is_collector"`
	CategoryId   int    `db:"category_id"`
	EditorId     int    `db:"editor_id"`
}

type ProductTag struct {
	ProductId int `db:"product_id"`
	TagId     int `db:"tag_id"`
}

type Picture struct {
	ID   int
	Link string
}

type Credential struct {
	ID         int    `db:"id"`
	PersonId   int    `db:"person_id"`
	Email      string `db:"email"`
	Password   string `db:"roles"`
	Roles      string `db:"password"`
	ResetToken string `db:"reset_token"`
}
type Person struct {
	ID           int    `db:"id"`
	CredentialId int    `db:"credential_id"`
	LastName     string `db:"last_name"`
	FirstName    string `db:"first_name"`
	PhoneNumber  string `db:"phone_number"`
}

type Comment struct {
	ID        int
	ProductID int
	PersonId  int
	Body      string
	Title     string
	Rate      int
	Date      time.Time
	Vote      int
}

type Status struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

type Address struct {
	ID           int    `db:"id"`
	PersonId     int    `db:"person_id"`
	StreetNumber int    `db:"street_number"`
	StreetName   string `db:"street_name"`
	City         string `db:"city"`
	Zipcode      int    `db:"zipcode"`
}

type Purchase struct {
	ID                   int       `db:"id"`
	PersonId             int       `db:"person_id"`
	StatusId             int       `db:"status_id"`
	AddressId            int       `db:"addresses_id"`
	DateDelivery         time.Time `db:"date_delivery"`
	DateExpectedDelivery time.Time `db:"date_expected_delivery"`
	DatePurchase         time.Time `db:"date_purchase"`
	Reference            string    `db:"reference"`
}

type Pick struct {
	ID         int
	ProductId  int
	PurchaseId int
	Quantity   int
	PriceItem  int
}

func (p *Product) ToString(categoriesBD []Category, editors []Editor) string {

	return fmt.Sprintf("(%v,%v,\"%v\",\"%v\",%v,\"%v\",%v,%v,%v,%v,%v,\"%v\",%v,%v),\n",
		getCategoryID(p.CategoryId, categoriesBD),
		getEditorID(p.EditorId, editors),
		removeDoubleQuotes(p.Name),
		removeDoubleQuotes(p.Reference),
		p.Price,
		removeDoubleQuotes(p.Description),
		p.Stock,
		p.Length,
		p.Height,
		p.Width,
		p.Weight,
		p.CreationDate,
		p.IsArchivedToInt(),
		p.IsCollectorToInt(),
	)
}

func getCategoryID(categoryID int, categories []Category) int {
	var categoryLabel string
	for _, c := range Categories {
		if c.ID == categoryID {
			categoryLabel = c.Label
			break
		}
	}

	for _, c := range categories {
		if c.Label == categoryLabel {
			return c.ID
		}
	}

	return 0
}

func getEditorID(editorID int, editors []Editor) int {
	var editorName string
	for _, editor := range Editors {
		if editorID == editor.ID {
			editorName = editor.Name
			break
		}
	}

	for _, editor := range editors {
		if editor.Name == editorName {
			return editor.ID
		}
	}

	return 0
}

func (p *Product) IsArchivedToInt() int {
	if p.IsArchived {
		return 1
	}
	return 0
}

func (p *Product) IsCollectorToInt() int {
	if p.IsCollector {
		return 1
	}
	return 0
}

func StructToList(items interface{}, fieldName string, haveLogo bool) string {
	var result string

	value := reflect.ValueOf(items)
	if value.Kind() != reflect.Slice {
		return "Not a slice"
	}

	structType := value.Type().Elem()
	_, found := structType.FieldByName(fieldName)
	if !found {
		return "Field not found"
	}

	for i := 0; i < value.Len(); i++ {
		item := value.Index(i)
		fieldValue := item.FieldByName(fieldName).Interface()

		if haveLogo {
			logoValue := item.FieldByName("LogoUrl").Interface()
			result = result + fmt.Sprintf("(\"%v\",\"%v\"),\n", fieldValue, logoValue)
			continue
		}

		result = result + fmt.Sprintf("(\"%v\"),\n", fieldValue)
	}

	result = strings.Trim(result, "\n,")

	return AddSemicolon(result)
}

func removeDoubleQuotes(input string) string {
	return strings.ReplaceAll(input, "\"", "")
}

func AddSemicolon(input string) string {
	return input + ";"
}

func LinkProductsAndCreators(products []Product, creators []Creator) []ProductCreator {
	productCreators := []ProductCreator{}

	for _, product := range products {
		for _, creator := range creators {
			update := false
			if strings.Contains(product.Name, creator.Name) {
				update = true
			} else if strings.Contains(product.Reference, strings.ToUpper(creator.Name)) {
				update = true
			}

			if update {
				productCreator := ProductCreator{
					ProductId: product.ID,
					CreatorId: creator.ID,
				}

				productCreators = append(productCreators, productCreator)
			}
		}
	}

	return productCreators
}

func JoinProductAndTag(products []Product, tags []Tag) []ProductTag {
	var productTags []ProductTag

	for _, product := range products {
		for _, tag := range tags {
			if strings.Contains(strings.ToLower(product.Name), strings.ToLower(tag.Name)) ||
				strings.Contains(strings.ToLower(product.Description), strings.ToLower(tag.Name)) {

				productTags = append(productTags, ProductTag{
					ProductId: product.ID,
					TagId:     tag.ID,
				})
			}
		}
	}

	return productTags
}
