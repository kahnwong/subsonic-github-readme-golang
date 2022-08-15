package main

import (
	"os"
)

type subsonicAuth struct {
	Username string `url:"u"`
	Token    string `url:"t"`
	Salt     string `url:"s"`
	Version  string `url:"v"`
	Client   string `url:"c"`
	Format   string `url:"f"`
}

var authEnv = subsonicAuth{
	Username: os.Getenv("USERNAME"),
	Token:    os.Getenv("TOKEN"),
	Salt:     os.Getenv("SALT"),
	Version:  "1.16.1",
	Client:   "github-readme",
	Format:   "json",
}
