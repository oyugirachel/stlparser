# STL File Parser

## Instructions

To build and run the STL parser using Docker , follow these steps:

1. Make sure you have Go installed.
2. Clone the repository or unzip the provided package.
- N/B In the case you want to use the docker image follow these steps:
- Build the Docker image:
`docker build -t stlparser .`
- Run the Docker container, providing the path to the STL file as an argument:
 `docker run stlparser /path/to/your/stl/file.stl`

3. Open a terminal and navigate to the directory containing the code.
4. Run the following command to parse an STL file:

`go run main.go /path/to/your/file.stl`


Replace `/path/to/your/file.stl` with the path to the STL file you want to analyze for example im using moon.stl file so im running the `go run main.go Moon.stl`

## Design Decisions

- The code is organized into multiple files for better modularity.
- We use the `bufio.Scanner` to efficiently read the STL file line by line.
- External package `github.com/quartercastle/vector` is used for efficient vector operations.
- We parse each triangle's vertices and calculate the surface area using vector operations.
- The main program reads the STL file, parses it, and calculates the surface area.

## Testing
- The tests were implemented using goldie files `github.com/sebdah/goldie/v2`
- You can run the test by navigating to the directory with the test file in this case to the stlparser then run:
  ` go test -update stlparser/stlparser` or `go test`

  All the tests are running successfully.

## Performance Improvements

- To handle models with millions of triangles, we can consider parallelizing the calculation of the surface area.
- We could also optimize memory usage by not storing all triangles in memory at once but calculating the surface area incrementally.

Feel free to reach out with any questions or feedback.
