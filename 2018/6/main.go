package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	row  int
	dist int
}

type Center struct {
	x int
	y int
}

func main() {
	delim := "\n"
	parts := strings.Split(input, delim)
	centers := make([]Center, 0)
	grid := make([][]Point, 1000)
	for i := range grid {
		grid[i] = make([]Point, 1000)
	}
	for _, p := range parts {
		xyParts := strings.Split(p, ", ")
		x, err := strconv.Atoi(xyParts[0])
		if err != nil {
			fmt.Printf("failed to parse %q", xyParts[0])
			return
		}
		y, err := strconv.Atoi(xyParts[1])
		if err != nil {
			fmt.Printf("failed to parse %q", xyParts[1])
			return
		}
		centers = append(centers, Center{x, y})
	}
	total := 0
	for x := -1000; x < 1000; x++ {
		for y := -1000; y < 1000; y++ {
			d := 0
			for _, c := range centers {
				d += int(math.Abs(float64(c.x-x)) + math.Abs(float64(c.y-y)))
			}
			if d < 10000 {
				total += 1
			}
		}
	}
	fmt.Printf("%v", total)
}

var input = `165, 169
334, 217
330, 227
317, 72
304, 232
115, 225
323, 344
161, 204
316, 259
63, 250
280, 205
84, 282
271, 158
190, 296
106, 349
171, 178
203, 108
89, 271
193, 254
111, 210
341, 343
349, 311
143, 172
170, 307
128, 157
183, 315
211, 297
74, 281
119, 164
266, 345
184, 62
96, 142
134, 61
117, 52
318, 72
338, 287
61, 215
323, 255
93, 171
325, 249
183, 171
71, 235
329, 306
322, 219
151, 298
180, 255
336, 291
72, 300
223, 286
179, 257`
