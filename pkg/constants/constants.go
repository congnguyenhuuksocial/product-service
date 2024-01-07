package constants

const (
	// EnvLogLevel is the log level for the application
	EnvLogLevel = "LOG_LEVEL"
	// EnvLogFormat is the log format for the application
	EnvLogFormat = "LOG_FORMAT"
	// EnvLogOutput is the log output for the application
	// EnvDbHost is the database host
	EnvDbHost = "DB_HOST"
	// EnvDbPort is the database port
	EnvDbPort = "DB_PORT"
	// EnvDbUser is the database user
	EnvDbUsername = "DB_USERNAME"
	// EnvDbPassword is the database password
	EnvDbPassword = "DB_PASSWORD"
	// EnvDbName is the database name
	EnvDbName = "DB_NAME"
	// EnvDbDialect is the database dialect
	EnvDbDialect = "DB_DIALECT"
	// EnvDbLog is the database log
	EnvDbConnMaxLifetimeSecond = "CONN_MAX_LIFETIME_SECOND"
	// EnvMaxIdleConns is the database max idle connections
	EnvMaxIdleConns = "MAX_IDLE_CONNS"
	// EnvMaxOpenConns is the database max open connections
	EnvMaxOpenConns = "MAX_OPEN_CONNS"
	// EnvDbDriver is the database driver
	EnvDbDriver = "DB_DRIVER"
	// EnvDbInitMaxRetryNumber is the database init max retry number
	EnvDbInitMaxRetryNumber = "DB_INIT_MAX_RETRY_NUMBER"
	// EnvDbInitRetryIntervalSecond is the database init retry interval second
	EnvDbInitRetryIntervalSecond = "DB_INIT_RETRY_INTERVAL_SECOND"
	// EnvGrpcHost is the grpc host
	EnvGrpcHost = "GRPC_HOST"
	// EnvGrpcPort is the grpc port
	EnvGrpcPort = "GRPC_PORT"
	// EnvServerHost is the server host
	EnvServerHost = "SERVER_HOST"
	// EnvServerPort is the server port
	EnvServerPort = "SERVER_PORT"
	// EnvServerEnv is the server env
	EnvServerEnv = "SERVER_ENV"
	// EnvSqsRegion is the sqs region
	EnvSqsRegion = "SQS_REGION"
	// EnvSqsName is the sqs name
	EnvSqsName = "SQS_NAME"
	// EnvSqsMaxNumberOfMessage is the sqs max number of message
	EnvSqsMaxNumberOfMessage = "SQS_MAX_NUMBER_OF_MESSAGE"
	// EnvSqsWaitTimeSecond is the sqs wait time second
	EnvSqsWaitTimeSecond = "SQS_WAIT_TIME_SECOND"
	// EnvSqsQueueUrl is the sqs queue url
	EnvSqsQueueUrl = "SQS_QUEUE_URL"
	// EnvSqsAccessKey is the sqs access key
	EnvSqsAccessKey = "SQS_ACCESS_KEY"
	// EnvSqsSecretKey is the sqs secret key
	EnvSqsSecretKey         = "SQS_SECRET_KEY"
	EnvSnsFlavorSearchTopic = "SNS_FLAVOR_SEARCH_TOPIC"
	// EnvSearchUrl is the search url
	EnvSearchUrl = "SEARCH_URL"
	// EnvSearchUsername is the search username
	EnvSearchUsername = "SEARCH_USERNAME"
	// EnvSearchPassword is the search password
	EnvSearchPassword = "SEARCH_PASSWORD"
	// EnvSearchMaxRetries is the search max retries
	EnvSearchMaxRetries = "SEARCH_MAX_RETRIES"
	// EnvSearchFlavorIndex is the search flavor index
	EnvSearchFlavorIndex = "SEARCH_FLAVOR_INDEX"

	// EnvSearchProfileIndex is the search profile index
	EnvSearchProfileIndex = "SEARCH_PROFILE_INDEX"

	// EnvGrpcProfileClientUrl is the grpc profile client url
	EnvGrpcProfileClientUrl = "GRPC_PROFILE_CLIENT_URL"

	// EnvBusinessFeatureMaxDays is the business feature max days
	EnvBusinessFeatureMaxDays = "BUSINESS_FEATURE_MAX_DAYS"
	// EnvBusinessFeatureMaxRecords is the business feature max records
	EnvBusinessFeatureMaxRecords = "BUSINESS_FEATURE_MAX_RECORDS"
	// EnvBusinessWatchesScore is the business watches score
	EnvBusinessWatchesScore = "BUSINESS_WATCHES_SCORE"
	// EnvBusinessLikesScore is the business likes score
	EnvBusinessLikesScore = "BUSINESS_LIKES_SCORE"
	// EnvBusinessDislikesScore is the business dislikes score
	EnvBusinessDislikesScore = "BUSINESS_DISLIKES_SCORE"
	// EnvBusinessBookmarksScore is the business bookmarks score
	EnvBusinessBookmarksScore = "BUSINESS_BOOKMARKS_SCORE"
	// EnvBusinessSharesScore is the business shares score
	EnvBusinessSharesScore = "BUSINESS_SHARES_SCORE"
	// EnvGrpcPostClientUrl is the grpc post client url
	EnvGrpcPostClientUrl = "GRPC_POST_CLIENT_URL"

	// EnvGeoNearByRadius is the geo near by radius
	EnvGeoNearByRadius = "GEO_NEAR_BY_RADIUS"

	// EnvSearchPostIndex is the search post index
	EnvSearchPostIndex = "SEARCH_POST_INDEX"
	// RedisHost is the redis host
	EnvRedisHost = "REDIS_HOST"
	// RedisPort is the redis port
	EnvRedisPort = "REDIS_PORT"
	// RedisURL is the redis url
	EnvRedisUri = "REDIS_URI"
	// EnvRedisTTL is the redis ttl
	EnvRedisTTL = "REDIS_TTL"
	// Cdn is the cdn
	EnvServerCdn = "CDN"

	EnvLogPrefix  = "LOG_PREFIX"
	EnvLogDevMode = "LOG_DEV_MODE"
	EnvLogEncoder = "LOG_ENCODER"
)

// Database dialects
const (
	Mysql    = "mysql"
	Postgres = "postgres"
)

// Logger constants
const (
	JSON        = "json"
	CONSOLE     = "console"
	Production  = "production"
	Development = "development"
)

const (
	EnvKafkaUri   = "KAFKA_URI"
	EnvKafkaTopic = "KAFKA_TOPIC"

	EnvKafkaGroup = "KAFKA_GROUP"
)
