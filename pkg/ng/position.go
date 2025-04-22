package ng

func IsAdjacent(x1, y1, x2, y2 int) bool {
	return absDiff(x1, x2) <= 1 && absDiff(y1, y2) <= 1
}
