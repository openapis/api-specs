Get Satisfaction
================

## Notes

1. The Swagger 2.0 YAML spec is currently built by hand
2. The Swagger 2.0 JSON spec is auto-converted using [Swagger Parser](http://bigstickcarpet.com/swagger-parser/)
3. The RAML 0.8 spec is auto-converted using [API Transformer](https://apitransformer.com/)

## Issues

1. The current Swagger specs are based on the JSON response format type. The XML format response uses the Atom format and provides the following additional response properties:
    1. paging links for `self`, `first`, `next`, and `last`
    2. result set properties including `question_count`, `idea_count`, `problem_count`, `talk_count`, `unanswered_count`
2. Swaggger UI: The Get Satisfaction API does not work with Swagger UI's "Try It Now" feature as it does not return the `Access-Control-Allow-Origin` header to support CORS. More information on this is available [here](https://github.com/swagger-api/swagger-ui/blob/master/README.md#cors-support).
3. Postman: As of now, RAML support is preferred when using Postman due to [Swagger 2.0 compatibility](https://github.com/postmanlabs/postman-app-support/issues/843). Specifically, Swagger 2.0 will substitute `{{variable}}` as default values causing mismatches in API calls. RAML does not load any default vaules.