// +build mkmanpage

package main

import (
	"log"
	"os"
)

type mkmanpage struct {
}

var mkmanpageOptions mkmanpage

func init() {
	_, err := optionsParser.AddCommand("mkmanpage", "make manpage", "renders manpage to stdout", &mkmanpageOptions)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *mkmanpage) Execute(args []string) error {
	optionsParser.WriteManPage(os.Stdout)
	return nil
}
