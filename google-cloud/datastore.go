package main

import (
	"context"
	"log"
	"sync"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"time"

	"cloud.google.com/go/datastore"
)

var wg sync.WaitGroup

type Product struct {
	Origin                  string        `datastore:"origin",bson:"origin"`
	BrowseCallOutSecondary  string        `bson:"browseCallOutSecondary"`
	EndDate                 time.Time     `bson:"endDate"`
	ParentCategoryID        string        `bson:"parentCategoryId"`
	Keywords                []string      `bson:"keywords"`
	Contents                []interface{} `bson:"contents"`
	BrowseCallOutPrimary    string        `bson:"browseCallOutPrimary"`
	Measurements            []interface{} `bson:"measurements"`
	HasVideo                bool          `bson:"hasVideo"`
	IsFeaturedInACatalog    bool          `bson:"isFeaturedInACatalog"`
	MerchandiseClass        string        `bson:"merchandiseClass"`
	ProductID               string        `bson:"productId"`
	ShippingAndReturns      bool          `bson:"shippingAndReturns"`
	StyleNumber             string        `bson:"styleNumber"`
	IPProductName           string        `bson:"ipProductName"`
	PdpCallOutTextSecondary string        `bson:"pdpCallOutTextSecondary"`
	LongDescription         string        `bson:"longDescription"`
	FamilyProducts          []interface{} `bson:"familyProducts"`
	IsActiveCatalog         bool          `bson:"isActiveCatalog"`
	IsGiftCard              bool          `bson:"isGiftCard"`
	WebExclusive            bool          `bson:"webExclusive"`
	Care                    string        `bson:"care"`
	PdpCallOutTextPrimary   string        `bson:"pdpCallOutTextPrimary"`
	DisplayName             string        `bson:"displayName"`
	IsVintage               bool          `bson:"isVintage"`
	IsEgiftCard             bool          `bson:"isEgiftCard"`
	UrbnExclusive           bool          `bson:"urbnExclusive"`
	IPStyleNumber           string        `bson:"ipStyleNumber"`
	IsAvailable             bool          `bson:"isAvailable"`
	ProductSlug             string        `bson:"productSlug"`
	DefaultImage            string        `bson:"defaultImage"`
}

type ProductData struct {
	ID       string `bson:"_id",datastore:"-"`
	Language string `bson:"language"`
	Product  `bson:"product"`
	SiteID   string `bson:"siteId"`
	Links    []struct {
		Locale      string `bson:"locale"`
		ProductSlug string `bson:"productSlug"`
	} `bson:"links"`
	LastModified   string `bson:"lastModified"`
	DocumentSource string `bson:"documentSource"`
	Reviews        struct {
		Count         int     `bson:"count"`
		AverageFit    int     `bson:"averageFit"`
		AverageRating float64 `bson:"averageRating"`
	} `bson:"reviews"`
	Social struct {
		Likes        int `bson:"likes"`
		CustomerPics int `bson:"customerPics"`
	} `bson:"social"`
	CreatedAt string `bson:"createdAt"`
}

func PutProduct(product *ProductData) {
	defer wg.Done()
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, "atg-replacement-poc")
	if err != nil {
		log.Fatalf("Error creating Google Datastore client: %s\n", err.Error())
	}

	key := datastore.IncompleteKey("Products", nil)
	key.Name = product.ProductID
	if _, err := client.Put(ctx, key, product); err != nil {
		log.Printf("Error storing in Google Datastore: %s\n", err.Error())
	}
}

//func GetProductById(id string) *ProductData {
//	ctx := context.Background()
//	client, err := datastore.NewClient(ctx, "atg-replacement-poc")
//	if err != nil {
//		log.Fatalf("Error creating Google Datastore client: %s\n", err.Error())
//	}
//
//	key := datastore.NameKey("Products", "FP-19486000-001", nil)
//	product := &ProductData{}
//	if err := client.Get(ctx, key, product); err != nil {
//		log.Fatalf("Error creating Google Datastore client: %s\n", err.Error())
//	}
//	return product
//}

func GetProductsFromMongo() []ProductData {
	session, err := mgo.Dial("d1shmondb01.phl.ecomm.urbanout.com:27017")
	if err != nil {
		panic(err)
	}

	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	// Collection People
	c := session.DB("CatalogService").C("product")

	// Query All
	var results []ProductData
	err = c.Find(bson.M{}).Limit(1000).All(&results)

	if err != nil {
		panic(err)
	}
	return results
}

func main() {
	products := GetProductsFromMongo()
	for _, p := range products {
		p.ID = ""
		wg.Add(1)
		go PutProduct(&p)
	}
	wg.Wait()
	//
	//product := GetProductById("FP-19486000-000")
	//log.Printf("Product from DCP: %+v\n", product)
}
