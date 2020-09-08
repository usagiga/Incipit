<template>
  <div>
    <v-row>
      <v-col>
        <h1>
          Login
        </h1>
        <v-form
          ref="loginForm"
          v-model="isValidLogin"
        >
          <v-text-field
            v-model="userName"
            label="User Name"
            :rules="userNameRules"
            outlined
            required
            clearable
          />
          <v-text-field
            v-model="password"
            label="Password"
            type="password"
            :rules="passwordRules"
            outlined
            required
            clearable
          />
          <v-btn :disabled="!isValidLogin" @click="submit">
            Login
          </v-btn>
        </v-form>
      </v-col>
    </v-row>
  </div>
</template>

<script lang="ts">
/* eslint-disable no-console,camelcase */

import { Vue, Component } from 'nuxt-property-decorator'
import { VForm } from '~/types/v-form'
import IncipitApi from '~/utils/incipit-api'
import TokenStore from '~/utils/token-store'

  @Component({
    layout: 'blank'
  })
export default class Login extends Vue {
    isValidLogin: boolean = false

    userName: string = ''
    userNameRules: ((v: any) => boolean | string)[] = [
      v => !!v || 'Name is required'
    ]

    password: string = ''
    passwordRules: ((v: any) => boolean | string)[] = [
      v => !!v || 'Password is required'
    ]

    get loginForm (): VForm {
      return this.$refs.loginForm as any
    }

    submit (): void {
      // Validation
      this.loginForm.validate()
      if (!this.isValidLogin) {
        return
      }

      // Login
      IncipitApi(this.$router)
        .login(this.userName, this.password)
        .then((resJson: any) => {
          // Set token pair
          TokenStore.accessToken = resJson?.access_token?.token
          TokenStore.refreshToken = resJson?.refresh_token?.token

          // Redirect
          this.$router.push('/link')
        })
        .catch((err: any) => {
          console.error(err)
        })
    }

    head () {
      return {
        title: 'Login'
      }
    }
}
</script>
