/**
 * TokenStore provides functions to treat access / refresh token
 */
class TokenStore {
  private _accessTokenKey: string = 'incipit_access_token'
  private _refreshTokenKey: string = 'incipit_refresh_token'

  static instance: TokenStore

  /**
   * Get its instance as singleton
   */
  static getInstance (): TokenStore {
    if (this.instance === undefined || this.instance == null) {
      this.instance = new TokenStore()
    }

    return this.instance
  }

  /**
   * Get access token from local storage.
   * If there's no token, return empty string.
   */
  get accessToken (): string {
    const token = localStorage.getItem(this._accessTokenKey)
    if (token === null) {
      return ''
    }

    return token
  }

  /**
   * Set access token into local storage.
   */
  set accessToken (token: string) {
    localStorage.setItem(this._accessTokenKey, token)
  }

  /**
   * Get refresh token from local storage.
   * If there's no token, return empty string.
   */
  get refreshToken (): string {
    const token = localStorage.getItem(this._refreshTokenKey)
    if (token === null) {
      return ''
    }

    return token
  }

  /**
   * Set refresh token into local storage.
   */
  set refreshToken (token: string) {
    localStorage.setItem(this._refreshTokenKey, token)
  }
}

// Export
export default TokenStore.getInstance()
