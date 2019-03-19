# token.go

Gets an access token from a JSON key file.

## Building and running the program
    
The source file, token.go, is in this directory:

    $GOPATH/github.com/steveperry-53/examples-go/18-token
    
Navigate to anywhere, for example $HOME, and enter this command;

    go install github.com/steveperry-53/examples-go/18-token
    
The executable file, 18-token, is in this directory:

    $GOPATH/bin
    
To run the executable, enter this command:

    $GOPATH/bin/18-token

## Discussion

Create a GCP service account, and download a JSON key file.

This program reads the JSON key file, and produces and access
token that you can use to call BigQuery APIs.

We call CredentialsFromJSON to get a Credentials object.
We supply a Context (in a simple form), the contents of the 
JSON key file, and a list of scopes.

creds, _ = google.CredentialsFromJSON(ctx, jsonKey, "https://www.googleapis.com/auth/bigquery")

From the Credentials object, we get a TokenSource. From that, we
get a Token. And from that, we get an AccessToken.

We can use the AccessToken to call a BigQuery API.

curl -H "Authorization: Bearer ya29.c.ElrRBu6lz3VQPrGhMrzymWCAYaLAgx0xqoxgWch22IG13Ys2ORszK78PlpCYm5TAYcMfadwsPNzOkg8A_FG1eHYlSzmLhvNCzbafNBnWG9-2O6ifKbLcb8iWIrk" https://www.googleapis.com/bigquery/v2/projects






