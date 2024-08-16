package sxid

// Defines the SXID error type.
// ref. https://docs.nvidia.com/datacenter/tesla/pdf/fabric-manager-user-guide.pdf
type Detail struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	PotentialFatal bool   `json:"potential_fatal"`
	AlwaysFatal    bool   `json:"always_fatal"`
	Impact         string `json:"impact"`
	Recovery       string `json:"recovery"`
	OtherImpact    string `json:"other_impact"`
}

// Returns the error if found.
// Otherwise, returns false.
func GetDetail(id int) (*Detail, bool) {
	e, ok := details[id]
	return &e, ok
}

// These are copied from:
// "D.4 Non-Fatal NVSwitch SXid Errors"
// "D.5 Fatal NVSwitch SXid Errors"
// "D.6 Always Fatal NVSwitch SXid Errors"
// "D.7 Other Notable NVSwitch SXid Errors"
// ref. https://docs.nvidia.com/datacenter/tesla/pdf/fabric-manager-user-guide.pdf
var details = map[int]Detail{
	// D.4 Non-Fatal NVSwitch SXid Errors
	11004: {
		ID:             11004,
		Name:           "Ingress invalid ACL",
		Description:    "This SXid error can happen only because of an incorrect FM partition configuration and is expected not to occur in the field.",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "Corresponding GPU NVLink traffic will be stalled, and the subsequent GPU access will hang. The GPU driver on the guest VM will abort CUDA jobs with Xid 45.",
		Recovery:       "Validate GPU/NVSwitch fabric partition routing information using the NVSwitch-audit tool. Restart the guest VM.",
		OtherImpact:    "If the error is observed on a Trunk port, partitions that are using NVSwitch trunk ports will be affected.",
	},
	11012: {
		ID:             11012,
		Name:           "Single bit ECC errors",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "No guest VM impact because the NVSwitch hardware will auto correct the ECC errors.",
		Recovery:       "Not Applicable.",
		OtherImpact:    "No Impact.",
	},
	11021: {
		ID:             11021,
		Name:           "Single bit ECC errors",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "No guest VM impact because the NVSwitch hardware will auto correct the ECC errors.",
		Recovery:       "Not Applicable.",
		OtherImpact:    "No Impact.",
	},
	11022: {
		ID:             11022,
		Name:           "Single bit ECC errors",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "No guest VM impact because the NVSwitch hardware will auto correct the ECC errors.",
		Recovery:       "Not Applicable.",
		OtherImpact:    "No Impact.",
	},
	11023: {
		ID:             11023,
		Name:           "Single bit ECC errors",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "No guest VM impact because the NVSwitch hardware will auto correct the ECC errors.",
		Recovery:       "Not Applicable.",
		OtherImpact:    "No Impact.",
	},
	12021: {
		ID:             12021,
		Name:           "Single bit ECC errors",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "No guest VM impact because the NVSwitch hardware will auto correct the ECC errors.",
		Recovery:       "Not Applicable.",
		OtherImpact:    "No Impact.",
	},
	12023: {
		ID:             12023,
		Name:           "Single bit ECC errors",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "No guest VM impact because the NVSwitch hardware will auto correct the ECC errors.",
		Recovery:       "Not Applicable.",
		OtherImpact:    "No Impact.",
	},
	15008: {
		ID:             15008,
		Name:           "Single bit ECC errors",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "No guest VM impact because the NVSwitch hardware will auto correct the ECC errors.",
		Recovery:       "Not Applicable.",
		OtherImpact:    "No Impact.",
	},
	15011: {
		ID:             15011,
		Name:           "Single bit ECC errors",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "No guest VM impact because the NVSwitch hardware will auto correct the ECC errors.",
		Recovery:       "Not Applicable.",
		OtherImpact:    "No Impact.",
	},
	19049: {
		ID:             19049,
		Name:           "Single bit ECC errors",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "No guest VM impact because the NVSwitch hardware will auto correct the ECC errors.",
		Recovery:       "Not Applicable.",
		OtherImpact:    "No Impact.",
	},
	19055: {
		ID:             19055,
		Name:           "Single bit ECC errors",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "No guest VM impact because the NVSwitch hardware will auto correct the ECC errors.",
		Recovery:       "Not Applicable.",
		OtherImpact:    "No Impact.",
	},
	19057: {
		ID:             19057,
		Name:           "Single bit ECC errors",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "No guest VM impact because the NVSwitch hardware will auto correct the ECC errors.",
		Recovery:       "Not Applicable.",
		OtherImpact:    "No Impact.",
	},
	19059: {
		ID:             19059,
		Name:           "Single bit ECC errors",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "No guest VM impact because the NVSwitch hardware will auto correct the ECC errors.",
		Recovery:       "Not Applicable.",
		OtherImpact:    "No Impact.",
	},
	19062: {
		ID:             19062,
		Name:           "Single bit ECC errors",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "No guest VM impact because the NVSwitch hardware will auto correct the ECC errors.",
		Recovery:       "Not Applicable.",
		OtherImpact:    "No Impact.",
	},
	19065: {
		ID:             19065,
		Name:           "Single bit ECC errors",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "No guest VM impact because the NVSwitch hardware will auto correct the ECC errors.",
		Recovery:       "Not Applicable.",
		OtherImpact:    "No Impact.",
	},
	19068: {
		ID:             19068,
		Name:           "Single bit ECC errors",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "No guest VM impact because the NVSwitch hardware will auto correct the ECC errors.",
		Recovery:       "Not Applicable.",
		OtherImpact:    "No Impact.",
	},
	19071: {
		ID:             19071,
		Name:           "Single bit ECC errors",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "No guest VM impact because the NVSwitch hardware will auto correct the ECC errors.",
		Recovery:       "Not Applicable.",
		OtherImpact:    "No Impact.",
	},
	24001: {
		ID:             24001,
		Name:           "Single bit ECC errors",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "No guest VM impact because the NVSwitch hardware will auto correct the ECC errors.",
		Recovery:       "Not Applicable.",
		OtherImpact:    "No Impact.",
	},
	24002: {
		ID:             24002,
		Name:           "Single bit ECC errors",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "No guest VM impact because the NVSwitch hardware will auto correct the ECC errors.",
		Recovery:       "Not Applicable.",
		OtherImpact:    "No Impact.",
	},
	24003: {
		ID:             24003,
		Name:           "Single bit ECC errors",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "No guest VM impact because the NVSwitch hardware will auto correct the ECC errors.",
		Recovery:       "Not Applicable.",
		OtherImpact:    "No Impact.",
	},
	20001: {
		ID:             20001,
		Name:           "TX Replay Error",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "NVLink packet needs to be retransmitted. This error might impact the NVLink throughput of the specified port.",
		Recovery:       "Not Applicable.",
		OtherImpact:    "If the error is observed on a Trunk port, the partitions that are using NVSwitch trunk ports might see throughput impact.",
	},
	12028: {
		ID:             12028,
		Name:           "egress nonposted PRIV error",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "Corresponding GPU NVLink traffic will be stalled, and subsequent GPU access will hang. The GPU driver on the guest VM will abort CUDA jobs with Xid 45.",
		Recovery:       "Restart Guest VM.",
		OtherImpact:    "If the error is observed on a Trunk port, the partitions that are using NVSwitch trunk ports will be affected.",
	},
	19084: {
		ID:             19084,
		Name:           "AN1 Heartbeat Timeout Error",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "This error is usually accompanied by a fatal SXid error that will affect the corresponding GPU NVLink traffic.",
		Recovery:       "Reset all GPUs and all NVSwitches (refer to section D.9).",
		OtherImpact:    "If the error is observed on a Trunk port, the partitions that are using NVSwitch trunk ports will be affected.",
	},
	22013: {
		ID:             22013,
		Name:           "Minion Link DLREQ interrupt",
		Description:    "",
		PotentialFatal: false,
		AlwaysFatal:    false,
		Impact:         "This SXid can be safely ignored.",
		Recovery:       "Not Applicable.",
		OtherImpact:    "No Impact.",
	},

	// D.5 Potential Fatal NVSwitch SXid Errors
	11001: {
		ID:             11001,
		Name:           "ingress invalid command",
		Description:    defaultPotentialFatalErr.Description,
		PotentialFatal: defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:    defaultPotentialFatalErr.AlwaysFatal,
		Impact:         defaultPotentialFatalErr.Impact,
		Recovery:       defaultPotentialFatalErr.Recovery,
		OtherImpact:    defaultPotentialFatalErr.OtherImpact,
	},
	11009: {
		ID:             11009,
		Name:           "ingress invalid VCSet",
		Description:    defaultPotentialFatalErr.Description,
		PotentialFatal: defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:    defaultPotentialFatalErr.AlwaysFatal,
		Impact:         defaultPotentialFatalErr.Impact,
		Recovery:       defaultPotentialFatalErr.Recovery,
		OtherImpact:    defaultPotentialFatalErr.OtherImpact,
	},
	11013: {
		ID:             11013,
		Name:           "ingress header DBE",
		Description:    defaultPotentialFatalErr.Description,
		PotentialFatal: defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:    defaultPotentialFatalErr.AlwaysFatal,
		Impact:         defaultPotentialFatalErr.Impact,
		Recovery:       defaultPotentialFatalErr.Recovery,
		OtherImpact:    defaultPotentialFatalErr.OtherImpact,
	},
	11018: {
		ID:             11018,
		Name:           "ingress RID DBE",
		Description:    defaultPotentialFatalErr.Description,
		PotentialFatal: defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:    defaultPotentialFatalErr.AlwaysFatal,
		Impact:         defaultPotentialFatalErr.Impact,
		Recovery:       defaultPotentialFatalErr.Recovery,
		OtherImpact:    defaultPotentialFatalErr.OtherImpact,
	},
	11019: {
		ID:             11019,
		Name:           "ingress RLAN DBE",
		Description:    defaultPotentialFatalErr.Description,
		PotentialFatal: defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:    defaultPotentialFatalErr.AlwaysFatal,
		Impact:         defaultPotentialFatalErr.Impact,
		Recovery:       defaultPotentialFatalErr.Recovery,
		OtherImpact:    defaultPotentialFatalErr.OtherImpact,
	},
	11020: {
		ID:             11020,
		Name:           "ingress control parity",
		Description:    defaultPotentialFatalErr.Description,
		PotentialFatal: defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:    defaultPotentialFatalErr.AlwaysFatal,
		Impact:         defaultPotentialFatalErr.Impact,
		Recovery:       defaultPotentialFatalErr.Recovery,
		OtherImpact:    defaultPotentialFatalErr.OtherImpact,
	},
	12001: {
		ID:             12001,
		Name:           "egress crossbar overflow",
		Description:    defaultPotentialFatalErr.Description,
		PotentialFatal: defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:    defaultPotentialFatalErr.AlwaysFatal,
		Impact:         defaultPotentialFatalErr.Impact,
		Recovery:       defaultPotentialFatalErr.Recovery,
		OtherImpact:    defaultPotentialFatalErr.OtherImpact,
	},
	12002: {
		ID:             12002,
		Name:           "egress packet route",
		Description:    defaultPotentialFatalErr.Description,
		PotentialFatal: defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:    defaultPotentialFatalErr.AlwaysFatal,
		Impact:         defaultPotentialFatalErr.Impact,
		Recovery:       defaultPotentialFatalErr.Recovery,
		OtherImpact:    defaultPotentialFatalErr.OtherImpact,
	},
	12022: {
		ID:             12022,
		Name:           "egress input ECC DBE error",
		Description:    defaultPotentialFatalErr.Description,
		PotentialFatal: defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:    defaultPotentialFatalErr.AlwaysFatal,
		Impact:         defaultPotentialFatalErr.Impact,
		Recovery:       defaultPotentialFatalErr.Recovery,
		OtherImpact:    defaultPotentialFatalErr.OtherImpact,
	},
	12024: {
		ID:             12024,
		Name:           "egress output ECC DBE error",
		Description:    defaultPotentialFatalErr.Description,
		PotentialFatal: defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:    defaultPotentialFatalErr.AlwaysFatal,
		Impact:         defaultPotentialFatalErr.Impact,
		Recovery:       defaultPotentialFatalErr.Recovery,
		OtherImpact:    defaultPotentialFatalErr.OtherImpact,
	},
	12025: {
		ID:             12025,
		Name:           "egress credit overflow",
		Description:    defaultPotentialFatalErr.Description,
		PotentialFatal: defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:    defaultPotentialFatalErr.AlwaysFatal,
		Impact:         defaultPotentialFatalErr.Impact,
		Recovery:       defaultPotentialFatalErr.Recovery,
		OtherImpact:    defaultPotentialFatalErr.OtherImpact,
	},
	12026: {
		ID:             12026,
		Name:           "egress destination request ID error",
		Description:    defaultPotentialFatalErr.Description,
		PotentialFatal: defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:    defaultPotentialFatalErr.AlwaysFatal,
		Impact:         defaultPotentialFatalErr.Impact,
		Recovery:       defaultPotentialFatalErr.Recovery,
		OtherImpact:    defaultPotentialFatalErr.OtherImpact,
	},
	12027: {
		ID:             12027,
		Name:           "egress destination response ID error",
		Description:    defaultPotentialFatalErr.Description,
		PotentialFatal: defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:    defaultPotentialFatalErr.AlwaysFatal,
		Impact:         defaultPotentialFatalErr.Impact,
		Recovery:       defaultPotentialFatalErr.Recovery,
		OtherImpact:    defaultPotentialFatalErr.OtherImpact,
	},
	12030: {
		ID:             12030,
		Name:           "egress control parity error",
		Description:    defaultPotentialFatalErr.Description,
		PotentialFatal: defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:    defaultPotentialFatalErr.AlwaysFatal,
		Impact:         defaultPotentialFatalErr.Impact,
		Recovery:       defaultPotentialFatalErr.Recovery,
		OtherImpact:    defaultPotentialFatalErr.OtherImpact,
	},
	12031: {
		ID:             12031,
		Name:           "egress credit parity error",
		Description:    defaultPotentialFatalErr.Description,
		PotentialFatal: defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:    defaultPotentialFatalErr.AlwaysFatal,
		Impact:         defaultPotentialFatalErr.Impact,
		Recovery:       defaultPotentialFatalErr.Recovery,
		OtherImpact:    defaultPotentialFatalErr.OtherImpact,
	},
	12032: {
		ID:             12032,
		Name:           "egress flit type mismatch",
		Description:    defaultPotentialFatalErr.Description,
		PotentialFatal: defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:    defaultPotentialFatalErr.AlwaysFatal,
		Impact:         defaultPotentialFatalErr.Impact,
		Recovery:       defaultPotentialFatalErr.Recovery,
		OtherImpact:    defaultPotentialFatalErr.OtherImpact,
	},
	14017: {
		ID:             14017,
		Name:           "TS ATO timeout",
		Description:    defaultPotentialFatalErr.Description,
		PotentialFatal: defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:    defaultPotentialFatalErr.AlwaysFatal,
		Impact:         defaultPotentialFatalErr.Impact,
		Recovery:       defaultPotentialFatalErr.Recovery,
		OtherImpact:    defaultPotentialFatalErr.OtherImpact,
	},
	15001: {
		ID:             15001,
		Name:           "route buffer over/underflow",
		Description:    defaultPotentialFatalErr.Description,
		PotentialFatal: defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:    defaultPotentialFatalErr.AlwaysFatal,
		Impact:         defaultPotentialFatalErr.Impact,
		Recovery:       defaultPotentialFatalErr.Recovery,
		OtherImpact:    defaultPotentialFatalErr.OtherImpact,
	},
	15006: {
		ID:             15006,
		Name:           "route transdone over/underflow",
		Description:    defaultPotentialFatalErr.Description,
		PotentialFatal: defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:    defaultPotentialFatalErr.AlwaysFatal,
		Impact:         defaultPotentialFatalErr.Impact,
		Recovery:       defaultPotentialFatalErr.Recovery,
		OtherImpact:    defaultPotentialFatalErr.OtherImpact,
	},
	15009: {
		ID:             15009,
		Name:           "route GLT DBE",
		Description:    defaultPotentialFatalErr.Description,
		PotentialFatal: defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:    defaultPotentialFatalErr.AlwaysFatal,
		Impact:         defaultPotentialFatalErr.Impact,
		Recovery:       defaultPotentialFatalErr.Recovery,
		OtherImpact:    defaultPotentialFatalErr.OtherImpact,
	},
	15010: {
		ID:             15010,
		Name:           "route parity",
		Description:    defaultPotentialFatalErr.Description,
		PotentialFatal: defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:    defaultPotentialFatalErr.AlwaysFatal,
		Impact:         defaultPotentialFatalErr.Impact,
		Recovery:       defaultPotentialFatalErr.Recovery,
		OtherImpact:    defaultPotentialFatalErr.OtherImpact,
	},
	15012: {
		ID:             15012,
		Name:           "route incoming DBE",
		Description:    defaultPotentialFatalErr.Description,
		PotentialFatal: defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:    defaultPotentialFatalErr.AlwaysFatal,
		Impact:         defaultPotentialFatalErr.Impact,
		Recovery:       defaultPotentialFatalErr.Recovery,
		OtherImpact:    defaultPotentialFatalErr.OtherImpact,
	},
	15013: {
		ID:             15013,
		Name:           "route credit parity",
		Description:    defaultPotentialFatalErr.Description,
		PotentialFatal: defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:    defaultPotentialFatalErr.AlwaysFatal,
		Impact:         defaultPotentialFatalErr.Impact,
		Recovery:       defaultPotentialFatalErr.Recovery,
		OtherImpact:    defaultPotentialFatalErr.OtherImpact,
	},
	19047: {
		ID:             19047,
		Name:           "NCISOC HDR ECC DBE Error",
		Description:    defaultPotentialFatalErr.Description,
		PotentialFatal: defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:    defaultPotentialFatalErr.AlwaysFatal,
		Impact:         defaultPotentialFatalErr.Impact,
		Recovery:       defaultPotentialFatalErr.Recovery,
		OtherImpact:    defaultPotentialFatalErr.OtherImpact,
	},
	19048: {
		ID:             19048,
		Name:           "NCISOC DAT ECC DBE Error",
		Description:    defaultPotentialFatalErr.Description,
		PotentialFatal: defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:    defaultPotentialFatalErr.AlwaysFatal,
		Impact:         defaultPotentialFatalErr.Impact,
		Recovery:       defaultPotentialFatalErr.Recovery,
		OtherImpact:    defaultPotentialFatalErr.OtherImpact,
	},
	19054: {
		ID:             19054,
		Name:           "HDR RAM ECC DBE Error",
		Description:    defaultPotentialFatalErr.Description,
		PotentialFatal: defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:    defaultPotentialFatalErr.AlwaysFatal,
		Impact:         defaultPotentialFatalErr.Impact,
		Recovery:       defaultPotentialFatalErr.Recovery,
		OtherImpact:    defaultPotentialFatalErr.OtherImpact,
	},
	19056: {
		ID:             19056,
		Name:           "DAT0 RAM ECC DBE Error",
		Description:    defaultPotentialFatalErr.Description,
		PotentialFatal: defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:    defaultPotentialFatalErr.AlwaysFatal,
		Impact:         defaultPotentialFatalErr.Impact,
		Recovery:       defaultPotentialFatalErr.Recovery,
		OtherImpact:    defaultPotentialFatalErr.OtherImpact,
	},
	19058: {
		ID:             19058,
		Name:           "DAT1 RAM ECC DBE Error",
		Description:    defaultPotentialFatalErr.Description,
		PotentialFatal: defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:    defaultPotentialFatalErr.AlwaysFatal,
		Impact:         defaultPotentialFatalErr.Impact,
		Recovery:       defaultPotentialFatalErr.Recovery,
		OtherImpact:    defaultPotentialFatalErr.OtherImpact,
	},
	19060: {
		ID:             19060,
		Name:           "CREQ RAM HDR ECC DBE Error",
		Description:    defaultPotentialFatalErr.Description,
		PotentialFatal: defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:    defaultPotentialFatalErr.AlwaysFatal,
		Impact:         defaultPotentialFatalErr.Impact,
		Recovery:       defaultPotentialFatalErr.Recovery,
		OtherImpact:    defaultPotentialFatalErr.OtherImpact,
	},
	19061: {
		ID:             19061,
		Name:           "CREQ RAM DAT ECC DBE Error",
		Description:    defaultPotentialFatalErr.Description,
		PotentialFatal: defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:    defaultPotentialFatalErr.AlwaysFatal,
		Impact:         defaultPotentialFatalErr.Impact,
		Recovery:       defaultPotentialFatalErr.Recovery,
		OtherImpact:    defaultPotentialFatalErr.OtherImpact,
	},
	19063: {
		ID:             19063,
		Name:           "Response RAM HDR ECC DBE Error",
		Description:    defaultPotentialFatalErr.Description,
		PotentialFatal: defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:    defaultPotentialFatalErr.AlwaysFatal,
		Impact:         defaultPotentialFatalErr.Impact,
		Recovery:       defaultPotentialFatalErr.Recovery,
		OtherImpact:    defaultPotentialFatalErr.OtherImpact,
	},
	19064: {
		ID:             19064,
		Name:           "Response RAM DAT ECC DBE Error",
		Description:    defaultPotentialFatalErr.Description,
		PotentialFatal: defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:    defaultPotentialFatalErr.AlwaysFatal,
		Impact:         defaultPotentialFatalErr.Impact,
		Recovery:       defaultPotentialFatalErr.Recovery,
		OtherImpact:    defaultPotentialFatalErr.OtherImpact,
	},
	19066: {
		ID:             19066,
		Name:           "COM RAM HDR ECC DBE Error",
		Description:    defaultPotentialFatalErr.Description,
		PotentialFatal: defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:    defaultPotentialFatalErr.AlwaysFatal,
		Impact:         defaultPotentialFatalErr.Impact,
		Recovery:       defaultPotentialFatalErr.Recovery,
		OtherImpact:    defaultPotentialFatalErr.OtherImpact,
	},
	19067: {
		ID:             19067,
		Name:           "COM RAM DAT ECC DBE Error",
		Description:    defaultPotentialFatalErr.Description,
		PotentialFatal: defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:    defaultPotentialFatalErr.AlwaysFatal,
		Impact:         defaultPotentialFatalErr.Impact,
		Recovery:       defaultPotentialFatalErr.Recovery,
		OtherImpact:    defaultPotentialFatalErr.OtherImpact,
	},
	19069: {
		ID:             19069,
		Name:           "RSP1 RAM HDR ECC DBE Error",
		Description:    defaultPotentialFatalErr.Description,
		PotentialFatal: defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:    defaultPotentialFatalErr.AlwaysFatal,
		Impact:         defaultPotentialFatalErr.Impact,
		Recovery:       defaultPotentialFatalErr.Recovery,
		OtherImpact:    defaultPotentialFatalErr.OtherImpact,
	},
	19070: {
		ID:             19070,
		Name:           "RSP1 RAM DAT ECC DBE Error",
		Description:    defaultPotentialFatalErr.Description,
		PotentialFatal: defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:    defaultPotentialFatalErr.AlwaysFatal,
		Impact:         defaultPotentialFatalErr.Impact,
		Recovery:       defaultPotentialFatalErr.Recovery,
		OtherImpact:    defaultPotentialFatalErr.OtherImpact,
	},
	20034: {
		ID:             20034,
		Name:           "LTSSM Fault Up",
		Description:    defaultPotentialFatalErr.Description,
		PotentialFatal: defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:    defaultPotentialFatalErr.AlwaysFatal,
		Impact: defaultPotentialFatalErr.Impact + `
Guest VM Impact: This SXid is triggered whenever the associated link has gone
down from active. This interrupt is usually associated with other NVLink errors.
`,
		Recovery: defaultPotentialFatalErr.Recovery + `
Guest VM Recovery: In case of A100, restart the VM. In case of H100, reset the
GPU (refer to section D.9). If issue persists, report GPU issues.
`,
		OtherImpact: defaultPotentialFatalErr.OtherImpact + `
Other Guest VM Impact: No impact if error is confined to a single GPU.
`,
	},
	22012: { // in both D.4 and D.5, treat it as potential fatal
		ID:             22012,
		Name:           "Minion Link NA interrupt",
		Description:    "",
		PotentialFatal: true,
		AlwaysFatal:    false,
		Impact:         "This error could occur due to a broken/inconsistent connection or uncoordinated shutdown.",
		Recovery:       "If this issue was not due to an uncoordinated shutdown, check link mechanical connections.",
		OtherImpact:    "No impact if error is confined to a single GPU.",
	},
	24004: {
		ID:             24004,
		Name:           "sourcetrack TCEN0 crubmstore DBE",
		Description:    defaultPotentialFatalErr.Description,
		PotentialFatal: defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:    defaultPotentialFatalErr.AlwaysFatal,
		Impact:         defaultPotentialFatalErr.Impact,
		Recovery:       defaultPotentialFatalErr.Recovery,
		OtherImpact:    defaultPotentialFatalErr.OtherImpact,
	},
	24005: {
		ID:             24005,
		Name:           "sourcetrack TCEN0 TD crubmstore DBE",
		Description:    defaultPotentialFatalErr.Description,
		PotentialFatal: defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:    defaultPotentialFatalErr.AlwaysFatal,
		Impact:         defaultPotentialFatalErr.Impact,
		Recovery:       defaultPotentialFatalErr.Recovery,
		OtherImpact:    defaultPotentialFatalErr.OtherImpact,
	},
	24006: {
		ID:             24006,
		Name:           "sourcetrack TCEN1 crubmstore DBE",
		Description:    defaultPotentialFatalErr.Description,
		PotentialFatal: defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:    defaultPotentialFatalErr.AlwaysFatal,
		Impact:         defaultPotentialFatalErr.Impact,
		Recovery:       defaultPotentialFatalErr.Recovery,
		OtherImpact:    defaultPotentialFatalErr.OtherImpact,
	},
	24007: {
		ID:             24007,
		Name:           "sourcetrack timeout error",
		Description:    defaultPotentialFatalErr.Description,
		PotentialFatal: defaultPotentialFatalErr.PotentialFatal,
		AlwaysFatal:    defaultPotentialFatalErr.AlwaysFatal,
		Impact:         defaultPotentialFatalErr.Impact,
		Recovery:       defaultPotentialFatalErr.Recovery,
		OtherImpact:    defaultPotentialFatalErr.OtherImpact,
	},

	// D.6 Always Fatal NVSwitch SXid Errors
	12020: {
		ID:             12020,
		Name:           "egress sequence ID error",
		Description:    defaultAlwaysFatalErr.Description,
		PotentialFatal: defaultAlwaysFatalErr.PotentialFatal,
		AlwaysFatal:    defaultAlwaysFatalErr.AlwaysFatal,
		Impact:         defaultAlwaysFatalErr.Impact,
		Recovery:       defaultAlwaysFatalErr.Recovery,
		OtherImpact:    defaultAlwaysFatalErr.OtherImpact,
	},
	22003: {
		ID:             22003,
		Name:           "Minion Halt",
		Description:    defaultAlwaysFatalErr.Description,
		PotentialFatal: defaultAlwaysFatalErr.PotentialFatal,
		AlwaysFatal:    defaultAlwaysFatalErr.AlwaysFatal,
		Impact:         defaultAlwaysFatalErr.Impact,
		Recovery:       defaultAlwaysFatalErr.Recovery,
		OtherImpact:    defaultAlwaysFatalErr.OtherImpact,
	},
	22011: {
		ID:             22011,
		Name:           "Minion exterror",
		Description:    defaultAlwaysFatalErr.Description,
		PotentialFatal: defaultAlwaysFatalErr.PotentialFatal,
		AlwaysFatal:    defaultAlwaysFatalErr.AlwaysFatal,
		Impact:         defaultAlwaysFatalErr.Impact,
		Recovery:       defaultAlwaysFatalErr.Recovery,
		OtherImpact:    defaultAlwaysFatalErr.OtherImpact,
	},
	23001: {
		ID:             23001,
		Name:           "ingress SRC-VC buffer overflow",
		Description:    defaultAlwaysFatalErr.Description,
		PotentialFatal: defaultAlwaysFatalErr.PotentialFatal,
		AlwaysFatal:    defaultAlwaysFatalErr.AlwaysFatal,
		Impact:         defaultAlwaysFatalErr.Impact,
		Recovery:       defaultAlwaysFatalErr.Recovery,
		OtherImpact:    defaultAlwaysFatalErr.OtherImpact,
	},
	23002: {
		ID:             23002,
		Name:           "ingress SRC-VC buffer underflow",
		Description:    defaultAlwaysFatalErr.Description,
		PotentialFatal: defaultAlwaysFatalErr.PotentialFatal,
		AlwaysFatal:    defaultAlwaysFatalErr.AlwaysFatal,
		Impact:         defaultAlwaysFatalErr.Impact,
		Recovery:       defaultAlwaysFatalErr.Recovery,
		OtherImpact:    defaultAlwaysFatalErr.OtherImpact,
	},
	23003: {
		ID:             23003,
		Name:           "egress DST-VC credit overflow",
		Description:    defaultAlwaysFatalErr.Description,
		PotentialFatal: defaultAlwaysFatalErr.PotentialFatal,
		AlwaysFatal:    defaultAlwaysFatalErr.AlwaysFatal,
		Impact:         defaultAlwaysFatalErr.Impact,
		Recovery:       defaultAlwaysFatalErr.Recovery,
		OtherImpact:    defaultAlwaysFatalErr.OtherImpact,
	},
	23004: {
		ID:             23004,
		Name:           "egress DST-VC credit underflow",
		Description:    defaultAlwaysFatalErr.Description,
		PotentialFatal: defaultAlwaysFatalErr.PotentialFatal,
		AlwaysFatal:    defaultAlwaysFatalErr.AlwaysFatal,
		Impact:         defaultAlwaysFatalErr.Impact,
		Recovery:       defaultAlwaysFatalErr.Recovery,
		OtherImpact:    defaultAlwaysFatalErr.OtherImpact,
	},
	23005: {
		ID:             23005,
		Name:           "ingress packet burst error",
		Description:    defaultAlwaysFatalErr.Description,
		PotentialFatal: defaultAlwaysFatalErr.PotentialFatal,
		AlwaysFatal:    defaultAlwaysFatalErr.AlwaysFatal,
		Impact:         defaultAlwaysFatalErr.Impact,
		Recovery:       defaultAlwaysFatalErr.Recovery,
		OtherImpact:    defaultAlwaysFatalErr.OtherImpact,
	},
	23006: {
		ID:             23006,
		Name:           "ingress packet sticky error",
		Description:    defaultAlwaysFatalErr.Description,
		PotentialFatal: defaultAlwaysFatalErr.PotentialFatal,
		AlwaysFatal:    defaultAlwaysFatalErr.AlwaysFatal,
		Impact:         defaultAlwaysFatalErr.Impact,
		Recovery:       defaultAlwaysFatalErr.Recovery,
		OtherImpact:    defaultAlwaysFatalErr.OtherImpact,
	},
	23007: {
		ID:             23007,
		Name:           "possible bubbles at ingress",
		Description:    defaultAlwaysFatalErr.Description,
		PotentialFatal: defaultAlwaysFatalErr.PotentialFatal,
		AlwaysFatal:    defaultAlwaysFatalErr.AlwaysFatal,
		Impact:         defaultAlwaysFatalErr.Impact,
		Recovery:       defaultAlwaysFatalErr.Recovery,
		OtherImpact:    defaultAlwaysFatalErr.OtherImpact,
	},
	23008: {
		ID:             23008,
		Name:           "ingress packet invalid dst error",
		Description:    defaultAlwaysFatalErr.Description,
		PotentialFatal: defaultAlwaysFatalErr.PotentialFatal,
		AlwaysFatal:    defaultAlwaysFatalErr.AlwaysFatal,
		Impact:         defaultAlwaysFatalErr.Impact,
		Recovery:       defaultAlwaysFatalErr.Recovery,
		OtherImpact:    defaultAlwaysFatalErr.OtherImpact,
	},
	23009: {
		ID:             23009,
		Name:           "ingress packet parity error",
		Description:    defaultAlwaysFatalErr.Description,
		PotentialFatal: defaultAlwaysFatalErr.PotentialFatal,
		AlwaysFatal:    defaultAlwaysFatalErr.AlwaysFatal,
		Impact:         defaultAlwaysFatalErr.Impact,
		Recovery:       defaultAlwaysFatalErr.Recovery,
		OtherImpact:    defaultAlwaysFatalErr.OtherImpact,
	},
	23010: {
		ID:             23010,
		Name:           "ingress SRC-VC buffer overflow",
		Description:    defaultAlwaysFatalErr.Description,
		PotentialFatal: defaultAlwaysFatalErr.PotentialFatal,
		AlwaysFatal:    defaultAlwaysFatalErr.AlwaysFatal,
		Impact:         defaultAlwaysFatalErr.Impact,
		Recovery:       defaultAlwaysFatalErr.Recovery,
		OtherImpact:    defaultAlwaysFatalErr.OtherImpact,
	},
	23011: {
		ID:             23011,
		Name:           "ingress SRC-VC buffer underflow",
		Description:    defaultAlwaysFatalErr.Description,
		PotentialFatal: defaultAlwaysFatalErr.PotentialFatal,
		AlwaysFatal:    defaultAlwaysFatalErr.AlwaysFatal,
		Impact:         defaultAlwaysFatalErr.Impact,
		Recovery:       defaultAlwaysFatalErr.Recovery,
		OtherImpact:    defaultAlwaysFatalErr.OtherImpact,
	},
	23012: {
		ID:             23012,
		Name:           "egress DST-VC credit overflow",
		Description:    defaultAlwaysFatalErr.Description,
		PotentialFatal: defaultAlwaysFatalErr.PotentialFatal,
		AlwaysFatal:    defaultAlwaysFatalErr.AlwaysFatal,
		Impact:         defaultAlwaysFatalErr.Impact,
		Recovery:       defaultAlwaysFatalErr.Recovery,
		OtherImpact:    defaultAlwaysFatalErr.OtherImpact,
	},
	23013: {
		ID:             23013,
		Name:           "egress DST-VC credit underflow",
		Description:    defaultAlwaysFatalErr.Description,
		PotentialFatal: defaultAlwaysFatalErr.PotentialFatal,
		AlwaysFatal:    defaultAlwaysFatalErr.AlwaysFatal,
		Impact:         defaultAlwaysFatalErr.Impact,
		Recovery:       defaultAlwaysFatalErr.Recovery,
		OtherImpact:    defaultAlwaysFatalErr.OtherImpact,
	},
	23014: {
		ID:             23014,
		Name:           "ingress packet burst error",
		Description:    defaultAlwaysFatalErr.Description,
		PotentialFatal: defaultAlwaysFatalErr.PotentialFatal,
		AlwaysFatal:    defaultAlwaysFatalErr.AlwaysFatal,
		Impact:         defaultAlwaysFatalErr.Impact,
		Recovery:       defaultAlwaysFatalErr.Recovery,
		OtherImpact:    defaultAlwaysFatalErr.OtherImpact,
	},
	23015: {
		ID:             23015,
		Name:           "ingress packet sticky error",
		Description:    defaultAlwaysFatalErr.Description,
		PotentialFatal: defaultAlwaysFatalErr.PotentialFatal,
		AlwaysFatal:    defaultAlwaysFatalErr.AlwaysFatal,
		Impact:         defaultAlwaysFatalErr.Impact,
		Recovery:       defaultAlwaysFatalErr.Recovery,
		OtherImpact:    defaultAlwaysFatalErr.OtherImpact,
	},
	23016: {
		ID:             23016,
		Name:           "possible bubbles at ingress",
		Description:    defaultAlwaysFatalErr.Description,
		PotentialFatal: defaultAlwaysFatalErr.PotentialFatal,
		AlwaysFatal:    defaultAlwaysFatalErr.AlwaysFatal,
		Impact:         defaultAlwaysFatalErr.Impact,
		Recovery:       defaultAlwaysFatalErr.Recovery,
		OtherImpact:    defaultAlwaysFatalErr.OtherImpact,
	},
	23017: {
		ID:             23017,
		Name:           "ingress credit parity error",
		Description:    defaultAlwaysFatalErr.Description,
		PotentialFatal: defaultAlwaysFatalErr.PotentialFatal,
		AlwaysFatal:    defaultAlwaysFatalErr.AlwaysFatal,
		Impact:         defaultAlwaysFatalErr.Impact,
		Recovery:       defaultAlwaysFatalErr.Recovery,
		OtherImpact:    defaultAlwaysFatalErr.OtherImpact,
	},

	// D.7 Other Notable NVSwitch SXid Error
	10001: {
		ID:             10001,
		Name:           "Host_priv_error",
		Description:    "The errors are not fatal to the fabric/system, but they might be followed by other fatal events.",
		PotentialFatal: true,
		AlwaysFatal:    false,
		Impact:         "",
		Recovery:       "",
		OtherImpact:    "",
	},
	10002: {
		ID:             10002,
		Name:           "Host_priv_timeout",
		Description:    "The errors are not fatal to the fabric/system, but they might be followed by other fatal events.",
		PotentialFatal: true,
		AlwaysFatal:    false,
		Impact:         "",
		Recovery:       "",
		OtherImpact:    "",
	},
	10003: {
		ID:             10003,
		Name:           "Host_unhandled_interrupt",
		Description:    "This SXid error is never expected to occur.",
		PotentialFatal: true,
		AlwaysFatal:    true,
		Impact:         "If it occurs, it is fatal to the fabric/system.",
		Recovery:       "To recover, it will require a reset to all GPUs and NVSwitches (refer to section D.9).",
		OtherImpact:    "If the error is observed on a Trunk port, the partitions that are using NVSwitch trunk ports will be affected.",
	},
	10004: {
		ID:             10004,
		Name:           "Host_thermal_event_start",
		Description:    "Related to thermal events, which are not directly fatal to the fabric/system, but they indicate that system cooling might be insufficient.",
		PotentialFatal: true,
		AlwaysFatal:    false,
		Impact:         "This error might force the specified NVSwitch Links to enter power saving mode (Single Lane Mode) and impact over the NVLink throughput.",
		Recovery:       "Ensure that the system cooling is sufficient.",
		OtherImpact:    "",
	},
	10005: {
		ID:             10005,
		Name:           "Host_thermal_event_end",
		Description:    "Related to thermal events, which are not directly fatal to the fabric/system, but they indicate that system cooling might be insufficient.",
		PotentialFatal: true,
		AlwaysFatal:    false,
		Impact:         "",
		Recovery:       "Ensure that the system cooling is sufficient.",
		OtherImpact:    "",
	},
}

// D.5 Fatal NVSwitch SXid Errors
var defaultPotentialFatalErr = Detail{
	Description:    "The hypervisor must track these SXid source ports (NVLink) to determine whether the error occurred on an NVSwitch trunk port or NVSwitch access port. The fatal SXid will be propagated to the GPU as Xid 74 when applicable.",
	PotentialFatal: true,
	AlwaysFatal:    false,
	Impact: `If the error occurred on an NVSwitch access port, the impact will be limited to the corresponding guest VM. To recover, shut down the guest VM.

If the errors occurred on an NVSwitch trunk port, to reset the trunk ports and recover, shut down the guest VM partitions that are crossing the trunk port. The partitions can be recreated. Currently, the partitions that are using NVSwitch trunk ports are the 16x GPU partition and the 8x GPU partitions with four GPUs per baseboard.
`,
	Recovery:    "Restart the guest VM.",
	OtherImpact: "",
}

// D.6 Always Fatal NVSwitch SXid Errors
var defaultAlwaysFatalErr = Detail{
	Description: `Always fatal to the entire fabric/system. After an always fatal SXid error has occurred, the guest VM partitions need to be shut down and one of the following tasks must occur:

1. The host needs to be restarted.
2. After the NVSwitches and GPUs are SBRed, restart the Service VM restart.

`,
	PotentialFatal: true,
	AlwaysFatal:    true,
	Impact:         `Always fatal to the entire fabric/system.`,
	Recovery:       "Restart the guest VM.",
	OtherImpact:    "",
}