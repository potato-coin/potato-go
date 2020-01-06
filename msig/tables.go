package msig

import potato "github.com/rise-worlds/potato-go"

type ProposalRow struct {
	ProposalName       potato.Name              `json:"proposal_name"`
	RequestedApprovals []potato.PermissionLevel `json:"requested_approvals"`
	ProvidedApprovals  []potato.PermissionLevel `json:"provided_approvals"`
	PackedTransaction  potato.HexBytes          `json:"packed_transaction"`
}
