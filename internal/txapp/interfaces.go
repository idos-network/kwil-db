package txapp

import (
	"context"

	"github.com/kwilteam/kwil-db/common/sql"
	"github.com/kwilteam/kwil-db/core/types"
	"github.com/kwilteam/kwil-db/internal/accounts"
	"github.com/kwilteam/kwil-db/internal/events"
	"github.com/kwilteam/kwil-db/internal/voting"
)

// Rebroadcaster is a service that marks events for rebroadcasting.
type Rebroadcaster interface {
	// MarkRebroadcast marks events for rebroadcasting.
	MarkRebroadcast(ctx context.Context, ids []types.UUID) error
}

// DB is the interface for the main SQL database.
type DB interface {
	sql.OuterTxMaker
	sql.ReadTxMaker
}

// package level funcs
// these can be overridden for testing
var (
	// getEvents gets all events, even if they have been
	// marked received
	getEvents = events.GetEvents

	// deleteEvent deletes an event. It will no longer
	// be broadcasted.
	deleteEvent  = events.DeleteEvent
	deleteEvents = events.DeleteEvents

	// voting
	setVoterPower                    = voting.SetValidatorPower
	getAllVoters                     = voting.GetValidators
	getResolutionsByThresholdAndType = voting.GetResolutionsByThresholdAndType
	deleteResolutions                = voting.DeleteResolutions
	markProcessed                    = voting.MarkProcessed
	getExpired                       = voting.GetExpired
	requiredPower                    = voting.RequiredPower
	getResolutionsByTypeAndProposer  = voting.GetResolutionIDsByTypeAndProposer
	createResolution                 = voting.CreateResolution
	approveResolution                = voting.ApproveResolution
	getVoterPower                    = voting.GetValidatorPower
	getDBHeight                      = voting.GetHeight
	updateDBHeight                   = voting.SetHeight

	// account functions
	getAccount = accounts.GetAccount
	credit     = accounts.Credit
	spend      = accounts.Spend
	transfer   = accounts.Transfer
)
