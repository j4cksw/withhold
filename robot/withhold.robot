*** Settings ***
Library    RequestsLibrary

*** Test cases ***
Return calculation withold rate of given amount
    Create Session    service    http://localhost:3333
    ${resp}=	Get Request    service    /withold/1000
    Should Be Equal As Strings    ${resp.status_code}    200
    Should Be Equal As Strings    ${resp.json()['value']}    30

Return calculation withold rate of given amount
    Create Session    service    http://localhost:3333
    ${resp}=	Get Request    service    /withold/10000
    Should Be Equal As Strings    ${resp.status_code}    200
    Should Be Equal As Strings    ${resp.json()['value']}    300