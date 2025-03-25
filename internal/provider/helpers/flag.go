package helpers

import (
	"context"
	"slices"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func IsFlagImporting(ctx context.Context, req resource.ReadRequest) (bool, diag.Diagnostics) {
	v, diags := req.Private.GetKey(ctx, "importing")

	return slices.Equal(v, []byte("1")), diags
}

// SetFlagImporting checks the respDiags and if they are error-free it sets the `importing` as a private flag inside
// SetKeyer. It appends its own results to respDiags.
//
// The caller must include in respDiags the result of state modification in the first place, to ensure consistency.
// The SetKeyer is something like resp.Private.
func SetFlagImporting(ctx context.Context, importing bool, sk SetKeyer, respDiags *diag.Diagnostics) {
	if respDiags.HasError() {
		return
	}

	b := []byte("0")
	if importing {
		b = []byte("1")
	}

	diags := sk.SetKey(ctx, "importing", b)
	respDiags.Append(diags...)
}

// SetKeyer is something like ReadResponse.Private or ImportStateResponse.Private.
type SetKeyer interface {
	SetKey(ctx context.Context, key string, value []byte) diag.Diagnostics
}

var (
	rr resource.ReadResponse
	ir resource.ImportStateResponse

	// ensure interface match
	_ SetKeyer = rr.Private
	_ SetKeyer = ir.Private
)
