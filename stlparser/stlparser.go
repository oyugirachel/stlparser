package stlparser

import (
	"bufio"
	"github.com/quartercastle/vector"
	"os"
	"strconv"
	"strings"
)

// Triangle represents a 3D triangle in an STL file.
type Triangle struct {
	V1, V2, V3 vector.Vector
}
// ParseSTLFile parses an STL file and returns the triangles.
func ParseSTLFile(filename string) (int, float64, error) {
    file, err := os.Open(filename)
    if err != nil {
        return 0, 0, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var triangles []Triangle
    vertexCount := 0

    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if strings.HasPrefix(line, "vertex") {
            parts := strings.Fields(line)
            if len(parts) != 4 {
                continue
            }

            v := vector.Vector{}
            for i := 1; i <= 3; i++ {
                part, err := strconv.ParseFloat(parts[i], 64)
                if err != nil {
                    return 0, 0, err
                }
                v = append(v, part)
            }

            if vertexCount%3 == 0 {
                triangles = append(triangles, Triangle{V1: v})
            } else {
                last := &triangles[len(triangles)-1]
                switch vertexCount % 3 {
                case 1:
                    last.V2 = v
                case 2:
                    last.V3 = v
                }
            }
            vertexCount++
        }
    }

    if err := scanner.Err(); err != nil {
        return 0, 0, err
    }

    numTriangles := len(triangles)
    surfaceArea := 0.0
    for _, t := range triangles {
        surfaceArea += t.CalculateArea()
    }

    return numTriangles, surfaceArea, nil
}


// CalculateArea calculates the surface area of a triangle.
func (t *Triangle) CalculateArea() float64 {
	v1 := vector.Vector{t.V2[0] - t.V1[0], t.V2[1] - t.V1[1], t.V2[2] - t.V1[2]}
	v2 := vector.Vector{t.V3[0] - t.V1[0], t.V3[1] - t.V1[1], t.V3[2] - t.V1[2]}

	crossProduct, err := vector.Vector.Cross(v1, v2)
	if err != nil {
		return 0 // Handles the error appropriately
	}
	triangleArea := vector.Vector.Magnitude(crossProduct) / 2
	return triangleArea
}
