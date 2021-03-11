// +build examples

/**
 * (C) Copyright IBM Corp. 2021.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package enterprisemanagementv1_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/enterprisemanagementv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the Enterprise Management service.
//
// The following configuration properties are assumed to be defined:
// ENTERPRISE_MANAGEMENT_URL=<service base url>
// ENTERPRISE_MANAGEMENT_AUTH_TYPE=iam
// ENTERPRISE_MANAGEMENT_APIKEY=<IAM apikey>
// ENTERPRISE_MANAGEMENT_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
const externalConfigFile = "../enterprise_management_v1.env"

var (
	enterpriseManagementService *enterprisemanagementv1.EnterpriseManagementV1
	config                      map[string]string
	configLoaded                bool = false
)

func shouldSkipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

var _ = Describe(`EnterpriseManagementV1 Examples Tests`, func() {
	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(enterprisemanagementv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}

			configLoaded = len(config) > 0
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			enterpriseManagementServiceOptions := &enterprisemanagementv1.EnterpriseManagementV1Options{}

			enterpriseManagementService, err = enterprisemanagementv1.NewEnterpriseManagementV1UsingExternalConfig(enterpriseManagementServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(enterpriseManagementService).ToNot(BeNil())
		})
	})

	Describe(`EnterpriseManagementV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateAccountGroup request example`, func() {
			// begin-create_account_group

			createAccountGroupOptions := enterpriseManagementService.NewCreateAccountGroupOptions(
				"testString",
				"testString",
				"testString",
			)

			createAccountGroupResponse, response, err := enterpriseManagementService.CreateAccountGroup(createAccountGroupOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(createAccountGroupResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_account_group

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createAccountGroupResponse).ToNot(BeNil())

		})
		It(`ListAccountGroups request example`, func() {
			// begin-list_account_groups

			listAccountGroupsOptions := enterpriseManagementService.NewListAccountGroupsOptions()

			listAccountGroupsResponse, response, err := enterpriseManagementService.ListAccountGroups(listAccountGroupsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(listAccountGroupsResponse, "", "  ")
			fmt.Println(string(b))

			// end-list_account_groups

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listAccountGroupsResponse).ToNot(BeNil())

		})
		It(`GetAccountGroup request example`, func() {
			// begin-get_account_group

			getAccountGroupOptions := enterpriseManagementService.NewGetAccountGroupOptions(
				"testString",
			)

			accountGroup, response, err := enterpriseManagementService.GetAccountGroup(getAccountGroupOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(accountGroup, "", "  ")
			fmt.Println(string(b))

			// end-get_account_group

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(accountGroup).ToNot(BeNil())

		})
		It(`UpdateAccountGroup request example`, func() {
			// begin-update_account_group

			updateAccountGroupOptions := enterpriseManagementService.NewUpdateAccountGroupOptions(
				"testString",
			)

			response, err := enterpriseManagementService.UpdateAccountGroup(updateAccountGroupOptions)
			if err != nil {
				panic(err)
			}

			// end-update_account_group

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`ImportAccountToEnterprise request example`, func() {
			// begin-import_account_to_enterprise

			importAccountToEnterpriseOptions := enterpriseManagementService.NewImportAccountToEnterpriseOptions(
				"testString",
				"testString",
			)

			response, err := enterpriseManagementService.ImportAccountToEnterprise(importAccountToEnterpriseOptions)
			if err != nil {
				panic(err)
			}

			// end-import_account_to_enterprise

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

		})
		It(`CreateAccount request example`, func() {
			// begin-create_account

			createAccountOptions := enterpriseManagementService.NewCreateAccountOptions(
				"testString",
				"testString",
				"testString",
			)

			createAccountResponse, response, err := enterpriseManagementService.CreateAccount(createAccountOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(createAccountResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_account

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(createAccountResponse).ToNot(BeNil())

		})
		It(`ListAccounts request example`, func() {
			// begin-list_accounts

			listAccountsOptions := enterpriseManagementService.NewListAccountsOptions()

			listAccountsResponse, response, err := enterpriseManagementService.ListAccounts(listAccountsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(listAccountsResponse, "", "  ")
			fmt.Println(string(b))

			// end-list_accounts

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listAccountsResponse).ToNot(BeNil())

		})
		It(`GetAccount request example`, func() {
			// begin-get_account

			getAccountOptions := enterpriseManagementService.NewGetAccountOptions(
				"testString",
			)

			account, response, err := enterpriseManagementService.GetAccount(getAccountOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(account, "", "  ")
			fmt.Println(string(b))

			// end-get_account

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(account).ToNot(BeNil())

		})
		It(`UpdateAccount request example`, func() {
			// begin-update_account

			updateAccountOptions := enterpriseManagementService.NewUpdateAccountOptions(
				"testString",
				"testString",
			)

			response, err := enterpriseManagementService.UpdateAccount(updateAccountOptions)
			if err != nil {
				panic(err)
			}

			// end-update_account

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`CreateEnterprise request example`, func() {
			// begin-create_enterprise

			createEnterpriseOptions := enterpriseManagementService.NewCreateEnterpriseOptions(
				"testString",
				"testString",
				"testString",
			)

			createEnterpriseResponse, response, err := enterpriseManagementService.CreateEnterprise(createEnterpriseOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(createEnterpriseResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_enterprise

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(createEnterpriseResponse).ToNot(BeNil())

		})
		It(`ListEnterprises request example`, func() {
			// begin-list_enterprises

			listEnterprisesOptions := enterpriseManagementService.NewListEnterprisesOptions()

			listEnterprisesResponse, response, err := enterpriseManagementService.ListEnterprises(listEnterprisesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(listEnterprisesResponse, "", "  ")
			fmt.Println(string(b))

			// end-list_enterprises

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listEnterprisesResponse).ToNot(BeNil())

		})
		It(`GetEnterprise request example`, func() {
			// begin-get_enterprise

			getEnterpriseOptions := enterpriseManagementService.NewGetEnterpriseOptions(
				"testString",
			)

			enterprise, response, err := enterpriseManagementService.GetEnterprise(getEnterpriseOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(enterprise, "", "  ")
			fmt.Println(string(b))

			// end-get_enterprise

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(enterprise).ToNot(BeNil())

		})
		It(`UpdateEnterprise request example`, func() {
			// begin-update_enterprise

			updateEnterpriseOptions := enterpriseManagementService.NewUpdateEnterpriseOptions(
				"testString",
			)

			response, err := enterpriseManagementService.UpdateEnterprise(updateEnterpriseOptions)
			if err != nil {
				panic(err)
			}

			// end-update_enterprise

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})
})
