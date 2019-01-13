package gin_example

type Config struct {
	// Common App variables
	AppName         string `split_words:"true"  default:"LocalOrderNumberGenerator"`
	GitHash         string `split_words:"true"  default:"Github-Hash"`
	Branch          string `split_words:"true"  default:"Github-Branch"`
	BuildNumber     string `split_words:"true"  default:"Jenkins-BuildNumber"`
	Environment     string `split_words:"true"  default:"LOCAL"`
	NewRelicLicense string `split_words:"true"`

	// Mongo Info
	MongoServerUrls string `split_words:"true"`
}
