*** Settings ***
Library    RequestsLibrary

*** Test cases ***
Return calculation withold rate of given amount
    Create Session    service    http://localhost:3333
    ${resp}=	Get Request    service    /
    Should Be Equal As Strings    ${resp.status_code}    200