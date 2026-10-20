Goal: Create a Redis copy using Neovim as the editor.

## October 15
I have to read about the protocol Redis use and what are its capabilities. 
I have no idea where to start, so I will start from memory. My first goal is to create an application that allows me to store integers for keys, like a hash in Ruby. For now, it should save it only in memory and be accessed by the terminal.
Reading https://pkg.go.dev/io@go1.27.2 to understand how to parse input.
I couldn't find anything in docs so I tried Google.
From [this page](https://tutorialedge.net/golang/reading-console-input-golang/), the secret is to use `bufio.NewReader(os.Stdin)`

## October 16
I was able to read something from the terminal and output it. I am still not sure what exactly is bufion and fmt, but I guess it's like scanf and printf in C.
My next goal is to find out which operations Redis has, then implement the simplest one, which probably is "add". 
SET and GET are the most basic ones. My goal now is to add those 2 commands that should keep the values while the program is running only. Storing them to the disk comes later.

## October 19
For now, I have given up on the idea to learn VIM as I go along. I will do some flashcards daily to at least have the basics. Meanwhile, I will continue with IntelliJ Goland.