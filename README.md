# Kusari Inspector Demo

This repo contains example code used to showcase the features of [Kusari Inspector](https://kusari.dev/inspector).
See instructions below for instructions on self-guided demos.


## Public uses of Kusari Inspector

* [Inspector recommends merging](https://github.com/search?q=is%3Apr+commenter%3Akusari-inspector%5Bbot%5D+-org%3Akusaridev+-org%3AKusari-Sandbox+%22%E2%9C%85+PROCEED%22&type=pullrequests&query=is%3Apr+commenter%3Akusari-inspector%5Bbot%5D&s=created&o=desc)
* [Inspector recommends not merging](https://github.com/search?q=is%3Apr+commenter%3Akusari-inspector%5Bbot%5D+-org%3Akusaridev+-org%3AKusari-Sandbox+%22DO+NOT+PROCEED%22&type=pullrequests&query=is%3Apr+commenter%3Akusari-inspector%5Bbot%5D&s=created&o=desc)

## Self-guided demos

To try Kusari Inspector:

1. [Create a new repo from this template](https://github.com/new?template_name=inspector-demo-vulnerable&template_owner=Kusari-Sandbox)
   - **Check the "Include all branches" box**
   - Select the namespace for the new repo
   - Add a name for the new repo
2. [Install Kusari Inspector](https://github.com/apps/kusari-inspector) to your new repo
3. Open a pull request to merge one of the branches below into `main`.
4. In a few seconds, you'll see what Kusari Inspector says!

You can also modify the code in the repo to try other variations.

## Demo branches

The following branches will demonstrate various Kusari Inspector use cases.

| Branch | Description
| ------ | -----------
| laravel_api_key     | Find secrets in a Laravel application environment
| python_maintenance  | Python libraries with non-blocking maintenance issues
| python_typosquat    | Typosquatted Python library
| workflow_secrets    | Various problems with GitHub Workflow secret handling

## Contributing

Since this is a product demo repo, we don't intend to receive external contributions.
If an example is not working as expected, please [open an issue](https://github.com/Kusari-Sandbox/inspector-demo/issues/new/choose) or email support@kusari.dev.
