S3 Club 7
==

Given a set of valid FT credentials, as recognised by Flex, accept a file upload and ~~reach for the stars~~ proxy into an S3 bucket.

This service replaces the kenny-logins project.

# Building

## Install dependencies

```
go get -u github.com/aws/aws-sdk-go
go get -u github.com/gorilla/securecookie
```

## Login into quay
```docker login -e="." -u="financialtimes+ftrobot" -p="XXX" quay.io```

## Build
```make dist```

## Running
```docker run quay.io/financialtimes/s3-club-7 -cluster={clustername} -flex-api=https://master-flex-{clustername}.ft.com/api```

Licence
--

Copyright (c) 2016 Financial Times

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
