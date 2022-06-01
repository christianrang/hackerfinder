package domain_test

import (
	        "log"
		        "os"
			        "testing"

				        "github.com/christianrang/find-bad-ip/pkg/vtsdk"
					        "github.com/christianrang/find-bad-ip/pkg/vtsdk/domain"
						        "github.com/stretchr/testify/assert"
						)

						var configuration = vtsdk.Configuration{}

						func createConfiguration() {
							        configuration.SetApiKey(getApiKey())
							}

							func getApiKey() string {
								        apiKey, ok := os.LookupEnv("VT_API_KEY")
									        if !ok {
											                log.Fatalln("Please export a valid virustotal API Key\n\texport VT_API_KEY='API key here'")
													        }

														        return apiKey
														}

														func TestQueryDomain(t *testing.T) {
															        createConfiguration()
