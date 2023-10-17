package stlparser

import( 
    "os"
    "testing"
    "github.com/sebdah/goldie/v2"
    "fmt"
	"github.com/quartercastle/vector"


)

func TestParseSTLFile(t *testing.T) {
    // Implementing test cases to parse sample STL files and check the number of triangles.
    tests := []struct {
		filename       string
		expectedOutput string
	}{
		{
			filename:       "../cube.stl",
			expectedOutput: "cube",
		},
		{
			filename:       "../Moon.stl",
			expectedOutput: "Moon",
		},
	}

	for _, test := range tests {
		t.Run(test.filename, func(t *testing.T) {
			file, err := os.Open(test.filename)
			if err != nil {
				t.Fatal(err)
			}
			defer file.Close()

			numTriangles, surfaceArea, err:= ParseSTLFile(test.filename)
			if err != nil {
				t.Fatal(err)
			}

			g := goldie.New(t)
			g.Assert(t, test.expectedOutput, []byte(fmt.Sprintf("Number of Triangles: %d\nSurface Area: %.4f\n", numTriangles, surfaceArea)))
		})
	}
}

func TestCalculateSurfaceArea(t *testing.T) {
    // Implementing test cases to calculate surface area for sample triangles.
	t.Run("TestCalculateSurfaceArea", func(t *testing.T) {
        triangle := Triangle{
            V1: vector.Vector{0, 0, 0},
            V2: vector.Vector{0, 1, 0},
            V3: vector.Vector{1, 0, 0},
        }

        expectedArea := 0.5
        actualArea := triangle.CalculateArea()

        if actualArea != expectedArea {
            t.Errorf("Expected area %f, but got %f", expectedArea, actualArea)
        }
    })
}
