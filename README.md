# hackerfinder

*WARNING*: If you are looking to use the pkg files to access the APIs. This repo is by no means stable and everything is subject to change.

Uses VirusTotal to look up malicious things

## Outputs

### IP Tables

#### Fields

|Header|Value|
|:----:|:---:|
|IP|Contains the IP address searched for|
|VT M| VirusTotal Malicious vote count field|
|VT S| VirusTotal Suspicious vote count field|
|VT H| VirusTotal Harmless vote count field|
|VT U| VirusTotal Undetected vote count field|
|AbuseIp Conf Score| abuseaipdb Confidence score (100% is malicious)|
|AbuseIp Report Count| abuseaipdb number of reports that the IP is malicious from users|
|AbuseIp Users|abuseaipdb unique users creating requests of the IP being malicious|
|AbuseIp Hostnames| abuseaipdb reported domain / hostnames for this IP|

## Configuration

The program looks for configuration in `$HOME/.config` and `$PWD`.

```yaml
virustotal:
  api_key: { your VirusTotal api key}
abuseaipdb:
  api_key: { your abuseaipdb api key}
```
