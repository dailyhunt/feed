package dsl

type ItemPipeline struct {
	Operations []ItemOperation
}

func (op *ItemPipeline) Filter() (*FilterOperation) {
	return nil
}

func (op *ItemPipeline) Score() (*ScoreOperation) {
	return nil
}

func (op *ItemPipeline) ScoreAndFilter() (*ScoreOperation) {
	return nil
}

func (op *ItemPipeline) Sort() (*SortOperation) {
	return nil
}
