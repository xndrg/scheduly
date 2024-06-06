package mau

type Scraper struct {
	host string
}

func New() *Scraper {
	return &Scraper{
		host: "https://www.mauniver.ru/student/timetable/new/schedule.php",
	}
}

func (s *Scraper) GetTodaySchedule(groupName string) (map[string]string, error) {
	return nil, nil
}

func (s *Scraper) GetTommorowSchedule(groupName string) (map[string]string, error) {
	return nil, nil
}
