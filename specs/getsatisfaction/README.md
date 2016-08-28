Get Satisfaction
================

## Notes

1. The current Swagger specs are based on the JSON response format type. The XML format response uses the Atom format and provides the following additional response properties:
    1. paging links for `self`, `first`, `next`, and `last`
    2. result set properties including `question_count`, `idea_count`, `problem_count`, `talk_count`, `unanswered_count`
2. The Get Satisfaction API does not work with Swagger UI's "Try It Now" feature as it does not return the `Access-Control-Allow-Origin` header to support CORS. More information on this is available [here](https://github.com/swagger-api/swagger-ui/blob/master/README.md#cors-support).