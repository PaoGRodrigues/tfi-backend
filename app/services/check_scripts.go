package services

type Check struct {
	Subdir    string
	ScriptKey string
}

var Checks = []Check{

	{
		Subdir:    "host",
		ScriptKey: "domain_names_contacts",
	},
	{
		Subdir:    "host",
		ScriptKey: "ntp_contacts",
	},
	{
		Subdir:    "host",
		ScriptKey: "syn_scan",
	},
	{
		Subdir:    "host",
		ScriptKey: "fin_scan",
	},
	{
		Subdir:    "host",
		ScriptKey: "dangerous_host",
	},
	{
		Subdir:    "host",
		ScriptKey: "",
	},
	{
		Subdir:    "host",
		ScriptKey: "",
	},
	{
		Subdir:    "host",
		ScriptKey: "",
	},
	{
		Subdir:    "host",
		ScriptKey: "scan_detection",
	},
	{
		Subdir:    "host",
		ScriptKey: "syn_flood",
	},
	{
		Subdir:    "host",
		ScriptKey: "dns_contacts",
	},
	{
		Subdir:    "host",
		ScriptKey: "flow_flood",
	},
	{
		Subdir:    "host",
		ScriptKey: "rst_scan",
	},
	{
		Subdir:    "host",
		ScriptKey: "smtp_contacts",
	},
	{
		Subdir:    "host",
		ScriptKey: "countries_contacts",
	},
	{
		Subdir:    "host",
		ScriptKey: "score_threshold",
	},
	{
		Subdir:    "host",
		ScriptKey: "icmp_flood",
	},
	{
		Subdir:    "network",
		ScriptKey: "ip_reassignment",
	},
	{
		Subdir:    "network",
		ScriptKey: "syn_flood_victim",
	},
	{
		Subdir:    "network",
		ScriptKey: "flow_flood_victim",
	},
	{
		Subdir:    "network",
		ScriptKey: "broadcast_domain_too_large",
	},
	{
		Subdir:    "network",
		ScriptKey: "network_discovery",
	},
	{
		Subdir:    "network",
		ScriptKey: "syn_scan_victim",
	},
	{
		Subdir:    "flow",
		ScriptKey: "blacklisted_client_contact",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_suspicious_entropy",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_http_suspicious_url",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_http_obsolete_server",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_binary_data_transfer",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_dns_large_packet",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_url_possible_rce_injection",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_http_suspicious_content",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_minor_issues",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_url_possible_sql_injection",
	},
	{
		Subdir:    "flow",
		ScriptKey: "device_protocol_not_allowed",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_url_possible_xss",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_smb_insecure_version",
	},
	{
		Subdir:    "flow",
		ScriptKey: "iec_unexpected_type_id",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_tls_alpn_sni_mismatch",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_invalid_characters",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_http_crawler_bot",
	},
	{
		Subdir:    "flow",
		ScriptKey: "remote_to_local_insecure_flow",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_periodic_flow",
	},
	{
		Subdir:    "flow",
		ScriptKey: "web_mining",
	},
	{
		Subdir:    "flow",
		ScriptKey: "unexpected_smtp",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_suspicious_dga_domain",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_desktop_or_file_sharing_session",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_tcp_issues",
	},
	{
		Subdir:    "flow",
		ScriptKey: "remote_to_remote",
	},
	{
		Subdir:    "flow",
		ScriptKey: "",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_http_suspicious_header",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_tls_suspicious_extension",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_malicious_sha1_certificate",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_dns_fragmented",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_tls_certificate_about_to_expire",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_numeric_ip_host",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_ssh_obsolete_server",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_tls_suspicious_esni_usage",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_malicious_ja3",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_http_suspicious_user_agent",
	},
	{
		Subdir:    "flow",
		ScriptKey: "unexpected_dns",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_ssh_obsolete_client",
	},
	{
		Subdir:    "flow",
		ScriptKey: "vlan_bidirectional_traffic",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_tls_fatal_alert",
	},
	{
		Subdir:    "flow",
		ScriptKey: "external_alert_check",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_risky_domain",
	},
	{
		Subdir:    "flow",
		ScriptKey: "iec_invalid_command_transition",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_anonymous_subscriber",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_tls_not_carrying_https",
	},
	{
		Subdir:    "flow",
		ScriptKey: "not_purged",
	},
	{
		Subdir:    "flow",
		ScriptKey: "custom_lua_script",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_unsafe_protocol",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_clear_text_credentials",
	},
	{
		Subdir:    "flow",
		ScriptKey: "blacklisted",
	},
	{
		Subdir:    "flow",
		ScriptKey: "broadcast_non_udp_traffic",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_possible_exploit",
	},
	{
		Subdir:    "flow",
		ScriptKey: "binary_application_transfer",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_probing_attempt",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_tls_uncommon_alpn",
	},
	{
		Subdir:    "flow",
		ScriptKey: "zero_tcp_window",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_malware_host_contacted",
	},
	{
		Subdir:    "flow",
		ScriptKey: "country_check",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_dns_suspicious_traffic",
	},
	{
		Subdir:    "flow",
		ScriptKey: "iec_invalid_transition",
	},
	{
		Subdir:    "flow",
		ScriptKey: "unexpected_dhcp",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_risky_asn",
	},
	{
		Subdir:    "flow",
		ScriptKey: "low_goodput",
	},
	{
		Subdir:    "flow",
		ScriptKey: "unexpected_ntp",
	},
	{
		Subdir:    "flow",
		ScriptKey: "remote_access",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_punicody_idn",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_malformed_packet",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_error_code_detected",
	},
	{
		Subdir:    "flow",
		ScriptKey: "tcp_no_data_exchanged",
	},
	{
		Subdir:    "flow",
		ScriptKey: "tcp_flow_reset",
	},
	{
		Subdir:    "flow",
		ScriptKey: "rare_destination",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_tls_missing_sni",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_unidirectional_traffic",
	},
	{
		Subdir:    "flow",
		ScriptKey: "blacklisted_server_contact",
	},
	{
		Subdir:    "flow",
		ScriptKey: "known_proto_on_non_std_port",
	},
	{
		Subdir:    "flow",
		ScriptKey: "ndpi_fully_encrypted",
	},
}
