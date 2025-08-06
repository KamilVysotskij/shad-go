//go:build !solution

package hotelbusiness

import "sort"

type Guest struct {
	CheckInDate  int
	CheckOutDate int
}

type Load struct {
	StartDate  int
	GuestCount int
}

func ComputeLoad(guests []Guest) []Load {
	// Собираем все дни с изменениями
	var result []Load
	dateChangeMap := make(map[int]int)
	for _, guest := range guests {
		dateChangeMap[guest.CheckInDate] += 1
		dateChangeMap[guest.CheckOutDate] -= 1
	}
	// Собираем все дни
	var dates []int
	for date := range dateChangeMap {
		dates = append(dates, date)
	}
	sort.Ints(dates)
	// Проходим по отсортированным дням и фиксируем кол-во гостей
	currentGuests := 0

	// Проходим по отсортированным датам и вычисляем текущее количество гостей
	for _, date := range dates {
		// Изменяем текущее количество гостей
		currentGuests += dateChangeMap[date]
		// Добавляем запись в результат, если количество гостей изменилось
		if len(result) == 0 || result[len(result)-1].GuestCount != currentGuests {
			result = append(result, Load{
				StartDate:  date,
				GuestCount: currentGuests,
			})
		}
	}
	return result
}
