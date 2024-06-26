package provider

import (
	"bitbucket.org/booqsoftware/terraform-provider-qbee/internal/qbee"
	"context"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func typeAndIdentifier(tag types.String, node types.String) (qbee.ConfigType, string) {
	var configType qbee.ConfigType
	var identifier string
	if tag.IsNull() {
		configType = qbee.ConfigForNode
		identifier = node.ValueString()
	} else {
		configType = qbee.ConfigForTag
		identifier = tag.ValueString()
	}
	return configType, identifier
}

func nullableStringValue(value string) basetypes.StringValue {
	if value == "" {
		return types.StringNull()
	} else {
		return types.StringValue(value)
	}
}

func listFromStructs[T interface{}](ctx context.Context, items []T, elemType attr.Type) (types.List, diag.Diagnostics) {
	if len(items) == 0 {
		return types.ListNull(elemType), nil
	}

	itemsValue, diags := types.ListValueFrom(
		ctx,
		elemType,
		items,
	)

	return itemsValue, diags
}
