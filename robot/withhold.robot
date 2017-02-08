*** Settings ***
Library    RequestsLibrary

*** Test cases ***
Return withholding of a given value
    [Template]    Return calculation withold rate of given amount
    100    3
    1000    30
    10000    300


*** Keywords ***
Return calculation withold rate of given amount
    [Arguments]    ${amount}    ${expected}
    Create Session    service    http://localhost:3333
    ${resp}=	Get Request    service    /withold/amount/${amount}
    Should Be Equal As Strings    ${resp.status_code}    200
    Should Be Equal As Strings    ${resp.json()['value']}    ${expected}