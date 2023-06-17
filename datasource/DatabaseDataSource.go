package datasource

type QueryLoggerDataSourceInterface interface {
	DataSource
	QueryLoggerInterface
}

type QueryLoggerInterface interface {
	LogQuery(model, query string, duration float32, bind map[string]interface{})
}

type mySQLStructure = struct {
	Model      string   `json:"model"`
	Query      string   `json:"query"`
	Duration   float32  `json:"duration"`
	Connection string   `json:"connection"`
	Tags       []string `json:"tags"`
}

type DatabaseDataSource struct {
	commands      []interface{}
	totalDuration float32
}

func (source *DatabaseDataSource) LogQuery(model, query string, duration float32, bind map[string]interface{}) {
	var tags []string

	if duration > 50 {
		tags = append(tags, "slow")
	} else {
		tags = []string{}
	}

	structure := mySQLStructure{
		Model:      model,
		Query:      query,
		Duration:   duration,
		Connection: "test-connection",
		Tags:       tags,
	}

	source.totalDuration += duration
	source.commands = append(source.commands, &structure)
}

func (source *DatabaseDataSource) Resolve(dataBuffer *DataBuffer) {
	dataBuffer.DatabaseQueries = source.commands
	dataBuffer.DatabaseDuration = source.totalDuration
	dataBuffer.DatabaseQueriesCount = len(dataBuffer.DatabaseQueries)
}
