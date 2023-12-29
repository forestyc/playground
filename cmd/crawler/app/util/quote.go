package util

import (
	"github.com/forestyc/playground/cmd/crawler/app/fippb"
	"regexp"
)

// GetVarietyFromContract 截取合约id中的品种id
func GetVarietyFromContract(exchange, contract string) string {
	if exchange != Czce {
		re := regexp.MustCompile(`(\d{2})(0[1-9]|1[0-2])`)
		matches := re.FindAllStringIndex(contract, -1)
		if len(matches) > 0 {
			if len(matches[0]) > 0 {
				return contract[0:matches[0][0]]
			}
		}
	} else {
		// 郑商所，SR401
		re := regexp.MustCompile(`(\d{1})(0[1-9]|1[0-2])`)
		matches := re.FindAllStringIndex(contract, -1)
		if len(matches) > 0 {
			if len(matches[0]) > 0 {
				return contract[0:matches[0][0]]
			}
		}
	}
	return ""
}

// GetOptionTypeFromContract 截取合约id中的期权类型
func GetOptionTypeFromContract(contract string) string {
	reg, _ := regexp.Compile(`[a-zA-Z]+\d+.*[P].+`) //匹配期权P
	if reg.Match([]byte(contract)) {
		return "P"
	}
	reg, _ = regexp.Compile(`[a-zA-Z]+\d+.*[C].+`) //匹配期权P
	if reg.Match([]byte(contract)) {
		return "C"
	}
	return ""
}

// GetQuoteTypeFromContract 获取行情类型
func GetQuoteTypeFromContract(contract string) int {
	if optionType := GetOptionTypeFromContract(contract); len(optionType) > 0 {
		return Option
	} else {
		return Future
	}
}

// ExchangeName 交易所Idz转换
func ExchangeName(exchangeId fippb.EExchgID) string {
	switch exchangeId {
	case fippb.EExchgID_DCE:
		return Dce
	case fippb.EExchgID_SHFE:
		return Shfe
	case fippb.EExchgID_CZCE:
		return Czce
	case fippb.EExchgID_CFFE:
		return Cffex
	case fippb.EExchgID_SGE:
		return Sge
	case fippb.EExchgID_INE:
		return Ine
	case fippb.EExchgID_NYMEX, fippb.EExchgID_NYMEX_REAL_REFINITIV, fippb.EExchgID_NYMEX_DELAY_REFINITIV, fippb.EExchgID_NYMEX_REFINITIV:
		return Nymex
	case fippb.EExchgID_COMEX, fippb.EExchgID_COMEX_REAL_REFINITIV, fippb.EExchgID_COMEX_DELAY_REFINITIV, fippb.EExchgID_COMEX_REFINITIV:
		return Comex
	case fippb.EExchgID_CBOT, fippb.EExchgID_CBOT_REAL_REFINITIV, fippb.EExchgID_CBOT_DELAY_REFINITIV, fippb.EExchgID_CBOT_REFINITIV:
		return Cbot
	case fippb.EExchgID_CME, fippb.EExchgID_CME_REFINITIV, fippb.EExchgID_CME_REAL_REFINITIV, fippb.EExchgID_CME_DELAY_REFINITIV:
		return Cme
	case fippb.EExchgID_CBOE:
		return Cboe
	case fippb.EExchgID_US_ICE:
		return UsIce
	case fippb.EExchgID_EUR_ICE:
		return EurIce
	case fippb.EExchgID_LME:
		return Lme
	case fippb.EExchgID_HKEX:
		return Hkex
	case fippb.EExchgID_SGX:
		return Sgx
	case fippb.EExchgID_TOCOM:
		return Tocom
	case fippb.EExchgID_BMD:
		return Bmd
	case fippb.EExchgID_APEX:
		return Apex
	case fippb.EExchgID_CSI:
		return Csi
	case fippb.EExchgID_ASX:
		return Asx
	case fippb.EExchgID_SFE:
		return Sfe
	case fippb.EExchgID_DGCX:
		return Dgcx
	case fippb.EExchgID_DME:
		return Dme
	case fippb.EExchgID_ICE_SG:
		return IceSg
	case fippb.EExchgID_TFEX:
		return Tfex
	case fippb.EExchgID_KRX:
		return Krx
	case fippb.EExchgID_JPX:
		return Jpx
	case fippb.EExchgID_XEUREX:
		return Xeurex
	case fippb.EExchgID_SHSE:
		return Shse
	case fippb.EExchgID_SZSE:
		return Szse
	case fippb.EExchgID_BJSE:
		return Bjse
	case fippb.EExchgID_NEEQ:
		return Neeq
	case fippb.EExchgID_GFEX:
		return Gfex
	default:
		return ""
	}
}
