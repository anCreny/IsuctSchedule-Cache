package repo

import (
	"fmt"
	"github.com/anCreny/IsuctSchedule-Packages/structs"
	"github.com/restream/reindexer/v3"
	"main/config"
	"slices"
	"strings"
	"time"
	_ "time/tzdata"
)

const (
	Location = "Europe/Moscow"
)

func CheckTeacher(name string) bool {
	names := GetNames()

	compName := strings.Join(strings.Split(name, "-"), " ")

	return slices.Contains(names.Names, compName)
}

func CheckGroup(groupNumber string) bool {
	_, found := r.Rx.Query(config.Cfg.RxCfg.Namespaces.Groups).
		Where("holder", reindexer.EQ, groupNumber).Get()
	return found
}

func GetGroupDay(groupNumber string, offset int) (structs.Day, error) {
	return getDay(groupNumber, config.Cfg.RxCfg.Namespaces.Groups, offset)
}

func GetTeacherDay(teacherName string, offset int) (structs.Day, error) {
	return getDay(teacherName, config.Cfg.RxCfg.Namespaces.Teachers, offset)
}

func getDay(holder, namespace string, offset int) (day structs.Day, err error) {
	weekDate := getWeekDate(offset)

	if weekDate.Weekday == 0 {
		day.Week = weekDate.Week
		return
	}

	q := r.Rx.Query(namespace).
		Where("holder", reindexer.EQ, holder)

	iterator, found := q.Get()
	if !found {
		err = fmt.Errorf("couldn't find %s", holder)
		return
	}

	timetable := iterator.(*structs.Timetable)

	for _, tDay := range timetable.Days {
		if tDay.Weekday == weekDate.Weekday &&
			tDay.Week == weekDate.Week {
			day = tDay
		}
	}

	return
}

func getWeekDate(offset int) WeekDate {
	weekDate := WeekDate{}

	location, err := time.LoadLocation(Location)
	if err != nil {
		panic(err)
	}

	day := time.Now().In(location).Add(24 * time.Duration(offset) * time.Hour)

	weekDate.Weekday = int(day.Weekday())
	_, weekN := day.ISOWeek()
	weekDate.Week = 3 - (weekN%2 + 1)

	return weekDate
}

func GetGroup(groupNumber string) (structs.Timetable, bool) {
	return getTimetable(groupNumber, config.Cfg.RxCfg.Namespaces.Groups)
}

func GetTeacher(teacherName string) (structs.Timetable, bool) {
	return getTimetable(teacherName, config.Cfg.RxCfg.Namespaces.Teachers)
}

func GetNames() structs.TeachersNames {
	item, _ := r.Rx.Query(config.Cfg.RxCfg.Namespaces.Names).Get()

	names := item.(*structs.TeachersNames)

	return *names
}

func GetCommonTeachers(name string) []string {
	namesArr := GetNames()
	var commonNames []string
	for _, teacherName := range namesArr.Names {
		if len(commonNames) == 5 {
			break
		}
		if strings.HasPrefix(strings.ToLower(teacherName), name) {
			commonNames = append(commonNames, teacherName)
		}
	}
	return commonNames
}

func getTimetable(holder, namespace string) (structs.Timetable, bool) {
	item, found := r.Rx.Query(namespace).
		Where("Holder", reindexer.EQ, holder).Get()

	if !found {
		return structs.Timetable{}, false
	}

	timetable := item.(*structs.Timetable)
	return *timetable, true
}
