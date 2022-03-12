package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

type Counter struct {
	Name string

	m     sync.RWMutex
	count int
}

func (c *Counter) Increment() int {
	c.m.Lock()
	defer c.m.Unlock()
	c.count++
	return c.count
}

func (c *Counter) View() int {
	c.m.RLock()
	defer c.m.RUnlock()
	return c.count
}

type secret struct {
	ID         string
	CreateTime time.Time

	token string
}

func (s *secret) Read(p []byte) (int, error) {
	return bytes.NewBuffer(p).WriteString(s.token)
}

func NewSecret() io.Reader {
	return &secret{
		ID:         "dummy_id",
		CreateTime: time.Now(),
		// 小文字はjson出力されない
		token: "dummy_token",
	}
}

type Chip struct {
	Number int
}

type Card struct {
	string
	Chip
	Number int
}

func (c *Chip) Scan() {
	fmt.Println(c.Number)
}

func main() {
	c := &Counter{Name: "Access"}
	fmt.Println(c.Increment())
	fmt.Println(c.View())

	s := NewSecret()
	err := json.NewEncoder(os.Stdout).Encode(s)
	if err != nil {
		fmt.Println("failed to json encode, error = ", err)
	}

	card := Card{
		string: "Credit",
		Chip: Chip{
			Number: 4242424242,
		},
		Number: 525252525,
	}

	// CardにはScanがないためChipのScanが呼ばれている
	// 4242424242が出力される
	card.Scan()
}
