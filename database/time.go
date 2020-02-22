package database

import (
	"fmt"
	"time"
)

type TimeGroup interface {
}

func (it *Db) ParseTime(t string) (time.Time, error) {
	switch it.Type {
	case "mysql":
		return time.Parse("2006-01-02 15:04:05", t)
	case "postgres":
		return time.Parse("2006-01-02T15:04:05Z", t)
	default:
		return time.Parse("2006-01-02 15:04:05", t)
	}
}

func (it *Db) FormatTime(t time.Time) string {
	switch it.Type {
	case "mysql":
		return t.UTC().Format("2006-01-02 15:04:05")
	case "postgres":
		return t.UTC().Format("2006-01-02 15:04:05.999999999")
	default:
		return t.UTC().Format("2006-01-02 15:04:05")
	}
}

func (it *Db) SelectByTime(increment string) string {
	switch it.Type {
	case "mysql":
		return fmt.Sprintf("CONCAT(date_format(created_at, '%s')) AS timeframe", it.correctTimestamp(increment))
	case "postgres":
		return fmt.Sprintf("date_trunc('%s', created_at) AS timeframe", increment)
	default:
		return fmt.Sprintf("strftime('%s', created_at, 'utc') as timeframe", it.correctTimestamp(increment))
	}
}

func (it *Db) correctTimestamp(increment string) string {
	var timestamper string
	switch increment {
	case "second":
		timestamper = "%Y-%m-%d %H:%M:%S"
	case "minute":
		timestamper = "%Y-%m-%d %H:%M:00"
	case "hour":
		timestamper = "%Y-%m-%d %H:00:00"
	case "day":
		timestamper = "%Y-%m-%d 00:00:00"
	case "month":
		timestamper = "%Y-%m-01 00:00:00"
	case "year":
		timestamper = "%Y-01-01 00:00:00"
	default:
		timestamper = "%Y-%m-%d 00:00:00"
	}

	switch it.Type {
	case "mysql":
	case "second":
		timestamper = "%Y-%m-%d %H:%i:%S"
	case "minute":
		timestamper = "%Y-%m-%d %H:%i:00"
	}

	return timestamper
}
