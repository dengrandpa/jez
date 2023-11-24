package timejez

type TimeFormat string

const (
	YYYYMMDDHHMMSS  TimeFormat = "2006-01-02 15:04:05"
	YYYYMMDDHHMM    TimeFormat = "2006-01-02 15:04"
	YYYYMMDDHH      TimeFormat = "2006-01-02 15"
	YYYYMMDD        TimeFormat = "2006-01-02"
	YYYYMM          TimeFormat = "2006-01"
	MMDD            TimeFormat = "01-02"
	DDMMYYHHMMSS    TimeFormat = "02-01-06 15:04:05"
	YYYYMMDDHHMMSS2 TimeFormat = "2006/01/02 15:04:05"
	YYYYMMDDHHMM2   TimeFormat = "2006/01/02 15:04"
	YYYYMMDDHH2     TimeFormat = "2006/01/02 15"
	YYYYMMDD2       TimeFormat = "2006/01/02"
	YYYYMM2         TimeFormat = "2006/01"
	MMDD2           TimeFormat = "01/02"
	DDMMYYHHMMSS2   TimeFormat = "02/01/06 15:04:05"
	YYYY            TimeFormat = "2006"
	MM              TimeFormat = "01"
	HHMMSS          TimeFormat = "15:04:05"
	HHMM            TimeFormat = "15:04"
	MMSS            TimeFormat = "04:05"
)

func (t TimeFormat) string() string {
	return string(t)
}

const (
	secondsInDay    = 86400
	secondsInHour   = 3600
	secondsInMinute = 60
)
