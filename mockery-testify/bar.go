package mockery_testify

type DataAccessLayer interface {
	Insert(collectionName string, docs ...interface{}) error
}

type Thing struct{}

func Bar(data DataAccessLayer) Thing {
	doc := Thing{}

	data.Insert("foo", doc)
	return doc
}
