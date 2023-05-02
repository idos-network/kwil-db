package sqlite

import (
	"math/rand"
	"time"
)

func (c *Connection) Savepoint() (*Savepoint, error) {
	return beginSavepoint(c)
}

func beginSavepoint(c *Connection) (*Savepoint, error) {
	saveName := "tx_" + randomSavepointName(8)

	err := c.Execute("SAVEPOINT " + saveName)
	if err != nil {
		return nil, err
	}
	return &Savepoint{conn: c, saveName: saveName}, nil
}

type Savepoint struct {
	conn     *Connection
	saveName string
}

// With both Commit and Rollback, if the checkpoint fails, it doesn't matter

func (sp *Savepoint) Commit() error {
	return sp.conn.Execute("RELEASE " + sp.saveName)
}

func (sp *Savepoint) CommitAndCheckpoint() error {
	err := sp.Commit()
	if err != nil {
		return err
	}

	return sp.conn.CheckpointWal()
}

func (sp *Savepoint) Rollback() error {
	err := sp.conn.Execute("ROLLBACK TO " + sp.saveName)
	if err != nil {
		return err
	}

	return sp.conn.Execute("RELEASE " + sp.saveName)
}

var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
var alphanumericRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomSavepointName(length int) string {
	if length < 2 {
		panic("Length must be at least 2 to generate a valid savepoint name.")
	}

	result := make([]rune, length)
	// First character must be a letter
	result[0] = letterRunes[rand.Intn(len(letterRunes))]

	// Rest of the characters can be alphanumeric
	for i := 1; i < length; i++ {
		result[i] = alphanumericRunes[rand.Intn(len(alphanumericRunes))]
	}

	return string(result)
}