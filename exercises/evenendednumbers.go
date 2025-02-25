package exercises

import "fmt"

func EvenEndedNumbersCount() int {
	count := 0


	for i := 1000; i <= 9999; i ++ {
		for j := i; j <= 9999; j ++ {
			m := i * j
			m_str := fmt.Sprintf("%d", m)

			if m_str[0] == m_str[len(m_str) - 1] {
				count ++
			}
		}
	}

	return count
}