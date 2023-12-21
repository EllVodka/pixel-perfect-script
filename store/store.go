package store

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"training.go/scriptPixelPerfect/models"
)

type Store interface {
	Open() error
	Close() error

	GetCreator() ([]models.Creator, error)
	GetProduct() ([]models.Product, error)
	GetPurchase() ([]models.Purchase, error)
	GetCategory() ([]models.Category, error)
	GetEditor() ([]models.Editor, error)
	GetTag() ([]models.Tag, error)
	GetPersons() ([]models.Person, error)
	GetCredentials() ([]models.Credential, error)
	GetAddress() ([]models.Address, error)
	GetStatus() ([]models.Status, error)
	DeleteAllInfo() error

	Insert(query string) error
}

type DbStore struct {
	db  *sqlx.DB
	cfg models.StoreCfg
}

func New(cfg models.StoreCfg) *DbStore {
	return &DbStore{
		cfg: cfg,
	}
}

func (store *DbStore) Open() error {
	db, err := sqlx.Connect(
		"mysql",
		fmt.Sprintf("%v:%v@(%v:%v)/%v?parseTime=true",
			store.cfg.User,
			store.cfg.Password,
			store.cfg.Server,
			store.cfg.Port,
			store.cfg.Database,
		),
	)
	if err != nil {
		return err
	}

	log.Println("Connected to DB")
	store.db = db
	return nil
}

func (store *DbStore) GetAddress() ([]models.Address, error) {
	var addresses []models.Address

	err := store.db.Select(&addresses, `SELECT * FROM address`)
	if err != nil {
		return nil, err
	}

	return addresses, nil
}

func (store *DbStore) GetProduct() ([]models.Product, error) {
	var products []models.Product

	err := store.db.Select(&products, `SELECT * FROM product`)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (store *DbStore) GetPurchase() ([]models.Purchase, error) {
	var purchases []models.Purchase

	err := store.db.Select(&purchases, `SELECT * FROM purchase`)
	if err != nil {
		return nil, err
	}

	return purchases, nil
}

func (store *DbStore) GetCreator() ([]models.Creator, error) {
	var creator []models.Creator

	err := store.db.Select(&creator, `SELECT * FROM creator`)
	if err != nil {
		return nil, err
	}

	return creator, nil
}

func (store *DbStore) GetCategory() ([]models.Category, error) {
	var category []models.Category

	err := store.db.Select(&category, "SELECT * FROM category")
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (store *DbStore) GetEditor() ([]models.Editor, error) {
	var editor []models.Editor

	err := store.db.Select(&editor, "SELECT * FROM editor")
	if err != nil {
		return nil, err
	}

	return editor, nil
}

func (store *DbStore) GetTag() ([]models.Tag, error) {
	var editor []models.Tag

	err := store.db.Select(&editor, "SELECT * FROM tag")
	if err != nil {
		return nil, err
	}

	return editor, nil
}
func (store *DbStore) GetStatus() ([]models.Status, error) {
	var status []models.Status

	err := store.db.Select(&status, "SELECT * FROM status")
	if err != nil {
		return nil, err
	}

	return status, nil
}

func (store *DbStore) GetPersons() ([]models.Person, error) {
	var persons []models.Person

	err := store.db.Select(&persons, "SELECT id, last_name, first_name, phone_number FROM person")
	if err != nil {
		return nil, err
	}

	return persons, nil
}

func (store *DbStore) GetCredentials() ([]models.Credential, error) {
	var credentials []models.Credential

	err := store.db.Select(&credentials, "SELECT id, email, roles, password, reset_token FROM credential")
	if err != nil {
		return nil, err
	}

	return credentials, nil
}

func (store *DbStore) DeleteAllInfo() error {
	_, err := store.db.Exec(`DELETE FROM picture;`)
	if err != nil {
		return err
	}

	_, err = store.db.Exec(`DELETE FROM creator;`)
	if err != nil {
		return err
	}

	_, err = store.db.Exec(`DELETE FROM pick;`)
	if err != nil {
		return err
	}

	_, err = store.db.Exec(`DELETE FROM comment;`)
	if err != nil {
		return err
	}

	_, err = store.db.Exec(`DELETE FROM product;`)
	if err != nil {
		return err
	}

	_, err = store.db.Exec(`DELETE FROM category;`)
	if err != nil {
		return err
	}
	_, err = store.db.Exec(`DELETE FROM tag;`)
	if err != nil {
		return err
	}

	_, err = store.db.Exec(`DELETE FROM editor;`)
	if err != nil {
		return err
	}

	_, err = store.db.Exec(`DELETE FROM product_creator;`)
	if err != nil {
		return err
	}

	_, err = store.db.Exec(`DELETE FROM comment;`)
	if err != nil {
		return err
	}

	_, err = store.db.Exec(`DELETE FROM purchase;`)
	if err != nil {
		return err
	}

	_, err = store.db.Exec(`DELETE FROM address;`)
	if err != nil {
		return err
	}

	_, err = store.db.Exec(`SET FOREIGN_KEY_CHECKS=0;`)
	if err != nil {
		return err
	}

	_, err = store.db.Exec(`DELETE FROM credential;`)
	if err != nil {
		return err
	}

	_, err = store.db.Exec(`DELETE FROM person;`)
	if err != nil {
		return err
	}

	_, err = store.db.Exec(`SET FOREIGN_KEY_CHECKS=1;`)
	if err != nil {
		return err
	}

	_, err = store.db.Exec(`DELETE FROM status;`)
	if err != nil {
		return err
	}

	return nil
}

func (store *DbStore) Close() error {
	return store.db.Close()
}

func (store *DbStore) Insert(query string) error {
	_, err := store.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
