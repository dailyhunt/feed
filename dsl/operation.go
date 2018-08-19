package dsl

type ItemOperation interface {
	Type() (string)
}

type FilterOperation struct {
	// TODO: what to filter on
}

type ScoreOperation struct {
	// TODO: what to score on
}

type SortOperation struct {
	// TODO: what to sort on
}

type TopKOperation struct {
	// TODO: Top K to collect
}

type BottomKOperation struct {
	// TODO: Bottom K to collect
}

func (op *FilterOperation) Type() (string) {
	return ""
}

func (op *ScoreOperation) Type() (string) {
	return ""
}

func (op *SortOperation) Type() (string) {
	return ""
}

func (op *TopKOperation) Type() (string) {
	return ""
}

func (op *BottomKOperation) Type() (string) {
	return ""
}

func (op *FilterOperation) Filter() (*FilterOperation) {
	return nil
}

func (op *FilterOperation) Score() (*ScoreOperation) {
	return nil
}

func (op *FilterOperation) ScoreAndFilter() (*ScoreOperation) {
	return nil
}

func (op *ScoreOperation) Filter() (*ScoreOperation) {
	return nil
}

func (op *ScoreOperation) Sort() (*SortOperation) {
	return nil
}

func (op *SortOperation) Sort() (*SortOperation) {
	return nil
}

func (op *SortOperation) TopK() (*TopKOperation) {
	return nil
}

func (op *SortOperation) BottomK() (*BottomKOperation) {
	return nil
}
