//go:build !solution

package hogwarts

import "fmt"

// GetCourseList выполняет топологическую сортировку курсов с учетом предварительных требований
func GetCourseList(prereqs map[string][]string) []string {
	return topologicalSort(prereqs)
}

// Константы для отслеживания состояния вершин при обходе
const (
	NOT_VISITED = 0 // Вершина еще не обрабатывалась
	IN_PROGRESS = 1 // Вершина в процессе обработки (обнаружение цикла)
	VISITED     = 2 // Вершина полностью обработана
)

// topologicalSort выполняет топологическую сортировку направленного графа
func topologicalSort(graph map[string][]string) []string {
	// Карта для отслеживания состояния каждой вершины
	visited := make(map[string]int)
	// Результирующий список курсов в правильном порядке
	var result []string

	// Собираем все вершины графа
	// Некоторые курсы могут быть только в предварительных требованиях,
	// но не быть ключами в карте, поэтому нужно учесть все
	allVertices := make(map[string]bool)
	for vertex, deps := range graph {
		allVertices[vertex] = true
		for _, dep := range deps {
			allVertices[dep] = true
		}
	}

	// Последовательно обрабатываем все вершины
	for vertex := range allVertices {
		if visited[vertex] == NOT_VISITED {
			// Если вершина не посещена, запускаем ее обработку
			if err := visit(graph, vertex, visited, &result); err != nil {
				panic(err) // При обнаружении цикла прекращаем выполнение
			}
		}
	}
	return result
}

// visit рекурсивно обрабатывает вершину и ее зависимости
func visit(graph map[string][]string, vertex string, visited map[string]int, result *[]string) error {
	// Обнаружен цикл - вершина уже в процессе обработки
	if visited[vertex] == IN_PROGRESS {
		return fmt.Errorf("циклическая зависимость")
	}

	// Вершина уже полностью обработана - выходим
	if visited[vertex] == VISITED {
		return nil
	}

	// Помечаем вершину как обрабатываемую
	visited[vertex] = IN_PROGRESS

	// Рекурсивно обрабатываем ВСЕ зависимости текущей вершины
	// В графе prereqs по ключу vertex хранятся курсы, которые должны быть пройдены ДО vertex
	for _, dep := range graph[vertex] {
		if err := visit(graph, dep, visited, result); err != nil {
			return err
		}
	}

	// После обработки всех зависимостей помечаем вершину как завершенную
	visited[vertex] = VISITED
	// Добавляем вершину в результат
	// На этом этапе все необходимые курсы уже добавлены в result
	*result = append(*result, vertex)
	return nil
}
