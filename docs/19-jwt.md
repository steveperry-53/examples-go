# jwt.go

Create base64-encoded header and claim set.

## Building and running the program
    
The source file, token.go, is in this directory:

    $GOPATH/github.com/steveperry-53/examples-go/19-jwt
    
Navigate to anywhere, for example $HOME, and enter this command;

    go install github.com/steveperry-53/examples-go/19-jwt
    
The executable file, 19-jwt, is in this directory:

    $GOPATH/bin
    
To run the executable, enter this command:

    $GOPATH/bin/19-jwt

## Discussion

The output of this program is

<base64-encoded header>.<base-64 encoded claim set>

Sign the output.
Base64 encode the signature.

Then create this JWT:
<base64-encoded header>.<base-64 encoded claim set>.<base64-encoded signature>

Pass the JWT to the auth server.

POST /token HTTP/1.1
Host: oauth2.googleapis.com
Content-Type: application/x-www-form-urlencoded

grant_type=urn%3Aietf%3Aparams%3Aoauth%3Agrant-type%3Ajwt-bearer&assertion=<jwt>

OR

curl -d 'grant_type=urn%3Aietf%3Aparams%3Aoauth%3Agrant-type%3Ajwt-bearer&assertion=<jwt>
' https://oauth2.googleapis.com/token


