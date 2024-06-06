package scraper

type Scraper interface {
	GetTodaySchedule(groupName string) (map[string]string, error)
	GetTommorowSchedule(groupName string) (map[string]string, error)
}
