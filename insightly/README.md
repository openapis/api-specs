# Insightly API v2.2

## `date-time` 

The Insightly API Swagger 2.0 spec uses the `date-time` format but the API does not used the required RFC-3339 time format. The time is specified below.

For code generation, a lower effort way to correct this is to remove the `date-time` `format` and have the dates converted to strings.

This is from the Insightly API docs:

https://api.insight.ly/v2.2/Help#!/Overview/Introduction

> Date Formatting
> 
> The preferred date format for a date parameter in a URL is M/d/yyyy h:mm:ss AM/PM. For example, 11/7/2015 8:07:05 AM. So an example URL would be:
> 
> `https://api.insight.ly/v2.2/Contacts/Search?updated_after_utc=11/7/2015 8:07:05 AM`
> 
> Since it is a URL, you may need to encode the spaces as %20, e.g.:
> 
> `https://api.insight.ly/v2.2/Contacts/Search?updated_after_utc=3/13/2016%208:40:18%20PM`
> 
> By contrast, the preferred date format for object data is yyyy-MM-dd HH:mm:ss. (Note that military time is used here instead of AM/PM.) For example, a task's REMINDER_DATE_UTC could be 2015-04-10 21:15:00.