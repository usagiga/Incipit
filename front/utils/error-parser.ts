/**
 * ErrorParser provides functions to treat API error
 */
class ErrorParser {
  static instance: ErrorParser

  /**
   * Map of error codes
   */
  static errorMap: Map<number, Map<number, string>> = new Map([
    // 101 Admin User Validation
    [101, new Map([
      [101, '101-101 Name is too short.'],
      [102, '101-102 Name is too long.'],
      [103, '101-103 Name has unavailable char.'],
      [201, '101-201 ScreenName is too short.'],
      [202, '101-202 ScreenName is too long.'],
      [203, '101-203 ScreenName has unavailable char.'],
      [301, '101-301 Password is too short.'],
      [302, '101-302 Password is too long.'],
      [303, '101-303 Password has unavailable char.']
    ])],
    // 102 Link Validation
    [102, new Map([
      [101, '102-101 URL is Incipit. Mustn\'t be itself.\n'],
      [102, '102-102 URL is invalid.']
    ])],
    // 201 CRED Admin
    [201, new Map([
      [101, '201-101 Failed add admin.'],
      [201, '201-201 Failed find admin.'],
      [202, '201-202 Finding admin not found.'],
      [301, '201-301 Failed update admin.'],
      [302, '201-302 Updating admin not found.'],
      [401, '201-401 Failed delete admin.']
    ])],
    // 202 Authorization(Admin)
    [202, new Map([
      [101, '202-101 Failed to find user.'],
      [102, '202-102 Unmatch password.'],
      [103, '202-103 Access token is expired.'],
      [104, '202-104 Failed to store generated token.']
    ])],
    // 203 Hash
    [203, new Map([
      [101, '203-101 Failed generate hash.'],
      [102, '203-102 Failed compare hash.']
    ])],
    // 204 Installer
    // None yet
    // 205 CRED Link
    [205, new Map([
      [101, '205-101 Failed add link.'],
      [201, '205-201 Failed find link.'],
      [202, '205-202 Finding link not found.'],
      [301, '205-301 Failed update link.'],
      [302, '205-302 Updating link not found.'],
      [401, '205-401 Failed delete link.']
    ])],
    // 301 Handle request of Admin User
    [301, new Map([
      [101, '301-101 Failed to bind JSON. Request body is not JSON or invalid.']
    ])],
    // 302 Handle request of Authorization (Admin)
    [302, new Map([
      [101, '302-101 Failed to bind JSON. Request body is not JSON or invalid.'],
      [102, '302-102 Need Authorization Header.']
    ])],
    // 303 Handle request of Link
    [303, new Map([
      [101, '303-101 Failed to bind JSON. Request body is not JSON or invalid.']
    ])],
    // 304 Handle request of Install
    [304, new Map([
      [101, '304-101 Failed to bind JSON. Request body is not JSON or invalid.']
    ])]
  ])

  /**
   * Get its instance as singleton
   */
  static getInstance(): ErrorParser {
    if (this.instance === undefined || this.instance == null) {
      this.instance = new ErrorParser()
    }

    return this.instance
  }

  /**
   * Get error message from error code pair
   * @param pCode Primary Code
   * @param sCode Secondary Code
   */
  getErrorMessage(pCode: number, sCode: number): string {
    const region = ErrorParser.errorMap.get(pCode)
    const errMsg = region?.get(sCode)

    return errMsg ?? 'Unknown error occurred.'
  }

  /**
   * Purify error before processing normal scenario
   * @param resJson JSON on intercepting response
   * @return Promise<response JSON>
   */
  interceptErrorResp(resJson: any): Promise<any> {
    if (resJson?.type !== 'error') {
      return resJson
    }

    // eslint-disable-next-line camelcase
    throw new Error(this.getErrorMessage(resJson?.p_code, resJson?.s_code))
  }
}

// Export
export default () => ErrorParser.getInstance()
