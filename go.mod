module github.com/NicklasWallgren/go-template

go 1.19

require (
	github.com/gin-gonic/gin v1.8.1
	github.com/spf13/cobra v1.4.0
	github.com/wagslane/go-rabbitmq v0.8.1
	go.uber.org/fx v1.19.0
	go.uber.org/zap v1.23.0
	gopkg.in/DataDog/dd-trace-go.v1 v1.38.1
	gorm.io/gorm v1.23.5
)

require (
	github.com/NicklasWallgren/sqlTemplate v0.11.0
	github.com/avast/retry-go v3.0.0+incompatible
	github.com/brianvoe/gofakeit/v6 v6.16.0
	github.com/gin-contrib/zap v0.0.2
	github.com/gkampitakis/go-snaps v0.3.2
	github.com/go-playground/validator/v10 v10.10.1
	github.com/go-sql-driver/mysql v1.6.0
	github.com/google/uuid v1.3.0
	github.com/iamolegga/enviper v1.4.0
	github.com/jackc/pgconn v1.12.1
	github.com/lib/pq v1.10.4
	github.com/mariomac/gostream v0.8.1
	github.com/pressly/goose/v3 v3.5.3
	github.com/rs/cors/wrapper/gin v0.0.0-20220223021805-a4a5ce87d5a2
	github.com/spf13/viper v1.10.1
	github.com/stretchr/testify v1.8.0
	github.com/swaggo/files v0.0.0-20220610200504-28940afbdbfe
	github.com/swaggo/gin-swagger v1.5.1
	github.com/swaggo/swag v1.8.4
	gorm.io/driver/mysql v1.3.3
	gorm.io/driver/postgres v1.3.7
	moul.io/zapgorm2 v1.1.3
)

replace github.com/wagslane/go-rabbitmq v0.8.1 => github.com/NicklasWallgren/go-rabbitmq v0.10.1-0.20220527095640-5ed506244369

require (
	github.com/DataDog/datadog-agent/pkg/obfuscate v0.0.0-20211129110424-6491aa3bf583 // indirect
	github.com/DataDog/datadog-go v4.8.2+incompatible // indirect
	github.com/DataDog/datadog-go/v5 v5.0.2 // indirect
	github.com/DataDog/sketches-go v1.0.0 // indirect
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/Microsoft/go-winio v0.5.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgraph-io/ristretto v0.1.0 // indirect
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/fsnotify/fsnotify v1.5.1 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gkampitakis/ciinfo v0.1.0 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.20.0 // indirect
	github.com/go-openapi/spec v0.20.5 // indirect
	github.com/go-openapi/swag v0.21.1 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/goccy/go-json v0.9.7 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.11.0 // indirect
	github.com/jackc/pgx/v4 v4.16.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/magiconair/properties v1.8.5 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mitchellh/mapstructure v1.4.3 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml v1.9.4 // indirect
	github.com/pelletier/go-toml/v2 v2.0.1 // indirect
	github.com/philhofer/fwd v1.1.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rabbitmq/amqp091-go v1.3.4 // indirect
	github.com/rogpeppe/go-internal v1.8.1 // indirect
	github.com/rs/cors v1.8.1 // indirect
	github.com/sergi/go-diff v1.2.0 // indirect
	github.com/spf13/afero v1.6.0 // indirect
	github.com/spf13/cast v1.4.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/objx v0.4.0 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	github.com/tinylib/msgp v1.1.2 // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/dig v1.16.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	golang.org/x/crypto v0.0.0-20220408190544-5352b0902921 // indirect
	golang.org/x/exp v0.0.0-20220328175248-053ad81199eb // indirect
	golang.org/x/net v0.0.0-20220425223048-2871e0cb64e4 // indirect
	golang.org/x/sys v0.0.0-20220408201424-a24fb2fb8a0f // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/time v0.0.0-20211116232009-f0f3c7e86c11 // indirect
	golang.org/x/tools v0.1.10 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
	gopkg.in/ini.v1 v1.66.2 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
