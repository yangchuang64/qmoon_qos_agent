package gov

import (
	"errors"
	"fmt"
	"github.com/QOSGroup/qbase/client/context"
	"github.com/QOSGroup/qos/module/gov/mapper"
	"github.com/QOSGroup/qos/module/gov/types"
	go_amino "github.com/tendermint/go-amino"
)

func QueryProposal(cdc *go_amino.Codec, pID int64) (result types.Proposal, err error) {
	cliCtx := context.NewCLIContext().WithCodec(cdc)
	//pID, err := strconv.ParseUint(args[0], 10, 64)
	//if err != nil {
	//	return fmt.Errorf("proposal id %s is not a valid uint value", args[0])
	//}

	path := mapper.BuildQueryProposalPath(pID)
	res, err := cliCtx.Query(path, []byte{})

	if err != nil {
		return
	}
	if len(res) == 0 {
		err = errors.New("no result found")
		return
	}

	err = cliCtx.Codec.UnmarshalJSON(res, &result)
	return
}

func QueryProposals(cdc *go_amino.Codec, /*limit int64, depositor, voter, statusStr string*/) (result []types.Proposal, err error) {
	cliCtx := context.NewCLIContext().WithCodec(cdc)

	//var depositorAddr btypes.AccAddress
	//var voterAddr btypes.AccAddress
	var status types.ProposalStatus

	//if d, err := qcliacc.GetAddrFromValue(cliCtx, depositor); err == nil {
	//	depositorAddr = d
	//}
	//
	//if d, err := qcliacc.GetAddrFromValue(cliCtx, voter); err == nil {
	//	voterAddr = d
	//}

	//status = toProposalStatus(statusStr)

	queryParam := mapper.QueryProposalsParam{
		Depositor: nil,
		Voter:     nil,
		Status:    status,
		Limit:     1000000,
	}
	data, err := cliCtx.Codec.MarshalJSON(queryParam)
	if err != nil {
		return
	}

	path := mapper.BuildQueryProposalsPath()
	res, err := cliCtx.Query(path, data)
	if err != nil {
		return
	}
	if len(res) == 0 {
		err = errors.New("no result found")
		return
	}

	if err = cliCtx.Codec.UnmarshalJSON(res, &result); err != nil {
		return
	}

	if len(result) == 0 {
		err = fmt.Errorf("no matching proposals found")
	}
	return
}

func toProposalStatus(statusStr string) types.ProposalStatus {
	switch statusStr {
	case "DepositPeriod", "deposit_period":
		return types.StatusDepositPeriod
	case "VotingPeriod", "voting_period":
		return types.StatusVotingPeriod
	case "Passed", "passed":
		return types.StatusPassed
	case "Rejected", "rejected":
		return types.StatusRejected
	default:
		return types.StatusNil
	}
}

//func QueryVote(cdc *go_amino.Codec, pID int64, addrStr string) (types.Vote, error) {
//	cliCtx := context.NewCLIContext().WithCodec(cdc)
//
//	addr, err := qcliacc.GetAddrFromValue(cliCtx, addrStr)
//	if err != nil {
//		return types.Vote{}, fmt.Errorf("voter %s is not a valid address value", addrStr)
//	}
//
//	path := mapper.BuildQueryVotePath(pID, addr.String())
//	res, err := cliCtx.Query(path, []byte{})
//	if err != nil {
//		return types.Vote{}, err
//	}
//
//	if len(res) == 0 {
//		return types.Vote{}, errors.New("no result found")
//	}
//
//	var vote types.Vote
//	if err := cliCtx.Codec.UnmarshalJSON(res, &vote); err != nil {
//		return types.Vote{}, err
//	}
//
//	return vote, err
//}

func QueryVotes(cdc *go_amino.Codec, pID int64) (votes []types.Vote, err error) {
	cliCtx := context.NewCLIContext().WithCodec(cdc)

	path := mapper.BuildQueryVotesPath(pID)
	res, err := cliCtx.Query(path, []byte{})
	if err != nil {
		return
	}

	if len(res) == 0 {
		err = errors.New("no result found")
		return
	}

	if err = cliCtx.Codec.UnmarshalJSON(res, &votes); err != nil {
		return
	}

	if len(votes) == 0 {
		err = errors.New("no votes found")
	}
	return
}

//func QueryDeposit(cdc *go_amino.Codec, pID int64, addrStr string) (types.Deposit, error) {
//	cliCtx := context.NewCLIContext().WithCodec(cdc)
//
//	addr, err := qcliacc.GetAddrFromValue(cliCtx, addrStr)
//	if err != nil {
//		return types.Deposit{}, fmt.Errorf("voter %s is not a valid address value", addrStr)
//	}
//
//	path := mapper.BuildQueryVotePath(pID, addr.String())
//	res, err := cliCtx.Query(path, []byte{})
//	if err != nil {
//		return types.Deposit{}, err
//	}
//
//	if len(res) == 0 {
//		return types.Deposit{}, errors.New("no result found")
//	}
//
//	var deposit types.Deposit
//	if err := cliCtx.Codec.UnmarshalJSON(res, &deposit); err != nil {
//		return types.Deposit{}, nil
//	}
//
//	return deposit, err
//}

func QueryDeposits(cdc *go_amino.Codec, pID int64) (deposits []types.Deposit, err error) {
	cliCtx := context.NewCLIContext().WithCodec(cdc)

	path := mapper.BuildQueryDepositsPath(pID)
	res, err := cliCtx.Query(path, []byte{})
	if err != nil {
		return
	}

	if len(res) == 0 {
		err = errors.New("no result found")
		return
	}

	err = cliCtx.Codec.UnmarshalJSON(res, &deposits)
	return
}

func QueryTally(cdc *go_amino.Codec, pID int64 /*, addrStr string*/) (result types.TallyResult, err error) {
	cliCtx := context.NewCLIContext().WithCodec(cdc)

	path := mapper.BuildQueryTallyPath(pID)
	res, err := cliCtx.Query(path, []byte{})
	if err != nil {
		return
	}

	if len(res) == 0 {
		err = errors.New("no result found")
		return
	}

	err = cliCtx.Codec.UnmarshalJSON(res, &result)
	return
}
