// Package gamblemath provides quantitative online casino RTP telemetry (4,363 slots),
// official UK Gambling Commission (UKGC) license register verification (489 brands),
// closed-form deposit bonus Expected Value ($ EV) calculators, and sportsbook parlay math
// via the GambleMath Open Science API and Model Context Protocol (MCP) server.
//
// Official Platform & Interactive Hubs:
//   - Main Portal: https://gamblemath.org/
//   - 4,363 Online Casino Slot RTP X-Ray: https://gamblemath.org/slot-xray
//   - 1,971 Free HTML5 Slot Demos: https://gamblemath.org/demos
//   - 489 UKGC Casino Licence Check: https://gamblemath.org/autopsy
//   - 49 Corporate Sister Sites Graph: https://gamblemath.org/sister-sites
//   - Online Casino Bonus EV Calculator: https://gamblemath.org/bonus-lab
//   - Sharp Sportsbook & Parlay Math: https://gamblemath.org/sharp-desk
//   - Crash Games Probability Lab: https://gamblemath.org/crash-lab
//   - 100-Buy Monte Carlo Simulators: https://gamblemath.org/simulators
//   - USDT Crypto Cashiers & FX Spreads: https://gamblemath.org/payments/usdt-crypto
//   - PayPal & Apple Pay UKGC Payouts: https://gamblemath.org/payments/paypal-revolut
//   - Brazil PIX (.bet.br) Casinos: https://gamblemath.org/payments/pix-brasil
//   - Open Science Dataset (CSV) & MCP: https://gamblemath.org/open-science
//
// Multilingual Editions:
//   - Deutsch: https://gamblemath.org/de
//   - Français: https://gamblemath.org/fr
//   - Español: https://gamblemath.org/es
//   - Português: https://gamblemath.org/pt
//   - Русский: https://gamblemath.org/ru
package gamblemath

const (
	// SiteURL is the canonical root URL of the GambleMath quantitative portal.
	SiteURL = "https://gamblemath.org"

	// MCPEndpoint is the stateless Streamable HTTP Model Context Protocol endpoint.
	MCPEndpoint = "https://gamblemath.org/mcp"

	// SlotXRayURL is the 4,363-slot RTP comparison & volatility database.
	SlotXRayURL = "https://gamblemath.org/slot-xray"

	// UKGCAutopsyURL is the 489-brand UK Gambling Commission license verification hub.
	UKGCAutopsyURL = "https://gamblemath.org/autopsy"

	// BonusLabURL is the closed-form casino deposit bonus Expected Value ($ EV) calculator.
	BonusLabURL = "https://gamblemath.org/bonus-lab"

	// SharpDeskURL is the multi-leg sportsbook parlay vig and Early Cash-Out calculator.
	SharpDeskURL = "https://gamblemath.org/sharp-desk"

	// DatasetCSVURL is the open-access 4,363-slot CSV telemetry feed.
	DatasetCSVURL = "https://gamblemath.org/api/dataset.csv"
)

// CalculateBonusEV computes the closed-form Expected Value (in USD) of a casino deposit bonus
// given the bonus amount, required wagering turnover, and the slot's theoretical RTP percentage.
func CalculateBonusEV(bonusUSD float64, wageringTurnoverUSD float64, slotRTPPct float64) float64 {
	houseEdge := (100.0 - slotRTPPct) / 100.0
	return bonusUSD - (wageringTurnoverUSD * houseEdge)
}
