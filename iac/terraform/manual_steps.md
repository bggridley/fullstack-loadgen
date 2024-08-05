## Steps


Go to subscriptions in Azure

go to settings > resource providers

Search and select cloudshell

Hit Register !

Now launch the cloud shell
Ignore storage account because that's lame !!

May need to change this later...

But now, run:

az ad sp create-for-rbac --name "terraform-sp" --role Contributor --scopes /subscriptions/<subscription-id> --sdk-auth


<!-- Now with that json output:

put ARM_CLIENT_ID as clientId
put ARM_CLIENT_SECRET as clientSecret
put ARM_TENANT_ID as tenantId

into github actions secrets! -->

Now add the output into AZURE_CREDENTIALS

