package engine

// 单任务版引擎
type SimpleEngine struct {
}

func (e *SimpleEngine) Run(queue ...Request) {

	for len(queue) > 0 {
		r := queue[0]
		queue = queue[1:]

		results, err := Worker(r)
		if err != nil {
			continue
		}
		for _, r := range results.Requests {
			queue = append(queue, r)
		}
		//queue = append(queue, results.Requests...)

	}
}
