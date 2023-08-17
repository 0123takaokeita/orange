package tasks

import (
	"fmt"
	// "io/ioutil"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

func ListUp(){
  fmt.Println("### Rake Taks List Up ###")
	rakeTasks, err := getRakeTasks()
	if err != nil {
		log.Fatalf("Failed to get Rake tasks: %s", err)
	}

	fmt.Println("Rake Tasks:")
	for _, task := range rakeTasks {
		fmt.Println(task)
	}
}

func getRakeTasks() ([]string, error) {
	// Rakeタスクを取得するためのコマンドを実行
	cmd := exec.Command("rake", "--tasks")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to execute Rake command: %s", err)
	}

	// 出力結果からタスク名を抽出
	taskLines := extractTaskLines(string(output))
	tasks := extractTasks(taskLines)

	return tasks, nil
}

func extractTaskLines(output string) []string {
	// タスク行を正規表現で抽出
	re := regexp.MustCompile(`\s*rake\s+([^\s]+)`)
	taskLines := re.FindAllStringSubmatch(output, -1)

	result := make([]string, len(taskLines))
	for i, line := range taskLines {
		result[i] = line[1]
	}

	return result
}

func extractTasks(taskLines []string) []string {
	// タスク名を抽出
	var tasks []string
	for _, line := range taskLines {
		task := strings.TrimSpace(line)
		if task != "" {
			tasks = append(tasks, task)
		}
	}

	return tasks
}
