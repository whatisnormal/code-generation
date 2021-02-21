# SMS Confirmations
2 phase application that generates a code for a unique identifier and sends sms. Useful for marketing campaigns or 2-factor security validations.

## Phase 1 - Code generation
1) Receives a cell phone number (msisdn) from an external client
2) Generates a unique alfanumeric code and associates with the msisdn.
3) Returns to the client

## Phase 2 - Confirmation
1) Receives the code generated in *phase 1* and checks with the last code for the given msisdn. If matches, returns ok, else incorrect input.

# Supported Notifier provider
## Twillio
Requires usage of the following command-line arguments
_TWILLIO_FROM_MSISDN_ - The registered phone number.
_TWILLIO_ACCOUNT_SID_ - Your Account service id.
_TWILLIO_AUTH_TOKEN_  - Authorization token.