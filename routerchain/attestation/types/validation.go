package types

import (
	"encoding/binary"
	"encoding/hex"
	math "math"
	"math/big"
	"slices"
	"sort"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/router-protocol/sdk-go/routerchain/util"

	errorsmod "cosmossdk.io/errors"
	multichainTypes "github.com/router-protocol/sdk-go/routerchain/multichain/types"
)

//////////////////////////////////////
//      BRIDGE VALIDATOR(S)         //
//////////////////////////////////////

// ToInternal transforms a BridgeValidator into its fully validated internal type
func (b BridgeValidator) ToInternal() (*InternalBridgeValidator, error) {
	return NewInternalBridgeValidator(b)
}

// BridgeValidators is the sorted set of validator data for Ethereum bridge MultiSig set
type BridgeValidators []BridgeValidator

func (b BridgeValidators) ToInternal() (*InternalBridgeValidators, error) {
	ret := make(InternalBridgeValidators, len(b))
	for i := range b {
		ibv, err := NewInternalBridgeValidator(b[i])
		if err != nil {
			return nil, errorsmod.Wrapf(err, "member %d", i)
		}
		ret[i] = ibv
	}
	return &ret, nil
}

// Equal checks that slice contents and order are equal
func (b BridgeValidators) Equal(o BridgeValidators) bool {
	if len(b) != len(o) {
		return false
	}

	for i, bv := range b {
		ov := o[i]
		if bv != ov {
			return false
		}
	}

	return true
}

// InternalBridgeValidator is a BridgeValidator but with validated EthereumAddress
type InternalBridgeValidator struct {
	Power           uint64
	EthereumAddress EthAddress
}

func NewInternalBridgeValidator(bridgeValidator BridgeValidator) (*InternalBridgeValidator, error) {
	ethAddr, err := NewEthAddress(bridgeValidator.EthereumAddress)
	if err != nil {
		return nil, errorsmod.Wrap(err, "invalid bridge validator eth address")
	}

	i := &InternalBridgeValidator{
		Power:           bridgeValidator.Power,
		EthereumAddress: *ethAddr,
	}
	if err := i.ValidateBasic(); err != nil {
		return nil, errorsmod.Wrap(err, "invalid bridge validator")
	}
	return i, nil
}

func (i InternalBridgeValidator) ValidateBasic() error {
	if err := i.EthereumAddress.ValidateBasic(); err != nil {
		return errorsmod.Wrap(err, "ethereum address")
	}
	return nil
}

func (i InternalBridgeValidator) ToExternal() BridgeValidator {
	return BridgeValidator{
		Power:           i.Power,
		EthereumAddress: i.EthereumAddress.GetAddress().Hex(),
	}
}

// InternalBridgeValidators is the sorted set of validator data for Ethereum bridge MultiSig set
type InternalBridgeValidators []*InternalBridgeValidator

func (i InternalBridgeValidators) ToExternal() BridgeValidators {
	bridgeValidators := make([]BridgeValidator, len(i))
	for b := range bridgeValidators {
		bridgeValidators[b] = i[b].ToExternal()
	}

	return BridgeValidators(bridgeValidators)
}

// Sort sorts the validators by power
func (b InternalBridgeValidators) Sort() {
	sort.Slice(b, func(i, j int) bool {
		if b[i].Power == b[j].Power {
			// Secondary sort on eth address in case powers are equal
			return EthAddrLessThan(b[i].EthereumAddress, b[j].EthereumAddress)
		}
		return b[i].Power > b[j].Power
	})
}

// PowerDiff returns the difference in power between two bridge validator sets
// note this is Gravity bridge power *not* Cosmos voting power. Cosmos voting
// power is based on the absolute number of tokens in the staking pool at any given
// time Gravity bridge power is normalized using the equation.
//
// validators cosmos voting power / total cosmos voting power in this block = gravity bridge power / u32_max
//
// As an example if someone has 52% of the Cosmos voting power when a validator set is created their Gravity
// bridge voting power is u32_max * .52
//
// Normalized voting power dramatically reduces how often we have to produce new validator set updates. For example
// if the total on chain voting power increases by 1% due to inflation, we shouldn't have to generate a new validator
// set, after all the validators retained their relative percentages during inflation and normalized Gravity bridge power
// shows no difference.
func (b InternalBridgeValidators) PowerDiff(c InternalBridgeValidators) float64 {
	powers := map[string]int64{}
	// loop over b and initialize the map with their powers
	for _, bv := range b {
		powers[bv.EthereumAddress.GetAddress().Hex()] = int64(bv.Power)
	}

	// subtract c powers from powers in the map, initializing
	// uninitialized keys with negative numbers
	for _, bv := range c {
		if val, ok := powers[bv.EthereumAddress.GetAddress().Hex()]; ok {
			powers[bv.EthereumAddress.GetAddress().Hex()] = val - int64(bv.Power)
		} else {
			powers[bv.EthereumAddress.GetAddress().Hex()] = -int64(bv.Power)
		}
	}

	var delta float64
	for _, v := range powers {
		// NOTE: we care about the absolute value of the changes
		delta += math.Abs(float64(v))
	}

	return math.Abs(delta / float64(math.MaxUint32))
}

// TotalPower returns the total power in the bridge validator set
func (b InternalBridgeValidators) TotalPower() (out uint64) {
	for _, v := range b {
		out += v.Power
	}
	return
}

// HasDuplicates returns true if there are duplicates in the set
func (b InternalBridgeValidators) HasDuplicates() bool {
	m := make(map[string]struct{}, len(b))
	// creates a hashmap then ensures that the hashmap and the array
	// have the same length, this acts as an O(n) duplicates check
	for i := range b {
		m[b[i].EthereumAddress.GetAddress().Hex()] = struct{}{}
	}
	return len(m) != len(b)
}

// GetPowers returns only the power values for all members
func (b InternalBridgeValidators) GetPowers() []uint64 {
	r := make([]uint64, len(b))
	for i := range b {
		r[i] = b[i].Power
	}
	return r
}

// ValidateBasic performs stateless checks
func (b InternalBridgeValidators) ValidateBasic() error {
	if len(b) == 0 {
		return ErrEmpty
	}
	for i := range b {
		if err := b[i].ValidateBasic(); err != nil {
			return errorsmod.Wrapf(err, "member %d", i)
		}
	}
	if b.HasDuplicates() {
		return errorsmod.Wrap(ErrDuplicate, "addresses")
	}
	return nil
}

//////////////////////////////////////
//             VALSETS              //
//////////////////////////////////////

// NewValset returns a new valset
func NewValset(nonce, height uint64, members InternalBridgeValidators) (*Valset, error) {
	if err := members.ValidateBasic(); err != nil {
		return nil, errorsmod.Wrap(err, "invalid members")
	}
	members.Sort()
	var mem []BridgeValidator
	for _, val := range members {
		mem = append(mem, val.ToExternal())
	}
	vs := Valset{Nonce: uint64(nonce), Members: mem, Height: height}
	return &vs,
		nil
}

func (v Valset) GetCheckpoint(destChainType multichainTypes.ChainType) ([]byte, error) {
	switch destChainType {
	case multichainTypes.CHAIN_TYPE_SOLANA:
		return v.GetSolanaCheckpoint()
	case multichainTypes.CHAIN_TYPE_SUI:
		return v.GetSuiCheckpoint()
	default:
		return v.GetEvmCheckpoint()
	}
}

// GetCheckpoint returns the checkpoint
func (v Valset) GetSolanaCheckpoint() ([]byte, error) {
	newValsetNonce := new(big.Int).SetUint64(v.Nonce)
	bufLength := 32 + 16 + len(v.Members)*20 + len(v.Members)*8
	data := make([]byte, bufLength)
	offset := 0
	// add discriminator
	discriminator := []byte{99, 104, 101, 99, 107, 112, 111, 105, 110, 116, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	copy(data[offset:], discriminator)
	offset += 32

	newValsetNonceBytes := newValsetNonce.Bytes()
	slices.Reverse(newValsetNonceBytes)
	copy(data[offset:], newValsetNonceBytes)
	offset += 16

	for _, member := range v.Members {
		validatorBytes := gethcommon.HexToAddress(member.EthereumAddress).Bytes()
		copy(data[offset:], validatorBytes)
		offset += len(validatorBytes)
	}

	for _, member := range v.Members {
		powerBytes := make([]byte, 8)
		binary.LittleEndian.PutUint64(powerBytes, member.Power)
		copy(data[offset:], powerBytes)
		offset += 8
	}
	return crypto.Keccak256Hash(data).Bytes(), nil
}

// GetCheckpoint returns the checkpoint
func (v Valset) GetEvmCheckpoint() ([]byte, error) {
	// error case here should not occur outside of testing since the above is a constant
	contractAbi, abiErr := abi.JSON(strings.NewReader(util.ValsetCheckpointABIJSON))
	if abiErr != nil {
		return nil, errorsmod.Wrap(abiErr, "invalid valset checkpoint abi")
	}

	// the contract argument is not a arbitrary length array but a fixed length 32 byte
	// array, therefore we have to utf8 encode the string (the default in this case) and
	// then copy the variable length encoded data into a fixed length array. This function
	// will panic if routerIDstring is too long to fit in 32 bytes
	// routerId, err := util.StrToFixByteArray(routerIdString)
	// if err != nil {
	// 	panic(err)
	// }

	checkpointBytes := []uint8("checkpoint")
	var checkpoint [32]uint8
	copy(checkpoint[:], checkpointBytes)

	memberAddresses := make([]gethcommon.Address, len(v.Members))
	convertedPowers := make([]*big.Int, len(v.Members))
	for i, m := range v.Members {
		memberAddresses[i] = gethcommon.HexToAddress(m.EthereumAddress)
		convertedPowers[i] = big.NewInt(int64(m.Power))
	}
	// the word 'checkpoint' needs to be the same as the 'name' above in the checkpointAbiJson
	// but other than that it's a constant that has no impact on the output. This is because
	// it gets encoded as a function name which we must then discard.
	bytes, packErr := contractAbi.Pack("checkpoint",
		checkpoint,
		big.NewInt(int64(v.Nonce)),
		memberAddresses,
		convertedPowers,
	)

	// this should never happen outside of test since any case that could crash on encoding
	// should be filtered above.
	if packErr != nil {
		return nil, errorsmod.Wrap(packErr, "Error packing checkpoint!")
	}

	// we hash the resulting encoded bytes discarding the first 4 bytes these 4 bytes are the constant
	// method name 'checkpoint'. If you were to replace the checkpoint constant in this code you would
	// then need to adjust how many bytes you truncate off the front to get the output of abi.encode()
	hash := crypto.Keccak256Hash(bytes[4:])
	return hash.Bytes(), nil
}


func (v Valset) GetSuiCheckpoint() ([]byte, error) {
	var checkpoint []byte

	checkpoint = append(checkpoint, []byte{
		99, 104, 101, 99, 107, 112, 111, 105, 110, 116, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}...)

	serializedNonce := serializeU256(v.Nonce)
	checkpoint = append(checkpoint, serializedNonce...)

	var serializedValidators, serializedPowers []byte
	for _, member := range v.Members {
		val,_:= hex.DecodeString(strings.TrimPrefix(member.EthereumAddress,"0x"))
		serializedValidator := serializeVectorU8(val)
		serializedValidators = append(serializedValidators, serializedValidator...)
	}

	for _, member := range v.Members {
		powerBytes := make([]byte, 8)
		binary.LittleEndian.PutUint64(powerBytes, member.Power) 
		serializedPowers = append(serializedPowers, powerBytes...)
	}

	validatorsWithLength := serializeVectorU8(serializedValidators)
	powersWithLength := serializeVectorU8(serializedPowers)

	checkpoint = append(checkpoint, validatorsWithLength...)
	checkpoint = append(checkpoint, powersWithLength...)

	return 	crypto.Keccak256Hash(checkpoint).Bytes() , nil

}

// Serialize a 256-bit unsigned integer (u256) in Little Endian
func serializeU256(value uint64) []byte {
	bytes := make([]byte, 32)
	binary.LittleEndian.PutUint64(bytes, value) // Little Endian padding
	return bytes
}

func serializeVectorU8(data []byte) []byte {
	length := len(data)
	result := []byte{byte(length)} // Add the length of the vector as the first byte
	result = append(result, data...) // Append the actual data
	return result
}


// WithoutEmptyMembers returns a new Valset without member that have 0 power or an empty Ethereum address.
func (v *Valset) WithoutEmptyMembers() *Valset {
	if v == nil {
		return nil
	}
	r := Valset{
		Nonce:   v.Nonce,
		Members: make([]BridgeValidator, 0, len(v.Members)),
		Height:  0,
	}
	for i := range v.Members {
		if _, err := v.Members[i].ToInternal(); err == nil {
			r.Members = append(r.Members, v.Members[i])
		}
	}
	return &r
}

// Equal compares all of the valset members, additionally returning an error explaining the problem
func (v Valset) Equal(o Valset) (bool, error) {
	if v.Height != o.Height {
		return false, errorsmod.Wrap(ErrInvalid, "valset heights mismatch")
	}

	if v.Nonce != o.Nonce {
		return false, errorsmod.Wrap(ErrInvalid, "valset nonces mismatch")
	}

	var bvs BridgeValidators = v.Members
	var ovs BridgeValidators = o.Members
	if !bvs.Equal(ovs) {
		return false, errorsmod.Wrap(ErrInvalid, "valset members mismatch")
	}

	return true, nil
}

// Valsets is a collection of valset
type Valsets []Valset

func (v Valsets) Len() int {
	return len(v)
}

func (v Valsets) Less(i, j int) bool {
	return v[i].Nonce > v[j].Nonce
}

func (v Valsets) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

// This interface is implemented by all the types that are used
// to create transactions on Ethereum that are signed by validators.
// The naming here could be improved.
type EthereumSigned interface {
	GetCheckpoint(gravityIDstring string) ([]byte, error)
}

// var _ EthereumSigned = &Valset{}
