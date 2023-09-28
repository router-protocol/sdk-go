package types

import "sort"

type OracleTask struct {
	OracleScriptID uint64
	Symbols        []string
}

func CalculateGas(base, each, n uint64) uint64 {
	return base + each*n
}

func ComputeOracleTasks(symbols []SymbolRequest, blockHeight int64) []OracleTask {
	symbolMap := make(map[uint64][]string)
	// stays the same
	for _, symbol := range symbols {
		if symbol.BlockInterval != 0 && blockHeight%int64(symbol.BlockInterval) == 0 {
			symbolMap[symbol.OracleScriptID] = append(symbolMap[symbol.OracleScriptID], symbol.Symbol)
		}
	}
	tasks := make([]OracleTask, 0, len(symbolMap))
	// directly create tasks
	for oracleScriptID, symbols := range symbolMap {
		tasks = append(tasks, OracleTask{
			OracleScriptID: oracleScriptID,
			Symbols:        symbols,
		})
	}
	// sort them
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].OracleScriptID < tasks[j].OracleScriptID
	})
	return tasks
}
