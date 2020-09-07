// To write Object for JSON as exact form
/* eslint-disable object-shorthand */

import { VueRouter } from 'vue-router/types/router'
import TokenStore from './token-store'

/**
 * IncipitApi provides functions to treat Incipit API
 */
class IncipitApi {
  private constructor (router: VueRouter) {
    const apiBaseUrl = process.env.apiBaseUrl
    if (apiBaseUrl === undefined) {
      throw new Error('Env var "INCIPIT_API" is not set.')
    }

    this.apiBaseUrl = apiBaseUrl
    this.router = router
  }

  static instance: IncipitApi

  /**
   * Get its instance as singleton
   * @param router VueRouter to redirect when authorizing
   */
  static getInstance (router: VueRouter): IncipitApi {
    if (this.instance === undefined || this.instance == null) {
      this.instance = new IncipitApi(router)
    }

    return this.instance
  }

  apiBaseUrl: string
  router: VueRouter

  /**
   * Get all links
   */
  getLinks (): Promise<void> {
    const url = new URL('link', this.apiBaseUrl)
    const req = new Request(
      url.href,
      {
        method: 'GET',
        headers: {
          authorization: 'Bearer ' + TokenStore.accessToken
        }
      }
    )

    return fetch(req)
      .then(res => res.json())
      .then(resJson => this.interceptInstall(resJson))
      .then(resJson => this.interceptAuthorize(resJson, () => this.getLinks()))
  }

  /**
   * Get specific link
   * @param shortId short ID of link
   */
  getLinkByShortID (shortId: string): Promise<void> {
    const url = new URL('link', this.apiBaseUrl)
    const reqBody = { shortId }
    const req = new Request(
      url.href,
      {
        method: 'GET',
        body: JSON.stringify(reqBody),
        headers: {
          'Content-Type': 'application/json',
          authorization: 'Bearer ' + TokenStore.accessToken
        }
      }
    )

    return fetch(req)
      .then(res => res.json())
      .then(resJson => this.interceptInstall(resJson))
  }

  /**
   * Register first administrator to set up
   * @param name Name of administrator
   * @param screenName Screen name of administrator
   * @param password Password of administrator
   */
  install (name: string, screenName: string, password: string): Promise<void> {
    const url = new URL('install', this.apiBaseUrl)
    const reqBody = {
      name: name,
      screen_name: screenName,
      password: password
    }
    const req = new Request(
      url.href,
      {
        method: 'POST',
        body: JSON.stringify(reqBody),
        headers: {
          'Content-Type': 'application/json'
        }
      }
    )

    return fetch(req)
      .then(res => res.json())
      .then(resJson => this.interceptInstalled(resJson))
  }

  /**
   * Login as administrator
   * @param name Name of administrator
   * @param password Password of administrator
   */
  login (name: string, password: string): Promise<void> {
    const url = new URL('login', this.apiBaseUrl)
    const reqBody = {
      name,
      password
    }
    const req = new Request(
      url.href,
      {
        method: 'POST',
        body: JSON.stringify(reqBody),
        headers: {
          'Content-Type': 'application/json'
        }
      }
    )

    return fetch(req)
      .then(res => res.json())
      .then(resJson => this.interceptInstall(resJson))
  }

  /**
   * Renew access / refresh token pair
   * @return Promise<response JSON>
   */
  renewToken (): Promise<void> {
    const url = new URL('login/refresh', this.apiBaseUrl)
    const reqBody = { refresh_token: TokenStore.refreshToken }
    const req = new Request(
      url.href,
      {
        method: 'POST',
        body: JSON.stringify(reqBody),
        headers: {
          'Content-Type': 'application/json'
        }
      }
    )

    return fetch(req)
      .then(res => res.json())
      .then(resJson => this.interceptInstall(resJson))
      .then((resJson) => {
        if (resJson.type !== 'refresh_token_admin') {
          throw new Error('Can\'t refresh token')
        }

        TokenStore.accessToken = resJson.access_token.token
        TokenStore.refreshToken = resJson.refresh_token.token

        return Promise.resolve()
      })
  }

  /**
   * Purify install error before processing normal scenario
   * @param resJson JSON on intercepting response
   * @return Promise<response JSON>
   */
  interceptInstall (resJson: any): Promise<any> {
    if (resJson.type !== 'needed_install') {
      return resJson
    }

    return this.router.push('/install')
  }

  /**
   * Purify installed error before processing normal scenario
   * @param resJson JSON on intercepting response
   * @return Promise<response JSON>
   */
  interceptInstalled (resJson: any): Promise<any> {
    if (resJson.type !== 'redundant_install') {
      return resJson
    }

    return this.router.push('/')
  }

  /**
   * Purify authorize error before processing normal scenario
   * @param resJson JSON on intercepting response
   * @param retryFunc If failed by expired tokens, run this handler after renewing it
   * @return Promise<response JSON> or Promise<Route>
   */
  interceptAuthorize (resJson: any, retryFunc: () => Promise<any>): Promise<any> {
    if (resJson.type !== 'error') {
      return resJson
    }

    // Renew access token
    if (resJson.p_code === 202 && resJson.s_code === 103) {
      return this.renewToken()
        .then(retryFunc)
        .catch(
          () => this.router.push('/login')
        )
    }

    return resJson
  }
}

/**
 * Get IncipitApi instance as singleton
 * @param router
 */
export default (router: VueRouter) => IncipitApi.getInstance(router)
