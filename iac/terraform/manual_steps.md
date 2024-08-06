## Steps


Go to subscriptions in Azure

go to settings > resource providers

Search and select cloudshell

Hit Register !

Now launch the cloud shell
Ignore storage account because that's lame !!

May need to change this later...

But now, run:

`az ad sp create-for-rbac --name "terraform-sp" --role Contributor --scopes /subscriptions/<subscription-id> --sdk-auth`

This is needed to login to Azure from github!
Go to secrets within Entra app terraform-sp, create a federated credential based on the branch, in this case 'terraform' under this repository.

Now do :

```
az role definition create --role-definition '{
    "Name": "CustomAuthorizationWriter",
    "Description": "Custom role for Microsoft.Authorization write operations",
    "Actions": [
        "Microsoft.Authorization/*/Write"
    ],
    "AssignableScopes": ["/subscriptions/<subscription id>"]
}'
```

And then this:

`az role assignment create --assignee "<client ID of terraform-sp>" --role "CustomAuthorizationWriter" --scope "/subscriptions/<subscription id>"`

<!-- Now with that json output:

put ARM_CLIENT_ID as clientId
put ARM_CLIENT_SECRET as clientSecret
put ARM_TENANT_ID as tenantId

into github actions secrets! -->

Now add the output into `AZURE_CREDENTIALS` secret on GitHub repo

### Create a storage account and a new container for terraform state file

Then create a resource group:

`az group create --name fl-rg --location WestUS`

Then create a storage account:

`az storage account create --name flterraformsa --resource-group fl-rg --location WestUS --sku Standard_LRS`

Then create a storage container:

`az storage container create --name terraform --account-name flterraformsa`

This is reflected in main.tf