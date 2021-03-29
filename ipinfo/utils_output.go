package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/ipinfo/go/v2/ipinfo"
	"github.com/jszwec/csvutil"
)

func outputJSON(d interface{}) error {
	jsonEnc := json.NewEncoder(os.Stdout)
	jsonEnc.SetIndent("", "  ")
	return jsonEnc.Encode(d)
}

func outputCSV(v interface{}) error {
	csvWriter := csv.NewWriter(os.Stdout)
	csvEnc := csvutil.NewEncoder(csvWriter)

	if err := csvEnc.Encode(v); err != nil {
		return err
	}
	csvWriter.Flush()

	return nil
}

func outputCSVBatchCore(core ipinfo.BatchCore) error {
	csvWriter := csv.NewWriter(os.Stdout)
	csvEnc := csvutil.NewEncoder(csvWriter)

	// print entries.
	for _, v := range core {
		if err := csvEnc.Encode(v); err != nil {
			return err
		}
		csvWriter.Flush()
	}

	return nil
}

func outputFriendlyCore(d *ipinfo.Core) {
	header := color.New(color.Bold, color.BgWhite, color.FgHiMagenta)

	header.Printf("                 CORE                 ")
	fmt.Println()
	fmt.Printf("IP              %s\n", d.IP.String())
	if d.Bogon {
		// exit early after printing bogon status.
		fmt.Printf("Bogon           %v\n", d.Bogon)
		return
	}
	fmt.Printf("Anycast         %v\n", d.Anycast)
	fmt.Printf("Hostname        %s\n", d.Hostname)
	fmt.Printf("City            %s\n", d.City)
	fmt.Printf("Region          %s\n", d.Region)
	fmt.Printf("Country         %s (%s)\n", d.CountryName, d.Country)
	fmt.Printf("Location        %s\n", d.Location)
	fmt.Printf("Organization    %s\n", d.Org)
	fmt.Printf("Postal          %s\n", d.Postal)
	fmt.Printf("Timezone        %s\n", d.Timezone)
	if d.ASN != nil {
		fmt.Println()
		header.Printf("                 ASN                  ")
		fmt.Println()
		fmt.Printf("ID              %s\n", d.ASN.ASN)
		fmt.Printf("Name            %s\n", d.ASN.Name)
		fmt.Printf("Domain          %s\n", d.ASN.Domain)
		fmt.Printf("Route           %s\n", d.ASN.Route)
		fmt.Printf("Type            %s\n", d.ASN.Type)
	}
	if d.Company != nil {
		fmt.Println()
		header.Printf("               COMPANY                ")
		fmt.Println()
		fmt.Printf("Name            %s\n", d.Company.Name)
		fmt.Printf("Domain          %s\n", d.Company.Domain)
		fmt.Printf("Type            %s\n", d.Company.Type)
	}
	if d.Carrier != nil {
		fmt.Println()
		header.Printf("               CARRIER                ")
		fmt.Println()
		fmt.Printf("Name            %s\n", d.Carrier.Name)
		fmt.Printf("MCC             %s\n", d.Carrier.MCC)
		fmt.Printf("MNC             %s\n", d.Carrier.MNC)
	}
	if d.Privacy != nil {
		fmt.Println()
		header.Printf("               PRIVACY                ")
		fmt.Println()
		fmt.Printf("VPN             %v\n", d.Privacy.VPN)
		fmt.Printf("Proxy           %v\n", d.Privacy.Proxy)
		fmt.Printf("Tor             %v\n", d.Privacy.Tor)
		fmt.Printf("Hosting         %v\n", d.Privacy.Hosting)
	}
	if d.Abuse != nil {
		fmt.Println()
		header.Printf("                ABUSE                 ")
		fmt.Println()
		fmt.Printf("Address         %s\n", d.Abuse.Address)
		fmt.Printf("Country         %s (%s)\n", d.Abuse.CountryName, d.Abuse.Country)
		fmt.Printf("Email           %s\n", d.Abuse.Email)
		fmt.Printf("Name            %s\n", d.Abuse.Name)
		fmt.Printf("Network         %s\n", d.Abuse.Network)
		fmt.Printf("Phone           %s\n", d.Abuse.Phone)
	}
	if d.Domains != nil && d.Domains.Total > 0 {
		fmt.Println()
		header.Printf("               DOMAINS                ")
		fmt.Println()
		fmt.Printf("Total           %v\n", d.Domains.Total)
		if len(d.Domains.Domains) > 0 {
			fmt.Printf("Examples     1: %s\n", d.Domains.Domains[0])
			if len(d.Domains.Domains) > 1 {
				for i, d := range d.Domains.Domains[1:] {
					fmt.Printf("             %v: %s\n", i+2, d)
				}
			}
		}
	}
}

func outputFieldBatchCore(
	core ipinfo.BatchCore,
	field string,
	header bool,
	fieldOnly bool,
) error {
	csvWriter := csv.NewWriter(os.Stdout)
	csvEnc := csvutil.NewEncoder(csvWriter)
	csvEnc.AutoHeader = false

	// TODO the dread of not having macros... we can simplify code length here
	// with reflection but until then this will have to do.
	switch field {
	case "ip":
		if header {
			fmt.Printf("ip\n")
		}

		for _, d := range core {
			fmt.Printf("%v\n", d.IP)
		}
	case "hostname":
		if header {
			if !fieldOnly {
				fmt.Printf("ip,")
			}
			fmt.Printf("hostname\n")
		}

		for _, d := range core {
			if !fieldOnly {
				fmt.Printf("%s,", d.IP)
			}
			fmt.Printf("%v\n", d.Hostname)
		}
	case "anycast":
		if header {
			if !fieldOnly {
				fmt.Printf("ip,")
			}
			fmt.Printf("anycast\n")
		}

		for _, d := range core {
			if !fieldOnly {
				fmt.Printf("%s,", d.IP)
			}
			fmt.Printf("%v\n", d.Anycast)
		}
	case "city":
		if header {
			if !fieldOnly {
				fmt.Printf("ip,")
			}
			fmt.Printf("city\n")
		}

		for _, d := range core {
			if !fieldOnly {
				fmt.Printf("%s,", d.IP)
			}
			fmt.Printf("%v\n", d.City)
		}
	case "region":
		if header {
			if !fieldOnly {
				fmt.Printf("ip,")
			}
			fmt.Printf("region\n")
		}

		for _, d := range core {
			if !fieldOnly {
				fmt.Printf("%s,", d.IP)
			}
			fmt.Printf("%v\n", d.Region)
		}
	case "country":
		if header {
			if !fieldOnly {
				fmt.Printf("ip,")
			}
			fmt.Printf("country\n")
		}

		for _, d := range core {
			if !fieldOnly {
				fmt.Printf("%s,", d.IP)
			}
			fmt.Printf("%v\n", d.Country)
		}
	case "country_name":
		if header {
			if !fieldOnly {
				fmt.Printf("ip,")
			}
			fmt.Printf("country_name\n")
		}

		for _, d := range core {
			if !fieldOnly {
				fmt.Printf("%s,", d.IP)
			}
			fmt.Printf("%v\n", d.CountryName)
		}
	case "loc":
		if header {
			if !fieldOnly {
				fmt.Printf("ip,")
			}
			fmt.Printf("loc\n")
		}

		for _, d := range core {
			if !fieldOnly {
				fmt.Printf("%s,", d.IP)
			}
			fmt.Printf("%v\n", d.Location)
		}
	case "org":
		if header {
			if !fieldOnly {
				fmt.Printf("ip,")
			}
			fmt.Printf("org\n")
		}

		for _, d := range core {
			if !fieldOnly {
				fmt.Printf("%s,", d.IP)
			}
			fmt.Printf("%v\n", d.Org)
		}
	case "postal":
		if header {
			if !fieldOnly {
				fmt.Printf("ip,")
			}
			fmt.Printf("postal\n")
		}

		for _, d := range core {
			if !fieldOnly {
				fmt.Printf("%s,", d.IP)
			}
			fmt.Printf("%v\n", d.Postal)
		}
	case "timezone":
		if header {
			if !fieldOnly {
				fmt.Printf("ip,")
			}
			fmt.Printf("timezone\n")
		}

		for _, d := range core {
			if !fieldOnly {
				fmt.Printf("%s,", d.IP)
			}
			fmt.Printf("%v\n", d.Timezone)
		}
	case "asn":
		if !fieldOnly {
			fmt.Printf("ip,")
		}
		if err := csvEnc.EncodeHeader(ipinfo.CoreASN{}); err != nil {
			return err
		}
		csvWriter.Flush()

		for _, d := range core {
			if d.ASN == nil {
				continue
			}

			if !fieldOnly {
				fmt.Printf("%s,", d.IP)
			}
			if err := csvEnc.Encode(d.ASN); err != nil {
				return err
			}
			csvWriter.Flush()
		}
	case "asn.id":
		if header {
			if !fieldOnly {
				fmt.Printf("ip,")
			}
			fmt.Printf("asn_id\n")
		}

		for _, d := range core {
			if d.ASN == nil {
				continue
			}

			if !fieldOnly {
				fmt.Printf("%s,", d.IP)
			}
			fmt.Printf("%v\n", d.ASN.ASN)
		}
	case "asn.name", "asn.asn":
		if header {
			if !fieldOnly {
				fmt.Printf("ip,")
			}
			fmt.Printf("asn_name\n")
		}

		for _, d := range core {
			if d.ASN == nil {
				continue
			}

			if !fieldOnly {
				fmt.Printf("%s,", d.IP)
			}
			fmt.Printf("%v\n", d.ASN.Name)
		}
	case "asn.domain":
		if header {
			if !fieldOnly {
				fmt.Printf("ip,")
			}
			fmt.Printf("asn_domain\n")
		}

		for _, d := range core {
			if d.ASN == nil {
				continue
			}

			if !fieldOnly {
				fmt.Printf("%s,", d.IP)
			}
			fmt.Printf("%v\n", d.ASN.Domain)
		}
	case "asn.route":
		if header {
			if !fieldOnly {
				fmt.Printf("ip,")
			}
			fmt.Printf("asn_route\n")
		}

		for _, d := range core {
			if d.ASN == nil {
				continue
			}

			if !fieldOnly {
				fmt.Printf("%s,", d.IP)
			}
			fmt.Printf("%v\n", d.ASN.Route)
		}
	case "asn.type":
		if header {
			if !fieldOnly {
				fmt.Printf("ip,")
			}
			fmt.Printf("asn_type\n")
		}

		for _, d := range core {
			if d.ASN == nil {
				continue
			}

			if !fieldOnly {
				fmt.Printf("%s,", d.IP)
			}
			fmt.Printf("%v\n", d.ASN.Type)
		}
	case "company":
		if !fieldOnly {
			fmt.Printf("ip,")
		}
		if err := csvEnc.EncodeHeader(ipinfo.CoreCompany{}); err != nil {
			return err
		}
		csvWriter.Flush()

		for _, d := range core {
			if d.Company == nil {
				continue
			}

			if !fieldOnly {
				fmt.Printf("%s,", d.IP)
			}
			if err := csvEnc.Encode(d.Company); err != nil {
				return err
			}
			csvWriter.Flush()
		}
	case "company.name":
		if header {
			if !fieldOnly {
				fmt.Printf("ip,")
			}
			fmt.Printf("company_name\n")
		}

		for _, d := range core {
			if d.Company == nil {
				continue
			}

			if !fieldOnly {
				fmt.Printf("%s,", d.IP)
			}
			fmt.Printf("%v\n", d.Company.Name)
		}
	case "company.domain":
		if header {
			if !fieldOnly {
				fmt.Printf("ip,")
			}
			fmt.Printf("company_domain\n")
		}

		for _, d := range core {
			if d.Company == nil {
				continue
			}

			if !fieldOnly {
				fmt.Printf("%s,", d.IP)
			}
			fmt.Printf("%v\n", d.Company.Domain)
		}
	case "company.type":
		if header {
			if !fieldOnly {
				fmt.Printf("ip,")
			}
			fmt.Printf("company_type\n")
		}

		for _, d := range core {
			if d.Company == nil {
				continue
			}

			if !fieldOnly {
				fmt.Printf("%s,", d.IP)
			}
			fmt.Printf("%v\n", d.Company.Type)
		}
	case "carrier":
		if !fieldOnly {
			fmt.Printf("ip,")
		}
		if err := csvEnc.EncodeHeader(ipinfo.CoreCarrier{}); err != nil {
			return err
		}
		csvWriter.Flush()

		for _, d := range core {
			if d.Carrier == nil {
				continue
			}

			if !fieldOnly {
				fmt.Printf("%s,", d.IP)
			}
			if err := csvEnc.Encode(d.Carrier); err != nil {
				return err
			}
			csvWriter.Flush()
		}
	case "carrier.name":
		if header {
			if !fieldOnly {
				fmt.Printf("ip,")
			}
			fmt.Printf("carrier_name\n")
		}

		for _, d := range core {
			if d.Carrier == nil {
				continue
			}

			if !fieldOnly {
				fmt.Printf("%s,", d.IP)
			}
			fmt.Printf("%v\n", d.Carrier.Name)
		}
	case "carrier.mcc":
		if header {
			if !fieldOnly {
				fmt.Printf("ip,")
			}
			fmt.Printf("carrier_mcc\n")
		}

		for _, d := range core {
			if d.Carrier == nil {
				continue
			}

			if !fieldOnly {
				fmt.Printf("%s,", d.IP)
			}
			fmt.Printf("%v\n", d.Carrier.MCC)
		}
	case "carrier.mnc":
		if header {
			if !fieldOnly {
				fmt.Printf("ip,")
			}
			fmt.Printf("carrier_mnc\n")
		}

		for _, d := range core {
			if d.Carrier == nil {
				continue
			}

			if !fieldOnly {
				fmt.Printf("%s,", d.IP)
			}
			fmt.Printf("%v\n", d.Carrier.MNC)
		}
	case "privacy":
		if !fieldOnly {
			fmt.Printf("ip,")
		}
		if err := csvEnc.EncodeHeader(ipinfo.CorePrivacy{}); err != nil {
			return err
		}
		csvWriter.Flush()

		for _, d := range core {
			if d.Privacy == nil {
				continue
			}

			if !fieldOnly {
				fmt.Printf("%s,", d.IP)
			}
			if err := csvEnc.Encode(d.Privacy); err != nil {
				return err
			}
			csvWriter.Flush()
		}
	case "privacy.vpn":
		if header {
			if !fieldOnly {
				fmt.Printf("ip,")
			}
			fmt.Printf("privacy_vpn\n")
		}

		for _, d := range core {
			if d.Privacy == nil {
				continue
			}

			if !fieldOnly {
				fmt.Printf("%s,", d.IP)
			}
			fmt.Printf("%v\n", d.Privacy.VPN)
		}
	case "privacy.proxy":
		if header {
			if !fieldOnly {
				fmt.Printf("ip,")
			}
			fmt.Printf("privacy_proxy\n")
		}

		for _, d := range core {
			if d.Privacy == nil {
				continue
			}

			if !fieldOnly {
				fmt.Printf("%s,", d.IP)
			}
			fmt.Printf("%v\n", d.Privacy.Proxy)
		}
	case "privacy.tor":
		if header {
			if !fieldOnly {
				fmt.Printf("ip,")
			}
			fmt.Printf("privacy_tor\n")
		}

		for _, d := range core {
			if d.Privacy == nil {
				continue
			}

			if !fieldOnly {
				fmt.Printf("%s,", d.IP)
			}
			fmt.Printf("%v\n", d.Privacy.Tor)
		}
	case "privacy.hosting":
		if header {
			if !fieldOnly {
				fmt.Printf("ip,")
			}
			fmt.Printf("privacy_hosting\n")
		}

		for _, d := range core {
			if d.Privacy == nil {
				continue
			}

			if !fieldOnly {
				fmt.Printf("%s,", d.IP)
			}
			fmt.Printf("%v\n", d.Privacy.Hosting)
		}
	case "abuse":
		if !fieldOnly {
			fmt.Printf("ip,")
		}
		if err := csvEnc.EncodeHeader(ipinfo.CoreAbuse{}); err != nil {
			return err
		}
		csvWriter.Flush()

		for _, d := range core {
			if d.Abuse == nil {
				continue
			}

			if !fieldOnly {
				fmt.Printf("%s,", d.IP)
			}
			if err := csvEnc.Encode(d.Abuse); err != nil {
				return err
			}
			csvWriter.Flush()
		}
	case "abuse.address":
		if header {
			if !fieldOnly {
				fmt.Printf("ip,")
			}
			fmt.Printf("abuse_address\n")
		}

		for _, d := range core {
			if d.Abuse == nil {
				continue
			}

			if !fieldOnly {
				fmt.Printf("%s,", d.IP)
			}
			fmt.Printf("%v\n", d.Abuse.Address)
		}
	case "abuse.country":
		if header {
			if !fieldOnly {
				fmt.Printf("ip,")
			}
			fmt.Printf("abuse_country\n")
		}

		for _, d := range core {
			if d.Abuse == nil {
				continue
			}

			if !fieldOnly {
				fmt.Printf("%s,", d.IP)
			}
			fmt.Printf("%s%v\"\n", d.Abuse.Country)
		}
	case "abuse.country_name":
		if header {
			if !fieldOnly {
				fmt.Printf("ip,")
			}
			fmt.Printf("abuse_country_name\n")
		}

		for _, d := range core {
			if d.Abuse == nil {
				continue
			}

			if !fieldOnly {
				fmt.Printf("%s,", d.IP)
			}
			fmt.Printf("%v\n", d.Abuse.CountryName)
		}
	case "abuse.email":
		if header {
			if !fieldOnly {
				fmt.Printf("ip,")
			}
			fmt.Printf("abuse_email\n")
		}

		for _, d := range core {
			if d.Abuse == nil {
				continue
			}

			if !fieldOnly {
				fmt.Printf("%s,", d.IP)
			}
			fmt.Printf("%v\n", d.Abuse.Email)
		}
	case "abuse.name":
		if header {
			if !fieldOnly {
				fmt.Printf("ip,")
			}
			fmt.Printf("abuse_name\n")
		}

		for _, d := range core {
			if d.Abuse == nil {
				continue
			}

			if !fieldOnly {
				fmt.Printf("%s,", d.IP)
			}
			fmt.Printf("%v\n", d.Abuse.Name)
		}
	case "abuse.network":
		if header {
			if !fieldOnly {
				fmt.Printf("ip,")
			}
			fmt.Printf("abuse_network\n")
		}

		for _, d := range core {
			if d.Abuse == nil {
				continue
			}

			if !fieldOnly {
				fmt.Printf("%s,", d.IP)
			}
			fmt.Printf("%v\n", d.Abuse.Network)
		}
	case "abuse.phone":
		if header {
			if !fieldOnly {
				fmt.Printf("ip,")
			}
			fmt.Printf("abuse_phone\n")
		}

		for _, d := range core {
			if d.Abuse == nil {
				continue
			}

			if !fieldOnly {
				fmt.Printf("%s,", d.IP)
			}
			fmt.Printf("%v\n", d.Abuse.Phone)
		}
	case "domains":
		if !fieldOnly {
			fmt.Printf("ip,")
		}
		if err := csvEnc.EncodeHeader(ipinfo.CoreDomains{}); err != nil {
			return err
		}
		csvWriter.Flush()

		for _, d := range core {
			if d.Domains == nil {
				continue
			}

			if !fieldOnly {
				fmt.Printf("%s,", d.IP)
			}
			if err := csvEnc.Encode(d.Domains); err != nil {
				return err
			}
			csvWriter.Flush()
		}
	case "domains.total":
		if header {
			if !fieldOnly {
				fmt.Printf("ip,")
			}
			fmt.Printf("domains_total\n")
		}

		for _, d := range core {
			if d.Domains == nil {
				continue
			}

			if !fieldOnly {
				fmt.Printf("%s,", d.IP)
			}
			fmt.Printf("%v\n", d.Domains.Total)
		}
	default:
		if header {
			if !fieldOnly {
				fmt.Printf("ip,")
			}
			fmt.Printf("%s\n", field)
		}
	}

	return nil
}

func outputFieldBatchASNDetails(
	asnDetails ipinfo.BatchASNDetails,
	field string,
	header bool,
	fieldOnly bool,
) error {
	csvWriter := csv.NewWriter(os.Stdout)
	csvEnc := csvutil.NewEncoder(csvWriter)
	csvEnc.AutoHeader = false

	switch field {
	case "name":
		if header {
			if !fieldOnly {
				fmt.Printf("asn,")
			}
			fmt.Printf("name\n")
		}

		for _, d := range asnDetails {
			if !fieldOnly {
				fmt.Printf("%s,", d.ASN)
			}
			fmt.Printf("%v\n", d.Name)
		}
	case "country":
		if header {
			if !fieldOnly {
				fmt.Printf("asn,")
			}
			fmt.Printf("country\n")
		}

		for _, d := range asnDetails {
			if !fieldOnly {
				fmt.Printf("%s,", d.ASN)
			}
			fmt.Printf("%v\n", d.Country)
		}
	case "country_name":
		if header {
			if !fieldOnly {
				fmt.Printf("asn,")
			}
			fmt.Printf("country_name\n")
		}

		for _, d := range asnDetails {
			if !fieldOnly {
				fmt.Printf("%s,", d.ASN)
			}
			fmt.Printf("%v\n", d.CountryName)
		}
	case "allocated":
		if header {
			if !fieldOnly {
				fmt.Printf("asn,")
			}
			fmt.Printf("allocated\n")
		}

		for _, d := range asnDetails {
			if !fieldOnly {
				fmt.Printf("%s,", d.ASN)
			}
			fmt.Printf("%v\n", d.Allocated)
		}
	case "registry":
		if header {
			if !fieldOnly {
				fmt.Printf("asn,")
			}
			fmt.Printf("registry\n")
		}

		for _, d := range asnDetails {
			if !fieldOnly {
				fmt.Printf("%s,", d.ASN)
			}
			fmt.Printf("%v\n", d.Registry)
		}
	case "domain":
		if header {
			if !fieldOnly {
				fmt.Printf("asn,")
			}
			fmt.Printf("domain\n")
		}

		for _, d := range asnDetails {
			if !fieldOnly {
				fmt.Printf("%s,", d.ASN)
			}
			fmt.Printf("%v\n", d.Domain)
		}
	case "num_ips":
		if header {
			if !fieldOnly {
				fmt.Printf("asn,")
			}
			fmt.Printf("num_ips\n")
		}

		for _, d := range asnDetails {
			if !fieldOnly {
				fmt.Printf("%s,", d.ASN)
			}
			fmt.Printf("%v\n", d.NumIPs)
		}
	case "prefixes":
		if header {
			if !fieldOnly {
				fmt.Printf("asn,")
			}
			fmt.Printf("prefixes\n")
		}

		for _, d := range asnDetails {
			if !fieldOnly {
				fmt.Printf("%s,", d.ASN)
			}
			fmt.Printf("%v\n", d.Prefixes)
		}
	case "prefixes6":
		if header {
			if !fieldOnly {
				fmt.Printf("asn,")
			}
			fmt.Printf("prefixes6\n")
		}

		for _, d := range asnDetails {
			if !fieldOnly {
				fmt.Printf("%s,", d.ASN)
			}
			fmt.Printf("%v\n", d.Prefixes6)
		}
	case "peers":
		if header {
			if !fieldOnly {
				fmt.Printf("asn,")
			}
			fmt.Printf("peers\n")
		}

		for _, d := range asnDetails {
			if !fieldOnly {
				fmt.Printf("%s,", d.ASN)
			}
			fmt.Printf("%v\n", d.Peers)
		}
	case "upstreams":
		if header {
			if !fieldOnly {
				fmt.Printf("asn,")
			}
			fmt.Printf("upstreams\n")
		}

		for _, d := range asnDetails {
			if !fieldOnly {
				fmt.Printf("%s,", d.ASN)
			}
			fmt.Printf("%v\n", d.Upstreams)
		}
	case "downstreams":
		if header {
			if !fieldOnly {
				fmt.Printf("asn,")
			}
			fmt.Printf("downstreams\n")
		}

		for _, d := range asnDetails {
			if !fieldOnly {
				fmt.Printf("%s,", d.ASN)
			}
			fmt.Printf("%v\n", d.Downstreams)
		}
	default:
		if header {
			if !fieldOnly {
				fmt.Printf("asn,")
			}
			fmt.Printf("%s\n", field)
		}
	}

	return nil
}
