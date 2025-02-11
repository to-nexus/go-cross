// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package core

import (
	"fmt"
	"io"
	"math/big"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	"github.com/ethereum/go-ethereum/consensus/istanbul/protocols"
	"github.com/ethereum/go-ethereum/rlp"
)

// Construct a new message set to accumulate messages for given sequence/view number.
func newMsgSet(valSet istanbul.ValidatorSet) *msgSet {
	return &msgSet{
		view: &istanbul.View{
			Round:    new(big.Int),
			Sequence: new(big.Int),
		},
		messagesMu: new(sync.Mutex),
		messages:   make(map[common.Address]protocols.Message),
		valSet:     valSet,
	}
}

// ----------------------------------------------------------------------------

type msgSet struct {
	view       *istanbul.View
	valSet     istanbul.ValidatorSet
	messagesMu *sync.Mutex
	messages   map[common.Address]protocols.Message
}

// msgMapAsStruct is a temporary holder struct to convert messages map to a slice when Encoding and Decoding msgSet
type msgMapAsStruct struct {
	Address common.Address
	Msg     protocols.Message
}

func (ms *msgSet) View() *istanbul.View {
	return ms.view
}

func (ms *msgSet) Add(msg protocols.Message) error {
	ms.messagesMu.Lock()
	defer ms.messagesMu.Unlock()
	ms.messages[msg.Source()] = msg
	return nil
}

func (ms *msgSet) Values() (result []protocols.Message) {
	ms.messagesMu.Lock()
	defer ms.messagesMu.Unlock()

	for _, v := range ms.messages {
		result = append(result, v)
	}

	return result
}

func (ms *msgSet) Size() int {
	ms.messagesMu.Lock()
	defer ms.messagesMu.Unlock()
	return len(ms.messages)
}

func (ms *msgSet) Get(addr common.Address) protocols.Message {
	ms.messagesMu.Lock()
	defer ms.messagesMu.Unlock()
	return ms.messages[addr]
}

// ----------------------------------------------------------------------------

func (ms *msgSet) String() string {
	ms.messagesMu.Lock()
	defer ms.messagesMu.Unlock()
	addresses := make([]string, 0, len(ms.messages))
	for _, v := range ms.messages {
		addresses = append(addresses, v.Source().String())
	}
	return fmt.Sprintf("[%v]", strings.Join(addresses, ", "))
}

// EncodeRLP serializes msgSet into Ethereum RLP format
// valSet is currently not being encoded.
func (ms *msgSet) EncodeRLP(w io.Writer) error {
	if ms == nil {
		return nil
	}
	ms.messagesMu.Lock()
	defer ms.messagesMu.Unlock()

	// maps cannot be RLP encoded, convert the map into a slice of struct and then encode it
	var messagesAsSlice []msgMapAsStruct
	for k, v := range ms.messages {
		msgMapAsStruct := msgMapAsStruct{
			Address: k,
			Msg:     v,
		}
		messagesAsSlice = append(messagesAsSlice, msgMapAsStruct)
	}

	return rlp.Encode(w, []interface{}{
		ms.view,
		messagesAsSlice,
	})
}

// DecodeRLP deserializes rlp stream into msgSet
// valSet is currently not being decoded
func (ms *msgSet) DecodeRLP(stream *rlp.Stream) error {
	// Don't decode msgSet if the size of the stream is 0
	_, size, _ := stream.Kind()
	if size == 0 {
		return nil
	}
	var msgSet struct {
		MsgView *istanbul.View
		//		valSet        istanbul.ValidatorSet
		MessagesSlice []msgMapAsStruct
	}
	if err := stream.Decode(&msgSet); err != nil {
		return err
	}

	// convert the messages struct slice back to map
	messages := make(map[common.Address]protocols.Message)
	for _, msgStruct := range msgSet.MessagesSlice {
		messages[msgStruct.Address] = msgStruct.Msg
	}

	ms.view = msgSet.MsgView
	//	ms.valSet = msgSet.valSet
	ms.messages = messages
	ms.messagesMu = new(sync.Mutex)

	return nil
}
