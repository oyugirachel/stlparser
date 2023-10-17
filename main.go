package main

import (
    "fmt"
    "os"
    "stlparser/stlparser"

)

func main() {
    // Check if an STL file is provided as a command-line argument
    if len(os.Args) != 2 {
        fmt.Println("Usage: stlparser <path_to_stl_file>")
        os.Exit(1)
    }

    filePath := os.Args[1]
    // Open and parse the STL file
	stlFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening STL file:", err)
		return
	}
	defer stlFile.Close()

    numTriangles, surfaceArea, err := stlparser.ParseSTLFile(filePath)
    if err != nil {
        fmt.Printf("Error reading STL file: %v\n", err)
        os.Exit(1)
    }

    fmt.Printf("Number of Triangles: %d\n", numTriangles)
    fmt.Printf("Surface Area: %f\n", surfaceArea)
}


