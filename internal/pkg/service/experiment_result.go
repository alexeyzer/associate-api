package service

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"sync"

	"github.com/alexeyzer/associate-api/internal/client"
	"github.com/alexeyzer/associate-api/internal/pkg/datastruct"
	"github.com/alexeyzer/associate-api/internal/pkg/repository"
)

type ExperimentResultService interface {
	CreateExperimentResult(ctx context.Context, req []*datastruct.ExperimentResult) ([]*datastruct.ExperimentResultResp, error)
	CalucateExperiment(ctx context.Context, experimentID int64) error
	GetCalculatedResulsts(ctx context.Context, experimentID int64, number, limit int64, names []string) ([]*datastruct.ExperimentResultCalculated, error)
	GetLongestAndWeighetst(ctx context.Context, experimentID int64, words []string) (*datastruct.ResultGetLongestAndWeighetst, error)
	FindPathsInExperiment(ctx context.Context, experimentID int64, word1, word2 string) ([][]string, error)
}

type experimentResultService struct {
	dao   repository.DAO
	redis client.RedisClient
}

// Функция для построения графа на основе данных из ExperimentResultCalculated
func buildGraph(data []*datastruct.ExperimentResultCalculated) map[string]*datastruct.Node {
	graph := make(map[string]*datastruct.Node)

	// Создаем узлы графа и строим ребра
	for _, result := range data {
		currentWord := result.StimuesWord
		nextWord := result.AssotiationWord
		amount := result.Amount

		currentNode, ok := graph[currentWord]
		if !ok {
			currentNode = &datastruct.Node{Word: currentWord}
			graph[currentWord] = currentNode
		}

		nextNode, ok := graph[nextWord]
		if !ok {
			nextNode = &datastruct.Node{Word: nextWord}
			graph[nextWord] = nextNode
		}

		currentNode.Children = append(currentNode.Children, nextNode)
		currentNode.Amount += amount
	}

	return graph
}

type NodeWithPath struct {
	Node *datastruct.Node // Узел графа
	Path []string         // Путь от корня до этого узла
}

// findAllChainsDFS ищет все цепочки ассоциаций с помощью нерекурсивного DFS и возвращает результат в [][]*Node

// findAllChainsDFS ищет все цепочки ассоциаций с помощью нерекурсивного DFS и возвращает результат в [][]*Node
func findAllChainsDFSnode(root *datastruct.Node) [][]*datastruct.Node {
	var allChains [][]*datastruct.Node
	var stack [][]*datastruct.Node

	// Начинаем с корневого узла
	stack = append(stack, []*datastruct.Node{root})

	for len(stack) > 0 {
		// Извлекаем последнюю цепочку из стека
		currentChain := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Текущий узел - последний элемент в цепочке
		currentNode := currentChain[len(currentChain)-1]

		// Проверяем, содержит ли цепочка циклы
		visited := make(map[*datastruct.Node]bool)
		for _, node := range currentChain {
			visited[node] = true
		}

		// Добавляем текущую цепочку в результат, если она состоит более чем из одного узла
		if len(currentChain) > 1 {
			allChains = append(allChains, append([]*datastruct.Node{}, currentChain...))
		}

		// Добавляем детей текущего узла в стек, создавая новые цепочки
		for _, child := range currentNode.Children {
			// Проверяем, чтобы не зациклиться
			if !visited[child] {
				newChain := append([]*datastruct.Node{}, currentChain...)
				newChain = append(newChain, child)
				stack = append(stack, newChain)
			}
		}
	}

	return allChains
}

// findAllChainsDFS ищет все цепочки ассоциаций с помощью нерекурсивного DFS и возвращает результат в [][]string
func findAllChainsDFS(root *datastruct.Node) [][]string {
	var allChains [][]string
	var stack [][]*datastruct.Node
	visited := make(map[*datastruct.Node]bool)

	stack = append(stack, []*datastruct.Node{root})

	for len(stack) > 0 {
		// Извлекаем последний элемент из стека
		currentChain := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Текущий узел - последний элемент в цепочке
		currentNode := currentChain[len(currentChain)-1]

		// Проверяем, посещен ли узел
		if visited[currentNode] {
			continue
		}
		visited[currentNode] = true

		// Добавляем текущую цепочку в результат, если она состоит более чем из одного узла
		if len(currentChain) > 1 {
			var chainWords []string
			for _, node := range currentChain {
				chainWords = append(chainWords, node.Word)
			}
			allChains = append(allChains, chainWords)
		}

		// Добавляем детей текущего узла в стек, создавая новые цепочки
		for _, child := range currentNode.Children {
			newChain := append([]*datastruct.Node{}, currentChain...)
			newChain = append(newChain, child)
			stack = append(stack, newChain)
		}
	}

	return allChains
}

// Структура для задачи поиска пути
type task struct {
	node *datastruct.Node
	path []string
}

func uniqueSecond(chains [][]string) [][]string {
	wordMap := make(map[string]struct{})
	newChain := make([][]string, 0, len(chains))

	for _, chain := range chains {
		if len(chain) > 2 {
			if _, ok := wordMap[chain[1]]; ok {
				continue
			}
			wordMap[chain[1]] = struct{}{}
			newChain = append(newChain, chain)
		} else {
			newChain = append(newChain, chain)
		}
	}
	return newChain
}

// findAllPaths принимает массив слов и находит все пути в графе для каждого слова
func findAllPaths(words []string, graph map[string]*datastruct.Node) map[string][][]string {
	results := make(map[string][][]string)

	for _, word := range words {
		if root, exists := graph[word]; exists {
			chains := findAllChainsDFS(root)
			print(len(chains))
			sortChainsByLength(&chains)
			results[word] = uniqueSecond(chains)
		} else {
			results[word] = nil // Если слово не найдено в графе, результат будет nil
		}
	}

	return results
}

// Поиск всех путей между двумя словами с помощью DFS
func findAllPathStartEnd(graph map[string]*datastruct.Node, start, end string) [][]string {
	var paths [][]string
	var wg sync.WaitGroup
	taskChan := make(chan task)
	resultChan := make(chan []string)
	numWorkers := 4 // Количество рабочих горутин

	// Если начального узла нет в графе, пути не существует
	startNode, startExists := graph[start]
	_, endExists := graph[end]
	if !startExists || !endExists {
		return paths
	}

	// Запуск рабочих горутин
	for i := 0; i < numWorkers; i++ {
		go func() {
			for t := range taskChan {
				currentNode := t.node
				currentPath := t.path

				// Если мы достигли конечного слова, добавляем путь к результатам
				if currentNode.Word == end {
					resultChan <- currentPath
				} else {
					// Проходим по всем соседям текущего узла
					for _, neighbor := range currentNode.Children {
						// Проверяем, что сосед ещё не находится в текущем пути, чтобы избежать циклов
						if contains(currentPath, neighbor.Word) {
							continue
						}

						// Создаем новую задачу и добавляем ее в канал задач
						newPath := append([]string{}, currentPath...)
						newPath = append(newPath, neighbor.Word)
						wg.Add(1)
						taskChan <- task{node: neighbor, path: newPath}
					}
				}
				wg.Done()
			}
		}()
	}

	// Инициализация начальной задачи
	wg.Add(1)
	taskChan <- task{node: startNode, path: []string{start}}

	// Запускаем горутину для закрытия канала задач после завершения всех задач
	go func() {
		wg.Wait()
		close(taskChan)
		close(resultChan)
	}()

	// Сбор результатов из канала
	for path := range resultChan {
		paths = append(paths, path)
	}

	return paths
}

// Проверяет, содержится ли слово в пути
func contains(path []string, word string) bool {
	for _, w := range path {
		if w == word {
			return true
		}
	}
	return false
}

func trimChains(chains [][]string, word2 string) [][]string {
	var trimmedChains [][]string
	for _, chain := range chains {
		for i, word := range chain {
			if word == word2 {
				trimmedChains = append(trimmedChains, chain[:i+1])
				break
			}
		}
	}
	return trimmedChains
}

func removeDuplicateChains(chains [][]string) [][]string {
	seen := make(map[string]struct{}) // используем карту для отслеживания уже встреченных цепочек
	var uniqueChains [][]string
	for _, chain := range chains {
		// преобразуем цепочку в строку для использования в качестве ключа карты
		chainString := fmt.Sprintf("%v", chain)
		// если цепочка не была ранее встречена, добавляем ее в список уникальных цепочек
		if _, ok := seen[chainString]; !ok {
			uniqueChains = append(uniqueChains, chain)
			seen[chainString] = struct{}{}
		}
	}
	return uniqueChains
}

func (e *experimentResultService) FindPathsInExperiment(ctx context.Context, experimentID int64, word1, word2 string) ([][]string, error) {
	var empty [][]string

	results, err := e.dao.ExperimentResultCalculatedQuery().GetCalculatedResulsts(ctx, 0, []int64{experimentID}, 0, 0, []string{})
	if err != nil {
		return nil, err
	}
	graph := buildGraph(results)
	if root, exists := graph[word1]; exists {
		print("found 1 word")
		chains := findAllChainsDFS(root)
		print("ifiafaf")
		return removeDuplicateChains(trimChains(chains, word2)), nil
	}
	return empty, nil
}

func (e *experimentResultService) GetLongestAndWeighetst(ctx context.Context, experimentID int64, words []string) (*datastruct.ResultGetLongestAndWeighetst, error) {
	result := &datastruct.ResultGetLongestAndWeighetst{
		TopLongest: make([][]string, 0),
		TopWeight:  make([][]string, 0),
	}

	marcheled, err := e.redis.GetValue(ctx, fmt.Sprintf("longest-%d", experimentID))
	if err != nil {
		return nil, err
	}
	if marcheled != nil {
		err = json.Unmarshal(marcheled, result)
		if err != nil {
			return nil, err
		}
		if err == nil {
			return result, nil
		}
	}

	results, err := e.dao.ExperimentResultCalculatedQuery().GetCalculatedResulsts(ctx, 0, []int64{experimentID}, 0, 0, []string{})
	if err != nil {
		return nil, err
	}

	fmt.Println("graphn?")
	// Строим граф на основе данных из ExperimentResultCalculated
	graph := buildGraph(results)
	fmt.Println("builded?")

	// Ищем все цепочки ассоциаций с помощью DFS
	var allChains [][]string
	fmt.Println("builded and?")
	for _, node := range graph {
		allChains = append(allChains, findAllChainsDFS(node)...)
	}

	fmt.Println("allchains?")

	// Сортируем цепочки по длине
	sortChainsByLength(&allChains)
	fmt.Println("here")
	wordMap := make(map[string]struct{}, 0)

	for _, chain := range allChains {
		if _, ok := wordMap[chain[0]]; ok {
			continue
		}
		wordMap[chain[0]] = struct{}{}
		result.TopLongest = append(result.TopLongest, chain)
	}
	fmt.Println("longest?")

	allPaths := findAllPaths(words, graph)
	result.FindWolrds = allPaths

	// // Ищем все цепочки ассоциаций с помощью DFS
	// var allChainsNodes [][]*datastruct.Node
	// for _, node := range graph {
	// 	allChainsNodes = append(allChainsNodes, findAllChainsDFSnode(node)...)
	// }

	// fmt.Println("all chain weighdet?")

	// topChains := findTopWeightedChains(&allChainsNodes)

	// for i, chain := range topChains {
	// 	if i >= 5 {
	// 		break
	// 	}
	// 	result.TopWeight = append(result.TopWeight, nodesToStrings(chain))
	// }

	marshaled, err := json.Marshal(result)
	if err != nil {
		return nil, err
	}

	err = e.redis.SetValue(ctx, fmt.Sprintf("longest%d", experimentID), marshaled)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func nodesToStrings(nodes []*datastruct.Node) []string {
	strings := make([]string, len(nodes))
	for i, node := range nodes {
		strings[i] = node.Word
	}
	return strings
}

// Функция для нахождения наиболее весомых цепочек
func findTopWeightedChains(allChains *[][]*datastruct.Node) [][]*datastruct.Node {
	sortChainsByWeight(allChains)
	return (*allChains)[:5]
}

// Функция для сортировки цепочек по сумме весов в порядке убывания
func sortChainsByWeight(chains *[][]*datastruct.Node) {
	sort.Slice(*chains, func(i, j int) bool {
		return getChainWeight((*chains)[i]) > getChainWeight((*chains)[j])
	})
}

// Функция для вычисления суммы весов цепочки
func getChainWeight(chain []*datastruct.Node) int64 {
	var weight int64
	for _, node := range chain {
		weight += node.Amount
	}
	return weight
}

// Функция для сортировки цепочек по длине в порядке убывания
func sortChainsByLength(chains *[][]string) {
	sort.Slice(*chains, func(i, j int) bool {
		return len((*chains)[i]) > len((*chains)[j])
	})
}

func (e *experimentResultService) GetCalculatedResulsts(ctx context.Context, experimentID int64, number, limit int64, names []string) ([]*datastruct.ExperimentResultCalculated, error) {
	return e.dao.ExperimentResultCalculatedQuery().GetCalculatedResulsts(ctx, 0, []int64{experimentID}, number, limit, names)
}

func (e *experimentResultService) CalucateExperiment(ctx context.Context, experimentID int64) error {
	results, err := e.dao.ExperimentResultQuery().List(ctx, 0, []int64{experimentID}, 0, 0, []string{})
	if err != nil {
		return err
	}
	grahp := make(map[int64]map[int64]int64, len(results))

	for _, r := range results {
		_, ok := grahp[r.StimusWordID]
		if !ok {
			grahp[r.StimusWordID] = make(map[int64]int64, 30)
		}
		grahp[r.StimusWordID][r.AssotiationWordID] = grahp[r.StimusWordID][r.AssotiationWordID] + 1
	}
	result := make([]*datastruct.ExperimentResultCalculated, 0, len(grahp))

	for key, mapp := range grahp {
		for associate, amount := range mapp {
			result = append(result, &datastruct.ExperimentResultCalculated{
				ExperimentID:      experimentID,
				StimusWordID:      key,
				AssotiationWordID: associate,
				Amount:            amount,
			})
		}
	}
	return e.dao.ExperimentResultCalculatedQuery().BatchCreate(ctx, result)
}

func (e *experimentResultService) CreateExperimentResult(ctx context.Context, req []*datastruct.ExperimentResult) ([]*datastruct.ExperimentResultResp, error) {
	for _, r := range req {
		res, err := e.dao.AssociateWordQuery().GetByName(ctx, r.AssotiationWord)
		if err == nil {
			r.AssotiationWordID = res.ID
			continue
		}

		res, err = e.dao.AssociateWordQuery().Create(ctx, datastruct.AssociateWord{
			Name: r.AssotiationWord,
		})
		if err != nil {
			return nil, err
		}
		r.AssotiationWordID = res.ID
	}

	repositoryes, err := e.dao.ExperimentResultQuery().BatchCreate(ctx, req)
	if err != nil {
		return nil, err
	}

	err = e.CalucateExperiment(ctx, req[0].ExperimentID)
	if err != nil {
		return nil, err
	}
	return repositoryes, nil
}

func NewExperimentResultService(dao repository.DAO, redis client.RedisClient) ExperimentResultService {
	return &experimentResultService{dao: dao, redis: redis}
}
