package datasource

import "time"

type QueryLoggerDataSourceInterface interface {
	DataSource
	QueryLoggerInterface
}

type QueryLoggerInterface interface {
	LogQuery(model, query string, duration float32, trace []map[string]interface{})
}

type mySQLStructure = struct {
	Model      string                   `json:"model"`
	Query      string                   `json:"query"`
	Duration   float32                  `json:"duration"`
	Connection string                   `json:"connection"`
	Tags       []string                 `json:"tags"`
	Trace      []map[string]interface{} `json:"trace"`
	Time       float64                  `json:"time"`
}

type DatabaseDataSource struct {
	commands      []interface{}
	totalDuration float32
}

func (source *DatabaseDataSource) LogQuery(model, query string, duration float32, trace []map[string]interface{}) {
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
		Trace:      trace,
		Time:       MicroTime(time.Now()),
	}

	source.totalDuration += duration
	source.commands = append(source.commands, &structure)
}

func (source *DatabaseDataSource) Resolve(dataBuffer *DataBuffer) {
	dataBuffer.DatabaseQueries = source.commands
	dataBuffer.DatabaseDuration = source.totalDuration
	dataBuffer.DatabaseQueriesCount = len(dataBuffer.DatabaseQueries)
}
