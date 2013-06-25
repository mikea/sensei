package classifiers

import (
	"github.com/mikea/sensei/ai"
	v "github.com/mikea/sensei/math/vector"
)

func NominalClassifierTrainerFromBinary(
	features []v.F64,
	labels []int,
	labelsCardinality int,
	binaryTrainer ai.ClassifierTrainer) ai.ClassifierTrainer {

	// classifiers := make([]BinaryClassifier, labelsCardinality)

	// // TODO(mike): parallelize
	// for i := 0; i < labelsCardinality; i++ {
	// 	boolLabels := make([]bool, len(labels))
	// 	for j, l := range labels {
	// 		boolLabels[j] = l == i
	// 	}

	// 	classifiers[i] = binaryTrainer(features, boolLabels)
	// }

	// return &compoundNominalClassifier{classifiers: classifiers}
	panic("Not Implemented")
}
