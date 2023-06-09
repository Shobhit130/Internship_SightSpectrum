package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Knetic/govaluate"
	"github.com/jdkato/prose/v2"
	"github.com/jonreiter/govader"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to the Chatbot!")
	fmt.Println("I can perform various tasks such as lowercase/uppercase conversion, text reversal, POS tagging, sentiment analysis, time/date retrieval, and mathematical calculations.")
	fmt.Println("-----------------------")

	for {
		fmt.Print("You: ")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)

		if userInput == "exit" {
			fmt.Println("Chatbot: Goodbye!")
			break
		}

		response := getResponse(userInput)
		fmt.Println("Chatbot:", response)
	}
}

func getResponse(userInput string) string {
	lowerInput := strings.ToLower(userInput)

	// Check for specific keywords or phrases in the user's input
	if strings.Contains(lowerInput, "lowercase") {
		text := extractTextInQuotes(lowerInput)
		if text == "" {
			return "Please provide a sentence in quotes for lowercasing."
		}

		lowercaseText := strings.ToLower(text)
		return fmt.Sprintf("Lowercase: %s", lowercaseText)
	} else if strings.Contains(lowerInput, "reverse") {
		text := extractTextInQuotes(lowerInput)
		if text == "" {
			return "Please provide a sentence in quotes for reversing."
		}
		reversedText := reverseString(text)
		return fmt.Sprintf("Reversed: %s", reversedText)
	} else if strings.Contains(lowerInput, "pos") {
		text := extractTextInQuotes(lowerInput)
		if text == "" {
			return "Please provide a sentence in quotes for POS tagging."
		}

		tokens := tokenizeText(text)
		posTags := performPOSTagging(tokens)

		response := "POS Tags:"
		for i, token := range tokens {
			response += fmt.Sprintf("\n%s - %s", token, posTags[i])
		}

		return response
	} else if strings.Contains(lowerInput, "uppercase") {
		text := extractTextInQuotes(lowerInput)
		if text == "" {
			return "Please provide a sentence in quotes for uppercasing."
		}

		uppercaseText := strings.ToUpper(text)
		return fmt.Sprintf("Uppercase: %s", uppercaseText)
	} else if strings.Contains(lowerInput, "sentiment") {
		text := extractTextInQuotes(lowerInput)
		if text == "" {
			return "Please provide a sentence in quotes for sentiment analysis."
		}

		sentiment := analyzeSentiment(text)
		return fmt.Sprintf("Sentiment: %s", sentiment)
	} else if strings.Contains(lowerInput, "hello") || strings.Contains(lowerInput, "hi") {
		return "Hello! How can I assist you today?"
	} else if strings.Contains(lowerInput, "time") {
		currentTime := time.Now().Format("15:04:05")
		return fmt.Sprintf("The current time is %s.", currentTime)
	} else if strings.Contains(lowerInput, "date") {
		currentDate := time.Now().Format("2006-01-02")
		return fmt.Sprintf("Today's date is %s.", currentDate)
	} else if strings.Contains(lowerInput, "calculate") {
		// Extract the mathematical expression from the user's input and perform the calculation
		expression := strings.ReplaceAll(lowerInput, "calculate", "")
		expression = strings.TrimSpace(expression)

		result, err := calculateExpression(expression)
		if err != nil {
			return "Sorry, I couldn't calculate that."
		}

		return fmt.Sprintf("The result is %f.", result)
	} else if strings.Contains(lowerInput, "thank") {
		return "You're welcome! If you have any more questions, feel free to ask."
	}

	return "I'm sorry, I don't have information about that."
}

func calculateExpression(expression string) (float64, error) {
	expression = strings.ReplaceAll(expression, " ", "") // Remove spaces from the expression

	// Create a new evaluator
	evaluator, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		return 0, err
	}

	// Evaluate the expression
	result, err := evaluator.Evaluate(nil)
	if err != nil {
		return 0, err
	}

	// Convert the result to float64
	resultFloat, ok := result.(float64)
	if !ok {
		return 0, fmt.Errorf("invalid result type")
	}

	return resultFloat, nil
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func tokenizeText(text string) []string {
	doc, _ := prose.NewDocument(text)
	var tokens []string
	for _, tok := range doc.Tokens() {
		tokens = append(tokens, tok.Text)
	}
	return tokens
}

func performPOSTagging(tokens []string) []string {
	doc, _ := prose.NewDocument(strings.Join(tokens, " "))
	var posTags []string
	for _, tok := range doc.Tokens() {
		posTags = append(posTags, tok.Tag)
	}
	return posTags
}

func extractTextInQuotes(input string) string {
	startIndex := strings.Index(input, "\"")
	endIndex := strings.LastIndex(input, "\"")

	if startIndex == -1 || endIndex == -1 || startIndex == endIndex {
		return ""
	}

	return strings.TrimSpace(input[startIndex+1 : endIndex])
}

func analyzeSentiment(text string) string {
	sentimentAnalyzer := govader.NewSentimentIntensityAnalyzer()

	// Analyze the sentiment of the text
	sentiment := sentimentAnalyzer.PolarityScores(text)

	// Determine the sentiment label based on the compound score
	if sentiment.Compound >= 0.05 {
		return "Positive"
	} else if sentiment.Compound <= -0.05 {
		return "Negative"
	}

	return "Neutral"
}
