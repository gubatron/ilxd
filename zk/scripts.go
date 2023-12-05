// Copyright (c) 2022 The illium developers
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file.

package zk

import (
	"embed"
	"errors"
	"fmt"
	"github.com/libp2p/go-libp2p/core/crypto"
	icrypto "github.com/project-illium/ilxd/crypto"
	"github.com/project-illium/ilxd/zk/lurk/macros"
)

//go:embed lurk/basic_transfer.lurk
var basicTransferScriptLurk embed.FS
var basicTransferScriptData string

//go:embed lurk/password_script.lurk
var passwordScriptLurk embed.FS
var passwordScriptData string

//go:embed lurk/multisig_script.lurk
var multisigScriptLurk embed.FS
var multisigScriptData string

//go:embed lurk/timelocked_multisig.lurk
var timelockedMultisigScriptLurk embed.FS
var timelockedMultisigScriptData string
var timeLockedMultisigCommitment []byte

func init() {
	mp, err := macros.NewMacroPreprocessor(macros.WithStandardLib(), macros.RemoveComments())
	if err != nil {
		panic(err)
	}
	data, err := basicTransferScriptLurk.ReadFile("lurk/basic_transfer.lurk")
	if err != nil {
		panic(err)
	}
	basicTransferScriptData, err = mp.Preprocess(string(data))
	if err != nil {
		panic(err)
	}

	data, err = passwordScriptLurk.ReadFile("lurk/password_script.lurk")
	if err != nil {
		panic(err)
	}
	passwordScriptData, err = mp.Preprocess(string(data))
	if err != nil {
		panic(err)
	}

	data, err = multisigScriptLurk.ReadFile("lurk/multisig_script.lurk")
	if err != nil {
		panic(err)
	}
	multisigScriptData, err = mp.Preprocess(string(data))
	if err != nil {
		panic(err)
	}

	data, err = timelockedMultisigScriptLurk.ReadFile("lurk/timelocked_multisig.lurk")
	if err != nil {
		panic(err)
	}
	timelockedMultisigScriptData, err = mp.Preprocess(string(data))
	if err != nil {
		panic(err)
	}
	timeLockedMultisigCommitment, err = LurkCommit(timelockedMultisigScriptData)
	if err != nil {
		panic(err)
	}
}

// BasicTransferScript returns the basic transfer lurk script
func BasicTransferScript() string {
	return basicTransferScriptData
}

// PasswordScript returns the password lurk script
func PasswordScript() string {
	return passwordScriptData
}

// MultisigScript returns the multsig lurk script
func MultisigScript() string {
	return multisigScriptData
}

// TimelockedMultisigScript returns the timelocked multisig lurk script
func TimelockedMultisigScript() string {
	return timelockedMultisigScriptData
}

// TimelockedMultisigScriptCommitment returns the script commitment hash
// for the timelocked multisig script.
func TimelockedMultisigScriptCommitment() []byte {
	ret := make([]byte, len(timeLockedMultisigCommitment))
	copy(ret, timeLockedMultisigCommitment)
	return ret
}

func MakeMultisigUnlockingParams(pubkeys [][]byte, sigs [][]byte, sigHash []byte) ([]byte, error) {
	sigCpy := make([][]byte, len(sigs))
	copy(sigCpy, sigs)

	keys := make([]crypto.PubKey, 0, len(pubkeys)/2)
	for i := 0; i < len(pubkeys); i += 2 {
		key, err := icrypto.PublicKeyFromXY(pubkeys[i], pubkeys[i+1])
		if err != nil {
			return nil, err
		}
		keys = append(keys, key)
	}

	keySelector := "(cons "
	for i, key := range keys {
		if len(sigs) > 0 {
			valid, err := key.Verify(sigHash, sigs[0])
			if err != nil {
				return nil, err
			}
			if valid {
				keySelector += "1 "
				sigs = sigs[1:]
			} else {
				keySelector += "0 "
			}
		} else {
			keySelector += "0 "
		}
		if i == len(keys)-1 {
			keySelector += "nil"
		} else {
			keySelector += "(cons "
		}
	}
	for i := 0; i < len(keys); i++ {
		keySelector += ")"
	}

	unlockignScript := "(cons " + keySelector + " "
	for _, sig := range sigCpy {
		unlockignScript += "(cons "
		if len(sig) != 64 {
			return nil, errors.New("invalid signature len")
		}
		unlockignScript += fmt.Sprintf("(cons 0x%x 0x%x) ", sig[:32], sig[32:])
	}
	unlockignScript += "nil)"
	for i := 0; i < len(sigCpy); i++ {
		unlockignScript += ")"
	}
	return []byte(unlockignScript), nil
}