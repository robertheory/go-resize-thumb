# Thumbnail Image Resizer

This is a simple image resizer that resizes images to a thumbnail size.

## Description

This project aims to demonstrate how to resize images to a thumbnail size using Go. Using `github.com/nfnt/resize` package, the program reads multiple images, resizes them to a thumbnail size, and saves them to a new directory.

## Concepts and Technologies

1. **Go**: The project is written in Go, a statically typed, compiled language that is designed for simplicity and efficiency.

   1. **Image Resizing**: The project uses the `github.com/nfnt/resize` package to resize the images to a thumbnail size.

   2. **File Handling**: The project uses the `os` package to read and write files.

   3. **Goroutines**: The project uses goroutines to handle the resizing of the images concurrently.

   4. **Channels**: The project uses channels to communicate between the goroutines and to handle the concurrency.

   5. **WaitGroup**: The project uses the WaitGroup to wait for all the goroutines to finish before exiting the program.

   6. **Worker Pool**: The project uses a worker pool to limit the number of goroutines running concurrently.
