package main

// GENERATED BY gen.go

var methodParamMap = map[string][]string{
	"encodeXML":    []string{"v"},
	"decodeXML":    []string{"buf", "takeFirstEl"},
	"buildRequest": []string{"urlstr", "v"},
	"doReq":        []string{"path", "v", "takeFirstEl"},
	"doReqString":  []string{"path", "v", "elName"},
	"doReqCheckOK": []string{"path", "v"},
	"Do":           []string{"path", "v"},
	"NewSessionAndTokenID": []string{},
	"SetSessionAndTokenID": []string{"sessionID", "tokenID"},
	"GlobalConfig":         []string{},
	"NetworkTypes":         []string{},
	"PCAssistantConfig":    []string{},
	"DeviceConfig":         []string{},
	"WebUIConfig":          []string{},
	"SmsConfig":            []string{},
	"WlanConfig":           []string{},
	"DhcpConfig":           []string{},
	"CradleStatusInfo":     []string{},
	"CradleMACSet":         []string{"addr"},
	"CradleMAC":            []string{},
	"AutorunVersion":       []string{},
	"DeviceBasicInfo":      []string{},
	"PublicKey":            []string{},
	"DeviceControl":        []string{"code"},
	"DeviceReboot":         []string{},
	"DeviceReset":          []string{},
	"DeviceBackup":         []string{},
	"DeviceShutdown":       []string{},
	"DeviceFeatures":       []string{},
	"DeviceInfo":           []string{},
	"DeviceMode":           []string{"mode"},
	"FastbootFeatures":     []string{},
	"PowerFeatures":        []string{},
	"TetheringFeatures":    []string{},
	"SignalInfo":           []string{},
	"ConnectionInfo":       []string{},
	"GlobalFeatures":       []string{},
	"Language":             []string{},
	"LanguageSet":          []string{"lang"},
	"NotificationInfo":     []string{},
	"SimInfo":              []string{},
	"StatusInfo":           []string{},
	"TrafficInfo":          []string{},
	"TrafficClear":         []string{},
	"MonthInfo":            []string{},
	"WlanMonthInfo":        []string{},
	"NetworkInfo":          []string{},
	"WifiFeatures":         []string{},
	"ModeList":             []string{},
	"ModeInfo":             []string{},
	"ModeNetworkInfo":      []string{},
	"ModeSet":              []string{"netMode", "netBand", "lteBand"},
	"PinInfo":              []string{},
	"doReqPin":             []string{"pt", "cur", "new", "puk"},
	"PinEnter":             []string{"pin"},
	"PinActivate":          []string{"pin"},
	"PinDeactivate":        []string{"pin"},
	"PinChange":            []string{"pin", "new"},
	"PinEnterPuk":          []string{"puk", "new"},
	"PinSaveInfo":          []string{},
	"PinSimlockInfo":       []string{},
	"Connect":              []string{},
	"Disconnect":           []string{},
	"ProfileInfo":          []string{},
	"SmsFeatures":          []string{},
	"SmsList":              []string{"boxType", "page", "count", "sortByName", "ascending", "unreadPreferred"},
	"SmsCount":             []string{},
	"SmsSend":              []string{"msg", "to"},
	"SmsSendStatus":        []string{},
	"SmsReadSet":           []string{"id"},
	"SmsDelete":            []string{"id"},
	"UssdStatus":           []string{},
	"UssdCode":             []string{"code"},
	"UssdContent":          []string{},
	"UssdRelease":          []string{},
	"DdnsList":             []string{},
	"LogPath":              []string{},
	"LogInfo":              []string{},
	"PhonebookGroupList":   []string{"page", "count", "sortByName", "ascending"},
	"PhonebookCount":       []string{},
	"PhonebookImport":      []string{"group"},
	"PhonebookDelete":      []string{"id"},
	"PhonebookList":        []string{"group", "page", "count", "sim", "sortByName", "ascending", "keyword"},
	"PhonebookCreate":      []string{"group", "name", "phone", "sim"},
	"FirewallFeatures":     []string{},
}

var methodCommentMap = map[string]string{
	"encodeXML":    "encodeXML encodes a map to standard XML values.",
	"decodeXML":    "decodeXML decodes buf into its simple xml values.",
	"buildRequest": "buildRequest creates a request for use with the Client.",
	"doReq":        "doReq sends a request to the server with the provided path. If data is nil,then GET will be used as the HTTP method, otherwise POST will be used.",
	"doReqString":  "doReqString wraps a request operation, returning the data of the specifiedchild node named elName as a string.",
	"doReqCheckOK": "doReqCheckOK wraps a request operation (ie, connect, disconnect, etc),checking success via the presence of 'OK' in the XML <response/>.",
	"Do":           "Do sends a request to the server with the provided path. If data is nil,then GET will be used as the HTTP method, otherwise POST will be used.",
	"NewSessionAndTokenID": "NewSessionAndTokenID starts a session with the server, and returns thesession and token.",
	"SetSessionAndTokenID": "SetSessionIDAndTokenID sets the sessionID for the Client.",
	"GlobalConfig":         "GlobalConfig retrieves global Hilink configuration.",
	"NetworkTypes":         "NetworkTypes retrieves available network types.",
	"PCAssistantConfig":    "PCAssistantConfig retrieves PC Assistant configuration.",
	"DeviceConfig":         "DeviceConfig retrieves device configuration.",
	"WebUIConfig":          "WebUIConfig retrieves WebUI configuration.",
	"SmsConfig":            "SmsConfig retrieves device SMS configuration.",
	"WlanConfig":           "WlanConfig retrieves basic WLAN settings.",
	"DhcpConfig":           "DhcpConfig retrieves DHCP configuration.",
	"CradleStatusInfo":     "CradleStatusInfo retrieves cradle status information.",
	"CradleMACSet":         "CradleMACSet sets the MAC address for the cradle.",
	"CradleMAC":            "CradleMAC retrieves cradle MAC address.",
	"AutorunVersion":       "AutorunVersion retrieves autorun version.",
	"DeviceBasicInfo":      "DeviceBasicInfo retrieves basic device information.",
	"PublicKey":            "PublicKey retrieves hilink public key.",
	"DeviceControl":        "DeviceControl sends a control code to the device.",
	"DeviceReboot":         "DeviceReboot restarts the Hilink device.",
	"DeviceReset":          "DeviceReset resets the device configuration.",
	"DeviceBackup":         "DeviceBackup backups device configuration and retrieves backed upconfiguration data as a base64 encoded string.",
	"DeviceShutdown":       "DeviceShutdown shuts down the device.",
	"DeviceFeatures":       "DeviceFeatures retrieves device feature information.",
	"DeviceInfo":           "DeviceInfo retrieves Hilink device information.",
	"DeviceMode":           "DeviceMode sets the device mode (0-project, 1-debug).",
	"FastbootFeatures":     "FastbootFeatures retrieves fastboot feature information.",
	"PowerFeatures":        "PowerFeatures retrieves power feature information.",
	"TetheringFeatures":    "TetheringFeatures retrieves USB tethering feature information.",
	"SignalInfo":           "SignalInfo retrieves signal information.",
	"ConnectionInfo":       "ConnectionInfo retrieves connection (dialup) information.",
	"GlobalFeatures":       "GlobalFeatures retrieves global feature information.",
	"Language":             "Language retrieves current language.",
	"LanguageSet":          "LanguageSet sets the language.",
	"NotificationInfo":     "NotificationInfo retrieves notification information.",
	"SimInfo":              "SimInfo retrieves SIM card information.",
	"StatusInfo":           "StatusInfo retrieves Hilink device and connection status information.",
	"TrafficInfo":          "TrafficInfo retrieves traffic statistic information.",
	"TrafficClear":         "TrafficClear clears the current traffic statistics.",
	"MonthInfo":            "MonthInfo retrieves the month download statistic information.",
	"WlanMonthInfo":        "WlanMonthInfo retrieves the WLAN month download statistic information.",
	"NetworkInfo":          "NetworkInfo retrieves network provider information.",
	"WifiFeatures":         "WifiFeatures retrieves wifi feature information.",
	"ModeList":             "ModeList retrieves available network modes.",
	"ModeInfo":             "ModeInfo retrieves network mode settings information.",
	"ModeNetworkInfo":      "ModeNetworkInfo retrieves current network mode information.",
	"ModeSet":              "ModeSet sets the network mode.",
	"PinInfo":              "PinInfo retrieves SIM PIN status information.",
	"doReqPin":             "doReqPin wraps a SIM PIN manipulation request.",
	"PinEnter":             "PinEnter enters a SIM PIN.",
	"PinActivate":          "PinActivate activates a SIM PIN.",
	"PinDeactivate":        "PinDeactivate deactivates a SIM PIN.",
	"PinChange":            "PinChange changes a SIM PIN.",
	"PinEnterPuk":          "PinEnterPuk enters a SIM PIN puk.",
	"PinSaveInfo":          "PinSaveInfo retrieves SIM PIN save information.",
	"PinSimlockInfo":       "PinSimlockInfo retrieves SIM lock information.",
	"Connect":              "Connect connects the Hilink device to the network provider.",
	"Disconnect":           "Disconnect disconnects the Hilink device from the network provider.",
	"ProfileInfo":          "ProfileInfo retrieves profile information (ie, APN).",
	"SmsFeatures":          "SmsFeatures retrieves SMS feature information.",
	"SmsList":              "SmsList retrieves list of SMS in an inbox.",
	"SmsCount":             "SmsCount retrieves count of SMS per inbox type.",
	"SmsSend":              "SmsSend sends an SMS.",
	"SmsSendStatus":        "SmsSendStatus retrieves SMS send status information.",
	"SmsReadSet":           "SmsReadSet sets the read status of a SMS.",
	"SmsDelete":            "SmsDelete deletes a specified SMS.",
	"UssdStatus":           "UssdStatus determines if the Hilink device is currently engaged in a USSDsession.",
	"UssdCode":             "UssdCode sends a USSD code to the Hilink device.",
	"UssdContent":          "UssdContent retrieves content buffer of the active USSD session.",
	"UssdRelease":          "UssdRelease releases the active USSD session.",
	"DdnsList":             "DdnsList retrieves list of DDNS providers.",
	"LogPath":              "LogPath retrieves device log path (URL).",
	"LogInfo":              "LogInfo retrieves current log setting information.",
	"PhonebookGroupList":   "PhonebookGroupList retrieves list of the phonebook groups.",
	"PhonebookCount":       "PhonebookCount retrieves count of phonebook entries per group.",
	"PhonebookImport":      "PhonebookImport imports SIM contacts into specified phonebook group.",
	"PhonebookDelete":      "PhonebookDelete deletes a specified phonebook entry.",
	"PhonebookList":        "PhonebookList retrieves list of phonebook entries from a specified group.",
	"PhonebookCreate":      "PhonebookCreate creates a new phonebook entry.",
	"FirewallFeatures":     "FirewallFeatures retrieves firewall security feature information.",
}
