package script

import (
	"fmt"
	"strings"

	"training.go/scriptPixelPerfect/models"
	"training.go/scriptPixelPerfect/store"
)

type Script struct {
	store store.Store
}

type Scripter interface {
	GetScript() (string, error)
}

func New(storer store.Store) *Script {
	return &Script{
		store: storer,
	}
}

// GetScript delete allInfo on database and add product and table in relation with product,
// return script of query making on database
func (s *Script) GetScript() (string, error) {

	if err := s.store.DeleteAllInfo(); err != nil {
		return "", err
	}

	fmt.Println("Delete product and relation")

	neededToProduct, err := s.addNeededForProduct()
	if err != nil {
		return "", err
	}

	productScript, err := s.addProduct()
	if err != nil {
		return "", err
	}

	insertProductCreator, err := s.addProductCreator()
	if err != nil {
		return "", err
	}

	insertProductTag, err := s.addProductTag()
	if err != nil {
		return "", err
	}

	insertPicture, err := s.linkProductAndPicture()
	if err != nil {
		return "", err
	}

	insertPersonAndCredential, err := s.addPersonAndCredential()
	if err != nil {
		return "", err
	}

	insertComment, err := s.addComment()
	if err != nil {
		return "", err
	}

	insertStatus, err := s.addStatus()
	if err != nil {
		return "", err
	}

	persons, err := s.store.GetPersons()
	if err != nil {
		return "", err
	}

	insertAddress, err := s.addAddress(persons)
	if err != nil {
		return "", err
	}

	insertPurchase, err := s.addPurchase(persons)
	if err != nil {
		return "", err
	}

	insertPick, err := s.addPick()
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("%v\n\n%v\n\n%v\n\n%v\n\n%v\n\n%v\n\n%v\n\n%v\n\n%v\n\n%v\n\n%v",
		neededToProduct,
		productScript,
		insertProductCreator,
		insertProductTag,
		insertPicture,
		insertPersonAndCredential,
		insertComment,
		insertStatus,
		insertAddress,
		insertPurchase,
		insertPick,
	)

	return result, nil
}

func (s *Script) addPick() (string, error) {
	insertPick := "INSERT INTO pick (product_id, purchase_id, quantity, priceitem) VALUES \n"

	purchases, err := s.store.GetPurchase()
	if err != nil {
		return "", err
	}

	products, err := s.store.GetProduct()
	if err != nil {
		return "", err
	}

	for _, p := range models.Picks {
		insertPick += fmt.Sprintf("(%v,%v,%v,%v),\n", products[p.ProductId-1].ID, purchases[p.PurchaseId-1].ID, p.Quantity, p.PriceItem)
	}

	insertPick = models.AddSemicolon(strings.Trim(insertPick, "\n,"))
	if err := s.store.Insert(insertPick); err != nil {
		return "", err
	}

	fmt.Println("Pick inserted")

	return insertPick, nil
}

func (s *Script) addPurchase(persons []models.Person) (string, error) {
	insertPurchase := "INSERT INTO purchase (person_id,status_id, addresses_id, date_delivery, date_expected_delivery, date_purchase, reference) VALUES\n"

	addresses, err := s.store.GetAddress()
	if err != nil {
		return "", err
	}

	status, err := s.store.GetStatus()

	for _, p := range models.Purchases {
		insertPurchase += fmt.Sprintf(
			"(%v,%v, %v, '%v','%v','%v','%v'),\n",
			persons[p.PersonId-1].ID,
			status[p.StatusId-1].ID,
			addresses[p.AddressId-1].ID,
			p.DateDelivery.Format("2006-01-02 15:04:05"),
			p.DateExpectedDelivery.Format("2006-01-02 15:04:05"),
			p.DatePurchase.Format("2006-01-02 15:04:05"),
			p.Reference,
		)
	}

	insertPurchase = models.AddSemicolon(strings.Trim(insertPurchase, "\n,"))
	if err := s.store.Insert(insertPurchase); err != nil {
		return "", err
	}

	fmt.Println("Purchase inserted")

	return insertPurchase, nil
}

func (s *Script) addAddress(persons []models.Person) (string, error) {
	insertAddress := "INSERT INTO address (person_id, street_number, street_name, city, zipcode) VALUES \n"

	for _, a := range models.Addresses {
		insertAddress += fmt.Sprintf("(%v,%v,'%v','%v',%v),\n", persons[a.PersonId-1].ID, a.StreetNumber, a.StreetName, a.City, a.Zipcode)
	}

	insertAddress = models.AddSemicolon(strings.Trim(insertAddress, "\n,"))
	if err := s.store.Insert(insertAddress); err != nil {
		return "", err
	}

	fmt.Println("Address inserted")

	return insertAddress, nil
}

func (s *Script) addStatus() (string, error) {
	insertStatus := "INSERT INTO status (name) VALUES \n"

	for _, s := range models.AllStatus {
		insertStatus += fmt.Sprintf("('%v'),\n", s.Name)
	}

	insertStatus = models.AddSemicolon(strings.Trim(insertStatus, "\n,"))
	if err := s.store.Insert(insertStatus); err != nil {
		return "", err
	}

	fmt.Println("Status inserted")

	return insertStatus, nil
}

func (s *Script) addComment() (string, error) {
	insertComment := "INSERT INTO comment (product_id, person_id, body, title, rate, date, vote) VALUES \n"

	products, err := s.store.GetProduct()
	if err != nil {
		return "", err
	}

	persons, err := s.store.GetPersons()
	if err != nil {
		return "", err
	}

	for _, c := range models.Comments {
		insertComment += fmt.Sprintf("(%v,%v,'%v','%v',%v,'%v',%v),", products[c.ProductID-1].ID, persons[c.PersonId-1].ID, c.Body, c.Title, c.Rate, c.Date.Format("2006-01-02 15:04:05"), c.Vote)
	}

	insertComment = models.AddSemicolon(strings.Trim(insertComment, "\n,"))
	if err := s.store.Insert(insertComment); err != nil {
		return "", err
	}

	fmt.Println("Comment inserted")

	return insertComment, nil
}

func (s *Script) addPersonAndCredential() (string, error) {
	insertCredential := "INSERT INTO credential (person_id, email, roles, password, is_verified, reset_token, is_activated, is_archived, is_blocked) VALUES\n"
	insertPerson := "INSERT INTO person (last_name, first_name, phone_number) VALUES\n"

	for _, p := range models.Persons {
		insertPerson += fmt.Sprintf("('%v','%v','%v'),", p.LastName, p.FirstName, p.PhoneNumber)
	}

	insertPerson = models.AddSemicolon(strings.Trim(insertPerson, "\n,"))

	if err := s.store.Insert(insertPerson); err != nil {
		return "", err
	}

	fmt.Println("Person inserted")

	persons, err := s.store.GetPersons()
	if err != nil {
		return "", err
	}

	for index, c := range models.Credentials {
		insertCredential += fmt.Sprintf("(%v,'%v','%v','%v',0,'%v',0,0,0),", persons[index].ID, c.Email, c.Roles, c.Password, c.ResetToken)
	}

	insertCredential = models.AddSemicolon(strings.Trim(insertCredential, "\n,"))

	if err := s.store.Insert(insertCredential); err != nil {
		return "", err
	}

	fmt.Println("Person inserted")

	credentials, err := s.store.GetCredentials()
	if err != nil {
		return "", err
	}

	results := fmt.Sprintf("%v\n\n%v\n\n", insertCredential, insertPerson)

	for index, c := range credentials {
		toUpdate := fmt.Sprintf("UPDATE person SET credential_id = %v WHERE id = %v;\n", c.ID, persons[index].ID)
		if err := s.store.Insert(toUpdate); err != nil {
			return "", err
		}

		results += fmt.Sprintf("%v\n", toUpdate)
	}

	fmt.Println("Person updated")

	return results, nil
}

func (s *Script) addProductTag() (string, error) {
	insertProductTag := "INSERT INTO product_tag  (product_id,tag_id) VALUES\n"

	products, err := s.store.GetProduct()
	if err != nil {
		return "", err
	}

	tags, err := s.store.GetTag()
	if err != nil {
		return "", err
	}

	productsTags := models.JoinProductAndTag(products, tags)

	for _, productTag := range productsTags {
		insertProductTag += fmt.Sprintf("(%v,%v),\n", productTag.ProductId, productTag.TagId)
	}

	insertProductTag = models.AddSemicolon(strings.Trim(insertProductTag, "\n,"))
	if err := s.store.Insert(insertProductTag); err != nil {
		return "", err
	}

	fmt.Println("ProductTag inserted")

	return insertProductTag, nil
}

func (s *Script) addProductCreator() (string, error) {
	insertProductCreator := "INSERT INTO product_creator (product_id, creator_id) VALUES\n"

	products, err := s.store.GetProduct()
	if err != nil {
		return "", err
	}

	creators, err := s.store.GetCreator()
	if err != nil {
		return "", err
	}

	for _, productCreator := range models.LinkProductsAndCreators(products, creators) {
		insertProductCreator += fmt.Sprintf("(%v,%v),\n", productCreator.ProductId, productCreator.CreatorId)
	}

	insertProductCreator = models.AddSemicolon(strings.Trim(insertProductCreator, "\n,"))
	if err := s.store.Insert(insertProductCreator); err != nil {
		return "", err
	}

	fmt.Println("ProductCreator inserted")

	return insertProductCreator, nil
}

func (s *Script) addNeededForProduct() (string, error) {
	if err := s.store.DeleteAllInfo(); err != nil {
		return "", err
	}

	insertCategories := fmt.Sprintf(
		"INSERT INTO category (label,logo_url) VALUES\n%v",
		models.StructToList(models.Categories, "Label", true),
	)
	if err := s.store.Insert(insertCategories); err != nil {
		return "", err
	}

	fmt.Println("Categorie inserted")

	insertTag := fmt.Sprintf(
		"INSERT INTO tag (`name`) VALUES\n%v",
		models.StructToList(models.Tags, "Name", false),
	)
	if err := s.store.Insert(insertTag); err != nil {
		return "", err
	}

	fmt.Println("Tag inserted")

	insertCreator := fmt.Sprintf(
		"INSERT INTO creator (`name`) VALUES\n%v",
		models.StructToList(models.Creators, "Name", false),
	)
	if err := s.store.Insert(insertCreator); err != nil {
		return "", err
	}

	fmt.Println("Creator inserted")

	insertEditor := fmt.Sprintf(
		"\n\nINSERT INTO editor (`name`) VALUES\n%v",
		models.StructToList(models.Editors, "Name", false),
	)
	if err := s.store.Insert(insertEditor); err != nil {
		return "", err
	}

	fmt.Println("Editor inserted")

	return fmt.Sprintf("%v\n\n%v\n\n%v\n\n%v", insertCategories, insertTag, insertCreator, insertEditor), nil
}

func (s *Script) addProduct() (string, error) {
	insertProduct := "INSERT INTO product (category_id, editor_id, name, reference, price, description, stock, length, height, width, weight, creation_date, is_archived, is_collector) VALUES\n"

	editors, err := s.store.GetEditor()
	if err != nil {
		return "", err
	}

	categories, err := s.store.GetCategory()
	if err != nil {
		return "", err
	}

	for _, product := range models.Products {
		insertProduct += fmt.Sprint(product.ToString(categories, editors))
	}

	insertProduct = models.AddSemicolon(strings.Trim(insertProduct, "\n,"))
	err = s.store.Insert(insertProduct)
	if err != nil {
		return "", err
	}

	fmt.Println("product inserted")

	return insertProduct, nil
}

func (s *Script) linkProductAndPicture() (string, error) {
	insertPicture := "INSERT INTO picture (product_id, name, url, alt) VALUES\n"

	products, err := s.store.GetProduct()
	if err != nil {
		return "", err
	}

	index := 0
	for _, product := range products {
		if index > len(models.Pictures)-1 {
			insertPicture += fmt.Sprintf("(%v,'no image','%v','no image'),", product.ID, models.NoImage)
		} else {
			insertPicture += fmt.Sprintf("(%v,\"%v\",\"%v\",\"%v\"),", product.ID, product.Name, models.Pictures[index].Link, product.Name)
		}
		index++
	}

	insertPicture = models.AddSemicolon(strings.Trim(insertPicture, "\n,"))
	err = s.store.Insert(insertPicture)
	if err != nil {
		return "", err
	}

	fmt.Println("Picture inserted")

	return insertPicture, nil
}
