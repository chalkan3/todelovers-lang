package tvm

type loop struct {
	start int64
	end   int64
}

func newLoop(start int64, end int64) *loop {
	return &loop{
		start: start,
		end:   end,
	}
}
