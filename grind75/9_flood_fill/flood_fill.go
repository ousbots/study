package flood_fill

type Pos struct {
	x int
	y int
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	q := []Pos{{x: sr, y: sc}}
	seen := map[Pos]struct{}{}

	checkColor := image[sr][sc]

	for len(q) != 0 {
		check := q[0]
		q = q[1:]

		if _, ok := seen[check]; ok {
			continue
		}
		seen[check] = struct{}{}

		image[check.x][check.y] = color

		if check.x < len(image)-1 {
			if image[check.x+1][check.y] == checkColor {
				q = append(q, Pos{check.x + 1, check.y})
			}
		}

		if check.x > 0 {
			if image[check.x-1][check.y] == checkColor {
				q = append(q, Pos{check.x - 1, check.y})
			}
		}

		if check.y < len(image[check.x])-1 {
			if image[check.x][check.y+1] == checkColor {
				q = append(q, Pos{check.x, check.y + 1})
			}
		}

		if check.y > 0 {
			if image[check.x][check.y-1] == checkColor {
				q = append(q, Pos{check.x, check.y - 1})
			}
		}
	}

	return image
}
