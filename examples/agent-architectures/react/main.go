package main

import (
	"fmt"
	"strings"

	"github.com/TobiasGleiter/ai-agents/pkg/llms/ollama"
	ChatColor "github.com/TobiasGleiter/ai-agents/internal/color"
)

type Env interface {
	Step(action string) (string, int, bool, map[string]interface{})
}

func main() {
	// ReAct (Reasoning and Acting)
	// - Reasoning traces and task-specific actions

	// Question from HotpotQA
	userInput := "Question: What is the elevation range for the area that the eastern sector of the Colorado orogeny extends into?" + " \n"

	instruction := `Solve a question answering task with interleaving Thought, Action, Observation steps. 
		Thought can reason about the current situation, and Action can be three types: 
		(1) Search[entity], which searches the exact entity on Wikipedia and returns the first paragraph if it exists. If not, it will return some similar entities to search.
		(2) Lookup[keyword], which returns the next sentence containing keyword in the current passage.
		(3) Finish[answer], which returns the answer and finishes the task.
		Here are some examples.
		`

	example := `
	Question: Musician and satirist Allie Goertz wrote a song about the \"The Simpsons\" character Milhouse, who Matt Groening named after who?\nThought 1: I only need to search Milhouse and find who it is named after.\nAction 1: Search[Milhouse]\nObservation 1: Milhouse Mussolini Van Houten is a recurring character in the Fox animated television series The Simpsons voiced by Pamela Hayden and created by Matt Groening. Milhouse is Bart Simpson's best friend in Mrs. Krabappel's fourth grade class at Springfield Elementary School. He is an insecure, gullible, and less popular child than Bart who is often led into trouble by Bart, who takes advantage of his friend's na\u00c3\u00afvet\u00c3\u00a9. Milhouse is a regular target for school bully Nelson Muntz and his friends Jimbo Jones, Dolph Starbeam and Kearney Zzyzwicz. Milhouse has a crush on Bart's sister, Lisa, a common plot element.\nThought 2: The paragraph does not tell who Milhouse is named after, maybe I can look up \"named after\".\nAction 2: Lookup[named after]\nObservation 2: (Result 1 / 1) Milhouse was designed by Matt Groening for a planned series on NBC, which was abandoned.[4] The design was then used for a Butterfinger commercial, and it was decided to use the character in the series.[5][6] Milhouse was named after U.S. president Richard Nixon, whose middle name was Milhous. The name was the most \"unfortunate name Matt Groening could think of for a kid\".[1] Years earlier, in a 1986 Life in Hell comic entitled \"What to Name the Baby\", Groening listed Milhouse as a name \"no longer recommended\".[7] Milhouse is a favorite among the staff as Al Jean noted \"most of the writers are more like Milhouse than Bart\".[1] His last name was given to him by Howard Gewirtz, a freelance writer who wrote the episode \"Homer Defined\"
	`

	var prompt = userInput
	var messages []ollama.ModelMessage
	messages = append(messages, ollama.ModelMessage{
		Role: "system",
		Content: fmt.Sprintf(`%s %s`, instruction, example),
	})

	var nCalls = 0
	//var nBadCall = 0

	iterationLimit := 10
	for i := 1; i < iterationLimit; i++ {
		nCalls++

		// Prompt the Model
		messages = append(messages, ollama.ModelMessage{
			Role: "user",
			Content: prompt+fmt.Sprintf("Thought %d:", i),
		})

		llamaRequest := ollama.Model{
			Model:  "llama3:8b",
			Messages: messages,
			Options: ollama.ModelOptions{
				Temperature: 1,
				NumCtx: 4096,
			},
			Stream: false,
			Stop:   []string{fmt.Sprintf("\nObservation %d:", i)},
		}

		thoughtAction, _ := ollama.ChatReAct(llamaRequest)
		parts := strings.Split(thoughtAction.Message.Content, fmt.Sprintf("\nAction %d: ", i))
		var thought, action string
		if len(parts) == 2 {
			thought = strings.TrimSpace(parts[0])
			action = strings.TrimSpace(parts[1])
		} else {
			fmt.Println("ohh...", thoughtAction.Message.Content)
		}

		// Search Wikipedia for information...
		observation := "The Colorado orogeny was an episode of mountain building (an orogeny) in Colorado and surrounding areas. This took place from 1780 to 1650 million years ago (Mya), during the Paleoproterozoic (Statherian Period). It is recorded in the Colorado orogen, a >500-km-wide belt of oceanic arc rock that extends southward into New Mexico. The Colorado orogeny was likely part of the larger Yavapai orogeny."
		//ChatColor.PrintColor(ChatColor.Cyan, thought + " " +  action)
		stepStr := fmt.Sprintf("Thought %d: %s\nAction %d: %s\nObservation %d: %s\n", i, thought, i, action, i, observation)
		
		messages = append(messages, ollama.ModelMessage{
			Role: "assistant",
			Content: stepStr,
		})
		ChatColor.PrintColor(ChatColor.Yellow, thoughtAction.Message.Content)
	}

}