package main

import (
	"fmt"
	"math"
	"os"
)

const (
	Height       = 40 / 2
	Width        = 80 / 2
	ThetaSpacing = 0.07
	PhiSpacing   = 0.02
	R1           = 1
	R2           = 2
	K2           = 5
)

var (
	K1 = float64(Width) * K2 * 3 / (8 * (R1 + R2))
)

func renderFrame(A, B float64) {
	output := make([][]rune, Width)
	for i := range output {
		output[i] = make([]rune, Height)
	}
	zbuffer := make([][]float64, Width)
	for i := range zbuffer {
		zbuffer[i] = make([]float64, Height)
	}

	cosA, sinA := math.Cos(A), math.Sin(A)
	cosB, sinB := math.Cos(B), math.Sin(B)

	for theta := 0.0; theta < 2*math.Pi; theta += ThetaSpacing {
		costheta, sintheta := math.Cos(theta), math.Sin(theta)

		for phi := 0.0; phi < 2*math.Pi; phi += PhiSpacing {
			cosphi, sinphi := math.Cos(phi), math.Sin(phi)

			circlex := R2 + R1*costheta
			circley := R1 * sintheta

			x := circlex*(cosB*cosphi+sinA*sinB*sinphi) - circley*cosA*sinB
			y := circlex*(sinB*cosphi-sinA*cosB*sinphi) + circley*cosA*cosB
			z := K2 + cosA*circlex*sinphi + circley*sinA
			ooz := 1 / z

			xp := int(Width/2 + K1*ooz*x)
			yp := int(Height/2 - K1*ooz*y)

			L := cosphi*costheta*sinB - cosA*costheta*sinphi - sinA*sintheta + cosB*(cosA*sintheta-costheta*sinA*sinphi)

			if L > 0 {
				if xp >= 0 && xp < Width && yp >= 0 && yp < Height {
					luminanceIndex := int(L * 2)
					if luminanceIndex >= 0 && luminanceIndex < len(".,-~:;=!*#$@") {
						if ooz > zbuffer[xp][yp] {
							zbuffer[xp][yp] = ooz
							output[xp][yp] = []rune(".,-~:;=!*#$@")[luminanceIndex]
						}
					}
				}
			}
		}
	}

	// Clear screen
	fmt.Print("\x1b[H\x1b[2J")

	// Print the output array
	for j := 0; j < Height; j++ {
		for i := 0; i < Width; i++ {
			fmt.Printf("%c", output[i][j])
		}
		fmt.Println()
	}

	// Flush the output
	os.Stdout.Sync()
}

func main() {
	for i := 0.0; i < 360+i; i += 0.0005 {
		renderFrame(i, 2*i)
	}
}
