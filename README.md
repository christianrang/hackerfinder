# hackerfinder

> _WARNING_: If you are looking to use the pkg files to access the APIs. This repo is not stable and everything is subject to change.

Uses VirusTotal and AbuseIP to look up malicious things

## Getting Started

- Navigate to [releases](https://github.com/christianrang/hackerfinder/releases)

- Select the `Assets` drop down for version you would like and the correct OS / Arch for your computer

- See the configuration section below for more information on creating the config

## Outputs

### IP Tables

#### Fields

|        Header        |                                Value                                |
| :------------------: | :-----------------------------------------------------------------: |
|          IP          |                Contains the IP address searched for                 |
|         VT M         |                VirusTotal Malicious vote count field                |
|         VT S         |               VirusTotal Suspicious vote count field                |
|         VT H         |                VirusTotal Harmless vote count field                 |
|         VT U         |               VirusTotal Undetected vote count field                |
|  AbuseIp Conf Score  |           abuseaipdb Confidence score (100% is malicious)           |
| AbuseIp Report Count |  abuseaipdb number of reports that the IP is malicious from users   |
|    AbuseIp Users     | abuseaipdb unique users creating requests of the IP being malicious |
|  AbuseIp Hostnames   |         abuseaipdb reported domain / hostnames for this IP          |

### Domain Tables

#### Fields

| Header |                 Value                  |
| :----: | :------------------------------------: |
| DOMAIN |  Contains the IP address searched for  |
|  VT M  | VirusTotal Malicious vote count field  |
|  VT S  | VirusTotal Suspicious vote count field |
|  VT H  |  VirusTotal Harmless vote count field  |
|  VT U  | VirusTotal Undetected vote count field |

### Hash Tables

#### Fields

| Header |                 Value                  |
| :----: | :------------------------------------: |
|  HASH  |  Contains the IP address searched for  |
|  VT M  | VirusTotal Malicious vote count field  |
|  VT S  | VirusTotal Suspicious vote count field |
|  VT H  |  VirusTotal Harmless vote count field  |
|  VT U  | VirusTotal Undetected vote count field |

## Configuration

The program looks for configuration in `$HOME/.config` and `$PWD`. Configuration filename must be `hackerfinder.yaml`.

```yaml
virustotal:
  api_key: { your VirusTotal api key }
  # This currently doesn't do anything. In the future it may be removed or used to cacluate runtime.
  premium: { bool "true/false" }
abuseaipdb:
  api_key: { your abuseaipdb api key }
```

## TODOs

- Create a subcommand for install

  - Should include configuration setup

  - Adding binary to path

  - Adding binary to path
