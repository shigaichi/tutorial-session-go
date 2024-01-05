package view

func Minus(a, b int) int { return a - b }

func Plus(a, b int) int { return a + b }

func Sequence(page Page, pageLinkMaxDispNum int) []int {
	begin := max(1, page.Page+1-pageLinkMaxDispNum/2)
	end := begin + (pageLinkMaxDispNum - 1)

	if end > page.GetTotalPages() {
		end = page.GetTotalPages()
		begin = max(1, end-(pageLinkMaxDispNum-1))
	}

	sequence := make([]int, 0, pageLinkMaxDispNum)
	for i := begin; i <= end; i++ {
		sequence = append(sequence, i)
	}
	return sequence
}
