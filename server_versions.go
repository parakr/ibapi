package ibapi

// Version is the server version
type Version = int

const (
	// MIN_SERVER_VER_REAL_TIME_BARS              Version = 34
	// MIN_SERVER_VER_SCALE_ORDERS                Version = 35
	// MIN_SERVER_VER_SNAPSHOT_MKT_DATA           Version = 35
	// MIN_SERVER_VER_SSHORT_COMBO_LEGS           Version = 35
	// MIN_SERVER_VER_WHAT_IF_ORDERS              Version = 36
	// MIN_SERVER_VER_CONTRACT_CONID              Version = 37
	MIN_SERVER_VER_PTA_ORDERS                         Version = 39
	MIN_SERVER_VER_FUNDAMENTAL_DATA                   Version = 40
	MIN_SERVER_VER_DELTA_NEUTRAL                      Version = 40
	MIN_SERVER_VER_CONTRACT_DATA_CHAIN                Version = 40
	MIN_SERVER_VER_SCALE_ORDERS2                      Version = 40
	MIN_SERVER_VER_ALGO_ORDERS                        Version = 41
	MIN_SERVER_VER_EXECUTION_DATA_CHAIN               Version = 42
	MIN_SERVER_VER_NOT_HELD                           Version = 44
	MIN_SERVER_VER_SEC_ID_TYPE                        Version = 45
	MIN_SERVER_VER_PLACE_ORDER_CONID                  Version = 46
	MIN_SERVER_VER_REQ_MKT_DATA_CONID                 Version = 47
	MIN_SERVER_VER_REQ_CALC_IMPLIED_VOLAT             Version = 49
	MIN_SERVER_VER_REQ_CALC_OPTION_PRICE              Version = 50
	MIN_SERVER_VER_CANCEL_CALC_IMPLIED_VOLAT          Version = 50
	MIN_SERVER_VER_CANCEL_CALC_OPTION_PRICE           Version = 50
	MIN_SERVER_VER_SSHORTX_OLD                        Version = 51
	MIN_SERVER_VER_SSHORTX                            Version = 52
	MIN_SERVER_VER_REQ_GLOBAL_CANCEL                  Version = 53
	MIN_SERVER_VER_HEDGE_ORDERS                       Version = 54
	MIN_SERVER_VER_REQ_MARKET_DATA_TYPE               Version = 55
	MIN_SERVER_VER_OPT_OUT_SMART_ROUTING              Version = 56
	MIN_SERVER_VER_SMART_COMBO_ROUTING_PARAMS         Version = 57
	MIN_SERVER_VER_DELTA_NEUTRAL_CONID                Version = 58
	MIN_SERVER_VER_SCALE_ORDERS3                      Version = 60
	MIN_SERVER_VER_ORDER_COMBO_LEGS_PRICE             Version = 61
	MIN_SERVER_VER_TRAILING_PERCENT                   Version = 62
	MIN_SERVER_VER_DELTA_NEUTRAL_OPEN_CLOSE           Version = 66
	MIN_SERVER_VER_POSITIONS                          Version = 67
	MIN_SERVER_VER_ACCOUNT_SUMMARY                    Version = 67
	MIN_SERVER_VER_TRADING_CLASS                      Version = 68
	MIN_SERVER_VER_SCALE_TABLE                        Version = 69
	MIN_SERVER_VER_LINKING                            Version = 70
	MIN_SERVER_VER_ALGO_ID                            Version = 71
	MIN_SERVER_VER_OPTIONAL_CAPABILITIES              Version = 72
	MIN_SERVER_VER_ORDER_SOLICITED                    Version = 73
	MIN_SERVER_VER_LINKING_AUTH                       Version = 74
	MIN_SERVER_VER_PRIMARYEXCH                        Version = 75
	MIN_SERVER_VER_RANDOMIZE_SIZE_AND_PRICE           Version = 76
	MIN_SERVER_VER_FRACTIONAL_POSITIONS               Version = 101
	MIN_SERVER_VER_PEGGED_TO_BENCHMARK                Version = 102
	MIN_SERVER_VER_MODELS_SUPPORT                     Version = 103
	MIN_SERVER_VER_SEC_DEF_OPT_PARAMS_REQ             Version = 104
	MIN_SERVER_VER_EXT_OPERATOR                       Version = 105
	MIN_SERVER_VER_SOFT_DOLLAR_TIER                   Version = 106
	MIN_SERVER_VER_REQ_FAMILY_CODES                   Version = 107
	MIN_SERVER_VER_REQ_MATCHING_SYMBOLS               Version = 108
	MIN_SERVER_VER_PAST_LIMIT                         Version = 109
	MIN_SERVER_VER_MD_SIZE_MULTIPLIER                 Version = 110
	MIN_SERVER_VER_CASH_QTY                           Version = 111
	MIN_SERVER_VER_REQ_MKT_DEPTH_EXCHANGES            Version = 112
	MIN_SERVER_VER_TICK_NEWS                          Version = 113
	MIN_SERVER_VER_REQ_SMART_COMPONENTS               Version = 114
	MIN_SERVER_VER_REQ_NEWS_PROVIDERS                 Version = 115
	MIN_SERVER_VER_REQ_NEWS_ARTICLE                   Version = 116
	MIN_SERVER_VER_REQ_HISTORICAL_NEWS                Version = 117
	MIN_SERVER_VER_REQ_HEAD_TIMESTAMP                 Version = 118
	MIN_SERVER_VER_REQ_HISTOGRAM                      Version = 119
	MIN_SERVER_VER_SERVICE_DATA_TYPE                  Version = 120
	MIN_SERVER_VER_AGG_GROUP                          Version = 121
	MIN_SERVER_VER_UNDERLYING_INFO                    Version = 122
	MIN_SERVER_VER_CANCEL_HEADTIMESTAMP               Version = 123
	MIN_SERVER_VER_SYNT_REALTIME_BARS                 Version = 124
	MIN_SERVER_VER_CFD_REROUTE                        Version = 125
	MIN_SERVER_VER_MARKET_RULES                       Version = 126
	MIN_SERVER_VER_PNL                                Version = 127
	MIN_SERVER_VER_NEWS_QUERY_ORIGINS                 Version = 128
	MIN_SERVER_VER_UNREALIZED_PNL                     Version = 129
	MIN_SERVER_VER_HISTORICAL_TICKS                   Version = 130
	MIN_SERVER_VER_MARKET_CAP_PRICE                   Version = 131
	MIN_SERVER_VER_PRE_OPEN_BID_ASK                   Version = 132
	MIN_SERVER_VER_REAL_EXPIRATION_DATE               Version = 134
	MIN_SERVER_VER_REALIZED_PNL                       Version = 135
	MIN_SERVER_VER_LAST_LIQUIDITY                     Version = 136
	MIN_SERVER_VER_TICK_BY_TICK                       Version = 137
	MIN_SERVER_VER_DECISION_MAKER                     Version = 138
	MIN_SERVER_VER_MIFID_EXECUTION                    Version = 139
	MIN_SERVER_VER_TICK_BY_TICK_IGNORE_SIZE           Version = 140
	MIN_SERVER_VER_AUTO_PRICE_FOR_HEDGE               Version = 141
	MIN_SERVER_VER_WHAT_IF_EXT_FIELDS                 Version = 142
	MIN_SERVER_VER_SCANNER_GENERIC_OPTS               Version = 143
	MIN_SERVER_VER_API_BIND_ORDER                     Version = 144
	MIN_SERVER_VER_ORDER_CONTAINER                    Version = 145
	MIN_SERVER_VER_SMART_DEPTH                        Version = 146
	MIN_SERVER_VER_REMOVE_NULL_ALL_CASTING            Version = 147
	MIN_SERVER_VER_D_PEG_ORDERS                       Version = 148
	MIN_SERVER_VER_MKT_DEPTH_PRIM_EXCHANGE            Version = 149
	MIN_SERVER_VER_COMPLETED_ORDERS                   Version = 150
	MIN_SERVER_VER_PRICE_MGMT_ALGO                    Version = 151
	MIN_SERVER_VER_STOCK_TYPE                         Version = 152
	MIN_SERVER_VER_ENCODE_MSG_ASCII7                  Version = 153
	MIN_SERVER_VER_SEND_ALL_FAMILY_CODES              Version = 154
	MIN_SERVER_VER_NO_DEFAULT_OPEN_CLOSE              Version = 155
	MIN_SERVER_VER_PRICE_BASED_VOLATILITY             Version = 156
	MIN_SERVER_VER_REPLACE_FA_END                     Version = 157
	MIN_SERVER_VER_DURATION                           Version = 158
	MIN_SERVER_VER_MARKET_DATA_IN_SHARES              Version = 159
	MIN_SERVER_VER_POST_TO_ATS                        Version = 160
	MIN_SERVER_VER_WSHE_CALENDAR                      Version = 161
	MIN_SERVER_VER_AUTO_CANCEL_PARENT                 Version = 162
	MIN_SERVER_VER_FRACTIONAL_SIZE_SUPPORT            Version = 163
	MIN_SERVER_VER_SIZE_RULES                         Version = 164
	MIN_SERVER_VER_HISTORICAL_SCHEDULE                Version = 165
	MIN_SERVER_VER_ADVANCED_ORDER_REJECT              Version = 166
	MIN_SERVER_VER_USER_INFO                          Version = 167
	MIN_SERVER_VER_CRYPTO_AGGREGATED_TRADES           Version = 168
	MIN_SERVER_VER_PEGBEST_PEGMID_OFFSETS             Version = 170
	MIN_SERVER_VER_WSH_EVENT_DATA_FILTERS             Version = 171
	MIN_SERVER_VER_IPO_PRICES                         Version = 172
	MIN_SERVER_VER_WSH_EVENT_DATA_FILTERS_DATE        Version = 173
	MIN_SERVER_VER_INSTRUMENT_TIMEZONE                Version = 174
	MIN_SERVER_VER_HMDS_MARKET_DATA_IN_SHARES         Version = 175
	MIN_SERVER_VER_MANUAL_ORDER_TIME                  Version = 169
	MIN_SERVER_VER_BOND_ISSUERID                      Version = 176
	MIN_SERVER_VER_FA_PROFILE_DESUPPORT               Version = 177
	MIN_SERVER_VER_PENDING_PRICE_REVISION             Version = 178
	MIN_SERVER_VER_FUND_DATA_FIELDS                   Version = 179
	MIN_SERVER_VER_MANUAL_ORDER_TIME_EXERCISE_OPTIONS Version = 180
	MIN_SERVER_VER_OPEN_ORDER_AD_STRATEGY             Version = 181
	MIN_SERVER_VER_LAST_TRADE_DATE                    Version = 182
	MIN_SERVER_VER_CUSTOMER_ACCOUNT                   Version = 183
	MIN_SERVER_VER_PROFESSIONAL_CUSTOMER              Version = 184
	MIN_SERVER_VER_BOND_ACCRUED_INTEREST              Version = 185
	MIN_SERVER_VER_INELIGIBILITY_REASONS              Version = 186
	MIN_SERVER_VER_RFQ_FIELDS                         Version = 187
	MIN_SERVER_VER_BOND_TRADING_HOURS                 Version = 188
	MIN_SERVER_VER_INCLUDE_OVERNIGHT                  Version = 189
	MIN_SERVER_VER_UNDO_RFQ_FIELDS                    Version = 190
	MIN_SERVER_VER_PERM_ID_AS_LONG                    Version = 191
	MIN_SERVER_VER_CME_TAGGING_FIELDS                 Version = 192
	MIN_SERVER_VER_CME_TAGGING_FIELDS_IN_OPEN_ORDER   Version = 193

	// 100+ messaging
	// 100 = enhanced handshake, msg length prefixes

	MIN_CLIENT_VER = 100
	MAX_CLIENT_VER = MIN_SERVER_VER_CME_TAGGING_FIELDS_IN_OPEN_ORDER
)
