# Tests
> Important to be run before push or PR


## Main Test
> The purpose of this test is to ensure the basic features are working as expected
>
> Requirements:
> - The database must be empty when starting this test (before compose up)
> - At least 3 Stripe Account are needed:
>   - Plateform using stripe connect - application fees (Bount.ing)
>   - Sponsor (User who set the bounties)
>   - Developer (User who resolve and claim the bounties)
>
> - At least 2 Github Account are needed:
>   - Sponsor (User who set the bounties)
>   - Developer (User who resolve and claim the bounties)
>
> - A Github OAuth App

### Use Cases
- General:
  - Anyone should be able to:
    - Create an account (SignUp)
    - Receive an activation code by mail
    - Set a Password
    - Log in (SignIn)
    - Reset a Password

- Stripe:
  - Anyone should be able to:
    - Link an account
    - Add a payment method (Card)
    - Payout a bounty

- Github:
  - Anyone should be able to:
    - Link an account
    - Get a list of issues of owned repos
    - Set a bounty on any (open) issue
    - Set a bounty on any (open) issue that are from a repo not owned via URL

- Invoices:
  - When a bounty is approved:
    - The dev must receive an invoice for the fees
    - The dev must be able to generate an invoice for the sponsor
    - The dev must be able to send the generated invoice for the sponsor
