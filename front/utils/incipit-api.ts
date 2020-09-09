// To write Object for JSON as exact form
/* eslint-disable object-shorthand,camelcase */

import { VueRouter } from 'vue-router/types/router'
import TokenStore from './token-store'
import ErrorParser from './error-parser'

/**
 * IncipitApi provides functions to treat Incipit API
 */
class IncipitApi {
  private constructor ($router: VueRouter) {
    const apiBaseUrl = process.env.apiBaseUrl
    if (apiBaseUrl === undefined) {
      throw new Error('Env var "INCIPIT_API" is not set.')
    }

    this.apiBaseUrl = apiBaseUrl
    this.$router = $router
  }

  static instance: IncipitApi

  /**
   * Get its instance as singleton
   * @param $router VueRouter to redirect when authorizing
   */
  static getInstance ($router: VueRouter): IncipitApi {
    if (this.instance === undefined || this.instance == null) {
      this.instance = new IncipitApi($router)
    }

    return this.instance
  }

  apiBaseUrl: string
  $router: VueRouter

  /**
   * Create link
   * @param creatingUrl updated link
   */
  createLink (creatingUrl: string): Promise<any> {
    const url = new URL('link', this.apiBaseUrl)
    const reqBody = {
      url: creatingUrl
    }
    const req = new Request(
      url.href,
      {
        method: 'POST',
        body: JSON.stringify(reqBody),
        headers: {
          authorization: 'Bearer ' + TokenStore.accessToken
        }
      }
    )

    return fetch(req)
      .then(res => res.json())
      .then(resJson => this.interceptInstall(resJson))
      .then(resJson => this.interceptAuthorize(resJson, () => this.getLinks()))
      .then(resJson => ErrorParser().interceptErrorResp(resJson))
      .then(resJson => this.interceptInvalidType(resJson, 'create_link'))
  }

  /**
   * Get all links
   */
  getLinks (): Promise<any> {
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
      .then(resJson => ErrorParser().interceptErrorResp(resJson))
      .then(resJson => this.interceptInvalidType(resJson, 'get_link'))
  }

  /**
   * Get specific link
   * @param shortId short ID of link
   */
  getLinkByShortID (shortId: string): Promise<any> {
    const url = new URL('link/shortened', this.apiBaseUrl)
    url.searchParams.append('short_id', shortId)

    const req = new Request(
      url.href,
      {
        method: 'GET'
      }
    )

    return fetch(req)
      .then(res => res.json())
      .then(resJson => this.interceptInstall(resJson))
      .then(resJson => ErrorParser().interceptErrorResp(resJson))
      .then(resJson => this.interceptInvalidType(resJson, 'get_link_by_short_id'))
  }

  /**
   * Update link
   * @param id updating ID
   * @param updatingUrl updated link
   */
  updateLink (id: number, updatingUrl: string): Promise<any> {
    const url = new URL('link', this.apiBaseUrl)
    const reqBody = {
      id: id,
      url: updatingUrl
    }
    const req = new Request(
      url.href,
      {
        method: 'PATCH',
        body: JSON.stringify(reqBody),
        headers: {
          authorization: 'Bearer ' + TokenStore.accessToken
        }
      }
    )

    return fetch(req)
      .then(res => res.json())
      .then(resJson => this.interceptInstall(resJson))
      .then(resJson => this.interceptAuthorize(resJson, () => this.getLinks()))
      .then(resJson => ErrorParser().interceptErrorResp(resJson))
      .then(resJson => this.interceptInvalidType(resJson, 'update_link'))
  }

  /**
   * Delete link
   * @param id deleting ID
   */
  deleteLink (id: number): Promise<any> {
    const url = new URL('link', this.apiBaseUrl)
    const reqBody = {
      id: id
    }
    const req = new Request(
      url.href,
      {
        method: 'DELETE',
        body: JSON.stringify(reqBody),
        headers: {
          authorization: 'Bearer ' + TokenStore.accessToken
        }
      }
    )

    return fetch(req)
      .then(res => res.json())
      .then(resJson => this.interceptInstall(resJson))
      .then(resJson => this.interceptAuthorize(resJson, () => this.getLinks()))
      .then(resJson => ErrorParser().interceptErrorResp(resJson))
      .then(resJson => this.interceptInvalidType(resJson, 'delete_link'))
  }

  /**
   * Create admin
   * @param creatingName creating admin name
   * @param creatingScreenName creating admin screen name
   * @param creatingPassword creating admin password
   */
  createAdmin (
    creatingName: string,
    creatingScreenName: string,
    creatingPassword: string
  ): Promise<any> {
    const url = new URL('admin', this.apiBaseUrl)
    const reqBody = {
      name: creatingName,
      screen_name: creatingScreenName,
      password: creatingPassword
    }
    const req = new Request(
      url.href,
      {
        method: 'POST',
        body: JSON.stringify(reqBody),
        headers: {
          authorization: 'Bearer ' + TokenStore.accessToken
        }
      }
    )

    return fetch(req)
      .then(res => res.json())
      .then(resJson => this.interceptInstall(resJson))
      .then(resJson => this.interceptAuthorize(resJson, () => this.getAdmins()))
      .then(resJson => ErrorParser().interceptErrorResp(resJson))
      .then(resJson => this.interceptInvalidType(resJson, 'create_admin'))
  }

  /**
   * Get all admins
   */
  getAdmins (): Promise<any> {
    const url = new URL('admin', this.apiBaseUrl)
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
      .then(resJson => this.interceptAuthorize(resJson, () => this.getAdmins()))
      .then(resJson => ErrorParser().interceptErrorResp(resJson))
      .then(resJson => this.interceptInvalidType(resJson, 'get_admin'))
  }

  /**
   * Delete admin
   * @param id deleting ID
   */
  deleteAdmin (id: number): Promise<any> {
    const url = new URL('admin', this.apiBaseUrl)
    const reqBody = {
      id: id
    }
    const req = new Request(
      url.href,
      {
        method: 'DELETE',
        body: JSON.stringify(reqBody),
        headers: {
          authorization: 'Bearer ' + TokenStore.accessToken
        }
      }
    )

    return fetch(req)
      .then(res => res.json())
      .then(resJson => this.interceptInstall(resJson))
      .then(resJson => this.interceptAuthorize(resJson, () => this.getAdmins()))
      .then(resJson => ErrorParser().interceptErrorResp(resJson))
      .then(resJson => this.interceptInvalidType(resJson, 'delete_admin'))
  }

  /**
   * Register first administrator to set up
   * @param name Name of administrator
   * @param screenName Screen name of administrator
   * @param password Password of administrator
   */
  install (name: string, screenName: string, password: string): Promise<any> {
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
      .then(resJson => ErrorParser().interceptErrorResp(resJson))
      .then(resJson => this.interceptInvalidType(resJson, 'install'))
  }

  /**
   * Get is logged in or not
   */
  isInstalled (): Promise<any> {
    const url = new URL('install', this.apiBaseUrl)
    const req = new Request(
      url.href,
      {
        method: 'GET'
      }
    )

    return fetch(req)
      .then(res => res.json())
      .then(resJson => this.interceptInstall(resJson))
      .then(resJson => ErrorParser().interceptErrorResp(resJson))
      .then(resJson => this.interceptInvalidType(resJson, 'is_installed'))
  }

  /**
   * Login as administrator
   * @param name Name of administrator
   * @param password Password of administrator
   */
  login (name: string, password: string): Promise<any> {
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
      .then(resJson => ErrorParser().interceptErrorResp(resJson))
      .then(resJson => this.interceptInvalidType(resJson, 'login_admin'))
  }

  /**
   * Get is logged in or not
   */
  isLogin (): Promise<any> {
    const url = new URL('login', this.apiBaseUrl)
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
      .then(resJson => this.interceptAuthorize(resJson, () => this.isLogin()))
      .then(resJson => ErrorParser().interceptErrorResp(resJson))
      .then(resJson => this.interceptInvalidType(resJson, 'is_login'))
  }

  /**
   * Renew access / refresh token pair
   * @return Promise<response JSON>
   */
  renewToken (): Promise<any> {
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
      .then(resJson => ErrorParser().interceptErrorResp(resJson))
      .then(resJson => this.interceptInvalidType(resJson, 'refresh_token_admin'))
      .then((resJson) => {
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
    if (resJson?.type !== 'needed_install') {
      return resJson
    }

    return this.$router.push('/install')
  }

  /**
   * Purify installed error before processing normal scenario
   * @param resJson JSON on intercepting response
   * @return Promise<response JSON>
   */
  interceptInstalled (resJson: any): Promise<any> {
    if (resJson?.type !== 'redundant_install') {
      return resJson
    }

    return this.$router.push('/')
  }

  /**
   * Purify authorize error before processing normal scenario
   * @param resJson JSON on intercepting response
   * @param retryFunc If failed by expired tokens, run this handler after renewing it
   * @return Promise<response JSON> or Promise<Route>
   */
  interceptAuthorize (resJson: any, retryFunc: () => Promise<any>): Promise<any> {
    if (resJson?.type !== 'error') {
      return resJson
    }

    // Need login
    if (
      (resJson?.p_code === 302 && resJson?.s_code === 102) || // No auth header
      (resJson?.p_code === 202 && resJson?.s_code === 101) // No such user
    ) {
      return this.$router.push('/login')
    }

    // Need renew access token
    if (resJson?.p_code === 202 && resJson?.s_code === 103) {
      return this.renewToken()
        .then(resJson => ErrorParser().interceptErrorResp(resJson))
        .then(retryFunc)
        .catch(
          () => this.$router.push('/login')
        )
    }

    return resJson
  }

  /**
   * Purify invalid response type before processing normal scenario
   * @param resJson JSON on intercepting response
   * @param resType expected type in response JSON
   * @return Promise<response JSON>
   */
  interceptInvalidType (resJson: any, resType: string): Promise<any> {
    const type = resJson?.type
    if (type !== resType) {
      throw new Error(`Return value is not expected. type: "${type}"`)
    }

    return resJson
  }
}

/**
 * Get IncipitApi instance as singleton
 * @param $router VueRouter
 */
export default ($router: VueRouter) => IncipitApi.getInstance($router)
