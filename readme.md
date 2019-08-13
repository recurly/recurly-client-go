# Recurly Golang Client

[APIv3 Documentation](https://partner-docs.recurly.com)

Do not add any external dependencies! Use standard golang libraries only.

### Example

Super early prototype

```
$ make test

Recurly 2019/08/12 21:35:40.094087 Requesting https://partner-api.recurly.com/sites/12345678/accounts/code-robot-overlord
Recurly 2019/08/12 21:35:40.619377 Response: 200, 0.525 sec
Recurly 2019/08/12 21:35:40.619452 Headers: status code: 200, request ID: 5057fbd20933cce2-EWR, version: recurly.v2018-08-09, limit: 1996 remaining of 2000
Body:
{"id":"lcdysuuhicby","object":"account","code":"robot-overlord","parent_account_id":null,"bill_to":"self","state":"active","username":"","email":"robots-rule@gmail.com","cc_emails":"","preferred_locale":"","first_name":"Robot","last_name":"Overloard","company":"","vat_number":"","tax_exempt":false,"exemption_certificate":null,"address":{"phone":"","street1":"","street2":"","city":"","region":"","postal_code":"","country":""},"billing_info":null,"shipping_addresses":[],"custom_fields":[],"created_at":"2019-08-13T03:53:15Z","updated_at":"2019-08-13T03:53:15Z","deleted_at":null}
2019/08/12 21:35:40.619584 Received account: {"code":"robot-overlord","username":"","email":"robots-rule@gmail.com","preferred_locale":"","cc_emails":"","first_name":"Robot","last_name":"Overloard","created_at":"2019-08-13T03:53:15Z","updated_at":"2019-08-13T03:53:15Z","deleted_at":null}
```
