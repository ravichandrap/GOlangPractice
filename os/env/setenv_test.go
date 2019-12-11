package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetEnv(t *testing.T) {
	setEnv()
	post := os.Getenv("PORT")
	fmt.Println(post)
	assert.NotEmpty(t, os.Getenv("PORT"))
}
