package stratum

import (
	"errors"
	"fmt"
	"slices"
)

// MiningConfigureParams is sent from the client to the pool.
// The client uses the message to advertise its features and to request/allow some protocol extensions.
// This should be the first message sent.
type MiningConfigureParams struct {
	Supported  []string
	Parameters map[string]any
}
type VersionRollingConfigurationRequest struct {
	Mask        uint32
	MinBitCount uint8
}
type InfoConfigurationRequest struct {
	ConnectionURL,
	HWVersion,
	SWVersion,
	HWID string
}

// FromRequest parses the [MiningConfigureParams] from a [Request].
func (p *MiningConfigureParams) FromRequest(r *Request) error {
	if r.Method != MethodMiningConfigure.String() {
		return errors.New("incorrect method")
	}
	l := len(r.Params)
	if l != 2 {
		return errors.New("incorrect parameter length; must be 2")
	}

	supported, ok := r.Params[0].([]any)
	if !ok {
		return errors.New("invalid supported format")
	}
	p.Supported = make([]string, len(supported))
	for idx, s := range supported {
		p.Supported[idx] = s.(string)
	}
	p.Parameters = r.Params[1].(map[string]any)
	return nil
}

// ToRequest creates a [Request] from the [MiningConfigureParams].
func (p *MiningConfigureParams) ToRequest(id MessageID) *Request {
	params := make([]any, 2)
	params[0] = p.Supported
	params[1] = p.Parameters
	return NewRequest(id, MethodMiningConfigure, params)
}

// Supports checks if the [MiningConfigureParams] contains the given extension.
func (p *MiningConfigureParams) Supports(extension Extension) bool {
	return slices.Contains(p.Supported, extension.String())
}

// GetVersionRollingMask returns the [ExtensionVersionRolling] config, if provided.
func (p *MiningConfigureParams) GetVersionRolling() (*VersionRollingConfigurationRequest, error) {
	rawMask, ok := p.Parameters["version-rolling.mask"]
	if !ok {
		return nil, errors.New("version-rolling.mask is not in the parameters")
	}
	mask, err := decodeBigEndian(rawMask.(string))
	if err != nil {
		return nil, fmt.Errorf("value: (%v) err: %w", rawMask, err)
	}
	b, ok := p.Parameters["version-rolling.min-bit-count"]
	mbc := uint64(0)
	if ok {
		mbc = b.(uint64)
		if mbc > 255 {
			return nil, errors.New("min-bit-count > 255")
		}
	}

	return &VersionRollingConfigurationRequest{
		Mask:        mask,
		MinBitCount: byte(mbc),
	}, nil
}

// SetVersionRolling adds [ExtensionVersionRolling] to the [MiningConfigureParams].
func (p *MiningConfigureParams) SetVersionRolling(x VersionRollingConfigurationRequest) error {
	if p.Supports(ExtensionVersionRolling) {
		return errors.New("request already contains version-rolling")
	}

	p.Supported = append(p.Supported, "version-rolling")

	p.Parameters["version-rolling.mask"] = encodeBigEndian(x.Mask)
	p.Parameters["version-rolling.min-bit-count"] = x.MinBitCount

	return nil
}

// GetMinimumDifficulty returns the minimum difficulty, if provided.
func (p *MiningConfigureParams) GetMinimumDifficulty() (float64, error) {
	if diff, ok := p.Parameters["minimum-difficulty.value"]; ok {
		if !validDifficulty(diff) {
			return 0, errors.New("invalid difficulty, must be a float64 and >= 0")
		}
		return diff.(float64), nil
	}
	return 0, errors.New("minimum-difficulty.value is not in the parameters")
}

// SetMinimumDifficulty adds [ExtensionMinimumDifficulty] to the [MiningConfigureParams].
func (p *MiningConfigureParams) SetMinimumDifficulty(diff float64) error {
	if p.Supports(ExtensionMinimumDifficulty) {
		return errors.New("request already contains minimum-difficulty")
	}

	p.Supported = append(p.Supported, "minimum-difficulty")

	p.Parameters["minimum-difficulty.value"] = diff

	return nil
}

// GetSubscribeExtranonce returns whether the client is capable of receiving [MethodMiningSetExtranonce] messages.
// Internally it just calls [MiningConfigureParams.Supports].
func (p *MiningConfigureParams) GetSubscribeExtranonce() bool {
	return p.Supports(ExtensionSubscribeExtranonce)
}

// SetSubscribeExtranonce adds [ExtensionSubscribeExtranonce] to the [MiningConfigureParams].
func (p *MiningConfigureParams) SetSubscribeExtranonce() error {
	if p.Supports(ExtensionSubscribeExtranonce) {
		return errors.New("request already contains subscribe-extranonce")
	}
	p.Supported = append(p.Supported, "subscribe-extranonce")

	return nil
}

// GetInfo returns the [ExtensionInfo] config, if provided.
func (p *MiningConfigureParams) GetInfo() (*InfoConfigurationRequest, error) {
	info := InfoConfigurationRequest{}

	if x, ok := p.Parameters["info.connection-url"]; ok {
		info.ConnectionURL, ok = x.(string)
		if !ok {
			return nil, errors.New("info.connection-url is not a string")
		}
	}

	if x, ok := p.Parameters["info.hw-version"]; ok {
		info.HWVersion, ok = x.(string)
		if !ok {
			return nil, errors.New("info.hw-version is not a string")
		}
	}

	if x, ok := p.Parameters["info.sw-version"]; ok {
		info.SWVersion, ok = x.(string)
		if !ok {
			return nil, errors.New("info.sw-version is not a string")
		}
	}

	if x, ok := p.Parameters["info.hw-id"]; ok {
		info.HWID, ok = x.(string)
		if !ok {
			return nil, errors.New("info.hw-id is not a string")
		}
	}

	return &info, nil
}

// SetInfo adds info to the [MiningConfigureParams].
func (p *MiningConfigureParams) SetInfo(params InfoConfigurationRequest) error {
	if p.Supports(ExtensionInfo) {
		return errors.New("request already contains info")
	}

	p.Supported = append(p.Supported, "info")

	p.Parameters["info.connection-url"] = params.ConnectionURL
	p.Parameters["info.hw-version"] = params.HWVersion
	p.Parameters["info.sw-version"] = params.SWVersion
	p.Parameters["info.hw-id"] = params.HWID

	return nil
}

type MiningConfigureResult map[string]any
type VersionRollingConfigurationResult struct {
	Accepted bool
	Mask     uint32
}

// FromResponse parses the [MiningConfigureResult] from a [Response].
func (p *MiningConfigureResult) FromResponse(r *Response) error {
	result, ok := r.Result.(MiningConfigureResult)
	if !ok {
		return errors.New("invalid result type; should be map[string]interface{}")
	}
	*p = result

	return nil
}

// ToResponse creates a [Response] from the [MiningConfigureResult].
func (p *MiningConfigureResult) ToResponse(id MessageID) *Response {
	return NewResponse(id, p)
}

// Supports checks if the [MiningConfigureResult] contains the given extension.
func (p *MiningConfigureResult) Supports(extension Extension) bool {
	_, ok := (*p)[extension.String()]
	return ok
}

// GetVersionRolling returns the [ExtensionVersionRolling] response from the [MiningConfigureResult].
func (p *MiningConfigureResult) GetVersionRolling() *VersionRollingConfigurationResult {
	v, ok := (*p)["version-rolling"]
	if !ok {
		return nil
	}

	accepted, ok := v.(bool)
	if !ok {
		return nil
	}

	if !accepted {
		return &VersionRollingConfigurationResult{
			Accepted: false,
			Mask:     0,
		}
	}

	m, ok := (*p)["version-rolling.mask"]
	if !ok {
		return nil
	}

	maskstr, ok := m.(string)
	if !ok {
		return nil
	}

	mask, err := decodeBigEndian(maskstr)
	if err != nil {
		return nil
	}

	return &VersionRollingConfigurationResult{
		Accepted: true,
		Mask:     mask,
	}
}

// SetVersionRolling adds the [ExtensionVersionRolling] response to the [MiningConfigureResult].
func (p *MiningConfigureResult) SetVersionRolling(x VersionRollingConfigurationResult) error {
	if _, ok := (*p)["version-rolling"]; ok {
		return errors.New("result already contains version-rolling")
	}

	(*p)["version-rolling"] = x.Accepted
	if x.Accepted {
		(*p)["version-rolling.mask"] = encodeBigEndian(x.Mask)
	}

	return nil
}

// GetMinimumDifficulty returns the [ExtensionMinimumDifficulty] response from the [MiningConfigureResult].
func (p *MiningConfigureResult) GetMinimumDifficulty() bool {
	v, ok := (*p)["minimum-difficulty"]
	if !ok {
		return false
	}

	accepted, ok := v.(bool)
	if !ok {
		return false
	}

	return accepted
}

// SetMinimumDifficulty adds the [ExtensionMinimumDifficulty] response to the [MiningConfigureResult].
func (p *MiningConfigureResult) SetMinimumDifficulty(accepted bool) error {
	if _, ok := (*p)["minimum-difficulty"]; ok {
		return errors.New("result already contains minimum-difficulty")
	}

	(*p)["minimum-difficulty"] = accepted

	return nil
}

// GetSubscribeExtranonce returns the [ExtensionSubscribeExtranonce] response from the [MiningConfigureResult].
func (p *MiningConfigureResult) GetSubscribeExtranonce() bool {
	v, ok := (*p)["subscribe-extranonce"]
	if !ok {
		return false
	}

	accepted, ok := v.(bool)
	if !ok {
		return false
	}

	return accepted
}

// SetSubscribeExtranonce adds the [ExtensionSubscribeExtranonce] response to the [MiningConfigureResult].
func (p *MiningConfigureResult) SetSubscribeExtranonce(accepted bool) error {
	if _, ok := (*p)["subscribe-extranonce"]; ok {
		return errors.New("result already contains subscribe-extranonce")
	}

	(*p)["subscribe-extranonce"] = accepted

	return nil
}

// GetInfo gets the [ExtensionInfo] response from the [MiningConfigureResult].
func (p *MiningConfigureResult) GetInfo() bool {
	v, ok := (*p)["info"]
	if !ok {
		return false
	}

	accepted, ok := v.(bool)
	if !ok {
		return false
	}

	return accepted
}

// SetInfo adds the [ExtensionInfo] response to the [MiningConfigureResult].
func (p *MiningConfigureResult) SetInfo(accepted bool) error {
	if _, ok := (*p)["info"]; ok {
		return errors.New("result already contains info")
	}

	(*p)["info"] = accepted

	return nil
}
