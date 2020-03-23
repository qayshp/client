// // Copyright 2020 Keybase, Inc. All rights reserved. Use of
// // this source code is governed by the included BSD license.

package engine

// import (
// 	"errors"
// 	"fmt"

// 	"github.com/keybase/client/go/libkb"
// 	keybase1 "github.com/keybase/client/go/protocol/keybase1"
// 	jsonw "github.com/keybase/go-jsonw"
// )

// type WotRevokeArg struct {
// 	Vouchee keybase1.UserVersion
// }

// // WotRevoke is an engine.
// type WotRevoke struct {
// 	arg *WotRevokeArg
// 	libkb.Contextified
// }

// // NewWotRevoke creates a WotRevoke engine.
// func NewWotRevoke(g *libkb.GlobalContext, arg *WotRevokeArg) *WotRevoke {
// 	return &WotRevoke{
// 		arg:          arg,
// 		Contextified: libkb.NewContextified(g),
// 	}
// }

// // Name is the unique engine name.
// func (e *WotRevoke) Name() string {
// 	return "WotRevoke"
// }

// // GetPrereqs returns the engine prereqs.
// func (e *WotRevoke) Prereqs() Prereqs {
// 	return Prereqs{Device: true}
// }

// // RequiredUIs returns the required UIs.
// func (e *WotRevoke) RequiredUIs() []libkb.UIKind {
// 	return []libkb.UIKind{}
// }

// // SubConsumers returns the other UI consumers for this engine.
// func (e *WotRevoke) SubConsumers() []libkb.UIConsumer {
// 	return nil
// }

// func (e *WotRevoke) fetchAttestationToRevoke(mctx libkb.MetaContext, vouchee *libkb.User) (res keybase1.WotVouch, err error) {
// 	defer mctx.TraceTimed(fmt.Sprintf("fetchAttestationToRevoke(%s)", vouchee.GetName()), func() error { return err })()
// 	voucherUsername := mctx.ActiveDevice.Username(mctx).String()
// 	voucheeUsername := vouchee.GetName()
// 	existingVouches, err := libkb.FetchWotVouches(mctx, libkb.FetchWotVouchesArg{Vouchee: voucheeUsername, Voucher: voucherUsername})
// 	if err != nil {
// 		return res, err
// 	}
// 	if len(existingVouches) > 1 {
// 		return res, fmt.Errorf("expected exactly 1 attestation by %s of %s but got %d", voucherUsername, voucheeUsername, len(existingVouches))
// 	}
// 	notFoundErr := fmt.Errorf("could not find attestation by %s of %s but got %d", voucherUsername, voucheeUsername)
// 	if len(existingVouches) == 0 {
// 		return res, notFoundErr
// 	}
// 	existingVouch := existingVouches[0]
// 	if !existingVouch.Voucher.Eq(mctx.ActiveDevice.UserVersion()) {
// 		return res, notFoundErr
// 	}
// 	if !existingVouch.Vouchee.ToUserVersion().Eq(vouchee.UserVersion) {
// 		return res, notFoundErr
// 	}
// 	if existingVouch.Status == keybase1.WotStatusType_REVOKED {
// 		return res, fmt.Errorf("attestation by %s of %s is already revoked", voucherUsername, voucheeUsername)
// 	}
// 	return existingVouch, nil
// }

// // Run starts the engine.
// func (e *WotRevoke) Run(mctx libkb.MetaContext) error {
// 	luArg := libkb.NewLoadUserArgWithMetaContext(mctx).WithUID(e.arg.Vouchee.Uid).WithStubMode(libkb.StubModeUnstubbed)
// 	them, err := libkb.LoadUser(luArg)
// 	if err != nil {
// 		return err
// 	}

// 	if them.GetCurrentEldestSeqno() != e.arg.Vouchee.EldestSeqno {
// 		return errors.New("vouchee has reset, make sure you still know them")
// 	}

// 	vouchToRevoke, err := e.fetchAttestationToRevoke(mctx, them)
// 	if err != nil {
// 		return err
// 	}

// 	signingKey, err := mctx.G().ActiveDevice.SigningKey()
// 	if err != nil {
// 		return err
// 	}

// 	sigVersion := libkb.KeybaseSignatureV2
// 	var inner []byte
// 	var sig string

// 	// ForcePoll is required.
// 	err = mctx.G().GetFullSelfer().WithSelfForcePoll(mctx.Ctx(), func(me *libkb.User) error {
// 		if me.GetUID() == e.arg.Vouchee.Uid {
// 			return libkb.InvalidArgumentError{Msg: "can't vouch for yourself"}
// 		}

// 		proof, err := me.WotVouchProof(mctx, signingKey, sigVersion, []byte{})
// 		if err != nil {
// 			return err
// 		}

// 		revokeSection := jsonw.NewDictionary()
// 		err := revokeSection.SetKey("sig_id", jsonw.NewString(sigToRevoke.String()))
// 		if err != nil {
// 			return nil, err
// 		}
// 		err = body.SetKey("revoke", revokeSection)
// 		if err != nil {
// 			return nil, err
// 		}

// 		inner, err = proof.J.Marshal()
// 		if err != nil {
// 			return err
// 		}

// 		sig, _, _, err = libkb.MakeSig(
// 			mctx,
// 			signingKey,
// 			libkb.LinkTypeWotVouch,
// 			inner,
// 			libkb.SigHasRevokes(true),
// 			keybase1.SeqType_PUBLIC,
// 			libkb.SigIgnoreIfUnsupported(true),
// 			me,
// 			sigVersion,
// 		)

// 		return err
// 	})

// 	if err != nil {
// 		return err
// 	}

// 	item := libkb.SigMultiItem{
// 		Sig:        sig,
// 		SigningKID: signingKey.GetKID(),
// 		Type:       string(libkb.LinkTypeWotVouch),
// 		SeqType:    keybase1.SeqType_PUBLIC,
// 		SigInner:   string(inner),
// 		Version:    sigVersion,
// 		Expansions: expansions,
// 	}

// 	payload := make(libkb.JSONPayload)
// 	payload["sigs"] = []interface{}{item}

// 	if _, err := e.G().API.PostJSON(mctx, libkb.APIArg{
// 		Endpoint:    "sig/multi",
// 		SessionType: libkb.APISessionTypeREQUIRED,
// 		JSONPayload: payload,
// 	}); err != nil {
// 		return err
// 	}
// 	voucherUsername := mctx.ActiveDevice().Username(mctx).String()
// 	return libkb.DismissWotNotifications(mctx, voucherUsername, them.GetName())
// }
