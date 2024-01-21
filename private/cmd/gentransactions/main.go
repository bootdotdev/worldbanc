package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

const (
	numTransactions = 10000
	startDate       = "2023-01-01"
	endDate         = "2023-12-31"
)

func main() {
	names := []string{"Alice", "Bob", "Charlie", "David", "Eva", "Frank", "Grace", "Hannah", "Ivan", "Julia", "Kyle", "Lily", "Mia", "Nora", "Oscar", "Paul", "Quinn", "Rachel", "Steve", "Tina"}

	start, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		log.Fatal(err)
	}
	end, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create(fmt.Sprintf("%v.csv", start.Year()))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"amount", "from_user_id", "to_user_id", "from_name", "to_name", "created_at"})

	for i := 0; i < numTransactions; i++ {
		amount := fmt.Sprintf("%.2f", rand.Float64()*1000)
		fromUserID := rand.Intn(20) + 1
		toUserID := rand.Intn(20) + 1
		fromName := names[fromUserID%len(names)]
		toName := names[toUserID%len(names)]
		createdAt := randomDate(start, end)
		writer.Write([]string{amount, fmt.Sprint(fromUserID), fmt.Sprint(toUserID), fromName, toName, createdAt.Format("2006-01-02 15:04:05")})
	}

	fmt.Println("CSV file 'transactions.csv' created successfully.")
}

func randomDate(start, end time.Time) time.Time {
	delta := end.Sub(start)
	sec := rand.Int63n(int64(delta.Seconds()))
	return start.Add(time.Second * time.Duration(sec))
}
