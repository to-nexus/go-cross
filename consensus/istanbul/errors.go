package istanbul

import "errors"

var (
	// ErrUnauthorizedAddress is returned when given address cannot be found in
	// current validator set.
	ErrUnauthorizedAddress = errors.New("unauthorized address")

	// ErrStoppedEngine is returned if the engine is stopped
	ErrStoppedEngine = errors.New("stopped engine")

	// ErrStartedEngine is returned if the engine is already started
	ErrStartedEngine = errors.New("started engine")

	// ErrInvalidProposal is returned when a prposal is malformed.
	ErrInvalidProposal = errors.New("invalid proposal")

	// ErrBlacklistedHash is returned if a block to import is on the blacklist.
	ErrBlacklistedHash = errors.New("blacklisted hash")

	// ErrInvalidSignature is returned when given signature is not signed by given
	// address.
	ErrInvalidSignature = errors.New("invalid signature")

	// ErrMismatchedRandomSignature is returned when the signature used
	// for random value generation does not match the block proposer's address.
	ErrMismatchedRandomSignature = errors.New("mismatched random signatures")

	// ErrUnknownBlock is returned when the list of validators is requested for a block
	// that is not part of the local blockchain.
	ErrUnknownBlock = errors.New("unknown block")

	// ErrUnauthorized is returned if a header is signed by a non authorized entity.
	ErrUnauthorized = errors.New("unauthorized")

	// ErrInvalidDifficulty is returned if the difficulty of a block is not 1
	ErrInvalidDifficulty = errors.New("invalid difficulty")

	// ErrInvalidExtraDataFormat is returned when the extra data format is incorrect
	ErrInvalidExtraDataFormat = errors.New("invalid extra data format")

	// ErrInvalidMixDigest is returned if a block's mix digest is not Istanbul digest.
	ErrInvalidMixDigest = errors.New("invalid Istanbul mix digest")

	// ErrInvalidUncleHash is returned if a block contains an non-empty uncle list.
	ErrInvalidUncleHash = errors.New("non empty uncle hash")

	// ErrInconsistentValidatorSet is returned if the validator set is inconsistent
	// ErrInconsistentValidatorSet = errors.New("non empty uncle hash")
	// ErrInvalidTimestamp is returned if the timestamp of a block is lower than the previous block's timestamp + the minimum block period.
	ErrInvalidTimestamp = errors.New("invalid timestamp")

	// ErrInvalidVotingChain is returned if an authorization list is attempted to
	// be modified via out-of-range or non-contiguous headers.
	ErrInvalidVotingChain = errors.New("invalid voting chain")

	// ErrInvalidVote is returned if a nonce value is something else that the two
	// allowed constants of 0x00..0 or 0xff..f.
	ErrInvalidVote = errors.New("vote nonce not 0x00..0 or 0xff..f")

	// ErrInvalidCommittedSeals is returned if the committed seal is not signed by any of parent validators.
	ErrInvalidCommittedSeals = errors.New("invalid committed seals")

	// ErrEmptyCommittedSeals is returned if the field of committed seals is zero.
	ErrEmptyCommittedSeals = errors.New("zero committed seals")

	// ErrMismatchTxhashes is returned if the TxHash in header is mismatch.
	ErrMismatchTxhashes = errors.New("mismatch transactions hashes")

	// ErrInvalidMessage is returned when the message is malformed.
	ErrInvalidMessage = errors.New("invalid message")

	// ErrFailedDecodePreprepare is returned when the PRE-PREPARE message is malformed.
	ErrFailedDecodePreprepare = errors.New("failed to decode PRE-PREPARE message")

	// ErrFailedDecodeCommit is returned when the COMMIT message is malformed.
	ErrFailedDecodeCommit = errors.New("failed to decode COMMIT message")

	// ErrFailedDecodeRoundChange is returned when the COMMIT message is malformed.
	ErrFailedDecodeRoundChange = errors.New("failed to decode ROUND-CHANGE message")
)
