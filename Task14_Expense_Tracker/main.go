package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Expense struct {
	Date        time.Time
	Amount      float64
	Category    string
	Description string
}

var expenses []Expense

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Expense Tracker")
		fmt.Println("1. Add Expense")
		fmt.Println("2. View Expenses")
		fmt.Println("3. Generate Report")
		fmt.Println("4. Search Expenses")
		fmt.Println("5. Exit")

		fmt.Print("Select an option: ")
		scanner.Scan()
		option := scanner.Text()

		switch option {
		case "1":
			addExpense(scanner)
		case "2":
			viewExpenses()
		case "3":
			generateReport(scanner)
		case "4":
			searchExpenses(scanner)
		case "5":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}

		fmt.Println()
	}
}

func addExpense(scanner *bufio.Scanner) {
	fmt.Println("Add Expense")

	var expense Expense

	fmt.Print("Enter the date (YYYY-MM-DD): ")
	scanner.Scan()
	dateStr := scanner.Text()
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		fmt.Println("Invalid date format. Please try again.")
		return
	}
	expense.Date = date

	fmt.Print("Enter the amount: ")
	scanner.Scan()
	amountStr := scanner.Text()
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		fmt.Println("Invalid amount. Please try again.")
		return
	}
	expense.Amount = amount

	fmt.Print("Enter the category: ")
	scanner.Scan()
	category := scanner.Text()
	expense.Category = category

	fmt.Print("Enter the description: ")
	scanner.Scan()
	description := scanner.Text()
	expense.Description = description

	expenses = append(expenses, expense)

	fmt.Println("Expense added successfully.")
}

func viewExpenses() {
	fmt.Println("Expense List")
	fmt.Println("-------------")

	if len(expenses) == 0 {
		fmt.Println("No expenses found.")
		return
	}

	for _, expense := range expenses {
		fmt.Printf("Date: %s\n", expense.Date.Format("2006-01-02"))
		fmt.Printf("Amount: %.2f\n", expense.Amount)
		fmt.Printf("Category: %s\n", expense.Category)
		fmt.Printf("Description: %s\n", expense.Description)
		fmt.Println("-------------")
	}
}

func generateReport(scanner *bufio.Scanner) {
	fmt.Println("Expense Report")
	fmt.Println("--------------")

	if len(expenses) == 0 {
		fmt.Println("No expenses found.")
		return
	}

	categoryExpenses := make(map[string][]float64)

	for _, expense := range expenses {
		categoryExpenses[expense.Category] = append(categoryExpenses[expense.Category], expense.Amount)
	}

	for category, amounts := range categoryExpenses {
		total := calculateTotal(amounts)
		average := calculateAverage(amounts)
		max := calculateMax(amounts)
		min := calculateMin(amounts)

		fmt.Printf("Category: %s\n", category)
		fmt.Printf("Total Expenses: %.2f\n", total)
		fmt.Printf("Average Expense: %.2f\n", average)
		fmt.Printf("Max Expense: %.2f\n", max)
		fmt.Printf("Min Expense: %.2f\n", min)
		fmt.Println("--------------")
	}
}

func searchExpenses(scanner *bufio.Scanner) {
	fmt.Println("Search Expenses")

	fmt.Print("Enter the category: ")
	scanner.Scan()
	category := scanner.Text()

	fmt.Println("Expense List")
	fmt.Println("-------------")

	found := false

	for _, expense := range expenses {
		if strings.EqualFold(expense.Category, category) {
			fmt.Printf("Date: %s\n", expense.Date.Format("2006-01-02"))
			fmt.Printf("Amount: %.2f\n", expense.Amount)
			fmt.Printf("Category: %s\n", expense.Category)
			fmt.Printf("Description: %s\n", expense.Description)
			fmt.Println("-------------")

			found = true
		}
	}

	if !found {
		fmt.Println("No expenses found for the given category.")
	}
}

func calculateTotal(amounts []float64) float64 {
	total := 0.0
	for _, amount := range amounts {
		total += amount
	}
	return total
}

func calculateAverage(amounts []float64) float64 {
	count := len(amounts)

	if count > 0 {
		return calculateTotal(amounts) / float64(count)
	}

	return 0
}

func calculateMax(amounts []float64) float64 {
	max := amounts[0]
	for _, amount := range amounts {
		if amount > max {
			max = amount
		}
	}
	return max
}

func calculateMin(amounts []float64) float64 {
	min := amounts[0]
	for _, amount := range amounts {
		if amount < min {
			min = amount
		}
	}
	return min
}
