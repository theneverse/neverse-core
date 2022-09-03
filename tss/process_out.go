package tss

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	btss "github.com/binance-chain/tss-lib/tss"
	"github.com/meshplus/bitxhub-core/tss/conversion"
	"github.com/meshplus/bitxhub-core/tss/message"
	"github.com/sirupsen/logrus"
)

// The TaskMessage body content
type BulkWireMsg struct {
	WiredBulkMsgs []byte
	MsgIdentifier string
	Routing       *btss.MessageRouting
}

func NewBulkWireMsg(msg []byte, id string, r *btss.MessageRouting) BulkWireMsg {
	return BulkWireMsg{
		WiredBulkMsgs: msg,
		MsgIdentifier: id,
		Routing:       r,
	}
}

// ProcessOutCh sends msg out ==========================================================================================
// add msgType to the library message, send it as requested or broadcast
func (t *TssInstance) ProcessOutCh(tssMsg btss.Message, msgType message.TssMsgType) error {
	// 1. get wire byte, include msg content and rout info
	msgData, r, err := tssMsg.WireBytes()
	// if we cannot get the wire share, the tss will fail, we just quit.
	if err != nil {
		return fmt.Errorf("fail to get wire bytes: %w", err)
	}

	// 2. store this message in cache
	cachedWiredMsg := NewBulkWireMsg(msgData, tssMsg.GetFrom().Moniker, r)
	if r.IsBroadcast {
		dat, ok := t.cachedWireBroadcastMsgLists.Load(tssMsg.Type())
		if !ok {
			l := []BulkWireMsg{cachedWiredMsg}
			t.cachedWireBroadcastMsgLists.Store(tssMsg.Type(), l)
		} else {
			cachedList := dat.([]BulkWireMsg)
			cachedList = append(cachedList, cachedWiredMsg)
			t.cachedWireBroadcastMsgLists.Store(tssMsg.Type(), cachedList)
		}
	} else {
		dat, ok := t.cachedWireUnicastMsgLists.Load(tssMsg.Type() + ":" + r.To[0].String())
		if !ok {
			l := []BulkWireMsg{cachedWiredMsg}
			t.cachedWireUnicastMsgLists.Store(tssMsg.Type()+":"+r.To[0].String(), l)
		} else {
			cachedList := dat.([]BulkWireMsg)
			cachedList = append(cachedList, cachedWiredMsg)
			t.cachedWireUnicastMsgLists.Store(tssMsg.Type()+":"+r.To[0].String(), cachedList)
		}
	}

	// 3. send broadcast msg
	t.cachedWireBroadcastMsgLists.Range(func(key, value interface{}) bool {
		wiredMsgList := value.([]BulkWireMsg)
		wiredMsgType := key.(string)
		if len(wiredMsgList) == t.msgNum {
			err := t.sendBulkMsg(wiredMsgType, msgType, wiredMsgList)
			if err != nil {
				t.logger.Errorf("error in send bulk message")
				return true
			}
			t.cachedWireBroadcastMsgLists.Delete(key)
		}
		return true
	})

	// 4. send unicast msg
	t.cachedWireUnicastMsgLists.Range(func(key, value interface{}) bool {
		wiredMsgList := value.([]BulkWireMsg)
		ret := strings.Split(key.(string), ":")
		wiredMsgType := ret[0]
		if len(wiredMsgList) == t.msgNum {
			err := t.sendBulkMsg(wiredMsgType, msgType, wiredMsgList)
			if err != nil {
				t.logger.Errorf("error in send bulk message")
				return true
			}
			t.cachedWireUnicastMsgLists.Delete(key)
		}
		return true
	})

	return nil
}

// sendBulkMsg packages the message with type and signature, put it into the network module to send it out
func (t *TssInstance) sendBulkMsg(wiredMsgType string, tssMsgType message.TssMsgType, wiredMsgList []BulkWireMsg) error {
	// 1. get msg rout info
	// since all the messages in the list is the same round, so it must have the same dest
	// we just need to get the routing info of the first message
	r := wiredMsgList[0].Routing

	// 2. msg marshal
	buf, err := json.Marshal(wiredMsgList)
	if err != nil {
		return fmt.Errorf("error in marshal the cachedWireMsg: %w", err)
	}

	// 3. sign with p2p privkey (receive a message requires the sign to certify that the message came from the source)
	sig, err := conversion.GenerateSignature(buf, t.msgID, t.localPrivK)
	if err != nil {
		t.logger.Errorf("fail to generate the share's signature")
		return err
	}

	// 4. package msg with routing and signature
	taskMsg := &message.TaskMessage{
		Routing:   r,
		RoundInfo: wiredMsgType,
		Message:   buf,
		Sig:       sig,
	}
	taskMsgBytes, err := json.Marshal(taskMsg)
	if err != nil {
		return fmt.Errorf("fail to convert tss msg to wire bytes: %w", err)
	}

	// 5. constructor a p2p msg with type
	wireMsg := &message.WireMessage{
		MsgID:   t.msgID,
		MsgType: tssMsgType,
		MsgData: taskMsgBytes,
	}

	// 6. get msg to info
	partiesID := []uint64{}
	if len(r.To) != 0 {
		idUint, err := strconv.ParseUint(r.To[0].Id, 10, 32)
		if err != nil {
			return fmt.Errorf("parse uint error: %v", err)
		}
		partiesID = append(partiesID, idUint)
	}

	// 7. set to network module
	t.renderToP2P(&message.SendMsgChan{
		WireMsg:   wireMsg,
		PartiesID: partiesID,
	})

	return nil
}

// NotifyTaskDone broadcasts a message, the current task is over  ======================================================
func (t *TssInstance) NotifyTaskDone() error {
	taskDoneMsg := &message.TssTaskNotifier{
		FromID:   t.localPartyID,
		TaskDone: true,
	}

	msgData, err := json.Marshal(taskDoneMsg)
	if err != nil {
		return fmt.Errorf("marshal tss task notifier error: %w", err)
	}

	wireMsg := &message.WireMessage{
		MsgID:   t.msgID,
		MsgType: message.TSSTaskDone,
		MsgData: msgData,
	}

	var parties []uint64

	for id, _ := range t.getPartyInfo().PartyIDMap {
		if id == t.localPartyID {
			continue
		}
		tmpId, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			t.logger.Errorf("get parties:%s", id)
			return fmt.Errorf("receiverBroadcastHashToPeers parse int error: %v", err)
		}
		parties = append(parties, tmpId)
	}
	t.logger.WithFields(logrus.Fields{"parties": t.getPartyInfo().PartyIDMap, "msgType": string(wireMsg.MsgType)}).
		Debug("send task Done")

	t.renderToP2P(&message.SendMsgChan{
		WireMsg:   wireMsg,
		PartiesID: parties,
	})
	return nil
}
