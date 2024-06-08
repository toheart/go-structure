/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import "github.com/toheart/go-structure/cmd"

//go:generate go run github.com/google/wire/cmd/wire ./cmd
func main() {
	cmd.Execute()
}
