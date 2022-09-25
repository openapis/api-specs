# API Specs

[![Build Status][build-status-svg]][build-status-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![License][license-svg]][license-url]

A collection of API specifications.

Reasons for building a set of API specifications include:

1. Easily use API tools like [Postman](https://www.getpostman.com/) and [MuleSoft](https://www.mulesoft.com/)
1. [Auto-generate SDKs](https://github.com/swagger-api/swagger-codegen)
1. Create [API Explorers](https://github.com/swagger-api/swagger-ui) when permitted

Some specifications are provided by API providers and others are built manually from documentation.

## Specs

| API | Type | Latest Local | Source |
|-----|------|--------------|--------|
| eBay | Official | [v1 beta 11.0](ebay) | [developer.ebay.com](https://developer.ebay.com/api-docs/buy/browse/overview.html) |
| Insightly | Official | [v2.3](insightly) | [api.insight.ly](https://api.insight.ly/) |
| Lyft | Official | v1.0.0 | [github.com](https://github.com/lyft/lyft-django-sample/blob/4f7b21926ad8081a0ce2fe79a3849cb6e46f6bf1/lyft-api.yml) |
| Salesforce IoT | Official | v43.0.0 | [developer.salesforce.com](https://developer.salesforce.com/docs/atlas.en-us.api_iot.meta/api_iot/intro_swagger_file.htm)

## Tools

* Validate Spec: https://apitools.dev/swagger-parser/online/
* Convert Swagger 2.0 to OpenAPI 3.0: https://www.npmjs.com/package/swagger2openapi

 [used-by-svg]: https://sourcegraph.com/github.com/openapis/api-specs/-/badge.svg
 [used-by-url]: https://sourcegraph.com/github.com/openapis/api-specs?badge
 [build-status-svg]: https://github.com/openapis/api-specs/workflows/test/badge.svg
 [build-status-url]: https://github.com/openapis/api-specs/actions
 [goreport-svg]: https://goreportcard.com/badge/github.com/openapis/api-specs
 [goreport-url]: https://goreportcard.com/report/github.com/openapis/api-specs
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/openapis/api-specs
 [docs-godoc-url]: https://pkg.go.dev/github.com/openapis/api-specs
 [loc-svg]: https://tokei.rs/b1/github/openapis/api-specs
 [repo-url]: https://github.com/openapis/api-specs
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/openapis/api-specs/blob/master/LICENSE
