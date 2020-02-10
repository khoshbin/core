// Code generated by swagger-to-go. DO NOT EDIT.
package companypb

import (
	"encoding/json"

	"elbix.dev/engine/pkg/assert"
	"elbix.dev/engine/pkg/grpcgw"
)

const paths = "{\"/v1/companies\":{\"post\":{\"operationId\":\"CreateCompany\",\"parameters\":[{\"in\":\"body\",\"name\":\"body\",\"required\":true,\"schema\":{\"$ref\":\"#/definitions/companyCreateCompanyRequest\"}}],\"responses\":{\"200\":{\"description\":\"A successful response.\",\"schema\":{\"$ref\":\"#/definitions/companyCreateCompanyResponse\"}}},\"security\":[{\"Authentication\":[]}],\"tags\":[\"CompanyService\"]}},\"/v1/companies/{id}\":{\"get\":{\"operationId\":\"GetCompany\",\"parameters\":[{\"format\":\"int64\",\"in\":\"path\",\"name\":\"id\",\"required\":true,\"type\":\"string\"}],\"responses\":{\"200\":{\"description\":\"A successful response.\",\"schema\":{\"$ref\":\"#/definitions/companyGetCompanyResponse\"}}},\"security\":[{\"Authentication\":[]}],\"tags\":[\"CompanyService\"]}}}"
const definitions = "{\"companyCompanyStatus\":{\"default\":\"COMPANY_STATUS_INVALID\",\"enum\":[\"COMPANY_STATUS_INVALID\",\"COMPANY_STATUS_PENDING\",\"COMPANY_STATUS_ACTIVE\",\"COMPANY_STATUS_BANNED\"],\"type\":\"string\"},\"companyCreateCompanyRequest\":{\"properties\":{\"id\":{\"format\":\"int64\",\"type\":\"string\"},\"name\":{\"type\":\"string\"},\"status\":{\"$ref\":\"#/definitions/companyCompanyStatus\"}},\"type\":\"object\"},\"companyCreateCompanyResponse\":{\"properties\":{\"id\":{\"format\":\"int64\",\"type\":\"string\"},\"name\":{\"type\":\"string\"},\"status\":{\"$ref\":\"#/definitions/companyCompanyStatus\"}},\"type\":\"object\"},\"companyGetCompanyResponse\":{\"properties\":{\"id\":{\"format\":\"int64\",\"type\":\"string\"},\"name\":{\"type\":\"string\"},\"status\":{\"$ref\":\"#/definitions/companyCompanyStatus\"}},\"type\":\"object\"}}"

func init() {
	var (
		p = make(map[string]interface{})
		d = make(map[string]interface{})
	)

	assert.Nil(json.Unmarshal([]byte(paths), &p))
	assert.Nil(json.Unmarshal([]byte(definitions), &d))

	grpcgw.RegisterSwagger(p, d)
}
