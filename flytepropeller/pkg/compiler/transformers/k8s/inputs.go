package k8s

import (
	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/core"
	"github.com/flyteorg/flyte/flytepropeller/pkg/compiler/common"
	"github.com/flyteorg/flyte/flytepropeller/pkg/compiler/errors"
	"github.com/flyteorg/flyte/flytepropeller/pkg/compiler/validators"
)

func validateInputs(nodeID common.NodeID, iface *core.TypedInterface, inputs core.LiteralMap, errs errors.CompileErrors) (ok bool) {
	if iface == nil {
		errs.Collect(errors.NewValueRequiredErr(nodeID, "interface"))
		return false
	}

	if iface.GetInputs() == nil {
		errs.Collect(errors.NewValueRequiredErr(nodeID, "interface.InputsRef"))
		return false
	}

	varMap := make(map[string]*core.Variable, len(iface.GetInputs().GetVariables()))
	requiredInputsSet := sets.String{}
	for name, v := range iface.GetInputs().GetVariables() {
		varMap[name] = v
		requiredInputsSet.Insert(name)
	}

	boundInputsSet := sets.String{}
	for inputVar, inputVal := range inputs.GetLiterals() {
		v, exists := varMap[inputVar]
		if !exists {
			errs.Collect(errors.NewVariableNameNotFoundErr(nodeID, "", inputVar))
			continue
		}

		inputType := validators.LiteralTypeForLiteral(inputVal)
		err := validators.ValidateLiteralType(inputType)
		if err != nil {
			errs.Collect(errors.NewInvalidLiteralTypeErr(nodeID, inputVar, err))
			continue
		}
		if !validators.AreTypesCastable(inputType, v.GetType()) {
			errs.Collect(errors.NewMismatchingTypesErr(nodeID, inputVar, common.LiteralTypeToStr(v.GetType()), common.LiteralTypeToStr(inputType)))
			continue
		}

		boundInputsSet.Insert(inputVar)
	}

	if diff := requiredInputsSet.Difference(boundInputsSet); len(diff) > 0 {
		for param := range diff {
			errs.Collect(errors.NewParameterNotBoundErr(nodeID, param))
		}
	}

	return !errs.HasErrors()
}
