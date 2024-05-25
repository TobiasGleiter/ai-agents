package main

import(
	"fmt"
	"encoding/json"
	"log"

	"github.com/TobiasGleiter/ai-agents/pkg/llms/ollama"
)

type Task struct {
    Todo string `json:"todo"`
    Done bool   `json:"done"`
}

type TaskList struct {
    Tasks []Task `json:"tasks"`
}

func main() {
	// Plan and Execute

	// 1. User input
	// 2. Planning LLM ("Planning prompt") comes up with a step-by-step approach completing the query
	// 3. Plan is saved into a "Task List"
	// 4. Agent receives task and either performs a tool search or generates
	// 5. Result from the Agent passed to a "replanner prompt" which updates the plan (is tasked finished?)
	// 6. Replanner decides if the user input is good enough.

	userInput := "Question: What is the elevation range for the area that the eastern sector of the Colorado orogeny extends into?"

	json_llama3_8b_model := ollama.OllamaModel{
		Model:  "llama3:8b",
		Options: ollama.ModelOptions{Temperature: 0.7, NumCtx: 4096},
		Stream: false,
		Format: "json",
	}

	llama3_8b_model := ollama.OllamaModel{
		Model:  "llama3:8b",
		Options: ollama.ModelOptions{Temperature: 0.7, NumCtx: 4096},
		Stream: false,
	}

	planner := ollama.NewOllamaClient(json_llama3_8b_model) // Creats a task list as JSON.
	agent := ollama.NewOllamaClient(llama3_8b_model) // Executes the plan from the task list
	replanner := ollama.NewOllamaClient(llama3_8b_model) // Replan or output the result to the user

	tasksJsonFormat := `
	{
		"tasks":[
		   {
			  "todo":"",
			  "done":false
		   },
		   {
			  "todo":"",
			  "done":false
		   },
		   {
			  "todo":"",
			  "done":false
		   }
		]
	 }
	`

	plannerSystemPrompt := fmt.Sprintf(`
	You are a planner assistant.
	Generate a setp by step plan to solve the request.
	Respond in JSON format like this: %s
	Output only JSON.
	`, tasksJsonFormat)

	planner.SetSystemPrompt(plannerSystemPrompt)
	res, err := planner.Chat(userInput) // 2. Planning LLM ("Planning prompt") comes up with a step-by-step approach completing the query
	
	var taskList TaskList // 3. Plan is saved into a "Task List"
	err = json.Unmarshal([]byte(res.Message.Content), &taskList)
	if err != nil {
		log.Fatalf("Error marshal json...")
	}

	// Prepare Agent
	agentSystemPrompt := fmt.Sprintf(`
	You have to do the tasks given by a planner.
	Output only one precise answer.
	`)
	agent.SetSystemPrompt(agentSystemPrompt)

	replannerSystemPrompt := fmt.Sprintf(`
	The agent gives you a prompt and you have to decide if the tasks is done.
	You update the task list and if all are done, you output a final response to the user.
	`)
	replanner.SetSystemPrompt(replannerSystemPrompt)

	for _, task := range taskList.Tasks {
		fmt.Println(task.Todo)
		// 4. Agent receives task and either performs a tool search or generates
		if task.Done == false {
			agentRes, _ := agent.Chat(task.Todo)
			// 5. Result from the Agent passed to a "replanner prompt" which updates the plan (is tasked finished?)
			prompt := fmt.Sprintf(`
				This is the answer from the agent: %s
				This is the todo he had to do: %s
				Is this answer good enough and can we set the todo to "done": true?
				`, agentRes.Message.Content, task.Todo)
			res, _ := replanner.Chat(prompt) // 6. Replanner decides if the user input is good enough.	
			fmt.Println(res.Message.Content)
		}
	}
}