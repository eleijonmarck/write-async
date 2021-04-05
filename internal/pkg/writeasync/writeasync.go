package writeasync

type WriterAsyncer struct{}

func NewWriteAsync() *WriterAsyncer {
	return &WriterAsyncer{}
}

type AddJobPayload struct {
	Name string `json:"name"`
}

func (w *WriterAsyncer) ProcessJob() {}
